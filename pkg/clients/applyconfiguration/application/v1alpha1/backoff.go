// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// BackoffApplyConfiguration represents an declarative configuration of the Backoff type for use
// with apply.
type BackoffApplyConfiguration struct {
	Duration    *string `json:"duration,omitempty"`
	Factor      *int64  `json:"factor,omitempty"`
	MaxDuration *string `json:"maxDuration,omitempty"`
}

// BackoffApplyConfiguration constructs an declarative configuration of the Backoff type for use with
// apply.
func Backoff() *BackoffApplyConfiguration {
	return &BackoffApplyConfiguration{}
}

// WithDuration sets the Duration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Duration field is set to the value of the last call.
func (b *BackoffApplyConfiguration) WithDuration(value string) *BackoffApplyConfiguration {
	b.Duration = &value
	return b
}

// WithFactor sets the Factor field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Factor field is set to the value of the last call.
func (b *BackoffApplyConfiguration) WithFactor(value int64) *BackoffApplyConfiguration {
	b.Factor = &value
	return b
}

// WithMaxDuration sets the MaxDuration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxDuration field is set to the value of the last call.
func (b *BackoffApplyConfiguration) WithMaxDuration(value string) *BackoffApplyConfiguration {
	b.MaxDuration = &value
	return b
}