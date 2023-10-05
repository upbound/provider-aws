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

type GlobalClusterInitParameters struct {

	// Name for an automatically created database on cluster creation.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to true. The default is false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Name of the database engine to be used for this DB cluster. Current Valid values: docdb. Defaults to docdb. Conflicts with source_db_cluster_identifier.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The global cluster identifier.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`

	// Specifies whether the DB cluster is encrypted. The default is false unless source_db_cluster_identifier is specified and encrypted.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
}

type GlobalClusterMembersInitParameters struct {
}

type GlobalClusterMembersObservation struct {

	// Amazon Resource Name (ARN) of member DB Cluster.
	DBClusterArn *string `json:"dbClusterArn,omitempty" tf:"db_cluster_arn,omitempty"`

	// Whether the member is the primary DB Cluster.
	IsWriter *bool `json:"isWriter,omitempty" tf:"is_writer,omitempty"`
}

type GlobalClusterMembersParameters struct {
}

type GlobalClusterObservation struct {

	// Global Cluster Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Name for an automatically created database on cluster creation.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to true. The default is false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Name of the database engine to be used for this DB cluster. Current Valid values: docdb. Defaults to docdb. Conflicts with source_db_cluster_identifier.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The global cluster identifier.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`

	// Set of objects containing Global Cluster members.
	GlobalClusterMembers []GlobalClusterMembersObservation `json:"globalClusterMembers,omitempty" tf:"global_cluster_members,omitempty"`

	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed.
	GlobalClusterResourceID *string `json:"globalClusterResourceId,omitempty" tf:"global_cluster_resource_id,omitempty"`

	// DocumentDB Global Cluster ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation.
	SourceDBClusterIdentifier *string `json:"sourceDbClusterIdentifier,omitempty" tf:"source_db_cluster_identifier,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies whether the DB cluster is encrypted. The default is false unless source_db_cluster_identifier is specified and encrypted.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
}

type GlobalClusterParameters struct {

	// Name for an automatically created database on cluster creation.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to true. The default is false.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Name of the database engine to be used for this DB cluster. Current Valid values: docdb. Defaults to docdb. Conflicts with source_db_cluster_identifier.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The global cluster identifier.
	// +kubebuilder:validation:Optional
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/docdb/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SourceDBClusterIdentifier *string `json:"sourceDbClusterIdentifier,omitempty" tf:"source_db_cluster_identifier,omitempty"`

	// Reference to a Cluster in docdb to populate sourceDbClusterIdentifier.
	// +kubebuilder:validation:Optional
	SourceDBClusterIdentifierRef *v1.Reference `json:"sourceDbClusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in docdb to populate sourceDbClusterIdentifier.
	// +kubebuilder:validation:Optional
	SourceDBClusterIdentifierSelector *v1.Selector `json:"sourceDbClusterIdentifierSelector,omitempty" tf:"-"`

	// Specifies whether the DB cluster is encrypted. The default is false unless source_db_cluster_identifier is specified and encrypted.
	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
}

// GlobalClusterSpec defines the desired state of GlobalCluster
type GlobalClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalClusterParameters `json:"forProvider"`
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
	InitProvider GlobalClusterInitParameters `json:"initProvider,omitempty"`
}

// GlobalClusterStatus defines the observed state of GlobalCluster.
type GlobalClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalCluster is the Schema for the GlobalClusters API. Manages a DocumentDB Global Cluster
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GlobalCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.globalClusterIdentifier) || (has(self.initProvider) && has(self.initProvider.globalClusterIdentifier))",message="spec.forProvider.globalClusterIdentifier is a required parameter"
	Spec   GlobalClusterSpec   `json:"spec"`
	Status GlobalClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalClusterList contains a list of GlobalClusters
type GlobalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalCluster `json:"items"`
}

// Repository type metadata.
var (
	GlobalCluster_Kind             = "GlobalCluster"
	GlobalCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalCluster_Kind}.String()
	GlobalCluster_KindAPIVersion   = GlobalCluster_Kind + "." + CRDGroupVersion.String()
	GlobalCluster_GroupVersionKind = CRDGroupVersion.WithKind(GlobalCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalCluster{}, &GlobalClusterList{})
}
