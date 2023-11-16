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

type BranchInitParameters struct {

	// ARN for a backend environment that is part of an Amplify app.
	BackendEnvironmentArn *string `json:"backendEnvironmentArn,omitempty" tf:"backend_environment_arn,omitempty"`

	// Description for the branch.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name for a branch. This is used as the default domain prefix.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Enables auto building for the branch.
	EnableAutoBuild *bool `json:"enableAutoBuild,omitempty" tf:"enable_auto_build,omitempty"`

	// Enables basic authorization for the branch.
	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty" tf:"enable_basic_auth,omitempty"`

	// Enables notifications for the branch.
	EnableNotification *bool `json:"enableNotification,omitempty" tf:"enable_notification,omitempty"`

	// Enables performance mode for the branch.
	EnablePerformanceMode *bool `json:"enablePerformanceMode,omitempty" tf:"enable_performance_mode,omitempty"`

	// Enables pull request previews for this branch.
	EnablePullRequestPreview *bool `json:"enablePullRequestPreview,omitempty" tf:"enable_pull_request_preview,omitempty"`

	// Environment variables for the branch.
	// +mapType:granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Framework for the branch.
	Framework *string `json:"framework,omitempty" tf:"framework,omitempty"`

	// Amplify environment name for the pull request.
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName,omitempty" tf:"pull_request_environment_name,omitempty"`

	// Describes the current stage for the branch. Valid values: PRODUCTION, BETA, DEVELOPMENT, EXPERIMENTAL, PULL_REQUEST.
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// Content Time To Live (TTL) for the website in seconds.
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BranchObservation struct {

	// Unique ID for an Amplify app.
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// ARN for the branch.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A list of custom resources that are linked to this branch.
	AssociatedResources []*string `json:"associatedResources,omitempty" tf:"associated_resources,omitempty"`

	// ARN for a backend environment that is part of an Amplify app.
	BackendEnvironmentArn *string `json:"backendEnvironmentArn,omitempty" tf:"backend_environment_arn,omitempty"`

	// Custom domains for the branch.
	CustomDomains []*string `json:"customDomains,omitempty" tf:"custom_domains,omitempty"`

	// Description for the branch.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Destination branch if the branch is a pull request branch.
	DestinationBranch *string `json:"destinationBranch,omitempty" tf:"destination_branch,omitempty"`

	// Display name for a branch. This is used as the default domain prefix.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Enables auto building for the branch.
	EnableAutoBuild *bool `json:"enableAutoBuild,omitempty" tf:"enable_auto_build,omitempty"`

	// Enables basic authorization for the branch.
	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty" tf:"enable_basic_auth,omitempty"`

	// Enables notifications for the branch.
	EnableNotification *bool `json:"enableNotification,omitempty" tf:"enable_notification,omitempty"`

	// Enables performance mode for the branch.
	EnablePerformanceMode *bool `json:"enablePerformanceMode,omitempty" tf:"enable_performance_mode,omitempty"`

	// Enables pull request previews for this branch.
	EnablePullRequestPreview *bool `json:"enablePullRequestPreview,omitempty" tf:"enable_pull_request_preview,omitempty"`

	// Environment variables for the branch.
	// +mapType:granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Framework for the branch.
	Framework *string `json:"framework,omitempty" tf:"framework,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amplify environment name for the pull request.
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName,omitempty" tf:"pull_request_environment_name,omitempty"`

	// Source branch if the branch is a pull request branch.
	SourceBranch *string `json:"sourceBranch,omitempty" tf:"source_branch,omitempty"`

	// Describes the current stage for the branch. Valid values: PRODUCTION, BETA, DEVELOPMENT, EXPERIMENTAL, PULL_REQUEST.
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// Content Time To Live (TTL) for the website in seconds.
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type BranchParameters struct {

	// Unique ID for an Amplify app.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/amplify/v1beta1.App
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// Reference to a App in amplify to populate appId.
	// +kubebuilder:validation:Optional
	AppIDRef *v1.Reference `json:"appIdRef,omitempty" tf:"-"`

	// Selector for a App in amplify to populate appId.
	// +kubebuilder:validation:Optional
	AppIDSelector *v1.Selector `json:"appIdSelector,omitempty" tf:"-"`

	// ARN for a backend environment that is part of an Amplify app.
	// +kubebuilder:validation:Optional
	BackendEnvironmentArn *string `json:"backendEnvironmentArn,omitempty" tf:"backend_environment_arn,omitempty"`

	// Basic authorization credentials for the branch.
	// +kubebuilder:validation:Optional
	BasicAuthCredentialsSecretRef *v1.SecretKeySelector `json:"basicAuthCredentialsSecretRef,omitempty" tf:"-"`

	// Description for the branch.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name for a branch. This is used as the default domain prefix.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Enables auto building for the branch.
	// +kubebuilder:validation:Optional
	EnableAutoBuild *bool `json:"enableAutoBuild,omitempty" tf:"enable_auto_build,omitempty"`

	// Enables basic authorization for the branch.
	// +kubebuilder:validation:Optional
	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty" tf:"enable_basic_auth,omitempty"`

	// Enables notifications for the branch.
	// +kubebuilder:validation:Optional
	EnableNotification *bool `json:"enableNotification,omitempty" tf:"enable_notification,omitempty"`

	// Enables performance mode for the branch.
	// +kubebuilder:validation:Optional
	EnablePerformanceMode *bool `json:"enablePerformanceMode,omitempty" tf:"enable_performance_mode,omitempty"`

	// Enables pull request previews for this branch.
	// +kubebuilder:validation:Optional
	EnablePullRequestPreview *bool `json:"enablePullRequestPreview,omitempty" tf:"enable_pull_request_preview,omitempty"`

	// Environment variables for the branch.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Framework for the branch.
	// +kubebuilder:validation:Optional
	Framework *string `json:"framework,omitempty" tf:"framework,omitempty"`

	// Amplify environment name for the pull request.
	// +kubebuilder:validation:Optional
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName,omitempty" tf:"pull_request_environment_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Describes the current stage for the branch. Valid values: PRODUCTION, BETA, DEVELOPMENT, EXPERIMENTAL, PULL_REQUEST.
	// +kubebuilder:validation:Optional
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// Content Time To Live (TTL) for the website in seconds.
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// BranchSpec defines the desired state of Branch
type BranchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BranchParameters `json:"forProvider"`
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
	InitProvider BranchInitParameters `json:"initProvider,omitempty"`
}

// BranchStatus defines the observed state of Branch.
type BranchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BranchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Branch is the Schema for the Branchs API. Provides an Amplify Branch resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Branch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BranchSpec   `json:"spec"`
	Status            BranchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BranchList contains a list of Branchs
type BranchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Branch `json:"items"`
}

// Repository type metadata.
var (
	Branch_Kind             = "Branch"
	Branch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Branch_Kind}.String()
	Branch_KindAPIVersion   = Branch_Kind + "." + CRDGroupVersion.String()
	Branch_GroupVersionKind = CRDGroupVersion.WithKind(Branch_Kind)
)

func init() {
	SchemeBuilder.Register(&Branch{}, &BranchList{})
}
