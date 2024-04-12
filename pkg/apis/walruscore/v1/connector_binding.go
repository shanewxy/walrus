package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const ProviderLabelPrefix = "walrus.seal.io/provider-"

// ConnectorBinding is the schema for the connectorbindings API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type ConnectorBinding struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectorBindingSpec   `json:"spec"`
	Status ConnectorBindingStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*ConnectorBinding)(nil)

// ConnectorBindingSpec defines the desired state of ConnectorBinding.
type ConnectorBindingSpec struct {
	// Connector is the reference to the connector.
	Connector ConnectorReference `json:"connector"`
}

// ConnectorBindingStatus defines the observed state of ConnectorBinding.
type ConnectorBindingStatus struct {
	// Type is the type of the connector.
	Type string `json:"Type"`

	// Category is the category of the connector.
	Category ConnectorCategory `json:"Category"`
}

// ConnectorBindingList contains a list of ConnectorBinding.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ConnectorBindingList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ConnectorBinding `json:"items"`
}

var _ runtime.Object = (*ConnectorList)(nil)
