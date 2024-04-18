// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

// SCMProviderGeneratorBitbucketApplyConfiguration represents an declarative configuration of the SCMProviderGeneratorBitbucket type for use
// with apply.
type SCMProviderGeneratorBitbucketApplyConfiguration struct {
	Owner          *string                      `json:"owner,omitempty"`
	User           *string                      `json:"user,omitempty"`
	AppPasswordRef *SecretRefApplyConfiguration `json:"appPasswordRef,omitempty"`
	AllBranches    *bool                        `json:"allBranches,omitempty"`
}

// SCMProviderGeneratorBitbucketApplyConfiguration constructs an declarative configuration of the SCMProviderGeneratorBitbucket type for use with
// apply.
func SCMProviderGeneratorBitbucket() *SCMProviderGeneratorBitbucketApplyConfiguration {
	return &SCMProviderGeneratorBitbucketApplyConfiguration{}
}

// WithOwner sets the Owner field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Owner field is set to the value of the last call.
func (b *SCMProviderGeneratorBitbucketApplyConfiguration) WithOwner(value string) *SCMProviderGeneratorBitbucketApplyConfiguration {
	b.Owner = &value
	return b
}

// WithUser sets the User field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the User field is set to the value of the last call.
func (b *SCMProviderGeneratorBitbucketApplyConfiguration) WithUser(value string) *SCMProviderGeneratorBitbucketApplyConfiguration {
	b.User = &value
	return b
}

// WithAppPasswordRef sets the AppPasswordRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AppPasswordRef field is set to the value of the last call.
func (b *SCMProviderGeneratorBitbucketApplyConfiguration) WithAppPasswordRef(value *SecretRefApplyConfiguration) *SCMProviderGeneratorBitbucketApplyConfiguration {
	b.AppPasswordRef = value
	return b
}

// WithAllBranches sets the AllBranches field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllBranches field is set to the value of the last call.
func (b *SCMProviderGeneratorBitbucketApplyConfiguration) WithAllBranches(value bool) *SCMProviderGeneratorBitbucketApplyConfiguration {
	b.AllBranches = &value
	return b
}