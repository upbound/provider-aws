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

type SizeConstraintSetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the WAF Size Constraint Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SizeConstraintSetParameters struct {

	// The name or description of the Size Constraint Set.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the parts of web requests that you want to inspect the size of.
	// +kubebuilder:validation:Optional
	SizeConstraints []SizeConstraintsParameters `json:"sizeConstraints,omitempty" tf:"size_constraints,omitempty"`
}

type SizeConstraintsFieldToMatchObservation struct {
}

type SizeConstraintsFieldToMatchParameters struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SizeConstraintsObservation struct {
}

type SizeConstraintsParameters struct {

	// The type of comparison you want to perform.
	// e.g., EQ, NE, LT, GT.
	// See docs for all supported values.
	// +kubebuilder:validation:Required
	ComparisonOperator *string `json:"comparisonOperator" tf:"comparison_operator,omitempty"`

	// Specifies where in a web request to look for the size constraint.
	// +kubebuilder:validation:Required
	FieldToMatch []SizeConstraintsFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// The size in bytes that you want to compare against the size of the specified field_to_match.
	// Valid values are between 0 - 21474836480 bytes (0 - 20 GB).
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on field_to_match before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	// Note: if you choose BODY as type, you must choose NONE because CloudFront forwards only the first 8192 bytes for inspection.
	// +kubebuilder:validation:Required
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

// SizeConstraintSetSpec defines the desired state of SizeConstraintSet
type SizeConstraintSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SizeConstraintSetParameters `json:"forProvider"`
}

// SizeConstraintSetStatus defines the observed state of SizeConstraintSet.
type SizeConstraintSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SizeConstraintSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SizeConstraintSet is the Schema for the SizeConstraintSets API. Provides an AWS WAF Regional Size Constraint Set resource for use with ALB.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   SizeConstraintSetSpec   `json:"spec"`
	Status SizeConstraintSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SizeConstraintSetList contains a list of SizeConstraintSets
type SizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SizeConstraintSet `json:"items"`
}

// Repository type metadata.
var (
	SizeConstraintSet_Kind             = "SizeConstraintSet"
	SizeConstraintSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SizeConstraintSet_Kind}.String()
	SizeConstraintSet_KindAPIVersion   = SizeConstraintSet_Kind + "." + CRDGroupVersion.String()
	SizeConstraintSet_GroupVersionKind = CRDGroupVersion.WithKind(SizeConstraintSet_Kind)
)

func init() {
	SchemeBuilder.Register(&SizeConstraintSet{}, &SizeConstraintSetList{})
}
