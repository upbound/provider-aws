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

type EBSBlockDeviceObservation struct {

	// Whether the volume should be destroyed
	// on instance termination (Default: true).
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// The name of the device to mount.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Whether the volume should be encrypted or not. Defaults to false.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The amount of provisioned
	// IOPS.
	// This must be set with a volume_type of "io1".
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Whether the device in the block device mapping of the AMI is suppressed.
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// The Snapshot ID to mount.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The throughput (MiBps) to provision for a gp3 volume.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// The size of the volume in gigabytes.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// The type of volume. Can be standard, gp2, gp3, st1, sc1 or io1.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EBSBlockDeviceParameters struct {

	// Whether the volume should be destroyed
	// on instance termination (Default: true).
	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// The name of the device to mount.
	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// Whether the volume should be encrypted or not. Defaults to false.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The amount of provisioned
	// IOPS.
	// This must be set with a volume_type of "io1".
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Whether the device in the block device mapping of the AMI is suppressed.
	// +kubebuilder:validation:Optional
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// The Snapshot ID to mount.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The throughput (MiBps) to provision for a gp3 volume.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// The size of the volume in gigabytes.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// The type of volume. Can be standard, gp2, gp3, st1, sc1 or io1.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EphemeralBlockDeviceObservation struct {

	// The name of the block device to mount on the instance.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Whether the device in the block device mapping of the AMI is suppressed.
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// The Instance Store Device Name.
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type EphemeralBlockDeviceParameters struct {

	// The name of the block device to mount on the instance.
	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// Whether the device in the block device mapping of the AMI is suppressed.
	// +kubebuilder:validation:Optional
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// The Instance Store Device Name.
	// +kubebuilder:validation:Optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type LaunchConfigurationObservation struct {

	// The Amazon Resource Name of the launch configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIPAddress *bool `json:"associatePublicIpAddress,omitempty" tf:"associate_public_ip_address,omitempty"`

	// Additional EBS block devices to attach to the instance. See Block Devices below for details.
	EBSBlockDevice []EBSBlockDeviceObservation `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// If true, the launched EC2 instance will be EBS-optimized.
	EBSOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring *bool `json:"enableMonitoring,omitempty" tf:"enable_monitoring,omitempty"`

	// Customize Ephemeral (also known as "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevice []EphemeralBlockDeviceObservation `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// The name attribute of the IAM instance profile to associate with launched instances.
	IAMInstanceProfile *string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`

	// The ID of the launch configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The EC2 image ID to launch.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The size of instance to launch.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The key name that should be used for the instance.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// The metadata options for the instance.
	MetadataOptions []MetadataOptionsObservation `json:"metadataOptions,omitempty" tf:"metadata_options,omitempty"`

	// The tenancy of the instance. Valid values are default or dedicated, see AWS's Create Launch Configuration for more details.
	PlacementTenancy *string `json:"placementTenancy,omitempty" tf:"placement_tenancy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Customize details about the root block device of the instance. See Block Devices below for details.
	RootBlockDevice []RootBlockDeviceObservation `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`

	// A list of associated security group IDS.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The maximum price to use for reserving spot instances.
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see user_data_base64 instead.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// Can be used instead of user_data to pass base64-encoded binary data directly. Use this instead of user_data whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 *string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`

	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. vpc-2730681a)
	VPCClassicLinkID *string `json:"vpcClassicLinkId,omitempty" tf:"vpc_classic_link_id,omitempty"`

	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. sg-46ae3d11).
	VPCClassicLinkSecurityGroups []*string `json:"vpcClassicLinkSecurityGroups,omitempty" tf:"vpc_classic_link_security_groups,omitempty"`
}

type LaunchConfigurationParameters struct {

	// Associate a public ip address with an instance in a VPC.
	// +kubebuilder:validation:Optional
	AssociatePublicIPAddress *bool `json:"associatePublicIpAddress,omitempty" tf:"associate_public_ip_address,omitempty"`

	// Additional EBS block devices to attach to the instance. See Block Devices below for details.
	// +kubebuilder:validation:Optional
	EBSBlockDevice []EBSBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// If true, the launched EC2 instance will be EBS-optimized.
	// +kubebuilder:validation:Optional
	EBSOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// Enables/disables detailed monitoring. This is enabled by default.
	// +kubebuilder:validation:Optional
	EnableMonitoring *bool `json:"enableMonitoring,omitempty" tf:"enable_monitoring,omitempty"`

	// Customize Ephemeral (also known as "Instance Store") volumes on the instance. See Block Devices below for details.
	// +kubebuilder:validation:Optional
	EphemeralBlockDevice []EphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// The name attribute of the IAM instance profile to associate with launched instances.
	// +kubebuilder:validation:Optional
	IAMInstanceProfile *string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`

	// The EC2 image ID to launch.
	// +kubebuilder:validation:Required
	ImageID *string `json:"imageId" tf:"image_id,omitempty"`

	// The size of instance to launch.
	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// The key name that should be used for the instance.
	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// The metadata options for the instance.
	// +kubebuilder:validation:Optional
	MetadataOptions []MetadataOptionsParameters `json:"metadataOptions,omitempty" tf:"metadata_options,omitempty"`

	// The tenancy of the instance. Valid values are default or dedicated, see AWS's Create Launch Configuration for more details.
	// +kubebuilder:validation:Optional
	PlacementTenancy *string `json:"placementTenancy,omitempty" tf:"placement_tenancy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Customize details about the root block device of the instance. See Block Devices below for details.
	// +kubebuilder:validation:Optional
	RootBlockDevice []RootBlockDeviceParameters `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`

	// A list of associated security group IDS.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The maximum price to use for reserving spot instances.
	// +kubebuilder:validation:Optional
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see user_data_base64 instead.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// Can be used instead of user_data to pass base64-encoded binary data directly. Use this instead of user_data whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	// +kubebuilder:validation:Optional
	UserDataBase64 *string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`

	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. vpc-2730681a)
	// +kubebuilder:validation:Optional
	VPCClassicLinkID *string `json:"vpcClassicLinkId,omitempty" tf:"vpc_classic_link_id,omitempty"`

	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. sg-46ae3d11).
	// +kubebuilder:validation:Optional
	VPCClassicLinkSecurityGroups []*string `json:"vpcClassicLinkSecurityGroups,omitempty" tf:"vpc_classic_link_security_groups,omitempty"`
}

type MetadataOptionsObservation struct {

	// The state of the metadata service: enabled, disabled.
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	// The desired HTTP PUT response hop limit for instance metadata requests.
	HTTPPutResponseHopLimit *float64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit,omitempty"`

	// If session tokens are required: optional, required.
	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens,omitempty"`
}

type MetadataOptionsParameters struct {

	// The state of the metadata service: enabled, disabled.
	// +kubebuilder:validation:Optional
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	// The desired HTTP PUT response hop limit for instance metadata requests.
	// +kubebuilder:validation:Optional
	HTTPPutResponseHopLimit *float64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit,omitempty"`

	// If session tokens are required: optional, required.
	// +kubebuilder:validation:Optional
	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens,omitempty"`
}

type RootBlockDeviceObservation struct {

	// Whether the volume should be destroyed on instance termination. Defaults to true.
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Whether the volume should be encrypted or not. Defaults to false.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The amount of provisioned IOPS. This must be set with a volume_type of io1.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The throughput (MiBps) to provision for a gp3 volume.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// The size of the volume in gigabytes.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// The type of volume. Can be standard, gp2, gp3, st1, sc1 or io1.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type RootBlockDeviceParameters struct {

	// Whether the volume should be destroyed on instance termination. Defaults to true.
	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Whether the volume should be encrypted or not. Defaults to false.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The amount of provisioned IOPS. This must be set with a volume_type of io1.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The throughput (MiBps) to provision for a gp3 volume.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// The size of the volume in gigabytes.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// The type of volume. Can be standard, gp2, gp3, st1, sc1 or io1.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

// LaunchConfigurationSpec defines the desired state of LaunchConfiguration
type LaunchConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LaunchConfigurationParameters `json:"forProvider"`
}

// LaunchConfigurationStatus defines the observed state of LaunchConfiguration.
type LaunchConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LaunchConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchConfiguration is the Schema for the LaunchConfigurations API. Provides a resource to create a new launch configuration, used for autoscaling groups.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LaunchConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LaunchConfigurationSpec   `json:"spec"`
	Status            LaunchConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchConfigurationList contains a list of LaunchConfigurations
type LaunchConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LaunchConfiguration `json:"items"`
}

// Repository type metadata.
var (
	LaunchConfiguration_Kind             = "LaunchConfiguration"
	LaunchConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LaunchConfiguration_Kind}.String()
	LaunchConfiguration_KindAPIVersion   = LaunchConfiguration_Kind + "." + CRDGroupVersion.String()
	LaunchConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(LaunchConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&LaunchConfiguration{}, &LaunchConfigurationList{})
}
