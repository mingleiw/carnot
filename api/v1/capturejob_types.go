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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CaptureJobSpec defines the desired state of CaptureJob
type CaptureJobSpec struct {
	//+kubebuilder:default=8080
	//+kubebuilder:validation:Minimum=0
	ListeningPort int32 `json:"listeningPort,omitempty"`

	// +kubebuilder:default=100
	// +kubebuilder:validation:Minimum=0
	TargetNumOfRequests int32 `json:"targetNumOfRequests,omitempty"`
}

// CaptureJobStatus defines the observed state of CaptureJob
type CaptureJobStatus struct {
	// The number of requests that job captured
	CapturedNumOfRequests int32 `json:"capturedNumOfRequests,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CaptureJob is the Schema for the capturejobs API
type CaptureJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CaptureJobSpec   `json:"spec,omitempty"`
	Status CaptureJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CaptureJobList contains a list of CaptureJob
type CaptureJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CaptureJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CaptureJob{}, &CaptureJobList{})
}
