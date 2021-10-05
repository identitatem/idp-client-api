// Copyright Red Hat

package v1alpha1

import (
	policyv1 "github.com/open-cluster-management/governance-policy-propagator/pkg/apis/policy/v1"
	openshiftconfigv1 "github.com/openshift/api/config/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AuthRealmSpec defines the desired state of AuthRealm
type AuthRealmSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// RouteSubDomain defines a string that will be used for the OAuth name displayed to the
	// user when they login. It will also be used for building the redirect URI callback.
	// The value must be a valid DNS-1123 subdomain. This field is immutable.
	// +required
	// +kubebuilder:validation:MinLength:=1
	// +kubebuilder:validation:MaxLength:=63
	RouteSubDomain string `json:"routeSubDomain"`

	// Placement defines a rule to select a set of ManagedClusters from the ManagedClusterSets bound
	// to the placement namespace.
	PlacementRef corev1.LocalObjectReference `json:"placementRef,omitempty"`

	//RemediateAction defines the remediation action to apply to the idp policy
	// +kubebuilder:validation:Enum=enforce;inform
	// +required
	RemediateAction policyv1.RemediationAction `json:"remediateAction,omitempty"`

	// +kubebuilder:validation:Enum=dex;rhsso
	// +required
	Type AuthProxyType `json:"type,omitempty"`
	//Host defines the url of the proxy
	// +required
	Host string `json:"host,omitempty"`
	//Certificates references a secret containing `ca.crt`, `tls.crt`, and `tls.key`
	CertificatesSecretRef corev1.LocalObjectReference `json:"certificatesSecretRef,omitempty"`
	// IdentityProviders reference an identity provider
	// +optional
	IdentityProviders []openshiftconfigv1.IdentityProvider `json:"identityProviders,omitempty"`
}

type AuthProxyType string

const (
	AuthProxyDex AuthProxyType = "dex"
	// AuthProxyRHSSO AuthProxyType = "rhsso"
)

// AuthRealmStatus defines the observed state of AuthRealm
type AuthRealmStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make generate-clients" to regenerate code after modifying this file

	// Conditions contains the different condition statuses for this AuthRealm.
	// +optional
	Conditions []metav1.Condition `json:"conditions"`
	// +optional
	Cluster []AuthRealmClusterStatus `json:"cluster,omitempty"`
}

const (
	//Applied when the authrealm was correct applied,
	//it does not guaranty that the OAuth gets updated on the managedcluster
	//for that you will have to check the cluster status
	AuthRealmApplied string = "Applied"
)

// AuthRealmClusterStatus defines the status for each cluster
type AuthRealmClusterStatus struct {
	// The name of the cluster
	Name string `json:"name"`
	// Conditions contains the different condition statuses for each cluster for this AuthRealm.
	Conditions metav1.Condition `json:"conditions"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AuthRealm is the Schema for the authrealms API
type AuthRealm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AuthRealmSpec   `json:"spec,omitempty"`
	Status AuthRealmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthRealmList contains a list of AuthRealm
type AuthRealmList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of AuthRealm.
	// +listType=set
	Items []AuthRealm `json:"items"`
}
