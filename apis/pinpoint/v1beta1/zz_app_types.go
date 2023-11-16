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

type AppInitParameters struct {

	// Specifies settings for invoking an AWS Lambda function that customizes a segment for a campaign
	CampaignHook []CampaignHookInitParameters `json:"campaignHook,omitempty" tf:"campaign_hook,omitempty"`

	// The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
	Limits []LimitsInitParameters `json:"limits,omitempty" tf:"limits,omitempty"`

	// The application name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
	QuietTime []QuietTimeInitParameters `json:"quietTime,omitempty" tf:"quiet_time,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppObservation struct {

	// The Application ID of the Pinpoint App.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Amazon Resource Name (ARN) of the PinPoint Application
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies settings for invoking an AWS Lambda function that customizes a segment for a campaign
	CampaignHook []CampaignHookObservation `json:"campaignHook,omitempty" tf:"campaign_hook,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
	Limits []LimitsObservation `json:"limits,omitempty" tf:"limits,omitempty"`

	// The application name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
	QuietTime []QuietTimeObservation `json:"quietTime,omitempty" tf:"quiet_time,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AppParameters struct {

	// Specifies settings for invoking an AWS Lambda function that customizes a segment for a campaign
	// +kubebuilder:validation:Optional
	CampaignHook []CampaignHookParameters `json:"campaignHook,omitempty" tf:"campaign_hook,omitempty"`

	// The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
	// +kubebuilder:validation:Optional
	Limits []LimitsParameters `json:"limits,omitempty" tf:"limits,omitempty"`

	// The application name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
	// +kubebuilder:validation:Optional
	QuietTime []QuietTimeParameters `json:"quietTime,omitempty" tf:"quiet_time,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CampaignHookInitParameters struct {

	// Lambda function name or ARN to be called for delivery. Conflicts with web_url
	LambdaFunctionName *string `json:"lambdaFunctionName,omitempty" tf:"lambda_function_name,omitempty"`

	// What mode Lambda should be invoked in. Valid values for this parameter are DELIVERY, FILTER.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Web URL to call for hook. If the URL has authentication specified it will be added as authentication to the request. Conflicts with lambda_function_name
	WebURL *string `json:"webUrl,omitempty" tf:"web_url,omitempty"`
}

type CampaignHookObservation struct {

	// Lambda function name or ARN to be called for delivery. Conflicts with web_url
	LambdaFunctionName *string `json:"lambdaFunctionName,omitempty" tf:"lambda_function_name,omitempty"`

	// What mode Lambda should be invoked in. Valid values for this parameter are DELIVERY, FILTER.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Web URL to call for hook. If the URL has authentication specified it will be added as authentication to the request. Conflicts with lambda_function_name
	WebURL *string `json:"webUrl,omitempty" tf:"web_url,omitempty"`
}

type CampaignHookParameters struct {

	// Lambda function name or ARN to be called for delivery. Conflicts with web_url
	// +kubebuilder:validation:Optional
	LambdaFunctionName *string `json:"lambdaFunctionName,omitempty" tf:"lambda_function_name,omitempty"`

	// What mode Lambda should be invoked in. Valid values for this parameter are DELIVERY, FILTER.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Web URL to call for hook. If the URL has authentication specified it will be added as authentication to the request. Conflicts with lambda_function_name
	// +kubebuilder:validation:Optional
	WebURL *string `json:"webUrl,omitempty" tf:"web_url,omitempty"`
}

type LimitsInitParameters struct {

	// The maximum number of messages that the campaign can send daily.
	Daily *float64 `json:"daily,omitempty" tf:"daily,omitempty"`

	// The length of time (in seconds) that the campaign can run before it ends and message deliveries stop. This duration begins at the scheduled start time for the campaign. The minimum value is 60.
	MaximumDuration *float64 `json:"maximumDuration,omitempty" tf:"maximum_duration,omitempty"`

	// The number of messages that the campaign can send per second. The minimum value is 50, and the maximum is 20000.
	MessagesPerSecond *float64 `json:"messagesPerSecond,omitempty" tf:"messages_per_second,omitempty"`

	// The maximum total number of messages that the campaign can send.
	Total *float64 `json:"total,omitempty" tf:"total,omitempty"`
}

type LimitsObservation struct {

	// The maximum number of messages that the campaign can send daily.
	Daily *float64 `json:"daily,omitempty" tf:"daily,omitempty"`

	// The length of time (in seconds) that the campaign can run before it ends and message deliveries stop. This duration begins at the scheduled start time for the campaign. The minimum value is 60.
	MaximumDuration *float64 `json:"maximumDuration,omitempty" tf:"maximum_duration,omitempty"`

	// The number of messages that the campaign can send per second. The minimum value is 50, and the maximum is 20000.
	MessagesPerSecond *float64 `json:"messagesPerSecond,omitempty" tf:"messages_per_second,omitempty"`

	// The maximum total number of messages that the campaign can send.
	Total *float64 `json:"total,omitempty" tf:"total,omitempty"`
}

type LimitsParameters struct {

	// The maximum number of messages that the campaign can send daily.
	// +kubebuilder:validation:Optional
	Daily *float64 `json:"daily,omitempty" tf:"daily,omitempty"`

	// The length of time (in seconds) that the campaign can run before it ends and message deliveries stop. This duration begins at the scheduled start time for the campaign. The minimum value is 60.
	// +kubebuilder:validation:Optional
	MaximumDuration *float64 `json:"maximumDuration,omitempty" tf:"maximum_duration,omitempty"`

	// The number of messages that the campaign can send per second. The minimum value is 50, and the maximum is 20000.
	// +kubebuilder:validation:Optional
	MessagesPerSecond *float64 `json:"messagesPerSecond,omitempty" tf:"messages_per_second,omitempty"`

	// The maximum total number of messages that the campaign can send.
	// +kubebuilder:validation:Optional
	Total *float64 `json:"total,omitempty" tf:"total,omitempty"`
}

type QuietTimeInitParameters struct {

	// The default end time for quiet time in ISO 8601 format. Required if start is set
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The default start time for quiet time in ISO 8601 format. Required if end is set
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type QuietTimeObservation struct {

	// The default end time for quiet time in ISO 8601 format. Required if start is set
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The default start time for quiet time in ISO 8601 format. Required if end is set
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type QuietTimeParameters struct {

	// The default end time for quiet time in ISO 8601 format. Required if start is set
	// +kubebuilder:validation:Optional
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The default start time for quiet time in ISO 8601 format. Required if end is set
	// +kubebuilder:validation:Optional
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

// AppSpec defines the desired state of App
type AppSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppParameters `json:"forProvider"`
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
	InitProvider AppInitParameters `json:"initProvider,omitempty"`
}

// AppStatus defines the observed state of App.
type AppStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// App is the Schema for the Apps API. Provides a Pinpoint App resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppSpec   `json:"spec"`
	Status            AppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppList contains a list of Apps
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

// Repository type metadata.
var (
	App_Kind             = "App"
	App_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: App_Kind}.String()
	App_KindAPIVersion   = App_Kind + "." + CRDGroupVersion.String()
	App_GroupVersionKind = CRDGroupVersion.WithKind(App_Kind)
)

func init() {
	SchemeBuilder.Register(&App{}, &AppList{})
}
