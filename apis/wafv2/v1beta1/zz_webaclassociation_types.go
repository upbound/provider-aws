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

type WebACLAssociationInitParameters struct {
}

type WebACLAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer, an Amazon API Gateway stage, or an Amazon Cognito User Pool.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebACLArn *string `json:"webAclArn,omitempty" tf:"web_acl_arn,omitempty"`
}

type WebACLAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer, an Amazon API Gateway stage, or an Amazon Cognito User Pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Stage
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Stage in apigateway to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Stage in apigateway to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`

	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/wafv2/v1beta1.WebACL
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	WebACLArn *string `json:"webAclArn,omitempty" tf:"web_acl_arn,omitempty"`

	// Reference to a WebACL in wafv2 to populate webAclArn.
	// +kubebuilder:validation:Optional
	WebACLArnRef *v1.Reference `json:"webAclArnRef,omitempty" tf:"-"`

	// Selector for a WebACL in wafv2 to populate webAclArn.
	// +kubebuilder:validation:Optional
	WebACLArnSelector *v1.Selector `json:"webAclArnSelector,omitempty" tf:"-"`
}

// WebACLAssociationSpec defines the desired state of WebACLAssociation
type WebACLAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebACLAssociationParameters `json:"forProvider"`
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
	InitProvider WebACLAssociationInitParameters `json:"initProvider,omitempty"`
}

// WebACLAssociationStatus defines the observed state of WebACLAssociation.
type WebACLAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebACLAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebACLAssociation is the Schema for the WebACLAssociations API. Creates a WAFv2 Web ACL Association.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WebACLAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebACLAssociationSpec   `json:"spec"`
	Status            WebACLAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebACLAssociationList contains a list of WebACLAssociations
type WebACLAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebACLAssociation `json:"items"`
}

// Repository type metadata.
var (
	WebACLAssociation_Kind             = "WebACLAssociation"
	WebACLAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebACLAssociation_Kind}.String()
	WebACLAssociation_KindAPIVersion   = WebACLAssociation_Kind + "." + CRDGroupVersion.String()
	WebACLAssociation_GroupVersionKind = CRDGroupVersion.WithKind(WebACLAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&WebACLAssociation{}, &WebACLAssociationList{})
}
