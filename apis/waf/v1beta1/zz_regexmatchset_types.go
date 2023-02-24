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

type RegexMatchSetObservation struct {

	// Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the WAF Regex Match Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description of the Regex Match Set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuple []RegexMatchTupleObservation `json:"regexMatchTuple,omitempty" tf:"regex_match_tuple,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type RegexMatchSetParameters struct {

	// The name or description of the Regex Match Set.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	// +kubebuilder:validation:Optional
	RegexMatchTuple []RegexMatchTupleParameters `json:"regexMatchTuple,omitempty" tf:"regex_match_tuple,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type RegexMatchTupleFieldToMatchObservation struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RegexMatchTupleFieldToMatchParameters struct {

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

type RegexMatchTupleObservation struct {

	// The part of a web request that you want to search, such as a specified header or a query string.
	FieldToMatch []RegexMatchTupleFieldToMatchObservation `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// The ID of a Regex Pattern Set.
	RegexPatternSetID *string `json:"regexPatternSetId,omitempty" tf:"regex_pattern_set_id,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

type RegexMatchTupleParameters struct {

	// The part of a web request that you want to search, such as a specified header or a query string.
	// +kubebuilder:validation:Required
	FieldToMatch []RegexMatchTupleFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// The ID of a Regex Pattern Set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/waf/v1beta1.RegexPatternSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RegexPatternSetID *string `json:"regexPatternSetId,omitempty" tf:"regex_pattern_set_id,omitempty"`

	// Reference to a RegexPatternSet in waf to populate regexPatternSetId.
	// +kubebuilder:validation:Optional
	RegexPatternSetIDRef *v1.Reference `json:"regexPatternSetIdRef,omitempty" tf:"-"`

	// Selector for a RegexPatternSet in waf to populate regexPatternSetId.
	// +kubebuilder:validation:Optional
	RegexPatternSetIDSelector *v1.Selector `json:"regexPatternSetIdSelector,omitempty" tf:"-"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Required
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

// RegexMatchSetSpec defines the desired state of RegexMatchSet
type RegexMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegexMatchSetParameters `json:"forProvider"`
}

// RegexMatchSetStatus defines the observed state of RegexMatchSet.
type RegexMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegexMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegexMatchSet is the Schema for the RegexMatchSets API. Provides a AWS WAF Regex Match Set resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RegexMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegexMatchSetSpec   `json:"spec"`
	Status            RegexMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegexMatchSetList contains a list of RegexMatchSets
type RegexMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegexMatchSet `json:"items"`
}

// Repository type metadata.
var (
	RegexMatchSet_Kind             = "RegexMatchSet"
	RegexMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegexMatchSet_Kind}.String()
	RegexMatchSet_KindAPIVersion   = RegexMatchSet_Kind + "." + CRDGroupVersion.String()
	RegexMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(RegexMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&RegexMatchSet{}, &RegexMatchSetList{})
}
