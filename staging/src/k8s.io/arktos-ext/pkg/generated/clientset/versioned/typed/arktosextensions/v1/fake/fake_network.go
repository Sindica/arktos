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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	arktosextensionsv1 "k8s.io/arktos-ext/pkg/apis/arktosextensions/v1"
	testing "k8s.io/client-go/testing"
)

// FakeNetworks implements NetworkInterface
type FakeNetworks struct {
	Fake *FakeArktosV1
	te   string
}

var networksResource = schema.GroupVersionResource{Group: "arktos.futurewei.com", Version: "v1", Resource: "networks"}

var networksKind = schema.GroupVersionKind{Group: "arktos.futurewei.com", Version: "v1", Kind: "Network"}

// Get takes name of the network, and returns the corresponding network object, and an error if there is any.
func (c *FakeNetworks) Get(name string, options v1.GetOptions) (result *arktosextensionsv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewTenantGetAction(networksResource, name, c.te), &arktosextensionsv1.Network{})

	if obj == nil {
		return nil, err
	}

	return obj.(*arktosextensionsv1.Network), err
}

// List takes label and field selectors, and returns the list of Networks that match those selectors.
func (c *FakeNetworks) List(opts v1.ListOptions) (result *arktosextensionsv1.NetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewTenantListAction(networksResource, networksKind, opts, c.te), &arktosextensionsv1.NetworkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &arktosextensionsv1.NetworkList{ListMeta: obj.(*arktosextensionsv1.NetworkList).ListMeta}
	for _, item := range obj.(*arktosextensionsv1.NetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networks.
func (c *FakeNetworks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewTenantWatchAction(networksResource, opts, c.te))
}

// Create takes the representation of a network and creates it.  Returns the server's representation of the network, and an error, if there is any.
func (c *FakeNetworks) Create(network *arktosextensionsv1.Network) (result *arktosextensionsv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewTenantCreateAction(networksResource, network, c.te), &arktosextensionsv1.Network{})

	if obj == nil {
		return nil, err
	}

	return obj.(*arktosextensionsv1.Network), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworks) UpdateStatus(network *arktosextensionsv1.Network) (*arktosextensionsv1.Network, error) {
	obj, err := c.Fake.
		Invokes(testing.NewTenantUpdateSubresourceAction(networksResource, "status", network, c.te), &arktosextensionsv1.Network{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arktosextensionsv1.Network), err
}
