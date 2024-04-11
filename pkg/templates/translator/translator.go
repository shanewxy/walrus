package translator

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/seal-io/utils/json"

	"github.com/seal-io/walrus/pkg/templates/translators/options"
	"github.com/seal-io/walrus/pkg/templates/translators/terraform"
)

// Translator translates between original template language and go types with openapi schema.
type Translator interface {
	// SchemaOfOriginalType generates openAPI schema from original type.
	SchemaOfOriginalType(typ any, opts options.Options) *openapi3.Schema
	// ToGoTypeValues converts values to go types.
	ToGoTypeValues(values map[string]json.RawMessage, schema openapi3.Schema) (map[string]any, error)
}

// SchemaOfType generates openAPI schema from original type.
func SchemaOfType(translatorName string, typ any, opts options.Options) (schema openapi3.Schema) {
	var (
		s  *openapi3.Schema
		tr Translator
	)

	switch translatorName {
	default:
		panic(fmt.Sprintf("unsupport translator %s", translatorName))
	case terraform.Name:
		tr = terraform.New()
	}

	s = tr.SchemaOfOriginalType(typ, opts)
	if s != nil {
		return *s
	}

	// Default unknown type.
	s = openapi3.NewSchema().
		WithDefault(opts.DefaultValue)
	s.Title = opts.Name
	s.Description = opts.Description
	s.WriteOnly = opts.Sensitive

	return *s
}

// ToGoTypeValues converts values to go types.
func ToGoTypeValues(values map[string]json.RawMessage, schema openapi3.Schema) (r map[string]any, err error) {
	// Terraform.
	tf := terraform.New()

	r, err = tf.ToGoTypeValues(values, schema)
	if err != nil {
		return nil, err
	}

	if r != nil {
		return r, nil
	}

	// Continue with other translator in the future.

	// No translator found.
	return nil, fmt.Errorf("no supported translator found for convert %v to go type", values)
}
