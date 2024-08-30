package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LastPassCustomSpec defines the desired state of LastPassCustom
type LastPassCustomSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	CustomRef  CustomRef  `json:"customRef,required"`
	SyncPolicy SyncPolicy `json:"syncPolicy,omitempty"`
}

type CustomRef struct {
	Group string `json:"group,omitempty"`
	Name  string `json:"name,required"`
}

// LastPassCustomStatus defines the observed state of LastPassCustom
type LastPassCustomStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LastPassCustom is the Schema for the lastpasscustoms API
type LastPassCustom struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LastPassCustomSpec   `json:"spec,omitempty"`
	Status LastPassCustomStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LastPassCustomList contains a list of LastPassCustom
type LastPassCustomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LastPassCustom `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LastPassCustom{}, &LastPassCustomList{})
}
