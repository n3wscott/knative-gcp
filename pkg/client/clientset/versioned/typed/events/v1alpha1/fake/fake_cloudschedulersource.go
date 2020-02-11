/*
Copyright 2020 Google LLC

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
	v1alpha1 "github.com/google/knative-gcp/pkg/apis/events/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudSchedulerSources implements CloudSchedulerSourceInterface
type FakeCloudSchedulerSources struct {
	Fake *FakeEventsV1alpha1
	ns   string
}

var cloudschedulersourcesResource = schema.GroupVersionResource{Group: "events.cloud.google.com", Version: "v1alpha1", Resource: "cloudschedulersources"}

var cloudschedulersourcesKind = schema.GroupVersionKind{Group: "events.cloud.google.com", Version: "v1alpha1", Kind: "CloudSchedulerSource"}

// Get takes name of the cloudSchedulerSource, and returns the corresponding cloudSchedulerSource object, and an error if there is any.
func (c *FakeCloudSchedulerSources) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudSchedulerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudschedulersourcesResource, c.ns, name), &v1alpha1.CloudSchedulerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerSource), err
}

// List takes label and field selectors, and returns the list of CloudSchedulerSources that match those selectors.
func (c *FakeCloudSchedulerSources) List(opts v1.ListOptions) (result *v1alpha1.CloudSchedulerSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudschedulersourcesResource, cloudschedulersourcesKind, c.ns, opts), &v1alpha1.CloudSchedulerSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudSchedulerSourceList{ListMeta: obj.(*v1alpha1.CloudSchedulerSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudSchedulerSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudSchedulerSources.
func (c *FakeCloudSchedulerSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudschedulersourcesResource, c.ns, opts))

}

// Create takes the representation of a cloudSchedulerSource and creates it.  Returns the server's representation of the cloudSchedulerSource, and an error, if there is any.
func (c *FakeCloudSchedulerSources) Create(cloudSchedulerSource *v1alpha1.CloudSchedulerSource) (result *v1alpha1.CloudSchedulerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudschedulersourcesResource, c.ns, cloudSchedulerSource), &v1alpha1.CloudSchedulerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerSource), err
}

// Update takes the representation of a cloudSchedulerSource and updates it. Returns the server's representation of the cloudSchedulerSource, and an error, if there is any.
func (c *FakeCloudSchedulerSources) Update(cloudSchedulerSource *v1alpha1.CloudSchedulerSource) (result *v1alpha1.CloudSchedulerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudschedulersourcesResource, c.ns, cloudSchedulerSource), &v1alpha1.CloudSchedulerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudSchedulerSources) UpdateStatus(cloudSchedulerSource *v1alpha1.CloudSchedulerSource) (*v1alpha1.CloudSchedulerSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudschedulersourcesResource, "status", c.ns, cloudSchedulerSource), &v1alpha1.CloudSchedulerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerSource), err
}

// Delete takes name of the cloudSchedulerSource and deletes it. Returns an error if one occurs.
func (c *FakeCloudSchedulerSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudschedulersourcesResource, c.ns, name), &v1alpha1.CloudSchedulerSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudSchedulerSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudschedulersourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudSchedulerSourceList{})
	return err
}

// Patch applies the patch and returns the patched cloudSchedulerSource.
func (c *FakeCloudSchedulerSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudSchedulerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudschedulersourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudSchedulerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudSchedulerSource), err
}
