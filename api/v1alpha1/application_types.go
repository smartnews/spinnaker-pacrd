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

// Permissions TODO
type Permissions struct {
	// +optional
	// Read TODO
	Read []string `json:"READ,omitempty"`
	// +optional
	// Write TODO
	Write []string `json:"WRITE,omitempty"`
	// +optional
	// Execute TODO
	Execute []string `json:"EXECUTE,omitempty"`
}

// DataSources TODO
type DataSources struct {
	// +optional
	Enabled []string `json:"enabled,omitempty"`
	// +optional
	Disabled []string `json:"disabled,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Email points to the e-mail user or list that owns this application.
	Email string `json:"email,omitempty"`
	// Description explains the purpose of this application.
	Description string `json:"description,omitempty"`
	// User is...? TODO
	User string `json:"user,omitempty"`

	// +optional
	// DataSources TODO
	DataSources `json:"dataSources,omitempty"`

	// +optional
	// DataSources TODO
	Permissions `json:"permissions,omitempty"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// +optional
	// LastConfigured represents the last time the operator updated this application in Spinnaker.
	LastConfigured metav1.Time `json:"lastConfigured,omitempty"`
	// Status represents the current status of this application.
	Status string `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".status.url",name=URL,type=string
// +kubebuilder:printcolumn:name="lastConfigured",type="date",JSONPath=".status.lastConfigured",description="Last Configured"
// +kubebuilder:printcolumn:name="project",type="string",JSONPath=".spec.project",description="Project"
// +kubebuilder:printcolumn:name="status",type="string",JSONPath=".status.status",description="Status"
// +kubebuilder:resource:path=applications,shortName=app

// Application is the Schema for the applications API
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
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
