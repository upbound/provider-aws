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

type CachePolicyInitParameters struct {

	// Description for the cache policy.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Amount of time, in seconds, that objects are allowed to remain in the CloudFront cache before CloudFront sends a new request to the origin server to check if the object has been updated.
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// Maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Minimum amount of time, in seconds, that objects should remain in the CloudFront cache before a new request is sent to the origin to check for updates.
	MinTTL *float64 `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`

	// Unique name used to identify the cache policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for including HTTP headers, cookies, and URL query strings in the cache key. For more information, refer to the Parameters In Cache Key And Forwarded To Origin section.
	ParametersInCacheKeyAndForwardedToOrigin []ParametersInCacheKeyAndForwardedToOriginInitParameters `json:"parametersInCacheKeyAndForwardedToOrigin,omitempty" tf:"parameters_in_cache_key_and_forwarded_to_origin,omitempty"`
}

type CachePolicyObservation struct {

	// Description for the cache policy.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Amount of time, in seconds, that objects are allowed to remain in the CloudFront cache before CloudFront sends a new request to the origin server to check if the object has been updated.
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// Current version of the cache policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Identifier for the cache policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Minimum amount of time, in seconds, that objects should remain in the CloudFront cache before a new request is sent to the origin to check for updates.
	MinTTL *float64 `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`

	// Unique name used to identify the cache policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for including HTTP headers, cookies, and URL query strings in the cache key. For more information, refer to the Parameters In Cache Key And Forwarded To Origin section.
	ParametersInCacheKeyAndForwardedToOrigin []ParametersInCacheKeyAndForwardedToOriginObservation `json:"parametersInCacheKeyAndForwardedToOrigin,omitempty" tf:"parameters_in_cache_key_and_forwarded_to_origin,omitempty"`
}

type CachePolicyParameters struct {

	// Description for the cache policy.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Amount of time, in seconds, that objects are allowed to remain in the CloudFront cache before CloudFront sends a new request to the origin server to check if the object has been updated.
	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// Maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	// +kubebuilder:validation:Optional
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Minimum amount of time, in seconds, that objects should remain in the CloudFront cache before a new request is sent to the origin to check for updates.
	// +kubebuilder:validation:Optional
	MinTTL *float64 `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`

	// Unique name used to identify the cache policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for including HTTP headers, cookies, and URL query strings in the cache key. For more information, refer to the Parameters In Cache Key And Forwarded To Origin section.
	// +kubebuilder:validation:Optional
	ParametersInCacheKeyAndForwardedToOrigin []ParametersInCacheKeyAndForwardedToOriginParameters `json:"parametersInCacheKeyAndForwardedToOrigin,omitempty" tf:"parameters_in_cache_key_and_forwarded_to_origin,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type CookiesConfigInitParameters struct {

	// Whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for cookie_behavior are none, whitelist, allExcept, and all.
	CookieBehavior *string `json:"cookieBehavior,omitempty" tf:"cookie_behavior,omitempty"`

	// Object that contains a list of cookie names. See Items for more information.
	Cookies []CookiesInitParameters `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type CookiesConfigObservation struct {

	// Whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for cookie_behavior are none, whitelist, allExcept, and all.
	CookieBehavior *string `json:"cookieBehavior,omitempty" tf:"cookie_behavior,omitempty"`

	// Object that contains a list of cookie names. See Items for more information.
	Cookies []CookiesObservation `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type CookiesConfigParameters struct {

	// Whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for cookie_behavior are none, whitelist, allExcept, and all.
	// +kubebuilder:validation:Optional
	CookieBehavior *string `json:"cookieBehavior,omitempty" tf:"cookie_behavior,omitempty"`

	// Object that contains a list of cookie names. See Items for more information.
	// +kubebuilder:validation:Optional
	Cookies []CookiesParameters `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type CookiesInitParameters struct {

	// List of item names, such as cookies, headers, or query strings.
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type CookiesObservation struct {

	// List of item names, such as cookies, headers, or query strings.
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type CookiesParameters struct {

	// List of item names, such as cookies, headers, or query strings.
	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersConfigInitParameters struct {

	// Whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for header_behavior are none and whitelist.
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	// Object contains a list of header names. See Items for more information.
	Headers []HeadersInitParameters `json:"headers,omitempty" tf:"headers,omitempty"`
}

type HeadersConfigObservation struct {

	// Whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for header_behavior are none and whitelist.
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	// Object contains a list of header names. See Items for more information.
	Headers []HeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`
}

type HeadersConfigParameters struct {

	// Whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for header_behavior are none and whitelist.
	// +kubebuilder:validation:Optional
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	// Object contains a list of header names. See Items for more information.
	// +kubebuilder:validation:Optional
	Headers []HeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`
}

type HeadersInitParameters struct {

	// List of item names, such as cookies, headers, or query strings.
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersObservation struct {

	// List of item names, such as cookies, headers, or query strings.
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersParameters struct {

	// List of item names, such as cookies, headers, or query strings.
	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type ParametersInCacheKeyAndForwardedToOriginInitParameters struct {

	// Whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig []CookiesConfigInitParameters `json:"cookiesConfig,omitempty" tf:"cookies_config,omitempty"`

	// Flag determines whether the Accept-Encoding HTTP header is included in the cache key and in requests that CloudFront sends to the origin.
	EnableAcceptEncodingBrotli *bool `json:"enableAcceptEncodingBrotli,omitempty" tf:"enable_accept_encoding_brotli,omitempty"`

	// Whether the Accept-Encoding HTTP header is included in the cache key and in requests sent to the origin by CloudFront.
	EnableAcceptEncodingGzip *bool `json:"enableAcceptEncodingGzip,omitempty" tf:"enable_accept_encoding_gzip,omitempty"`

	// Whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig []HeadersConfigInitParameters `json:"headersConfig,omitempty" tf:"headers_config,omitempty"`

	// Whether any URL query strings in viewer requests are included in the cache key. It also automatically includes these query strings in requests that CloudFront sends to the origin. Please refer to the Query String Config for more information.
	QueryStringsConfig []QueryStringsConfigInitParameters `json:"queryStringsConfig,omitempty" tf:"query_strings_config,omitempty"`
}

type ParametersInCacheKeyAndForwardedToOriginObservation struct {

	// Whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig []CookiesConfigObservation `json:"cookiesConfig,omitempty" tf:"cookies_config,omitempty"`

	// Flag determines whether the Accept-Encoding HTTP header is included in the cache key and in requests that CloudFront sends to the origin.
	EnableAcceptEncodingBrotli *bool `json:"enableAcceptEncodingBrotli,omitempty" tf:"enable_accept_encoding_brotli,omitempty"`

	// Whether the Accept-Encoding HTTP header is included in the cache key and in requests sent to the origin by CloudFront.
	EnableAcceptEncodingGzip *bool `json:"enableAcceptEncodingGzip,omitempty" tf:"enable_accept_encoding_gzip,omitempty"`

	// Whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig []HeadersConfigObservation `json:"headersConfig,omitempty" tf:"headers_config,omitempty"`

	// Whether any URL query strings in viewer requests are included in the cache key. It also automatically includes these query strings in requests that CloudFront sends to the origin. Please refer to the Query String Config for more information.
	QueryStringsConfig []QueryStringsConfigObservation `json:"queryStringsConfig,omitempty" tf:"query_strings_config,omitempty"`
}

type ParametersInCacheKeyAndForwardedToOriginParameters struct {

	// Whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	// +kubebuilder:validation:Optional
	CookiesConfig []CookiesConfigParameters `json:"cookiesConfig,omitempty" tf:"cookies_config,omitempty"`

	// Flag determines whether the Accept-Encoding HTTP header is included in the cache key and in requests that CloudFront sends to the origin.
	// +kubebuilder:validation:Optional
	EnableAcceptEncodingBrotli *bool `json:"enableAcceptEncodingBrotli,omitempty" tf:"enable_accept_encoding_brotli,omitempty"`

	// Whether the Accept-Encoding HTTP header is included in the cache key and in requests sent to the origin by CloudFront.
	// +kubebuilder:validation:Optional
	EnableAcceptEncodingGzip *bool `json:"enableAcceptEncodingGzip,omitempty" tf:"enable_accept_encoding_gzip,omitempty"`

	// Whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	// +kubebuilder:validation:Optional
	HeadersConfig []HeadersConfigParameters `json:"headersConfig,omitempty" tf:"headers_config,omitempty"`

	// Whether any URL query strings in viewer requests are included in the cache key. It also automatically includes these query strings in requests that CloudFront sends to the origin. Please refer to the Query String Config for more information.
	// +kubebuilder:validation:Optional
	QueryStringsConfig []QueryStringsConfigParameters `json:"queryStringsConfig,omitempty" tf:"query_strings_config,omitempty"`
}

type QueryStringsConfigInitParameters struct {

	// Whether URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for query_string_behavior are none, whitelist, allExcept, and all.
	QueryStringBehavior *string `json:"queryStringBehavior,omitempty" tf:"query_string_behavior,omitempty"`

	// Configuration parameter that contains a list of query string names. See Items for more information.
	QueryStrings []QueryStringsInitParameters `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type QueryStringsConfigObservation struct {

	// Whether URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for query_string_behavior are none, whitelist, allExcept, and all.
	QueryStringBehavior *string `json:"queryStringBehavior,omitempty" tf:"query_string_behavior,omitempty"`

	// Configuration parameter that contains a list of query string names. See Items for more information.
	QueryStrings []QueryStringsObservation `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type QueryStringsConfigParameters struct {

	// Whether URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for query_string_behavior are none, whitelist, allExcept, and all.
	// +kubebuilder:validation:Optional
	QueryStringBehavior *string `json:"queryStringBehavior,omitempty" tf:"query_string_behavior,omitempty"`

	// Configuration parameter that contains a list of query string names. See Items for more information.
	// +kubebuilder:validation:Optional
	QueryStrings []QueryStringsParameters `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type QueryStringsInitParameters struct {

	// List of item names, such as cookies, headers, or query strings.
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type QueryStringsObservation struct {

	// List of item names, such as cookies, headers, or query strings.
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type QueryStringsParameters struct {

	// List of item names, such as cookies, headers, or query strings.
	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

// CachePolicySpec defines the desired state of CachePolicy
type CachePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CachePolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider CachePolicyInitParameters `json:"initProvider,omitempty"`
}

// CachePolicyStatus defines the observed state of CachePolicy.
type CachePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CachePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CachePolicy is the Schema for the CachePolicys API. Use the
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CachePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.parametersInCacheKeyAndForwardedToOrigin) || has(self.initProvider.parametersInCacheKeyAndForwardedToOrigin)",message="parametersInCacheKeyAndForwardedToOrigin is a required parameter"
	Spec   CachePolicySpec   `json:"spec"`
	Status CachePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CachePolicyList contains a list of CachePolicys
type CachePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CachePolicy `json:"items"`
}

// Repository type metadata.
var (
	CachePolicy_Kind             = "CachePolicy"
	CachePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CachePolicy_Kind}.String()
	CachePolicy_KindAPIVersion   = CachePolicy_Kind + "." + CRDGroupVersion.String()
	CachePolicy_GroupVersionKind = CRDGroupVersion.WithKind(CachePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&CachePolicy{}, &CachePolicyList{})
}
