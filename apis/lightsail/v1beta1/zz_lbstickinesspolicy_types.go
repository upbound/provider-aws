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

type LBStickinessPolicyObservation struct {

	// The cookie duration in seconds. This determines the length of the session stickiness.
	CookieDuration *float64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// - The Session Stickiness state of the load balancer. true to activate session stickiness or false to deactivate session stickiness.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name used for this load balancer (matches lb_name).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type LBStickinessPolicyParameters struct {

	// The cookie duration in seconds. This determines the length of the session stickiness.
	// +kubebuilder:validation:Required
	CookieDuration *float64 `json:"cookieDuration" tf:"cookie_duration,omitempty"`

	// - The Session Stickiness state of the load balancer. true to activate session stickiness or false to deactivate session stickiness.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LBStickinessPolicySpec defines the desired state of LBStickinessPolicy
type LBStickinessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBStickinessPolicyParameters `json:"forProvider"`
}

// LBStickinessPolicyStatus defines the observed state of LBStickinessPolicy.
type LBStickinessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBStickinessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBStickinessPolicy is the Schema for the LBStickinessPolicys API. Configures Session Stickiness for a Lightsail Load Balancer
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBStickinessPolicySpec   `json:"spec"`
	Status            LBStickinessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBStickinessPolicyList contains a list of LBStickinessPolicys
type LBStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBStickinessPolicy `json:"items"`
}

// Repository type metadata.
var (
	LBStickinessPolicy_Kind             = "LBStickinessPolicy"
	LBStickinessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBStickinessPolicy_Kind}.String()
	LBStickinessPolicy_KindAPIVersion   = LBStickinessPolicy_Kind + "." + CRDGroupVersion.String()
	LBStickinessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(LBStickinessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LBStickinessPolicy{}, &LBStickinessPolicyList{})
}
