package v1

import (
	"errors"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ProjectSubjects holds the list of ProjectSubject.
//
// ProjectSubjects is the subresource of Project to manage the subjects.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"],shortName=["projsub"]
type ProjectSubjects struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	// +patchStrategy=merge
	// +patchMergeKey=name
	// +listType=map
	// +listMapKey=name
	Items []ProjectSubject `json:"items,omitempty" patchStrategy:"merge" patchMergeKey:"name"`
}

var _ runtime.Object = (*ProjectSubjects)(nil)

// ProjectRole describes the role of project subject.
// +enum
type ProjectRole string

const (
	// ProjectRoleViewer is the role for project viewer.
	ProjectRoleViewer ProjectRole = "Viewer"
	// ProjectRoleMember is the role for project member.
	ProjectRoleMember ProjectRole = "Member"
	// ProjectRoleOwner is the role for project owner.
	ProjectRoleOwner ProjectRole = "Owner"
)

func (in ProjectRole) String() string {
	return string(in)
}

func (in ProjectRole) Validate() error {
	switch in {
	case ProjectRoleViewer, ProjectRoleMember, ProjectRoleOwner:
		return nil
	default:
		return errors.New("invalid project role")
	}
}

// ProjectSubject is the schema for the project subject API.
type ProjectSubject struct {
	SubjectReference `json:",inline"`

	// Role is the project role of the subject.
	//
	// +k8s:validation:enum=["Viewer","Member","Owner"]
	Role ProjectRole `json:"role"`
}
