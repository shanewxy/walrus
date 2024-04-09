// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package fake

import (
	v1 "github.com/seal-io/walrus/pkg/clients/clientset/typed/discovery/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDiscoveryV1 struct {
	*testing.Fake
}

func (c *FakeDiscoveryV1) EndpointSlices(namespace string) v1.EndpointSliceInterface {
	return &FakeEndpointSlices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDiscoveryV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}