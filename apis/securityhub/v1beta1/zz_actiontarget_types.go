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

type ActionTargetInitParameters struct {

	// The name of the custom action target.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The description for the custom action target.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ActionTargetObservation struct {

	// Amazon Resource Name (ARN) of the Security Hub custom action target.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the custom action target.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The description for the custom action target.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ActionTargetParameters struct {

	// The name of the custom action target.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The description for the custom action target.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ActionTargetSpec defines the desired state of ActionTarget
type ActionTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActionTargetParameters `json:"forProvider"`
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
	InitProvider ActionTargetInitParameters `json:"initProvider,omitempty"`
}

// ActionTargetStatus defines the observed state of ActionTarget.
type ActionTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActionTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ActionTarget is the Schema for the ActionTargets API. Creates Security Hub custom action.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ActionTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ActionTargetSpec   `json:"spec"`
	Status ActionTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActionTargetList contains a list of ActionTargets
type ActionTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionTarget `json:"items"`
}

// Repository type metadata.
var (
	ActionTarget_Kind             = "ActionTarget"
	ActionTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActionTarget_Kind}.String()
	ActionTarget_KindAPIVersion   = ActionTarget_Kind + "." + CRDGroupVersion.String()
	ActionTarget_GroupVersionKind = CRDGroupVersion.WithKind(ActionTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&ActionTarget{}, &ActionTargetList{})
}
