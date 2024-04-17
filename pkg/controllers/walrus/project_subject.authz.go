package walrus

import (
	"context"
	"fmt"
	"time"

	rbac "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlbuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlevent "sigs.k8s.io/controller-runtime/pkg/event"
	ctrlhandler "sigs.k8s.io/controller-runtime/pkg/handler"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/systemauthz"
	"github.com/seal-io/walrus/pkg/systemkuberes"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// ProjectSubjectAuthzReconciler reconciles a rbac RoleBinding object below the Project namespace.
//
// ProjectSubjectAuthzReconciler works like a dispatcher,
// which listens to the role bindings created under the Project namespace.
// And then, it copies the role bindings to the related Environments.
//
// ProjectSubjectAuthzReconciler will be requeue if a new Environment of the Project is created,
// so that we will not miss any assigned ProjectSubject.
type ProjectSubjectAuthzReconciler struct {
	Client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*ProjectSubjectAuthzReconciler)(nil)

func (r *ProjectSubjectAuthzReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrllog.FromContext(ctx)

	// Fetch.
	rb := new(rbac.RoleBinding)
	err := r.Client.Get(ctx, req.NamespacedName, rb)
	if err != nil {
		logger.Error(err, "fetch role binding")
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	// Revoke if deleted.
	if rb.DeletionTimestamp != nil {
		// Return if already unlocked.
		if systemmeta.Unlock(rb) {
			return ctrl.Result{}, nil
		}

		// List related role bindings.
		rbList := new(rbac.RoleBindingList)
		err = r.Client.List(ctx, rbList,
			ctrlcli.MatchingFields{
				"rolebindings[scope=environment].name": rb.Name,
			})
		if err != nil {
			logger.Error(err, "list related role bindings")
			return ctrl.Result{}, err
		}

		// Delete related role bindings.
		for i := range rbList.Items {
			eRb := &rbList.Items[i]
			err = r.Client.Delete(ctx, eRb)
			if err != nil && !kerrors.IsNotFound(err) {
				logger.Error(err, "delete related role binding", "rolebinding", kubemeta.GetNamespacedNameKey(eRb))
				return ctrl.Result{}, err
			}
		}

		// Unlock.
		_, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, rb)
		if err != nil {
			logger.Error(err, "unlock role binding")
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}

		return ctrl.Result{}, nil
	}

	// Lock if not.
	if !systemmeta.Lock(rb) {
		rb, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, rb)
		if err != nil {
			logger.Error(err, "lock role binding")
			return ctrl.Result{}, err
		}
	}

	// Get subject.
	subj := new(walrus.Subject)
	{
		var ns, n string
		for _, s := range rb.Subjects {
			if s.Kind != rbac.ServiceAccountKind {
				continue
			}
			ns = s.Namespace
			if ns == "" {
				continue
			}
			n = systemauthz.ConvertSubjectNameFromServiceAccountName(s.Name)
			if n == "" {
				continue
			}
			break
		}
		if ns == "" || n == "" {
			return ctrl.Result{}, nil
		}

		err = r.Client.Get(ctx, types.NamespacedName{Namespace: ns, Name: n}, subj)
		if err != nil {
			// Revoke if the subject is not found
			err = kubeclientset.DeleteWithCtrlClient(ctx, r.Client, rb)
			if err != nil {
				logger.Error(err, "delete role binding of not found subject")
			}
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}
	}

	// List related environments.
	projKey := kubemeta.ParseNamespacedNameKey(systemmeta.DescribeResourceNote(rb, "project"))
	envList := new(walrus.EnvironmentList)
	err = r.Client.List(ctx, envList,
		ctrlcli.InNamespace(projKey.Name))
	if err != nil {
		logger.Error(err, "list related environments")
		return ctrl.Result{}, err
	}

	// Grant: copy to related environments.
	for i := range envList.Items {
		env := &envList.Items[i]
		if env.DeletionTimestamp != nil {
			continue
		}

		// Degrade the project role if the subject is a viewer but the environment is production.
		eRb := &rbac.RoleBinding{
			ObjectMeta: meta.ObjectMeta{
				Namespace: env.Name,
				Name:      rb.Name,
			},
			RoleRef:  rb.RoleRef,
			Subjects: rb.Subjects,
		}
		if env.Spec.Type == walruscore.EnvironmentTypeProduction && subj.Spec.Role == walrus.SubjectRoleUser {
			eRb.RoleRef.Name = systemauthz.ConvertClusterRoleNameFromProjectRole(walrus.ProjectRoleViewer)
		}
		systemmeta.NoteResource(eRb, "rolebindings", map[string]string{
			"scope":       "environment",
			"environment": kubemeta.GetNamespacedNameKey(env),
			"subject":     kubemeta.GetNamespacedNameKey(subj),
		})

		// Create.
		_, err = kubeclientset.CreateWithCtrlClient(ctx, r.Client, eRb,
			kubeclientset.WithRecreateIfDuplicated(kubeclientset.NewRbacRoleBindingCompareFunc(eRb)))
		if err != nil {
			return ctrl.Result{RequeueAfter: time.Second}, fmt.Errorf("create role binding: %w", err)
		}
	}

	return ctrl.Result{}, nil
}

func (r *ProjectSubjectAuthzReconciler) SetupController(ctx context.Context, opts controller.SetupOptions) error {
	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err := fi.IndexField(ctx, &rbac.RoleBinding{}, "rolebindings[scope=environment].name",
		func(obj ctrlcli.Object) []string {
			if obj == nil {
				return nil
			}
			resType, notes := systemmeta.DescribeResource(obj)
			if resType != "rolebindings" {
				return nil
			}
			if notes["scope"] != "environment" {
				return nil
			}
			return []string{obj.GetName()}
		})
	if err != nil {
		return fmt.Errorf("index role binding 'rolebindings[scope=environment].name': %w", err)
	}

	r.Client = opts.Manager.GetClient()

	// Filter out non-project role bindings.
	rbFilter := ctrlpredicate.NewPredicateFuncs(func(obj ctrlcli.Object) bool {
		resType, notes := systemmeta.DescribeResource(obj)
		return resType == "rolebindings" &&
			notes["scope"] == "project" &&
			kubemeta.ContainsNameInNamespacedNameKey(obj.GetNamespace(), notes["project"])
	})

	// Filter out creating environment.
	envFilter := ctrlpredicate.Not(ctrlpredicate.Funcs{
		CreateFunc: func(_ ctrlevent.CreateEvent) bool { return false },
	})

	return ctrl.NewControllerManagedBy(opts.Manager).
		Named("project_subject.authz").
		For(
			// Focus on the role bindings under the Project namespace.
			&rbac.RoleBinding{},
			ctrlbuilder.WithPredicates(rbFilter),
		).
		Watches(
			// Requeue when creating an Environment.
			&meta.PartialObjectMetadata{
				TypeMeta: meta.TypeMeta{
					APIVersion: walrus.SchemeGroupVersion.String(),
					Kind:       "Environment",
				},
			},
			ctrlhandler.EnqueueRequestsFromMapFunc(r.findObjectsWhenEnvironmentCreating),
			ctrlbuilder.WithPredicates(envFilter),
		).
		Complete(r)
}

func (r *ProjectSubjectAuthzReconciler) findObjectsWhenEnvironmentCreating(ctx context.Context, env ctrlcli.Object) []ctrlreconcile.Request {
	logger := ctrllog.FromContext(ctx)

	projSubjs := new(walrus.ProjectSubjects)
	{
		proj := &walrus.Project{
			ObjectMeta: meta.ObjectMeta{
				Namespace: systemkuberes.SystemNamespaceName,
				Name:      env.GetNamespace(),
			},
		}
		err := r.Client.Get(ctx, ctrlcli.ObjectKeyFromObject(proj), proj)
		if err != nil {
			logger.Error(err, "get project")
			return []ctrlreconcile.Request{}
		}
		err = r.Client.SubResource("subjects").Get(ctx, proj, projSubjs)
		if err != nil {
			logger.Error(err, "get project subjects")
			return []ctrlreconcile.Request{}
		}
	}

	reqs := make([]ctrlreconcile.Request, len(projSubjs.Items))
	for i, item := range projSubjs.Items {
		reqs[i] = ctrlreconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: env.GetNamespace(),
				Name:      systemauthz.GetProjectSubjectRoleBindingName(ptr.To(item.SubjectReference)),
			},
		}
	}
	return reqs
}
