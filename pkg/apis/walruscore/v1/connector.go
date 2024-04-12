package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Connector is the schema for the connectors API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type Connector struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectorSpec   `json:"spec"`
	Status ConnectorStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Connector)(nil)

// ConnectorSpec defines the desired state of Connector.
type ConnectorSpec struct {
	// ApplicableEnvironmentType is the environment type that the connector is applicable to.
	//
	// +k8s:validation:enum=["Development","Staging","Production"]
	ApplicableEnvironmentType EnvironmentType `json:"applicableEnvironmentType,omitempty"`

	// Category is the category of the connector.
	//
	// +k8s:validation:enum=["Docker","Kubernetes","Custom","CloudProvider"]
	Category ConnectorCategory `json:"category"`

	// Type is the type of the connector.
	Type string `json:"type"`

	// Config is the configuration of the connector.
	Config ConnectorConfig `json:"config"`

	// Description is the description of the connector.
	Description string `json:"description,omitempty"`

	// SecretName is the auto-generated secret name for the connector configuration. Will be overridden if set.
	SecretName string `json:"secretName,omitempty"`
}

type ConnectorConfig struct {
	Version string                          `json:"version"`
	Data    map[string]ConnectorConfigEntry `json:"data"`
}

type ConnectorConfigEntry struct {
	Value   string `json:"value"`
	Visible bool   `json:"visible"`
}

// ConnectorStatus defines the observed state of Connector.
type ConnectorStatus struct {
	// StatusDescriptor defines the status of the Connector.
	StatusDescriptor `json:",inline"`

	// Project is the project that the connector belongs to.
	Project string `json:"project,omitempty"`
}

// ConnectorList holds the list of Connector.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ConnectorList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Connector `json:"items"`
}

var _ runtime.Object = (*ConnectorList)(nil)

// ConnectorCategory is the category of the connector.
//
// +enum
type ConnectorCategory string

const (
	ConnectorCategoryDocker        ConnectorCategory = "Docker"
	ConnectorCategoryKubernetes    ConnectorCategory = "Kubernetes"
	ConnectorCategoryCustom        ConnectorCategory = "Custom"
	ConnectorCategoryCloudProvider ConnectorCategory = "CloudProvider"
)

// ConnectorReference is a reference to a connector.
type ConnectorReference struct {
	// Name is the name of the connector.
	Name string `json:"name"`

	// Namespace is the namespace of the connector.
	Namespace string `json:"namespace"`
}
