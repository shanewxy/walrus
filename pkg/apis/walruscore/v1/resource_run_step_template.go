package v1

import (
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ResourceRunStepTemplate is the schema for the resource run step templates API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type ResourceRunStepTemplate struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceRunStepTemplateSpec   `json:"spec,omitempty"`
	Status ResourceRunStepTemplateStatus `json:"status,omitempty"`
}

type ResourceRunStepTemplateReference struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

var _ runtime.Object = (*ResourceRunStepTemplate)(nil)

// ResourceRunStepTemplateSpec defines the desired state of ResourceRunStepTemplate.
type ResourceRunStepTemplateSpec struct {
	// Container is the main container image to run in the resource run step template.
	Container *core.Container `json:"container,omitempty"`

	// Approval is the approval process for the resource run step template.
	Approval *ResourceRunStepApprovalTemplate `json:"approval,omitempty"`
}

// ResourceRunStepTemplateStatus defines the observed state of ResourceRunStepTemplate.
type ResourceRunStepTemplateStatus struct{}

// ResourceRunStepTemplateList holds the list of ResourceRunStepTemplate.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourceRunStepTemplateList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ResourceRunStepTemplate `json:"items"`
}

var _ runtime.Object = (*ResourceRunStepTemplateList)(nil)

// ResourceRunStepTemplateApprovalType orchesrates the approval process for the resource run step template.
// +enum
type ResourceRunStepTemplateApprovalType string

const (
	// ResourceRunStepTemplateApprovalTypeAny means step is approved
	// if any of the approval users approves it.
	ResourceRunStepTemplateApprovalTypeAny ResourceRunStepTemplateApprovalType = "any"
	// ResourceRunStepTemplateApprovalTypeAnd means step is approved
	// only all of the approval users approve it.
	ResourceRunStepTemplateApprovalTypeAnd ResourceRunStepTemplateApprovalType = "and"
)

// ResourceRunStepApprovalTemplate orchesrates the approval process for the resource run step template.
type ResourceRunStepApprovalTemplate struct {
	// Type is the type of the approval process.
	Type ResourceRunStepTemplateApprovalType `json:"type"`

	// Users is the users to approve the step.
	Users []string `json:"users"`
}
