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

type AvailabilityZoneGroupObservation struct {

	// Name of the Availability Zone Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AvailabilityZoneGroupParameters struct {

	// Indicates whether to enable or disable Availability Zone Group. Valid values: opted-in or not-opted-in.
	// +kubebuilder:validation:Optional
	OptInStatus *string `json:"optInStatus,omitempty" tf:"opt_in_status,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AvailabilityZoneGroupSpec defines the desired state of AvailabilityZoneGroup
type AvailabilityZoneGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AvailabilityZoneGroupParameters `json:"forProvider"`
}

// AvailabilityZoneGroupStatus defines the observed state of AvailabilityZoneGroup.
type AvailabilityZoneGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AvailabilityZoneGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AvailabilityZoneGroup is the Schema for the AvailabilityZoneGroups API. Manages an EC2 Availability Zone Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AvailabilityZoneGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.optInStatus)",message="optInStatus is a required parameter"
	Spec   AvailabilityZoneGroupSpec   `json:"spec"`
	Status AvailabilityZoneGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AvailabilityZoneGroupList contains a list of AvailabilityZoneGroups
type AvailabilityZoneGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AvailabilityZoneGroup `json:"items"`
}

// Repository type metadata.
var (
	AvailabilityZoneGroup_Kind             = "AvailabilityZoneGroup"
	AvailabilityZoneGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AvailabilityZoneGroup_Kind}.String()
	AvailabilityZoneGroup_KindAPIVersion   = AvailabilityZoneGroup_Kind + "." + CRDGroupVersion.String()
	AvailabilityZoneGroup_GroupVersionKind = CRDGroupVersion.WithKind(AvailabilityZoneGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AvailabilityZoneGroup{}, &AvailabilityZoneGroupList{})
}
