/*
Copyright 2021 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/device-localpv/pkg/apis/openebs.io/device/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeviceNodes implements DeviceNodeInterface
type FakeDeviceNodes struct {
	Fake *FakeLocalV1alpha1
	ns   string
}

var devicenodesResource = schema.GroupVersionResource{Group: "local.openebs.io", Version: "v1alpha1", Resource: "devicenodes"}

var devicenodesKind = schema.GroupVersionKind{Group: "local.openebs.io", Version: "v1alpha1", Kind: "DeviceNode"}

// Get takes name of the deviceNode, and returns the corresponding deviceNode object, and an error if there is any.
func (c *FakeDeviceNodes) Get(name string, options v1.GetOptions) (result *v1alpha1.DeviceNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devicenodesResource, c.ns, name), &v1alpha1.DeviceNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceNode), err
}

// List takes label and field selectors, and returns the list of DeviceNodes that match those selectors.
func (c *FakeDeviceNodes) List(opts v1.ListOptions) (result *v1alpha1.DeviceNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devicenodesResource, devicenodesKind, c.ns, opts), &v1alpha1.DeviceNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeviceNodeList{ListMeta: obj.(*v1alpha1.DeviceNodeList).ListMeta}
	for _, item := range obj.(*v1alpha1.DeviceNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deviceNodes.
func (c *FakeDeviceNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devicenodesResource, c.ns, opts))

}

// Create takes the representation of a deviceNode and creates it.  Returns the server's representation of the deviceNode, and an error, if there is any.
func (c *FakeDeviceNodes) Create(deviceNode *v1alpha1.DeviceNode) (result *v1alpha1.DeviceNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devicenodesResource, c.ns, deviceNode), &v1alpha1.DeviceNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceNode), err
}

// Update takes the representation of a deviceNode and updates it. Returns the server's representation of the deviceNode, and an error, if there is any.
func (c *FakeDeviceNodes) Update(deviceNode *v1alpha1.DeviceNode) (result *v1alpha1.DeviceNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devicenodesResource, c.ns, deviceNode), &v1alpha1.DeviceNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceNode), err
}

// Delete takes name of the deviceNode and deletes it. Returns an error if one occurs.
func (c *FakeDeviceNodes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devicenodesResource, c.ns, name), &v1alpha1.DeviceNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeviceNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devicenodesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeviceNodeList{})
	return err
}

// Patch applies the patch and returns the patched deviceNode.
func (c *FakeDeviceNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeviceNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devicenodesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DeviceNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceNode), err
}
