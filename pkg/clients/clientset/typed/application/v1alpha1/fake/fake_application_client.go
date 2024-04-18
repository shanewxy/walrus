// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/seal-io/walrus/pkg/clients/clientset/typed/application/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeArgoprojV1alpha1 struct {
	*testing.Fake
}

func (c *FakeArgoprojV1alpha1) AppProjects(namespace string) v1alpha1.AppProjectInterface {
	return &FakeAppProjects{c, namespace}
}

func (c *FakeArgoprojV1alpha1) Applications(namespace string) v1alpha1.ApplicationInterface {
	return &FakeApplications{c, namespace}
}

func (c *FakeArgoprojV1alpha1) ApplicationSets(namespace string) v1alpha1.ApplicationSetInterface {
	return &FakeApplicationSets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeArgoprojV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}