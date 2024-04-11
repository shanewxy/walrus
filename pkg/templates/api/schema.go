package api

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/seal-io/utils/json"
	"k8s.io/klog/v2"

	"github.com/seal-io/walrus/pkg/templates/openapi"
)

const (
	VariableSchemaKey = "variables"
	OutputSchemaKey   = "outputs"
)

type SchemaGroup struct {
	Template *TemplateSchema `json:"template,omitempty"`
	UI       *UISchema       `json:"ui,omitempty"`
}

// UISchema include the UI schema that users can customize.
type UISchema = WrapSchema

// TemplateSchema include the internal template variables schema and template data.
type TemplateSchema struct {
	WrapSchema `json:",inline"`

	// DefaultValue specifies the default value of the schema.
	DefaultValue []byte `json:"defaultValue,omitempty"`

	// Metadata specifies the metadata of this template.
	Data *TemplateData `json:"data,omitempty"`
}

// Validate reports if the schema is valid.
func (s *TemplateSchema) Validate() error {
	return s.WrapSchema.Validate()
}

type TemplateData struct {
	// Readme specifies the readme of this template.
	Readme string `json:"readme,omitempty"`

	Terraform *TerraformMetadata `json:"terraform,omitempty"`
}

// TerraformMetadata include the terraform metadata of this template version.
type TerraformMetadata struct {
	// RequiredProviders specifies the required providers of this template.
	RequiredProviders []ProviderRequirement `json:"requiredProviders,omitempty"`
}

// ProviderRequirement include the required provider.
type ProviderRequirement struct {
	*tfconfig.ProviderRequirement

	Name string `json:"name,omitempty"`
}

// WrapSchema wrap the openAPI schema with variables and outputs.
type WrapSchema struct {
	Schema *openapi3.T `json:"schema"`
}

// Validate reports if the schema is valid.
func (s *WrapSchema) Validate() error {
	if s.Schema == nil {
		return nil
	}

	// workaround: inject paths and version since kin-openapi/openapi3 need it.
	s.Schema.Paths = &openapi3.Paths{}
	if s.Schema.Info != nil && s.Schema.Info.Version == "" {
		s.Schema.Info.Version = "v0.0.0"
	}

	if err := s.Schema.Validate(
		context.Background(),
		openapi3.DisableSchemaDefaultsValidation(),
	); err != nil {
		return err
	}

	return nil
}

func (s *WrapSchema) IsEmpty() bool {
	return s.Schema == nil ||
		s.Schema.Components == nil ||
		len(s.Schema.Components.Schemas) == 0
}

// Expose returns the UI schema of the schema.
func (s *WrapSchema) Expose(skipProps ...string) WrapSchema {
	vs := s.VariableSchema()
	if vs == nil {
		return WrapSchema{}
	}

	// In order to prevent the remove ext affect the original schema, serialize and deserialize to copy the schema.
	b, err := json.Marshal(vs)
	if err != nil {
		klog.Warningf("error marshal variable schema while expost: %v", err)
		return WrapSchema{}
	}

	var cps openapi3.Schema

	err = json.Unmarshal(b, &cps)
	if err != nil {
		klog.Warningf("error unmarshal variable schema while expost: %v", err)
		return WrapSchema{}
	}

	for _, v := range skipProps {
		delete(cps.Properties, v)
	}

	return WrapSchema{
		Schema: &openapi3.T{
			OpenAPI: s.Schema.OpenAPI,
			Info:    s.Schema.Info,
			Components: &openapi3.Components{
				Schemas: map[string]*openapi3.SchemaRef{
					VariableSchemaKey: {
						Value: openapi.RemoveExtOriginal(&cps),
					},
				},
			},
		},
	}
}

// VariableSchema returns the variables' schema.
func (s *WrapSchema) VariableSchema() *openapi3.Schema {
	if s.Schema == nil ||
		s.Schema.Components == nil ||
		s.Schema.Components.Schemas == nil ||
		s.Schema.Components.Schemas[VariableSchemaKey] == nil ||
		s.Schema.Components.Schemas[VariableSchemaKey].Value == nil {
		return nil
	}

	return s.Schema.Components.Schemas[VariableSchemaKey].Value
}

func (s *WrapSchema) SetVariableSchema(v *openapi3.Schema) {
	s.ensureInit()
	s.Schema.Components.Schemas[VariableSchemaKey].Value = v
}

func (s *WrapSchema) RemoveVariableContext() {
	if s.IsEmpty() {
		return
	}

	variableSchema := openapi.RemoveVariableContext(s.VariableSchema())
	s.SetVariableSchema(variableSchema)
}

func (s *WrapSchema) SetOutputSchema(v *openapi3.Schema) {
	s.ensureInit()
	s.Schema.Components.Schemas[OutputSchemaKey].Value = v
}

// OutputSchema returns the outputs' schema.
func (s *WrapSchema) OutputSchema() *openapi3.Schema {
	if s.Schema == nil ||
		s.Schema.Components == nil ||
		s.Schema.Components.Schemas == nil ||
		s.Schema.Components.Schemas[OutputSchemaKey] == nil ||
		s.Schema.Components.Schemas[OutputSchemaKey].Value == nil {
		return nil
	}

	return s.Schema.Components.Schemas[OutputSchemaKey].Value
}

// Intersect sets variables & outputs schema of s to intersection of s and s2.
func (s *WrapSchema) Intersect(s2 *WrapSchema) {
	if s2.Schema == nil {
		return
	}

	variableSchema := openapi.IntersectSchema(s.VariableSchema(), s2.VariableSchema())
	s.SetVariableSchema(variableSchema)
	outputSchema := openapi.IntersectSchema(s.OutputSchema(), s2.OutputSchema())
	s.SetOutputSchema(outputSchema)
}

func (s *WrapSchema) ensureInit() {
	if s.Schema == nil {
		s.Schema = &openapi3.T{}
	}

	if s.Schema.Components == nil {
		s.Schema.Components = &openapi3.Components{}
	}

	if s.Schema.Components.Schemas == nil {
		s.Schema.Components.Schemas = openapi3.Schemas{}
	}
}
