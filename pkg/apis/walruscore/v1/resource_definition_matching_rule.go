package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ResourceDefinitionMatchingRule is the schema for the resource definition matching rules API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type ResourceDefinitionMatchingRule struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceDefinitionMatchingRuleSpec   `json:"spec,omitempty"`
	Status ResourceDefinitionMatchingRuleStatus `json:"status,omitempty"`
}

type ResourceDefinitionMatchingRuleReference struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

var _ runtime.Object = (*ResourceDefinitionMatchingRule)(nil)

// ResourceDefinitionMatchingRuleSpec defines the desired state of ResourceDefinitionMatchingRule.
type ResourceDefinitionMatchingRuleSpec struct {
	// TODO
}

// ResourceDefinitionMatchingRuleStatus defines the observed state of ResourceDefinitionMatchingRule.
type ResourceDefinitionMatchingRuleStatus struct {
	// TODO
}

// ResourceDefinitionMatchingRuleList holds the list of ResourceDefinitionMatchingRule.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourceDefinitionMatchingRuleList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ResourceDefinitionMatchingRule `json:"items"`
}

var _ runtime.Object = (*ResourceDefinitionMatchingRuleList)(nil)
