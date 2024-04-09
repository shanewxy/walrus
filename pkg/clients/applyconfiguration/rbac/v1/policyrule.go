// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// PolicyRuleApplyConfiguration represents an declarative configuration of the PolicyRule type for use
// with apply.
type PolicyRuleApplyConfiguration struct {
	Verbs           []string `json:"verbs,omitempty"`
	APIGroups       []string `json:"apiGroups,omitempty"`
	Resources       []string `json:"resources,omitempty"`
	ResourceNames   []string `json:"resourceNames,omitempty"`
	NonResourceURLs []string `json:"nonResourceURLs,omitempty"`
}

// PolicyRuleApplyConfiguration constructs an declarative configuration of the PolicyRule type for use with
// apply.
func PolicyRule() *PolicyRuleApplyConfiguration {
	return &PolicyRuleApplyConfiguration{}
}

// WithVerbs adds the given value to the Verbs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Verbs field.
func (b *PolicyRuleApplyConfiguration) WithVerbs(values ...string) *PolicyRuleApplyConfiguration {
	for i := range values {
		b.Verbs = append(b.Verbs, values[i])
	}
	return b
}

// WithAPIGroups adds the given value to the APIGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIGroups field.
func (b *PolicyRuleApplyConfiguration) WithAPIGroups(values ...string) *PolicyRuleApplyConfiguration {
	for i := range values {
		b.APIGroups = append(b.APIGroups, values[i])
	}
	return b
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *PolicyRuleApplyConfiguration) WithResources(values ...string) *PolicyRuleApplyConfiguration {
	for i := range values {
		b.Resources = append(b.Resources, values[i])
	}
	return b
}

// WithResourceNames adds the given value to the ResourceNames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourceNames field.
func (b *PolicyRuleApplyConfiguration) WithResourceNames(values ...string) *PolicyRuleApplyConfiguration {
	for i := range values {
		b.ResourceNames = append(b.ResourceNames, values[i])
	}
	return b
}

// WithNonResourceURLs adds the given value to the NonResourceURLs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NonResourceURLs field.
func (b *PolicyRuleApplyConfiguration) WithNonResourceURLs(values ...string) *PolicyRuleApplyConfiguration {
	for i := range values {
		b.NonResourceURLs = append(b.NonResourceURLs, values[i])
	}
	return b
}