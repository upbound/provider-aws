// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type EmailIdentityFeedbackAttributesInitParameters struct {

	// Sets the feedback forwarding configuration for the identity.
	EmailForwardingEnabled *bool `json:"emailForwardingEnabled,omitempty" tf:"email_forwarding_enabled,omitempty"`
}

type EmailIdentityFeedbackAttributesObservation struct {

	// Sets the feedback forwarding configuration for the identity.
	EmailForwardingEnabled *bool `json:"emailForwardingEnabled,omitempty" tf:"email_forwarding_enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
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
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider EmailIdentityFeedbackAttributesInitParameters `json:"initProvider,omitempty"`
}

// EmailIdentityFeedbackAttributesStatus defines the observed state of EmailIdentityFeedbackAttributes.
type EmailIdentityFeedbackAttributesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailIdentityFeedbackAttributesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EmailIdentityFeedbackAttributes is the Schema for the EmailIdentityFeedbackAttributess API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
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
