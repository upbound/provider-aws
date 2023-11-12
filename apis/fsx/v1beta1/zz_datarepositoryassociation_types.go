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

type AutoExportPolicyInitParameters struct {

	// A list of file event types to automatically export to your linked S3 bucket or import from the linked S3 bucket. Valid values are NEW, CHANGED, DELETED. Max of 3.
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`
}

type AutoExportPolicyObservation struct {

	// A list of file event types to automatically export to your linked S3 bucket or import from the linked S3 bucket. Valid values are NEW, CHANGED, DELETED. Max of 3.
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`
}

type AutoExportPolicyParameters struct {

	// A list of file event types to automatically export to your linked S3 bucket or import from the linked S3 bucket. Valid values are NEW, CHANGED, DELETED. Max of 3.
	// +kubebuilder:validation:Optional
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`
}

type AutoImportPolicyInitParameters struct {

	// A list of file event types to automatically export to your linked S3 bucket or import from the linked S3 bucket. Valid values are NEW, CHANGED, DELETED. Max of 3.
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`
}

type AutoImportPolicyObservation struct {

	// A list of file event types to automatically export to your linked S3 bucket or import from the linked S3 bucket. Valid values are NEW, CHANGED, DELETED. Max of 3.
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`
}

type AutoImportPolicyParameters struct {

	// A list of file event types to automatically export to your linked S3 bucket or import from the linked S3 bucket. Valid values are NEW, CHANGED, DELETED. Max of 3.
	// +kubebuilder:validation:Optional
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`
}

type DataRepositoryAssociationInitParameters struct {

	// Set to true to run an import data repository task to import metadata from the data repository to the file system after the data repository association is created. Defaults to false.
	BatchImportMetaDataOnCreate *bool `json:"batchImportMetaDataOnCreate,omitempty" tf:"batch_import_meta_data_on_create,omitempty"`

	// The path to the Amazon S3 data repository that will be linked to the file system. The path must be an S3 bucket s3://myBucket/myPrefix/. This path specifies where in the S3 data repository files will be imported from or exported to. The same S3 bucket cannot be linked more than once to the same file system.
	DataRepositoryPath *string `json:"dataRepositoryPath,omitempty" tf:"data_repository_path,omitempty"`

	// Set to true to delete files from the file system upon deleting this data repository association. Defaults to false.
	DeleteDataInFilesystem *bool `json:"deleteDataInFilesystem,omitempty" tf:"delete_data_in_filesystem,omitempty"`

	// A path on the file system that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with data_repository_path. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path /ns1/, then you cannot link another data repository with file system path /ns1/ns2. This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
	FileSystemPath *string `json:"fileSystemPath,omitempty" tf:"file_system_path,omitempty"`

	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize,omitempty" tf:"imported_file_chunk_size,omitempty"`

	// See the s3 configuration block. Max of 1.
	// The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
	S3 []S3InitParameters `json:"s3,omitempty" tf:"s3,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DataRepositoryAssociationObservation struct {

	// Amazon Resource Name of the file system.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Identifier of the data repository association, e.g., dra-12345678
	AssociationID *string `json:"associationId,omitempty" tf:"association_id,omitempty"`

	// Set to true to run an import data repository task to import metadata from the data repository to the file system after the data repository association is created. Defaults to false.
	BatchImportMetaDataOnCreate *bool `json:"batchImportMetaDataOnCreate,omitempty" tf:"batch_import_meta_data_on_create,omitempty"`

	// The path to the Amazon S3 data repository that will be linked to the file system. The path must be an S3 bucket s3://myBucket/myPrefix/. This path specifies where in the S3 data repository files will be imported from or exported to. The same S3 bucket cannot be linked more than once to the same file system.
	DataRepositoryPath *string `json:"dataRepositoryPath,omitempty" tf:"data_repository_path,omitempty"`

	// Set to true to delete files from the file system upon deleting this data repository association. Defaults to false.
	DeleteDataInFilesystem *bool `json:"deleteDataInFilesystem,omitempty" tf:"delete_data_in_filesystem,omitempty"`

	// The ID of the Amazon FSx file system to on which to create a data repository association.
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// A path on the file system that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with data_repository_path. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path /ns1/, then you cannot link another data repository with file system path /ns1/ns2. This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
	FileSystemPath *string `json:"fileSystemPath,omitempty" tf:"file_system_path,omitempty"`

	// Identifier of the data repository association, e.g., dra-12345678
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize,omitempty" tf:"imported_file_chunk_size,omitempty"`

	// See the s3 configuration block. Max of 1.
	// The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
	S3 []S3Observation `json:"s3,omitempty" tf:"s3,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DataRepositoryAssociationParameters struct {

	// Set to true to run an import data repository task to import metadata from the data repository to the file system after the data repository association is created. Defaults to false.
	// +kubebuilder:validation:Optional
	BatchImportMetaDataOnCreate *bool `json:"batchImportMetaDataOnCreate,omitempty" tf:"batch_import_meta_data_on_create,omitempty"`

	// The path to the Amazon S3 data repository that will be linked to the file system. The path must be an S3 bucket s3://myBucket/myPrefix/. This path specifies where in the S3 data repository files will be imported from or exported to. The same S3 bucket cannot be linked more than once to the same file system.
	// +kubebuilder:validation:Optional
	DataRepositoryPath *string `json:"dataRepositoryPath,omitempty" tf:"data_repository_path,omitempty"`

	// Set to true to delete files from the file system upon deleting this data repository association. Defaults to false.
	// +kubebuilder:validation:Optional
	DeleteDataInFilesystem *bool `json:"deleteDataInFilesystem,omitempty" tf:"delete_data_in_filesystem,omitempty"`

	// The ID of the Amazon FSx file system to on which to create a data repository association.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/fsx/v1beta1.LustreFileSystem
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Reference to a LustreFileSystem in fsx to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDRef *v1.Reference `json:"fileSystemIdRef,omitempty" tf:"-"`

	// Selector for a LustreFileSystem in fsx to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDSelector *v1.Selector `json:"fileSystemIdSelector,omitempty" tf:"-"`

	// A path on the file system that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with data_repository_path. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path /ns1/, then you cannot link another data repository with file system path /ns1/ns2. This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
	// +kubebuilder:validation:Optional
	FileSystemPath *string `json:"fileSystemPath,omitempty" tf:"file_system_path,omitempty"`

	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
	// +kubebuilder:validation:Optional
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize,omitempty" tf:"imported_file_chunk_size,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// See the s3 configuration block. Max of 1.
	// The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
	// +kubebuilder:validation:Optional
	S3 []S3Parameters `json:"s3,omitempty" tf:"s3,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type S3InitParameters struct {

	// Specifies the type of updated objects that will be automatically exported from your file system to the linked S3 bucket. See the events configuration block.
	AutoExportPolicy []AutoExportPolicyInitParameters `json:"autoExportPolicy,omitempty" tf:"auto_export_policy,omitempty"`

	// Specifies the type of updated objects that will be automatically imported from the linked S3 bucket to your file system. See the events configuration block.
	AutoImportPolicy []AutoImportPolicyInitParameters `json:"autoImportPolicy,omitempty" tf:"auto_import_policy,omitempty"`
}

type S3Observation struct {

	// Specifies the type of updated objects that will be automatically exported from your file system to the linked S3 bucket. See the events configuration block.
	AutoExportPolicy []AutoExportPolicyObservation `json:"autoExportPolicy,omitempty" tf:"auto_export_policy,omitempty"`

	// Specifies the type of updated objects that will be automatically imported from the linked S3 bucket to your file system. See the events configuration block.
	AutoImportPolicy []AutoImportPolicyObservation `json:"autoImportPolicy,omitempty" tf:"auto_import_policy,omitempty"`
}

type S3Parameters struct {

	// Specifies the type of updated objects that will be automatically exported from your file system to the linked S3 bucket. See the events configuration block.
	// +kubebuilder:validation:Optional
	AutoExportPolicy []AutoExportPolicyParameters `json:"autoExportPolicy,omitempty" tf:"auto_export_policy,omitempty"`

	// Specifies the type of updated objects that will be automatically imported from the linked S3 bucket to your file system. See the events configuration block.
	// +kubebuilder:validation:Optional
	AutoImportPolicy []AutoImportPolicyParameters `json:"autoImportPolicy,omitempty" tf:"auto_import_policy,omitempty"`
}

// DataRepositoryAssociationSpec defines the desired state of DataRepositoryAssociation
type DataRepositoryAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataRepositoryAssociationParameters `json:"forProvider"`
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
	InitProvider DataRepositoryAssociationInitParameters `json:"initProvider,omitempty"`
}

// DataRepositoryAssociationStatus defines the observed state of DataRepositoryAssociation.
type DataRepositoryAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataRepositoryAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataRepositoryAssociation is the Schema for the DataRepositoryAssociations API. Manages a FSx for Lustre Data Repository Association.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DataRepositoryAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataRepositoryPath) || (has(self.initProvider) && has(self.initProvider.dataRepositoryPath))",message="spec.forProvider.dataRepositoryPath is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fileSystemPath) || (has(self.initProvider) && has(self.initProvider.fileSystemPath))",message="spec.forProvider.fileSystemPath is a required parameter"
	Spec   DataRepositoryAssociationSpec   `json:"spec"`
	Status DataRepositoryAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataRepositoryAssociationList contains a list of DataRepositoryAssociations
type DataRepositoryAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataRepositoryAssociation `json:"items"`
}

// Repository type metadata.
var (
	DataRepositoryAssociation_Kind             = "DataRepositoryAssociation"
	DataRepositoryAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataRepositoryAssociation_Kind}.String()
	DataRepositoryAssociation_KindAPIVersion   = DataRepositoryAssociation_Kind + "." + CRDGroupVersion.String()
	DataRepositoryAssociation_GroupVersionKind = CRDGroupVersion.WithKind(DataRepositoryAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&DataRepositoryAssociation{}, &DataRepositoryAssociationList{})
}
