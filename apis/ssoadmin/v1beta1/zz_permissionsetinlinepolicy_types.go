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

type PermissionSetInlinePolicyObservation struct {

	// The Amazon Resource Names (ARNs) of the Permission Set and SSO Instance, separated by a comma (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PermissionSetInlinePolicyParameters struct {

	// The IAM inline policy to attach to a Permission Set.
	// +kubebuilder:validation:Optional
	InlinePolicy *string `json:"inlinePolicy,omitempty" tf:"inline_policy,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssoadmin/v1beta1.PermissionSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("instance_arn",false)
	// +kubebuilder:validation:Optional
	InstanceArn *string `json:"instanceArn,omitempty" tf:"instance_arn,omitempty"`

	// Reference to a PermissionSet in ssoadmin to populate instanceArn.
	// +kubebuilder:validation:Optional
	InstanceArnRef *v1.Reference `json:"instanceArnRef,omitempty" tf:"-"`

	// Selector for a PermissionSet in ssoadmin to populate instanceArn.
	// +kubebuilder:validation:Optional
	InstanceArnSelector *v1.Selector `json:"instanceArnSelector,omitempty" tf:"-"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssoadmin/v1beta1.PermissionSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	PermissionSetArn *string `json:"permissionSetArn,omitempty" tf:"permission_set_arn,omitempty"`

	// Reference to a PermissionSet in ssoadmin to populate permissionSetArn.
	// +kubebuilder:validation:Optional
	PermissionSetArnRef *v1.Reference `json:"permissionSetArnRef,omitempty" tf:"-"`

	// Selector for a PermissionSet in ssoadmin to populate permissionSetArn.
	// +kubebuilder:validation:Optional
	PermissionSetArnSelector *v1.Selector `json:"permissionSetArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// PermissionSetInlinePolicySpec defines the desired state of PermissionSetInlinePolicy
type PermissionSetInlinePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionSetInlinePolicyParameters `json:"forProvider"`
}

// PermissionSetInlinePolicyStatus defines the observed state of PermissionSetInlinePolicy.
type PermissionSetInlinePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionSetInlinePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionSetInlinePolicy is the Schema for the PermissionSetInlinePolicys API. Manages an IAM inline policy for a Single Sign-On (SSO) Permission Set
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PermissionSetInlinePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.inlinePolicy)",message="inlinePolicy is a required parameter"
	Spec   PermissionSetInlinePolicySpec   `json:"spec"`
	Status PermissionSetInlinePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionSetInlinePolicyList contains a list of PermissionSetInlinePolicys
type PermissionSetInlinePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PermissionSetInlinePolicy `json:"items"`
}

// Repository type metadata.
var (
	PermissionSetInlinePolicy_Kind             = "PermissionSetInlinePolicy"
	PermissionSetInlinePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PermissionSetInlinePolicy_Kind}.String()
	PermissionSetInlinePolicy_KindAPIVersion   = PermissionSetInlinePolicy_Kind + "." + CRDGroupVersion.String()
	PermissionSetInlinePolicy_GroupVersionKind = CRDGroupVersion.WithKind(PermissionSetInlinePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&PermissionSetInlinePolicy{}, &PermissionSetInlinePolicyList{})
}
