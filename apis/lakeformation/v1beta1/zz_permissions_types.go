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

type DataLocationObservation struct {
}

type DataLocationParameters struct {

	// –  Amazon Resource Name (ARN) that uniquely identifies the data location resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lakeformation/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Reference to a Resource in lakeformation to populate arn.
	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// Selector for a Resource in lakeformation to populate arn.
	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// Identifier for the Data Catalog where the location is registered with Lake Formation. By default, it is the account ID of the caller.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`
}

type DatabaseObservation struct {
}

type DatabaseParameters struct {

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// –  Name of the database resource. Unique to the Data Catalog.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.CatalogDatabase
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a CatalogDatabase in glue to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a CatalogDatabase in glue to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

type ExpressionObservation struct {
}

type ExpressionParameters struct {

	// name of an LF-Tag.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// A list of possible values of an LF-Tag.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type LfTagObservation struct {
}

type LfTagParameters struct {

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// name for the tag.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// A list of possible values an attribute can take.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type LfTagPolicyObservation struct {
}

type LfTagPolicyParameters struct {

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// A list of tag conditions that apply to the resource's tag policy. Configuration block for tag conditions that apply to the policy. See expression below.
	// +kubebuilder:validation:Required
	Expression []ExpressionParameters `json:"expression" tf:"expression,omitempty"`

	// –  The resource type for which the tag policy applies. Valid values are DATABASE and TABLE.
	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`
}

type PermissionsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PermissionsParameters struct {

	// –  Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Whether the permissions are to be granted for the Data Catalog. Defaults to false.
	// +kubebuilder:validation:Optional
	CatalogResource *bool `json:"catalogResource,omitempty" tf:"catalog_resource,omitempty"`

	// Configuration block for a data location resource. Detailed below.
	// +kubebuilder:validation:Optional
	DataLocation []DataLocationParameters `json:"dataLocation,omitempty" tf:"data_location,omitempty"`

	// Configuration block for a database resource. Detailed below.
	// +kubebuilder:validation:Optional
	Database []DatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// Configuration block for an LF-tag resource. Detailed below.
	// +kubebuilder:validation:Optional
	LfTag []LfTagParameters `json:"lfTag,omitempty" tf:"lf_tag,omitempty"`

	// Configuration block for an LF-tag policy resource. Detailed below.
	// +kubebuilder:validation:Optional
	LfTagPolicy []LfTagPolicyParameters `json:"lfTagPolicy,omitempty" tf:"lf_tag_policy,omitempty"`

	// –  List of permissions granted to the principal. Valid values may include ALL, ALTER, ASSOCIATE, CREATE_DATABASE, CREATE_TABLE, DATA_LOCATION_ACCESS, DELETE, DESCRIBE, DROP, INSERT, and SELECT. For details on each permission, see Lake Formation Permissions Reference.
	// +kubebuilder:validation:Required
	Permissions []*string `json:"permissions" tf:"permissions,omitempty"`

	// Subset of permissions which the principal can pass.
	// +kubebuilder:validation:Optional
	PermissionsWithGrantOption []*string `json:"permissionsWithGrantOption,omitempty" tf:"permissions_with_grant_option,omitempty"`

	// account permissions. For more information, see Lake Formation Permissions Reference.
	// +kubebuilder:validation:Required
	Principal *string `json:"principal" tf:"principal,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for a table resource. Detailed below.
	// +kubebuilder:validation:Optional
	Table []TableParameters `json:"table,omitempty" tf:"table,omitempty"`

	// Configuration block for a table with columns resource. Detailed below.
	// +kubebuilder:validation:Optional
	TableWithColumns []TableWithColumnsParameters `json:"tableWithColumns,omitempty" tf:"table_with_columns,omitempty"`
}

type TableObservation struct {
}

type TableParameters struct {

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// –  Name of the database for the table. Unique to a Data Catalog.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// Name of the table.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to use a wildcard representing every table under a database. Defaults to false.
	// +kubebuilder:validation:Optional
	Wildcard *bool `json:"wildcard,omitempty" tf:"wildcard,omitempty"`
}

type TableWithColumnsObservation struct {
}

type TableWithColumnsParameters struct {

	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Set of column names for the table.
	// +kubebuilder:validation:Optional
	ColumnNames []*string `json:"columnNames,omitempty" tf:"column_names,omitempty"`

	// –  Name of the database for the table with columns resource. Unique to the Data Catalog.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// Set of column names for the table to exclude.
	// +kubebuilder:validation:Optional
	ExcludedColumnNames []*string `json:"excludedColumnNames,omitempty" tf:"excluded_column_names,omitempty"`

	// –  Name of the table resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.CatalogTable
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a CatalogTable in glue to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a CatalogTable in glue to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// Whether to use a column wildcard.
	// +kubebuilder:validation:Optional
	Wildcard *bool `json:"wildcard,omitempty" tf:"wildcard,omitempty"`
}

// PermissionsSpec defines the desired state of Permissions
type PermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionsParameters `json:"forProvider"`
}

// PermissionsStatus defines the observed state of Permissions.
type PermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Permissions is the Schema for the Permissionss API. Grants permissions to the principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Permissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionsSpec   `json:"spec"`
	Status            PermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionsList contains a list of Permissionss
type PermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Permissions `json:"items"`
}

// Repository type metadata.
var (
	Permissions_Kind             = "Permissions"
	Permissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Permissions_Kind}.String()
	Permissions_KindAPIVersion   = Permissions_Kind + "." + CRDGroupVersion.String()
	Permissions_GroupVersionKind = CRDGroupVersion.WithKind(Permissions_Kind)
)

func init() {
	SchemeBuilder.Register(&Permissions{}, &PermissionsList{})
}
