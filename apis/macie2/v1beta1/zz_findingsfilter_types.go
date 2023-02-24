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

type CriterionObservation struct {

	// The value for the property matches (equals) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	Eq []*string `json:"eq,omitempty" tf:"eq,omitempty"`

	// The value for the property exclusively matches (equals an exact match for) all the specified values. If you specify multiple values, Amazon Macie uses AND logic to join the values.
	EqExactMatch []*string `json:"eqExactMatch,omitempty" tf:"eq_exact_match,omitempty"`

	// The name of the field to be evaluated.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// The value for the property is greater than the specified value.
	Gt *string `json:"gt,omitempty" tf:"gt,omitempty"`

	// The value for the property is greater than or equal to the specified value.
	Gte *string `json:"gte,omitempty" tf:"gte,omitempty"`

	// The value for the property is less than the specified value.
	Lt *string `json:"lt,omitempty" tf:"lt,omitempty"`

	// The value for the property is less than or equal to the specified value.
	Lte *string `json:"lte,omitempty" tf:"lte,omitempty"`

	// The value for the property doesn't match (doesn't equal) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	Neq []*string `json:"neq,omitempty" tf:"neq,omitempty"`
}

type CriterionParameters struct {

	// The value for the property matches (equals) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	// +kubebuilder:validation:Optional
	Eq []*string `json:"eq,omitempty" tf:"eq,omitempty"`

	// The value for the property exclusively matches (equals an exact match for) all the specified values. If you specify multiple values, Amazon Macie uses AND logic to join the values.
	// +kubebuilder:validation:Optional
	EqExactMatch []*string `json:"eqExactMatch,omitempty" tf:"eq_exact_match,omitempty"`

	// The name of the field to be evaluated.
	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// The value for the property is greater than the specified value.
	// +kubebuilder:validation:Optional
	Gt *string `json:"gt,omitempty" tf:"gt,omitempty"`

	// The value for the property is greater than or equal to the specified value.
	// +kubebuilder:validation:Optional
	Gte *string `json:"gte,omitempty" tf:"gte,omitempty"`

	// The value for the property is less than the specified value.
	// +kubebuilder:validation:Optional
	Lt *string `json:"lt,omitempty" tf:"lt,omitempty"`

	// The value for the property is less than or equal to the specified value.
	// +kubebuilder:validation:Optional
	Lte *string `json:"lte,omitempty" tf:"lte,omitempty"`

	// The value for the property doesn't match (doesn't equal) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	// +kubebuilder:validation:Optional
	Neq []*string `json:"neq,omitempty" tf:"neq,omitempty"`
}

type FindingCriteriaObservation struct {

	// A condition that specifies the property, operator, and one or more values to use to filter the results.  (documented below)
	Criterion []CriterionObservation `json:"criterion,omitempty" tf:"criterion,omitempty"`
}

type FindingCriteriaParameters struct {

	// A condition that specifies the property, operator, and one or more values to use to filter the results.  (documented below)
	// +kubebuilder:validation:Optional
	Criterion []CriterionParameters `json:"criterion,omitempty" tf:"criterion,omitempty"`
}

type FindingsFilterObservation struct {

	// The action to perform on findings that meet the filter criteria (finding_criteria). Valid values are: ARCHIVE, suppress (automatically archive) the findings; and, NOOP, don't perform any action on the findings.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The Amazon Resource Name (ARN) of the Findings Filter.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A custom description of the filter. The description can contain as many as 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The criteria to use to filter findings.
	FindingCriteria []FindingCriteriaObservation `json:"findingCriteria,omitempty" tf:"finding_criteria,omitempty"`

	// The unique identifier (ID) of the macie Findings Filter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A custom name for the filter. The name must contain at least 3 characters and can contain as many as 64 characters. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The position of the filter in the list of saved filters on the Amazon Macie console. This value also determines the order in which the filter is applied to findings, relative to other filters that are also applied to the findings.
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FindingsFilterParameters struct {

	// The action to perform on findings that meet the filter criteria (finding_criteria). Valid values are: ARCHIVE, suppress (automatically archive) the findings; and, NOOP, don't perform any action on the findings.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// A custom description of the filter. The description can contain as many as 512 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The criteria to use to filter findings.
	// +kubebuilder:validation:Required
	FindingCriteria []FindingCriteriaParameters `json:"findingCriteria" tf:"finding_criteria,omitempty"`

	// A custom name for the filter. The name must contain at least 3 characters and can contain as many as 64 characters. Conflicts with name_prefix.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The position of the filter in the list of saved filters on the Amazon Macie console. This value also determines the order in which the filter is applied to findings, relative to other filters that are also applied to the findings.
	// +kubebuilder:validation:Optional
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// FindingsFilterSpec defines the desired state of FindingsFilter
type FindingsFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FindingsFilterParameters `json:"forProvider"`
}

// FindingsFilterStatus defines the observed state of FindingsFilter.
type FindingsFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FindingsFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FindingsFilter is the Schema for the FindingsFilters API. Provides a resource to manage an Amazon Macie Findings Filter.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FindingsFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FindingsFilterSpec   `json:"spec"`
	Status            FindingsFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FindingsFilterList contains a list of FindingsFilters
type FindingsFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FindingsFilter `json:"items"`
}

// Repository type metadata.
var (
	FindingsFilter_Kind             = "FindingsFilter"
	FindingsFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FindingsFilter_Kind}.String()
	FindingsFilter_KindAPIVersion   = FindingsFilter_Kind + "." + CRDGroupVersion.String()
	FindingsFilter_GroupVersionKind = CRDGroupVersion.WithKind(FindingsFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&FindingsFilter{}, &FindingsFilterList{})
}
