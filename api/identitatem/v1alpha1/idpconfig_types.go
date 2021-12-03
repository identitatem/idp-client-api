// Copyright Red Hat

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IDPConfigSpec defines the desired state of IDPConfig
type IDPConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

}

// IDPConfigStatus defines the observed state of IDPConfig
type IDPConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make generate-clients" to regenerate code after modifying this file

	// Conditions contains the different condition statuses for this IDPConfig.
	// +optional
	Conditions []metav1.Condition `json:"conditions"`
}

const (
	//Applied when the IDPConfig was correct applied,
	//it does not guaranty that the OAuth gets updated on the managedcluster
	//for that you will have to check the cluster status
	IDPConfigApplied string = "Applied"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// IDPConfig is the Schema for the idpconfigs API
type IDPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IDPConfigSpec   `json:"spec,omitempty"`
	Status IDPConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IDPConfigList contains a list of IDPConfig
type IDPConfigList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of IDPConfig.
	// +listType=set
	Items []IDPConfig `json:"items"`
}
