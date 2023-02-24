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

type ConnectionAssociationObservation struct {

	// The ID of the connection.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the LAG with which to associate the connection.
	LagID *string `json:"lagId,omitempty" tf:"lag_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type ConnectionAssociationParameters struct {

	// The ID of the connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.Connection
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Reference to a Connection in directconnect to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// Selector for a Connection in directconnect to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// The ID of the LAG with which to associate the connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.Lag
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LagID *string `json:"lagId,omitempty" tf:"lag_id,omitempty"`

	// Reference to a Lag in directconnect to populate lagId.
	// +kubebuilder:validation:Optional
	LagIDRef *v1.Reference `json:"lagIdRef,omitempty" tf:"-"`

	// Selector for a Lag in directconnect to populate lagId.
	// +kubebuilder:validation:Optional
	LagIDSelector *v1.Selector `json:"lagIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ConnectionAssociationSpec defines the desired state of ConnectionAssociation
type ConnectionAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionAssociationParameters `json:"forProvider"`
}

// ConnectionAssociationStatus defines the observed state of ConnectionAssociation.
type ConnectionAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionAssociation is the Schema for the ConnectionAssociations API. Associates a Direct Connect Connection with a LAG.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConnectionAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionAssociationSpec   `json:"spec"`
	Status            ConnectionAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionAssociationList contains a list of ConnectionAssociations
type ConnectionAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionAssociation `json:"items"`
}

// Repository type metadata.
var (
	ConnectionAssociation_Kind             = "ConnectionAssociation"
	ConnectionAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectionAssociation_Kind}.String()
	ConnectionAssociation_KindAPIVersion   = ConnectionAssociation_Kind + "." + CRDGroupVersion.String()
	ConnectionAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ConnectionAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectionAssociation{}, &ConnectionAssociationList{})
}
