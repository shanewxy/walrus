package walrus

import (
	"context"
	"fmt"

	"golang.org/x/exp/maps"
	rbac "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/extensionapi"
	"github.com/seal-io/walrus/pkg/systemauthz"
	"github.com/seal-io/walrus/pkg/systemkuberes"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// ProjectSubjectsHandler is a handler for v1.ProjectSubjects objects,
// which is a subresource of v1.Project objects.
//
// ProjectSubjectsHandler maps the rbac RoleBinding objects to the walrus v1.ProjectSubjects objects.
type ProjectSubjectsHandler struct {
	extensionapi.ObjectInfo
	extensionapi.GetOperation
	extensionapi.UpdateOperation

	Client    ctrlcli.Client
	APIReader ctrlcli.Reader
}

func newProjectSubjectsHandler(opts extensionapi.SetupOptions) *ProjectSubjectsHandler {
	h := &ProjectSubjectsHandler{}

	// As storage.
	h.ObjectInfo = &walrus.ProjectSubjects{}
	h.GetOperation = extensionapi.WithGet(h)
	h.UpdateOperation = extensionapi.WithUpdate(h)

	// Set client.
	h.Client = opts.Manager.GetClient()
	h.APIReader = opts.Manager.GetAPIReader()

	return h
}

var (
	_ rest.Storage = (*ProjectSubjectsHandler)(nil)
	_ rest.Getter  = (*ProjectSubjectsHandler)(nil)
	_ rest.Updater = (*ProjectSubjectsHandler)(nil)
	_ rest.Patcher = (*ProjectSubjectsHandler)(nil)
)

func (h *ProjectSubjectsHandler) New() runtime.Object {
	return &walrus.ProjectSubjects{}
}

func (h *ProjectSubjectsHandler) Destroy() {}

func (h *ProjectSubjectsHandler) OnGet(ctx context.Context, key types.NamespacedName, opts ctrlcli.GetOptions) (runtime.Object, error) {
	// Validate.
	if key.Namespace != systemkuberes.SystemNamespaceName {
		return nil, kerrors.NewNotFound(walrus.SchemeResource("projectsubjects"), key.Name)
	}

	// List.
	rbList := new(rbac.RoleBindingList)
	err := h.APIReader.List(ctx, rbList,
		ctrlcli.InNamespace(key.Name),
		ctrlcli.MatchingLabelsSelector{
			Selector: systemmeta.GetResourcesLabelSelectorOfType("rolebindings"),
		})
	if err != nil {
		return nil, kerrors.NewInternalError(err)
	}
	rbList = systemmeta.FilterResourceListByNotes(rbList, "project", key.String())

	// Convert.
	psbjs := convertProjectSubjectsFromRoleBindingList(rbList)
	if psbjs == nil {
		return nil, kerrors.NewNotFound(walrus.SchemeResource("projectsubjects"), key.Name)
	}

	// Get and refill.
	proj := new(walrus.Project)
	err = h.Client.Get(ctx, key, proj, &opts)
	if err != nil {
		return nil, kerrors.NewInternalError(err)
	}
	psbjs.ObjectMeta = proj.ObjectMeta

	return psbjs, nil
}

func (h *ProjectSubjectsHandler) OnUpdate(ctx context.Context, obj, objOld runtime.Object, _ ctrlcli.UpdateOptions) (runtime.Object, error) {
	psbjs, psbjsOld := obj.(*walrus.ProjectSubjects), objOld.(*walrus.ProjectSubjects)

	// Validate.
	{
		var errs field.ErrorList
		for i, psbj := range psbjs.Items {
			err := h.Client.Get(ctx, psbj.ToNamespacedName(), new(walrus.Subject))
			if err != nil {
				errs = append(errs, field.Invalid(
					field.NewPath(fmt.Sprintf("items[%d]", i)), psbj.SubjectReference, err.Error()),
				)
			}
			if err := psbj.Role.Validate(); err != nil {
				errs = append(errs, field.Invalid(
					field.NewPath(fmt.Sprintf("items[%d].role", i)), psbj.Role, err.Error()))
			}
		}
		if len(errs) > 0 {
			return nil, kerrors.NewInvalid(walrus.SchemeKind("projectsubjects"), psbjs.Name, errs)
		}
	}

	// Figure out delta.
	psbjsReverseIndex := make(map[walrus.ProjectSubject]int)
	for i := range psbjs.Items {
		psbjsReverseIndex[psbjs.Items[i]] = i
	}
	psbjsOldSet := make(map[walrus.ProjectSubject]sets.Empty)
	for i := range psbjsOld.Items {
		psbjsOldSet[psbjsOld.Items[i]] = sets.Empty{}
	}
	for psbj := range psbjsReverseIndex {
		// Delete the one exists in both of the new set and old set,
		// then the remaining items of the new set are need to create,
		// and the remaining items of the old set are need to delete.
		if _, existed := psbjsOldSet[psbj]; existed {
			delete(psbjsReverseIndex, psbj)
			delete(psbjsOldSet, psbj)
		}
	}

	// NB(thxCode): we revoke the old permission from the Project namespace first,
	// then ProjectSubjectAuthzReconciler will take care of revoking the old permission from the Environment namespace.
	psbjsOld.Items = maps.Keys(psbjsOldSet)
	err := systemauthz.RevokeProjectSubjects(ctx, h.Client, psbjsOld)
	if err != nil {
		return nil, kerrors.NewInternalError(fmt.Errorf("revoke project subject: %w", err))
	}

	// NB(thxCode): we grant the new permission to the Project namespace first,
	// then ProjectSubjectAuthzReconciler will take care of granting the new permission to the Environment namespace.
	psbjs.Items = maps.Keys(psbjsReverseIndex)
	err = systemauthz.GrantProjectSubjects(ctx, h.Client, psbjs)
	if err != nil {
		return nil, kerrors.NewInternalError(fmt.Errorf("grant project subject: %w", err))
	}

	// Get.
	return h.OnGet(ctx, ctrlcli.ObjectKeyFromObject(psbjs),
		ctrlcli.GetOptions{
			Raw: &meta.GetOptions{ResourceVersion: "0"},
		})
}

func convertProjectSubjectFromRoleBinding(rb *rbac.RoleBinding) *walrus.ProjectSubject {
	if rb == nil || rb.RoleRef.Kind != "ClusterRole" {
		return nil
	}

	r := systemauthz.ConvertProjectRoleFromClusterRoleName(rb.RoleRef.Name)
	if r.Validate() != nil {
		return nil
	}

	var ns, n string
	for _, subj := range rb.Subjects {
		if subj.Kind != rbac.ServiceAccountKind {
			continue
		}
		ns = subj.Namespace
		if ns == "" {
			continue
		}
		n = systemauthz.ConvertSubjectNameFromServiceAccountName(subj.Name)
		if n == "" {
			continue
		}
	}
	if ns == "" || n == "" {
		return nil
	}

	psbj := &walrus.ProjectSubject{
		SubjectReference: walrus.SubjectReference{
			Namespace: ns,
			Name:      n,
		},
		Role: r,
	}
	return psbj
}

func convertProjectSubjectsFromRoleBindingList(rbList *rbac.RoleBindingList) *walrus.ProjectSubjects {
	if rbList == nil {
		return nil
	}

	psbjs := &walrus.ProjectSubjects{
		Items: make([]walrus.ProjectSubject, 0, len(rbList.Items)),
	}

	for i := range rbList.Items {
		psbj := convertProjectSubjectFromRoleBinding(&rbList.Items[i])
		if psbj == nil {
			continue
		}
		psbjs.Items = append(psbjs.Items, *psbj)
	}

	return psbjs
}
