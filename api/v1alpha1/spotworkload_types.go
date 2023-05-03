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

// Component defines the desired state of a component in a workload.
type Component struct {
	// The amount of memory (in GB) allocated to the component.
	Memory int `json:"memory"`
	// The number of virtual CPUs allocated to the component.
	VCPUs int `json:"vCPUs"`
	// The amount of network bandwidth (in Gbps) allocated to the component.
	Network int `json:"network"`
	// The behavior to perform when the component terminates.
	Behavior string `json:"behavior"`
	// The frequency (in minutes) at which to perform the behavior.
	Frequency string `json:"frequency"`
	// The type of storage to use for the component (optional).
	StorageType string `json:"storageType,omitempty"`
	// The name of the component to which this component is affinity with (optional).
	Affinity string `json:"affinity,omitempty"`
	// The name of the component.
	Name string `json:"name"`
	// Whether the component is allowed to burst above its resource allocation (optional).
	Burstable bool `json:"burstable,omitempty"`
	// The name of the component to which this component is anti-affinity with (optional).
	AntiAffinity string `json:"anti-affinity,omitempty"`
}

// SpotWorkloadSpec defines the desired state of SpotWorkload
type SpotWorkloadSpec struct {
	// The name of the application.
	App string `json:"app"`
	// Whether the workload components share resources.
	Share bool `json:"share"`
	// The list of components in the workload.
	Components []Component `json:"components"`
}

// SpotWorkloadStatus defines the observed state of SpotWorkload
type SpotWorkloadStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Stage is the stage of the lifecycle of the workload
	Stage string `json:"stage,omitempty" enum:"Pending|Scheduled|Deployed|Evacuating"`
	// SchedulingTarget is the name of the node the instance is scheduled to
	SchedulingTarget string `json:"scheduling-target,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SpotWorkload is the Schema for the spotworkloads API
type SpotWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpotWorkloadSpec   `json:"spec,omitempty"`
	Status SpotWorkloadStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpotWorkloadList contains a list of SpotWorkload
type SpotWorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotWorkload `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpotWorkload{}, &SpotWorkloadList{})
}
