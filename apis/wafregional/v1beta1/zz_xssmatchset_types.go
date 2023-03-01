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

type XSSMatchSetObservation struct {

	// The ID of the Regional WAF XSS Match Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type XSSMatchSetParameters struct {

	// The name of the set
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	// +kubebuilder:validation:Optional
	XSSMatchTuple []XSSMatchTupleParameters `json:"xssMatchTuple,omitempty" tf:"xss_match_tuple,omitempty"`
}

type XSSMatchTupleFieldToMatchObservation struct {
}

type XSSMatchTupleFieldToMatchParameters struct {

	// When the value of type is HEADER, enter the name of the header that you want the WAF to search, for example, User-Agent or Referer. If the value of type is any other value, omit data.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified stringE.g., HEADER or METHOD
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type XSSMatchTupleObservation struct {
}

type XSSMatchTupleParameters struct {

	// Specifies where in a web request to look for cross-site scripting attacks.
	// +kubebuilder:validation:Required
	FieldToMatch []XSSMatchTupleFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// Which text transformation, if any, to perform on the web request before inspecting the request for cross-site scripting attacks.
	// +kubebuilder:validation:Required
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

// XSSMatchSetSpec defines the desired state of XSSMatchSet
type XSSMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     XSSMatchSetParameters `json:"forProvider"`
}

// XSSMatchSetStatus defines the observed state of XSSMatchSet.
type XSSMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        XSSMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// XSSMatchSet is the Schema for the XSSMatchSets API. Provides an AWS WAF Regional XSS Match Set resource for use with ALB.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type XSSMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   XSSMatchSetSpec   `json:"spec"`
	Status XSSMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// XSSMatchSetList contains a list of XSSMatchSets
type XSSMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []XSSMatchSet `json:"items"`
}

// Repository type metadata.
var (
	XSSMatchSet_Kind             = "XSSMatchSet"
	XSSMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: XSSMatchSet_Kind}.String()
	XSSMatchSet_KindAPIVersion   = XSSMatchSet_Kind + "." + CRDGroupVersion.String()
	XSSMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(XSSMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&XSSMatchSet{}, &XSSMatchSetList{})
}
