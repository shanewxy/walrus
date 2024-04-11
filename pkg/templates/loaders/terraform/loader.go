package terraform

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/ext/typeexpr"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/seal-io/utils/json"
	"github.com/seal-io/utils/osx"
	"github.com/seal-io/utils/pools/gopool"
	"github.com/seal-io/utils/stringx"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"

	"github.com/seal-io/walrus/pkg/templates/api"
	"github.com/seal-io/walrus/pkg/templates/openapi"
	"github.com/seal-io/walrus/pkg/templates/translator"
	"github.com/seal-io/walrus/pkg/templates/translators/options"
	"github.com/seal-io/walrus/pkg/templates/translators/terraform"
)

const (
	defaultGroup = "Basic"
)

// Loader for load terraform template schema and data.
type Loader struct {
	translator translator.Translator
}

// New creates a new terraform loader.
func New() *Loader {
	return &Loader{
		translator: terraform.New(),
	}
}

// Load loads the internal template version schema and data.
func (l *Loader) Load(
	rootDir, templateName string,
) (*api.SchemaGroup, error) {
	mod, err := l.loadMod(rootDir)
	if err != nil {
		return nil, err
	}

	ts, fs, err := l.loadSchema(rootDir, mod, templateName)
	if err != nil {
		return nil, err
	}

	return &api.SchemaGroup{
		Template: ts,
		UI:       fs,
	}, nil
}

// loadMod load the terraform module.
func (l *Loader) loadMod(rootDir string) (*tfconfig.Module, error) {
	if !tfconfig.IsModuleDir(rootDir) {
		return nil, fmt.Errorf("no terraform configuration files found")
	}

	mod, diags := tfconfig.LoadModule(rootDir)
	if diags.HasErrors() {
		return nil, diags.Err()
	}

	return mod, nil
}

// loadSchema loads the internal template version schema.
func (l *Loader) loadSchema(
	rootDir string,
	mod *tfconfig.Module,
	template string,
) (*api.TemplateSchema, *api.UISchema, error) {
	var (
		s  = &api.TemplateSchema{}
		us = api.UISchema{}
	)

	var (
		ts  *openapi3.T
		err error
	)

	// TemplateSchema.
	{
		// OpenAPISchema.
		ts, err = l.getSchemaFromTerraform(mod, template)
		if err != nil {
			return nil, nil, err
		}
		s.WrapSchema = api.WrapSchema{
			Schema: ts,
		}

		// TemplateSchema Default.
		tsDefault, err := openapi.GenSchemaDefaultPatch(context.Background(), s.WrapSchema.VariableSchema())
		if err != nil {
			return nil, nil, err
		}
		s.DefaultValue = tsDefault

		// Data.
		data, err := l.loadData(rootDir, mod)
		if err != nil {
			return nil, nil, err
		}
		s.Data = data
	}

	// UI TemplateSchema.
	{
		fs, err := l.getSchemaFromFile(rootDir, ts)
		if err != nil {
			return nil, nil, err
		}

		wfs := api.TemplateSchema{
			WrapSchema: api.WrapSchema{
				Schema: fs,
			},
		}

		us = s.Expose(openapi.WalrusContextVariableName)
		if fs != nil && !wfs.IsEmpty() {
			us = wfs.Expose()
		}
	}

	return s, &us, nil
}

func (l *Loader) getSchemaFromTerraform(mod *tfconfig.Module, template string) (*openapi3.T, error) {
	varsSchema, err := l.getVariableSchemaFromTerraform(mod)
	if err != nil {
		return nil, err
	}

	outputsSchema, err := l.getOutputSchemaFromTerraform(mod)
	if err != nil {
		return nil, err
	}

	// OpenAPI OpenAPISchema.
	t := &openapi3.T{
		OpenAPI: openapi.OpenAPIVersion,
		Info: &openapi3.Info{
			Title: fmt.Sprintf("OpenAPI schema for template %s", template),
		},
		Components: &openapi3.Components{
			Schemas: openapi3.Schemas{},
		},
	}

	if varsSchema != nil {
		t.Components.Schemas["variables"] = varsSchema.NewRef()
	}

	if outputsSchema != nil {
		t.Components.Schemas["outputs"] = outputsSchema.NewRef()
	}

	return t, nil
}

// getSchemaFromFile get openapi schema from schema.yaml.
func (l *Loader) getSchemaFromFile(rootDir string, originalSchema *openapi3.T) (*openapi3.T, error) {
	schemaFile := filepath.Join(rootDir, "schema.yaml")
	if !osx.Exists(schemaFile) {
		if schemaFile = filepath.Join(rootDir, "schema.yml"); !osx.Exists(schemaFile) {
			return nil, nil
		}
	}

	content, err := os.ReadFile(schemaFile)
	if err != nil {
		return nil, fmt.Errorf("error reading schema file %s: %w", schemaFile, err)
	}

	if len(content) == 0 {
		return nil, nil
	}

	// Openapi loader will cache the data with file path as key if we use LoadFromFile,
	// since the repo with different tag the schema.yaml file is the same, so we use LoadFromData to skip the cache.
	it, err := openapi3.NewLoader().LoadFromData(content)
	if err != nil {
		return nil, fmt.Errorf("error loading schema file %s: %w", schemaFile, err)
	}

	if it.Components == nil ||
		it.Components.Schemas == nil ||
		it.Components.Schemas["variables"] == nil ||
		it.Components.Schemas["variables"].Value == nil {
		return nil, nil
	}

	// Inject.
	var (
		varsSchema         = it.Components.Schemas["variables"].Value
		originalVarsSchema *openapi3.Schema
	)

	if originalSchema.Components != nil && originalSchema.Components.Schemas["variables"] != nil {
		originalVarsSchema = originalSchema.Components.Schemas["variables"].Value
	}

	varsSchema = l.applyMissingConfig(originalVarsSchema, varsSchema)
	l.injectExts(varsSchema)
	it.Components.Schemas["variables"].Value = varsSchema

	return it, nil
}

// applyMissingConfig apply the missing config to schema generate from schema.yaml.
func (l *Loader) applyMissingConfig(generated, customized *openapi3.Schema) *openapi3.Schema {
	if customized == nil {
		return nil
	}

	if generated == nil {
		return customized
	}

	s := *customized
	if len(s.Extensions) == 0 && len(generated.Extensions) != 0 {
		s.Extensions = generated.Extensions
	}

	for n, v := range s.Properties {
		in := generated.Properties[n]
		if in == nil || in.Value == nil {
			continue
		}

		// Title.
		if v.Value.Title == "" {
			s.Properties[n].Value.Title = generated.Properties[n].Value.Title
		}

		// Extensions.
		var (
			genExt = openapi.NewExtFromMap(in.Value.Extensions)
			ext    = openapi.NewExtFromMap(v.Value.Extensions)
		)

		ext.WithOriginal(in.Value.Extensions[openapi.ExtOriginalKey])

		if ext.ExtUI.Order == 0 {
			ext.WithUIOrder(genExt.Order)
		}

		if ext.ExtUI.ColSpan == 0 {
			ext.WithUIColSpan(genExt.ColSpan)
		}

		s.Properties[n].Value.Extensions = ext.Export()
	}

	return &s
}

// getVariableSchemaFromTerraform generate variable schemas from terraform files.
func (l *Loader) getVariableSchemaFromTerraform(mod *tfconfig.Module) (*openapi3.Schema, error) {
	if len(mod.Variables) == 0 {
		return nil, nil
	}

	var (
		varSchemas = openapi3.NewObjectSchema()
		required   []string
		keys       = make([]string, len(mod.Variables))
	)

	// Variables.
	for i, v := range sortVariables(mod.Variables) {
		// Parse tf expression from type.
		var (
			tfType       = cty.DynamicPseudoType
			defValue     = v.Default
			defObj       any
			order        = i + 1
			tyExpression any
		)

		// Required and keys.
		if v.Required {
			required = append(required, v.Name)
		}

		keys[i] = v.Name

		// Generate json schema from tf type or default value.
		if v.Type != "" {
			// Type exists.
			expr, diags := hclsyntax.ParseExpression(stringx.ToBytes(&v.Type), "", hcl.Pos{Line: 1, Column: 1})
			if diags.HasErrors() {
				return nil, fmt.Errorf("error parsing expression: %w", diags)
			}

			tfType, defObj, diags = typeexpr.TypeConstraintWithDefaults(expr)
			if diags.HasErrors() {
				return nil, fmt.Errorf("error getting type: %w", diags)
			}

			tyExpression = expr
		} else if v.Default != nil {
			// Empty type, use default value to get type.
			b, err := json.Marshal(v.Default)
			if err != nil {
				return nil, fmt.Errorf("error while marshal terraform variable default value: %w", err)
			}

			var sjv ctyjson.SimpleJSONValue

			err = sjv.UnmarshalJSON(b)
			if err != nil {
				return nil, fmt.Errorf("error while unmarshal terraform variable default value: %w", err)
			}
			tfType = sjv.Type()
		}

		varSchemas.WithProperty(
			v.Name,
			l.translator.SchemaOfOriginalType(
				tfType,
				options.Options{
					Name:          v.Name,
					DefaultValue:  defValue,
					DefaultObject: defObj,
					Description:   v.Description,
					Sensitive:     v.Sensitive,
					Order:         order,
					TypeExpress:   tyExpression,
				}))
	}

	// Inject extension sequence.
	sort.Strings(required)
	varSchemas.Required = required
	varSchemas.Extensions = openapi.NewExtFromMap(varSchemas.Extensions).
		WithOriginalVariablesSequence(keys).
		Export()

	// Inject extensions.
	l.injectExts(varSchemas)

	return varSchemas, nil
}

// getOutputSchemaFromTerraform generate output schemas from terraform files.
func (l *Loader) getOutputSchemaFromTerraform(mod *tfconfig.Module) (*openapi3.Schema, error) {
	if len(mod.Outputs) == 0 {
		return nil, nil
	}

	var (
		filenames     = sets.Set[string]{}
		outputSchemas = openapi3.NewObjectSchema()
	)

	for i, v := range sortOutput(mod.Outputs) {
		order := i + 1
		// Use dynamic type for output.
		outputSchemas.WithProperty(
			v.Name,
			l.translator.SchemaOfOriginalType(
				cty.DynamicPseudoType,
				options.Options{
					Name:        v.Name,
					Description: v.Description,
					Sensitive:   v.Sensitive,
					Order:       order,
				}))

		filenames.Insert(v.Pos.Filename)
	}

	values, err := getOutputValues(filenames)
	if err != nil {
		return nil, err
	}

	for n, v := range values {
		ext := outputSchemas.Properties[n].Value.Extensions
		outputSchemas.Properties[n].Value.Extensions = openapi.NewExtFromMap(ext).
			WithOriginalValueExpression(v).
			WithOriginalType(cty.DynamicPseudoType).
			Export()
	}

	outputSchemas.Extensions = openapi.NewExt().
		Export()

	return outputSchemas, nil
}

// getOutputValues gets the output values from output configuration files.
func getOutputValues(filenames sets.Set[string]) (map[string][]byte, error) {
	var (
		mu      sync.Mutex
		wg      = gopool.Group()
		outputs = make(map[string][]byte)
	)

	for _, filename := range filenames.UnsortedList() {
		wg.Go(func() error {
			b, err := os.ReadFile(filename)
			if err != nil {
				return fmt.Errorf("error read output configuration file %s: %w", filename, err)
			}

			var (
				file   *hcl.File
				diag   hcl.Diagnostics
				parser = hclparse.NewParser()
			)

			if strings.HasSuffix(filename, ".json") {
				file, diag = parser.ParseJSON(b, filename)
			} else {
				file, diag = parser.ParseHCL(b, filename)
			}

			if diag.HasErrors() {
				klog.Warningf("error parse output configuration file %s: %s", filename, diag.Error())
				return nil
			}

			if file == nil {
				return nil
			}

			o := getOutputValueFromFile(file)

			mu.Lock()
			for on, oe := range o {
				outputs[on] = oe
			}
			mu.Unlock()

			return nil
		})
	}

	if err := wg.Wait(); err != nil {
		return nil, err
	}

	return outputs, nil
}

// getOutputValueFromFile gets the output value from output configuration file.
func getOutputValueFromFile(file *hcl.File) map[string][]byte {
	var (
		rootSchema = &hcl.BodySchema{
			Blocks: []hcl.BlockHeaderSchema{
				{
					Type:       "output",
					LabelNames: []string{"name"},
				},
			},
		}
		outputSchema = &hcl.BodySchema{
			Attributes: []hcl.AttributeSchema{
				{
					Name: "value",
				},
			},
		}
	)

	var (
		outputs       = make(map[string][]byte)
		content, _, _ = file.Body.PartialContent(rootSchema)
	)

	for _, block := range content.Blocks {
		if block.Type == "output" {
			ct, _, _ := block.Body.PartialContent(outputSchema)
			name := block.Labels[0]

			if attr, defined := ct.Attributes["value"]; defined {
				outputs[name] = attr.Expr.Range().SliceBytes(file.Bytes)
			}
		}
	}

	return outputs
}

// loadData loads the internal template version data.
func (l *Loader) loadData(rootDir string, mod *tfconfig.Module) (
	*api.TemplateData, error,
) {
	// Readme.
	readme, err := l.getReadme(rootDir)
	if err != nil {
		return nil, err
	}

	// Providers.
	requiredProviders := l.getRequiredProviders(mod.RequiredProviders)
	if len(requiredProviders) == 0 && readme == "" {
		return nil, nil
	}

	data := &api.TemplateData{
		Readme: readme,
	}

	if len(requiredProviders) > 0 {
		data.Terraform = &api.TerraformMetadata{
			RequiredProviders: requiredProviders,
		}
	}
	return data, nil
}

// getReadme gets the readme content.
func (l *Loader) getReadme(rootDir string) (string, error) {
	path := filepath.Join(rootDir, "README.md")
	if osx.Exists(path) {
		content, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}

		return string(content), nil
	}

	return "", nil
}

// getRequiredProviders gets the required providers.
func (l *Loader) getRequiredProviders(
	m map[string]*tfconfig.ProviderRequirement,
) (s []api.ProviderRequirement) {
	if len(m) == 0 {
		return
	}

	for k, v := range m {
		s = append(s, api.ProviderRequirement{
			Name:                k,
			ProviderRequirement: v,
		})
	}

	sort.SliceStable(s, func(i, j int) bool {
		return s[i].Name < s[j].Name
	})

	return
}

// injectExts injects extension for variables.
func (l *Loader) injectExts(vs *openapi3.Schema) {
	if vs == nil {
		return
	}

	groupOrder := make(map[string]int)

	for n, v := range vs.Properties {
		if v.Value == nil || v.Value.IsEmpty() {
			continue
		}

		// Group.
		extUI := openapi.GetExtUI(v.Value.Extensions)
		if extUI.Group == "" {
			extUI.Group = defaultGroup

			vs.Properties[n].Value.Extensions = openapi.NewExtFromMap(vs.Properties[n].Value.Extensions).
				WithUIGroup(defaultGroup).
				Export()
		}

		od, ok := groupOrder[extUI.Group]
		if !ok || od > extUI.Order {
			groupOrder[extUI.Group] = extUI.Order
		}
	}

	vsExtUI := openapi.GetExtUI(vs.Extensions)
	if len(vsExtUI.GroupOrder) == 0 {
		ep := openapi.NewExtFromMap(vs.Extensions).
			WithUIGroupOrder(sortMapValue(groupOrder)...).
			Export()
		vs.Extensions = ep
	}
}

func sortMapValue(m map[string]int) []string {
	type keyValue struct {
		Key   string
		Value int
	}

	s := make([]keyValue, 0)

	for key, value := range m {
		s = append(s, keyValue{Key: key, Value: value})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Value < s[j].Value
	})

	keys := make([]string, len(s))
	for i, kv := range s {
		keys[i] = kv.Key
	}

	return keys
}

func sortVariables(m map[string]*tfconfig.Variable) (s []*tfconfig.Variable) {
	s = make([]*tfconfig.Variable, 0, len(m))
	for k := range m {
		s = append(s, m[k])
	}

	sort.SliceStable(s, func(i, j int) bool {
		return judgeSourcePos(&s[i].Pos, &s[j].Pos)
	})

	return
}

func sortOutput(m map[string]*tfconfig.Output) (s []*tfconfig.Output) {
	s = make([]*tfconfig.Output, 0, len(m))
	for k := range m {
		s = append(s, m[k])
	}

	sort.SliceStable(s, func(i, j int) bool {
		return judgeSourcePos(&s[i].Pos, &s[j].Pos)
	})

	return
}

func judgeSourcePos(i, j *tfconfig.SourcePos) bool {
	switch {
	case i.Filename < j.Filename:
		return false
	case i.Filename > j.Filename:
		return true
	}

	return i.Line < j.Line
}
