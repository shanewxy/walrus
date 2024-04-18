// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// BearerTokenBitbucketCloudApplyConfiguration represents an declarative configuration of the BearerTokenBitbucketCloud type for use
// with apply.
type BearerTokenBitbucketCloudApplyConfiguration struct {
	TokenRef *SecretRefApplyConfiguration `json:"tokenRef,omitempty"`
}

// BearerTokenBitbucketCloudApplyConfiguration constructs an declarative configuration of the BearerTokenBitbucketCloud type for use with
// apply.
func BearerTokenBitbucketCloud() *BearerTokenBitbucketCloudApplyConfiguration {
	return &BearerTokenBitbucketCloudApplyConfiguration{}
}

// WithTokenRef sets the TokenRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TokenRef field is set to the value of the last call.
func (b *BearerTokenBitbucketCloudApplyConfiguration) WithTokenRef(value *SecretRefApplyConfiguration) *BearerTokenBitbucketCloudApplyConfiguration {
	b.TokenRef = value
	return b
}