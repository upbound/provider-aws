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

type RouteTableObservation_2 struct {

	// The ARN of the route table.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the routing table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the AWS account that owns the route table.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A list of virtual gateways for propagation.
	PropagatingVgws []*string `json:"propagatingVgws,omitempty" tf:"propagating_vgws,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A list of route objects. Their keys are documented below. This argument is processed in attribute-as-blocks mode.
	// This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	Route []RouteTableRouteObservation_2 `json:"route,omitempty" tf:"route,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The VPC ID.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type RouteTableParameters_2 struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type RouteTableRouteObservation_2 struct {

	// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayID *string `json:"carrierGatewayId,omitempty" tf:"carrier_gateway_id,omitempty"`

	// The CIDR block of the route.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn,omitempty"`

	// The ID of a managed prefix list destination of the route.
	DestinationPrefixListID *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id,omitempty"`

	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayID *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id,omitempty"`

	// Identifier of a VPC internet gateway or a virtual private gateway.
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// The Ipv6 CIDR block of the route.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// Identifier of an EC2 instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Identifier of a Outpost local gateway.
	LocalGatewayID *string `json:"localGatewayId,omitempty" tf:"local_gateway_id,omitempty"`

	// Identifier of a VPC NAT gateway.
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Identifier of an EC2 network interface.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Identifier of an EC2 Transit Gateway.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Identifier of a VPC Endpoint.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// Identifier of a VPC peering connection.
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

type RouteTableRouteParameters_2 struct {
}

// RouteTableSpec defines the desired state of RouteTable
type RouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteTableParameters_2 `json:"forProvider"`
}

// RouteTableStatus defines the observed state of RouteTable.
type RouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteTableObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTable is the Schema for the RouteTables API. Provides a resource to create a VPC routing table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableList contains a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// Repository type metadata.
var (
	RouteTable_Kind             = "RouteTable"
	RouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteTable_Kind}.String()
	RouteTable_KindAPIVersion   = RouteTable_Kind + "." + CRDGroupVersion.String()
	RouteTable_GroupVersionKind = CRDGroupVersion.WithKind(RouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
