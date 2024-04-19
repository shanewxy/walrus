package manager

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/pprof"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/seal-io/utils/funcx"
	"github.com/seal-io/utils/httpx"
	"github.com/seal-io/utils/pools/gopool"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/apiserver/pkg/server/routes"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	ctrlhealthz "sigs.k8s.io/controller-runtime/pkg/healthz"
	ctrlmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	"github.com/seal-io/walrus/pkg/apis"
	"github.com/seal-io/walrus/pkg/controllers"
	"github.com/seal-io/walrus/pkg/kuberest"
	"github.com/seal-io/walrus/pkg/system"
)

type Manager struct {
	CtrlManager CtrlManager
	sentinel    _CtrlManagerSentinel
}

// Prepare prepares the runtime for the manager,
// including installing CRDs and registering metric collectors.
func (m *Manager) Prepare(ctx context.Context) error {
	loopbackKubeCli := system.LoopbackKubeClient.Get()

	// Initialize CRDs.
	err := apis.InstallCustomResourceDefinitions(ctx, loopbackKubeCli)
	if err != nil {
		return fmt.Errorf("install CRDs: %w", err)
	}

	// Register metric collectors.
	{
		reg := ctrlmetrics.Registry
		cs := []prometheus.Collector{
			collectors.NewBuildInfoCollector(),
			gopool.NewStatsCollector(),
		}
		for i := range cs {
			err = reg.Register(cs[i])
			if err != nil {
				return fmt.Errorf("register metric collector: %w", err)
			}
		}
	}

	return nil
}

// Start starts the manager.
//
// Start sets up controllers and registers assistant routes,
// before starting the controller manager.
func (m *Manager) Start(ctx context.Context) error {
	cm := m.CtrlManager
	ms := cm.GetWebhookServer()

	// Setup controllers.
	err := controllers.Setup(ctx, cm)
	if err != nil {
		return fmt.Errorf("setup controllers: %w", err)
	}

	// Register /metrics.
	{
		h := promhttp.HandlerOpts{
			ErrorLog:      klog.NewStandardLogger("WARNING"),
			ErrorHandling: promhttp.HTTPErrorOnError,
		}
		ms.Register("/metrics", promhttp.HandlerFor(ctrlmetrics.Registry, h))
	}

	// Register /readyz.
	{
		p := "/readyz"
		h := &ctrlhealthz.Handler{
			Checks: map[string]ctrlhealthz.Checker{
				"ping": ctrlhealthz.Ping,
				"log":  healthz.LogHealthz.Check,
				"informer": func(r *http.Request) error {
					if cm.GetCache().WaitForCacheSync(r.Context()) {
						return nil
					}
					return errors.New("informer cache is not synced yet")
				},
			},
		}
		ms.Register(p, http.StripPrefix(p, h))
	}

	// Register /livez.
	{
		p := "/livez"
		h := &ctrlhealthz.Handler{
			Checks: map[string]ctrlhealthz.Checker{
				"ping": ctrlhealthz.Ping,
				"log":  healthz.LogHealthz.Check,
				"gopool": func(r *http.Request) error {
					return gopool.IsHealthy()
				},
				"loopback": func(r *http.Request) error {
					restCli := funcx.MustNoError(
						rest.UnversionedRESTClientForConfigAndClient(
							dynamic.ConfigFor(cm.GetConfig()),
							cm.GetHTTPClient(),
						),
					)
					return kuberest.IsAvailable(r.Context(), restCli)
				},
			},
		}
		ms.Register(p, http.StripPrefix(p, h))
	}

	// Register /debug.
	{
		runtime.SetBlockProfileRate(1)
		ms.Register("/debug/pprof/", httpx.LoopbackAccessHandlerFunc(pprof.Index))
		ms.Register("/debug/pprof/cmdline", httpx.LoopbackAccessHandlerFunc(pprof.Cmdline))
		ms.Register("/debug/pprof/profile", httpx.LoopbackAccessHandlerFunc(pprof.Profile))
		ms.Register("/debug/pprof/symbol", httpx.LoopbackAccessHandlerFunc(pprof.Symbol))
		ms.Register("/debug/pprof/trace", httpx.LoopbackAccessHandlerFunc(pprof.Trace))
		ms.Register("/debug/flags/v", httpx.LoopbackAccessHandlerFunc(routes.StringFlagPutHandler(logs.GlogSetter)))
	}

	// Start.
	klog.Info("starting controller manager")
	return cm.Start(ctx)
}

// WaitForReady waits for the manager to be ready.
func (m *Manager) WaitForReady(ctx context.Context) error {
	// Wait for controller manager to start.
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-m.sentinel.Done():
	}

	// Wait for cache sync.
	if !m.CtrlManager.GetCache().WaitForCacheSync(ctx) {
		return errors.New("cache is not synced yet")
	}

	return nil
}
