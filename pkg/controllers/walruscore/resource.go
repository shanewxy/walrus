package walruscore

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/controller"
)

// ResourceReconciler reconciles a v1.Resource object.
type ResourceReconciler struct {
	client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*ResourceReconciler)(nil)

func (r *ResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrllog.FromContext(ctx)
	obj := new(walruscore.Resource)

	// Fetch
	err := r.client.Get(ctx, req.NamespacedName, obj)
	if err != nil {
		logger.Error(err, "fetch resource")

		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

func (r *ResourceReconciler) SetupController(_ context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&walruscore.Resource{}).
		Complete(r)
}
