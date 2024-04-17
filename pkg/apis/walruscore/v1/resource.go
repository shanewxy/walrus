package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Resource is the schema for the resources API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type Resource struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceSpec   `json:"spec,omitempty"`
	Status ResourceStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Resource)(nil)

// ResourceSpec defines the desired state of Resource.
type ResourceSpec struct {
	// Attributes to configure the template.
	Attributes runtime.RawExtension `json:"attributes,omitempty"`

	// TemplateVersion template version to which the resource belongs.
	TemplateVersion *TempalteVersionReference `json:"templateVersionReference,omitempty"`

	// Type is a resource definition type.
	Type string `json:"type,omitempty"`

	// Draft indicates whether the resource is a draft.
	Draft bool `json:"draft,omitempty"`

	// Stop indicates whether to stop the resource.
	Stop bool `json:"stop,omitempty"`
}

// ResourceStatus defines the observed state of Resource.
type ResourceStatus struct {
	// StatusDescriptor defines the status of the resource.
	StatusDescriptor `json:",inline"`

	// Project is the project to which the resource belongs.
	Project string `json:"project"`

	// ComputedAttributes generated from attributes and schemas.
	ComputedAttributes runtime.RawExtension `json:"computedAttributes"`

	// Dependencies of the resource.
	Dependencies []string `json:"dependencies,omitempty"`

	// ResourceHook is a reference to a resource hook.
	// For one resource, the resource hook is unique.
	ResourceHook *ResourceHookReference `json:"resourceHook"`

	// ResourceDefinition is a reference to a resource definition.
	ResourceDefinition *ResourceDefinitionReference `json:"resourceDefinition,omitempty"`

	// ResourceDefinitionMatchingRule is a reference to a resource definition matching rule.
	ResourceDefinitionMatchingRule *ResourceDefinitionMatchingRuleReference `json:"resourceDefinitionMatchingRule,omitempty"`

	// Endpoints of the resource.
	Endpoints []string `json:"endpoints,omitempty"`
}

// ResourceList holds the list of Resource.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourceList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Resource `json:"items"`
}

var _ runtime.Object = (*ResourceList)(nil)
