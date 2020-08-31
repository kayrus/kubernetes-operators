/*
Copyright 2017 SAP SE

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
	"context"

	seederv1 "github.com/sapcc/kubernetes-operators/openstack-seeder/pkg/apis/seeder/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOpenstackSeeds implements OpenstackSeedInterface
type FakeOpenstackSeeds struct {
	Fake *FakeOpenstackV1
	ns   string
}

var openstackseedsResource = schema.GroupVersionResource{Group: "openstack.stable.sap.cc", Version: "v1", Resource: "openstackseeds"}

var openstackseedsKind = schema.GroupVersionKind{Group: "openstack.stable.sap.cc", Version: "v1", Kind: "OpenstackSeed"}

// Get takes name of the openstackSeed, and returns the corresponding openstackSeed object, and an error if there is any.
func (c *FakeOpenstackSeeds) Get(ctx context.Context, name string, options v1.GetOptions) (result *seederv1.OpenstackSeed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(openstackseedsResource, c.ns, name), &seederv1.OpenstackSeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*seederv1.OpenstackSeed), err
}

// List takes label and field selectors, and returns the list of OpenstackSeeds that match those selectors.
func (c *FakeOpenstackSeeds) List(ctx context.Context, opts v1.ListOptions) (result *seederv1.OpenstackSeedList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(openstackseedsResource, openstackseedsKind, c.ns, opts), &seederv1.OpenstackSeedList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &seederv1.OpenstackSeedList{ListMeta: obj.(*seederv1.OpenstackSeedList).ListMeta}
	for _, item := range obj.(*seederv1.OpenstackSeedList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested openstackSeeds.
func (c *FakeOpenstackSeeds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(openstackseedsResource, c.ns, opts))

}

// Create takes the representation of a openstackSeed and creates it.  Returns the server's representation of the openstackSeed, and an error, if there is any.
func (c *FakeOpenstackSeeds) Create(ctx context.Context, openstackSeed *seederv1.OpenstackSeed, opts v1.CreateOptions) (result *seederv1.OpenstackSeed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(openstackseedsResource, c.ns, openstackSeed), &seederv1.OpenstackSeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*seederv1.OpenstackSeed), err
}

// Update takes the representation of a openstackSeed and updates it. Returns the server's representation of the openstackSeed, and an error, if there is any.
func (c *FakeOpenstackSeeds) Update(ctx context.Context, openstackSeed *seederv1.OpenstackSeed, opts v1.UpdateOptions) (result *seederv1.OpenstackSeed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(openstackseedsResource, c.ns, openstackSeed), &seederv1.OpenstackSeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*seederv1.OpenstackSeed), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpenstackSeeds) UpdateStatus(ctx context.Context, openstackSeed *seederv1.OpenstackSeed, opts v1.UpdateOptions) (*seederv1.OpenstackSeed, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(openstackseedsResource, "status", c.ns, openstackSeed), &seederv1.OpenstackSeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*seederv1.OpenstackSeed), err
}

// Delete takes name of the openstackSeed and deletes it. Returns an error if one occurs.
func (c *FakeOpenstackSeeds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(openstackseedsResource, c.ns, name), &seederv1.OpenstackSeed{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpenstackSeeds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(openstackseedsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &seederv1.OpenstackSeedList{})
	return err
}

// Patch applies the patch and returns the patched openstackSeed.
func (c *FakeOpenstackSeeds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *seederv1.OpenstackSeed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(openstackseedsResource, c.ns, name, pt, data, subresources...), &seederv1.OpenstackSeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*seederv1.OpenstackSeed), err
}