// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// CreateS3BucketOptionsApplyConfiguration represents an declarative configuration of the CreateS3BucketOptions type for use
// with apply.
type CreateS3BucketOptionsApplyConfiguration struct {
	ObjectLocking *bool `json:"objectLocking,omitempty"`
}

// CreateS3BucketOptionsApplyConfiguration constructs an declarative configuration of the CreateS3BucketOptions type for use with
// apply.
func CreateS3BucketOptions() *CreateS3BucketOptionsApplyConfiguration {
	return &CreateS3BucketOptionsApplyConfiguration{}
}

// WithObjectLocking sets the ObjectLocking field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObjectLocking field is set to the value of the last call.
func (b *CreateS3BucketOptionsApplyConfiguration) WithObjectLocking(value bool) *CreateS3BucketOptionsApplyConfiguration {
	b.ObjectLocking = &value
	return b
}