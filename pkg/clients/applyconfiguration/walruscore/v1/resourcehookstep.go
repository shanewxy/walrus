// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	corev1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
)

// ResourceHookStepApplyConfiguration represents an declarative configuration of the ResourceHookStep type for use
// with apply.
type ResourceHookStepApplyConfiguration struct {
	Name                    *string                                             `json:"name,omitempty"`
	ResourceRunStepTemplate *ResourceRunStepTemplateReferenceApplyConfiguration `json:"resourceRunStepTemplate,omitempty"`
	Container               *corev1.ContainerApplyConfiguration                 `json:"container,omitempty"`
}

// ResourceHookStepApplyConfiguration constructs an declarative configuration of the ResourceHookStep type for use with
// apply.
func ResourceHookStep() *ResourceHookStepApplyConfiguration {
	return &ResourceHookStepApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceHookStepApplyConfiguration) WithName(value string) *ResourceHookStepApplyConfiguration {
	b.Name = &value
	return b
}

// WithResourceRunStepTemplate sets the ResourceRunStepTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceRunStepTemplate field is set to the value of the last call.
func (b *ResourceHookStepApplyConfiguration) WithResourceRunStepTemplate(value *ResourceRunStepTemplateReferenceApplyConfiguration) *ResourceHookStepApplyConfiguration {
	b.ResourceRunStepTemplate = value
	return b
}

// WithContainer sets the Container field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Container field is set to the value of the last call.
func (b *ResourceHookStepApplyConfiguration) WithContainer(value *corev1.ContainerApplyConfiguration) *ResourceHookStepApplyConfiguration {
	b.Container = value
	return b
}