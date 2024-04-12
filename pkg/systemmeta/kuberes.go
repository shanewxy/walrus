package systemmeta

import (
	"context"
	"errors"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/system"
)

type (
	// NamespaceKind represents the kind of Walrus namespace.
	NamespaceKind string

	// Namespace represents a Walrus namespace.
	Namespace interface {
		// Kind returns the kind of the Walrus namespace.
		Kind() NamespaceKind
		// RetrieveOwner returns the Walrus namespace of the owner.
		RetrieveOwner(ctx context.Context) (Namespace, error)
		// OwnerName returns the name of the owner.
		OwnerName() string
		// Metadata returns the metadata of the Walrus namespace.
		Metadata() meta.PartialObjectMetadata
	}
)

const (
	// NamespaceKindSystem represents the system namespace.
	NamespaceKindSystem NamespaceKind = "System"
	// NamespaceKindProject represents the project namespace.
	NamespaceKindProject NamespaceKind = "Project"
	// NamespaceKindEnvironment represents the environment namespace.
	NamespaceKindEnvironment NamespaceKind = "Environment"
)

type _Namespace struct {
	ownerName string
	metadata  meta.PartialObjectMetadata
}

func (n _Namespace) Kind() NamespaceKind {
	switch n.metadata.Kind {
	case "Project":
		return NamespaceKindProject
	case "Environment":
		return NamespaceKindEnvironment
	}
	return NamespaceKindSystem
}

func (n _Namespace) RetrieveOwner(ctx context.Context) (Namespace, error) {
	if n.Kind() == NamespaceKindSystem {
		return nil, errors.New("system namespace has no owner")
	}

	return ReflectNamespace(ctx, system.LoopbackCtrlClient.Get(), n.ownerName)
}

func (n _Namespace) OwnerName() string {
	return n.ownerName
}

func (n _Namespace) Metadata() meta.PartialObjectMetadata {
	return n.metadata
}

// ReflectNamespace reflects a core namespace name to a Walrus namespace.
func ReflectNamespace(ctx context.Context, cli ctrlcli.Client, namespaceName string) (Namespace, error) {
	metadata := &meta.PartialObjectMetadata{
		TypeMeta: meta.TypeMeta{
			APIVersion: core.SchemeGroupVersion.String(),
			Kind:       "Namespace",
		},
		ObjectMeta: meta.ObjectMeta{
			Name: namespaceName,
		},
	}
	err := cli.Get(ctx, ctrlcli.ObjectKeyFromObject(metadata), metadata)
	if err != nil {
		return nil, err
	}

	if namespaceName == system.NamespaceName {
		return _Namespace{metadata: *metadata}, nil
	}

	resType := DescribeResourceType(metadata)
	switch resType {
	case "projects":
		metadata.APIVersion = walrus.SchemeGroupVersion.String()
		metadata.Kind = "Project"
		return _Namespace{ownerName: system.NamespaceName, metadata: *metadata}, nil
	case "environments":
		proj := kubemeta.GetOwnerRefOfNoCopy(metadata, walrus.SchemeGroupVersionKind("Project"))
		if proj == nil {
			return nil, errors.New("invalid environment: incomplete project")
		}
		metadata.APIVersion = walrus.SchemeGroupVersion.String()
		metadata.Kind = "Environment"
		return _Namespace{ownerName: proj.Name, metadata: *metadata}, nil
	}
	return nil, errors.New("not found")
}
