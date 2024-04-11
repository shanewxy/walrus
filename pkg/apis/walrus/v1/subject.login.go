package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// SubjectLogin is the subresource of the Subject resource for login request.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"]
type SubjectLogin struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubjectLoginSpec   `json:"spec"`
	Status SubjectLoginStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*SubjectLogin)(nil)

// SubjectLoginSpec defines the desired state of SubjectLogin.
type SubjectLoginSpec struct {
	// Credential is the credential of the subject,
	// it is provided as a write-only input field.
	//
	// +k8s:validation:format="password"
	Credential string `json:"credential"`
}

// SubjectLoginStatus defines the observed state of SubjectLogin.
type SubjectLoginStatus struct {
	// Token is the token of the SubjectLogin.
	Token string `json:"token"`
}
