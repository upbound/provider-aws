/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EnablerObservation struct {

	// Set of account IDs.
	AccountIds []*string `json:"accountIds,omitempty" tf:"account_ids,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Type of resources to scan. Valid values are EC2, ECR, and LAMBDA.
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type EnablerParameters struct {

	// Set of account IDs.
	// +kubebuilder:validation:Optional
	AccountIds []*string `json:"accountIds,omitempty" tf:"account_ids,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Type of resources to scan. Valid values are EC2, ECR, and LAMBDA.
	// +kubebuilder:validation:Optional
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

// EnablerSpec defines the desired state of Enabler
type EnablerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnablerParameters `json:"forProvider"`
}

// EnablerStatus defines the observed state of Enabler.
type EnablerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnablerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Enabler is the Schema for the Enablers API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Enabler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.accountIds)",message="accountIds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.resourceTypes)",message="resourceTypes is a required parameter"
	Spec   EnablerSpec   `json:"spec"`
	Status EnablerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnablerList contains a list of Enablers
type EnablerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Enabler `json:"items"`
}

// Repository type metadata.
var (
	Enabler_Kind             = "Enabler"
	Enabler_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Enabler_Kind}.String()
	Enabler_KindAPIVersion   = Enabler_Kind + "." + CRDGroupVersion.String()
	Enabler_GroupVersionKind = CRDGroupVersion.WithKind(Enabler_Kind)
)

func init() {
	SchemeBuilder.Register(&Enabler{}, &EnablerList{})
}
