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

type HostedPublicVirtualInterfaceAccepterObservation struct {

	// The ARN of the virtual interface.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the virtual interface.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type HostedPublicVirtualInterfaceAccepterParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Direct Connect virtual interface to accept.
	// +crossplane:generate:reference:type=HostedPublicVirtualInterface
	// +kubebuilder:validation:Optional
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`

	// Reference to a HostedPublicVirtualInterface to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDRef *v1.Reference `json:"virtualInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a HostedPublicVirtualInterface to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDSelector *v1.Selector `json:"virtualInterfaceIdSelector,omitempty" tf:"-"`
}

// HostedPublicVirtualInterfaceAccepterSpec defines the desired state of HostedPublicVirtualInterfaceAccepter
type HostedPublicVirtualInterfaceAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedPublicVirtualInterfaceAccepterParameters `json:"forProvider"`
}

// HostedPublicVirtualInterfaceAccepterStatus defines the observed state of HostedPublicVirtualInterfaceAccepter.
type HostedPublicVirtualInterfaceAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedPublicVirtualInterfaceAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPublicVirtualInterfaceAccepter is the Schema for the HostedPublicVirtualInterfaceAccepters API. Provides a resource to manage the accepter's side of a Direct Connect hosted public virtual interface.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HostedPublicVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedPublicVirtualInterfaceAccepterSpec   `json:"spec"`
	Status            HostedPublicVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPublicVirtualInterfaceAccepterList contains a list of HostedPublicVirtualInterfaceAccepters
type HostedPublicVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedPublicVirtualInterfaceAccepter `json:"items"`
}

// Repository type metadata.
var (
	HostedPublicVirtualInterfaceAccepter_Kind             = "HostedPublicVirtualInterfaceAccepter"
	HostedPublicVirtualInterfaceAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedPublicVirtualInterfaceAccepter_Kind}.String()
	HostedPublicVirtualInterfaceAccepter_KindAPIVersion   = HostedPublicVirtualInterfaceAccepter_Kind + "." + CRDGroupVersion.String()
	HostedPublicVirtualInterfaceAccepter_GroupVersionKind = CRDGroupVersion.WithKind(HostedPublicVirtualInterfaceAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedPublicVirtualInterfaceAccepter{}, &HostedPublicVirtualInterfaceAccepterList{})
}