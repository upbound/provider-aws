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

type EventFilterInitParameters struct {

	// Source of the events.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type EventFilterObservation struct {

	// Source of the events.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type EventFilterParameters struct {

	// Source of the events.
	// +kubebuilder:validation:Optional
	Source *string `json:"source" tf:"source,omitempty"`
}

type EventIntegrationInitParameters struct {

	// Description of the Event Integration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
	EventFilter []EventFilterInitParameters `json:"eventFilter,omitempty" tf:"event_filter,omitempty"`

	// EventBridge bus.
	EventbridgeBus *string `json:"eventbridgeBus,omitempty" tf:"eventbridge_bus,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventIntegrationObservation struct {

	// ARN of the Event Integration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the Event Integration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
	EventFilter []EventFilterObservation `json:"eventFilter,omitempty" tf:"event_filter,omitempty"`

	// EventBridge bus.
	EventbridgeBus *string `json:"eventbridgeBus,omitempty" tf:"eventbridge_bus,omitempty"`

	// Identifier of the Event Integration which is the name of the Event Integration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EventIntegrationParameters struct {

	// Description of the Event Integration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
	// +kubebuilder:validation:Optional
	EventFilter []EventFilterParameters `json:"eventFilter,omitempty" tf:"event_filter,omitempty"`

	// EventBridge bus.
	// +kubebuilder:validation:Optional
	EventbridgeBus *string `json:"eventbridgeBus,omitempty" tf:"eventbridge_bus,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EventIntegrationSpec defines the desired state of EventIntegration
type EventIntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventIntegrationParameters `json:"forProvider"`
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
	InitProvider EventIntegrationInitParameters `json:"initProvider,omitempty"`
}

// EventIntegrationStatus defines the observed state of EventIntegration.
type EventIntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventIntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventIntegration is the Schema for the EventIntegrations API. Provides details about a specific Amazon AppIntegrations Event Integration
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EventIntegration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventFilter) || (has(self.initProvider) && has(self.initProvider.eventFilter))",message="spec.forProvider.eventFilter is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventbridgeBus) || (has(self.initProvider) && has(self.initProvider.eventbridgeBus))",message="spec.forProvider.eventbridgeBus is a required parameter"
	Spec   EventIntegrationSpec   `json:"spec"`
	Status EventIntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventIntegrationList contains a list of EventIntegrations
type EventIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventIntegration `json:"items"`
}

// Repository type metadata.
var (
	EventIntegration_Kind             = "EventIntegration"
	EventIntegration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventIntegration_Kind}.String()
	EventIntegration_KindAPIVersion   = EventIntegration_Kind + "." + CRDGroupVersion.String()
	EventIntegration_GroupVersionKind = CRDGroupVersion.WithKind(EventIntegration_Kind)
)

func init() {
	SchemeBuilder.Register(&EventIntegration{}, &EventIntegrationList{})
}
