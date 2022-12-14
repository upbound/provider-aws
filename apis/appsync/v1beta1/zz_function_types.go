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

type FunctionObservation struct {

	// The ARN of the Function object.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A unique ID representing the Function object.
	FunctionID *string `json:"functionId,omitempty" tf:"function_id,omitempty"`

	// API Function ID (Formatted as ApiId-FunctionId)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FunctionParameters struct {

	// The ID of the associated AppSync API.
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

	// The Function DataSource name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta1.Datasource
	// +kubebuilder:validation:Optional
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Reference to a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceRef *v1.Reference `json:"dataSourceRef,omitempty" tf:"-"`

	// Selector for a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceSelector *v1.Selector `json:"dataSourceSelector,omitempty" tf:"-"`

	// The Function description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The version of the request mapping template. Currently the supported value is 2018-05-29.
	// +kubebuilder:validation:Optional
	FunctionVersion *string `json:"functionVersion,omitempty" tf:"function_version,omitempty"`

	// The maximum batching size for a resolver. Valid values are between 0 and 2000.
	// +kubebuilder:validation:Optional
	MaxBatchSize *float64 `json:"maxBatchSize,omitempty" tf:"max_batch_size,omitempty"`

	// The Function name. The function name does not have to be unique.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	// +kubebuilder:validation:Required
	RequestMappingTemplate *string `json:"requestMappingTemplate" tf:"request_mapping_template,omitempty"`

	// The Function response mapping template.
	// +kubebuilder:validation:Required
	ResponseMappingTemplate *string `json:"responseMappingTemplate" tf:"response_mapping_template,omitempty"`

	// Describes a Sync configuration for a resolver. See Sync Config.
	// +kubebuilder:validation:Optional
	SyncConfig []SyncConfigParameters `json:"syncConfig,omitempty" tf:"sync_config,omitempty"`
}

type LambdaConflictHandlerConfigObservation struct {
}

type LambdaConflictHandlerConfigParameters struct {

	// The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.
	// +kubebuilder:validation:Optional
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn,omitempty" tf:"lambda_conflict_handler_arn,omitempty"`
}

type SyncConfigObservation struct {
}

type SyncConfigParameters struct {

	// The Conflict Detection strategy to use. Valid values are NONE and VERSION.
	// +kubebuilder:validation:Optional
	ConflictDetection *string `json:"conflictDetection,omitempty" tf:"conflict_detection,omitempty"`

	// The Conflict Resolution strategy to perform in the event of a conflict. Valid values are NONE, OPTIMISTIC_CONCURRENCY, AUTOMERGE, and LAMBDA.
	// +kubebuilder:validation:Optional
	ConflictHandler *string `json:"conflictHandler,omitempty" tf:"conflict_handler,omitempty"`

	// The Lambda Conflict Handler Config when configuring LAMBDA as the Conflict Handler. See Lambda Conflict Handler Config.
	// +kubebuilder:validation:Optional
	LambdaConflictHandlerConfig []LambdaConflictHandlerConfigParameters `json:"lambdaConflictHandlerConfig,omitempty" tf:"lambda_conflict_handler_config,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API. Provides an AppSync Function.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
