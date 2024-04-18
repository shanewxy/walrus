// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	walruscorev1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/walruscore/v1"
	scheme "github.com/seal-io/walrus/pkg/clients/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceRunTemplatesGetter has a method to return a ResourceRunTemplateInterface.
// A group's client should implement this interface.
type ResourceRunTemplatesGetter interface {
	ResourceRunTemplates(namespace string) ResourceRunTemplateInterface
}

// ResourceRunTemplateInterface has methods to work with ResourceRunTemplate resources.
type ResourceRunTemplateInterface interface {
	Create(ctx context.Context, resourceRunTemplate *v1.ResourceRunTemplate, opts metav1.CreateOptions) (*v1.ResourceRunTemplate, error)
	Update(ctx context.Context, resourceRunTemplate *v1.ResourceRunTemplate, opts metav1.UpdateOptions) (*v1.ResourceRunTemplate, error)
	UpdateStatus(ctx context.Context, resourceRunTemplate *v1.ResourceRunTemplate, opts metav1.UpdateOptions) (*v1.ResourceRunTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ResourceRunTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ResourceRunTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ResourceRunTemplate, err error)
	Apply(ctx context.Context, resourceRunTemplate *walruscorev1.ResourceRunTemplateApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceRunTemplate, err error)
	ApplyStatus(ctx context.Context, resourceRunTemplate *walruscorev1.ResourceRunTemplateApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceRunTemplate, err error)
	ResourceRunTemplateExpansion
}

// resourceRunTemplates implements ResourceRunTemplateInterface
type resourceRunTemplates struct {
	client rest.Interface
	ns     string
}

// newResourceRunTemplates returns a ResourceRunTemplates
func newResourceRunTemplates(c *WalruscoreV1Client, namespace string) *resourceRunTemplates {
	return &resourceRunTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceRunTemplate, and returns the corresponding resourceRunTemplate object, and an error if there is any.
func (c *resourceRunTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ResourceRunTemplate, err error) {
	result = &v1.ResourceRunTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceruntemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceRunTemplates that match those selectors.
func (c *resourceRunTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ResourceRunTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ResourceRunTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceruntemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceRunTemplates.
func (c *resourceRunTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourceruntemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a resourceRunTemplate and creates it.  Returns the server's representation of the resourceRunTemplate, and an error, if there is any.
func (c *resourceRunTemplates) Create(ctx context.Context, resourceRunTemplate *v1.ResourceRunTemplate, opts metav1.CreateOptions) (result *v1.ResourceRunTemplate, err error) {
	result = &v1.ResourceRunTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourceruntemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceRunTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a resourceRunTemplate and updates it. Returns the server's representation of the resourceRunTemplate, and an error, if there is any.
func (c *resourceRunTemplates) Update(ctx context.Context, resourceRunTemplate *v1.ResourceRunTemplate, opts metav1.UpdateOptions) (result *v1.ResourceRunTemplate, err error) {
	result = &v1.ResourceRunTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourceruntemplates").
		Name(resourceRunTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceRunTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *resourceRunTemplates) UpdateStatus(ctx context.Context, resourceRunTemplate *v1.ResourceRunTemplate, opts metav1.UpdateOptions) (result *v1.ResourceRunTemplate, err error) {
	result = &v1.ResourceRunTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourceruntemplates").
		Name(resourceRunTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceRunTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the resourceRunTemplate and deletes it. Returns an error if one occurs.
func (c *resourceRunTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceruntemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceRunTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceruntemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched resourceRunTemplate.
func (c *resourceRunTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ResourceRunTemplate, err error) {
	result = &v1.ResourceRunTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourceruntemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied resourceRunTemplate.
func (c *resourceRunTemplates) Apply(ctx context.Context, resourceRunTemplate *walruscorev1.ResourceRunTemplateApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceRunTemplate, err error) {
	if resourceRunTemplate == nil {
		return nil, fmt.Errorf("resourceRunTemplate provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(resourceRunTemplate)
	if err != nil {
		return nil, err
	}
	name := resourceRunTemplate.Name
	if name == nil {
		return nil, fmt.Errorf("resourceRunTemplate.Name must be provided to Apply")
	}
	result = &v1.ResourceRunTemplate{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("resourceruntemplates").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *resourceRunTemplates) ApplyStatus(ctx context.Context, resourceRunTemplate *walruscorev1.ResourceRunTemplateApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ResourceRunTemplate, err error) {
	if resourceRunTemplate == nil {
		return nil, fmt.Errorf("resourceRunTemplate provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(resourceRunTemplate)
	if err != nil {
		return nil, err
	}

	name := resourceRunTemplate.Name
	if name == nil {
		return nil, fmt.Errorf("resourceRunTemplate.Name must be provided to Apply")
	}

	result = &v1.ResourceRunTemplate{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("resourceruntemplates").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}