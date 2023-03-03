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

type CachePolicyObservation struct {

	// A comment to describe the cache policy.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// The current version of the cache policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The identifier for the cache policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MinTTL *float64 `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`

	// A unique name to identify the cache policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
	ParametersInCacheKeyAndForwardedToOrigin []ParametersInCacheKeyAndForwardedToOriginObservation `json:"parametersInCacheKeyAndForwardedToOrigin,omitempty" tf:"parameters_in_cache_key_and_forwarded_to_origin,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type CachePolicyParameters struct {

	// A comment to describe the cache policy.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	// +kubebuilder:validation:Optional
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	// +kubebuilder:validation:Optional
	MinTTL *float64 `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`

	// A unique name to identify the cache policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
	// +kubebuilder:validation:Optional
	ParametersInCacheKeyAndForwardedToOrigin []ParametersInCacheKeyAndForwardedToOriginParameters `json:"parametersInCacheKeyAndForwardedToOrigin,omitempty" tf:"parameters_in_cache_key_and_forwarded_to_origin,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type CookiesConfigObservation struct {

	// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values are none, whitelist, allExcept, all.
	CookieBehavior *string `json:"cookieBehavior,omitempty" tf:"cookie_behavior,omitempty"`

	// Object that contains a list of cookie names. See Items for more information.
	Cookies []CookiesObservation `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type CookiesConfigParameters struct {

	// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values are none, whitelist, allExcept, all.
	// +kubebuilder:validation:Required
	CookieBehavior *string `json:"cookieBehavior" tf:"cookie_behavior,omitempty"`

	// Object that contains a list of cookie names. See Items for more information.
	// +kubebuilder:validation:Optional
	Cookies []CookiesParameters `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type CookiesObservation struct {

	// A list of item names (cookies, headers, or query strings).
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type CookiesParameters struct {

	// A list of item names (cookies, headers, or query strings).
	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersConfigObservation struct {

	// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values are none, whitelist.
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	// Object that contains a list of header names. See Items for more information.
	Headers []HeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`
}

type HeadersConfigParameters struct {

	// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values are none, whitelist.
	// +kubebuilder:validation:Optional
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	// Object that contains a list of header names. See Items for more information.
	// +kubebuilder:validation:Optional
	Headers []HeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`
}

type HeadersObservation struct {

	// A list of item names (cookies, headers, or query strings).
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersParameters struct {

	// A list of item names (cookies, headers, or query strings).
	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type ParametersInCacheKeyAndForwardedToOriginObservation struct {

	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig []CookiesConfigObservation `json:"cookiesConfig,omitempty" tf:"cookies_config,omitempty"`

	// A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	EnableAcceptEncodingBrotli *bool `json:"enableAcceptEncodingBrotli,omitempty" tf:"enable_accept_encoding_brotli,omitempty"`

	// A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	EnableAcceptEncodingGzip *bool `json:"enableAcceptEncodingGzip,omitempty" tf:"enable_accept_encoding_gzip,omitempty"`

	// Object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig []HeadersConfigObservation `json:"headersConfig,omitempty" tf:"headers_config,omitempty"`

	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfig []QueryStringsConfigObservation `json:"queryStringsConfig,omitempty" tf:"query_strings_config,omitempty"`
}

type ParametersInCacheKeyAndForwardedToOriginParameters struct {

	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	// +kubebuilder:validation:Required
	CookiesConfig []CookiesConfigParameters `json:"cookiesConfig" tf:"cookies_config,omitempty"`

	// A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	// +kubebuilder:validation:Optional
	EnableAcceptEncodingBrotli *bool `json:"enableAcceptEncodingBrotli,omitempty" tf:"enable_accept_encoding_brotli,omitempty"`

	// A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	// +kubebuilder:validation:Optional
	EnableAcceptEncodingGzip *bool `json:"enableAcceptEncodingGzip,omitempty" tf:"enable_accept_encoding_gzip,omitempty"`

	// Object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	// +kubebuilder:validation:Required
	HeadersConfig []HeadersConfigParameters `json:"headersConfig" tf:"headers_config,omitempty"`

	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	// +kubebuilder:validation:Required
	QueryStringsConfig []QueryStringsConfigParameters `json:"queryStringsConfig" tf:"query_strings_config,omitempty"`
}

type QueryStringsConfigObservation struct {

	// Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values are none, whitelist, allExcept, all.
	QueryStringBehavior *string `json:"queryStringBehavior,omitempty" tf:"query_string_behavior,omitempty"`

	// Object that contains a list of query string names. See Items for more information.
	QueryStrings []QueryStringsObservation `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type QueryStringsConfigParameters struct {

	// Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values are none, whitelist, allExcept, all.
	// +kubebuilder:validation:Required
	QueryStringBehavior *string `json:"queryStringBehavior" tf:"query_string_behavior,omitempty"`

	// Object that contains a list of query string names. See Items for more information.
	// +kubebuilder:validation:Optional
	QueryStrings []QueryStringsParameters `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type QueryStringsObservation struct {

	// A list of item names (cookies, headers, or query strings).
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type QueryStringsParameters struct {

	// A list of item names (cookies, headers, or query strings).
	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

// CachePolicySpec defines the desired state of CachePolicy
type CachePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CachePolicyParameters `json:"forProvider"`
}

// CachePolicyStatus defines the observed state of CachePolicy.
type CachePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CachePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CachePolicy is the Schema for the CachePolicys API. Provides a cache policy for a CloudFront ditribution. When it’s attached to a cache behavior, the cache policy determines the values that CloudFront includes in the cache key. These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer. It also determines the default, minimum, and maximum time to live (TTL) values that you want objects to stay in the CloudFront cache.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CachePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.parametersInCacheKeyAndForwardedToOrigin)",message="parametersInCacheKeyAndForwardedToOrigin is a required parameter"
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
