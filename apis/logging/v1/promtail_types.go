/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PromtailSpec defines the desired state of Promtail
type PromtailSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Promtail. Edit Promtail_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// PromtailStatus defines the observed state of Promtail
type PromtailStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Promtail is the Schema for the promtails API
type Promtail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PromtailSpec   `json:"spec,omitempty"`
	Status PromtailStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PromtailList contains a list of Promtail
type PromtailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Promtail `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Promtail{}, &PromtailList{})
}
