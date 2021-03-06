// Copyright Red Hat

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/identitatem/idp-client-api/api/identitatem/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IDPConfigLister helps list IDPConfigs.
// All objects returned here must be treated as read-only.
type IDPConfigLister interface {
	// List lists all IDPConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IDPConfig, err error)
	// IDPConfigs returns an object that can list and get IDPConfigs.
	IDPConfigs(namespace string) IDPConfigNamespaceLister
	IDPConfigListerExpansion
}

// iDPConfigLister implements the IDPConfigLister interface.
type iDPConfigLister struct {
	indexer cache.Indexer
}

// NewIDPConfigLister returns a new IDPConfigLister.
func NewIDPConfigLister(indexer cache.Indexer) IDPConfigLister {
	return &iDPConfigLister{indexer: indexer}
}

// List lists all IDPConfigs in the indexer.
func (s *iDPConfigLister) List(selector labels.Selector) (ret []*v1alpha1.IDPConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IDPConfig))
	})
	return ret, err
}

// IDPConfigs returns an object that can list and get IDPConfigs.
func (s *iDPConfigLister) IDPConfigs(namespace string) IDPConfigNamespaceLister {
	return iDPConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IDPConfigNamespaceLister helps list and get IDPConfigs.
// All objects returned here must be treated as read-only.
type IDPConfigNamespaceLister interface {
	// List lists all IDPConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IDPConfig, err error)
	// Get retrieves the IDPConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IDPConfig, error)
	IDPConfigNamespaceListerExpansion
}

// iDPConfigNamespaceLister implements the IDPConfigNamespaceLister
// interface.
type iDPConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IDPConfigs in the indexer for a given namespace.
func (s iDPConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IDPConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IDPConfig))
	})
	return ret, err
}

// Get retrieves the IDPConfig from the indexer for a given namespace and name.
func (s iDPConfigNamespaceLister) Get(name string) (*v1alpha1.IDPConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("idpconfig"), name)
	}
	return obj.(*v1alpha1.IDPConfig), nil
}
