/*

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
	v1alpha1 "github.com/oam-dev/oam-go-sdk/apis/core.oam.dev/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationConfigurations implements ApplicationConfigurationInterface
type FakeApplicationConfigurations struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var applicationconfigurationsResource = schema.GroupVersionResource{Group: "core.oam.dev", Version: "v1alpha1", Resource: "applicationconfigurations"}

var applicationconfigurationsKind = schema.GroupVersionKind{Group: "core.oam.dev", Version: "v1alpha1", Kind: "ApplicationConfiguration"}

// Get takes name of the applicationConfiguration, and returns the corresponding applicationConfiguration object, and an error if there is any.
func (c *FakeApplicationConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationconfigurationsResource, c.ns, name), &v1alpha1.ApplicationConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationConfiguration), err
}

// List takes label and field selectors, and returns the list of ApplicationConfigurations that match those selectors.
func (c *FakeApplicationConfigurations) List(opts v1.ListOptions) (result *v1alpha1.ApplicationConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationconfigurationsResource, applicationconfigurationsKind, c.ns, opts), &v1alpha1.ApplicationConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationConfigurationList{ListMeta: obj.(*v1alpha1.ApplicationConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationConfigurations.
func (c *FakeApplicationConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a applicationConfiguration and creates it.  Returns the server's representation of the applicationConfiguration, and an error, if there is any.
func (c *FakeApplicationConfigurations) Create(applicationConfiguration *v1alpha1.ApplicationConfiguration) (result *v1alpha1.ApplicationConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationconfigurationsResource, c.ns, applicationConfiguration), &v1alpha1.ApplicationConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationConfiguration), err
}

// Update takes the representation of a applicationConfiguration and updates it. Returns the server's representation of the applicationConfiguration, and an error, if there is any.
func (c *FakeApplicationConfigurations) Update(applicationConfiguration *v1alpha1.ApplicationConfiguration) (result *v1alpha1.ApplicationConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationconfigurationsResource, c.ns, applicationConfiguration), &v1alpha1.ApplicationConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationConfigurations) UpdateStatus(applicationConfiguration *v1alpha1.ApplicationConfiguration) (*v1alpha1.ApplicationConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationconfigurationsResource, "status", c.ns, applicationConfiguration), &v1alpha1.ApplicationConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationConfiguration), err
}

// Delete takes name of the applicationConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeApplicationConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationconfigurationsResource, c.ns, name), &v1alpha1.ApplicationConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationconfigurationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched applicationConfiguration.
func (c *FakeApplicationConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApplicationConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationConfiguration), err
}
