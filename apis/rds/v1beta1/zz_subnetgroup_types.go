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

type SubnetGroupObservation struct {

	// The ARN of the db subnet group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description of the DB subnet group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The db subnet group name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A list of VPC subnet IDs.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The network type of the db subnet group.
	SupportedNetworkTypes []*string `json:"supportedNetworkTypes,omitempty" tf:"supported_network_types,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SubnetGroupParameters struct {

	// The description of the DB subnet group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of VPC subnet IDs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SubnetGroupSpec defines the desired state of SubnetGroup
type SubnetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetGroupParameters `json:"forProvider"`
}

// SubnetGroupStatus defines the observed state of SubnetGroup.
type SubnetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetGroup is the Schema for the SubnetGroups API. Provides an RDS DB subnet group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetGroupSpec   `json:"spec"`
	Status            SubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetGroupList contains a list of SubnetGroups
type SubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetGroup `json:"items"`
}

// Repository type metadata.
var (
	SubnetGroup_Kind             = "SubnetGroup"
	SubnetGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetGroup_Kind}.String()
	SubnetGroup_KindAPIVersion   = SubnetGroup_Kind + "." + CRDGroupVersion.String()
	SubnetGroup_GroupVersionKind = CRDGroupVersion.WithKind(SubnetGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetGroup{}, &SubnetGroupList{})
}
