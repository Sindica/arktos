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
	v1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// JobLister helps list Jobs.
type JobLister interface {
	// List lists all Jobs in the indexer.
	List(selector labels.Selector) (ret []*v1.Job, err error)
	// Jobs returns an object that can list and get Jobs.
	Jobs(namespace string) JobNamespaceLister
	JobsWithMultiTenancy(namespace string, tenant string) JobNamespaceLister
	JobListerExpansion
}

// jobLister implements the JobLister interface.
type jobLister struct {
	indexer cache.Indexer
}

// NewJobLister returns a new JobLister.
func NewJobLister(indexer cache.Indexer) JobLister {
	return &jobLister{indexer: indexer}
}

// List lists all Jobs in the indexer.
func (s *jobLister) List(selector labels.Selector) (ret []*v1.Job, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Job))
	})
	return ret, err
}

// Jobs returns an object that can list and get Jobs.
func (s *jobLister) Jobs(namespace string) JobNamespaceLister {
	return jobNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: "default"}
}

func (s *jobLister) JobsWithMultiTenancy(namespace string, tenant string) JobNamespaceLister {
	return jobNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: tenant}
}

// JobNamespaceLister helps list and get Jobs.
type JobNamespaceLister interface {
	// List lists all Jobs in the indexer for a given tenant/namespace.
	List(selector labels.Selector) (ret []*v1.Job, err error)
	// Get retrieves the Job from the indexer for a given tenant/namespace and name.
	Get(name string) (*v1.Job, error)
	JobNamespaceListerExpansion
}

// jobNamespaceLister implements the JobNamespaceLister
// interface.
type jobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
	tenant    string
}

// List lists all Jobs in the indexer for a given namespace.
func (s jobNamespaceLister) List(selector labels.Selector) (ret []*v1.Job, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.tenant, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Job))
	})
	return ret, err
}

// Get retrieves the Job from the indexer for a given namespace and name.
func (s jobNamespaceLister) Get(name string) (*v1.Job, error) {
	key := s.tenant + "/" + s.namespace + "/" + name
	if s.tenant == "default" {
		key = s.namespace + "/" + name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("job"), name)
	}
	return obj.(*v1.Job), nil
}
