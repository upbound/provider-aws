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

type OrganizationalUnitAccountsObservation struct {

	// ARN of the account
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Email of the account
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Identifier of the account
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name for the organizational unit
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type OrganizationalUnitAccountsParameters struct {
}

type OrganizationalUnitObservation struct {

	// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
	Accounts []OrganizationalUnitAccountsObservation `json:"accounts,omitempty" tf:"accounts,omitempty"`

	// ARN of the account
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Identifier of the account
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name for the organizational unit
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the parent organizational unit, which may be the root
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type OrganizationalUnitParameters struct {

	// The name for the organizational unit
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the parent organizational unit, which may be the root
	// +kubebuilder:validation:Optional
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// OrganizationalUnitSpec defines the desired state of OrganizationalUnit
type OrganizationalUnitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationalUnitParameters `json:"forProvider"`
}

// OrganizationalUnitStatus defines the observed state of OrganizationalUnit.
type OrganizationalUnitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationalUnitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationalUnit is the Schema for the OrganizationalUnits API. Provides a resource to create an organizational unit.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OrganizationalUnit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.parentId)",message="parentId is a required parameter"
	Spec   OrganizationalUnitSpec   `json:"spec"`
	Status OrganizationalUnitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationalUnitList contains a list of OrganizationalUnits
type OrganizationalUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationalUnit `json:"items"`
}

// Repository type metadata.
var (
	OrganizationalUnit_Kind             = "OrganizationalUnit"
	OrganizationalUnit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationalUnit_Kind}.String()
	OrganizationalUnit_KindAPIVersion   = OrganizationalUnit_Kind + "." + CRDGroupVersion.String()
	OrganizationalUnit_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationalUnit_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationalUnit{}, &OrganizationalUnitList{})
}
