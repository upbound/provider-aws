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

type InstancePublicPortsObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Lightsail Instance.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Configuration block with port information. AWS closes all currently open ports that are not included in the port_info. Detailed below.
	PortInfo []PortInfoObservation `json:"portInfo,omitempty" tf:"port_info,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type InstancePublicPortsParameters struct {

	// Name of the Lightsail Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta1.Instance
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// Configuration block with port information. AWS closes all currently open ports that are not included in the port_info. Detailed below.
	// +kubebuilder:validation:Required
	PortInfo []PortInfoParameters `json:"portInfo" tf:"port_info,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type PortInfoObservation struct {

	// Set of CIDR aliases that define access for a preconfigured range of IP addresses.
	CidrListAliases []*string `json:"cidrListAliases,omitempty" tf:"cidr_list_aliases,omitempty"`

	// Set of CIDR blocks.
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`

	// First port in a range of open ports on an instance.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	IPv6Cidrs []*string `json:"ipv6Cidrs,omitempty" tf:"ipv6_cidrs,omitempty"`

	// IP protocol name. Valid values are tcp, all, udp, and icmp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Last port in a range of open ports on an instance.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type PortInfoParameters struct {

	// Set of CIDR aliases that define access for a preconfigured range of IP addresses.
	// +kubebuilder:validation:Optional
	CidrListAliases []*string `json:"cidrListAliases,omitempty" tf:"cidr_list_aliases,omitempty"`

	// Set of CIDR blocks.
	// +kubebuilder:validation:Optional
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`

	// First port in a range of open ports on an instance.
	// +kubebuilder:validation:Required
	FromPort *float64 `json:"fromPort" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Cidrs []*string `json:"ipv6Cidrs,omitempty" tf:"ipv6_cidrs,omitempty"`

	// IP protocol name. Valid values are tcp, all, udp, and icmp.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Last port in a range of open ports on an instance.
	// +kubebuilder:validation:Required
	ToPort *float64 `json:"toPort" tf:"to_port,omitempty"`
}

// InstancePublicPortsSpec defines the desired state of InstancePublicPorts
type InstancePublicPortsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstancePublicPortsParameters `json:"forProvider"`
}

// InstancePublicPortsStatus defines the observed state of InstancePublicPorts.
type InstancePublicPortsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstancePublicPortsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstancePublicPorts is the Schema for the InstancePublicPortss API. Provides an Lightsail Instance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstancePublicPorts struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstancePublicPortsSpec   `json:"spec"`
	Status            InstancePublicPortsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstancePublicPortsList contains a list of InstancePublicPortss
type InstancePublicPortsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstancePublicPorts `json:"items"`
}

// Repository type metadata.
var (
	InstancePublicPorts_Kind             = "InstancePublicPorts"
	InstancePublicPorts_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstancePublicPorts_Kind}.String()
	InstancePublicPorts_KindAPIVersion   = InstancePublicPorts_Kind + "." + CRDGroupVersion.String()
	InstancePublicPorts_GroupVersionKind = CRDGroupVersion.WithKind(InstancePublicPorts_Kind)
)

func init() {
	SchemeBuilder.Register(&InstancePublicPorts{}, &InstancePublicPortsList{})
}
