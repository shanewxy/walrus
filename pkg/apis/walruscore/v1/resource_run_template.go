package v1

import (
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ResourceRunTemplate is the schema for the resource run templates API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type ResourceRunTemplate struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceRunTemplateSpec   `json:"spec,omitempty"`
	Status ResourceRunTemplateStatus `json:"status,omitempty"`
}

type ResourceRunTemplateReference struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

var _ runtime.Object = (*ResourceRunTemplate)(nil)

// ResourceRunTemplateSpec defines the desired state of ResourceRunTemplate.
// It is a template for a resource run, which defines the previous and post step for resource operation like plan and apply.
// The order of the steps will be:
//
// template pre-plan--> resource pre-plan --> plan --> resource post-plan
// --> template post-plan --> template pre-apply --> resource pre-apply --> apply
// --> resource post-apply --> template post-apply.
type ResourceRunTemplateSpec struct {
	// Plan defines the template steps to run in before and after the resource run template.
	Plan ResourceOperationHook `json:"plan,omitempty"`

	// Apply defines the template steps to apply in the resource run template.
	Apply ResourceOperationHook `json:"apply,omitempty"`
}

// ResourceRunTemplateStatus defines the observed state of ResourceRunTemplate.
type ResourceRunTemplateStatus struct{}

// ResourceRunTemplateList holds the list of ResourceRunTemplate.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourceRunTemplateList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ResourceRunTemplate `json:"items"`
}

var _ runtime.Object = (*ResourceRunTemplateList)(nil)

// ResourceRunTemplateStep defines the desired state of ResourceRunTemplateStep.
type ResourceRunTemplateStep struct {
	// Name is the name of the resource run template step.
	Name string `json:"name"`

	// Description is the description of the resource run template step.
	Description string `json:"description,omitempty"`

	// ResourceRunStepTemplate is the reference to the resource run step template.
	ResourceRunStepTemplate *ResourceRunTemplateReference `json:"resourceRunStepTemplate,omitempty"`

	// Container is the main container image to run in the resource run template step.
	// It is used when ResourceRunStepTemplate is not provided.
	Container *core.Container `json:"container,omitempty"`
}
