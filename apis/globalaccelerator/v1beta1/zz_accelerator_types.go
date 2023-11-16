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

type AcceleratorInitParameters struct {

	// The attributes of the accelerator. Fields documented below.
	Attributes []AttributesInitParameters `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// Indicates whether the accelerator is enabled. Defaults to true. Valid values: true, false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The value for the address type. Defaults to IPV4. Valid values: IPV4, DUAL_STACK.
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// The name of the accelerator.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AcceleratorObservation struct {

	// The attributes of the accelerator. Fields documented below.
	Attributes []AttributesObservation `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The DNS name of the accelerator. For example, a5d53ff5ee6bca4ce.awsglobalaccelerator.com.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The Domain Name System (DNS) name that Global Accelerator creates that points to a dual-stack accelerator's four static IP addresses: two IPv4 addresses and two IPv6 addresses. For example, a1234567890abcdef.dualstack.awsglobalaccelerator.com.
	DualStackDNSName *string `json:"dualStackDnsName,omitempty" tf:"dual_stack_dns_name,omitempty"`

	// Indicates whether the accelerator is enabled. Defaults to true. Valid values: true, false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// -  The Global Accelerator Route 53 zone ID that can be used to
	// route an Alias Resource Record Set to the Global Accelerator. This attribute
	// is simply an alias for the zone ID Z2BJ6XQ5FK7U4H.
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// The Amazon Resource Name (ARN) of the accelerator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The value for the address type. Defaults to IPV4. Valid values: IPV4, DUAL_STACK.
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// IP address set associated with the accelerator.
	IPSets []IPSetsObservation `json:"ipSets,omitempty" tf:"ip_sets,omitempty"`

	// The name of the accelerator.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AcceleratorParameters struct {

	// The attributes of the accelerator. Fields documented below.
	// +kubebuilder:validation:Optional
	Attributes []AttributesParameters `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// Indicates whether the accelerator is enabled. Defaults to true. Valid values: true, false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The value for the address type. Defaults to IPV4. Valid values: IPV4, DUAL_STACK.
	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
	// +kubebuilder:validation:Optional
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// The name of the accelerator.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AttributesInitParameters struct {

	// Indicates whether flow logs are enabled. Defaults to false. Valid values: true, false.
	FlowLogsEnabled *bool `json:"flowLogsEnabled,omitempty" tf:"flow_logs_enabled,omitempty"`

	// The name of the Amazon S3 bucket for the flow logs. Required if flow_logs_enabled is true.
	FlowLogsS3Bucket *string `json:"flowLogsS3Bucket,omitempty" tf:"flow_logs_s3_bucket,omitempty"`

	// The prefix for the location in the Amazon S3 bucket for the flow logs. Required if flow_logs_enabled is true.
	FlowLogsS3Prefix *string `json:"flowLogsS3Prefix,omitempty" tf:"flow_logs_s3_prefix,omitempty"`
}

type AttributesObservation struct {

	// Indicates whether flow logs are enabled. Defaults to false. Valid values: true, false.
	FlowLogsEnabled *bool `json:"flowLogsEnabled,omitempty" tf:"flow_logs_enabled,omitempty"`

	// The name of the Amazon S3 bucket for the flow logs. Required if flow_logs_enabled is true.
	FlowLogsS3Bucket *string `json:"flowLogsS3Bucket,omitempty" tf:"flow_logs_s3_bucket,omitempty"`

	// The prefix for the location in the Amazon S3 bucket for the flow logs. Required if flow_logs_enabled is true.
	FlowLogsS3Prefix *string `json:"flowLogsS3Prefix,omitempty" tf:"flow_logs_s3_prefix,omitempty"`
}

type AttributesParameters struct {

	// Indicates whether flow logs are enabled. Defaults to false. Valid values: true, false.
	// +kubebuilder:validation:Optional
	FlowLogsEnabled *bool `json:"flowLogsEnabled,omitempty" tf:"flow_logs_enabled,omitempty"`

	// The name of the Amazon S3 bucket for the flow logs. Required if flow_logs_enabled is true.
	// +kubebuilder:validation:Optional
	FlowLogsS3Bucket *string `json:"flowLogsS3Bucket,omitempty" tf:"flow_logs_s3_bucket,omitempty"`

	// The prefix for the location in the Amazon S3 bucket for the flow logs. Required if flow_logs_enabled is true.
	// +kubebuilder:validation:Optional
	FlowLogsS3Prefix *string `json:"flowLogsS3Prefix,omitempty" tf:"flow_logs_s3_prefix,omitempty"`
}

type IPSetsInitParameters struct {
}

type IPSetsObservation struct {

	// The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// The type of IP addresses included in this IP set.
	IPFamily *string `json:"ipFamily,omitempty" tf:"ip_family,omitempty"`
}

type IPSetsParameters struct {
}

// AcceleratorSpec defines the desired state of Accelerator
type AcceleratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AcceleratorParameters `json:"forProvider"`
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
	InitProvider AcceleratorInitParameters `json:"initProvider,omitempty"`
}

// AcceleratorStatus defines the observed state of Accelerator.
type AcceleratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AcceleratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Accelerator is the Schema for the Accelerators API. Provides a Global Accelerator accelerator.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Accelerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AcceleratorSpec   `json:"spec"`
	Status AcceleratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AcceleratorList contains a list of Accelerators
type AcceleratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Accelerator `json:"items"`
}

// Repository type metadata.
var (
	Accelerator_Kind             = "Accelerator"
	Accelerator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Accelerator_Kind}.String()
	Accelerator_KindAPIVersion   = Accelerator_Kind + "." + CRDGroupVersion.String()
	Accelerator_GroupVersionKind = CRDGroupVersion.WithKind(Accelerator_Kind)
)

func init() {
	SchemeBuilder.Register(&Accelerator{}, &AcceleratorList{})
}
