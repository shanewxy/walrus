// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// SubjectProviderStatusApplyConfiguration represents an declarative configuration of the SubjectProviderStatus type for use
// with apply.
type SubjectProviderStatusApplyConfiguration struct {
	LoginWithPassword *bool `json:"loginWithPassword,omitempty"`
}

// SubjectProviderStatusApplyConfiguration constructs an declarative configuration of the SubjectProviderStatus type for use with
// apply.
func SubjectProviderStatus() *SubjectProviderStatusApplyConfiguration {
	return &SubjectProviderStatusApplyConfiguration{}
}

// WithLoginWithPassword sets the LoginWithPassword field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LoginWithPassword field is set to the value of the last call.
func (b *SubjectProviderStatusApplyConfiguration) WithLoginWithPassword(value bool) *SubjectProviderStatusApplyConfiguration {
	b.LoginWithPassword = &value
	return b
}