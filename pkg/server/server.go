package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/seal-io/utils/contextx"
	"github.com/seal-io/utils/funcx"
	"github.com/seal-io/utils/httpx"
	"github.com/seal-io/utils/pools/gopool"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/apiserver/pkg/server/routes"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/component-base/logs"
	"k8s.io/component-base/metrics/legacyregistry"
	"k8s.io/klog/v2"
	apireg "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	"k8s.io/utils/ptr"
	ctrlmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	"github.com/seal-io/walrus/pkg/apis"
	"github.com/seal-io/walrus/pkg/extensionapis"
	"github.com/seal-io/walrus/pkg/kuberest"
	"github.com/seal-io/walrus/pkg/manager"
	"github.com/seal-io/walrus/pkg/server/webserver"
	"github.com/seal-io/walrus/pkg/servers/serverset/scheme"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemapp"
	"github.com/seal-io/walrus/pkg/systemauthz"
	"github.com/seal-io/walrus/pkg/systemkuberes"
	"github.com/seal-io/walrus/pkg/systemsetting"
	"github.com/seal-io/walrus/pkg/webhooks"
)

func init() {
	ctrlmetrics.Registry = struct {
		prometheus.Registerer
		prometheus.Gatherer
	}{
		Registerer: legacyregistry.Registerer(),
		Gatherer:   legacyregistry.DefaultGatherer,
	}
}

type Server struct {
	Manager   *manager.Manager
	APIServer *genericapiserver.GenericAPIServer
}

// Prepare prepares the runtime for the server,
// including installing system resources, etc.
func (s *Server) Prepare(ctx context.Context) error {
	err := s.Manager.Prepare(ctx)
	if err != nil {
		return err
	}

	loopbackKubeCli := system.LoopbackKubeClient.Get()
	_, bindPort, _ := s.APIServer.SecureServingInfo.HostPort()

	// Install system namespace.
	err = systemkuberes.InstallSystemNamespace(ctx, loopbackKubeCli)
	if err != nil {
		return fmt.Errorf("install system namespace: %w", err)
	}

	// Install fake system routing service if needed.
	if !system.LoopbackKubeInside.Get() && system.LoopbackKubeNearby.Get() {
		// NB(thxCode): Need to enable the loopback Kubernetes APIServer's `--enable-aggregator-routing` flag also.
		err = systemkuberes.InstallFakeSystemRoutingService(ctx, loopbackKubeCli, bindPort)
		if err != nil {
			return fmt.Errorf("install fake system routing service: %w", err)
		}
	}

	// Install extension API services.
	{
		// By default, we hope to deploy in HA mode(by all-in-one YAML or Helm Chart),
		// so the system Kubernetes Service is created before webhook server start.
		cc := apireg.ServiceReference{
			Namespace: systemkuberes.SystemNamespaceName,
			Name:      systemkuberes.SystemRoutingServiceName,
			Port:      ptr.To(int32(bindPort)),
		}
		err = apis.InstallAPIServices(ctx, loopbackKubeCli, cc, nil)
		if err != nil {
			return fmt.Errorf("install extension API services: %w", err)
		}
	}

	// Initialize settings.
	err = systemsetting.Initialize(ctx, loopbackKubeCli)
	if err != nil {
		return fmt.Errorf("install settings: %w", err)
	}

	// Install authorization.
	err = systemauthz.Initialize(ctx, loopbackKubeCli)
	if err != nil {
		return fmt.Errorf("install authorization: %w", err)
	}

	// Install application.
	err = systemapp.Install(ctx, system.LoopbackKubeClientConfig.Get(), system.DisableApplications.Get())
	if err != nil {
		return fmt.Errorf("install application: %w", err)
	}

	// Initialize builtin resources after cache synced.
	err = s.APIServer.AddPostStartHook("install-builtin-resources", func(phc genericapiserver.PostStartHookContext) error {
		ctx := contextx.Background(phc.StopCh)

		// NB(thxCode): we install builtin resources after manager is ready,
		// which allows the installer to use the extension api resources.
		err = s.Manager.WaitForReady(ctx)
		if err != nil {
			return fmt.Errorf("wait for manager to be ready: %w", err)
		}

		// Initialize default subject provider.
		err = systemkuberes.InstallDefaultSubjectProvider(ctx, loopbackKubeCli)
		if err != nil {
			return err
		}

		// Initialize default subject.
		err = systemkuberes.InstallAdminSubject(ctx, loopbackKubeCli, system.BootstrapPassword.Get())
		if err != nil {
			return err
		}

		// Initialize default project.
		err = systemkuberes.InstallDefaultProject(ctx, loopbackKubeCli)
		if err != nil {
			return err
		}

		// Initialize default environment.
		err = systemkuberes.InstallDefaultEnvironment(ctx, loopbackKubeCli)
		if err != nil {
			return err
		}

		// Initialize built catalog.
		err = systemkuberes.InstallBuiltinCatalog(ctx, loopbackKubeCli, system.BuiltinCatalogVCSPlatform.Get())
		if err != nil {
			return err
		}

		klog.Info("!!! everything is ready !!!")
		return nil
	})
	if err != nil {
		return fmt.Errorf("add post-start hook: %w", err)
	}

	return nil
}

// Start starts the server.
//
// Start sets up extension apis, webhooks and registers assistant routes,
// before starting the manager and the API server.
func (s *Server) Start(ctx context.Context) error {
	cm := s.Manager.CtrlManager
	mu := s.APIServer.Handler.NonGoRestfulMux

	// Setup extension API handlers.
	err := extensionapis.Setup(ctx, s.APIServer, scheme.Scheme, scheme.ParameterCodec, scheme.Codecs, cm)
	if err != nil {
		return fmt.Errorf("setup extension API handlers: %w", err)
	}

	// Register /validate-*, /mutate-*.
	err = webhooks.Setup(ctx, cm, mu)
	if err != nil {
		return fmt.Errorf("setup webhooks: %w", err)
	}

	// Register /metrics.
	mu.Handle("/metrics", legacyregistry.Handler())

	// Register /readyz.
	{
		err = s.APIServer.AddReadyzChecks(
			healthz.NamedCheck("manager", func(r *http.Request) error {
				err := s.Manager.WaitForReady(ctx)
				if err != nil {
					return fmt.Errorf("wait for manager to be ready: %w", err)
				}
				return nil
			}))
		if err != nil {
			return fmt.Errorf("add readyz checks: %w", err)
		}
	}

	// Register /livez.
	{
		err = s.APIServer.AddLivezChecks(10*time.Second,
			healthz.NamedCheck("gopool", func(r *http.Request) error {
				return gopool.IsHealthy()
			}),
			healthz.NamedCheck("loopback", func(r *http.Request) error {
				restCli := funcx.MustNoError(
					rest.UnversionedRESTClientForConfigAndClient(
						dynamic.ConfigFor(cm.GetConfig()),
						cm.GetHTTPClient(),
					),
				)
				return kuberest.IsAvailable(r.Context(), restCli)
			}),
		)
		if err != nil {
			return fmt.Errorf("add livez checks: %w", err)
		}
	}

	// Register /debug.
	{
		runtime.SetBlockProfileRate(1)
		mu.Handle("/debug/pprof/", httpx.LoopbackAccessHandlerFunc(pprof.Index))
		mu.Handle("/debug/pprof/cmdline", httpx.LoopbackAccessHandlerFunc(pprof.Cmdline))
		mu.Handle("/debug/pprof/profile", httpx.LoopbackAccessHandlerFunc(pprof.Profile))
		mu.Handle("/debug/pprof/symbol", httpx.LoopbackAccessHandlerFunc(pprof.Symbol))
		mu.Handle("/debug/pprof/trace", httpx.LoopbackAccessHandlerFunc(pprof.Trace))
		mu.Handle("/debug/flags/v", httpx.LoopbackAccessHandlerFunc(routes.StringFlagPutHandler(logs.GlogSetter)))
	}

	// Register extension apis.
	mu.NotFoundHandler(webserver.Index())

	// Start.
	gp := gopool.GroupWithContextIn(ctx)
	gp.Go(func(ctx context.Context) error {
		// NB(thxCode): we start the manager after extension api is ready,
		// which allows the controller to index the extension api resources.
		err := apis.WaitForAPIServicesReady(ctx, system.LoopbackKubeClient.Get())
		if err != nil {
			return fmt.Errorf("wait for extension API services to be ready: %w", err)
		}
		return s.Manager.Start(ctx)
	})
	gp.Go(func(ctx context.Context) error {
		klog.Info("starting api server")
		return s.APIServer.PrepareRun().Run(ctx.Done())
	})
	return gp.Wait()
}
