// Copyright Red Hat

package v1alpha1

import (
	openshiftconfigv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterOAuthSpec defines the desired state of ClusterOAuth
type ClusterOAuthSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	OAuth              *openshiftconfigv1.OAuth `json:"oauth,omitempty"`
	AuthRealmReference RelatedObjectReference   `json:"authRealmReference,omitempty"`
	StrategyReference  RelatedObjectReference   `json:"strategyReference,omitempty"`
	DexClientReference RelatedObjectReference   `json:"dexClientReference,omitempty"`
}

type RelatedObjectReference struct {
	// the Kind of the referenced resource
	Kind string `json:"kind,omitempty"`
	// The name of the referenced object
	Name string `json:"name,omitempty"`
	// The namespace of the referenced object
	Namespace string `json:"namespace,omitempty"`
}

// ClusterOAuthStatus defines the observed state of ClusterOAuth
type ClusterOAuthStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make generate-clients" to regenerate code after modifying this file

	// Conditions contains the different condition statuses for this ClusterOAuth.
	// +optional
	Conditions []metav1.Condition `json:"conditions"`
}

const (
	//Applied when the ClusterOAuth was correct applied,
	//it does not guaranty that the OAuth gets updated on the managedcluster
	//for that you will have to check the cluster status
	ClusterOAuthApplied string = "Applied"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterOAuth is the Schema for the authrealms API
type ClusterOAuth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterOAuthSpec   `json:"spec,omitempty"`
	Status ClusterOAuthStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterOAuthList contains a list of ClusterOAuth
type ClusterOAuthList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of ClusterOAuth.
	// +listType=set
	Items []ClusterOAuth `json:"items"`
}
