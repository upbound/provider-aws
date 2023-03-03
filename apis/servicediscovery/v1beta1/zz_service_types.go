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

type DNSConfigObservation struct {

	// An array that contains one DnsRecord object for each resource record set.
	DNSRecords []DNSRecordsObservation `json:"dnsRecords,omitempty" tf:"dns_records,omitempty"`

	// The ID of the namespace to use for DNS configuration.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// The routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED
	RoutingPolicy *string `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`
}

type DNSConfigParameters struct {

	// An array that contains one DnsRecord object for each resource record set.
	// +kubebuilder:validation:Required
	DNSRecords []DNSRecordsParameters `json:"dnsRecords" tf:"dns_records,omitempty"`

	// The ID of the namespace to use for DNS configuration.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/servicediscovery/v1beta1.PrivateDNSNamespace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a PrivateDNSNamespace in servicediscovery to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSNamespace in servicediscovery to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// The routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED
	// +kubebuilder:validation:Optional
	RoutingPolicy *string `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`
}

type DNSRecordsObservation struct {

	// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The type of the resource, which indicates the value that Amazon Route 53 returns in response to DNS queries. Valid Values: A, AAAA, SRV, CNAME
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DNSRecordsParameters struct {

	// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set.
	// +kubebuilder:validation:Required
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`

	// The type of the resource, which indicates the value that Amazon Route 53 returns in response to DNS queries. Valid Values: A, AAAA, SRV, CNAME
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type HealthCheckConfigObservation struct {

	// The number of consecutive health checks. Maximum value of 10.
	FailureThreshold *float64 `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`

	// The path that you want Route 53 to request when performing health checks. Route 53 automatically adds the DNS name for the service. If you don't specify a value, the default value is /.
	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path,omitempty"`

	// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthCheckConfigParameters struct {

	// The number of consecutive health checks. Maximum value of 10.
	// +kubebuilder:validation:Optional
	FailureThreshold *float64 `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`

	// The path that you want Route 53 to request when performing health checks. Route 53 automatically adds the DNS name for the service. If you don't specify a value, the default value is /.
	// +kubebuilder:validation:Optional
	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path,omitempty"`

	// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthCheckCustomConfigObservation struct {

	// The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance.  Maximum value of 10.
	FailureThreshold *float64 `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`
}

type HealthCheckCustomConfigParameters struct {

	// The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance.  Maximum value of 10.
	// +kubebuilder:validation:Optional
	FailureThreshold *float64 `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`
}

type ServiceObservation struct {

	// The ARN of the service.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
	DNSConfig []DNSConfigObservation `json:"dnsConfig,omitempty" tf:"dns_config,omitempty"`

	// The description of the service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A boolean that indicates all instances should be deleted from the service so that the service can be destroyed without error. These instances are not recoverable.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
	HealthCheckConfig []HealthCheckConfigObservation `json:"healthCheckConfig,omitempty" tf:"health_check_config,omitempty"`

	// A complex type that contains settings for ECS managed health checks.
	HealthCheckCustomConfig []HealthCheckCustomConfigObservation `json:"healthCheckCustomConfig,omitempty" tf:"health_check_custom_config,omitempty"`

	// The ID of the service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the service.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the namespace that you want to use to create the service.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// If present, specifies that the service instances are only discoverable using the DiscoverInstances API operation. No DNS records is registered for the service instances. The only valid value is HTTP.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServiceParameters struct {

	// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
	// +kubebuilder:validation:Optional
	DNSConfig []DNSConfigParameters `json:"dnsConfig,omitempty" tf:"dns_config,omitempty"`

	// The description of the service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A boolean that indicates all instances should be deleted from the service so that the service can be destroyed without error. These instances are not recoverable.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
	// +kubebuilder:validation:Optional
	HealthCheckConfig []HealthCheckConfigParameters `json:"healthCheckConfig,omitempty" tf:"health_check_config,omitempty"`

	// A complex type that contains settings for ECS managed health checks.
	// +kubebuilder:validation:Optional
	HealthCheckCustomConfig []HealthCheckCustomConfigParameters `json:"healthCheckCustomConfig,omitempty" tf:"health_check_custom_config,omitempty"`

	// The name of the service.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the namespace that you want to use to create the service.
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// If present, specifies that the service instances are only discoverable using the DiscoverInstances API operation. No DNS records is registered for the service instances. The only valid value is HTTP.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API. Provides a Service Discovery Service resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   ServiceSpec   `json:"spec"`
	Status ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
