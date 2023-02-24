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

type ConditionObservation struct {
	StringEquals []StringEqualsObservation `json:"stringEquals,omitempty" tf:"string_equals,omitempty"`

	StringLike []StringLikeObservation `json:"stringLike,omitempty" tf:"string_like,omitempty"`

	StringNotEquals []StringNotEqualsObservation `json:"stringNotEquals,omitempty" tf:"string_not_equals,omitempty"`

	StringNotLike []StringNotLikeObservation `json:"stringNotLike,omitempty" tf:"string_not_like,omitempty"`
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	StringEquals []StringEqualsParameters `json:"stringEquals,omitempty" tf:"string_equals,omitempty"`

	// +kubebuilder:validation:Optional
	StringLike []StringLikeParameters `json:"stringLike,omitempty" tf:"string_like,omitempty"`

	// +kubebuilder:validation:Optional
	StringNotEquals []StringNotEqualsParameters `json:"stringNotEquals,omitempty" tf:"string_not_equals,omitempty"`

	// +kubebuilder:validation:Optional
	StringNotLike []StringNotLikeParameters `json:"stringNotLike,omitempty" tf:"string_not_like,omitempty"`
}

type SelectionObservation struct {

	// A list of conditions that you define to assign resources to your backup plans using tags.
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the AWS Backup Developer Guide for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Backup Selection identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The display name of a resource selection document.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to exclude from a backup plan.
	NotResources []*string `json:"notResources,omitempty" tf:"not_resources,omitempty"`

	// The backup plan ID to be associated with the selection of resources.
	PlanID *string `json:"planId,omitempty" tf:"plan_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan.
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`

	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTag []SelectionTagObservation `json:"selectionTag,omitempty" tf:"selection_tag,omitempty"`
}

type SelectionParameters struct {

	// A list of conditions that you define to assign resources to your backup plans using tags.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the AWS Backup Developer Guide for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`

	// The display name of a resource selection document.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to exclude from a backup plan.
	// +kubebuilder:validation:Optional
	NotResources []*string `json:"notResources,omitempty" tf:"not_resources,omitempty"`

	// The backup plan ID to be associated with the selection of resources.
	// +crossplane:generate:reference:type=Plan
	// +kubebuilder:validation:Optional
	PlanID *string `json:"planId,omitempty" tf:"plan_id,omitempty"`

	// Reference to a Plan to populate planId.
	// +kubebuilder:validation:Optional
	PlanIDRef *v1.Reference `json:"planIdRef,omitempty" tf:"-"`

	// Selector for a Plan to populate planId.
	// +kubebuilder:validation:Optional
	PlanIDSelector *v1.Selector `json:"planIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan.
	// +kubebuilder:validation:Optional
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`

	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	// +kubebuilder:validation:Optional
	SelectionTag []SelectionTagParameters `json:"selectionTag,omitempty" tf:"selection_tag,omitempty"`
}

type SelectionTagObservation struct {

	// The key in a key-value pair.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// An operation, such as StringEquals, that is applied to a key-value pair used to filter resources in a selection.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value in a key-value pair.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SelectionTagParameters struct {

	// The key in a key-value pair.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// An operation, such as StringEquals, that is applied to a key-value pair used to filter resources in a selection.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The value in a key-value pair.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type StringEqualsObservation struct {

	// The key in a key-value pair.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value in a key-value pair.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StringEqualsParameters struct {

	// The key in a key-value pair.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The value in a key-value pair.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type StringLikeObservation struct {

	// The key in a key-value pair.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value in a key-value pair.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StringLikeParameters struct {

	// The key in a key-value pair.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The value in a key-value pair.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type StringNotEqualsObservation struct {

	// The key in a key-value pair.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value in a key-value pair.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StringNotEqualsParameters struct {

	// The key in a key-value pair.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The value in a key-value pair.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type StringNotLikeObservation struct {

	// The key in a key-value pair.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value in a key-value pair.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StringNotLikeParameters struct {

	// The key in a key-value pair.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The value in a key-value pair.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// SelectionSpec defines the desired state of Selection
type SelectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SelectionParameters `json:"forProvider"`
}

// SelectionStatus defines the observed state of Selection.
type SelectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SelectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Selection is the Schema for the Selections API. Manages selection conditions for AWS Backup plan resources.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Selection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SelectionSpec   `json:"spec"`
	Status            SelectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SelectionList contains a list of Selections
type SelectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Selection `json:"items"`
}

// Repository type metadata.
var (
	Selection_Kind             = "Selection"
	Selection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Selection_Kind}.String()
	Selection_KindAPIVersion   = Selection_Kind + "." + CRDGroupVersion.String()
	Selection_GroupVersionKind = CRDGroupVersion.WithKind(Selection_Kind)
)

func init() {
	SchemeBuilder.Register(&Selection{}, &SelectionList{})
}
