/*
Copyright 2021 Flant CJSC

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

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodLoggingConfig specify target for kubernetes pods logs collecting in specified namespace
type PodLoggingConfig struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a namespaced log source.
	Spec PodLoggingConfigSpec `json:"spec"`

	// Most recently observed status of a namespaced log source.
	Status PodLoggingConfigStatus `json:"status,omitempty"`
}

type PodLoggingConfigSpec struct {
	// LabelSelector filter pods by label
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty"`

	// ClusterDestinationRefs slice of ClusterLogDestination names
	ClusterDestinationRefs []string `json:"clusterDestinationRefs,omitempty"`
}

type PodLoggingConfigStatus struct {
}
