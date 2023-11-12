// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type IPGroupInitParameters struct {

	// The description of the IP group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the IP group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	Rules []RulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IPGroupObservation struct {

	// The description of the IP group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP group identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the IP group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type IPGroupParameters struct {

	// The description of the IP group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the IP group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RulesInitParameters struct {

	// The description of the IP group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP address range, in CIDR notation, e.g., 10.0.0.0/16
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type RulesObservation struct {

	// The description of the IP group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP address range, in CIDR notation, e.g., 10.0.0.0/16
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type RulesParameters struct {

	// The description of the IP group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP address range, in CIDR notation, e.g., 10.0.0.0/16
	// +kubebuilder:validation:Optional
	Source *string `json:"source" tf:"source,omitempty"`
}

// IPGroupSpec defines the desired state of IPGroup
type IPGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPGroupParameters `json:"forProvider"`
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
	InitProvider IPGroupInitParameters `json:"initProvider,omitempty"`
}

// IPGroupStatus defines the observed state of IPGroup.
type IPGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPGroup is the Schema for the IPGroups API. Provides an IP access control group in AWS WorkSpaces Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IPGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IPGroupSpec   `json:"spec"`
	Status IPGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPGroupList contains a list of IPGroups
type IPGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPGroup `json:"items"`
}

// Repository type metadata.
var (
	IPGroup_Kind             = "IPGroup"
	IPGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPGroup_Kind}.String()
	IPGroup_KindAPIVersion   = IPGroup_Kind + "." + CRDGroupVersion.String()
	IPGroup_GroupVersionKind = CRDGroupVersion.WithKind(IPGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&IPGroup{}, &IPGroupList{})
}
