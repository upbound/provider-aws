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

type ThingInitParameters struct {

	// Map of attributes of the thing.
	// +mapType:granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The thing type name.
	ThingTypeName *string `json:"thingTypeName,omitempty" tf:"thing_type_name,omitempty"`
}

type ThingObservation struct {

	// The ARN of the thing.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Map of attributes of the thing.
	// +mapType:granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The default client ID.
	DefaultClientID *string `json:"defaultClientId,omitempty" tf:"default_client_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The thing type name.
	ThingTypeName *string `json:"thingTypeName,omitempty" tf:"thing_type_name,omitempty"`

	// The current version of the thing record in the registry.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ThingParameters struct {

	// Map of attributes of the thing.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The thing type name.
	// +kubebuilder:validation:Optional
	ThingTypeName *string `json:"thingTypeName,omitempty" tf:"thing_type_name,omitempty"`
}

// ThingSpec defines the desired state of Thing
type ThingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ThingParameters `json:"forProvider"`
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
	InitProvider ThingInitParameters `json:"initProvider,omitempty"`
}

// ThingStatus defines the observed state of Thing.
type ThingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ThingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Thing is the Schema for the Things API. Creates and manages an AWS IoT Thing.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Thing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ThingSpec   `json:"spec"`
	Status            ThingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThingList contains a list of Things
type ThingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Thing `json:"items"`
}

// Repository type metadata.
var (
	Thing_Kind             = "Thing"
	Thing_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Thing_Kind}.String()
	Thing_KindAPIVersion   = Thing_Kind + "." + CRDGroupVersion.String()
	Thing_GroupVersionKind = CRDGroupVersion.WithKind(Thing_Kind)
)

func init() {
	SchemeBuilder.Register(&Thing{}, &ThingList{})
}
