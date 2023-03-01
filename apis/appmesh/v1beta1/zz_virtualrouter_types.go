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

type SpecListenerPortMappingObservation struct {
}

type SpecListenerPortMappingParameters struct {

	// Port used for the port mapping.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Protocol used for the port mapping. Valid values are http,http2, tcp and grpc.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type VirtualRouterObservation struct {

	// ARN of the virtual router.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Creation date of the virtual router.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// ID of the virtual router.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Last update date of the virtual router.
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	// Resource owner's AWS account ID.
	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VirtualRouterParameters struct {

	// Name of the service mesh in which to create the virtual router. Must be between 1 and 255 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appmesh/v1beta1.Mesh
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MeshName *string `json:"meshName,omitempty" tf:"mesh_name,omitempty"`

	// Reference to a Mesh in appmesh to populate meshName.
	// +kubebuilder:validation:Optional
	MeshNameRef *v1.Reference `json:"meshNameRef,omitempty" tf:"-"`

	// Selector for a Mesh in appmesh to populate meshName.
	// +kubebuilder:validation:Optional
	MeshNameSelector *v1.Selector `json:"meshNameSelector,omitempty" tf:"-"`

	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	// +kubebuilder:validation:Optional
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner,omitempty"`

	// Name to use for the virtual router. Must be between 1 and 255 characters in length.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Virtual router specification to apply.
	// +kubebuilder:validation:Optional
	Spec []VirtualRouterSpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualRouterSpecListenerObservation struct {
}

type VirtualRouterSpecListenerParameters struct {

	// Port mapping information for the listener.
	// +kubebuilder:validation:Required
	PortMapping []SpecListenerPortMappingParameters `json:"portMapping" tf:"port_mapping,omitempty"`
}

type VirtualRouterSpecObservation struct {
}

type VirtualRouterSpecParameters struct {

	// configuration block to the spec argument.
	// +kubebuilder:validation:Required
	Listener []VirtualRouterSpecListenerParameters `json:"listener" tf:"listener,omitempty"`
}

// VirtualRouterSpec defines the desired state of VirtualRouter
type VirtualRouterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualRouterParameters `json:"forProvider"`
}

// VirtualRouterStatus defines the observed state of VirtualRouter.
type VirtualRouterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualRouterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualRouter is the Schema for the VirtualRouters API. Provides an AWS App Mesh virtual router resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VirtualRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.spec)",message="spec is a required parameter"
	Spec   VirtualRouterSpec   `json:"spec"`
	Status VirtualRouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualRouterList contains a list of VirtualRouters
type VirtualRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualRouter `json:"items"`
}

// Repository type metadata.
var (
	VirtualRouter_Kind             = "VirtualRouter"
	VirtualRouter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualRouter_Kind}.String()
	VirtualRouter_KindAPIVersion   = VirtualRouter_Kind + "." + CRDGroupVersion.String()
	VirtualRouter_GroupVersionKind = CRDGroupVersion.WithKind(VirtualRouter_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualRouter{}, &VirtualRouterList{})
}
