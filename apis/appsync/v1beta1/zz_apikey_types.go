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

type APIKeyObservation struct {

	// ID of the associated AppSync API
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// API key description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// API Key ID (Formatted as ApiId:Key)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type APIKeyParameters struct {

	// ID of the associated AppSync API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta1.GraphQLAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// API key description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	// +kubebuilder:validation:Optional
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// APIKeySpec defines the desired state of APIKey
type APIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIKeyParameters `json:"forProvider"`
}

// APIKeyStatus defines the observed state of APIKey.
type APIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIKey is the Schema for the APIKeys API. Provides an AppSync API Key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type APIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIKeySpec   `json:"spec"`
	Status            APIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIKeyList contains a list of APIKeys
type APIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIKey `json:"items"`
}

// Repository type metadata.
var (
	APIKey_Kind             = "APIKey"
	APIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIKey_Kind}.String()
	APIKey_KindAPIVersion   = APIKey_Kind + "." + CRDGroupVersion.String()
	APIKey_GroupVersionKind = CRDGroupVersion.WithKind(APIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&APIKey{}, &APIKeyList{})
}
