// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	v1 "github.com/seal-io/walrus/pkg/apis/walrus/v1"
)

// SubjectProviderGithubConfigApplyConfiguration represents an declarative configuration of the SubjectProviderGithubConfig type for use
// with apply.
type SubjectProviderGithubConfigApplyConfiguration struct {
	ClientID     *string                             `json:"clientID,omitempty"`
	ClientSecret *string                             `json:"clientSecret,omitempty"`
	Groups       *v1.SubjectProviderGitGroupsMatcher `json:"groups,omitempty"`
}

// SubjectProviderGithubConfigApplyConfiguration constructs an declarative configuration of the SubjectProviderGithubConfig type for use with
// apply.
func SubjectProviderGithubConfig() *SubjectProviderGithubConfigApplyConfiguration {
	return &SubjectProviderGithubConfigApplyConfiguration{}
}

// WithClientID sets the ClientID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientID field is set to the value of the last call.
func (b *SubjectProviderGithubConfigApplyConfiguration) WithClientID(value string) *SubjectProviderGithubConfigApplyConfiguration {
	b.ClientID = &value
	return b
}

// WithClientSecret sets the ClientSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientSecret field is set to the value of the last call.
func (b *SubjectProviderGithubConfigApplyConfiguration) WithClientSecret(value string) *SubjectProviderGithubConfigApplyConfiguration {
	b.ClientSecret = &value
	return b
}

// WithGroups sets the Groups field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Groups field is set to the value of the last call.
func (b *SubjectProviderGithubConfigApplyConfiguration) WithGroups(value v1.SubjectProviderGitGroupsMatcher) *SubjectProviderGithubConfigApplyConfiguration {
	b.Groups = &value
	return b
}
