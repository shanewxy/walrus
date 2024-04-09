// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

// LimitRangeSpecApplyConfiguration represents an declarative configuration of the LimitRangeSpec type for use
// with apply.
type LimitRangeSpecApplyConfiguration struct {
	Limits []LimitRangeItemApplyConfiguration `json:"limits,omitempty"`
}

// LimitRangeSpecApplyConfiguration constructs an declarative configuration of the LimitRangeSpec type for use with
// apply.
func LimitRangeSpec() *LimitRangeSpecApplyConfiguration {
	return &LimitRangeSpecApplyConfiguration{}
}

// WithLimits adds the given value to the Limits field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Limits field.
func (b *LimitRangeSpecApplyConfiguration) WithLimits(values ...*LimitRangeItemApplyConfiguration) *LimitRangeSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithLimits")
		}
		b.Limits = append(b.Limits, *values[i])
	}
	return b
}