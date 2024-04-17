package walruscore

import (
	"context"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlbuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/resourcehandlers"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// ConnectorReconciler reconciles a v1.Connector object.
type ConnectorReconciler struct {
	client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*ConnectorReconciler)(nil)

func (r *ConnectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrllog.FromContext(ctx)

	obj := new(walruscore.Connector)
	err := r.client.Get(ctx, req.NamespacedName, obj)
	if err != nil {
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	if obj.DeletionTimestamp != nil {
		// Return if already unlocked.
		if systemmeta.Unlock(obj) {
			return ctrl.Result{}, nil
		}

		sec := &core.Secret{
			ObjectMeta: meta.ObjectMeta{
				Namespace: obj.Namespace,
				Name:      obj.Spec.SecretName,
			},
		}

		err = r.client.Get(ctx, ctrlcli.ObjectKeyFromObject(sec), sec)
		if err != nil {
			return ctrl.Result{}, err
		}

		if sec.DeletionTimestamp == nil {
			err = r.client.Delete(ctx, sec)
			if err != nil {
				return ctrl.Result{}, err
			}
		}

		// Unlock.
		_, err = kubeclientset.UpdateWithCtrlClient(ctx, r.client, obj)
		if err != nil {
			logger.Error(err, "failed to unlock connector", "connector", obj)
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	if apistatus.ConnectorConditionConnected.IsTrue(obj) || apistatus.ConnectorConditionConnected.IsFalse(obj) {
		return ctrl.Result{}, nil
	}

	// Lock if not.
	if !systemmeta.Lock(obj) {
		obj, err = kubeclientset.UpdateWithCtrlClient(ctx, r.client, obj)
		if err != nil {
			logger.Error(err, "failed to lock connector", "connector", obj)
			return ctrl.Result{}, err
		}
	}

	rh, err := resourcehandlers.Get(ctx, resourcehandler.CreateOptions{
		Connector: *obj,
	})
	if err != nil {
		logger.Error(err, "fetch resource handler")
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	err = rh.IsConnected(ctx)
	if err != nil {
		apistatus.ConnectorConditionConnected.False(obj, apistatus.ConnectorConditionReasonDisconnected, err.Error())
	} else {
		apistatus.ConnectorConditionConnected.True(obj, "", "")
		obj.Status.Project = systemmeta.GetProjectName(obj.Namespace)
	}

	// Update status.
	{
		err = r.updateStatus(ctx, obj)
		if err != nil {
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

// updateStatus updates the status of the given connector.
func (r *ConnectorReconciler) updateStatus(ctx context.Context, obj *walruscore.Connector) error {
	existed := new(walruscore.Connector)
	err := r.client.Get(ctx, ctrlcli.ObjectKeyFromObject(obj), existed)
	if err != nil {
		return err
	}

	existed.Status = obj.Status
	existed.Status.ConditionSummary = *apistatus.WalkConnector(&existed.Status.StatusDescriptor)

	return r.client.Status().Update(ctx, existed)
}

func (r *ConnectorReconciler) SetupController(_ context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&walruscore.Connector{},
			ctrlbuilder.WithPredicates(ctrlpredicate.GenerationChangedPredicate{}),
		).
		Complete(r)
}
