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

type MethodResponseObservation struct {

	// HTTP Method (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY)
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// API resource ID
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Map of the API models used for the response's content type
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// Map of response parameters that can be sent to the caller.
	// For example: response_parameters = { "method.response.header.X-Some-Header" = true }
	// would define that the header X-Some-Header can be provided on the response.
	ResponseParameters map[string]*bool `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// ID of the associated REST API
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// HTTP status code
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type MethodResponseParameters struct {

	// HTTP Method (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY)
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Method
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("http_method",false)
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// Reference to a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodRef *v1.Reference `json:"httpMethodRef,omitempty" tf:"-"`

	// Selector for a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodSelector *v1.Selector `json:"httpMethodSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// API resource ID
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// Map of the API models used for the response's content type
	// +kubebuilder:validation:Optional
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// Map of response parameters that can be sent to the caller.
	// For example: response_parameters = { "method.response.header.X-Some-Header" = true }
	// would define that the header X-Some-Header can be provided on the response.
	// +kubebuilder:validation:Optional
	ResponseParameters map[string]*bool `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

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

	// HTTP status code
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

// MethodResponseSpec defines the desired state of MethodResponse
type MethodResponseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MethodResponseParameters `json:"forProvider"`
}

// MethodResponseStatus defines the observed state of MethodResponse.
type MethodResponseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MethodResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MethodResponse is the Schema for the MethodResponses API. Provides an HTTP Method Response for an API Gateway Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MethodResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.statusCode)",message="statusCode is a required parameter"
	Spec   MethodResponseSpec   `json:"spec"`
	Status MethodResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MethodResponseList contains a list of MethodResponses
type MethodResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MethodResponse `json:"items"`
}

// Repository type metadata.
var (
	MethodResponse_Kind             = "MethodResponse"
	MethodResponse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MethodResponse_Kind}.String()
	MethodResponse_KindAPIVersion   = MethodResponse_Kind + "." + CRDGroupVersion.String()
	MethodResponse_GroupVersionKind = CRDGroupVersion.WithKind(MethodResponse_Kind)
)

func init() {
	SchemeBuilder.Register(&MethodResponse{}, &MethodResponseList{})
}
