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

type IdentityPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IdentityPolicyParameters struct {

	// Name or Amazon Resource Name (ARN) of the SES Identity.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ses/v1beta1.DomainIdentity
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// Reference to a DomainIdentity in ses to populate identity.
	// +kubebuilder:validation:Optional
	IdentityRef *v1.Reference `json:"identityRef,omitempty" tf:"-"`

	// Selector for a DomainIdentity in ses to populate identity.
	// +kubebuilder:validation:Optional
	IdentitySelector *v1.Selector `json:"identitySelector,omitempty" tf:"-"`

	// Name of the policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// JSON string of the policy.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// IdentityPolicySpec defines the desired state of IdentityPolicy
type IdentityPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityPolicyParameters `json:"forProvider"`
}

// IdentityPolicyStatus defines the observed state of IdentityPolicy.
type IdentityPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityPolicy is the Schema for the IdentityPolicys API. Manages a SES Identity Policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IdentityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.policy)",message="policy is a required parameter"
	Spec   IdentityPolicySpec   `json:"spec"`
	Status IdentityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityPolicyList contains a list of IdentityPolicys
type IdentityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityPolicy `json:"items"`
}

// Repository type metadata.
var (
	IdentityPolicy_Kind             = "IdentityPolicy"
	IdentityPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityPolicy_Kind}.String()
	IdentityPolicy_KindAPIVersion   = IdentityPolicy_Kind + "." + CRDGroupVersion.String()
	IdentityPolicy_GroupVersionKind = CRDGroupVersion.WithKind(IdentityPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityPolicy{}, &IdentityPolicyList{})
}
