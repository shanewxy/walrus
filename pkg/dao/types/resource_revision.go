package types

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/seal-io/walrus/pkg/dao/types/property"
)

type OutputValue struct {
	Name   string          `json:"name,omitempty"`
	Value  property.Value  `json:"value,omitempty"`
	Schema openapi3.Schema `json:"schema,omitempty"`
}

const (
	ResourceActionTypeCreate   string = "create"
	ResourceActionTypeDelete   string = "delete"
	ResourceActionTypeUpgrade  string = "upgrade"
	ResourceActionTypeRollback string = "rollback"
	ResourceActionTypeStop     string = "stop"
	ResourceActionTypeStart    string = "start"
)
