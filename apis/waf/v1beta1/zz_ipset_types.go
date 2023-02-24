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

type IPSetDescriptorsObservation struct {

	// Type of the IP address - IPV4 or IPV6.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// An IPv4 or IPv6 address specified via CIDR notationE.g., 192.0.2.44/32 or 1111:0000:0000:0000:0000:0000:0000:0000/64
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type IPSetDescriptorsParameters struct {

	// Type of the IP address - IPV4 or IPV6.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// An IPv4 or IPv6 address specified via CIDR notationE.g., 192.0.2.44/32 or 1111:0000:0000:0000:0000:0000:0000:0000/64
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type IPSetObservation struct {

	// The ARN of the WAF IPSet.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the WAF IPSet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IPSetDescriptors []IPSetDescriptorsObservation `json:"ipSetDescriptors,omitempty" tf:"ip_set_descriptors,omitempty"`

	// The name or description of the IPSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type IPSetParameters struct {

	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	// +kubebuilder:validation:Optional
	IPSetDescriptors []IPSetDescriptorsParameters `json:"ipSetDescriptors,omitempty" tf:"ip_set_descriptors,omitempty"`

	// The name or description of the IPSet.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// IPSetSpec defines the desired state of IPSet
type IPSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPSetParameters `json:"forProvider"`
}

// IPSetStatus defines the observed state of IPSet.
type IPSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPSet is the Schema for the IPSets API. Provides a AWS WAF IPSet resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IPSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPSetSpec   `json:"spec"`
	Status            IPSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPSetList contains a list of IPSets
type IPSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPSet `json:"items"`
}

// Repository type metadata.
var (
	IPSet_Kind             = "IPSet"
	IPSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPSet_Kind}.String()
	IPSet_KindAPIVersion   = IPSet_Kind + "." + CRDGroupVersion.String()
	IPSet_GroupVersionKind = CRDGroupVersion.WithKind(IPSet_Kind)
)

func init() {
	SchemeBuilder.Register(&IPSet{}, &IPSetList{})
}
