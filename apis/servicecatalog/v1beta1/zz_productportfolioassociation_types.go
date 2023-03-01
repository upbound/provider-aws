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

type ProductPortfolioAssociationObservation struct {

	// Identifier of the association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProductPortfolioAssociationParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Portfolio identifier.
	// +crossplane:generate:reference:type=Portfolio
	// +kubebuilder:validation:Optional
	PortfolioID *string `json:"portfolioId,omitempty" tf:"portfolio_id,omitempty"`

	// Reference to a Portfolio to populate portfolioId.
	// +kubebuilder:validation:Optional
	PortfolioIDRef *v1.Reference `json:"portfolioIdRef,omitempty" tf:"-"`

	// Selector for a Portfolio to populate portfolioId.
	// +kubebuilder:validation:Optional
	PortfolioIDSelector *v1.Selector `json:"portfolioIdSelector,omitempty" tf:"-"`

	// Product identifier.
	// +crossplane:generate:reference:type=Product
	// +kubebuilder:validation:Optional
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Reference to a Product to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDRef *v1.Reference `json:"productIdRef,omitempty" tf:"-"`

	// Selector for a Product to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDSelector *v1.Selector `json:"productIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Identifier of the source portfolio.
	// +kubebuilder:validation:Optional
	SourcePortfolioID *string `json:"sourcePortfolioId,omitempty" tf:"source_portfolio_id,omitempty"`
}

// ProductPortfolioAssociationSpec defines the desired state of ProductPortfolioAssociation
type ProductPortfolioAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProductPortfolioAssociationParameters `json:"forProvider"`
}

// ProductPortfolioAssociationStatus defines the observed state of ProductPortfolioAssociation.
type ProductPortfolioAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProductPortfolioAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProductPortfolioAssociation is the Schema for the ProductPortfolioAssociations API. Manages a Service Catalog Product Portfolio Association
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ProductPortfolioAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.acceptLanguage)",message="acceptLanguage is a required parameter"
	Spec   ProductPortfolioAssociationSpec   `json:"spec"`
	Status ProductPortfolioAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProductPortfolioAssociationList contains a list of ProductPortfolioAssociations
type ProductPortfolioAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProductPortfolioAssociation `json:"items"`
}

// Repository type metadata.
var (
	ProductPortfolioAssociation_Kind             = "ProductPortfolioAssociation"
	ProductPortfolioAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProductPortfolioAssociation_Kind}.String()
	ProductPortfolioAssociation_KindAPIVersion   = ProductPortfolioAssociation_Kind + "." + CRDGroupVersion.String()
	ProductPortfolioAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ProductPortfolioAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ProductPortfolioAssociation{}, &ProductPortfolioAssociationList{})
}
