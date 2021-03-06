// Copyright Red Hat

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/identitatem/idp-client-api/api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AuthRealms returns a AuthRealmInformer.
	AuthRealms() AuthRealmInformer
	// ClusterOAuths returns a ClusterOAuthInformer.
	ClusterOAuths() ClusterOAuthInformer
	// IDPConfigs returns a IDPConfigInformer.
	IDPConfigs() IDPConfigInformer
	// Strategies returns a StrategyInformer.
	Strategies() StrategyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AuthRealms returns a AuthRealmInformer.
func (v *version) AuthRealms() AuthRealmInformer {
	return &authRealmInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterOAuths returns a ClusterOAuthInformer.
func (v *version) ClusterOAuths() ClusterOAuthInformer {
	return &clusterOAuthInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IDPConfigs returns a IDPConfigInformer.
func (v *version) IDPConfigs() IDPConfigInformer {
	return &iDPConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Strategies returns a StrategyInformer.
func (v *version) Strategies() StrategyInformer {
	return &strategyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
