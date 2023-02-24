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

type ResourceObservation struct {

	// –  Amazon Resource Name (ARN) of the resource, an S3 path.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time the resource was last modified in RFC 3339 format.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// linked role must exist and is used.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type ResourceParameters struct {

	// –  Amazon Resource Name (ARN) of the resource, an S3 path.
	// +kubebuilder:validation:Required
	Arn *string `json:"arn" tf:"arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// linked role must exist and is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

// ResourceSpec defines the desired state of Resource
type ResourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceParameters `json:"forProvider"`
}

// ResourceStatus defines the observed state of Resource.
type ResourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Resource is the Schema for the Resources API. Registers a Lake Formation resource as managed by the Data Catalog.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Resource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceSpec   `json:"spec"`
	Status            ResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceList contains a list of Resources
type ResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resource `json:"items"`
}

// Repository type metadata.
var (
	Resource_Kind             = "Resource"
	Resource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Resource_Kind}.String()
	Resource_KindAPIVersion   = Resource_Kind + "." + CRDGroupVersion.String()
	Resource_GroupVersionKind = CRDGroupVersion.WithKind(Resource_Kind)
)

func init() {
	SchemeBuilder.Register(&Resource{}, &ResourceList{})
}
