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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BlobContainerSpec defines the desired state of BlobContainer
type BlobContainerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string `json:"location"`
	ResourceGroup string `json:"resourcegroup,omitempty"`
	AccountName   string `json:"accountname,omitempty"`
	ContainerName string `json:"containername,omitempty"`
}

// BlobContainerStatus defines the observed state of BlobContainer
type BlobContainerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool   `json:"provisioning,omitempty"`
	Provisioned  bool   `json:"provisioned,omitempty"`
	State        string `json:"state,omitempty"`
	Message      string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true

// BlobContainer is the Schema for the blobcontainers API
type BlobContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BlobContainerSpec   `json:"spec,omitempty"`
	Status BlobContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlobContainerList contains a list of BlobContainer
type BlobContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlobContainer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BlobContainer{}, &BlobContainerList{})
}

func (s *BlobContainer) IsSubmitted() bool {
	return s.Status.Provisioned || s.Status.Provisioning
}

func (s *BlobContainer) IsProvisioned() bool {
	return s.Status.Provisioned
}
