// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type RouteResponseInitParameters struct {

	// The model selection expression for the route response.
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression,omitempty"`

	// Response models for the route response.
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// Route response key.
	RouteResponseKey *string `json:"routeResponseKey,omitempty" tf:"route_response_key,omitempty"`
}

type RouteResponseObservation struct {

	// API identifier.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Route response identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The model selection expression for the route response.
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression,omitempty"`

	// Response models for the route response.
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// Identifier of the aws_apigatewayv2_route.
	RouteID *string `json:"routeId,omitempty" tf:"route_id,omitempty"`

	// Route response key.
	RouteResponseKey *string `json:"routeResponseKey,omitempty" tf:"route_response_key,omitempty"`
}

type RouteResponseParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=API
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// The model selection expression for the route response.
	// +kubebuilder:validation:Optional
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Response models for the route response.
	// +kubebuilder:validation:Optional
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// Identifier of the aws_apigatewayv2_route.
	// +crossplane:generate:reference:type=Route
	// +kubebuilder:validation:Optional
	RouteID *string `json:"routeId,omitempty" tf:"route_id,omitempty"`

	// Reference to a Route to populate routeId.
	// +kubebuilder:validation:Optional
	RouteIDRef *v1.Reference `json:"routeIdRef,omitempty" tf:"-"`

	// Selector for a Route to populate routeId.
	// +kubebuilder:validation:Optional
	RouteIDSelector *v1.Selector `json:"routeIdSelector,omitempty" tf:"-"`

	// Route response key.
	// +kubebuilder:validation:Optional
	RouteResponseKey *string `json:"routeResponseKey,omitempty" tf:"route_response_key,omitempty"`
}

// RouteResponseSpec defines the desired state of RouteResponse
type RouteResponseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteResponseParameters `json:"forProvider"`
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
	InitProvider RouteResponseInitParameters `json:"initProvider,omitempty"`
}

// RouteResponseStatus defines the observed state of RouteResponse.
type RouteResponseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteResponse is the Schema for the RouteResponses API. Manages an Amazon API Gateway Version 2 route response.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RouteResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routeResponseKey) || has(self.initProvider.routeResponseKey)",message="routeResponseKey is a required parameter"
	Spec   RouteResponseSpec   `json:"spec"`
	Status RouteResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteResponseList contains a list of RouteResponses
type RouteResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteResponse `json:"items"`
}

// Repository type metadata.
var (
	RouteResponse_Kind             = "RouteResponse"
	RouteResponse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteResponse_Kind}.String()
	RouteResponse_KindAPIVersion   = RouteResponse_Kind + "." + CRDGroupVersion.String()
	RouteResponse_GroupVersionKind = CRDGroupVersion.WithKind(RouteResponse_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteResponse{}, &RouteResponseList{})
}
