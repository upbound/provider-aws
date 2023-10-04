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

type NetworkProfileInitParameters struct {

	// The description of the network profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600. Default value is 104857600.
	DownlinkBandwidthBits *int64 `json:"downlinkBandwidthBits,omitempty" tf:"downlink_bandwidth_bits,omitempty"`

	// Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.
	DownlinkDelayMs *int64 `json:"downlinkDelayMs,omitempty" tf:"downlink_delay_ms,omitempty"`

	// Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.
	DownlinkJitterMs *int64 `json:"downlinkJitterMs,omitempty" tf:"downlink_jitter_ms,omitempty"`

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	DownlinkLossPercent *int64 `json:"downlinkLossPercent,omitempty" tf:"downlink_loss_percent,omitempty"`

	// The name for the network profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of network profile to create. Valid values are listed are PRIVATE and CURATED.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600. Default value is 104857600.
	UplinkBandwidthBits *int64 `json:"uplinkBandwidthBits,omitempty" tf:"uplink_bandwidth_bits,omitempty"`

	// Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.
	UplinkDelayMs *int64 `json:"uplinkDelayMs,omitempty" tf:"uplink_delay_ms,omitempty"`

	// Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.
	UplinkJitterMs *int64 `json:"uplinkJitterMs,omitempty" tf:"uplink_jitter_ms,omitempty"`

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	UplinkLossPercent *int64 `json:"uplinkLossPercent,omitempty" tf:"uplink_loss_percent,omitempty"`
}

type NetworkProfileObservation struct {

	// The Amazon Resource Name of this network profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description of the network profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600. Default value is 104857600.
	DownlinkBandwidthBits *int64 `json:"downlinkBandwidthBits,omitempty" tf:"downlink_bandwidth_bits,omitempty"`

	// Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.
	DownlinkDelayMs *int64 `json:"downlinkDelayMs,omitempty" tf:"downlink_delay_ms,omitempty"`

	// Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.
	DownlinkJitterMs *int64 `json:"downlinkJitterMs,omitempty" tf:"downlink_jitter_ms,omitempty"`

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	DownlinkLossPercent *int64 `json:"downlinkLossPercent,omitempty" tf:"downlink_loss_percent,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name for the network profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the project for the network profile.
	ProjectArn *string `json:"projectArn,omitempty" tf:"project_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The type of network profile to create. Valid values are listed are PRIVATE and CURATED.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600. Default value is 104857600.
	UplinkBandwidthBits *int64 `json:"uplinkBandwidthBits,omitempty" tf:"uplink_bandwidth_bits,omitempty"`

	// Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.
	UplinkDelayMs *int64 `json:"uplinkDelayMs,omitempty" tf:"uplink_delay_ms,omitempty"`

	// Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.
	UplinkJitterMs *int64 `json:"uplinkJitterMs,omitempty" tf:"uplink_jitter_ms,omitempty"`

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	UplinkLossPercent *int64 `json:"uplinkLossPercent,omitempty" tf:"uplink_loss_percent,omitempty"`
}

type NetworkProfileParameters struct {

	// The description of the network profile.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600. Default value is 104857600.
	// +kubebuilder:validation:Optional
	DownlinkBandwidthBits *int64 `json:"downlinkBandwidthBits,omitempty" tf:"downlink_bandwidth_bits,omitempty"`

	// Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.
	// +kubebuilder:validation:Optional
	DownlinkDelayMs *int64 `json:"downlinkDelayMs,omitempty" tf:"downlink_delay_ms,omitempty"`

	// Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.
	// +kubebuilder:validation:Optional
	DownlinkJitterMs *int64 `json:"downlinkJitterMs,omitempty" tf:"downlink_jitter_ms,omitempty"`

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	// +kubebuilder:validation:Optional
	DownlinkLossPercent *int64 `json:"downlinkLossPercent,omitempty" tf:"downlink_loss_percent,omitempty"`

	// The name for the network profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the project for the network profile.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/devicefarm/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ProjectArn *string `json:"projectArn,omitempty" tf:"project_arn,omitempty"`

	// Reference to a Project in devicefarm to populate projectArn.
	// +kubebuilder:validation:Optional
	ProjectArnRef *v1.Reference `json:"projectArnRef,omitempty" tf:"-"`

	// Selector for a Project in devicefarm to populate projectArn.
	// +kubebuilder:validation:Optional
	ProjectArnSelector *v1.Selector `json:"projectArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of network profile to create. Valid values are listed are PRIVATE and CURATED.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600. Default value is 104857600.
	// +kubebuilder:validation:Optional
	UplinkBandwidthBits *int64 `json:"uplinkBandwidthBits,omitempty" tf:"uplink_bandwidth_bits,omitempty"`

	// Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.
	// +kubebuilder:validation:Optional
	UplinkDelayMs *int64 `json:"uplinkDelayMs,omitempty" tf:"uplink_delay_ms,omitempty"`

	// Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.
	// +kubebuilder:validation:Optional
	UplinkJitterMs *int64 `json:"uplinkJitterMs,omitempty" tf:"uplink_jitter_ms,omitempty"`

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	// +kubebuilder:validation:Optional
	UplinkLossPercent *int64 `json:"uplinkLossPercent,omitempty" tf:"uplink_loss_percent,omitempty"`
}

// NetworkProfileSpec defines the desired state of NetworkProfile
type NetworkProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkProfileParameters `json:"forProvider"`
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
	InitProvider NetworkProfileInitParameters `json:"initProvider,omitempty"`
}

// NetworkProfileStatus defines the observed state of NetworkProfile.
type NetworkProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkProfile is the Schema for the NetworkProfiles API. Provides a Devicefarm network profile
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   NetworkProfileSpec   `json:"spec"`
	Status NetworkProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkProfileList contains a list of NetworkProfiles
type NetworkProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkProfile `json:"items"`
}

// Repository type metadata.
var (
	NetworkProfile_Kind             = "NetworkProfile"
	NetworkProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkProfile_Kind}.String()
	NetworkProfile_KindAPIVersion   = NetworkProfile_Kind + "." + CRDGroupVersion.String()
	NetworkProfile_GroupVersionKind = CRDGroupVersion.WithKind(NetworkProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkProfile{}, &NetworkProfileList{})
}
