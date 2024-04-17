// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// ResourceDefinitionReferenceApplyConfiguration represents an declarative configuration of the ResourceDefinitionReference type for use
// with apply.
type ResourceDefinitionReferenceApplyConfiguration struct {
	Namespace *string `json:"namespace,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// ResourceDefinitionReferenceApplyConfiguration constructs an declarative configuration of the ResourceDefinitionReference type for use with
// apply.
func ResourceDefinitionReference() *ResourceDefinitionReferenceApplyConfiguration {
	return &ResourceDefinitionReferenceApplyConfiguration{}
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ResourceDefinitionReferenceApplyConfiguration) WithNamespace(value string) *ResourceDefinitionReferenceApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceDefinitionReferenceApplyConfiguration) WithName(value string) *ResourceDefinitionReferenceApplyConfiguration {
	b.Name = &value
	return b
}
