package v1

import (
	"k8s.io/apimachinery/pkg/runtime"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

// ConnectorBinding is the schema for the connectorbindings API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"],shortName=["cb"],subResources=["status"]
type ConnectorBinding walruscore.ConnectorBinding

var _ runtime.Object = (*ConnectorBinding)(nil)

// ConnectorBindingList contains a list of ConnectorBinding.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ConnectorBindingList walruscore.ConnectorBindingList

var _ runtime.Object = (*ConnectorBindingList)(nil)
