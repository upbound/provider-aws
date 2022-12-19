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

type DNSTargetResourceObservation struct {
}

type DNSTargetResourceParameters struct {

	// DNS Name that acts as the ingress point to a portion of application.
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// Hosted Zone ARN that contains the DNS record with the provided name of target resource.
	// +kubebuilder:validation:Optional
	HostedZoneArn *string `json:"hostedZoneArn,omitempty" tf:"hosted_zone_arn,omitempty"`

	// Route53 record set id to uniquely identify a record given a domain_name and a record_type.
	// +kubebuilder:validation:Optional
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`

	// Type of DNS Record of target resource.
	// +kubebuilder:validation:Optional
	RecordType *string `json:"recordType,omitempty" tf:"record_type,omitempty"`

	// Target resource the R53 record specified with the above params points to.
	// +kubebuilder:validation:Optional
	TargetResource []TargetResourceParameters `json:"targetResource,omitempty" tf:"target_resource,omitempty"`
}

type NlbResourceObservation struct {
}

type NlbResourceParameters struct {

	// NLB resource ARN.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type R53ResourceObservation struct {
}

type R53ResourceParameters struct {

	// Domain name that is targeted.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Resource record set ID that is targeted.
	// +kubebuilder:validation:Optional
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`
}

type ResourceSetObservation struct {

	// ARN of the resource set
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of resources to add to this resource set. See below.
	// +kubebuilder:validation:Required
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ResourceSetParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Type of the resources in the resource set.
	// +kubebuilder:validation:Required
	ResourceSetType *string `json:"resourceSetType" tf:"resource_set_type,omitempty"`

	// List of resources to add to this resource set. See below.
	// +kubebuilder:validation:Required
	Resources []ResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ResourcesObservation struct {

	// Unique identified for DNS Target Resources, use for readiness checks.
	ComponentID *string `json:"componentId,omitempty" tf:"component_id,omitempty"`
}

type ResourcesParameters struct {

	// Component for DNS/Routing Control Readiness Checks.
	// +kubebuilder:validation:Optional
	DNSTargetResource []DNSTargetResourceParameters `json:"dnsTargetResource,omitempty" tf:"dns_target_resource,omitempty"`

	// Recovery group ARN or cell ARN that contains this resource set.
	// +kubebuilder:validation:Optional
	ReadinessScopes []*string `json:"readinessScopes,omitempty" tf:"readiness_scopes,omitempty"`

	// ARN of the resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatch/v1beta1.MetricAlarm
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a MetricAlarm in cloudwatch to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a MetricAlarm in cloudwatch to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

type TargetResourceObservation struct {
}

type TargetResourceParameters struct {

	// NLB resource a DNS Target Resource points to. Required if r53_resource is not set.
	// +kubebuilder:validation:Optional
	NlbResource []NlbResourceParameters `json:"nlbResource,omitempty" tf:"nlb_resource,omitempty"`

	// Route53 resource a DNS Target Resource record points to.
	// +kubebuilder:validation:Optional
	R53Resource []R53ResourceParameters `json:"r53Resource,omitempty" tf:"r53_resource,omitempty"`
}

// ResourceSetSpec defines the desired state of ResourceSet
type ResourceSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceSetParameters `json:"forProvider"`
}

// ResourceSetStatus defines the observed state of ResourceSet.
type ResourceSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceSet is the Schema for the ResourceSets API. Provides an AWS Route 53 Recovery Readiness Resource Set
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourceSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceSetSpec   `json:"spec"`
	Status            ResourceSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceSetList contains a list of ResourceSets
type ResourceSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceSet `json:"items"`
}

// Repository type metadata.
var (
	ResourceSet_Kind             = "ResourceSet"
	ResourceSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceSet_Kind}.String()
	ResourceSet_KindAPIVersion   = ResourceSet_Kind + "." + CRDGroupVersion.String()
	ResourceSet_GroupVersionKind = CRDGroupVersion.WithKind(ResourceSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceSet{}, &ResourceSetList{})
}
