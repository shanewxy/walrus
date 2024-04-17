package manager

import (
	"context"
	"net"
	"net/url"
	"slices"
	"strings"

	"github.com/seal-io/utils/netx"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlapiutil "sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/seal-io/walrus/pkg/system"
)

func isLoopbackClusterNearby(restCfg *rest.Config) bool {
	// Extract host from rest config.
	var host string
	if strings.Contains(restCfg.Host, "://") {
		u, _ := url.Parse(restCfg.Host)
		host = u.Host
	} else {
		host = restCfg.Host
	}
	if h, _, err := net.SplitHostPort(host); err == nil {
		host = h
	} else if strings.Contains(host, ":") {
		host = strings.Split(host, ":")[0]
	}

	// Detect host in a fast pass way.
	knownLoopbackHosts := []string{
		"kubernetes.docker.internal",
		"host.docker.internal",
		"localhost",
		"127.0.0.1",
		"[::1]",
		"[::1%lo0]",
	}
	if slices.Contains(knownLoopbackHosts, host) {
		return true
	}

	// Detect host in a slow pass way.
	subnets := make([]netx.IPv4, 0, system.Subnets.Get().Len())
	for _, v := range system.Subnets.Get().List() {
		sn := netx.MustIPv4FromCIDR(v)
		subnets = append(subnets, sn)
	}

	// IP detect.
	if ip := net.ParseIP(host); ip != nil {
		for j := range subnets {
			if subnets[j].Contains(ip) {
				return true
			}
		}

		return false
	}

	// Or DNS lookup.
	ips, err := net.LookupIP(host)
	if err != nil {
		return false
	}

	for i := range ips {
		if ips[i].IsLoopback() {
			return true
		}
		for j := range subnets {
			if subnets[j].Contains(ips[i]) {
				return true
			}
		}
	}

	return false
}

type (
	// CtrlManager is a wrapper around ctrl.Manager.
	CtrlManager struct {
		ctrl.Manager
		disableCache  bool
		indexedFields sets.Set[string]
	}

	// UnCachedCtrlClient is a wrapper around ctrl.Client.
	UnCachedCtrlClient struct {
		ctrlcli.Client
		apiReader ctrlcli.Reader
	}

	// RepeatableCtrlFieldIndexer is a wrapper around ctrlcli.FieldIndexer.
	RepeatableCtrlFieldIndexer struct {
		ctrl.Manager
		indexedFields sets.Set[string]
	}
)

func (m CtrlManager) GetClient() ctrlcli.Client {
	if !m.disableCache {
		return m.Manager.GetClient()
	}
	return UnCachedCtrlClient{
		Client:    m.Manager.GetClient(),
		apiReader: m.Manager.GetAPIReader(),
	}
}

func (m CtrlManager) Start(ctx context.Context) error {
	return m.Manager.Start(ctx)
}

func (m CtrlManager) CacheDisabled() bool {
	return m.disableCache
}

func (i UnCachedCtrlClient) Get(ctx context.Context, key ctrlcli.ObjectKey, obj ctrlcli.Object, opts ...ctrlcli.GetOption) error {
	return i.apiReader.Get(ctx, key, obj, opts...)
}

func (i UnCachedCtrlClient) List(ctx context.Context, list ctrlcli.ObjectList, opts ...ctrlcli.ListOption) error {
	return i.apiReader.List(ctx, list, opts...)
}

func (i UnCachedCtrlClient) Watch(ctx context.Context, obj ctrlcli.ObjectList, opts ...ctrlcli.ListOption) (watch.Interface, error) {
	return i.Client.(ctrlcli.WithWatch).Watch(ctx, obj, opts...)
}

func (i UnCachedCtrlClient) CacheDisabled() bool {
	return true
}

func (m CtrlManager) GetFieldIndexer() ctrlcli.FieldIndexer {
	return RepeatableCtrlFieldIndexer{
		Manager:       m.Manager,
		indexedFields: m.indexedFields,
	}
}

func (i RepeatableCtrlFieldIndexer) IndexField(ctx context.Context, obj ctrlcli.Object, field string, extractValue ctrlcli.IndexerFunc) error {
	logger := ctrllog.FromContext(ctx)
	gvk, err := ctrlapiutil.GVKForObject(obj, i.Manager.GetScheme())
	if err != nil {
		return err
	}
	key := gvk.String() + "/" + field
	if i.indexedFields.Has(key) {
		// If the field is already indexed, skip.
		logger.Info("field is already indexed, skipping", "field", field, "gvk", gvk)
		return nil
	}
	i.indexedFields.Insert(key)
	return i.Manager.GetFieldIndexer().IndexField(ctx, obj, field, extractValue)
}

// _CtrlManagerSentinel is a ctrlmgr.Runnable implementation for observing
// whether the ctrl.Manager is started.
type _CtrlManagerSentinel struct {
	done chan struct{}
}

func (s _CtrlManagerSentinel) Start(ctx context.Context) error {
	close(s.done)
	<-ctx.Done()
	return ctx.Err()
}

func (s _CtrlManagerSentinel) NeedLeaderElection() bool {
	return false
}

func (s _CtrlManagerSentinel) Done() <-chan struct{} {
	return s.done
}
