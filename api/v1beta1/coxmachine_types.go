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

package v1beta1

import (
	"github.com/platform9/cluster-api-provider-cox/pkg/cloud/coxedge"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MachineFinalizer allows ReconcileCoxMachine to clean up Cox resources
	// associated with CoxCluster before removing it from the apiserver.
	MachineFinalizer = "coxmachine.infrastructure.cluster.x-k8s.io"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CoxMachineSpec defines the desired state of CoxMachine
type CoxMachineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ProviderID is the unique identifier as specified by the cloud provider.
	// +optional
	ProviderID string `json:"providerID,omitempty"`

	// Type represents CoxEdge workload type VM or CONTAINER
	Type string `json:"type,omitempty"`

	// AddAnyCastIPAddress enables AnyCast IP Address
	// +optional
	AddAnyCastIPAddress bool `json:"addanycastipaddress,omitempty"`

	// PersistentStorages mount storage volumes to your workload instances.
	// +optional
	PersistentStorages []coxedge.PersistentStorage `json:"persistentStorages,omitempty"`

	// Expose any ports required by your workload instances
	Ports []coxedge.Port `json:"ports,omitempty"`

	// SSHAuthorizedKeys contains the public SSH keys that should be added to
	// the machine on first boot. In the CoxEdge API this field is equivalent
	// to `firstBootSSHKey`.
	SSHAuthorizedKeys []string `json:"sshAuthorizedKeys,omitempty"`

	// Deployment targets
	Deployments []coxedge.Deployment `json:"deployments,omitempty"`
	Specs       string               `json:"specs,omitempty"`

	// Image is used if Type is set to container then Docker image that will be run in a container. The version can be specified (i.e. nginx:latest).
	Image string `json:"image,omitempty"`

	// Container command
	// +optional
	Commands []string `json:"commands,omitempty"`
	// User data compatible with cloud-init
	UserData string `json:"userData,omitempty"`
}

// CoxMachineStatus defines the observed state of CoxMachine
type CoxMachineStatus struct {
	// Important: Run "make" to regenerate code after modifying this file

	Ready        bool    `json:"ready,omitempty"`
	ErrorMessage *string `json:"errormessage,omitempty"`

	// Addresses contains the IP and/or DNS addresses of the CoxEdge instances.
	Addresses []corev1.NodeAddress `json:"addresses,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".metadata.labels.cluster\\.x-k8s\\.io/cluster-name",description="Cluster to which this CoxMachine belongs"
// +kubebuilder:printcolumn:name="Machine",type="string",JSONPath=".metadata.ownerReferences[?(@.kind==\"Machine\")].name",description="Machine object which owns with this CoxMachine"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.ready",description="Machine ready status"

// CoxMachine is the Schema for the coxmachines API
type CoxMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CoxMachineSpec   `json:"spec,omitempty"`
	Status CoxMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CoxMachineList contains a list of CoxMachine
type CoxMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CoxMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CoxMachine{}, &CoxMachineList{})
}
