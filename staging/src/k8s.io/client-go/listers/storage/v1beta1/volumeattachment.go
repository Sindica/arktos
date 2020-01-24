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

package v1beta1

import (
	v1beta1 "k8s.io/api/storage/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeAttachmentLister helps list VolumeAttachments.
type VolumeAttachmentLister interface {
	// List lists all VolumeAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.VolumeAttachment, err error)
	// VolumeAttachments returns an object that can list and get VolumeAttachments.
	VolumeAttachments() VolumeAttachmentTenantLister
	VolumeAttachmentsWithMultiTenancy(tenant string) VolumeAttachmentTenantLister
	// Get retrieves the VolumeAttachment from the index for a given name.
	Get(name string) (*v1beta1.VolumeAttachment, error)
	VolumeAttachmentListerExpansion
}

// volumeAttachmentLister implements the VolumeAttachmentLister interface.
type volumeAttachmentLister struct {
	indexer cache.Indexer
}

// NewVolumeAttachmentLister returns a new VolumeAttachmentLister.
func NewVolumeAttachmentLister(indexer cache.Indexer) VolumeAttachmentLister {
	return &volumeAttachmentLister{indexer: indexer}
}

// List lists all VolumeAttachments in the indexer.
func (s *volumeAttachmentLister) List(selector labels.Selector) (ret []*v1beta1.VolumeAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeAttachment))
	})
	return ret, err
}

// Get retrieves the VolumeAttachment from the index for a given name.
func (s *volumeAttachmentLister) Get(name string) (*v1beta1.VolumeAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("volumeattachment"), name)
	}
	return obj.(*v1beta1.VolumeAttachment), nil
}

// VolumeAttachments returns an object that can list and get VolumeAttachments.
func (s *volumeAttachmentLister) VolumeAttachments() VolumeAttachmentTenantLister {
	return volumeAttachmentTenantLister{indexer: s.indexer, tenant: "default"}
}

func (s *volumeAttachmentLister) VolumeAttachmentsWithMultiTenancy(tenant string) VolumeAttachmentTenantLister {
	return volumeAttachmentTenantLister{indexer: s.indexer, tenant: tenant}
}

// VolumeAttachmentTenantLister helps list and get VolumeAttachments.
type VolumeAttachmentTenantLister interface {
	// List lists all VolumeAttachments in the indexer for a given tenant/tenant.
	List(selector labels.Selector) (ret []*v1beta1.VolumeAttachment, err error)
	// Get retrieves the VolumeAttachment from the indexer for a given tenant/tenant and name.
	Get(name string) (*v1beta1.VolumeAttachment, error)
	VolumeAttachmentTenantListerExpansion
}

// volumeAttachmentTenantLister implements the VolumeAttachmentTenantLister
// interface.
type volumeAttachmentTenantLister struct {
	indexer cache.Indexer
	tenant  string
}

// List lists all VolumeAttachments in the indexer for a given tenant.
func (s volumeAttachmentTenantLister) List(selector labels.Selector) (ret []*v1beta1.VolumeAttachment, err error) {
	err = cache.ListAllByTenant(s.indexer, s.tenant, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeAttachment))
	})
	return ret, err
}

// Get retrieves the VolumeAttachment from the indexer for a given tenant and name.
func (s volumeAttachmentTenantLister) Get(name string) (*v1beta1.VolumeAttachment, error) {
	key := s.tenant + "/" + name
	if s.tenant == "default" {
		key = name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("volumeattachment"), name)
	}
	return obj.(*v1beta1.VolumeAttachment), nil
}
