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

type CognitoIdentityProvidersInitParameters struct {

	// The provider name for an Amazon Cognito Identity User Pool.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Whether server-side token validation is enabled for the identity provider’s token or not.
	ServerSideTokenCheck *bool `json:"serverSideTokenCheck,omitempty" tf:"server_side_token_check,omitempty"`
}

type CognitoIdentityProvidersObservation struct {

	// The client ID for the Amazon Cognito Identity User Pool.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The provider name for an Amazon Cognito Identity User Pool.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Whether server-side token validation is enabled for the identity provider’s token or not.
	ServerSideTokenCheck *bool `json:"serverSideTokenCheck,omitempty" tf:"server_side_token_check,omitempty"`
}

type CognitoIdentityProvidersParameters struct {

	// The client ID for the Amazon Cognito Identity User Pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolClient
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The provider name for an Amazon Cognito Identity User Pool.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Whether server-side token validation is enabled for the identity provider’s token or not.
	// +kubebuilder:validation:Optional
	ServerSideTokenCheck *bool `json:"serverSideTokenCheck,omitempty" tf:"server_side_token_check,omitempty"`
}

type PoolInitParameters struct {

	// Enables or disables the classic / basic authentication flow. Default is false.
	AllowClassicFlow *bool `json:"allowClassicFlow,omitempty" tf:"allow_classic_flow,omitempty"`

	// Whether the identity pool supports unauthenticated logins or not.
	AllowUnauthenticatedIdentities *bool `json:"allowUnauthenticatedIdentities,omitempty" tf:"allow_unauthenticated_identities,omitempty"`

	// An array of Amazon Cognito Identity user pools and their client IDs.
	CognitoIdentityProviders []CognitoIdentityProvidersInitParameters `json:"cognitoIdentityProviders,omitempty" tf:"cognito_identity_providers,omitempty"`

	// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
	// backend and the Cognito service to communicate about the developer provider.
	DeveloperProviderName *string `json:"developerProviderName,omitempty" tf:"developer_provider_name,omitempty"`

	// The Cognito Identity Pool name.
	IdentityPoolName *string `json:"identityPoolName,omitempty" tf:"identity_pool_name,omitempty"`

	// Set of OpendID Connect provider ARNs.
	OpenIDConnectProviderArns []*string `json:"openidConnectProviderArns,omitempty" tf:"openid_connect_provider_arns,omitempty"`

	// Key-Value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]*string `json:"supportedLoginProviders,omitempty" tf:"supported_login_providers,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PoolObservation struct {

	// Enables or disables the classic / basic authentication flow. Default is false.
	AllowClassicFlow *bool `json:"allowClassicFlow,omitempty" tf:"allow_classic_flow,omitempty"`

	// Whether the identity pool supports unauthenticated logins or not.
	AllowUnauthenticatedIdentities *bool `json:"allowUnauthenticatedIdentities,omitempty" tf:"allow_unauthenticated_identities,omitempty"`

	// The ARN of the identity pool.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// An array of Amazon Cognito Identity user pools and their client IDs.
	CognitoIdentityProviders []CognitoIdentityProvidersObservation `json:"cognitoIdentityProviders,omitempty" tf:"cognito_identity_providers,omitempty"`

	// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
	// backend and the Cognito service to communicate about the developer provider.
	DeveloperProviderName *string `json:"developerProviderName,omitempty" tf:"developer_provider_name,omitempty"`

	// An identity pool ID, e.g. us-west-2_abc123.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Cognito Identity Pool name.
	IdentityPoolName *string `json:"identityPoolName,omitempty" tf:"identity_pool_name,omitempty"`

	// Set of OpendID Connect provider ARNs.
	OpenIDConnectProviderArns []*string `json:"openidConnectProviderArns,omitempty" tf:"openid_connect_provider_arns,omitempty"`

	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
	SAMLProviderArns []*string `json:"samlProviderArns,omitempty" tf:"saml_provider_arns,omitempty"`

	// Key-Value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]*string `json:"supportedLoginProviders,omitempty" tf:"supported_login_providers,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PoolParameters struct {

	// Enables or disables the classic / basic authentication flow. Default is false.
	// +kubebuilder:validation:Optional
	AllowClassicFlow *bool `json:"allowClassicFlow,omitempty" tf:"allow_classic_flow,omitempty"`

	// Whether the identity pool supports unauthenticated logins or not.
	// +kubebuilder:validation:Optional
	AllowUnauthenticatedIdentities *bool `json:"allowUnauthenticatedIdentities,omitempty" tf:"allow_unauthenticated_identities,omitempty"`

	// An array of Amazon Cognito Identity user pools and their client IDs.
	// +kubebuilder:validation:Optional
	CognitoIdentityProviders []CognitoIdentityProvidersParameters `json:"cognitoIdentityProviders,omitempty" tf:"cognito_identity_providers,omitempty"`

	// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
	// backend and the Cognito service to communicate about the developer provider.
	// +kubebuilder:validation:Optional
	DeveloperProviderName *string `json:"developerProviderName,omitempty" tf:"developer_provider_name,omitempty"`

	// The Cognito Identity Pool name.
	// +kubebuilder:validation:Optional
	IdentityPoolName *string `json:"identityPoolName,omitempty" tf:"identity_pool_name,omitempty"`

	// Set of OpendID Connect provider ARNs.
	// +kubebuilder:validation:Optional
	OpenIDConnectProviderArns []*string `json:"openidConnectProviderArns,omitempty" tf:"openid_connect_provider_arns,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.SAMLProvider
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SAMLProviderArns []*string `json:"samlProviderArns,omitempty" tf:"saml_provider_arns,omitempty"`

	// References to SAMLProvider in iam to populate samlProviderArns.
	// +kubebuilder:validation:Optional
	SAMLProviderArnsRefs []v1.Reference `json:"samlProviderArnsRefs,omitempty" tf:"-"`

	// Selector for a list of SAMLProvider in iam to populate samlProviderArns.
	// +kubebuilder:validation:Optional
	SAMLProviderArnsSelector *v1.Selector `json:"samlProviderArnsSelector,omitempty" tf:"-"`

	// Key-Value pairs mapping provider names to provider app IDs.
	// +kubebuilder:validation:Optional
	SupportedLoginProviders map[string]*string `json:"supportedLoginProviders,omitempty" tf:"supported_login_providers,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PoolSpec defines the desired state of Pool
type PoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PoolParameters `json:"forProvider"`
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
	InitProvider PoolInitParameters `json:"initProvider,omitempty"`
}

// PoolStatus defines the observed state of Pool.
type PoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pool is the Schema for the Pools API. Provides an AWS Cognito Identity Pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Pool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityPoolName) || (has(self.initProvider) && has(self.initProvider.identityPoolName))",message="spec.forProvider.identityPoolName is a required parameter"
	Spec   PoolSpec   `json:"spec"`
	Status PoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PoolList contains a list of Pools
type PoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pool `json:"items"`
}

// Repository type metadata.
var (
	Pool_Kind             = "Pool"
	Pool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pool_Kind}.String()
	Pool_KindAPIVersion   = Pool_Kind + "." + CRDGroupVersion.String()
	Pool_GroupVersionKind = CRDGroupVersion.WithKind(Pool_Kind)
)

func init() {
	SchemeBuilder.Register(&Pool{}, &PoolList{})
}
