// Copyright Red Hat

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/identitatem/idp-client-api/api/client/clientset/versioned"
	internalinterfaces "github.com/identitatem/idp-client-api/api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/identitatem/idp-client-api/api/client/listers/identitatem/v1alpha1"
	identitatemv1alpha1 "github.com/identitatem/idp-client-api/api/identitatem/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterOAuthInformer provides access to a shared informer and lister for
// ClusterOAuths.
type ClusterOAuthInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterOAuthLister
}

type clusterOAuthInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterOAuthInformer constructs a new informer for ClusterOAuth type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterOAuthInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterOAuthInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterOAuthInformer constructs a new informer for ClusterOAuth type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterOAuthInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IdentityconfigV1alpha1().ClusterOAuths(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IdentityconfigV1alpha1().ClusterOAuths(namespace).Watch(context.TODO(), options)
			},
		},
		&identitatemv1alpha1.ClusterOAuth{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterOAuthInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterOAuthInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterOAuthInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&identitatemv1alpha1.ClusterOAuth{}, f.defaultInformer)
}

func (f *clusterOAuthInformer) Lister() v1alpha1.ClusterOAuthLister {
	return v1alpha1.NewClusterOAuthLister(f.Informer().GetIndexer())
}
