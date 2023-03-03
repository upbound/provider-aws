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

type AuthorizerObservation struct {

	// ARN of the API Gateway Authorizer
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Credentials required for the authorizer. To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials *string `json:"authorizerCredentials,omitempty" tf:"authorizer_credentials,omitempty"`

	// TTL of cached authorizer results in seconds. Defaults to 300.
	AuthorizerResultTTLInSeconds *float64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// Authorizer's Uniform Resource Identifier (URI). This must be a well-formed Lambda function URI in the form of arn:aws:apigateway:{region}:lambda:path/{service_api},
	// e.g., arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// Authorizer identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Source of the identity in an incoming request. Defaults to method.request.header.Authorization. For REQUEST type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g., "method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"
	IdentitySource *string `json:"identitySource,omitempty" tf:"identity_source,omitempty"`

	// Validation expression for the incoming identity. For TOKEN type, this value should be a regular expression. The incoming token from the client is matched against this expression, and will proceed if the token matches. If the token doesn't match, the client receives a 401 Unauthorized response.
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`

	// Name of the authorizer
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of the Amazon Cognito user pool ARNs. Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}.
	ProviderArns []*string `json:"providerArns,omitempty" tf:"provider_arns,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// ID of the associated REST API
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Type of the authorizer. Possible values are TOKEN for a Lambda function using a single authorization token submitted in a custom header, REQUEST for a Lambda function using incoming request parameters, or COGNITO_USER_POOLS for using an Amazon Cognito user pool. Defaults to TOKEN.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuthorizerParameters struct {

	// Credentials required for the authorizer. To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	AuthorizerCredentials *string `json:"authorizerCredentials,omitempty" tf:"authorizer_credentials,omitempty"`

	// Reference to a Role in iam to populate authorizerCredentials.
	// +kubebuilder:validation:Optional
	AuthorizerCredentialsRef *v1.Reference `json:"authorizerCredentialsRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate authorizerCredentials.
	// +kubebuilder:validation:Optional
	AuthorizerCredentialsSelector *v1.Selector `json:"authorizerCredentialsSelector,omitempty" tf:"-"`

	// TTL of cached authorizer results in seconds. Defaults to 300.
	// +kubebuilder:validation:Optional
	AuthorizerResultTTLInSeconds *float64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// Authorizer's Uniform Resource Identifier (URI). This must be a well-formed Lambda function URI in the form of arn:aws:apigateway:{region}:lambda:path/{service_api},
	// e.g., arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("invoke_arn",true)
	// +kubebuilder:validation:Optional
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// Reference to a Function in lambda to populate authorizerUri.
	// +kubebuilder:validation:Optional
	AuthorizerURIRef *v1.Reference `json:"authorizerUriRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate authorizerUri.
	// +kubebuilder:validation:Optional
	AuthorizerURISelector *v1.Selector `json:"authorizerUriSelector,omitempty" tf:"-"`

	// Source of the identity in an incoming request. Defaults to method.request.header.Authorization. For REQUEST type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g., "method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"
	// +kubebuilder:validation:Optional
	IdentitySource *string `json:"identitySource,omitempty" tf:"identity_source,omitempty"`

	// Validation expression for the incoming identity. For TOKEN type, this value should be a regular expression. The incoming token from the client is matched against this expression, and will proceed if the token matches. If the token doesn't match, the client receives a 401 Unauthorized response.
	// +kubebuilder:validation:Optional
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`

	// Name of the authorizer
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of the Amazon Cognito user pool ARNs. Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}.
	// +kubebuilder:validation:Optional
	ProviderArns []*string `json:"providerArns,omitempty" tf:"provider_arns,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the associated REST API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// Type of the authorizer. Possible values are TOKEN for a Lambda function using a single authorization token submitted in a custom header, REQUEST for a Lambda function using incoming request parameters, or COGNITO_USER_POOLS for using an Amazon Cognito user pool. Defaults to TOKEN.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// AuthorizerSpec defines the desired state of Authorizer
type AuthorizerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthorizerParameters `json:"forProvider"`
}

// AuthorizerStatus defines the observed state of Authorizer.
type AuthorizerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthorizerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Authorizer is the Schema for the Authorizers API. Provides an API Gateway Authorizer.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Authorizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   AuthorizerSpec   `json:"spec"`
	Status AuthorizerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthorizerList contains a list of Authorizers
type AuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Authorizer `json:"items"`
}

// Repository type metadata.
var (
	Authorizer_Kind             = "Authorizer"
	Authorizer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Authorizer_Kind}.String()
	Authorizer_KindAPIVersion   = Authorizer_Kind + "." + CRDGroupVersion.String()
	Authorizer_GroupVersionKind = CRDGroupVersion.WithKind(Authorizer_Kind)
)

func init() {
	SchemeBuilder.Register(&Authorizer{}, &AuthorizerList{})
}
