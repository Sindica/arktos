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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CSINodeLister helps list CSINodes.
type CSINodeLister interface {
	// List lists all CSINodes in the indexer.
	List(selector labels.Selector) (ret []*v1.CSINode, err error)
	// CSINodes returns an object that can list and get CSINodes.
	CSINodes() CSINodeTenantLister
	CSINodesWithMultiTenancy(tenant string) CSINodeTenantLister
	// Get retrieves the CSINode from the index for a given name.
	Get(name string) (*v1.CSINode, error)
	CSINodeListerExpansion
}

// cSINodeLister implements the CSINodeLister interface.
type cSINodeLister struct {
	indexer cache.Indexer
}

// NewCSINodeLister returns a new CSINodeLister.
func NewCSINodeLister(indexer cache.Indexer) CSINodeLister {
	return &cSINodeLister{indexer: indexer}
}

// List lists all CSINodes in the indexer.
func (s *cSINodeLister) List(selector labels.Selector) (ret []*v1.CSINode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CSINode))
	})
	return ret, err
}

// Get retrieves the CSINode from the index for a given name.
func (s *cSINodeLister) Get(name string) (*v1.CSINode, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("csinode"), name)
	}
	return obj.(*v1.CSINode), nil
}

// CSINodes returns an object that can list and get CSINodes.
func (s *cSINodeLister) CSINodes() CSINodeTenantLister {
	return cSINodeTenantLister{indexer: s.indexer, tenant: "system"}
}

func (s *cSINodeLister) CSINodesWithMultiTenancy(tenant string) CSINodeTenantLister {
	return cSINodeTenantLister{indexer: s.indexer, tenant: tenant}
}

// CSINodeTenantLister helps list and get CSINodes.
type CSINodeTenantLister interface {
	// List lists all CSINodes in the indexer for a given tenant/tenant.
	List(selector labels.Selector) (ret []*v1.CSINode, err error)
	// Get retrieves the CSINode from the indexer for a given tenant/tenant and name.
	Get(name string) (*v1.CSINode, error)
	CSINodeTenantListerExpansion
}

// cSINodeTenantLister implements the CSINodeTenantLister
// interface.
type cSINodeTenantLister struct {
	indexer cache.Indexer
	tenant  string
}

// List lists all CSINodes in the indexer for a given tenant.
func (s cSINodeTenantLister) List(selector labels.Selector) (ret []*v1.CSINode, err error) {
	err = cache.ListAllByTenant(s.indexer, s.tenant, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CSINode))
	})
	return ret, err
}

// Get retrieves the CSINode from the indexer for a given tenant and name.
func (s cSINodeTenantLister) Get(name string) (*v1.CSINode, error) {
	key := s.tenant + "/" + name
	if s.tenant == "system" {
		key = name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("csinode"), name)
	}
	return obj.(*v1.CSINode), nil
}
