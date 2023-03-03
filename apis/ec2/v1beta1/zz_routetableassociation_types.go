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

type RouteTableAssociationObservation struct {

	// The gateway ID to create an association. Conflicts with subnet_id.
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// The ID of the association
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ID of the routing table to associate with.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The subnet ID to create an association. Conflicts with gateway_id.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type RouteTableAssociationParameters struct {

	// The gateway ID to create an association. Conflicts with subnet_id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.InternetGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// Reference to a InternetGateway in ec2 to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDRef *v1.Reference `json:"gatewayIdRef,omitempty" tf:"-"`

	// Selector for a InternetGateway in ec2 to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDSelector *v1.Selector `json:"gatewayIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the routing table to associate with.
	// +crossplane:generate:reference:type=RouteTable
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// The subnet ID to create an association. Conflicts with gateway_id.
	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// RouteTableAssociationSpec defines the desired state of RouteTableAssociation
type RouteTableAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteTableAssociationParameters `json:"forProvider"`
}

// RouteTableAssociationStatus defines the observed state of RouteTableAssociation.
type RouteTableAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteTableAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAssociation is the Schema for the RouteTableAssociations API. Provides a resource to create an association between a route table and a subnet or a route table and an internet gateway or virtual private gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableAssociationSpec   `json:"spec"`
	Status            RouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAssociationList contains a list of RouteTableAssociations
type RouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTableAssociation `json:"items"`
}

// Repository type metadata.
var (
	RouteTableAssociation_Kind             = "RouteTableAssociation"
	RouteTableAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteTableAssociation_Kind}.String()
	RouteTableAssociation_KindAPIVersion   = RouteTableAssociation_Kind + "." + CRDGroupVersion.String()
	RouteTableAssociation_GroupVersionKind = CRDGroupVersion.WithKind(RouteTableAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteTableAssociation{}, &RouteTableAssociationList{})
}
