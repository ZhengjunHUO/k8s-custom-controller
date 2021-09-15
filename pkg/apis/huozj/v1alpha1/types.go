package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Fufu struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FufuSpec   `json:"spec"`
	Status FufuStatus `json:"status"`
}

type FufuSpec struct {
	Color	string `json:"color"`
	Weight	string `json:"weight"`
}

type FufuStatus struct {
	LastPosition	string `json:"lastPosition"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FufuList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Fufu	`json:"items"`
}
