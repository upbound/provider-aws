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

type ByteMatchSetObservation struct {

	// The ID of the WAF Byte Match Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ByteMatchSetParameters struct {

	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	// +kubebuilder:validation:Optional
	ByteMatchTuples []ByteMatchTuplesParameters `json:"byteMatchTuples,omitempty" tf:"byte_match_tuples,omitempty"`

	// The name or description of the Byte Match Set.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type ByteMatchTuplesObservation struct {
}

type ByteMatchTuplesParameters struct {

	// The part of a web request that you want to search, such as a specified header or a query string.
	// +kubebuilder:validation:Required
	FieldToMatch []FieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// Within the portion of a web request that you want to search
	// (for example, in the query string, if any), specify where you want to search.
	// e.g., CONTAINS, CONTAINS_WORD or EXACTLY.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Required
	PositionalConstraint *string `json:"positionalConstraint" tf:"positional_constraint,omitempty"`

	// The value that you want to search for within the field specified by field_to_match, e.g., badrefer1.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Optional
	TargetString *string `json:"targetString,omitempty" tf:"target_string,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on target_string before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Required
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

type FieldToMatchObservation struct {
}

type FieldToMatchParameters struct {

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

// ByteMatchSetSpec defines the desired state of ByteMatchSet
type ByteMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ByteMatchSetParameters `json:"forProvider"`
}

// ByteMatchSetStatus defines the observed state of ByteMatchSet.
type ByteMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ByteMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ByteMatchSet is the Schema for the ByteMatchSets API. Provides a AWS WAF Byte Match Set resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ByteMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   ByteMatchSetSpec   `json:"spec"`
	Status ByteMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ByteMatchSetList contains a list of ByteMatchSets
type ByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ByteMatchSet `json:"items"`
}

// Repository type metadata.
var (
	ByteMatchSet_Kind             = "ByteMatchSet"
	ByteMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ByteMatchSet_Kind}.String()
	ByteMatchSet_KindAPIVersion   = ByteMatchSet_Kind + "." + CRDGroupVersion.String()
	ByteMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(ByteMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ByteMatchSet{}, &ByteMatchSetList{})
}
