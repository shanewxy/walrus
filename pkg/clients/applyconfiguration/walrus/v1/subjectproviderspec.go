// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	v1 "github.com/seal-io/walrus/pkg/apis/walrus/v1"
)

// SubjectProviderSpecApplyConfiguration represents an declarative configuration of the SubjectProviderSpec type for use
// with apply.
type SubjectProviderSpecApplyConfiguration struct {
	Type           *v1.SubjectProviderType                          `json:"type,omitempty"`
	DisplayName    *string                                          `json:"displayName,omitempty"`
	Description    *string                                          `json:"description,omitempty"`
	ExternalConfig *SubjectProviderExternalConfigApplyConfiguration `json:"externalConfig,omitempty"`
}

// SubjectProviderSpecApplyConfiguration constructs an declarative configuration of the SubjectProviderSpec type for use with
// apply.
func SubjectProviderSpec() *SubjectProviderSpecApplyConfiguration {
	return &SubjectProviderSpecApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *SubjectProviderSpecApplyConfiguration) WithType(value v1.SubjectProviderType) *SubjectProviderSpecApplyConfiguration {
	b.Type = &value
	return b
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *SubjectProviderSpecApplyConfiguration) WithDisplayName(value string) *SubjectProviderSpecApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *SubjectProviderSpecApplyConfiguration) WithDescription(value string) *SubjectProviderSpecApplyConfiguration {
	b.Description = &value
	return b
}

// WithExternalConfig sets the ExternalConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalConfig field is set to the value of the last call.
func (b *SubjectProviderSpecApplyConfiguration) WithExternalConfig(value *SubjectProviderExternalConfigApplyConfiguration) *SubjectProviderSpecApplyConfiguration {
	b.ExternalConfig = value
	return b
}