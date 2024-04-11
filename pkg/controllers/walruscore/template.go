package walruscore

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/kubemeta"
	tmplfetcher "github.com/seal-io/walrus/pkg/templates/fetcher"
)

// TemplateReconciler reconciles a v1.Template object.
type TemplateReconciler struct {
	client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*TemplateReconciler)(nil)

// Reconcile reconciles the template.
func (r *TemplateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	// Fetch.
	obj := new(walruscore.Template)
	err = r.client.Get(ctx, req.NamespacedName, obj)
	if err != nil {
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	// Skip.
	{
		// Skip deletion.
		if obj.DeletionTimestamp != nil {
			return ctrl.Result{}, nil
		}

		// Skip template already reconciling.
		if apistatus.TemplateConditionReady.IsUnknown(obj) && !apistatus.TemplateConditionRefresh.IsUnknown(obj) {
			return ctrl.Result{}, nil
		}

		// Skip template finished reconciling.
		if (apistatus.TemplateConditionReady.IsTrue(obj) || apistatus.TemplateConditionReady.IsFalse(obj)) &&
			!apistatus.TemplateConditionRefresh.IsUnknown(obj) {
			return ctrl.Result{}, nil
		}

		// Skip template own by catalog and without manually refresh.
		if kubemeta.IsControlledByGVK(obj, walruscore.SchemeGroupVersionKind("Catalog")) &&
			!apistatus.TemplateConditionRefresh.IsUnknown(obj) {
			return ctrl.Result{}, nil
		}
	}

	// Initialize status.
	{
		apistatus.TemplateConditionReady.Unknown(obj, apistatus.TemplateConditionReasonPreparing, "Preparing")
		if apistatus.TemplateConditionRefresh.IsUnknown(obj) {
			apistatus.TemplateConditionRefresh.False(obj, "", "")
		}
		err = r.updateStatus(ctx, obj)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	defer func() {
		if err != nil {
			apistatus.TemplateConditionReady.False(obj, apistatus.TemplateConditionReasonError, err.Error())
			_ = r.updateStatus(ctx, obj)
		}
	}()

	// Reconcile and update status.
	{
		_, err = tmplfetcher.Fetch(ctx, obj)
		return ctrl.Result{}, err
	}
}

// updateStatus updates the status of the template.
func (r *TemplateReconciler) updateStatus(ctx context.Context, obj *walruscore.Template) error {
	existed := new(walruscore.Template)
	err := r.client.Get(ctx, ctrlcli.ObjectKeyFromObject(obj), existed)
	if err != nil {
		return err
	}

	existed.Status = obj.Status
	existed.Status.ConditionSummary = *apistatus.WalkTemplate(&existed.Status.StatusDescriptor)
	err = r.client.Status().Update(ctx, existed)
	return err
}

// SetupController sets up the controller.
func (r *TemplateReconciler) SetupController(_ context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&walruscore.Template{}).
		Complete(r)
}
