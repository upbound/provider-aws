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

type AppMonitorConfigurationInitParameters struct {

	// If you set this to true, RUM web client sets two cookies, a session cookie  and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
	AllowCookies *bool `json:"allowCookies,omitempty" tf:"allow_cookies,omitempty"`

	// If you set this to true, RUM enables X-Ray tracing for the user sessions  that RUM samples. RUM adds an X-Ray trace header to allowed HTTP requests. It also records an X-Ray segment for allowed HTTP requests.
	EnableXray *bool `json:"enableXray,omitempty" tf:"enable_xray,omitempty"`

	// A list of URLs in your website or application to exclude from RUM data collection.
	ExcludedPages []*string `json:"excludedPages,omitempty" tf:"excluded_pages,omitempty"`

	// A list of pages in the CloudWatch RUM console that are to be displayed with a "favorite" icon.
	FavoritePages []*string `json:"favoritePages,omitempty" tf:"favorite_pages,omitempty"`

	// The ARN of the guest IAM role that is attached to the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
	GuestRoleArn *string `json:"guestRoleArn,omitempty" tf:"guest_role_arn,omitempty"`

	// The ID of the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
	IdentityPoolID *string `json:"identityPoolId,omitempty" tf:"identity_pool_id,omitempty"`

	// If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.
	IncludedPages []*string `json:"includedPages,omitempty" tf:"included_pages,omitempty"`

	// Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. Default value is 0.1.
	SessionSampleRate *float64 `json:"sessionSampleRate,omitempty" tf:"session_sample_rate,omitempty"`

	// An array that lists the types of telemetry data that this app monitor is to collect. Valid values are errors, performance, and http.
	Telemetries []*string `json:"telemetries,omitempty" tf:"telemetries,omitempty"`
}

type AppMonitorConfigurationObservation struct {

	// If you set this to true, RUM web client sets two cookies, a session cookie  and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
	AllowCookies *bool `json:"allowCookies,omitempty" tf:"allow_cookies,omitempty"`

	// If you set this to true, RUM enables X-Ray tracing for the user sessions  that RUM samples. RUM adds an X-Ray trace header to allowed HTTP requests. It also records an X-Ray segment for allowed HTTP requests.
	EnableXray *bool `json:"enableXray,omitempty" tf:"enable_xray,omitempty"`

	// A list of URLs in your website or application to exclude from RUM data collection.
	ExcludedPages []*string `json:"excludedPages,omitempty" tf:"excluded_pages,omitempty"`

	// A list of pages in the CloudWatch RUM console that are to be displayed with a "favorite" icon.
	FavoritePages []*string `json:"favoritePages,omitempty" tf:"favorite_pages,omitempty"`

	// The ARN of the guest IAM role that is attached to the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
	GuestRoleArn *string `json:"guestRoleArn,omitempty" tf:"guest_role_arn,omitempty"`

	// The ID of the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
	IdentityPoolID *string `json:"identityPoolId,omitempty" tf:"identity_pool_id,omitempty"`

	// If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.
	IncludedPages []*string `json:"includedPages,omitempty" tf:"included_pages,omitempty"`

	// Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. Default value is 0.1.
	SessionSampleRate *float64 `json:"sessionSampleRate,omitempty" tf:"session_sample_rate,omitempty"`

	// An array that lists the types of telemetry data that this app monitor is to collect. Valid values are errors, performance, and http.
	Telemetries []*string `json:"telemetries,omitempty" tf:"telemetries,omitempty"`
}

type AppMonitorConfigurationParameters struct {

	// If you set this to true, RUM web client sets two cookies, a session cookie  and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
	// +kubebuilder:validation:Optional
	AllowCookies *bool `json:"allowCookies,omitempty" tf:"allow_cookies,omitempty"`

	// If you set this to true, RUM enables X-Ray tracing for the user sessions  that RUM samples. RUM adds an X-Ray trace header to allowed HTTP requests. It also records an X-Ray segment for allowed HTTP requests.
	// +kubebuilder:validation:Optional
	EnableXray *bool `json:"enableXray,omitempty" tf:"enable_xray,omitempty"`

	// A list of URLs in your website or application to exclude from RUM data collection.
	// +kubebuilder:validation:Optional
	ExcludedPages []*string `json:"excludedPages,omitempty" tf:"excluded_pages,omitempty"`

	// A list of pages in the CloudWatch RUM console that are to be displayed with a "favorite" icon.
	// +kubebuilder:validation:Optional
	FavoritePages []*string `json:"favoritePages,omitempty" tf:"favorite_pages,omitempty"`

	// The ARN of the guest IAM role that is attached to the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
	// +kubebuilder:validation:Optional
	GuestRoleArn *string `json:"guestRoleArn,omitempty" tf:"guest_role_arn,omitempty"`

	// The ID of the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
	// +kubebuilder:validation:Optional
	IdentityPoolID *string `json:"identityPoolId,omitempty" tf:"identity_pool_id,omitempty"`

	// If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.
	// +kubebuilder:validation:Optional
	IncludedPages []*string `json:"includedPages,omitempty" tf:"included_pages,omitempty"`

	// Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. Default value is 0.1.
	// +kubebuilder:validation:Optional
	SessionSampleRate *float64 `json:"sessionSampleRate,omitempty" tf:"session_sample_rate,omitempty"`

	// An array that lists the types of telemetry data that this app monitor is to collect. Valid values are errors, performance, and http.
	// +kubebuilder:validation:Optional
	Telemetries []*string `json:"telemetries,omitempty" tf:"telemetries,omitempty"`
}

type AppMonitorInitParameters struct {

	// configuration data for the app monitor. See app_monitor_configuration below.
	AppMonitorConfiguration []AppMonitorConfigurationInitParameters `json:"appMonitorConfiguration,omitempty" tf:"app_monitor_configuration,omitempty"`

	// Specifies whether this app monitor allows the web client to define and send custom events. If you omit this parameter, custom events are DISABLED. See custom_events below.
	CustomEvents []CustomEventsInitParameters `json:"customEvents,omitempty" tf:"custom_events,omitempty"`

	// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter  specifies whether RUM sends a copy of this telemetry data to Amazon CloudWatch Logs in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur Amazon CloudWatch Logs charges. Default value is false.
	CwLogEnabled *bool `json:"cwLogEnabled,omitempty" tf:"cw_log_enabled,omitempty"`

	// The top-level internet domain name for which your application has administrative authority.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppMonitorObservation struct {

	// configuration data for the app monitor. See app_monitor_configuration below.
	AppMonitorConfiguration []AppMonitorConfigurationObservation `json:"appMonitorConfiguration,omitempty" tf:"app_monitor_configuration,omitempty"`

	// The unique ID of the app monitor. Useful for JS templates.
	AppMonitorID *string `json:"appMonitorId,omitempty" tf:"app_monitor_id,omitempty"`

	// The Amazon Resource Name (ARN) specifying the app monitor.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies whether this app monitor allows the web client to define and send custom events. If you omit this parameter, custom events are DISABLED. See custom_events below.
	CustomEvents []CustomEventsObservation `json:"customEvents,omitempty" tf:"custom_events,omitempty"`

	// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter  specifies whether RUM sends a copy of this telemetry data to Amazon CloudWatch Logs in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur Amazon CloudWatch Logs charges. Default value is false.
	CwLogEnabled *bool `json:"cwLogEnabled,omitempty" tf:"cw_log_enabled,omitempty"`

	// The name of the log group where the copies are stored.
	CwLogGroup *string `json:"cwLogGroup,omitempty" tf:"cw_log_group,omitempty"`

	// The top-level internet domain name for which your application has administrative authority.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The CloudWatch RUM name as it is the identifier of a RUM.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AppMonitorParameters struct {

	// configuration data for the app monitor. See app_monitor_configuration below.
	// +kubebuilder:validation:Optional
	AppMonitorConfiguration []AppMonitorConfigurationParameters `json:"appMonitorConfiguration,omitempty" tf:"app_monitor_configuration,omitempty"`

	// Specifies whether this app monitor allows the web client to define and send custom events. If you omit this parameter, custom events are DISABLED. See custom_events below.
	// +kubebuilder:validation:Optional
	CustomEvents []CustomEventsParameters `json:"customEvents,omitempty" tf:"custom_events,omitempty"`

	// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter  specifies whether RUM sends a copy of this telemetry data to Amazon CloudWatch Logs in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur Amazon CloudWatch Logs charges. Default value is false.
	// +kubebuilder:validation:Optional
	CwLogEnabled *bool `json:"cwLogEnabled,omitempty" tf:"cw_log_enabled,omitempty"`

	// The top-level internet domain name for which your application has administrative authority.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CustomEventsInitParameters struct {

	// Specifies whether this app monitor allows the web client to define and send custom events. The default is for custom events to be DISABLED. Valid values are DISABLED and ENABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type CustomEventsObservation struct {

	// Specifies whether this app monitor allows the web client to define and send custom events. The default is for custom events to be DISABLED. Valid values are DISABLED and ENABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type CustomEventsParameters struct {

	// Specifies whether this app monitor allows the web client to define and send custom events. The default is for custom events to be DISABLED. Valid values are DISABLED and ENABLED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// AppMonitorSpec defines the desired state of AppMonitor
type AppMonitorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppMonitorParameters `json:"forProvider"`
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
	InitProvider AppMonitorInitParameters `json:"initProvider,omitempty"`
}

// AppMonitorStatus defines the observed state of AppMonitor.
type AppMonitorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppMonitorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppMonitor is the Schema for the AppMonitors API. Provides a CloudWatch RUM App Monitor resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AppMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	Spec   AppMonitorSpec   `json:"spec"`
	Status AppMonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppMonitorList contains a list of AppMonitors
type AppMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppMonitor `json:"items"`
}

// Repository type metadata.
var (
	AppMonitor_Kind             = "AppMonitor"
	AppMonitor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppMonitor_Kind}.String()
	AppMonitor_KindAPIVersion   = AppMonitor_Kind + "." + CRDGroupVersion.String()
	AppMonitor_GroupVersionKind = CRDGroupVersion.WithKind(AppMonitor_Kind)
)

func init() {
	SchemeBuilder.Register(&AppMonitor{}, &AppMonitorList{})
}
