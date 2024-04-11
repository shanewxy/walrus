package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// SubjectToken is the subresource of the Subject resource for token request.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"]
type SubjectToken struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec   SubjectTokenSpec   `json:"spec"`
	Status SubjectTokenStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*SubjectToken)(nil)

// SubjectTokenSpec defines the desired state of SubjectToken.
type SubjectTokenSpec struct {
	// ExpirationSeconds is the requested duration of validity of the request. The
	// token issuer may return a token with a different validity duration so a
	// client needs to check the 'expiration' field in a response.
	//
	// The value must be non-negative.
	// The maximum value is controlled by the loopback Kubernetes Cluster ApiServer.
	//
	// +optional
	// +k8s:validation:minimum=0
	// +k8s:validation:exclusiveMinimum
	ExpirationSeconds *int64 `json:"expirationSeconds,omitempty"`
}

// SubjectTokenStatus defines the observed state of SubjectToken.
type SubjectTokenStatus struct {
	// Token is the token of the SubjectToken.
	Token string `json:"token"`

	// ExpirationTimestamp is the time of expiration of the returned token.
	ExpirationTimestamp meta.Time `json:"expirationTimestamp"`
}
