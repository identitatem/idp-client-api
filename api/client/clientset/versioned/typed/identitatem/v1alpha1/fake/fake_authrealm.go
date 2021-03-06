// Copyright Red Hat

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/identitatem/idp-client-api/api/identitatem/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuthRealms implements AuthRealmInterface
type FakeAuthRealms struct {
	Fake *FakeIdentityconfigV1alpha1
	ns   string
}

var authrealmsResource = schema.GroupVersionResource{Group: "identityconfig.identitatem.io", Version: "v1alpha1", Resource: "authrealms"}

var authrealmsKind = schema.GroupVersionKind{Group: "identityconfig.identitatem.io", Version: "v1alpha1", Kind: "AuthRealm"}

// Get takes name of the authRealm, and returns the corresponding authRealm object, and an error if there is any.
func (c *FakeAuthRealms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AuthRealm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(authrealmsResource, c.ns, name), &v1alpha1.AuthRealm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthRealm), err
}

// List takes label and field selectors, and returns the list of AuthRealms that match those selectors.
func (c *FakeAuthRealms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AuthRealmList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(authrealmsResource, authrealmsKind, c.ns, opts), &v1alpha1.AuthRealmList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AuthRealmList{ListMeta: obj.(*v1alpha1.AuthRealmList).ListMeta}
	for _, item := range obj.(*v1alpha1.AuthRealmList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authRealms.
func (c *FakeAuthRealms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(authrealmsResource, c.ns, opts))

}

// Create takes the representation of a authRealm and creates it.  Returns the server's representation of the authRealm, and an error, if there is any.
func (c *FakeAuthRealms) Create(ctx context.Context, authRealm *v1alpha1.AuthRealm, opts v1.CreateOptions) (result *v1alpha1.AuthRealm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(authrealmsResource, c.ns, authRealm), &v1alpha1.AuthRealm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthRealm), err
}

// Update takes the representation of a authRealm and updates it. Returns the server's representation of the authRealm, and an error, if there is any.
func (c *FakeAuthRealms) Update(ctx context.Context, authRealm *v1alpha1.AuthRealm, opts v1.UpdateOptions) (result *v1alpha1.AuthRealm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(authrealmsResource, c.ns, authRealm), &v1alpha1.AuthRealm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthRealm), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAuthRealms) UpdateStatus(ctx context.Context, authRealm *v1alpha1.AuthRealm, opts v1.UpdateOptions) (*v1alpha1.AuthRealm, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(authrealmsResource, "status", c.ns, authRealm), &v1alpha1.AuthRealm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthRealm), err
}

// Delete takes name of the authRealm and deletes it. Returns an error if one occurs.
func (c *FakeAuthRealms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(authrealmsResource, c.ns, name, opts), &v1alpha1.AuthRealm{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuthRealms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(authrealmsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AuthRealmList{})
	return err
}

// Patch applies the patch and returns the patched authRealm.
func (c *FakeAuthRealms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AuthRealm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(authrealmsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AuthRealm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthRealm), err
}
