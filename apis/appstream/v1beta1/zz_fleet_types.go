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

type ComputeCapacityObservation struct {

	// Number of currently available instances that can be used to stream sessions.
	Available *float64 `json:"available,omitempty" tf:"available,omitempty"`

	// Desired number of streaming instances.
	DesiredInstances *float64 `json:"desiredInstances,omitempty" tf:"desired_instances,omitempty"`

	// Number of instances in use for streaming.
	InUse *float64 `json:"inUse,omitempty" tf:"in_use,omitempty"`

	// Total number of simultaneous streaming instances that are running.
	Running *float64 `json:"running,omitempty" tf:"running,omitempty"`
}

type ComputeCapacityParameters struct {

	// Desired number of streaming instances.
	// +kubebuilder:validation:Required
	DesiredInstances *float64 `json:"desiredInstances" tf:"desired_instances,omitempty"`
}

type DomainJoinInfoObservation struct {

	// Fully qualified name of the directory (for example, corp.example.com).
	DirectoryName *string `json:"directoryName,omitempty" tf:"directory_name,omitempty"`

	// Distinguished name of the organizational unit for computer accounts.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name,omitempty"`
}

type DomainJoinInfoParameters struct {

	// Fully qualified name of the directory (for example, corp.example.com).
	// +kubebuilder:validation:Optional
	DirectoryName *string `json:"directoryName,omitempty" tf:"directory_name,omitempty"`

	// Distinguished name of the organizational unit for computer accounts.
	// +kubebuilder:validation:Optional
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name,omitempty"`
}

type FleetObservation struct {

	// ARN of the appstream fleet.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration block for the desired capacity of the fleet. See below.
	ComputeCapacity []ComputeCapacityObservation `json:"computeCapacity,omitempty" tf:"compute_capacity,omitempty"`

	// Date and time, in UTC and extended RFC 3339 format, when the fleet was created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// Description to display.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amount of time that a streaming session remains active after users disconnect.
	DisconnectTimeoutInSeconds *float64 `json:"disconnectTimeoutInSeconds,omitempty" tf:"disconnect_timeout_in_seconds,omitempty"`

	// Human-readable friendly name for the AppStream fleet.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Configuration block for the name of the directory and organizational unit (OU) to use to join the fleet to a Microsoft Active Directory domain. See below.
	DomainJoinInfo []DomainJoinInfoObservation `json:"domainJoinInfo,omitempty" tf:"domain_join_info,omitempty"`

	// Enables or disables default internet access for the fleet.
	EnableDefaultInternetAccess *bool `json:"enableDefaultInternetAccess,omitempty" tf:"enable_default_internet_access,omitempty"`

	// Fleet type. Valid values are: ON_DEMAND, ALWAYS_ON
	FleetType *string `json:"fleetType,omitempty" tf:"fleet_type,omitempty"`

	// ARN of the IAM role to apply to the fleet.
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Unique identifier (ID) of the appstream fleet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the disconnect_timeout_in_seconds time interval begins.
	IdleDisconnectTimeoutInSeconds *float64 `json:"idleDisconnectTimeoutInSeconds,omitempty" tf:"idle_disconnect_timeout_in_seconds,omitempty"`

	// ARN of the public, private, or shared image to use.
	ImageArn *string `json:"imageArn,omitempty" tf:"image_arn,omitempty"`

	// Name of the image used to create the fleet.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// Instance type to use when launching fleet instances.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Maximum amount of time that a streaming session can remain active, in seconds.
	MaxUserDurationInSeconds *float64 `json:"maxUserDurationInSeconds,omitempty" tf:"max_user_duration_in_seconds,omitempty"`

	// Unique name for the fleet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// State of the fleet. Can be STARTING, RUNNING, STOPPING or STOPPED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// AppStream 2.0 view that is displayed to your users when they stream from the fleet. When APP is specified, only the windows of applications opened by users display. When DESKTOP is specified, the standard desktop that is provided by the operating system displays. If not specified, defaults to APP.
	StreamView *string `json:"streamView,omitempty" tf:"stream_view,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Configuration block for the VPC configuration for the image builder. See below.
	VPCConfig []VPCConfigObservation `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type FleetParameters struct {

	// Configuration block for the desired capacity of the fleet. See below.
	// +kubebuilder:validation:Optional
	ComputeCapacity []ComputeCapacityParameters `json:"computeCapacity,omitempty" tf:"compute_capacity,omitempty"`

	// Description to display.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amount of time that a streaming session remains active after users disconnect.
	// +kubebuilder:validation:Optional
	DisconnectTimeoutInSeconds *float64 `json:"disconnectTimeoutInSeconds,omitempty" tf:"disconnect_timeout_in_seconds,omitempty"`

	// Human-readable friendly name for the AppStream fleet.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Configuration block for the name of the directory and organizational unit (OU) to use to join the fleet to a Microsoft Active Directory domain. See below.
	// +kubebuilder:validation:Optional
	DomainJoinInfo []DomainJoinInfoParameters `json:"domainJoinInfo,omitempty" tf:"domain_join_info,omitempty"`

	// Enables or disables default internet access for the fleet.
	// +kubebuilder:validation:Optional
	EnableDefaultInternetAccess *bool `json:"enableDefaultInternetAccess,omitempty" tf:"enable_default_internet_access,omitempty"`

	// Fleet type. Valid values are: ON_DEMAND, ALWAYS_ON
	// +kubebuilder:validation:Optional
	FleetType *string `json:"fleetType,omitempty" tf:"fleet_type,omitempty"`

	// ARN of the IAM role to apply to the fleet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`

	// Amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the disconnect_timeout_in_seconds time interval begins.
	// +kubebuilder:validation:Optional
	IdleDisconnectTimeoutInSeconds *float64 `json:"idleDisconnectTimeoutInSeconds,omitempty" tf:"idle_disconnect_timeout_in_seconds,omitempty"`

	// ARN of the public, private, or shared image to use.
	// +kubebuilder:validation:Optional
	ImageArn *string `json:"imageArn,omitempty" tf:"image_arn,omitempty"`

	// Name of the image used to create the fleet.
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// Instance type to use when launching fleet instances.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Maximum amount of time that a streaming session can remain active, in seconds.
	// +kubebuilder:validation:Optional
	MaxUserDurationInSeconds *float64 `json:"maxUserDurationInSeconds,omitempty" tf:"max_user_duration_in_seconds,omitempty"`

	// Unique name for the fleet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// AppStream 2.0 view that is displayed to your users when they stream from the fleet. When APP is specified, only the windows of applications opened by users display. When DESKTOP is specified, the standard desktop that is provided by the operating system displays. If not specified, defaults to APP.
	// +kubebuilder:validation:Optional
	StreamView *string `json:"streamView,omitempty" tf:"stream_view,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for the VPC configuration for the image builder. See below.
	// +kubebuilder:validation:Optional
	VPCConfig []VPCConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type VPCConfigObservation struct {

	// Identifiers of the security groups for the fleet or image builder.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Identifiers of the subnets to which a network interface is attached from the fleet instance or image builder instance.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type VPCConfigParameters struct {

	// Identifiers of the security groups for the fleet or image builder.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Identifiers of the subnets to which a network interface is attached from the fleet instance or image builder instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

// FleetSpec defines the desired state of Fleet
type FleetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FleetParameters `json:"forProvider"`
}

// FleetStatus defines the observed state of Fleet.
type FleetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FleetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Fleet is the Schema for the Fleets API. Provides an AppStream fleet
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Fleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.computeCapacity)",message="computeCapacity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.instanceType)",message="instanceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   FleetSpec   `json:"spec"`
	Status FleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FleetList contains a list of Fleets
type FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Fleet `json:"items"`
}

// Repository type metadata.
var (
	Fleet_Kind             = "Fleet"
	Fleet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Fleet_Kind}.String()
	Fleet_KindAPIVersion   = Fleet_Kind + "." + CRDGroupVersion.String()
	Fleet_GroupVersionKind = CRDGroupVersion.WithKind(Fleet_Kind)
)

func init() {
	SchemeBuilder.Register(&Fleet{}, &FleetList{})
}
