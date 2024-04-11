package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Schema API for the template's version.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced"
type Schema struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Status SchemaStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Schema)(nil)

// SchemaStatus defines the template version's schema.
type SchemaStatus struct {
	// Value is the current value of the schema.
	Value runtime.RawExtension `json:"value"`

	// Project is the project that the catalog belongs to.
	Project string `json:"project,omitempty"`

	// Conditions holds the conditions for the schema.
	Conditions []Condition `json:"conditions,omitempty"`
}

// SchemaList holds the list of Schema.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SchemaList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Schema `json:"items"`
}

var _ runtime.Object = (*SchemaList)(nil)

const (
	NameSuffixTemplateSchema   = "template-schema"
	NameSuffixUISchema         = "ui-schema"
	NameSuffixOriginalUISchema = "original-ui-schema"
)
