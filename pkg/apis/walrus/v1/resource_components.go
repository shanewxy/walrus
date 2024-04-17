package v1

import (
	"k8s.io/apimachinery/pkg/runtime"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

// ResourceComponents is the schema for the resource components API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"],shortName=["rescomps"]
type ResourceComponents walruscore.ResourceComponents

var _ runtime.Object = (*ResourceComponents)(nil)

// ResourceComponentsList holds the list of ResourceComponents.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourceComponentsList walruscore.ResourceComponentsList

var _ runtime.Object = (*ResourceComponentsList)(nil)
