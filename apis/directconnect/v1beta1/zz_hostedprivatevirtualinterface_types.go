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

type HostedPrivateVirtualInterfaceObservation struct {
	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	// The ARN of the virtual interface.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`

	// The ID of the virtual interface.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`
}

type HostedPrivateVirtualInterfaceParameters struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	// +kubebuilder:validation:Optional
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	// +kubebuilder:validation:Optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	// +kubebuilder:validation:Optional
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	// +kubebuilder:validation:Optional
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	// +crossplane:generate:reference:type=Connection
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Reference to a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// Selector for a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	// +kubebuilder:validation:Optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either 1500 or 9001 (jumbo frames). Default is 1500.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name for the virtual interface.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The AWS account that will own the new virtual interface.
	// +kubebuilder:validation:Optional
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The VLAN ID.
	// +kubebuilder:validation:Optional
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

// HostedPrivateVirtualInterfaceSpec defines the desired state of HostedPrivateVirtualInterface
type HostedPrivateVirtualInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedPrivateVirtualInterfaceParameters `json:"forProvider"`
}

// HostedPrivateVirtualInterfaceStatus defines the observed state of HostedPrivateVirtualInterface.
type HostedPrivateVirtualInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedPrivateVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPrivateVirtualInterface is the Schema for the HostedPrivateVirtualInterfaces API. Provides a Direct Connect hosted private virtual interface resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HostedPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.addressFamily)",message="addressFamily is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.bgpAsn)",message="bgpAsn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.ownerAccountId)",message="ownerAccountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vlan)",message="vlan is a required parameter"
	Spec   HostedPrivateVirtualInterfaceSpec   `json:"spec"`
	Status HostedPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPrivateVirtualInterfaceList contains a list of HostedPrivateVirtualInterfaces
type HostedPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedPrivateVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	HostedPrivateVirtualInterface_Kind             = "HostedPrivateVirtualInterface"
	HostedPrivateVirtualInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedPrivateVirtualInterface_Kind}.String()
	HostedPrivateVirtualInterface_KindAPIVersion   = HostedPrivateVirtualInterface_Kind + "." + CRDGroupVersion.String()
	HostedPrivateVirtualInterface_GroupVersionKind = CRDGroupVersion.WithKind(HostedPrivateVirtualInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedPrivateVirtualInterface{}, &HostedPrivateVirtualInterfaceList{})
}
