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

type DefaultSecurityGroupEgressInitParameters struct {

	// List of CIDR blocks.
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	// Description of this rule.
	Description *string `json:"description,omitempty" tf:"description"`

	// Start port (or ICMP type number if protocol is icmp)
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port"`

	// List of IPv6 CIDR blocks.
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks"`

	// List of prefix list IDs (for allowing access to VPC endpoints)
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids"`

	// Protocol. If you select a protocol of "-1" (semantically equivalent to all, which is not a valid value here), you must specify a from_port and to_port equal to 0. If not icmp, tcp, udp, or -1 use the protocol number.
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups"`

	// Whether the security group itself will be added as a source to this egress rule.
	Self *bool `json:"self,omitempty" tf:"self"`

	// End range port (or ICMP code if protocol is icmp).
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port"`
}

type DefaultSecurityGroupEgressObservation struct {

	// List of CIDR blocks.
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// Description of this rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Start port (or ICMP type number if protocol is icmp)
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// List of IPv6 CIDR blocks.
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks,omitempty"`

	// List of prefix list IDs (for allowing access to VPC endpoints)
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids,omitempty"`

	// Protocol. If you select a protocol of "-1" (semantically equivalent to all, which is not a valid value here), you must specify a from_port and to_port equal to 0. If not icmp, tcp, udp, or -1 use the protocol number.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Whether the security group itself will be added as a source to this egress rule.
	Self *bool `json:"self,omitempty" tf:"self,omitempty"`

	// End range port (or ICMP code if protocol is icmp).
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type DefaultSecurityGroupEgressParameters struct {

	// List of CIDR blocks.
	// +kubebuilder:validation:Optional
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	// Description of this rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// Start port (or ICMP type number if protocol is icmp)
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port"`

	// List of IPv6 CIDR blocks.
	// +kubebuilder:validation:Optional
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks"`

	// List of prefix list IDs (for allowing access to VPC endpoints)
	// +kubebuilder:validation:Optional
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids"`

	// Protocol. If you select a protocol of "-1" (semantically equivalent to all, which is not a valid value here), you must specify a from_port and to_port equal to 0. If not icmp, tcp, udp, or -1 use the protocol number.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups"`

	// Whether the security group itself will be added as a source to this egress rule.
	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self"`

	// End range port (or ICMP code if protocol is icmp).
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port"`
}

type DefaultSecurityGroupIngressInitParameters struct {

	// List of CIDR blocks.
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	// Description of this rule.
	Description *string `json:"description,omitempty" tf:"description"`

	// Start port (or ICMP type number if protocol is icmp)
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port"`

	// List of IPv6 CIDR blocks.
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks"`

	// List of prefix list IDs (for allowing access to VPC endpoints)
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids"`

	// Protocol. If you select a protocol of "-1" (semantically equivalent to all, which is not a valid value here), you must specify a from_port and to_port equal to 0. If not icmp, tcp, udp, or -1 use the protocol number.
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups"`

	// Whether the security group itself will be added as a source to this egress rule.
	Self *bool `json:"self,omitempty" tf:"self"`

	// End range port (or ICMP code if protocol is icmp).
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port"`
}

type DefaultSecurityGroupIngressObservation struct {

	// List of CIDR blocks.
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// Description of this rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Start port (or ICMP type number if protocol is icmp)
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// List of IPv6 CIDR blocks.
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks,omitempty"`

	// List of prefix list IDs (for allowing access to VPC endpoints)
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids,omitempty"`

	// Protocol. If you select a protocol of "-1" (semantically equivalent to all, which is not a valid value here), you must specify a from_port and to_port equal to 0. If not icmp, tcp, udp, or -1 use the protocol number.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Whether the security group itself will be added as a source to this egress rule.
	Self *bool `json:"self,omitempty" tf:"self,omitempty"`

	// End range port (or ICMP code if protocol is icmp).
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type DefaultSecurityGroupIngressParameters struct {

	// List of CIDR blocks.
	// +kubebuilder:validation:Optional
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	// Description of this rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// Start port (or ICMP type number if protocol is icmp)
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port"`

	// List of IPv6 CIDR blocks.
	// +kubebuilder:validation:Optional
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks"`

	// List of prefix list IDs (for allowing access to VPC endpoints)
	// +kubebuilder:validation:Optional
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids"`

	// Protocol. If you select a protocol of "-1" (semantically equivalent to all, which is not a valid value here), you must specify a from_port and to_port equal to 0. If not icmp, tcp, udp, or -1 use the protocol number.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups"`

	// Whether the security group itself will be added as a source to this egress rule.
	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self"`

	// End range port (or ICMP code if protocol is icmp).
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port"`
}

type DefaultSecurityGroupInitParameters struct {

	// Configuration block. Detailed below.
	Egress []DefaultSecurityGroupEgressInitParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// Configuration block. Detailed below.
	Ingress []DefaultSecurityGroupIngressInitParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	RevokeRulesOnDelete *bool `json:"revokeRulesOnDelete,omitempty" tf:"revoke_rules_on_delete,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DefaultSecurityGroupObservation struct {

	// ARN of the security group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of this rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block. Detailed below.
	Egress []DefaultSecurityGroupEgressObservation `json:"egress,omitempty" tf:"egress,omitempty"`

	// ID of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block. Detailed below.
	Ingress []DefaultSecurityGroupIngressObservation `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Name of the security group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner ID.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	RevokeRulesOnDelete *bool `json:"revokeRulesOnDelete,omitempty" tf:"revoke_rules_on_delete,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// VPC ID. Note that changing the  It will be left in its current state.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DefaultSecurityGroupParameters struct {

	// Configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	Egress []DefaultSecurityGroupEgressParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// Configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	Ingress []DefaultSecurityGroupIngressParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RevokeRulesOnDelete *bool `json:"revokeRulesOnDelete,omitempty" tf:"revoke_rules_on_delete,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// VPC ID. Note that changing the  It will be left in its current state.
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

// DefaultSecurityGroupSpec defines the desired state of DefaultSecurityGroup
type DefaultSecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultSecurityGroupParameters `json:"forProvider"`
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
	InitProvider DefaultSecurityGroupInitParameters `json:"initProvider,omitempty"`
}

// DefaultSecurityGroupStatus defines the observed state of DefaultSecurityGroup.
type DefaultSecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSecurityGroup is the Schema for the DefaultSecurityGroups API. Manage a default security group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DefaultSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSecurityGroupSpec   `json:"spec"`
	Status            DefaultSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSecurityGroupList contains a list of DefaultSecurityGroups
type DefaultSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	DefaultSecurityGroup_Kind             = "DefaultSecurityGroup"
	DefaultSecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultSecurityGroup_Kind}.String()
	DefaultSecurityGroup_KindAPIVersion   = DefaultSecurityGroup_Kind + "." + CRDGroupVersion.String()
	DefaultSecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(DefaultSecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultSecurityGroup{}, &DefaultSecurityGroupList{})
}
