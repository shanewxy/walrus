package manager

import (
	"context"
	"fmt"
	"math"
	"net"
	"net/http"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlcache "sigs.k8s.io/controller-runtime/pkg/cache"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlmetricsrv "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/clients/clientset/scheme"
	"github.com/seal-io/walrus/pkg/manager/leaderelection"
	"github.com/seal-io/walrus/pkg/manager/webhookserver"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemkuberes"
)

type Config struct {
	InformerCacheResyncPeriod time.Duration
	KubeConfigPath            string
	KubeClientConfig          rest.Config
	KubeHTTPClient            *http.Client
	KubeClient                clientset.Interface
	KubeLeaderElection        bool
	KubeLeaderLease           time.Duration
	KubeLeaderRenewTimeout    time.Duration
	ServeListenerCertDir      string
	ServeListener             net.Listener
	DisableCache              bool
}

func (c *Config) Apply(_ context.Context) (*Manager, error) {
	ctrlMgrOpts := ctrl.Options{
		// General.
		GracefulShutdownTimeout: ptr.To(30 * time.Second),
		Scheme:                  scheme.Scheme,
		Logger:                  klog.Background().WithName("ctrl"),

		// Client.
		Client: ctrlcli.Options{
			HTTPClient: c.KubeHTTPClient,
		},
		NewClient: func(config *rest.Config, options ctrlcli.Options) (ctrlcli.Client, error) {
			return ctrlcli.NewWithWatch(config, options)
		},

		// Cache.
		Cache: ctrlcache.Options{
			HTTPClient: c.KubeHTTPClient,
			SyncPeriod: ptr.To(c.InformerCacheResyncPeriod),
		},

		// Leader election.
		LeaderElectionReleaseOnCancel: true,
		LeaderElectionNamespace:       systemkuberes.SystemNamespaceName,
		LeaderElectionID:              "walrus-leader",
		LeaderElection:                c.KubeLeaderElection,
		LeaseDuration:                 ptr.To(c.KubeLeaderLease),
		RenewDeadline:                 ptr.To(c.KubeLeaderRenewTimeout),
		RetryPeriod:                   ptr.To(2 * time.Second),

		// Disable default webhook server.
		WebhookServer: webhookserver.Dummy(),
		// Disable default metrics service.
		Metrics: ctrlmetricsrv.Options{BindAddress: "0"},
		// Disable default healthcheck service.
		HealthProbeBindAddress: "0",
		// Disable default profiling service.
		PprofBindAddress: "0",
	}

	// Inject a lease locker that will never succeed to prevent controller from starting.
	if c.DisableCache {
		ctrlMgrOpts.LeaderElection = true
		ctrlMgrOpts.RetryPeriod = ptr.To(time.Duration(math.MaxInt64)) // Set longest retry period.
		ctrlMgrOpts.LeaderElectionResourceLockInterface = leaderelection.Dummy()
	}

	// Enable webhook serving,
	// includes configurations installation.
	if c.ServeListener != nil {
		ctrlMgrOpts.WebhookServer = webhookserver.Enhance(c.ServeListener, c.ServeListenerCertDir, c.KubeClient)
	}

	// Create controller manager and wrap it.
	rawCtrlManager, err := ctrl.NewManager(rest.CopyConfig(&c.KubeClientConfig), ctrlMgrOpts)
	if err != nil {
		return nil, fmt.Errorf("create controller manager: %w", err)
	}
	// Add manager sentinel.
	sentinel := _CtrlManagerSentinel{done: make(chan struct{})}
	err = rawCtrlManager.Add(sentinel)
	if err != nil {
		return nil, fmt.Errorf("add manager sentinel: %w", err)
	}

	ctrlManager := CtrlManager{
		Manager:       rawCtrlManager,
		disableCache:  c.DisableCache,
		indexedFields: sets.Set[string]{},
	}
	system.ConfigureLoopbackCtrlRuntime(ctrlManager.GetClient(), ctrlManager.GetAPIReader())

	return &Manager{
		CtrlManager: ctrlManager,
		sentinel:    sentinel,
	}, nil
}
