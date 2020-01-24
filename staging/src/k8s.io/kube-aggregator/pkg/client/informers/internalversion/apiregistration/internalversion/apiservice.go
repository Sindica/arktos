/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Alkaid - file modified.

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

package internalversion

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	apiregistration "k8s.io/kube-aggregator/pkg/apis/apiregistration"
	internalclientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/kube-aggregator/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "k8s.io/kube-aggregator/pkg/client/listers/apiregistration/internalversion"
)

// APIServiceInformer provides access to a shared informer and lister for
// APIServices.
type APIServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.APIServiceLister
}

type aPIServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAPIServiceInformer constructs a new informer for APIService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIServiceInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIServiceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIServiceInformer constructs a new informer for APIService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIServiceInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Apiregistration().APIServices().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Apiregistration().APIServices().Watch(options)
			},
		},
		&apiregistration.APIService{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIServiceInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIServiceInformer(client, resyncPeriod, cache.Indexers{}, f.tweakListOptions)
}

func (f *aPIServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiregistration.APIService{}, f.defaultInformer)
}

func (f *aPIServiceInformer) Lister() internalversion.APIServiceLister {
	return internalversion.NewAPIServiceLister(f.Informer().GetIndexer())
}
