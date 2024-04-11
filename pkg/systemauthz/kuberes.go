package systemauthz

import (
	"context"
	"fmt"

	rbac "k8s.io/api/rbac/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubeclientset/review"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// InitializedNamespaceName is the name indicates which Kubernetes Namespace storing system resources.
const InitializedNamespaceName = system.NamespaceName

const (
	// AnonymousClusterRoleName is the name of the Kubernetes ClusterRole for system anonymous.
	AnonymousClusterRoleName = "walrus-anonymous"
	// ViewerClusterRoleName is the name of the Kubernetes ClusterRole for system viewer.
	ViewerClusterRoleName = "walrus-viewer"
	// EditorClusterRoleName is the name of the Kubernetes ClusterRole for system editor.
	EditorClusterRoleName = "walrus-editor"
	// AdminClusterRoleName is the name of the Kubernetes ClusterRole for system administrator.
	AdminClusterRoleName = "walrus-admin"
)

// Initialize initializes Kubernetes resources for authorization.
//
// Initialize creates Kubernetes ClusterRole/ClusterRoleBinding/RoleBinding for system.
func Initialize(ctx context.Context, cli clientset.Interface) error {
	err := review.CanDoCreate(ctx,
		cli.AuthorizationV1().SelfSubjectAccessReviews(),
		review.Simples{
			{
				Group:    rbac.SchemeGroupVersion.Group,
				Version:  rbac.SchemeGroupVersion.Version,
				Resource: "clusterroles",
			},
			{
				Group:    rbac.SchemeGroupVersion.Group,
				Version:  rbac.SchemeGroupVersion.Version,
				Resource: "rolebindings",
			},
		},
		review.WithUpdateIfExisted(),
	)
	if err != nil {
		return err
	}

	crCli := cli.RbacV1().ClusterRoles()
	eCrs := []*rbac.ClusterRole{
		// Anonymous.
		{
			ObjectMeta: meta.ObjectMeta{
				Name: AnonymousClusterRoleName,
			},
			Rules: []rbac.PolicyRule{
				// Read limited resources include:
				// - Specific settings.
				{
					APIGroups: []string{
						walrus.GroupName,
					},
					Resources: []string{
						"settings",
					},
					ResourceNames: []string{
						"bootstrap-password-provision",
						"serve-url",
					},
					Verbs: []string{
						"get",
					},
				},
			},
		},
		// Viewer.
		{
			ObjectMeta: meta.ObjectMeta{
				Name: ViewerClusterRoleName,
			},
			Rules: []rbac.PolicyRule{
				// View all resources exclude:
				// - Subject Providers
				{
					APIGroups: []string{
						walrus.GroupName,
					},
					Resources: []string{
						"catalogs",
						"connectors",
						"environments",
						"fileexamples",
						"projects",
						"projects/subjects",
						"resources",
						"resources/components",
						"resourcedefinitions",
						"resourceruns",
						"settings",
						"subjects",
						// "subjectproviders", // NB(thxCode): Not included.
						"templates",
						"variables",
					},
					Verbs: []string{
						"get",
						"list",
						"watch",
					},
				},
				// Manage self Project.
				{
					APIGroups: []string{
						walrus.GroupName,
					},
					Resources: []string{
						"projects",
					},
					Verbs: []string{
						rbac.VerbAll,
					},
				},
			},
		},
		// Editor.
		{
			ObjectMeta: meta.ObjectMeta{
				Name: EditorClusterRoleName,
			},
			Rules: []rbac.PolicyRule{
				// Manage all resources exclude:
				// - Subject
				// - Subject Login
				// - Subject Token
				// - Subject Providers
				{
					APIGroups: []string{
						walrus.GroupName,
					},
					Resources: []string{
						"catalogs",
						"connectors",
						"environments",
						"resources",
						"resources/components",
						"resourcedefinitions",
						"templates",
						"variables",
					},
					Verbs: []string{
						rbac.VerbAll,
					},
				},
			},
		},
		// Admin.
		{
			ObjectMeta: meta.ObjectMeta{
				Name: AdminClusterRoleName,
			},
			Rules: []rbac.PolicyRule{
				// Manage all resources exclude:
				// - Subject Login
				// - Subject Token
				{
					APIGroups: []string{
						walrus.GroupName,
					},
					Resources: []string{
						"catalogs",
						"connectors",
						"environments",
						"fileexamples",
						"projects",
						"projects/subjects",
						"resources",
						"resources/components",
						"resourcedefinitions",
						"resourceruns",
						"settings",
						"subjects",
						"subjectproviders",
						"templates",
						"variables",
					},
					Verbs: []string{
						rbac.VerbAll,
					},
				},
			},
		},
	}
	for i := range eCrs {
		systemmeta.NoteResource(eCrs[i], "roles", nil)

		// Create.
		_, err = kubeclientset.Create(ctx, crCli, eCrs[i],
			kubeclientset.WithUpdateIfExisted(kubeclientset.NewRbacClusterRoleAlignFunc(eCrs[i])))
		if err != nil {
			return fmt.Errorf("install cluster role %q: %w", eCrs[i].Name, err)
		}
	}

	rbCli := cli.RbacV1().RoleBindings(InitializedNamespaceName)
	eRbs := []*rbac.RoleBinding{
		// Fro system anonymous.
		{
			ObjectMeta: meta.ObjectMeta{
				Namespace: InitializedNamespaceName,
				Name:      AnonymousClusterRoleName,
			},
			RoleRef: rbac.RoleRef{
				APIGroup: rbac.GroupName,
				Kind:     "ClusterRole",
				Name:     AnonymousClusterRoleName,
			},
			Subjects: []rbac.Subject{
				{
					APIGroup: rbac.GroupName,
					Kind:     rbac.GroupKind,
					Name:     "system:unauthenticated",
				},
			},
		},
		// For system user.
		{
			ObjectMeta: meta.ObjectMeta{
				Namespace: InitializedNamespaceName,
				Name:      ViewerClusterRoleName,
			},
			RoleRef: rbac.RoleRef{
				APIGroup: rbac.GroupName,
				Kind:     "ClusterRole",
				Name:     ViewerClusterRoleName,
			},
			Subjects: []rbac.Subject{
				{
					APIGroup: rbac.GroupName,
					Kind:     rbac.GroupKind,
					Name:     "system:authenticated",
				},
			},
		},
	}
	for i := range eRbs {
		// Create.
		_, err = kubeclientset.Create(ctx, rbCli, eRbs[i],
			kubeclientset.WithRecreateIfDuplicated(kubeclientset.NewRbacRoleBindingCompareFunc(eRbs[i])))
		if err != nil {
			return fmt.Errorf("install role binding %q: %w", eRbs[i].Name, err)
		}
	}

	return nil
}
