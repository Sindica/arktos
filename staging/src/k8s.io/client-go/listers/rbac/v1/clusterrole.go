/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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
	v1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterRoleLister helps list ClusterRoles.
type ClusterRoleLister interface {
	// List lists all ClusterRoles in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterRole, err error)
	// ClusterRoles returns an object that can list and get ClusterRoles.
	ClusterRoles() ClusterRoleTenantLister
	ClusterRolesWithMultiTenancy(tenant string) ClusterRoleTenantLister
	// Get retrieves the ClusterRole from the index for a given name.
	Get(name string) (*v1.ClusterRole, error)
	ClusterRoleListerExpansion
}

// clusterRoleLister implements the ClusterRoleLister interface.
type clusterRoleLister struct {
	indexer cache.Indexer
}

// NewClusterRoleLister returns a new ClusterRoleLister.
func NewClusterRoleLister(indexer cache.Indexer) ClusterRoleLister {
	return &clusterRoleLister{indexer: indexer}
}

// List lists all ClusterRoles in the indexer.
func (s *clusterRoleLister) List(selector labels.Selector) (ret []*v1.ClusterRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterRole))
	})
	return ret, err
}

// Get retrieves the ClusterRole from the index for a given name.
func (s *clusterRoleLister) Get(name string) (*v1.ClusterRole, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterrole"), name)
	}
	return obj.(*v1.ClusterRole), nil
}

// ClusterRoles returns an object that can list and get ClusterRoles.
func (s *clusterRoleLister) ClusterRoles() ClusterRoleTenantLister {
	return clusterRoleTenantLister{indexer: s.indexer, tenant: "system"}
}

func (s *clusterRoleLister) ClusterRolesWithMultiTenancy(tenant string) ClusterRoleTenantLister {
	return clusterRoleTenantLister{indexer: s.indexer, tenant: tenant}
}

// ClusterRoleTenantLister helps list and get ClusterRoles.
type ClusterRoleTenantLister interface {
	// List lists all ClusterRoles in the indexer for a given tenant/tenant.
	List(selector labels.Selector) (ret []*v1.ClusterRole, err error)
	// Get retrieves the ClusterRole from the indexer for a given tenant/tenant and name.
	Get(name string) (*v1.ClusterRole, error)
	ClusterRoleTenantListerExpansion
}

// clusterRoleTenantLister implements the ClusterRoleTenantLister
// interface.
type clusterRoleTenantLister struct {
	indexer cache.Indexer
	tenant  string
}

// List lists all ClusterRoles in the indexer for a given tenant.
func (s clusterRoleTenantLister) List(selector labels.Selector) (ret []*v1.ClusterRole, err error) {
	err = cache.ListAllByTenant(s.indexer, s.tenant, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterRole))
	})
	return ret, err
}

// Get retrieves the ClusterRole from the indexer for a given tenant and name.
func (s clusterRoleTenantLister) Get(name string) (*v1.ClusterRole, error) {
	key := s.tenant + "/" + name
	if s.tenant == "system" {
		key = name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterrole"), name)
	}
	return obj.(*v1.ClusterRole), nil
}
