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

type EmailIdentityFeedbackAttributesObservation struct {

	// Sets the feedback forwarding configuration for the identity.
	EmailForwardingEnabled *bool `json:"emailForwardingEnabled,omitempty" tf:"email_forwarding_enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type EmailIdentityFeedbackAttributesParameters struct {

	// Sets the feedback forwarding configuration for the identity.
	// +kubebuilder:validation:Optional
	EmailForwardingEnabled *bool `json:"emailForwardingEnabled,omitempty" tf:"email_forwarding_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// EmailIdentityFeedbackAttributesSpec defines the desired state of EmailIdentityFeedbackAttributes
type EmailIdentityFeedbackAttributesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailIdentityFeedbackAttributesParameters `json:"forProvider"`
}

// EmailIdentityFeedbackAttributesStatus defines the observed state of EmailIdentityFeedbackAttributes.
type EmailIdentityFeedbackAttributesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailIdentityFeedbackAttributesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIdentityFeedbackAttributes is the Schema for the EmailIdentityFeedbackAttributess API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EmailIdentityFeedbackAttributes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmailIdentityFeedbackAttributesSpec   `json:"spec"`
	Status            EmailIdentityFeedbackAttributesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIdentityFeedbackAttributesList contains a list of EmailIdentityFeedbackAttributess
type EmailIdentityFeedbackAttributesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailIdentityFeedbackAttributes `json:"items"`
}

// Repository type metadata.
var (
	EmailIdentityFeedbackAttributes_Kind             = "EmailIdentityFeedbackAttributes"
	EmailIdentityFeedbackAttributes_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailIdentityFeedbackAttributes_Kind}.String()
	EmailIdentityFeedbackAttributes_KindAPIVersion   = EmailIdentityFeedbackAttributes_Kind + "." + CRDGroupVersion.String()
	EmailIdentityFeedbackAttributes_GroupVersionKind = CRDGroupVersion.WithKind(EmailIdentityFeedbackAttributes_Kind)
)

func init() {
	SchemeBuilder.Register(&EmailIdentityFeedbackAttributes{}, &EmailIdentityFeedbackAttributesList{})
}
