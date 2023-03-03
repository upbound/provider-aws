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

type DefaultVPCDHCPOptionsObservation struct {

	// The ARN of the DHCP Options Set.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	DomainNameServers *string `json:"domainNameServers,omitempty" tf:"domain_name_servers,omitempty"`

	// The ID of the DHCP Options Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of NETBIOS name servers.
	NetbiosNameServers *string `json:"netbiosNameServers,omitempty" tf:"netbios_name_servers,omitempty"`

	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see RFC 2132.
	NetbiosNodeType *string `json:"netbiosNodeType,omitempty" tf:"netbios_node_type,omitempty"`

	NtpServers *string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`

	// The ID of the AWS account that owns the DHCP options set.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DefaultVPCDHCPOptionsParameters struct {

	// The ID of the AWS account that owns the DHCP options set.
	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DefaultVPCDHCPOptionsSpec defines the desired state of DefaultVPCDHCPOptions
type DefaultVPCDHCPOptionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultVPCDHCPOptionsParameters `json:"forProvider"`
}

// DefaultVPCDHCPOptionsStatus defines the observed state of DefaultVPCDHCPOptions.
type DefaultVPCDHCPOptionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultVPCDHCPOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultVPCDHCPOptions is the Schema for the DefaultVPCDHCPOptionss API. Manage the default VPC DHCP Options resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DefaultVPCDHCPOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultVPCDHCPOptionsSpec   `json:"spec"`
	Status            DefaultVPCDHCPOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultVPCDHCPOptionsList contains a list of DefaultVPCDHCPOptionss
type DefaultVPCDHCPOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultVPCDHCPOptions `json:"items"`
}

// Repository type metadata.
var (
	DefaultVPCDHCPOptions_Kind             = "DefaultVPCDHCPOptions"
	DefaultVPCDHCPOptions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultVPCDHCPOptions_Kind}.String()
	DefaultVPCDHCPOptions_KindAPIVersion   = DefaultVPCDHCPOptions_Kind + "." + CRDGroupVersion.String()
	DefaultVPCDHCPOptions_GroupVersionKind = CRDGroupVersion.WithKind(DefaultVPCDHCPOptions_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultVPCDHCPOptions{}, &DefaultVPCDHCPOptionsList{})
}
