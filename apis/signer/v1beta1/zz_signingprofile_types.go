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

type SignatureValidityPeriodInitParameters struct {
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type SignatureValidityPeriodObservation struct {
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type SignatureValidityPeriodParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type SigningProfileInitParameters struct {

	// The ID of the platform that is used by the target signing profile.
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// The validity period for a signing job.
	SignatureValidityPeriod []SignatureValidityPeriodInitParameters `json:"signatureValidityPeriod,omitempty" tf:"signature_validity_period,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SigningProfileObservation struct {

	// The Amazon Resource Name (ARN) for the signing profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A human-readable name for the signing platform associated with the signing profile.
	PlatformDisplayName *string `json:"platformDisplayName,omitempty" tf:"platform_display_name,omitempty"`

	// The ID of the platform that is used by the target signing profile.
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// Revocation information for a signing profile.
	RevocationRecord []SigningProfileRevocationRecordObservation `json:"revocationRecord,omitempty" tf:"revocation_record,omitempty"`

	// The validity period for a signing job.
	SignatureValidityPeriod []SignatureValidityPeriodObservation `json:"signatureValidityPeriod,omitempty" tf:"signature_validity_period,omitempty"`

	// The status of the target signing profile.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The current version of the signing profile.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The signing profile ARN, including the profile version.
	VersionArn *string `json:"versionArn,omitempty" tf:"version_arn,omitempty"`
}

type SigningProfileParameters struct {

	// The ID of the platform that is used by the target signing profile.
	// +kubebuilder:validation:Optional
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The validity period for a signing job.
	// +kubebuilder:validation:Optional
	SignatureValidityPeriod []SignatureValidityPeriodParameters `json:"signatureValidityPeriod,omitempty" tf:"signature_validity_period,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SigningProfileRevocationRecordInitParameters struct {
}

type SigningProfileRevocationRecordObservation struct {
	RevocationEffectiveFrom *string `json:"revocationEffectiveFrom,omitempty" tf:"revocation_effective_from,omitempty"`

	RevokedAt *string `json:"revokedAt,omitempty" tf:"revoked_at,omitempty"`

	RevokedBy *string `json:"revokedBy,omitempty" tf:"revoked_by,omitempty"`
}

type SigningProfileRevocationRecordParameters struct {
}

// SigningProfileSpec defines the desired state of SigningProfile
type SigningProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SigningProfileParameters `json:"forProvider"`
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
	InitProvider SigningProfileInitParameters `json:"initProvider,omitempty"`
}

// SigningProfileStatus defines the observed state of SigningProfile.
type SigningProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SigningProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SigningProfile is the Schema for the SigningProfiles API. Creates a Signer Signing Profile.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SigningProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.platformId) || (has(self.initProvider) && has(self.initProvider.platformId))",message="spec.forProvider.platformId is a required parameter"
	Spec   SigningProfileSpec   `json:"spec"`
	Status SigningProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SigningProfileList contains a list of SigningProfiles
type SigningProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SigningProfile `json:"items"`
}

// Repository type metadata.
var (
	SigningProfile_Kind             = "SigningProfile"
	SigningProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SigningProfile_Kind}.String()
	SigningProfile_KindAPIVersion   = SigningProfile_Kind + "." + CRDGroupVersion.String()
	SigningProfile_GroupVersionKind = CRDGroupVersion.WithKind(SigningProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&SigningProfile{}, &SigningProfileList{})
}
