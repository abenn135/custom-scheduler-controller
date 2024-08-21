/*
Copyright 2024.

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
	schedv1 "k8s.io/kube-scheduler/config/v1"
)

// SchedulerControllerSpec defines the desired state of SchedulerController.
type SchedulerControllerSpec struct {
	// This specifies the custom scheduler configuration.
	//
	// N.B.: Besides the default scheduler policy (LeastAllocated), only two custom scheduler configurations are supported:
	//
	// 1. Bin-packing
	// 2. Co-scheduling
	//
	// See documentation for more details.
	Spec schedv1.KubeSchedulerConfiguration `json:"spec,omitempty" protobuf:"bytes,1,opt,name=spec,proto3"`
}

// SchedulerControllerStatus defines the observed state of SchedulerController
type SchedulerControllerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SchedulerController is the Schema for the schedulercontrollers API
type SchedulerController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SchedulerControllerSpec   `json:"spec,omitempty"`
	Status SchedulerControllerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchedulerControllerList contains a list of SchedulerController
type SchedulerControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SchedulerController `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SchedulerController{}, &SchedulerControllerList{})
}
