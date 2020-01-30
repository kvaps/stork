/*
Copyright 2018 Openstorage.org

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
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataExports implements DataExportInterface
type FakeDataExports struct {
	Fake *FakeStorkV1alpha1
	ns   string
}

var dataexportsResource = schema.GroupVersionResource{Group: "stork.libopenstorage.org", Version: "v1alpha1", Resource: "dataexports"}

var dataexportsKind = schema.GroupVersionKind{Group: "stork.libopenstorage.org", Version: "v1alpha1", Kind: "DataExport"}

// Get takes name of the dataExport, and returns the corresponding dataExport object, and an error if there is any.
func (c *FakeDataExports) Get(name string, options v1.GetOptions) (result *v1alpha1.DataExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataexportsResource, c.ns, name), &v1alpha1.DataExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataExport), err
}

// List takes label and field selectors, and returns the list of DataExports that match those selectors.
func (c *FakeDataExports) List(opts v1.ListOptions) (result *v1alpha1.DataExportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataexportsResource, dataexportsKind, c.ns, opts), &v1alpha1.DataExportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataExportList{ListMeta: obj.(*v1alpha1.DataExportList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataExportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataExports.
func (c *FakeDataExports) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataexportsResource, c.ns, opts))

}

// Create takes the representation of a dataExport and creates it.  Returns the server's representation of the dataExport, and an error, if there is any.
func (c *FakeDataExports) Create(dataExport *v1alpha1.DataExport) (result *v1alpha1.DataExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataexportsResource, c.ns, dataExport), &v1alpha1.DataExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataExport), err
}

// Update takes the representation of a dataExport and updates it. Returns the server's representation of the dataExport, and an error, if there is any.
func (c *FakeDataExports) Update(dataExport *v1alpha1.DataExport) (result *v1alpha1.DataExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataexportsResource, c.ns, dataExport), &v1alpha1.DataExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataExport), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataExports) UpdateStatus(dataExport *v1alpha1.DataExport) (*v1alpha1.DataExport, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataexportsResource, "status", c.ns, dataExport), &v1alpha1.DataExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataExport), err
}

// Delete takes name of the dataExport and deletes it. Returns an error if one occurs.
func (c *FakeDataExports) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dataexportsResource, c.ns, name), &v1alpha1.DataExport{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataExports) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataexportsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataExportList{})
	return err
}

// Patch applies the patch and returns the patched dataExport.
func (c *FakeDataExports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataexportsResource, c.ns, name, data, subresources...), &v1alpha1.DataExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataExport), err
}
