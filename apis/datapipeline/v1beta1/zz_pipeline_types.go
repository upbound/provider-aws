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

type PipelineObservation struct {

	// The description of Pipeline.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the client certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of Pipeline.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PipelineParameters struct {

	// The description of Pipeline.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of Pipeline.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PipelineSpec defines the desired state of Pipeline
type PipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PipelineParameters `json:"forProvider"`
}

// PipelineStatus defines the observed state of Pipeline.
type PipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pipeline is the Schema for the Pipelines API. Provides a AWS DataPipeline Pipeline.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   PipelineSpec   `json:"spec"`
	Status PipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PipelineList contains a list of Pipelines
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}

// Repository type metadata.
var (
	Pipeline_Kind             = "Pipeline"
	Pipeline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pipeline_Kind}.String()
	Pipeline_KindAPIVersion   = Pipeline_Kind + "." + CRDGroupVersion.String()
	Pipeline_GroupVersionKind = CRDGroupVersion.WithKind(Pipeline_Kind)
)

func init() {
	SchemeBuilder.Register(&Pipeline{}, &PipelineList{})
}
