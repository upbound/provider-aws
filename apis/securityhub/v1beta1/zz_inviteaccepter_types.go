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

type InviteAccepterInitParameters struct {
}

type InviteAccepterObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the invitation.
	InvitationID *string `json:"invitationId,omitempty" tf:"invitation_id,omitempty"`

	// The account ID of the master Security Hub account whose invitation you're accepting.
	MasterID *string `json:"masterId,omitempty" tf:"master_id,omitempty"`
}

type InviteAccepterParameters struct {

	// The account ID of the master Security Hub account whose invitation you're accepting.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/securityhub/v1beta1.Member
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("master_id",true)
	// +kubebuilder:validation:Optional
	MasterID *string `json:"masterId,omitempty" tf:"master_id,omitempty"`

	// Reference to a Member in securityhub to populate masterId.
	// +kubebuilder:validation:Optional
	MasterIDRef *v1.Reference `json:"masterIdRef,omitempty" tf:"-"`

	// Selector for a Member in securityhub to populate masterId.
	// +kubebuilder:validation:Optional
	MasterIDSelector *v1.Selector `json:"masterIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// InviteAccepterSpec defines the desired state of InviteAccepter
type InviteAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InviteAccepterParameters `json:"forProvider"`
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
	InitProvider InviteAccepterInitParameters `json:"initProvider,omitempty"`
}

// InviteAccepterStatus defines the observed state of InviteAccepter.
type InviteAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InviteAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InviteAccepter is the Schema for the InviteAccepters API. Accepts a Security Hub invitation.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InviteAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InviteAccepterSpec   `json:"spec"`
	Status            InviteAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InviteAccepterList contains a list of InviteAccepters
type InviteAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InviteAccepter `json:"items"`
}

// Repository type metadata.
var (
	InviteAccepter_Kind             = "InviteAccepter"
	InviteAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InviteAccepter_Kind}.String()
	InviteAccepter_KindAPIVersion   = InviteAccepter_Kind + "." + CRDGroupVersion.String()
	InviteAccepter_GroupVersionKind = CRDGroupVersion.WithKind(InviteAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&InviteAccepter{}, &InviteAccepterList{})
}
