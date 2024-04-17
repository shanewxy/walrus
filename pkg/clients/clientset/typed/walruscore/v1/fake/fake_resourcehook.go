// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	walruscorev1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/walruscore/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourceHooks implements ResourceHookInterface
type FakeResourceHooks struct {
	Fake *FakeWalruscoreV1
	ns   string
}

var resourcehooksResource = v1.SchemeGroupVersion.WithResource("resourcehooks")

var resourcehooksKind = v1.SchemeGroupVersion.WithKind("ResourceHook")

// Get takes name of the resourceHook, and returns the corresponding resourceHook object, and an error if there is any.
func (c *FakeResourceHooks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ResourceHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcehooksResource, c.ns, name), &v1.ResourceHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceHook), err
}

// List takes label and field selectors, and returns the list of ResourceHooks that match those selectors.
func (c *FakeResourceHooks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ResourceHookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcehooksResource, resourcehooksKind, c.ns, opts), &v1.ResourceHookList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ResourceHookList{ListMeta: obj.(*v1.ResourceHookList).ListMeta}
	for _, item := range obj.(*v1.ResourceHookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceHooks.
func (c *FakeResourceHooks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcehooksResource, c.ns, opts))

}

// Create takes the representation of a resourceHook and creates it.  Returns the server's representation of the resourceHook, and an error, if there is any.
func (c *FakeResourceHooks) Create(ctx context.Context, resourceHook *v1.ResourceHook, opts metav1.CreateOptions) (result *v1.ResourceHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcehooksResource, c.ns, resourceHook), &v1.ResourceHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceHook), err
}

// Update takes the representation of a resourceHook and updates it. Returns the server's representation of the resourceHook, and an error, if there is any.
func (c *FakeResourceHooks) Update(ctx context.Context, resourceHook *v1.ResourceHook, opts metav1.UpdateOptions) (result *v1.ResourceHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcehooksResource, c.ns, resourceHook), &v1.ResourceHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceHook), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceHooks) UpdateStatus(ctx context.Context, resourceHook *v1.ResourceHook, opts metav1.UpdateOptions) (*v1.ResourceHook, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcehooksResource, "status", c.ns, resourceHook), &v1.ResourceHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceHook), err
}

// Delete takes name of the resourceHook and deletes it. Returns an error if one occurs.
func (c *FakeResourceHooks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourcehooksResource, c.ns, name, opts), &v1.ResourceHook{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceHooks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcehooksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ResourceHookList{})
	return err
}

// Patch applies the patch and returns the patched resourceHook.
func (c *FakeResourceHooks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ResourceHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcehooksResource, c.ns, name, pt, data, subresources...), &v1.ResourceHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceHook), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied resourceHook.
func (c *FakeResourceHooks) Apply(ctx context.Context, resourceHook *walruscorev1.ResourceHookApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceHook, err error) {
	if resourceHook == nil {
		return nil, fmt.Errorf("resourceHook provided to Apply must not be nil")
	}
	data, err := json.Marshal(resourceHook)
	if err != nil {
		return nil, err
	}
	name := resourceHook.Name
	if name == nil {
		return nil, fmt.Errorf("resourceHook.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcehooksResource, c.ns, *name, types.ApplyPatchType, data), &v1.ResourceHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceHook), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeResourceHooks) ApplyStatus(ctx context.Context, resourceHook *walruscorev1.ResourceHookApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceHook, err error) {
	if resourceHook == nil {
		return nil, fmt.Errorf("resourceHook provided to Apply must not be nil")
	}
	data, err := json.Marshal(resourceHook)
	if err != nil {
		return nil, err
	}
	name := resourceHook.Name
	if name == nil {
		return nil, fmt.Errorf("resourceHook.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcehooksResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.ResourceHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceHook), err
}
