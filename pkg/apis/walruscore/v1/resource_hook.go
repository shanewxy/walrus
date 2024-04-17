package v1

import (
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ResourceHook is the schema for the resource hooks API,
// which orchestrates the previous and post steps around resource operations like plan and apply.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type ResourceHook struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceHookSpec   `json:"spec,omitempty"`
	Status ResourceHookStatus `json:"status,omitempty"`
}

type ResourceHookReference struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

var _ runtime.Object = (*ResourceHook)(nil)

// ResourceHookSpec defines the desired state of ResourceHook.
type ResourceHookSpec struct {
	// Plan defines the before and after steps to run in the resource hook.
	Plan ResourceOperationHook `json:"plan,omitempty"`

	// Apply defines the before and after steps to apply in the resource hook.
	Apply ResourceOperationHook `json:"apply,omitempty"`
}

// ResourceHookStatus defines the observed state of ResourceHook.
type ResourceHookStatus struct{}

// ResourceHookList holds the list of ResourceHook.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourceHookList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ResourceHook `json:"items"`
}

// ResourceOperationHook defines the previous and post steps to run in the resource hook for a resource operation.
type ResourceOperationHook struct {
	// Pre defines the previous steps to run in the operation hook.
	Pre []ResourceHookStep `json:"pre,omitempty"`

	// Post defines the post steps to run in the operation hook.
	Post []ResourceHookStep `json:"post,omitempty"`
}

// ResourceHookStep defines the desired state of ResourceHookStep.
type ResourceHookStep struct {
	// Name is the name of the step.
	Name string `json:"name"`

	// ResourceRunStepTemplate is a reference to a step template.
	ResourceRunStepTemplate *ResourceRunStepTemplateReference `json:"resourceRunStepTemplate,omitempty"`

	// Container defines the container step task to run in the resource hook step.
	// It is a reference to a container object. If the step template is not defined,
	// the container could be used to run the task.
	Container *core.Container `json:"container,omitempty"`
}
