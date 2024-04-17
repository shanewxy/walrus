package walruscore

import (
	"context"
	"fmt"
	"strings"
	"time"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlbuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// ConnectorBindingReconciler reconciles a v1.ConnectorBinding object.
type ConnectorBindingReconciler struct {
	client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*ConnectorBindingReconciler)(nil)

func (r *ConnectorBindingReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrllog.FromContext(ctx)

	cb := new(walruscore.ConnectorBinding)
	err := r.client.Get(ctx, req.NamespacedName, cb)
	if err != nil {
		logger.Error(err, "failed to get ConnectorBinding", "namespace", req.Namespace, "name", req.Name)
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	if cb.DeletionTimestamp != nil {
		if systemmeta.Unlock(cb) {
			return ctrl.Result{}, nil
		}

		// Unlabel environment.
		envList := &walrus.EnvironmentList{}
		if err := r.client.List(ctx, envList, ctrlcli.MatchingFields{"metadata.name": cb.Namespace}); err != nil {
			logger.Error(err, "failed to list Environments for ConnectorBinding", "namespace", cb.Namespace)
			return ctrl.Result{}, err
		}

		if len(envList.Items) == 1 {
			env := envList.Items[0]

			labels := env.GetLabels()
			delete(labels, walruscore.ProviderLabelPrefix+strings.ToLower(cb.Status.Type))
			env.SetLabels(labels)

			err = r.client.Update(ctx, &env)
			if err != nil {
				logger.Error(err, "failed to update Environment", "name", env.Name)
				return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
			}
		}

		// Unlock.
		_, err = kubeclientset.UpdateWithCtrlClient(ctx, r.client, cb)
		if err != nil {
			logger.Error(err, "failed to unlock ConnectorBinding", "namespace", cb.Namespace, "name", cb.Name)
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}

		return ctrl.Result{}, nil
	}

	// Lock if not.
	if !systemmeta.Lock(cb) {
		cb, err = kubeclientset.UpdateWithCtrlClient(ctx, r.client, cb)
		if err != nil {
			logger.Error(err, "failed to lock ConnectorBinding", "namespace", cb.Namespace, "name", cb.Name)
			return ctrl.Result{}, err
		}
	}

	conn := &walruscore.Connector{
		ObjectMeta: meta.ObjectMeta{
			Name:      cb.Spec.Connector.Name,
			Namespace: cb.Spec.Connector.Namespace,
		},
	}

	err = r.client.Get(ctx, ctrlcli.ObjectKeyFromObject(conn), conn)
	if err != nil {
		return ctrl.Result{}, err
	}

	if cb.Status.Type != conn.Spec.Type || cb.Status.Category != conn.Spec.Category {
		cb.Status.Type = conn.Spec.Type
		cb.Status.Category = conn.Spec.Category

		err = r.client.Status().Update(ctx, cb)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	// Label environment.
	envList := &walrus.EnvironmentList{}
	if err := r.client.List(ctx, envList, ctrlcli.MatchingFields{"metadata.name": cb.Namespace}); err != nil {
		logger.Error(err, "failed to list Environments for ConnectorBinding", "namespace", cb.Namespace)
		return ctrl.Result{}, err
	}

	if len(envList.Items) != 1 {
		// NB: we should never reach here.
		logger.Error(nil, "cannot fetch corresponding environment")
		return ctrl.Result{RequeueAfter: time.Second}, nil
	}

	env := envList.Items[0]

	labels := env.Labels
	if labels == nil {
		labels = make(map[string]string)
	}

	labels[walruscore.ProviderLabelPrefix+strings.ToLower(cb.Status.Type)] = "true"
	env.SetLabels(labels)

	err = r.client.Update(ctx, &env)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ConnectorBindingReconciler) SetupController(ctx context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err := fi.IndexField(ctx, &walrus.Environment{}, "metadata.name",
		func(obj ctrlcli.Object) []string {
			if obj == nil {
				return nil
			}
			return []string{obj.GetName()}
		})
	if err != nil {
		return fmt.Errorf("index environment 'metadata.name': %w", err)
	}

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&walruscore.ConnectorBinding{},
			ctrlbuilder.WithPredicates(ctrlpredicate.GenerationChangedPredicate{}),
		).
		Complete(r)
}
