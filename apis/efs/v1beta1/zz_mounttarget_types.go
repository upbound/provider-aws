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

type MountTargetObservation struct {

	// The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneID *string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id,omitempty"`

	// The name of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name,omitempty"`

	// The DNS name for the EFS file system.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// Amazon Resource Name of the file system.
	FileSystemArn *string `json:"fileSystemArn,omitempty" tf:"file_system_arn,omitempty"`

	// The ID of the file system for which the mount target is intended.
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// The ID of the mount target.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The address (within the address range of the specified subnet) at
	// which the file system may be mounted via the mount target.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The DNS name for the given subnet/AZ per documented convention.
	MountTargetDNSName *string `json:"mountTargetDnsName,omitempty" tf:"mount_target_dns_name,omitempty"`

	// The ID of the network interface that Amazon EFS created when it created the mount target.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// AWS account ID that owns the resource.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A list of up to 5 VPC security group IDs (that must
	// be for the same VPC as subnet specified) in effect for the mount target.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The ID of the subnet to add the mount target in.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type MountTargetParameters struct {

	// The ID of the file system for which the mount target is intended.
	// +crossplane:generate:reference:type=FileSystem
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Reference to a FileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDRef *v1.Reference `json:"fileSystemIdRef,omitempty" tf:"-"`

	// Selector for a FileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDSelector *v1.Selector `json:"fileSystemIdSelector,omitempty" tf:"-"`

	// The address (within the address range of the specified subnet) at
	// which the file system may be mounted via the mount target.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A list of up to 5 VPC security group IDs (that must
	// be for the same VPC as subnet specified) in effect for the mount target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsRefs []v1.Reference `json:"securityGroupsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsSelector *v1.Selector `json:"securityGroupsSelector,omitempty" tf:"-"`

	// The ID of the subnet to add the mount target in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// MountTargetSpec defines the desired state of MountTarget
type MountTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MountTargetParameters `json:"forProvider"`
}

// MountTargetStatus defines the observed state of MountTarget.
type MountTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MountTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MountTarget is the Schema for the MountTargets API. Provides an Elastic File System (EFS) mount target.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MountTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MountTargetSpec   `json:"spec"`
	Status            MountTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MountTargetList contains a list of MountTargets
type MountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MountTarget `json:"items"`
}

// Repository type metadata.
var (
	MountTarget_Kind             = "MountTarget"
	MountTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MountTarget_Kind}.String()
	MountTarget_KindAPIVersion   = MountTarget_Kind + "." + CRDGroupVersion.String()
	MountTarget_GroupVersionKind = CRDGroupVersion.WithKind(MountTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&MountTarget{}, &MountTargetList{})
}
