package v1

import (
	"errors"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Subject is the schema for the subjects API.
//
// +genclient
// +genclient:onlyVerbs=create,get,list,watch,apply,update,patch,delete,deleteCollection
// +genclient:method=Login,verb=create,subresource=login,input=SubjectLogin,result=SubjectLogin
// +genclient:method=CreateToken,verb=create,subresource=token,input=SubjectToken,result=SubjectToken
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"],shortName=["subj"]
type Subject struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec SubjectSpec `json:"spec"`
}

var _ runtime.Object = (*Subject)(nil)

// SubjectReference is the reference of the subject.
type SubjectReference struct {
	// Namespace is the namespace of the subject.
	Namespace string `json:"namespace"`
	// Name is the name of the subject.
	Name string `json:"name"`
}

func (in *SubjectReference) ToNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: in.Namespace,
		Name:      in.Name,
	}
}

// SubjectRole describes the role of subject.
// +enum
type SubjectRole string

const (
	// SubjectRoleUser is the subject role for subject user.
	SubjectRoleUser SubjectRole = "User"
	// SubjectRoleManager is the subject role for subject manager.
	SubjectRoleManager SubjectRole = "Manager"
	// SubjectRoleAdmin is the subject role for subject admin.
	SubjectRoleAdmin SubjectRole = "Admin"
)

func (in SubjectRole) String() string {
	return string(in)
}

func (in SubjectRole) Validate() error {
	switch in {
	case SubjectRoleUser, SubjectRoleManager, SubjectRoleAdmin:
		return nil
	default:
		return errors.New("invalid subject role")
	}
}

// SubjectSpec defines the desired state of Subject.
type SubjectSpec struct {
	// Provider is the name of subject provider who provides this subject,
	// which is immutable.
	Provider string `json:"provider"`

	// Role is the role of the subject.
	//
	// +k8s:validation:enum=["User","Manager","Admin"]
	Role SubjectRole `json:"role"`

	// DisplayName is the display name of the subject.
	DisplayName string `json:"displayName,omitempty"`

	// Description is the description of the subject.
	Description string `json:"description,omitempty"`

	// Email is the email of the subject.
	//
	// +k8s:validation:format="email"
	Email string `json:"email"`

	// Groups is the groups that the subject belongs to.
	//
	// +k8s:validation:uniqueItems=true
	Groups []string `json:"groups,omitempty"`

	// Credential is the credential of the subject,
	// it is provided as a write-only input field.
	//
	// +k8s:validation:format="password"
	Credential *string `json:"credential,omitempty"`
}

// SubjectList holds the list of Subject.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SubjectList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Subject `json:"items"`
}

var _ runtime.Object = (*SubjectList)(nil)
