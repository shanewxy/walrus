//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	"bytes"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	rest "k8s.io/apiserver/pkg/registry/rest"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
)

type (
	WithStatusSubResource interface {
		v1.ObjectMetaAccessor
		runtime.Object
		CopyStatusTo(runtime.Object)
	}

	WithScaleSubResource interface {
		v1.ObjectMetaAccessor
		runtime.Object
		GetScale() *autoscalingv1.Scale // TODO: Main struct needs to implement this.
		SetScale(*autoscalingv1.Scale)  // TODO: Main struct needs to implement this.
	}
)

func GetAPIService(svc apiregistrationv1.ServiceReference, ca []byte) *apiregistrationv1.APIService {
	return &apiregistrationv1.APIService{
		TypeMeta: v1.TypeMeta{
			APIVersion: "apiregistration.k8s.io/v1",
			Kind:       "APIService",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "v1.walrus.seal.io",
		},
		Spec: apiregistrationv1.APIServiceSpec{
			Service:               svc.DeepCopy(),
			Group:                 "walrus.seal.io",
			Version:               "v1",
			InsecureSkipTLSVerify: true,
			CABundle:              bytes.Clone(ca),
			GroupPriorityMinimum:  100,
			VersionPriority:       100,
		},
	}
}

var _ rest.Scoper = (*Catalog)(nil)

func (*Catalog) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*Catalog)(nil)

func (*Catalog) Kind() string {
	return "Catalog"
}

var _ rest.SingularNameProvider = (*Catalog)(nil)

func (*Catalog) GetSingularName() string {
	return "catalog"
}

var _ rest.ShortNamesProvider = (*Catalog)(nil)

func (*Catalog) ShortNames() []string {
	return []string{}
}

var _ rest.CategoriesProvider = (*Catalog)(nil)

func (*Catalog) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ WithStatusSubResource = (*Catalog)(nil)

func (in *Catalog) CopyStatusTo(out runtime.Object) {
	out.(*Catalog).Status = in.Status
}

var _ rest.Scoper = (*Connector)(nil)

func (*Connector) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*Connector)(nil)

func (*Connector) Kind() string {
	return "Connector"
}

var _ rest.SingularNameProvider = (*Connector)(nil)

func (*Connector) GetSingularName() string {
	return "connector"
}

var _ rest.ShortNamesProvider = (*Connector)(nil)

func (*Connector) ShortNames() []string {
	return []string{
		"conn",
	}
}

var _ rest.CategoriesProvider = (*Connector)(nil)

func (*Connector) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ WithStatusSubResource = (*Connector)(nil)

func (in *Connector) CopyStatusTo(out runtime.Object) {
	out.(*Connector).Status = in.Status
}

var _ rest.Scoper = (*Environment)(nil)

func (*Environment) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*Environment)(nil)

func (*Environment) Kind() string {
	return "Environment"
}

var _ rest.SingularNameProvider = (*Environment)(nil)

func (*Environment) GetSingularName() string {
	return "environment"
}

var _ rest.ShortNamesProvider = (*Environment)(nil)

func (*Environment) ShortNames() []string {
	return []string{
		"env",
	}
}

var _ rest.CategoriesProvider = (*Environment)(nil)

func (*Environment) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*FileExample)(nil)

func (*FileExample) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*FileExample)(nil)

func (*FileExample) Kind() string {
	return "FileExample"
}

var _ rest.SingularNameProvider = (*FileExample)(nil)

func (*FileExample) GetSingularName() string {
	return "fileexample"
}

var _ rest.ShortNamesProvider = (*FileExample)(nil)

func (*FileExample) ShortNames() []string {
	return []string{}
}

var _ rest.CategoriesProvider = (*FileExample)(nil)

func (*FileExample) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*Project)(nil)

func (*Project) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*Project)(nil)

func (*Project) Kind() string {
	return "Project"
}

var _ rest.SingularNameProvider = (*Project)(nil)

func (*Project) GetSingularName() string {
	return "project"
}

var _ rest.ShortNamesProvider = (*Project)(nil)

func (*Project) ShortNames() []string {
	return []string{
		"proj",
	}
}

var _ rest.CategoriesProvider = (*Project)(nil)

func (*Project) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*ProjectSubjects)(nil)

func (*ProjectSubjects) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*ProjectSubjects)(nil)

func (*ProjectSubjects) Kind() string {
	return "ProjectSubjects"
}

var _ rest.SingularNameProvider = (*ProjectSubjects)(nil)

func (*ProjectSubjects) GetSingularName() string {
	return "projectsubjects"
}

var _ rest.ShortNamesProvider = (*ProjectSubjects)(nil)

func (*ProjectSubjects) ShortNames() []string {
	return []string{
		"projsub",
	}
}

var _ rest.CategoriesProvider = (*ProjectSubjects)(nil)

func (*ProjectSubjects) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*Resource)(nil)

func (*Resource) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*Resource)(nil)

func (*Resource) Kind() string {
	return "Resource"
}

var _ rest.SingularNameProvider = (*Resource)(nil)

func (*Resource) GetSingularName() string {
	return "resource"
}

var _ rest.ShortNamesProvider = (*Resource)(nil)

func (*Resource) ShortNames() []string {
	return []string{
		"res",
	}
}

var _ rest.CategoriesProvider = (*Resource)(nil)

func (*Resource) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ WithStatusSubResource = (*Resource)(nil)

func (in *Resource) CopyStatusTo(out runtime.Object) {
	out.(*Resource).Status = in.Status
}

var _ rest.Scoper = (*ResourceComponents)(nil)

func (*ResourceComponents) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*ResourceComponents)(nil)

func (*ResourceComponents) Kind() string {
	return "ResourceComponents"
}

var _ rest.SingularNameProvider = (*ResourceComponents)(nil)

func (*ResourceComponents) GetSingularName() string {
	return "resourcecomponents"
}

var _ rest.ShortNamesProvider = (*ResourceComponents)(nil)

func (*ResourceComponents) ShortNames() []string {
	return []string{
		"rescomp",
	}
}

var _ rest.CategoriesProvider = (*ResourceComponents)(nil)

func (*ResourceComponents) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*ResourceDefinition)(nil)

func (*ResourceDefinition) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*ResourceDefinition)(nil)

func (*ResourceDefinition) Kind() string {
	return "ResourceDefinition"
}

var _ rest.SingularNameProvider = (*ResourceDefinition)(nil)

func (*ResourceDefinition) GetSingularName() string {
	return "resourcedefinition"
}

var _ rest.ShortNamesProvider = (*ResourceDefinition)(nil)

func (*ResourceDefinition) ShortNames() []string {
	return []string{
		"resdef",
	}
}

var _ rest.CategoriesProvider = (*ResourceDefinition)(nil)

func (*ResourceDefinition) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ WithStatusSubResource = (*ResourceDefinition)(nil)

func (in *ResourceDefinition) CopyStatusTo(out runtime.Object) {
	out.(*ResourceDefinition).Status = in.Status
}

var _ rest.Scoper = (*ResourceRun)(nil)

func (*ResourceRun) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*ResourceRun)(nil)

func (*ResourceRun) Kind() string {
	return "ResourceRun"
}

var _ rest.SingularNameProvider = (*ResourceRun)(nil)

func (*ResourceRun) GetSingularName() string {
	return "resourcerun"
}

var _ rest.ShortNamesProvider = (*ResourceRun)(nil)

func (*ResourceRun) ShortNames() []string {
	return []string{
		"resrun",
	}
}

var _ rest.CategoriesProvider = (*ResourceRun)(nil)

func (*ResourceRun) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ WithStatusSubResource = (*ResourceRun)(nil)

func (in *ResourceRun) CopyStatusTo(out runtime.Object) {
	out.(*ResourceRun).Status = in.Status
}

var _ rest.Scoper = (*Setting)(nil)

func (*Setting) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*Setting)(nil)

func (*Setting) Kind() string {
	return "Setting"
}

var _ rest.SingularNameProvider = (*Setting)(nil)

func (*Setting) GetSingularName() string {
	return "setting"
}

var _ rest.ShortNamesProvider = (*Setting)(nil)

func (*Setting) ShortNames() []string {
	return []string{
		"set",
	}
}

var _ rest.CategoriesProvider = (*Setting)(nil)

func (*Setting) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*Subject)(nil)

func (*Subject) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*Subject)(nil)

func (*Subject) Kind() string {
	return "Subject"
}

var _ rest.SingularNameProvider = (*Subject)(nil)

func (*Subject) GetSingularName() string {
	return "subject"
}

var _ rest.ShortNamesProvider = (*Subject)(nil)

func (*Subject) ShortNames() []string {
	return []string{
		"subj",
	}
}

var _ rest.CategoriesProvider = (*Subject)(nil)

func (*Subject) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*SubjectLogin)(nil)

func (*SubjectLogin) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*SubjectLogin)(nil)

func (*SubjectLogin) Kind() string {
	return "SubjectLogin"
}

var _ rest.SingularNameProvider = (*SubjectLogin)(nil)

func (*SubjectLogin) GetSingularName() string {
	return "subjectlogin"
}

var _ rest.ShortNamesProvider = (*SubjectLogin)(nil)

func (*SubjectLogin) ShortNames() []string {
	return []string{}
}

var _ rest.CategoriesProvider = (*SubjectLogin)(nil)

func (*SubjectLogin) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*SubjectProvider)(nil)

func (*SubjectProvider) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*SubjectProvider)(nil)

func (*SubjectProvider) Kind() string {
	return "SubjectProvider"
}

var _ rest.SingularNameProvider = (*SubjectProvider)(nil)

func (*SubjectProvider) GetSingularName() string {
	return "subjectprovider"
}

var _ rest.ShortNamesProvider = (*SubjectProvider)(nil)

func (*SubjectProvider) ShortNames() []string {
	return []string{
		"subjprov",
	}
}

var _ rest.CategoriesProvider = (*SubjectProvider)(nil)

func (*SubjectProvider) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*SubjectToken)(nil)

func (*SubjectToken) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*SubjectToken)(nil)

func (*SubjectToken) Kind() string {
	return "SubjectToken"
}

var _ rest.SingularNameProvider = (*SubjectToken)(nil)

func (*SubjectToken) GetSingularName() string {
	return "subjecttoken"
}

var _ rest.ShortNamesProvider = (*SubjectToken)(nil)

func (*SubjectToken) ShortNames() []string {
	return []string{}
}

var _ rest.CategoriesProvider = (*SubjectToken)(nil)

func (*SubjectToken) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ rest.Scoper = (*Template)(nil)

func (*Template) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*Template)(nil)

func (*Template) Kind() string {
	return "Template"
}

var _ rest.SingularNameProvider = (*Template)(nil)

func (*Template) GetSingularName() string {
	return "template"
}

var _ rest.ShortNamesProvider = (*Template)(nil)

func (*Template) ShortNames() []string {
	return []string{
		"tpl",
	}
}

var _ rest.CategoriesProvider = (*Template)(nil)

func (*Template) Categories() []string {
	return []string{
		"walrus",
	}
}

var _ WithStatusSubResource = (*Template)(nil)

func (in *Template) CopyStatusTo(out runtime.Object) {
	out.(*Template).Status = in.Status
}

var _ rest.Scoper = (*Variable)(nil)

func (*Variable) NamespaceScoped() bool {
	return true
}

var _ rest.KindProvider = (*Variable)(nil)

func (*Variable) Kind() string {
	return "Variable"
}

var _ rest.SingularNameProvider = (*Variable)(nil)

func (*Variable) GetSingularName() string {
	return "variable"
}

var _ rest.ShortNamesProvider = (*Variable)(nil)

func (*Variable) ShortNames() []string {
	return []string{
		"var",
	}
}

var _ rest.CategoriesProvider = (*Variable)(nil)

func (*Variable) Categories() []string {
	return []string{
		"walrus",
	}
}
