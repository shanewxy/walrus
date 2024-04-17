// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// ResourceRunStatusApplyConfiguration represents an declarative configuration of the ResourceRunStatus type for use
// with apply.
type ResourceRunStatusApplyConfiguration struct {
	StatusDescriptorApplyConfiguration `json:",inline"`
	ComputedAttributes                 *runtime.RawExtension                             `json:"computedAttributes,omitempty"`
	TemplateFormat                     *string                                           `json:"templateFormat,omitempty"`
	ConfigSecretName                   *string                                           `json:"configSecretName,omitempty"`
	ComponentChanges                   []byte                                            `json:"componentChanges,omitempty"`
	ComponentChangeSummary             *ResourceComponentChangeSummaryApplyConfiguration `json:"componentChangeSummary,omitempty"`
	ResourceRunTemplate                *ResourceRunTemplateReferenceApplyConfiguration   `json:"resourceRunTemplate,omitempty"`
	Steps                              []ResourceRunStepApplyConfiguration               `json:"steps,omitempty"`
}

// ResourceRunStatusApplyConfiguration constructs an declarative configuration of the ResourceRunStatus type for use with
// apply.
func ResourceRunStatus() *ResourceRunStatusApplyConfiguration {
	return &ResourceRunStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *ResourceRunStatusApplyConfiguration) WithPhase(value string) *ResourceRunStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithPhaseMessage sets the PhaseMessage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PhaseMessage field is set to the value of the last call.
func (b *ResourceRunStatusApplyConfiguration) WithPhaseMessage(value string) *ResourceRunStatusApplyConfiguration {
	b.PhaseMessage = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ResourceRunStatusApplyConfiguration) WithConditions(values ...*ConditionApplyConfiguration) *ResourceRunStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithComputedAttributes sets the ComputedAttributes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ComputedAttributes field is set to the value of the last call.
func (b *ResourceRunStatusApplyConfiguration) WithComputedAttributes(value runtime.RawExtension) *ResourceRunStatusApplyConfiguration {
	b.ComputedAttributes = &value
	return b
}

// WithTemplateFormat sets the TemplateFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TemplateFormat field is set to the value of the last call.
func (b *ResourceRunStatusApplyConfiguration) WithTemplateFormat(value string) *ResourceRunStatusApplyConfiguration {
	b.TemplateFormat = &value
	return b
}

// WithConfigSecretName sets the ConfigSecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigSecretName field is set to the value of the last call.
func (b *ResourceRunStatusApplyConfiguration) WithConfigSecretName(value string) *ResourceRunStatusApplyConfiguration {
	b.ConfigSecretName = &value
	return b
}

// WithComponentChanges adds the given value to the ComponentChanges field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ComponentChanges field.
func (b *ResourceRunStatusApplyConfiguration) WithComponentChanges(values ...byte) *ResourceRunStatusApplyConfiguration {
	for i := range values {
		b.ComponentChanges = append(b.ComponentChanges, values[i])
	}
	return b
}

// WithComponentChangeSummary sets the ComponentChangeSummary field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ComponentChangeSummary field is set to the value of the last call.
func (b *ResourceRunStatusApplyConfiguration) WithComponentChangeSummary(value *ResourceComponentChangeSummaryApplyConfiguration) *ResourceRunStatusApplyConfiguration {
	b.ComponentChangeSummary = value
	return b
}

// WithResourceRunTemplate sets the ResourceRunTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceRunTemplate field is set to the value of the last call.
func (b *ResourceRunStatusApplyConfiguration) WithResourceRunTemplate(value *ResourceRunTemplateReferenceApplyConfiguration) *ResourceRunStatusApplyConfiguration {
	b.ResourceRunTemplate = value
	return b
}

// WithSteps adds the given value to the Steps field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Steps field.
func (b *ResourceRunStatusApplyConfiguration) WithSteps(values ...*ResourceRunStepApplyConfiguration) *ResourceRunStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSteps")
		}
		b.Steps = append(b.Steps, *values[i])
	}
	return b
}
