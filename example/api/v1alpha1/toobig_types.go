/*

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TooBigSpec defines the desired state of TooBig
type TooBigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of TooBig. Edit TooBig_types.go to remove/update
	PodA corev1.PodTemplateSpec `json:"podA"`
	PodB corev1.PodTemplateSpec `json:"podB"`
	PodC corev1.PodTemplateSpec `json:"podC"`
	Foo  string                 `json:"foo,omitempty"`
}

// TooBigStatus defines the observed state of TooBig
type TooBigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// TooBig is the Schema for the toobigs API
type TooBig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TooBigSpec   `json:"spec,omitempty"`
	Status TooBigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TooBigList contains a list of TooBig
type TooBigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TooBig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TooBig{}, &TooBigList{})
}
