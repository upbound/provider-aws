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

type SegmentInitParameters struct {

	// Specifies the description of the segment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The pattern to use for the segment. For more information about pattern syntax, see Segment rule pattern syntax.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SegmentObservation struct {

	// The ARN of the segment.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date and time that the segment is created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// Specifies the description of the segment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The number of experiments that this segment is used in. This count includes all current experiments, not just those that are currently running.
	ExperimentCount *float64 `json:"experimentCount,omitempty" tf:"experiment_count,omitempty"`

	// The ID has the same value as the ARN of the segment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time that this segment was most recently updated.
	LastUpdatedTime *string `json:"lastUpdatedTime,omitempty" tf:"last_updated_time,omitempty"`

	// The number of launches that this segment is used in. This count includes all current launches, not just those that are currently running.
	LaunchCount *float64 `json:"launchCount,omitempty" tf:"launch_count,omitempty"`

	// The pattern to use for the segment. For more information about pattern syntax, see Segment rule pattern syntax.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SegmentParameters struct {

	// Specifies the description of the segment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The pattern to use for the segment. For more information about pattern syntax, see Segment rule pattern syntax.
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SegmentSpec defines the desired state of Segment
type SegmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SegmentParameters `json:"forProvider"`
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
	InitProvider SegmentInitParameters `json:"initProvider,omitempty"`
}

// SegmentStatus defines the observed state of Segment.
type SegmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SegmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Segment is the Schema for the Segments API. Provides a CloudWatch Evidently Segment resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Segment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pattern) || (has(self.initProvider) && has(self.initProvider.pattern))",message="spec.forProvider.pattern is a required parameter"
	Spec   SegmentSpec   `json:"spec"`
	Status SegmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SegmentList contains a list of Segments
type SegmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Segment `json:"items"`
}

// Repository type metadata.
var (
	Segment_Kind             = "Segment"
	Segment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Segment_Kind}.String()
	Segment_KindAPIVersion   = Segment_Kind + "." + CRDGroupVersion.String()
	Segment_GroupVersionKind = CRDGroupVersion.WithKind(Segment_Kind)
)

func init() {
	SchemeBuilder.Register(&Segment{}, &SegmentList{})
}
