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

type DomainDKIMObservation struct {

	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// Find out more about verifying domains in Amazon SES
	// in the AWS SES docs.
	DKIMTokens []*string `json:"dkimTokens,omitempty" tf:"dkim_tokens,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type DomainDKIMParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainDKIMSpec defines the desired state of DomainDKIM
type DomainDKIMSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainDKIMParameters `json:"forProvider"`
}

// DomainDKIMStatus defines the observed state of DomainDKIM.
type DomainDKIMStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainDKIMObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainDKIM is the Schema for the DomainDKIMs API. Provides an SES domain DKIM generation resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainDKIM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainDKIMSpec   `json:"spec"`
	Status            DomainDKIMStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainDKIMList contains a list of DomainDKIMs
type DomainDKIMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainDKIM `json:"items"`
}

// Repository type metadata.
var (
	DomainDKIM_Kind             = "DomainDKIM"
	DomainDKIM_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainDKIM_Kind}.String()
	DomainDKIM_KindAPIVersion   = DomainDKIM_Kind + "." + CRDGroupVersion.String()
	DomainDKIM_GroupVersionKind = CRDGroupVersion.WithKind(DomainDKIM_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainDKIM{}, &DomainDKIMList{})
}
