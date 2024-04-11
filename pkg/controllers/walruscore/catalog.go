package walruscore

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/seal-io/utils/pools/gopool"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/systemmeta"
	"github.com/seal-io/walrus/pkg/templates/fetcher"
	tmpllister "github.com/seal-io/walrus/pkg/templates/lister"
)

// CatalogReconciler reconciles a v1.Catalog object.
type CatalogReconciler struct {
	client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*CatalogReconciler)(nil)

// Reconcile reconciles the catalog.
func (r *CatalogReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	logger := ctrllog.FromContext(ctx)

	// Fetch.
	obj := new(walruscore.Catalog)
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

		// Skip catalog already reconciling.
		if apistatus.CatalogConditionReady.IsUnknown(obj) && !apistatus.CatalogConditionRefresh.IsUnknown(obj) {
			return ctrl.Result{}, nil
		}

		// Skip catalog finished reconciling.
		if (apistatus.CatalogConditionReady.IsTrue(obj) || apistatus.CatalogConditionReady.IsFalse(obj)) &&
			!apistatus.CatalogConditionRefresh.IsUnknown(obj) {
			return ctrl.Result{}, nil
		}
	}

	// Initialize status.
	{
		apistatus.CatalogConditionReady.Unknown(obj, apistatus.CatalogConditionReasonPreparing, "Preparing")
		if apistatus.CatalogConditionRefresh.IsUnknown(obj) {
			apistatus.CatalogConditionRefresh.False(obj, "", "")
		}

		err = r.client.Status().Update(ctx, obj)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	defer func() {
		if err != nil {
			apistatus.CatalogConditionReady.False(obj, apistatus.CatalogConditionReasonError, err.Error())
			_ = r.updateStatus(ctx, obj)
		}
	}()

	// Ensure templates.
	var (
		succeed = int32(0)
		failed  = int32(0)
		invalid = int32(0)
	)

	{
		var (
			wg        = gopool.GroupWithContextIn(ctx)
			batchSize = 10
		)

		tmpls, err := tmpllister.List(ctx, obj)
		if err != nil {
			return ctrl.Result{}, err
		}

		for i := 0; i < batchSize; i++ {
			s := i

			wg.Go(func(ctx context.Context) error {
				for j := s; j < len(tmpls); j += batchSize {
					tmpl, err := fetcher.Fetch(ctx, &tmpls[j])
					if err != nil {
						atomic.AddInt32(&failed, 1)
						return fmt.Errorf("failed to fetch template %s/%s: %w", tmpls[j].Namespace, tmpls[j].Name, err)
					}

					if tmpl == nil {
						atomic.AddInt32(&invalid, 1)
						continue
					}

					atomic.AddInt32(&succeed, 1)
				}

				return nil
			})
		}

		err = wg.Wait()
		if err != nil {
			return ctrl.Result{}, err
		}

		logger.Infof("synced catalog \"%s/%s\", total: %d, succeed: %d, failed: %d, invalid: %d",
			obj.Namespace, obj.Name, len(tmpls), succeed, failed, invalid)
	}

	// Update status.
	{
		apistatus.CatalogConditionReady.True(obj, "", "")
		obj.Status.Project = systemmeta.GetProjectName(obj.Namespace)
		obj.Status.URL = obj.Spec.VCSSource.URL
		obj.Status.TemplateCount = int64(succeed)
		obj.Status.LastSyncTime = meta.Now()

		err = r.updateStatus(ctx, obj)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// updateStatus updates the status of the catalog.
func (r *CatalogReconciler) updateStatus(ctx context.Context, obj *walruscore.Catalog) error {
	existed := new(walruscore.Catalog)
	err := r.client.Get(ctx, ctrlcli.ObjectKeyFromObject(obj), existed)
	if err != nil {
		return err
	}

	existed.Status = obj.Status
	existed.Status.ConditionSummary = *apistatus.WalkCatalog(&existed.Status.StatusDescriptor)
	err = r.client.Status().Update(ctx, existed)
	return err
}

// SetupController sets up the controller.
func (r *CatalogReconciler) SetupController(_ context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&walruscore.Catalog{}).
		Complete(r)
}
