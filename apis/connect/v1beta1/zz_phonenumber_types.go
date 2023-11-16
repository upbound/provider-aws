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

type PhoneNumberInitParameters struct {

	// The ISO country code. For a list of Valid values, refer to PhoneNumberCountryCode.
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The description of the phone number.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain + as part of the country code. Do not specify this argument when importing the resource.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of phone number. Valid Values: TOLL_FREE | DID.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PhoneNumberObservation struct {

	// The ARN of the phone number.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ISO country code. For a list of Valid values, refer to PhoneNumberCountryCode.
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The description of the phone number.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the phone number.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The phone number. Phone numbers are formatted [+] [country code] [subscriber number including area code].
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`

	// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain + as part of the country code. Do not specify this argument when importing the resource.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// A block that specifies status of the phone number. Documented below.
	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The Amazon Resource Name (ARN) for Amazon Connect instances that phone numbers are claimed to.
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`

	// The type of phone number. Valid Values: TOLL_FREE | DID.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PhoneNumberParameters struct {

	// The ISO country code. For a list of Valid values, refer to PhoneNumberCountryCode.
	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The description of the phone number.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain + as part of the country code. Do not specify this argument when importing the resource.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Amazon Resource Name (ARN) for Amazon Connect instances that phone numbers are claimed to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`

	// Reference to a Instance in connect to populate targetArn.
	// +kubebuilder:validation:Optional
	TargetArnRef *v1.Reference `json:"targetArnRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate targetArn.
	// +kubebuilder:validation:Optional
	TargetArnSelector *v1.Selector `json:"targetArnSelector,omitempty" tf:"-"`

	// The type of phone number. Valid Values: TOLL_FREE | DID.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StatusInitParameters struct {
}

type StatusObservation struct {

	// The status message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The status of the phone number. Valid Values: CLAIMED | IN_PROGRESS | FAILED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type StatusParameters struct {
}

// PhoneNumberSpec defines the desired state of PhoneNumber
type PhoneNumberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PhoneNumberParameters `json:"forProvider"`
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
	InitProvider PhoneNumberInitParameters `json:"initProvider,omitempty"`
}

// PhoneNumberStatus defines the observed state of PhoneNumber.
type PhoneNumberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PhoneNumberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PhoneNumber is the Schema for the PhoneNumbers API. Provides details about a specific Amazon Connect Phone Number.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PhoneNumber struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.countryCode) || (has(self.initProvider) && has(self.initProvider.countryCode))",message="spec.forProvider.countryCode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   PhoneNumberSpec   `json:"spec"`
	Status PhoneNumberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PhoneNumberList contains a list of PhoneNumbers
type PhoneNumberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PhoneNumber `json:"items"`
}

// Repository type metadata.
var (
	PhoneNumber_Kind             = "PhoneNumber"
	PhoneNumber_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PhoneNumber_Kind}.String()
	PhoneNumber_KindAPIVersion   = PhoneNumber_Kind + "." + CRDGroupVersion.String()
	PhoneNumber_GroupVersionKind = CRDGroupVersion.WithKind(PhoneNumber_Kind)
)

func init() {
	SchemeBuilder.Register(&PhoneNumber{}, &PhoneNumberList{})
}
