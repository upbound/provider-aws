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

type ParameterGroupInitParameters struct {

	// The description of the ElastiCache parameter group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The family of the ElastiCache parameter group.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The name of the ElastiCache parameter group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of ElastiCache parameters to apply.
	Parameter []ParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ParameterGroupObservation struct {

	// The AWS ARN associated with the parameter group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description of the ElastiCache parameter group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The family of the ElastiCache parameter group.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The ElastiCache parameter group name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the ElastiCache parameter group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of ElastiCache parameters to apply.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ParameterGroupParameters struct {

	// The description of the ElastiCache parameter group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The family of the ElastiCache parameter group.
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The name of the ElastiCache parameter group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of ElastiCache parameters to apply.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ParameterInitParameters struct {

	// The name of the ElastiCache parameter group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the ElastiCache parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterObservation struct {

	// The name of the ElastiCache parameter group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the ElastiCache parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterParameters struct {

	// The name of the ElastiCache parameter group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of the ElastiCache parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// ParameterGroupSpec defines the desired state of ParameterGroup
type ParameterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ParameterGroupParameters `json:"forProvider"`
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
	InitProvider ParameterGroupInitParameters `json:"initProvider,omitempty"`
}

// ParameterGroupStatus defines the observed state of ParameterGroup.
type ParameterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterGroup is the Schema for the ParameterGroups API. Provides an ElastiCache parameter group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.family) || (has(self.initProvider) && has(self.initProvider.family))",message="spec.forProvider.family is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ParameterGroupSpec   `json:"spec"`
	Status ParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterGroupList contains a list of ParameterGroups
type ParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	ParameterGroup_Kind             = "ParameterGroup"
	ParameterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ParameterGroup_Kind}.String()
	ParameterGroup_KindAPIVersion   = ParameterGroup_Kind + "." + CRDGroupVersion.String()
	ParameterGroup_GroupVersionKind = CRDGroupVersion.WithKind(ParameterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ParameterGroup{}, &ParameterGroupList{})
}
