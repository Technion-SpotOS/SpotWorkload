/*
Copyright 2023.

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

// ComponentSpec defines the requirements of a component in a workload.
type ComponentSpec struct {
	// VCPUs is the required number of virtual CPUs for component.
	VCPUs int `json:"vCPUs"`
	// Memory is the required memory for component (GB).
	Memory int `json:"memory"`
	// Network is the required network bandwidth for component (Gbps).
	Network int `json:"network"`
	// Behavior is the required interruption behavior: options: terminate,stop,hibernation
	Behavior string `json:"behavior" enum:"terminate|stop|hibernation" default:"terminate"`
	// Frequency is the limit interruption frequency of the instances. options: 0-4.
	Frequency int `json:"frequency" enum:"0|1|2|3|4" default:"0"`
	// The type of storage to use for the component (optional).
	StorageType string `json:"storageType,omitempty"`
	// Affinity is the components names that must be on the same instance.
	Affinity []string `json:"affinity,omitempty"`
	// AntiAffinity is the components names that must be on different instances.
	AntiAffinity []string `json:"anti-affinity,omitempty"`
}

// ComponentStatus defines the observed state of a component in a workload.
type ComponentStatus struct {
	// Stage is the stage of the lifecycle of the workload
	Stage string `json:"stage,omitempty" enum:"pending|scheduled|deployed|evacuated"`
	// InstanceName is the name of the node the instance is scheduled to
	InstanceName string `json:"instance-name,omitempty"`
}

// SpotWorkloadSpec defines the desired state of SpotWorkload.
type SpotWorkloadSpec struct {
	// The name of the application.
	App string `json:"app"`
	// Whether the workload components share resources.
	Share bool `json:"share"`
	// The list of components in the workload. Component names (keys) must match relevant deployment names.
	Components map[string]ComponentSpec `json:"components"`
}

// SpotWorkloadStatus defines the observed state of SpotWorkload.
type SpotWorkloadStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Components is the status of the components in the workload.
	Components map[string]ComponentStatus `json:"components"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SpotWorkload is the Schema for the spotworkloads API.
type SpotWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpotWorkloadSpec   `json:"spec,omitempty"`
	Status SpotWorkloadStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpotWorkloadList contains a list of SpotWorkload.
type SpotWorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotWorkload `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpotWorkload{}, &SpotWorkloadList{})
}
