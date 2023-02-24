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

type IdentityProviderConfigObservation struct {

	// Amazon Resource Name (ARN) of the EKS Identity Provider Configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// –  Name of the EKS Cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// EKS Cluster name and EKS Identity Provider Configuration name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Nested attribute containing OpenID Connect identity provider information for the cluster. Detailed below.
	Oidc []IdentityProviderConfigOidcObservation `json:"oidc,omitempty" tf:"oidc,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Status of the EKS Identity Provider Configuration.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type IdentityProviderConfigOidcObservation struct {

	// –  Client ID for the OpenID Connect identity provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The JWT claim that the provider will use to return groups.
	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// A prefix that is prepended to group claims e.g., oidc:.
	GroupsPrefix *string `json:"groupsPrefix,omitempty" tf:"groups_prefix,omitempty"`

	// Issuer URL for the OpenID Connect identity provider.
	IssuerURL *string `json:"issuerUrl,omitempty" tf:"issuer_url,omitempty"`

	// The key value pairs that describe required claims in the identity token.
	RequiredClaims map[string]*string `json:"requiredClaims,omitempty" tf:"required_claims,omitempty"`

	// The JWT claim that the provider will use as the username.
	UsernameClaim *string `json:"usernameClaim,omitempty" tf:"username_claim,omitempty"`

	// A prefix that is prepended to username claims.
	UsernamePrefix *string `json:"usernamePrefix,omitempty" tf:"username_prefix,omitempty"`
}

type IdentityProviderConfigOidcParameters struct {

	// –  Client ID for the OpenID Connect identity provider.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// The JWT claim that the provider will use to return groups.
	// +kubebuilder:validation:Optional
	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// A prefix that is prepended to group claims e.g., oidc:.
	// +kubebuilder:validation:Optional
	GroupsPrefix *string `json:"groupsPrefix,omitempty" tf:"groups_prefix,omitempty"`

	// Issuer URL for the OpenID Connect identity provider.
	// +kubebuilder:validation:Required
	IssuerURL *string `json:"issuerUrl" tf:"issuer_url,omitempty"`

	// The key value pairs that describe required claims in the identity token.
	// +kubebuilder:validation:Optional
	RequiredClaims map[string]*string `json:"requiredClaims,omitempty" tf:"required_claims,omitempty"`

	// The JWT claim that the provider will use as the username.
	// +kubebuilder:validation:Optional
	UsernameClaim *string `json:"usernameClaim,omitempty" tf:"username_claim,omitempty"`

	// A prefix that is prepended to username claims.
	// +kubebuilder:validation:Optional
	UsernamePrefix *string `json:"usernamePrefix,omitempty" tf:"username_prefix,omitempty"`
}

type IdentityProviderConfigParameters struct {

	// –  Name of the EKS Cluster.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// Nested attribute containing OpenID Connect identity provider information for the cluster. Detailed below.
	// +kubebuilder:validation:Required
	Oidc []IdentityProviderConfigOidcParameters `json:"oidc" tf:"oidc,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// IdentityProviderConfigSpec defines the desired state of IdentityProviderConfig
type IdentityProviderConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderConfigParameters `json:"forProvider"`
}

// IdentityProviderConfigStatus defines the observed state of IdentityProviderConfig.
type IdentityProviderConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderConfig is the Schema for the IdentityProviderConfigs API. Manages an EKS Identity Provider Configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IdentityProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IdentityProviderConfigSpec   `json:"spec"`
	Status            IdentityProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderConfigList contains a list of IdentityProviderConfigs
type IdentityProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderConfig `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderConfig_Kind             = "IdentityProviderConfig"
	IdentityProviderConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderConfig_Kind}.String()
	IdentityProviderConfig_KindAPIVersion   = IdentityProviderConfig_Kind + "." + CRDGroupVersion.String()
	IdentityProviderConfig_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProviderConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProviderConfig{}, &IdentityProviderConfigList{})
}
