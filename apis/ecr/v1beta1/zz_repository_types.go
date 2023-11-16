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

type EncryptionConfigurationInitParameters struct {

	// The encryption type to use for the repository. Valid values are AES256 or KMS. Defaults to AES256.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`
}

type EncryptionConfigurationObservation struct {

	// The encryption type to use for the repository. Valid values are AES256 or KMS. Defaults to AES256.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// The ARN of the KMS key to use when encryption_type is KMS. If not specified, uses the default AWS managed key for ECR.
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`
}

type EncryptionConfigurationParameters struct {

	// The encryption type to use for the repository. Valid values are AES256 or KMS. Defaults to AES256.
	// +kubebuilder:validation:Optional
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// The ARN of the KMS key to use when encryption_type is KMS. If not specified, uses the default AWS managed key for ECR.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// Reference to a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`
}

type ImageScanningConfigurationInitParameters struct {

	// Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
	ScanOnPush *bool `json:"scanOnPush,omitempty" tf:"scan_on_push,omitempty"`
}

type ImageScanningConfigurationObservation struct {

	// Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
	ScanOnPush *bool `json:"scanOnPush,omitempty" tf:"scan_on_push,omitempty"`
}

type ImageScanningConfigurationParameters struct {

	// Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
	// +kubebuilder:validation:Optional
	ScanOnPush *bool `json:"scanOnPush" tf:"scan_on_push,omitempty"`
}

type RepositoryInitParameters struct {

	// Encryption configuration for the repository. See below for schema.
	EncryptionConfiguration []EncryptionConfigurationInitParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// If true, will delete the repository even if it contains images.
	// Defaults to false.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the ECR User Guide for more information about image scanning.
	ImageScanningConfiguration []ImageScanningConfigurationInitParameters `json:"imageScanningConfiguration,omitempty" tf:"image_scanning_configuration,omitempty"`

	// The tag mutability setting for the repository. Must be one of: MUTABLE or IMMUTABLE. Defaults to MUTABLE.
	ImageTagMutability *string `json:"imageTagMutability,omitempty" tf:"image_tag_mutability,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RepositoryObservation struct {

	// Full ARN of the repository.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Encryption configuration for the repository. See below for schema.
	EncryptionConfiguration []EncryptionConfigurationObservation `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// If true, will delete the repository even if it contains images.
	// Defaults to false.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the ECR User Guide for more information about image scanning.
	ImageScanningConfiguration []ImageScanningConfigurationObservation `json:"imageScanningConfiguration,omitempty" tf:"image_scanning_configuration,omitempty"`

	// The tag mutability setting for the repository. Must be one of: MUTABLE or IMMUTABLE. Defaults to MUTABLE.
	ImageTagMutability *string `json:"imageTagMutability,omitempty" tf:"image_tag_mutability,omitempty"`

	// The registry ID where the repository was created.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// The URL of the repository (in the form aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName).
	RepositoryURL *string `json:"repositoryUrl,omitempty" tf:"repository_url,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RepositoryParameters struct {

	// Encryption configuration for the repository. See below for schema.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// If true, will delete the repository even if it contains images.
	// Defaults to false.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the ECR User Guide for more information about image scanning.
	// +kubebuilder:validation:Optional
	ImageScanningConfiguration []ImageScanningConfigurationParameters `json:"imageScanningConfiguration,omitempty" tf:"image_scanning_configuration,omitempty"`

	// The tag mutability setting for the repository. Must be one of: MUTABLE or IMMUTABLE. Defaults to MUTABLE.
	// +kubebuilder:validation:Optional
	ImageTagMutability *string `json:"imageTagMutability,omitempty" tf:"image_tag_mutability,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
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
	InitProvider RepositoryInitParameters `json:"initProvider,omitempty"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Repository is the Schema for the Repositorys API. Provides an Elastic Container Registry Repository.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
