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

type AuthenticationConfigurationInitParameters struct {

	// A valid CIDR block for IP filtering. Required for IP.
	AllowedIPRange *string `json:"allowedIpRange,omitempty" tf:"allowed_ip_range,omitempty"`
}

type AuthenticationConfigurationObservation struct {

	// A valid CIDR block for IP filtering. Required for IP.
	AllowedIPRange *string `json:"allowedIpRange,omitempty" tf:"allowed_ip_range,omitempty"`
}

type AuthenticationConfigurationParameters struct {

	// A valid CIDR block for IP filtering. Required for IP.
	// +kubebuilder:validation:Optional
	AllowedIPRange *string `json:"allowedIpRange,omitempty" tf:"allowed_ip_range,omitempty"`

	// The shared secret for the GitHub repository webhook. Set this as secret in your github_repository_webhook's configuration block. Required for GITHUB_HMAC.
	// +kubebuilder:validation:Optional
	SecretTokenSecretRef *v1.SecretKeySelector `json:"secretTokenSecretRef,omitempty" tf:"-"`
}

type FilterInitParameters struct {

	// The JSON path to filter on.
	JSONPath *string `json:"jsonPath,omitempty" tf:"json_path,omitempty"`

	// The value to match on (e.g., refs/heads/{Branch}). See AWS docs for details.
	MatchEquals *string `json:"matchEquals,omitempty" tf:"match_equals,omitempty"`
}

type FilterObservation struct {

	// The JSON path to filter on.
	JSONPath *string `json:"jsonPath,omitempty" tf:"json_path,omitempty"`

	// The value to match on (e.g., refs/heads/{Branch}). See AWS docs for details.
	MatchEquals *string `json:"matchEquals,omitempty" tf:"match_equals,omitempty"`
}

type FilterParameters struct {

	// The JSON path to filter on.
	// +kubebuilder:validation:Optional
	JSONPath *string `json:"jsonPath" tf:"json_path,omitempty"`

	// The value to match on (e.g., refs/heads/{Branch}). See AWS docs for details.
	// +kubebuilder:validation:Optional
	MatchEquals *string `json:"matchEquals" tf:"match_equals,omitempty"`
}

type WebhookInitParameters struct {

	// The type of authentication  to use. One of IP, GITHUB_HMAC, or UNAUTHENTICATED.
	Authentication *string `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// An auth block. Required for IP and GITHUB_HMAC. Auth blocks are documented below.
	AuthenticationConfiguration []AuthenticationConfigurationInitParameters `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`

	// One or more filter blocks. Filter blocks are documented below.
	Filter []FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction *string `json:"targetAction,omitempty" tf:"target_action,omitempty"`
}

type WebhookObservation struct {

	// The CodePipeline webhook's ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The type of authentication  to use. One of IP, GITHUB_HMAC, or UNAUTHENTICATED.
	Authentication *string `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// An auth block. Required for IP and GITHUB_HMAC. Auth blocks are documented below.
	AuthenticationConfiguration []AuthenticationConfigurationObservation `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`

	// One or more filter blocks. Filter blocks are documented below.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// The CodePipeline webhook's ARN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction *string `json:"targetAction,omitempty" tf:"target_action,omitempty"`

	// The name of the pipeline.
	TargetPipeline *string `json:"targetPipeline,omitempty" tf:"target_pipeline,omitempty"`

	// The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type WebhookParameters struct {

	// The type of authentication  to use. One of IP, GITHUB_HMAC, or UNAUTHENTICATED.
	// +kubebuilder:validation:Optional
	Authentication *string `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// An auth block. Required for IP and GITHUB_HMAC. Auth blocks are documented below.
	// +kubebuilder:validation:Optional
	AuthenticationConfiguration []AuthenticationConfigurationParameters `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`

	// One or more filter blocks. Filter blocks are documented below.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	// +kubebuilder:validation:Optional
	TargetAction *string `json:"targetAction,omitempty" tf:"target_action,omitempty"`

	// The name of the pipeline.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codepipeline/v1beta1.Codepipeline
	// +kubebuilder:validation:Optional
	TargetPipeline *string `json:"targetPipeline,omitempty" tf:"target_pipeline,omitempty"`

	// Reference to a Codepipeline in codepipeline to populate targetPipeline.
	// +kubebuilder:validation:Optional
	TargetPipelineRef *v1.Reference `json:"targetPipelineRef,omitempty" tf:"-"`

	// Selector for a Codepipeline in codepipeline to populate targetPipeline.
	// +kubebuilder:validation:Optional
	TargetPipelineSelector *v1.Selector `json:"targetPipelineSelector,omitempty" tf:"-"`
}

// WebhookSpec defines the desired state of Webhook
type WebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebhookParameters `json:"forProvider"`
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
	InitProvider WebhookInitParameters `json:"initProvider,omitempty"`
}

// WebhookStatus defines the observed state of Webhook.
type WebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Webhook is the Schema for the Webhooks API. Provides a CodePipeline Webhook
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Webhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authentication) || (has(self.initProvider) && has(self.initProvider.authentication))",message="spec.forProvider.authentication is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filter) || (has(self.initProvider) && has(self.initProvider.filter))",message="spec.forProvider.filter is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetAction) || (has(self.initProvider) && has(self.initProvider.targetAction))",message="spec.forProvider.targetAction is a required parameter"
	Spec   WebhookSpec   `json:"spec"`
	Status WebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebhookList contains a list of Webhooks
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webhook `json:"items"`
}

// Repository type metadata.
var (
	Webhook_Kind             = "Webhook"
	Webhook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Webhook_Kind}.String()
	Webhook_KindAPIVersion   = Webhook_Kind + "." + CRDGroupVersion.String()
	Webhook_GroupVersionKind = CRDGroupVersion.WithKind(Webhook_Kind)
)

func init() {
	SchemeBuilder.Register(&Webhook{}, &WebhookList{})
}
