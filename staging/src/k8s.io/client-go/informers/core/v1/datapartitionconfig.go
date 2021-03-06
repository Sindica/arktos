/*
Copyright 2020 Authors of Arktos.

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

package v1

import (
	time "time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
	cache "k8s.io/client-go/tools/cache"
)

// DataPartitionConfigInformer provides access to a shared informer and lister for
// DataPartitionConfigs.
type DataPartitionConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DataPartitionConfigLister
}

type dataPartitionConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDataPartitionConfigInformer constructs a new informer for DataPartitionConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDataPartitionConfigInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDataPartitionConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDataPartitionConfigInformer constructs a new informer for DataPartitionConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDataPartitionConfigInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().DataPartitionConfigs().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().DataPartitionConfigs().Watch(options)
			},
		},
		&corev1.DataPartitionConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *dataPartitionConfigInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDataPartitionConfigInformer(client, resyncPeriod, cache.Indexers{}, f.tweakListOptions)
}

func (f *dataPartitionConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.DataPartitionConfig{}, f.defaultInformer)
}

func (f *dataPartitionConfigInformer) Lister() v1.DataPartitionConfigLister {
	return v1.NewDataPartitionConfigLister(f.Informer().GetIndexer())
}
