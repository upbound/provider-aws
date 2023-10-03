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

type LoggingConfigurationInitParameters struct {

	// The ARN of the CloudWatch log group to which the vended log data will be published. This log group must exist.
	LogGroupArn *string `json:"logGroupArn,omitempty" tf:"log_group_arn,omitempty"`
}

type LoggingConfigurationObservation struct {

	// The ARN of the CloudWatch log group to which the vended log data will be published. This log group must exist.
	LogGroupArn *string `json:"logGroupArn,omitempty" tf:"log_group_arn,omitempty"`
}

type LoggingConfigurationParameters struct {

	// The ARN of the CloudWatch log group to which the vended log data will be published. This log group must exist.
	// +kubebuilder:validation:Optional
	LogGroupArn *string `json:"logGroupArn" tf:"log_group_arn,omitempty"`
}

type WorkspaceInitParameters struct {

	// The alias of the prometheus workspace. See more in AWS Docs.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Logging configuration for the workspace. See Logging Configuration below for details.
	LoggingConfiguration []LoggingConfigurationInitParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WorkspaceObservation struct {

	// The alias of the prometheus workspace. See more in AWS Docs.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Amazon Resource Name (ARN) of the workspace.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Identifier of the workspace
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Logging configuration for the workspace. See Logging Configuration below for details.
	LoggingConfiguration []LoggingConfigurationObservation `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// Prometheus endpoint available for this workspace.
	PrometheusEndpoint *string `json:"prometheusEndpoint,omitempty" tf:"prometheus_endpoint,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkspaceParameters struct {

	// The alias of the prometheus workspace. See more in AWS Docs.
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Logging configuration for the workspace. See Logging Configuration below for details.
	// +kubebuilder:validation:Optional
	LoggingConfiguration []LoggingConfigurationParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
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
	InitProvider WorkspaceInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workspace is the Schema for the Workspaces API. Manages an Amazon Managed Service for Prometheus (AMP) Workspace
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Repository type metadata.
var (
	Workspace_Kind             = "Workspace"
	Workspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workspace_Kind}.String()
	Workspace_KindAPIVersion   = Workspace_Kind + "." + CRDGroupVersion.String()
	Workspace_GroupVersionKind = CRDGroupVersion.WithKind(Workspace_Kind)
)

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
