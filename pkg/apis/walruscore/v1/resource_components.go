package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ResourceComponents is the schema for the resource components API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type ResourceComponents struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceComponentsSpec   `json:"spec,omitempty"`
	Status ResourceComponentsStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*ResourceComponents)(nil)

// ResourceComponentsSpec defines the desired state of ResourceComponents.
type ResourceComponentsSpec struct{}

// ResourceComponentsStatus defines the observed state of ResourceComponents.
type ResourceComponentsStatus struct {
	// StatusDescriptor defines the status of the resource components.
	StatusDescriptor `json:",inline"`

	// Project is the project of the resource components.
	Project string `json:"project"`

	// ResourceName is the resource name of the resource components.
	ResourceName string `json:"resource"`

	// TemplateVersion template version to which is used to create the resource components.
	TemplateVersion *TempalteVersionReference `json:"templateVersionReference"`

	// ComputedAttributes stores the computed attributes of the component.
	// It stores the attributes of the resource that used to create the component.
	ComputedAttributes runtime.RawExtension `json:"computedAttributes"`

	// Components store the components of the resource components.
	Components []ResourceComponent `json:"components"`

	// Dependencies store the dependencies of the resource components.
	Dependencies []ResourceComponentDependency `json:"dependencies"`
}

// ResourceComponentsList holds the list of ResourceComponents.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourceComponentsList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ResourceComponents `json:"items"`
}

// ResourceComponentAddress is the address of the resource component.
type ResourceComponentAddress = string

// ResourceComponentMode is the mode of the resource component.
// +enum
type ResourceComponentMode string

const (
	// ResourceComponentModeManaged indicates the resource created to target platform,
	// it is writable(update or delete).
	ResourceComponentModeManaged ResourceComponentMode = "managed"

	// ResourceComponentModeData indicates the resource read from target platform,
	// it is read-only.
	ResourceComponentModeData ResourceComponentMode = "data"
)

type ResourceComponent struct {
	// Name is the name of the component.
	Name string `json:"name"`

	// Type is the type of the component, example: "aws_instance" for aws_instance.foo.
	Type string `json:"type"`

	// Mode is the mode of the component.
	Mode ResourceComponentMode `json:"mode"`

	// Shape is the shape of the component.
	Shape ResourceComponentShape `json:"shape"`

	// Address is the address of the component.
	Address ResourceComponentAddress `json:"address"`

	// Connector of the resource component.
	Connector *ConnectorReference `json:"connector"`
}

// ResourceComponentShape is the shape of the resource component.
// +enum
type ResourceComponentShape string

const (
	// ResourceComponentShapeClass defines the resource component as class.
	ResourceComponentShapeClass ResourceComponentShape = "class"

	// ResourceComponentShapeInstance defines the resource component as instance.
	ResourceComponentShapeInstance ResourceComponentShape = "instance"

	// ResourceComponentShapeComposition defines the resource component as composition.
	ResourceComponentShapeComposition ResourceComponentShape = "composition"
)

type ResourceComponentDependency struct {
	From ResourceComponentAddress `json:"from"`
	To   ResourceComponentAddress `json:"to"`
}

// ResourceComponentAppendix stores the appendix of the resource component.
type ResourceComponentAppendix struct {
	// OperationKeys stores the operation keys of the resource component.
	OperationKeys ResourceComponentOperationKeys `json:"operationKeys,omitempty"`
}

// ResourceComponentOperationKeys stores the operation keys of the resource component.
type ResourceComponentOperationKeys struct {
	// Labels stores label of layer,
	// its length means each key contains levels with the same value as level.
	Labels []string `json:"labels,omitempty"`
	// Keys stores key in tree.
	Keys []ResourceComponentOperationKey `json:"keys,omitempty"`
}

// ResourceComponentOperationKey holds hierarchy query keys.
type ResourceComponentOperationKey struct {
	// Name indicates the name of the key.
	Name string `json:"name"`
	// Value indicates the value of the key;
	// usually, it should be valued in leaves.
	Value string `json:"value,omitempty"`
	// Loggable indicates whether to be able to get log.
	Loggable *bool `json:"loggable,omitempty"`
	// Executable indicates whether to be able to execute remote command.
	Executable *bool `json:"executable,omitempty"`
}
