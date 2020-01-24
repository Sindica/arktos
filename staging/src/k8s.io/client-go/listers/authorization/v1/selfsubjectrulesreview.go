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
	v1 "k8s.io/api/authorization/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SelfSubjectRulesReviewLister helps list SelfSubjectRulesReviews.
type SelfSubjectRulesReviewLister interface {
	// List lists all SelfSubjectRulesReviews in the indexer.
	List(selector labels.Selector) (ret []*v1.SelfSubjectRulesReview, err error)
	// SelfSubjectRulesReviews returns an object that can list and get SelfSubjectRulesReviews.
	SelfSubjectRulesReviews() SelfSubjectRulesReviewTenantLister
	SelfSubjectRulesReviewsWithMultiTenancy(tenant string) SelfSubjectRulesReviewTenantLister
	// Get retrieves the SelfSubjectRulesReview from the index for a given name.
	Get(name string) (*v1.SelfSubjectRulesReview, error)
	SelfSubjectRulesReviewListerExpansion
}

// selfSubjectRulesReviewLister implements the SelfSubjectRulesReviewLister interface.
type selfSubjectRulesReviewLister struct {
	indexer cache.Indexer
}

// NewSelfSubjectRulesReviewLister returns a new SelfSubjectRulesReviewLister.
func NewSelfSubjectRulesReviewLister(indexer cache.Indexer) SelfSubjectRulesReviewLister {
	return &selfSubjectRulesReviewLister{indexer: indexer}
}

// List lists all SelfSubjectRulesReviews in the indexer.
func (s *selfSubjectRulesReviewLister) List(selector labels.Selector) (ret []*v1.SelfSubjectRulesReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SelfSubjectRulesReview))
	})
	return ret, err
}

// Get retrieves the SelfSubjectRulesReview from the index for a given name.
func (s *selfSubjectRulesReviewLister) Get(name string) (*v1.SelfSubjectRulesReview, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("selfsubjectrulesreview"), name)
	}
	return obj.(*v1.SelfSubjectRulesReview), nil
}

// SelfSubjectRulesReviews returns an object that can list and get SelfSubjectRulesReviews.
func (s *selfSubjectRulesReviewLister) SelfSubjectRulesReviews() SelfSubjectRulesReviewTenantLister {
	return selfSubjectRulesReviewTenantLister{indexer: s.indexer, tenant: "default"}
}

func (s *selfSubjectRulesReviewLister) SelfSubjectRulesReviewsWithMultiTenancy(tenant string) SelfSubjectRulesReviewTenantLister {
	return selfSubjectRulesReviewTenantLister{indexer: s.indexer, tenant: tenant}
}

// SelfSubjectRulesReviewTenantLister helps list and get SelfSubjectRulesReviews.
type SelfSubjectRulesReviewTenantLister interface {
	// List lists all SelfSubjectRulesReviews in the indexer for a given tenant/tenant.
	List(selector labels.Selector) (ret []*v1.SelfSubjectRulesReview, err error)
	// Get retrieves the SelfSubjectRulesReview from the indexer for a given tenant/tenant and name.
	Get(name string) (*v1.SelfSubjectRulesReview, error)
	SelfSubjectRulesReviewTenantListerExpansion
}

// selfSubjectRulesReviewTenantLister implements the SelfSubjectRulesReviewTenantLister
// interface.
type selfSubjectRulesReviewTenantLister struct {
	indexer cache.Indexer
	tenant  string
}

// List lists all SelfSubjectRulesReviews in the indexer for a given tenant.
func (s selfSubjectRulesReviewTenantLister) List(selector labels.Selector) (ret []*v1.SelfSubjectRulesReview, err error) {
	err = cache.ListAllByTenant(s.indexer, s.tenant, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SelfSubjectRulesReview))
	})
	return ret, err
}

// Get retrieves the SelfSubjectRulesReview from the indexer for a given tenant and name.
func (s selfSubjectRulesReviewTenantLister) Get(name string) (*v1.SelfSubjectRulesReview, error) {
	key := s.tenant + "/" + name
	if s.tenant == "default" {
		key = name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("selfsubjectrulesreview"), name)
	}
	return obj.(*v1.SelfSubjectRulesReview), nil
}
