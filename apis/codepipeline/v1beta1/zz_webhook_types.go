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

type AuthenticationConfigurationObservation struct {

	// A valid CIDR block for IP filtering. Required for IP.
	AllowedIPRange *string `json:"allowedIpRange,omitempty" tf:"allowed_ip_range,omitempty"`

	// The shared secret for the GitHub repository webhook. Set this as secret in your github_repository_webhook's configuration block. Required for GITHUB_HMAC.
	SecretTokenSecretRef *v1.SecretKeySelector `json:"secretTokenSecretRef,omitempty" tf:"-"`
}

type AuthenticationConfigurationParameters struct {

	// A valid CIDR block for IP filtering. Required for IP.
	// +kubebuilder:validation:Optional
	AllowedIPRange *string `json:"allowedIpRange,omitempty" tf:"allowed_ip_range,omitempty"`

	// The shared secret for the GitHub repository webhook. Set this as secret in your github_repository_webhook's configuration block. Required for GITHUB_HMAC.
	// +kubebuilder:validation:Optional
	SecretTokenSecretRef *v1.SecretKeySelector `json:"secretTokenSecretRef,omitempty" tf:"-"`
}

type FilterObservation struct {

	// The JSON path to filter on.
	JSONPath *string `json:"jsonPath,omitempty" tf:"json_path,omitempty"`

	// The value to match on (e.g., refs/heads/{Branch}). See AWS docs for details.
	MatchEquals *string `json:"matchEquals,omitempty" tf:"match_equals,omitempty"`
}

type FilterParameters struct {

	// The JSON path to filter on.
	// +kubebuilder:validation:Required
	JSONPath *string `json:"jsonPath" tf:"json_path,omitempty"`

	// The value to match on (e.g., refs/heads/{Branch}). See AWS docs for details.
	// +kubebuilder:validation:Required
	MatchEquals *string `json:"matchEquals" tf:"match_equals,omitempty"`
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

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

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
	// +kubebuilder:validation:Required
	Authentication *string `json:"authentication" tf:"authentication,omitempty"`

	// An auth block. Required for IP and GITHUB_HMAC. Auth blocks are documented below.
	// +kubebuilder:validation:Optional
	AuthenticationConfiguration []AuthenticationConfigurationParameters `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`

	// One or more filter blocks. Filter blocks are documented below.
	// +kubebuilder:validation:Required
	Filter []FilterParameters `json:"filter" tf:"filter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	// +kubebuilder:validation:Required
	TargetAction *string `json:"targetAction" tf:"target_action,omitempty"`

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
	Spec              WebhookSpec   `json:"spec"`
	Status            WebhookStatus `json:"status,omitempty"`
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
