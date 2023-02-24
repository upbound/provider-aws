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

type CustomerGatewayAssociationObservation struct {

	// The Amazon Resource Name (ARN) of the customer gateway.
	CustomerGatewayArn *string `json:"customerGatewayArn,omitempty" tf:"customer_gateway_arn,omitempty"`

	// The ID of the device.
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// The ID of the global network.
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the link.
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type CustomerGatewayAssociationParameters struct {

	// The Amazon Resource Name (ARN) of the customer gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.CustomerGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	CustomerGatewayArn *string `json:"customerGatewayArn,omitempty" tf:"customer_gateway_arn,omitempty"`

	// Reference to a CustomerGateway in ec2 to populate customerGatewayArn.
	// +kubebuilder:validation:Optional
	CustomerGatewayArnRef *v1.Reference `json:"customerGatewayArnRef,omitempty" tf:"-"`

	// Selector for a CustomerGateway in ec2 to populate customerGatewayArn.
	// +kubebuilder:validation:Optional
	CustomerGatewayArnSelector *v1.Selector `json:"customerGatewayArnSelector,omitempty" tf:"-"`

	// The ID of the device.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.Device
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// Reference to a Device in networkmanager to populate deviceId.
	// +kubebuilder:validation:Optional
	DeviceIDRef *v1.Reference `json:"deviceIdRef,omitempty" tf:"-"`

	// Selector for a Device in networkmanager to populate deviceId.
	// +kubebuilder:validation:Optional
	DeviceIDSelector *v1.Selector `json:"deviceIdSelector,omitempty" tf:"-"`

	// The ID of the global network.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.GlobalNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	// Reference to a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDRef *v1.Reference `json:"globalNetworkIdRef,omitempty" tf:"-"`

	// Selector for a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDSelector *v1.Selector `json:"globalNetworkIdSelector,omitempty" tf:"-"`

	// The ID of the link.
	// +kubebuilder:validation:Optional
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// CustomerGatewayAssociationSpec defines the desired state of CustomerGatewayAssociation
type CustomerGatewayAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomerGatewayAssociationParameters `json:"forProvider"`
}

// CustomerGatewayAssociationStatus defines the observed state of CustomerGatewayAssociation.
type CustomerGatewayAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomerGatewayAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomerGatewayAssociation is the Schema for the CustomerGatewayAssociations API. Associates a customer gateway with a device and optionally, with a link.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CustomerGatewayAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomerGatewayAssociationSpec   `json:"spec"`
	Status            CustomerGatewayAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomerGatewayAssociationList contains a list of CustomerGatewayAssociations
type CustomerGatewayAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomerGatewayAssociation `json:"items"`
}

// Repository type metadata.
var (
	CustomerGatewayAssociation_Kind             = "CustomerGatewayAssociation"
	CustomerGatewayAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomerGatewayAssociation_Kind}.String()
	CustomerGatewayAssociation_KindAPIVersion   = CustomerGatewayAssociation_Kind + "." + CRDGroupVersion.String()
	CustomerGatewayAssociation_GroupVersionKind = CRDGroupVersion.WithKind(CustomerGatewayAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomerGatewayAssociation{}, &CustomerGatewayAssociationList{})
}
