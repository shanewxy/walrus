package v1

import (
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ResourceRun is the schema for the resource runs API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type ResourceRun struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceRunSpec   `json:"spec,omitempty"`
	Status ResourceRunStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*ResourceRun)(nil)

// ResourceRunSpec defines the desired state of ResourceRun.
type ResourceRunSpec struct {
	// Project is the project of the resource run.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	Project string `json:"project"`

	// ResourceName is the resource name of the resource run.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	ResourceName string `json:"resource"`

	// Type is the type of the resource run.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	Type ResourceRunType `json:"type"`

	// Attributes is the attributes of the resource run.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	Attributes runtime.RawExtension `json:"attributes"`

	// TemplateVersion template version to which the resource belongs.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	TemplateVersion *TempalteVersionReference `json:"templateVersionReference"`
}

// ResourceRunStatus defines the observed state of ResourceRun.
type ResourceRunStatus struct {
	// StatusDescriptor defines the status of the resource run.
	StatusDescriptor `json:",inline"`

	// ComputedAttributes is the computed attributes of the resource run.
	ComputedAttributes runtime.RawExtension `json:"computedAttributes"`

	// TemplateFormat is the format of template version.
	TemplateFormat string `json:"templateFormat"`

	// ConfigSecretName is the name of generated secret stores configs for the resource run.
	ConfigSecretName string `json:"configSecretName"`

	// ComponentChanges is the changes in the components of the resource run.
	ComponentChanges []byte `json:"componentChanges,omitempty"`

	// ComponentChangeSummary is the summary of the component changes.
	ComponentChangeSummary ResourceComponentChangeSummary `json:"componentChangeSummary,omitempty"`

	// ResourceRunTemplate is a reference to a resource run template.
	ResourceRunTemplate *ResourceRunTemplateReference `json:"resourceRunTemplate,omitempty"`

	// Steps stores the step run results of the resource run.
	Steps []ResourceRunStep `json:"steps"`
}

// ResourceRunList holds the list of ResourceRun.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourceRunList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ResourceRun `json:"items"`
}

var _ runtime.Object = (*ResourceRunList)(nil)

// ResourceRunType describes the type of the resource run.
// +enum
type ResourceRunType string

const (
	ResourceRunTypeCreate   ResourceRunType = "Create"
	ResourceRunTypeUpdate   ResourceRunType = "Update"
	ResourceRunTypeDelete   ResourceRunType = "Delete"
	ResourceRunTypeStart    ResourceRunType = "Start"
	ResourceRunTypeStop     ResourceRunType = "Stop"
	ResourceRunTypeRollback ResourceRunType = "Rollback"
)

// ResourceRunOperationType describes the type of the system operation of the resource run to be performed with deployer.
// +enum
type ResourceRunOperationType string

const (
	ResourceRunOperationTypePlan  ResourceRunOperationType = "Plan"
	ResourceRunOperationTypeApply ResourceRunOperationType = "Apply"
)

func (t ResourceRunOperationType) String() string {
	return string(t)
}

// ResourceRunStepType describes the type of the resource run step.
// +enum
type ResourceRunStepType string

const (
	// ResourceRunStepTypePlan is the type of the plan resource, it is system generated.
	ResourceRunStepTypePlan ResourceRunStepType = "Plan"
	// ResourceRunStepTypeApply is the type of the apply resource, it is system generated.
	ResourceRunStepTypeApply ResourceRunStepType = "Apply"

	// ResourceRunStepTypeContainer is the type of the container resource, it depends on the resource run hooks and resource run template.
	ResourceRunStepTypeContainer ResourceRunStepType = "Container"
	// ResourceRunStepTypeApproval is the type of the approval resource, it depends on the resource run hooks and resource run template.
	ResourceRunStepTypeApproval ResourceRunStepType = "Approval"
	ResourceRunStepTypeUnknown  ResourceRunStepType = "Unknown"
)

// ResourceRunStep stores the step run result of the resource run.
// The step may be the plan, apply, and hook step defined in the resource hook or resource run template.
type ResourceRunStep struct {
	// StatusDescriptor defines the status of the step.
	StatusDescriptor `json:",inline"`

	// Name is the name of the step.
	Name string `json:"name"`

	// Type is the type of the step.
	Type ResourceRunStepType `json:"type"`

	// Template is a reference to a resource step template.
	Template *ResourceRunStepTemplateReference `json:"template,omitempty"`

	// StartTime is the time when the step started.
	StartTime meta.Time `json:"startTime,omitempty"`

	// FinishTime is the time when the step finished.
	FinishTime meta.Time `json:"finishTime,omitempty"`

	// DurationInSeconds is the duration of the step.
	DurationInSeconds int64 `json:"durationInSeconds,omitempty"`

	// Container is the container configs to run in the resource run step.
	// Only used when ResourceStepTemplate is not provided.
	Container *core.Container `json:"container,omitempty"`

	// ApprovalRecord stores the approval records of the step.
	// Only the step template reference that has approval spec will have approval record.
	ApprovalRecord *ResourceRunStepApprovalRecord `json:"approvalRecord,omitempty"`
}

// ResourceRunStepApprovalUserActionType describes the type of the user action.
// +enum
type ResourceRunStepApprovalUserActionType string

const (
	// ResourceRunStepApprovalUserActionTypeApprove is the type of the approve action.
	ResourceRunStepApprovalUserActionTypeApprove ResourceRunStepApprovalUserActionType = "Approve"

	// ResourceRunStepApprovalUserActionTypeReject is the type of the reject action.
	ResourceRunStepApprovalUserActionTypeReject ResourceRunStepApprovalUserActionType = "Reject"
)

type ResourceRunStepApprovalUserAction struct {
	// Type is the type of the user action.
	Type ResourceRunStepApprovalUserActionType `json:"type"`
	// User is the user who performed the action.
	User string `json:"user"`
	// Time is the time when the action performed.
	Time meta.Time `json:"time"`
	// Comment is the comment of the action.
	Comment string `json:"comment"`
}

// ResourceRunStepApprovalRecord stores the approval records of target resource run step.
type ResourceRunStepApprovalRecord struct {
	// Type is the type of the approval process.
	Type ResourceRunStepTemplateApprovalType `json:"type"`

	// Users is the users to approve the step.
	//
	// +listType=set
	Users []string `json:"users"`

	// Actions is the user actions of the approval process.
	//
	// +listType=map
	// +listMapKey=user
	Actions []ResourceRunStepApprovalUserAction `json:"actions"`
}

// Check checks if the approval is approved and ready.
func (s *ResourceRunStepApprovalRecord) Check() (approval, ready bool) {
	if len(s.Actions) == 0 {
		return false, false
	}

	switch s.Type {
	case ResourceRunStepTemplateApprovalTypeAny:
		for i := range s.Actions {
			// If any user has approved, the step is already approved.
			if s.Actions[i].Type == ResourceRunStepApprovalUserActionTypeApprove {
				return true, true
			}
		}

		// If all users have rejected, the approval is already rejected.
		if len(s.Actions) == len(s.Users) {
			return false, true
		}

		return false, false

	case ResourceRunStepTemplateApprovalTypeAnd:
		for i := range s.Actions {
			// If any user has rejected, the approval is already rejected.
			if s.Actions[i].Type == ResourceRunStepApprovalUserActionTypeReject {
				return false, true
			}
		}

		// If not all users have approved, the step is not ready.
		if len(s.Actions) < len(s.Users) {
			return false, false
		}

		// If all user approved, the approval is already approved.
		return true, true
	}

	return false, false
}

func (s *ResourceRunStepApprovalRecord) SetUserAction(user, comment string, action ResourceRunStepApprovalUserActionType) error {
	s.Actions = append(s.Actions, ResourceRunStepApprovalUserAction{
		Type:    action,
		User:    user,
		Comment: comment,
		Time:    meta.Now(),
	})

	return nil
}

func (s *ResourceRunStepApprovalRecord) SetApprovedUser(user, comment string) error {
	return s.SetUserAction(user, comment, ResourceRunStepApprovalUserActionTypeApprove)
}

func (s *ResourceRunStepApprovalRecord) SetRejectedUser(user, comment string) error {
	return s.SetUserAction(user, comment, ResourceRunStepApprovalUserActionTypeReject)
}

// ResourceComponentChangeSummary is the summary of the component changes.
type ResourceComponentChangeSummary struct {
	// Created is the number of created components.
	Created *int64 `json:"created"`

	// Updated is the number of updated components.
	Updated *int64 `json:"updated"`

	// Deleted is the number of deleted components.
	Deleted *int64 `json:"deleted"`
}
