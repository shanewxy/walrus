package v1

import (
	"errors"
	"reflect"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Variable is the schema for the variables API.
//
// +genclient
// +genclient:onlyVerbs=create,get,list,watch,apply,update,patch,delete,deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"],shortName=["var"]
type Variable struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   VariableSpec   `json:"spec"`
	Status VariableStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Variable)(nil)

// VariableSpec defines the desired state of Variable.
type VariableSpec struct {
	// Value contains the configuration data,
	// it is provided as a write-only input field.
	Value *string `json:"value,omitempty"`

	// Sensitive indicates whether the variable is sensitive.
	Sensitive bool `json:"sensitive,omitempty"`
}

// VariableScope defines the scope of the variable.
// +enum
type VariableScope string

const (
	// VariableScopeSystem represents the system scope.
	VariableScopeSystem VariableScope = "System"
	// VariableScopeProject represents the project scope.
	VariableScopeProject VariableScope = "Project"
	// VariableScopeEnvironment represents the environment scope.
	VariableScopeEnvironment VariableScope = "Environment"
)

func (in VariableScope) String() string {
	return string(in)
}

func (in VariableScope) Validate() error {
	switch in {
	case VariableScopeSystem, VariableScopeProject, VariableScopeEnvironment:
		return nil
	default:
		return errors.New("invalid variable scope")
	}
}

func (in VariableScope) Priority() int {
	switch in {
	case VariableScopeProject:
		return 1
	case VariableScopeEnvironment:
		return 2
	}
	return 0
}

// VariableStatus defines the observed state of Variable.
type VariableStatus struct {
	// Scope is the scope of the variable.
	//
	// +k8s:validation:enum=["System","Project","Environment"]
	Scope VariableScope `json:"scope"`

	// Value is the current value of the setting,
	// it is provided as a read-only output field.
	//
	// "(sensitive)" returns if the variable is sensitive.
	Value string `json:"value"`
	// Value_ is the shadow of the Value,
	// it is provided for system processing only.
	//
	// DO NOT EXPOSE AND STORE IT.
	Value_ string `json:"-"`
}

func (in *Variable) Equal(in2 *Variable) bool {
	return reflect.DeepEqual(in.Spec, in2.Spec) &&
		in.Status.Scope == in2.Status.Scope &&
		in.Status.Value_ == in2.Status.Value_
}

// VariableList holds the list of Variable.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VariableList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Variable `json:"items"`
}

var _ runtime.Object = (*VariableList)(nil)
