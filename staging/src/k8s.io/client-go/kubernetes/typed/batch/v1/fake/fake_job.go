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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeJobs implements JobInterface
type FakeJobs struct {
	Fake *FakeBatchV1
	ns   string
	te   string
}

var jobsResource = schema.GroupVersionResource{Group: "batch", Version: "v1", Resource: "jobs"}

var jobsKind = schema.GroupVersionKind{Group: "batch", Version: "v1", Kind: "Job"}

// Get takes name of the job, and returns the corresponding job object, and an error if there is any.
func (c *FakeJobs) Get(name string, options v1.GetOptions) (result *batchv1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithMultiTenancy(jobsResource, c.ns, name, c.te), &batchv1.Job{})

	if obj == nil {
		return nil, err
	}

	return obj.(*batchv1.Job), err
}

// List takes label and field selectors, and returns the list of Jobs that match those selectors.
func (c *FakeJobs) List(opts v1.ListOptions) (result *batchv1.JobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithMultiTenancy(jobsResource, jobsKind, c.ns, opts, c.te), &batchv1.JobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &batchv1.JobList{ListMeta: obj.(*batchv1.JobList).ListMeta}
	for _, item := range obj.(*batchv1.JobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jobs.
func (c *FakeJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithMultiTenancy(jobsResource, c.ns, opts, c.te))

}

// Create takes the representation of a job and creates it.  Returns the server's representation of the job, and an error, if there is any.
func (c *FakeJobs) Create(job *batchv1.Job) (result *batchv1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithMultiTenancy(jobsResource, c.ns, job, c.te), &batchv1.Job{})

	if obj == nil {
		return nil, err
	}

	return obj.(*batchv1.Job), err
}

// Update takes the representation of a job and updates it. Returns the server's representation of the job, and an error, if there is any.
func (c *FakeJobs) Update(job *batchv1.Job) (result *batchv1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithMultiTenancy(jobsResource, c.ns, job, c.te), &batchv1.Job{})

	if obj == nil {
		return nil, err
	}

	return obj.(*batchv1.Job), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJobs) UpdateStatus(job *batchv1.Job) (*batchv1.Job, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithMultiTenancy(jobsResource, "status", c.ns, job, c.te), &batchv1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1.Job), err
}

// Delete takes name of the job and deletes it. Returns an error if one occurs.
func (c *FakeJobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithMultiTenancy(jobsResource, c.ns, name, c.te), &batchv1.Job{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithMultiTenancy(jobsResource, c.ns, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &batchv1.JobList{})
	return err
}

// Patch applies the patch and returns the patched job.
func (c *FakeJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batchv1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithMultiTenancy(jobsResource, c.te, c.ns, name, pt, data, subresources...), &batchv1.Job{})

	if obj == nil {
		return nil, err
	}

	return obj.(*batchv1.Job), err
}
