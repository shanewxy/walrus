package systemauthz

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/seal-io/utils/stringx"
	rbac "k8s.io/api/rbac/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	authnuser "k8s.io/apiserver/pkg/authentication/user"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/systemkuberes"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// ConvertClusterRoleNameFromProjectRole converts the cluster role name from the project subject role.
func ConvertClusterRoleNameFromProjectRole(role walrus.ProjectRole) (clusterRoleName string) {
	switch role {
	case walrus.ProjectRoleOwner:
		return AdminClusterRoleName
	case walrus.ProjectRoleMember:
		return EditorClusterRoleName
	default:
		return ViewerClusterRoleName
	}
}

// ConvertProjectRoleFromClusterRoleName converts the project role from the cluster role name.
//
// If the cluster role name is not recognized, it returns an empty string.
func ConvertProjectRoleFromClusterRoleName(clusterRoleName string) (role walrus.ProjectRole) {
	switch clusterRoleName {
	case AdminClusterRoleName:
		return walrus.ProjectRoleOwner
	case EditorClusterRoleName:
		return walrus.ProjectRoleMember
	case ViewerClusterRoleName:
		return walrus.ProjectRoleViewer
	}
	return ""
}

// GrantProjectSubjects (re)grants the given project role to the corresponding subjects.
func GrantProjectSubjects(ctx context.Context, cli ctrlcli.Client, projSubjs *walrus.ProjectSubjects) error {
	if projSubjs == nil || len(projSubjs.Items) == 0 {
		return nil
	}

	for i := range projSubjs.Items {
		item := &projSubjs.Items[i]

		eRb := &rbac.RoleBinding{
			ObjectMeta: meta.ObjectMeta{
				Namespace: projSubjs.Name,
				Name:      GetProjectSubjectRoleBindingName(&item.SubjectReference),
			},
			RoleRef: rbac.RoleRef{
				APIGroup: rbac.GroupName,
				Kind:     "ClusterRole",
				Name:     ConvertClusterRoleNameFromProjectRole(item.Role),
			},
			Subjects: []rbac.Subject{
				{
					Kind:      rbac.ServiceAccountKind,
					Namespace: item.Namespace,
					Name:      ConvertServiceAccountNameFromSubjectName(item.Name),
				},
				{
					APIGroup: rbac.GroupName,
					Kind:     rbac.UserKind,
					Name:     ConvertImpersonateUserFromSubjectName(item.Namespace, item.Name),
				},
			},
		}
		systemmeta.NoteResource(eRb, "rolebindings", map[string]string{
			"scope":   "project",
			"project": kubemeta.GetNamespacedNameKey(projSubjs),
			"subject": kubemeta.GetNamespacedNameKey(item.SubjectReference.ToNamespacedName()),
		})

		// Create.
		_, err := kubeclientset.CreateWithCtrlClient(ctx, cli, eRb,
			kubeclientset.WithRecreateIfDuplicated(kubeclientset.NewRbacRoleBindingCompareFunc(eRb)))
		if err != nil {
			return fmt.Errorf("create role binding: %w", err)
		}
	}

	return nil
}

// RevokeProjectSubjects revokes the project role from the corresponding subjects.
func RevokeProjectSubjects(ctx context.Context, cli ctrlcli.Client, projSubjs *walrus.ProjectSubjects) error {
	if projSubjs == nil || len(projSubjs.Items) == 0 {
		return nil
	}

	for i := range projSubjs.Items {
		item := &projSubjs.Items[i]

		eRb := &rbac.RoleBinding{
			ObjectMeta: meta.ObjectMeta{
				Namespace: projSubjs.Name,
				Name:      GetProjectSubjectRoleBindingName(&item.SubjectReference),
			},
		}

		// Delete.
		err := kubeclientset.DeleteWithCtrlClient(ctx, cli, eRb)
		if err != nil {
			return fmt.Errorf("delete role binding: %w", err)
		}
	}

	return nil
}

// GrantProjectSubjectRole (re)grants the given project role for the request user.
func GrantProjectSubjectRole(ctx context.Context, cli ctrlcli.Client, proj *walrus.Project, role walrus.ProjectRole) error {
	ui, ok := genericapirequest.UserFrom(ctx)
	if !ok {
		return errors.New("request user not found")
	}

	// Don't bind the walrus admin, system:admin user or system:master group.
	{
		if un, ug := ui.GetName(), ui.GetGroups(); un == "system:admin" || slices.Contains(ug, "system:master") {
			return nil
		}
		if ns, n, ok := ConvertSubjectNamesFromAuthnUser(ui); ok && ns == systemkuberes.SystemNamespaceName && n == systemkuberes.AdminSubjectName {
			return nil
		}
	}

	return GrantProjectSubjectRoleFor(ctx, cli, proj, role, ui)
}

// GrantProjectSubjectRoleFor (re)grants the given for the specified user.
func GrantProjectSubjectRoleFor(ctx context.Context, cli ctrlcli.Client, proj *walrus.Project, role walrus.ProjectRole, user authnuser.Info) error { // nolint:lll
	// Validate.
	if proj == nil || proj.Name == "" {
		return errors.New("empty project")
	}
	if err := role.Validate(); err != nil {
		return err
	}

	// Convert.
	var subjRef walrus.SubjectReference
	{
		subjNamespace, subjName, ok := ConvertSubjectNamesFromAuthnUser(user)
		if !ok {
			return errors.New("incomplete user")
		}
		subjRef = walrus.SubjectReference{
			Namespace: subjNamespace,
			Name:      subjName,
		}
	}

	eRb := &rbac.RoleBinding{
		ObjectMeta: meta.ObjectMeta{
			Namespace: proj.Name,
			Name:      GetProjectSubjectRoleBindingName(&subjRef),
		},
		RoleRef: rbac.RoleRef{
			APIGroup: rbac.GroupName,
			Kind:     "ClusterRole",
			Name:     ConvertClusterRoleNameFromProjectRole(role),
		},
		Subjects: []rbac.Subject{
			{
				Kind:      rbac.ServiceAccountKind,
				Namespace: subjRef.Namespace,
				Name:      ConvertServiceAccountNameFromSubjectName(subjRef.Name),
			},
			{
				APIGroup: rbac.GroupName,
				Kind:     rbac.UserKind,
				Name:     ConvertImpersonateUserFromSubjectName(subjRef.Namespace, subjRef.Name),
			},
		},
	}
	systemmeta.NoteResource(eRb, "rolebindings", map[string]string{
		"scope":   "project",
		"project": kubemeta.GetNamespacedNameKey(proj),
		"subject": kubemeta.GetNamespacedNameKey(subjRef.ToNamespacedName()),
	})

	// Create.
	_, err := kubeclientset.CreateWithCtrlClient(ctx, cli, eRb,
		kubeclientset.WithRecreateIfDuplicated(kubeclientset.NewRbacRoleBindingCompareFunc(eRb)))
	if err != nil {
		return fmt.Errorf("create role binding: %w", err)
	}
	return nil
}

// GetProjectSubjectRoleBindingName returns the role binding name for the project subject.
func GetProjectSubjectRoleBindingName(subj *walrus.SubjectReference) string {
	return fmt.Sprintf("walrus-project-subject-%s",
		stringx.SumByFNV64a(subj.Namespace, subj.Name))
}
