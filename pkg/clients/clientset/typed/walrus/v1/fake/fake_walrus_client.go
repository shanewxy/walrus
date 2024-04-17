// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package fake

import (
	v1 "github.com/seal-io/walrus/pkg/clients/clientset/typed/walrus/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeWalrusV1 struct {
	*testing.Fake
}

func (c *FakeWalrusV1) Catalogs(namespace string) v1.CatalogInterface {
	return &FakeCatalogs{c, namespace}
}

func (c *FakeWalrusV1) Connectors(namespace string) v1.ConnectorInterface {
	return &FakeConnectors{c, namespace}
}

func (c *FakeWalrusV1) Environments(namespace string) v1.EnvironmentInterface {
	return &FakeEnvironments{c, namespace}
}

func (c *FakeWalrusV1) FileExamples(namespace string) v1.FileExampleInterface {
	return &FakeFileExamples{c, namespace}
}

func (c *FakeWalrusV1) Projects(namespace string) v1.ProjectInterface {
	return &FakeProjects{c, namespace}
}

func (c *FakeWalrusV1) Resources(namespace string) v1.ResourceInterface {
	return &FakeResources{c, namespace}
}

func (c *FakeWalrusV1) ResourceComponents(namespace string) v1.ResourceComponentsInterface {
	return &FakeResourceComponents{c, namespace}
}

func (c *FakeWalrusV1) ResourceDefinitions(namespace string) v1.ResourceDefinitionInterface {
	return &FakeResourceDefinitions{c, namespace}
}

func (c *FakeWalrusV1) ResourceRuns(namespace string) v1.ResourceRunInterface {
	return &FakeResourceRuns{c, namespace}
}

func (c *FakeWalrusV1) Schemas(namespace string) v1.SchemaInterface {
	return &FakeSchemas{c, namespace}
}

func (c *FakeWalrusV1) Settings(namespace string) v1.SettingInterface {
	return &FakeSettings{c, namespace}
}

func (c *FakeWalrusV1) Subjects(namespace string) v1.SubjectInterface {
	return &FakeSubjects{c, namespace}
}

func (c *FakeWalrusV1) SubjectProviders(namespace string) v1.SubjectProviderInterface {
	return &FakeSubjectProviders{c, namespace}
}

func (c *FakeWalrusV1) Templates(namespace string) v1.TemplateInterface {
	return &FakeTemplates{c, namespace}
}

func (c *FakeWalrusV1) Variables(namespace string) v1.VariableInterface {
	return &FakeVariables{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeWalrusV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
