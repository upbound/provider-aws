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

type NetworkInterfaceAttachmentInitParameters struct {

	// Network interface index (int).
	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`
}

type NetworkInterfaceAttachmentObservation struct {

	// The ENI Attachment ID.
	AttachmentID *string `json:"attachmentId,omitempty" tf:"attachment_id,omitempty"`

	// Network interface index (int).
	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance ID to attach.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// ENI ID to attach.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The status of the Network Interface Attachment.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NetworkInterfaceAttachmentParameters struct {

	// Network interface index (int).
	// +kubebuilder:validation:Optional
	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`

	// Instance ID to attach.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// ENI ID to attach.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface in ec2 to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface in ec2 to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// NetworkInterfaceAttachmentSpec defines the desired state of NetworkInterfaceAttachment
type NetworkInterfaceAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceAttachmentParameters `json:"forProvider"`
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
	InitProvider NetworkInterfaceAttachmentInitParameters `json:"initProvider,omitempty"`
}

// NetworkInterfaceAttachmentStatus defines the observed state of NetworkInterfaceAttachment.
type NetworkInterfaceAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceAttachment is the Schema for the NetworkInterfaceAttachments API. Attach an Elastic network interface (ENI) resource with EC2 instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkInterfaceAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.deviceIndex) || (has(self.initProvider) && has(self.initProvider.deviceIndex))",message="spec.forProvider.deviceIndex is a required parameter"
	Spec   NetworkInterfaceAttachmentSpec   `json:"spec"`
	Status NetworkInterfaceAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceAttachmentList contains a list of NetworkInterfaceAttachments
type NetworkInterfaceAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceAttachment `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceAttachment_Kind             = "NetworkInterfaceAttachment"
	NetworkInterfaceAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterfaceAttachment_Kind}.String()
	NetworkInterfaceAttachment_KindAPIVersion   = NetworkInterfaceAttachment_Kind + "." + CRDGroupVersion.String()
	NetworkInterfaceAttachment_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterfaceAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceAttachment{}, &NetworkInterfaceAttachmentList{})
}
