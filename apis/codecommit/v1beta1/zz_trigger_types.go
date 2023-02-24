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

type TriggerObservation struct {

	// System-generated unique identifier.
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	Trigger []TriggerTriggerObservation `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type TriggerParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name for the repository. This needs to be less than 100 characters.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codecommit/v1beta1.Repository
	// +kubebuilder:validation:Optional
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Reference to a Repository in codecommit to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameRef *v1.Reference `json:"repositoryNameRef,omitempty" tf:"-"`

	// Selector for a Repository in codecommit to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameSelector *v1.Selector `json:"repositoryNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Trigger []TriggerTriggerParameters `json:"trigger" tf:"trigger,omitempty"`
}

type TriggerTriggerObservation struct {

	// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
	Branches []*string `json:"branches,omitempty" tf:"branches,omitempty"`

	// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
	CustomData *string `json:"customData,omitempty" tf:"custom_data,omitempty"`

	// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: all, updateReference, createReference, deleteReference.
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// The name of the trigger.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TriggerTriggerParameters struct {

	// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
	// +kubebuilder:validation:Optional
	Branches []*string `json:"branches,omitempty" tf:"branches,omitempty"`

	// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
	// +kubebuilder:validation:Optional
	CustomData *string `json:"customData,omitempty" tf:"custom_data,omitempty"`

	// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// Reference to a Topic in sns to populate destinationArn.
	// +kubebuilder:validation:Optional
	DestinationArnRef *v1.Reference `json:"destinationArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate destinationArn.
	// +kubebuilder:validation:Optional
	DestinationArnSelector *v1.Selector `json:"destinationArnSelector,omitempty" tf:"-"`

	// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: all, updateReference, createReference, deleteReference.
	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// The name of the trigger.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// TriggerSpec defines the desired state of Trigger
type TriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerParameters `json:"forProvider"`
}

// TriggerStatus defines the observed state of Trigger.
type TriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Trigger is the Schema for the Triggers API. Provides a CodeCommit Trigger Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Trigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerSpec   `json:"spec"`
	Status            TriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerList contains a list of Triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trigger `json:"items"`
}

// Repository type metadata.
var (
	Trigger_Kind             = "Trigger"
	Trigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Trigger_Kind}.String()
	Trigger_KindAPIVersion   = Trigger_Kind + "." + CRDGroupVersion.String()
	Trigger_GroupVersionKind = CRDGroupVersion.WithKind(Trigger_Kind)
)

func init() {
	SchemeBuilder.Register(&Trigger{}, &TriggerList{})
}
