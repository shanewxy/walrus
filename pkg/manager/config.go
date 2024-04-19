package manager

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	kmeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlcache "sigs.k8s.io/controller-runtime/pkg/cache"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlapiutil "sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	ctrlcfg "sigs.k8s.io/controller-runtime/pkg/config"
	ctrlmetricsrv "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/clients/clientset/scheme"
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
	DisableController         bool
}

func (c *Config) Apply(ctx context.Context) (*Manager, error) {
	// Set controller logger.
	ctrl.SetLogger(klog.Background().WithName("ctrl"))

	ctrlMgrOpts := ctrl.Options{
		// General.
		GracefulShutdownTimeout: ptr.To(30 * time.Second),
		Scheme:                  scheme.Scheme,
		Logger:                  klog.Background().V(4).WithName("ctrlmgr"),

		// Context.
		BaseContext: func() context.Context {
			return ctx
		},

		// Mapper.
		MapperProvider: func(config *rest.Config, _ *http.Client) (kmeta.RESTMapper, error) {
			// NB(thxCode): since mapper provider cannot reuse the singleton http client,
			// we have to pass the http client to the provider here.
			return ctrlapiutil.NewDynamicRESTMapper(config, c.KubeHTTPClient)
		},

		// Client.
		Client: ctrlcli.Options{
			// Reuse http client.
			HTTPClient: c.KubeHTTPClient,
		},
		NewClient: func(config *rest.Config, options ctrlcli.Options) (ctrlcli.Client, error) {
			// Create watchable controller client here.
			return ctrlcli.NewWithWatch(config, options)
		},

		// Cache.
		Cache: ctrlcache.Options{
			// Reuse http client.
			HTTPClient: c.KubeHTTPClient,
			// Set resync period to underlay informer.
			SyncPeriod: ptr.To(c.InformerCacheResyncPeriod),
		},

		// Controller.
		Controller: ctrlcfg.Controller{
			// How long we should wait for the watching cache to be synced before controller starts.
			CacheSyncTimeout: 2 * time.Minute,
			// Recover from panic.
			RecoverPanic: ptr.To(true),
			// NB(thxCode): make controllers run with leader election by default,
			// this configuration does not affect the manager leader election.
			NeedLeaderElection: ptr.To(true),
		},

		// Controller manager leader election.
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

	// Disable controller manager election if needed
	if c.DisableController {
		ctrlMgrOpts.LeaderElection = false
	}

	// Enable webhook serving, includes configurations installation.
	if c.ServeListener != nil {
		ctrlMgrOpts.WebhookServer = webhookserver.Enhance(c.ServeListener, c.ServeListenerCertDir, c.KubeClient)
	}

	// Create controller manager and wrap it.
	var ctrlManager CtrlManager
	{
		rawCtrlManager, err := ctrl.NewManager(rest.CopyConfig(&c.KubeClientConfig), ctrlMgrOpts)
		if err != nil {
			return nil, fmt.Errorf("create controller manager: %w", err)
		}
		ctrlManager = CtrlManager{
			Manager:           rawCtrlManager,
			httpClient:        c.KubeHTTPClient,
			disableController: c.DisableController,
			indexedFields:     sets.Set[string]{},
		}
	}

	// Add controller manager sentinel.
	sentinel := _CtrlManagerSentinel{done: make(chan struct{})}
	err := ctrlManager.Add(sentinel)
	if err != nil {
		return nil, fmt.Errorf("add controller manager sentinel: %w", err)
	}

	// Configure loopback controller runtime,
	// including the client and direct reading client.
	system.ConfigureLoopbackCtrlRuntime(ctrlManager.GetClient(), ctrlManager.GetAPIReader())

	return &Manager{
		CtrlManager: ctrlManager,
		sentinel:    sentinel,
	}, nil
}
