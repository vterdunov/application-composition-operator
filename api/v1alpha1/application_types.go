/*
Copyright 2025.

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

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	// foo is an example field of Application. Edit application_types.go to remove/update
	// +optional
	Foo *string `json:"foo,omitempty"`

	// +kubebuilder:validation:Required
	Image string `json:"image"`

	PreDeploy *PreDeploySpec `json:"preDeploy,omitempty"`

	// +kubebuilder:validation:Required
	Components []ComponentSpec `json:"components"`
}

type PreDeploySpec struct {
	Command []string `json:"command"`
}

type ComponentSpec struct {
	Name      string                       `json:"name"`
	Replicas  *int32                       `json:"replicas,omitempty"`
	Command   []string                     `json:"command,omitempty"`
	Ports     []PortSpec                   `json:"ports,omitempty"`
	Metrics   *MetricsSpec                 `json:"metrics,omitempty"`
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

type PortSpec struct {
	ContainerPort int32 `json:"containerPort"`
	ExposedPort   int32 `json:"exposedPort,omitempty"`
}

type MetricsSpec struct {
	Enabled *bool  `json:"enabled,omitempty"`
	Port    int32  `json:"port,omitempty"`
	Path    string `json:"path,omitempty"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Application is the Schema for the applications API
type Application struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of Application
	// +required
	Spec ApplicationSpec `json:"spec"`

	// status defines the observed state of Application
	// +optional
	Status ApplicationStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
