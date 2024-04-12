// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	walrusv1 "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	clientset "github.com/seal-io/walrus/pkg/clients/clientset"
	internalinterfaces "github.com/seal-io/walrus/pkg/clients/informers/internalinterfaces"
	v1 "github.com/seal-io/walrus/pkg/clients/listers/walrus/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConnectorBindingInformer provides access to a shared informer and lister for
// ConnectorBindings.
type ConnectorBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ConnectorBindingLister
}

type connectorBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewConnectorBindingInformer constructs a new informer for ConnectorBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConnectorBindingInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConnectorBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredConnectorBindingInformer constructs a new informer for ConnectorBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConnectorBindingInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WalrusV1().ConnectorBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WalrusV1().ConnectorBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&walrusv1.ConnectorBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *connectorBindingInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConnectorBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *connectorBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&walrusv1.ConnectorBinding{}, f.defaultInformer)
}

func (f *connectorBindingInformer) Lister() v1.ConnectorBindingLister {
	return v1.NewConnectorBindingLister(f.Informer().GetIndexer())
}
