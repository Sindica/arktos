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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NamespaceLister helps list Namespaces.
type NamespaceLister interface {
	// List lists all Namespaces in the indexer.
	List(selector labels.Selector) (ret []*v1.Namespace, err error)
	// Namespaces returns an object that can list and get Namespaces.
	Namespaces() NamespaceTenantLister
	NamespacesWithMultiTenancy(tenant string) NamespaceTenantLister
	// Get retrieves the Namespace from the index for a given name.
	Get(name string) (*v1.Namespace, error)
	NamespaceListerExpansion
}

// namespaceLister implements the NamespaceLister interface.
type namespaceLister struct {
	indexer cache.Indexer
}

// NewNamespaceLister returns a new NamespaceLister.
func NewNamespaceLister(indexer cache.Indexer) NamespaceLister {
	return &namespaceLister{indexer: indexer}
}

// List lists all Namespaces in the indexer.
func (s *namespaceLister) List(selector labels.Selector) (ret []*v1.Namespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Namespace))
	})
	return ret, err
}

// Get retrieves the Namespace from the index for a given name.
func (s *namespaceLister) Get(name string) (*v1.Namespace, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("namespace"), name)
	}
	return obj.(*v1.Namespace), nil
}

// Namespaces returns an object that can list and get Namespaces.
func (s *namespaceLister) Namespaces() NamespaceTenantLister {
	return namespaceTenantLister{indexer: s.indexer, tenant: "default"}
}

func (s *namespaceLister) NamespacesWithMultiTenancy(tenant string) NamespaceTenantLister {
	return namespaceTenantLister{indexer: s.indexer, tenant: tenant}
}

// NamespaceTenantLister helps list and get Namespaces.
type NamespaceTenantLister interface {
	// List lists all Namespaces in the indexer for a given tenant/tenant.
	List(selector labels.Selector) (ret []*v1.Namespace, err error)
	// Get retrieves the Namespace from the indexer for a given tenant/tenant and name.
	Get(name string) (*v1.Namespace, error)
	NamespaceTenantListerExpansion
}

// namespaceTenantLister implements the NamespaceTenantLister
// interface.
type namespaceTenantLister struct {
	indexer cache.Indexer
	tenant  string
}

// List lists all Namespaces in the indexer for a given tenant.
func (s namespaceTenantLister) List(selector labels.Selector) (ret []*v1.Namespace, err error) {
	err = cache.ListAllByTenant(s.indexer, s.tenant, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Namespace))
	})
	return ret, err
}

// Get retrieves the Namespace from the indexer for a given tenant and name.
func (s namespaceTenantLister) Get(name string) (*v1.Namespace, error) {
	key := s.tenant + "/" + name
	if s.tenant == "default" {
		key = name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("namespace"), name)
	}
	return obj.(*v1.Namespace), nil
}
