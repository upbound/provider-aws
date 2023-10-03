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

type DocumentationPartInitParameters struct {

	// Location of the targeted API entity of the to-be-created documentation part. See below.
	Location []LocationInitParameters `json:"location,omitempty" tf:"location,omitempty"`

	// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ "description": "The API does ..." }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties *string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type DocumentationPartObservation struct {

	// Unique ID of the Documentation Part
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Location of the targeted API entity of the to-be-created documentation part. See below.
	Location []LocationObservation `json:"location,omitempty" tf:"location,omitempty"`

	// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ "description": "The API does ..." }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties *string `json:"properties,omitempty" tf:"properties,omitempty"`

	// ID of the associated Rest API
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`
}

type DocumentationPartParameters struct {

	// Location of the targeted API entity of the to-be-created documentation part. See below.
	// +kubebuilder:validation:Optional
	Location []LocationParameters `json:"location,omitempty" tf:"location,omitempty"`

	// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ "description": "The API does ..." }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	// +kubebuilder:validation:Optional
	Properties *string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the associated Rest API
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
}

type LocationInitParameters struct {

	// HTTP verb of a method. The default value is * for any method.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Name of the targeted API entity.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// URL path of the target. The default value is / for the root resource.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// HTTP status code of a response. The default value is * for any status code.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// Type of API entity to which the documentation content appliesE.g., API, METHOD or REQUEST_BODY
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LocationObservation struct {

	// HTTP verb of a method. The default value is * for any method.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Name of the targeted API entity.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// URL path of the target. The default value is / for the root resource.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// HTTP status code of a response. The default value is * for any status code.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// Type of API entity to which the documentation content appliesE.g., API, METHOD or REQUEST_BODY
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LocationParameters struct {

	// HTTP verb of a method. The default value is * for any method.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Name of the targeted API entity.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// URL path of the target. The default value is / for the root resource.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// HTTP status code of a response. The default value is * for any status code.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// Type of API entity to which the documentation content appliesE.g., API, METHOD or REQUEST_BODY
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// DocumentationPartSpec defines the desired state of DocumentationPart
type DocumentationPartSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DocumentationPartParameters `json:"forProvider"`
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
	InitProvider DocumentationPartInitParameters `json:"initProvider,omitempty"`
}

// DocumentationPartStatus defines the observed state of DocumentationPart.
type DocumentationPartStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DocumentationPartObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationPart is the Schema for the DocumentationParts API. Provides a settings of an API Gateway Documentation Part.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DocumentationPart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.properties) || (has(self.initProvider) && has(self.initProvider.properties))",message="spec.forProvider.properties is a required parameter"
	Spec   DocumentationPartSpec   `json:"spec"`
	Status DocumentationPartStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationPartList contains a list of DocumentationParts
type DocumentationPartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocumentationPart `json:"items"`
}

// Repository type metadata.
var (
	DocumentationPart_Kind             = "DocumentationPart"
	DocumentationPart_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DocumentationPart_Kind}.String()
	DocumentationPart_KindAPIVersion   = DocumentationPart_Kind + "." + CRDGroupVersion.String()
	DocumentationPart_GroupVersionKind = CRDGroupVersion.WithKind(DocumentationPart_Kind)
)

func init() {
	SchemeBuilder.Register(&DocumentationPart{}, &DocumentationPartList{})
}
