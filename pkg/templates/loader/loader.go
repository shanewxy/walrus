package loader

import (
	"fmt"

	"github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/templates/api"
	"github.com/seal-io/walrus/pkg/templates/loaders/terraform"
)

// SchemaLoader define the interface for loading schema from template.
type SchemaLoader interface {
	Load(rootDir, templateName string) (*api.SchemaGroup, error)
}

// LoadSchema loads schema from template.
func LoadSchema(rootDir, templateName, templateFormat string) (s *api.SchemaGroup, err error) {
	// Terraform.
	switch templateFormat {
	default:
		return nil, fmt.Errorf("unsupport template format %s", templateFormat)
	case v1.TemplateFormatTerraform:
		tf := terraform.New()
		return tf.Load(rootDir, templateName)
	}
}
