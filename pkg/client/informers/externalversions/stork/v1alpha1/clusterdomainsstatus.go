/*
Copyright 2018 Openstorage.org

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	storkv1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	versioned "github.com/libopenstorage/stork/pkg/client/clientset/versioned"
	internalinterfaces "github.com/libopenstorage/stork/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/libopenstorage/stork/pkg/client/listers/stork/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterDomainsStatusInformer provides access to a shared informer and lister for
// ClusterDomainsStatuses.
type ClusterDomainsStatusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterDomainsStatusLister
}

type clusterDomainsStatusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterDomainsStatusInformer constructs a new informer for ClusterDomainsStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterDomainsStatusInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterDomainsStatusInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterDomainsStatusInformer constructs a new informer for ClusterDomainsStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterDomainsStatusInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorkV1alpha1().ClusterDomainsStatuses().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorkV1alpha1().ClusterDomainsStatuses().Watch(options)
			},
		},
		&storkv1alpha1.ClusterDomainsStatus{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterDomainsStatusInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterDomainsStatusInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterDomainsStatusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storkv1alpha1.ClusterDomainsStatus{}, f.defaultInformer)
}

func (f *clusterDomainsStatusInformer) Lister() v1alpha1.ClusterDomainsStatusLister {
	return v1alpha1.NewClusterDomainsStatusLister(f.Informer().GetIndexer())
}