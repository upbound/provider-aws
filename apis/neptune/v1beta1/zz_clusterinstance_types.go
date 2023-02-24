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

type ClusterInstanceObservation struct {

	// The hostname of the instance. See also endpoint and port.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Specifies whether any instance modifications
	// are applied immediately, or during the next maintenance window. Default isfalse.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Amazon Resource Name (ARN) of neptune instance
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is true.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// The EC2 Availability Zone that the neptune instance is created in.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The identifier of the aws_neptune_cluster in which to launch this instance.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// The region-unique, immutable identifier for the neptune instance.
	DbiResourceID *string `json:"dbiResourceId,omitempty" tf:"dbi_resource_id,omitempty"`

	// The connection endpoint in address:port format.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The name of the database engine to be used for the neptune instance. Defaults to neptune. Valid Values: neptune.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The neptune engine version.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The Instance identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance class to use.
	InstanceClass *string `json:"instanceClass,omitempty" tf:"instance_class,omitempty"`

	// The ARN for the KMS encryption key if one is set to the neptune cluster.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// The name of the neptune parameter group to associate with this instance.
	NeptuneParameterGroupName *string `json:"neptuneParameterGroupName,omitempty" tf:"neptune_parameter_group_name,omitempty"`

	// A subnet group to associate with this neptune instance. NOTE: This must match the neptune_subnet_group_name of the attached aws_neptune_cluster.
	NeptuneSubnetGroupName *string `json:"neptuneSubnetGroupName,omitempty" tf:"neptune_subnet_group_name,omitempty"`

	// The port on which the DB accepts connections. Defaults to 8182.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00"
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier *float64 `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`

	// Bool to control if instance is publicly accessible. Default is false.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Specifies whether the neptune cluster is encrypted.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// – Boolean indicating if this instance is writable. False indicates this instance is a read replica.
	Writer *bool `json:"writer,omitempty" tf:"writer,omitempty"`
}

type ClusterInstanceParameters struct {

	// Specifies whether any instance modifications
	// are applied immediately, or during the next maintenance window. Default isfalse.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is true.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// The EC2 Availability Zone that the neptune instance is created in.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The identifier of the aws_neptune_cluster in which to launch this instance.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// The name of the database engine to be used for the neptune instance. Defaults to neptune. Valid Values: neptune.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The neptune engine version.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The instance class to use.
	// +kubebuilder:validation:Required
	InstanceClass *string `json:"instanceClass" tf:"instance_class,omitempty"`

	// The name of the neptune parameter group to associate with this instance.
	// +crossplane:generate:reference:type=ParameterGroup
	// +kubebuilder:validation:Optional
	NeptuneParameterGroupName *string `json:"neptuneParameterGroupName,omitempty" tf:"neptune_parameter_group_name,omitempty"`

	// Reference to a ParameterGroup to populate neptuneParameterGroupName.
	// +kubebuilder:validation:Optional
	NeptuneParameterGroupNameRef *v1.Reference `json:"neptuneParameterGroupNameRef,omitempty" tf:"-"`

	// Selector for a ParameterGroup to populate neptuneParameterGroupName.
	// +kubebuilder:validation:Optional
	NeptuneParameterGroupNameSelector *v1.Selector `json:"neptuneParameterGroupNameSelector,omitempty" tf:"-"`

	// A subnet group to associate with this neptune instance. NOTE: This must match the neptune_subnet_group_name of the attached aws_neptune_cluster.
	// +crossplane:generate:reference:type=SubnetGroup
	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupName *string `json:"neptuneSubnetGroupName,omitempty" tf:"neptune_subnet_group_name,omitempty"`

	// Reference to a SubnetGroup to populate neptuneSubnetGroupName.
	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupNameRef *v1.Reference `json:"neptuneSubnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup to populate neptuneSubnetGroupName.
	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupNameSelector *v1.Selector `json:"neptuneSubnetGroupNameSelector,omitempty" tf:"-"`

	// The port on which the DB accepts connections. Defaults to 8182.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00"
	// +kubebuilder:validation:Optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	// +kubebuilder:validation:Optional
	PromotionTier *float64 `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`

	// Bool to control if instance is publicly accessible. Default is false.
	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ClusterInstanceSpec defines the desired state of ClusterInstance
type ClusterInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterInstanceParameters `json:"forProvider"`
}

// ClusterInstanceStatus defines the observed state of ClusterInstance.
type ClusterInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInstance is the Schema for the ClusterInstances API. Provides an Neptune Cluster Resource Instance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterInstanceSpec   `json:"spec"`
	Status            ClusterInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInstanceList contains a list of ClusterInstances
type ClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterInstance `json:"items"`
}

// Repository type metadata.
var (
	ClusterInstance_Kind             = "ClusterInstance"
	ClusterInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterInstance_Kind}.String()
	ClusterInstance_KindAPIVersion   = ClusterInstance_Kind + "." + CRDGroupVersion.String()
	ClusterInstance_GroupVersionKind = CRDGroupVersion.WithKind(ClusterInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterInstance{}, &ClusterInstanceList{})
}
