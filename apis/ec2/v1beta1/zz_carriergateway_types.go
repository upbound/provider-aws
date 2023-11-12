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

type CarrierGatewayInitParameters struct {

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CarrierGatewayObservation struct {

	// The ARN of the carrier gateway.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the carrier gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS account ID of the owner of the carrier gateway.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the VPC to associate with the carrier gateway.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type CarrierGatewayParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the VPC to associate with the carrier gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// CarrierGatewaySpec defines the desired state of CarrierGateway
type CarrierGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CarrierGatewayParameters `json:"forProvider"`
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
	InitProvider CarrierGatewayInitParameters `json:"initProvider,omitempty"`
}

// CarrierGatewayStatus defines the observed state of CarrierGateway.
type CarrierGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CarrierGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CarrierGateway is the Schema for the CarrierGateways API. Manages an EC2 Carrier Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CarrierGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CarrierGatewaySpec   `json:"spec"`
	Status            CarrierGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CarrierGatewayList contains a list of CarrierGateways
type CarrierGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CarrierGateway `json:"items"`
}

// Repository type metadata.
var (
	CarrierGateway_Kind             = "CarrierGateway"
	CarrierGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CarrierGateway_Kind}.String()
	CarrierGateway_KindAPIVersion   = CarrierGateway_Kind + "." + CRDGroupVersion.String()
	CarrierGateway_GroupVersionKind = CRDGroupVersion.WithKind(CarrierGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&CarrierGateway{}, &CarrierGatewayList{})
}
