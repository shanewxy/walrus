package v1

import (
	"k8s.io/apimachinery/pkg/runtime"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

// Schema is the schema for the schemas API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"]
type Schema walruscore.Schema

var _ runtime.Object = (*Schema)(nil)

// SchemaList holds the list of Schema.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SchemaList walruscore.SchemaList

var _ runtime.Object = (*SchemaList)(nil)
