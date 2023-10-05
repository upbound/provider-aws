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

type AccountSettingDefaultInitParameters struct {

	// Name of the account setting to set. Valid values are serviceLongArnFormat, taskLongArnFormat, containerInstanceLongArnFormat, awsvpcTrunking and containerInsights.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// State of the setting. Valid values are enabled and disabled.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AccountSettingDefaultObservation struct {

	// ARN that identifies the account setting.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the account setting to set. Valid values are serviceLongArnFormat, taskLongArnFormat, containerInstanceLongArnFormat, awsvpcTrunking and containerInsights.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`

	// State of the setting. Valid values are enabled and disabled.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AccountSettingDefaultParameters struct {

	// Name of the account setting to set. Valid values are serviceLongArnFormat, taskLongArnFormat, containerInstanceLongArnFormat, awsvpcTrunking and containerInsights.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// State of the setting. Valid values are enabled and disabled.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// AccountSettingDefaultSpec defines the desired state of AccountSettingDefault
type AccountSettingDefaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountSettingDefaultParameters `json:"forProvider"`
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
	InitProvider AccountSettingDefaultInitParameters `json:"initProvider,omitempty"`
}

// AccountSettingDefaultStatus defines the observed state of AccountSettingDefault.
type AccountSettingDefaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountSettingDefaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountSettingDefault is the Schema for the AccountSettingDefaults API. Provides an ECS Default account setting.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AccountSettingDefault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   AccountSettingDefaultSpec   `json:"spec"`
	Status AccountSettingDefaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountSettingDefaultList contains a list of AccountSettingDefaults
type AccountSettingDefaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountSettingDefault `json:"items"`
}

// Repository type metadata.
var (
	AccountSettingDefault_Kind             = "AccountSettingDefault"
	AccountSettingDefault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountSettingDefault_Kind}.String()
	AccountSettingDefault_KindAPIVersion   = AccountSettingDefault_Kind + "." + CRDGroupVersion.String()
	AccountSettingDefault_GroupVersionKind = CRDGroupVersion.WithKind(AccountSettingDefault_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountSettingDefault{}, &AccountSettingDefaultList{})
}
