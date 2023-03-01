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

type AccessLogsObservation struct {
}

type AccessLogsParameters struct {

	// The S3 bucket name to store the logs in.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// The S3 bucket prefix. Logs are stored in the root if not configured.
	// +kubebuilder:validation:Optional
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// Boolean to enable / disable access_logs. Default is true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The publishing interval in minutes. Valid values: 5 and 60. Default: 60
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`
}

type ELBObservation struct {

	// The ARN of the ELB
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The DNS name of the ELB
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The name of the ELB
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the security group that you can use as
	// part of your inbound rules for your load balancer's back-end application
	// instances. Only available on ELBs launched in a VPC.
	SourceSecurityGroupID *string `json:"sourceSecurityGroupId,omitempty" tf:"source_security_group_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The canonical hosted zone ID of the ELB (to be used in a Route 53 Alias record)
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ELBParameters struct {

	// An Access Logs block. Access Logs documented below.
	// +kubebuilder:validation:Optional
	AccessLogs []AccessLogsParameters `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`

	// The AZ's to serve traffic in.
	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Boolean to enable connection draining. Default: false
	// +kubebuilder:validation:Optional
	ConnectionDraining *bool `json:"connectionDraining,omitempty" tf:"connection_draining,omitempty"`

	// The time in seconds to allow for connections to drain. Default: 300
	// +kubebuilder:validation:Optional
	ConnectionDrainingTimeout *float64 `json:"connectionDrainingTimeout,omitempty" tf:"connection_draining_timeout,omitempty"`

	// Enable cross-zone load balancing. Default: true
	// +kubebuilder:validation:Optional
	CrossZoneLoadBalancing *bool `json:"crossZoneLoadBalancing,omitempty" tf:"cross_zone_load_balancing,omitempty"`

	// Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are monitor, defensive (default), strictest.
	// +kubebuilder:validation:Optional
	DesyncMitigationMode *string `json:"desyncMitigationMode,omitempty" tf:"desync_mitigation_mode,omitempty"`

	// A health_check block. Health Check documented below.
	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The time in seconds that the connection is allowed to be idle. Default: 60
	// +kubebuilder:validation:Optional
	IdleTimeout *float64 `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	// A list of instance ids to place in the ELB pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +kubebuilder:validation:Optional
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// References to Instance in ec2 to populate instances.
	// +kubebuilder:validation:Optional
	InstancesRefs []v1.Reference `json:"instancesRefs,omitempty" tf:"-"`

	// Selector for a list of Instance in ec2 to populate instances.
	// +kubebuilder:validation:Optional
	InstancesSelector *v1.Selector `json:"instancesSelector,omitempty" tf:"-"`

	// If true, ELB will be an internal ELB.
	// +kubebuilder:validation:Optional
	Internal *bool `json:"internal,omitempty" tf:"internal,omitempty"`

	// A list of listener blocks. Listeners documented below.
	// +kubebuilder:validation:Optional
	Listener []ListenerParameters `json:"listener,omitempty" tf:"listener,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A list of security group IDs to assign to the ELB.
	// Only valid if creating an ELB within a VPC
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The name of the security group that you can use as
	// part of your inbound rules for your load balancer's back-end application
	// instances. Use this for Classic or Default VPC only.
	// +kubebuilder:validation:Optional
	SourceSecurityGroup *string `json:"sourceSecurityGroup,omitempty" tf:"source_security_group,omitempty"`

	// A list of subnet IDs to attach to the ELB.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// References to Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsRefs []v1.Reference `json:"subnetsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsSelector *v1.Selector `json:"subnetsSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HealthCheckObservation struct {
}

type HealthCheckParameters struct {

	// The number of checks before the instance is declared healthy.
	// +kubebuilder:validation:Required
	HealthyThreshold *float64 `json:"healthyThreshold" tf:"healthy_threshold,omitempty"`

	// The publishing interval in minutes. Valid values: 5 and 60. Default: 60
	// +kubebuilder:validation:Required
	Interval *float64 `json:"interval" tf:"interval,omitempty"`

	// The target of the check. Valid pattern is "${PROTOCOL}:${PORT}${PATH}", where PROTOCOL
	// values are:
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`

	// The length of time before the check times out.
	// +kubebuilder:validation:Required
	Timeout *float64 `json:"timeout" tf:"timeout,omitempty"`

	// The number of checks before the instance is declared unhealthy.
	// +kubebuilder:validation:Required
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" tf:"unhealthy_threshold,omitempty"`
}

type ListenerObservation struct {
}

type ListenerParameters struct {

	// The port on the instance to route to
	// +kubebuilder:validation:Required
	InstancePort *float64 `json:"instancePort" tf:"instance_port,omitempty"`

	// The protocol to use to the instance. Valid
	// values are HTTP, HTTPS, TCP, or SSL
	// +kubebuilder:validation:Required
	InstanceProtocol *string `json:"instanceProtocol" tf:"instance_protocol,omitempty"`

	// The port to listen on for the load balancer
	// +kubebuilder:validation:Required
	LBPort *float64 `json:"lbPort" tf:"lb_port,omitempty"`

	// The protocol to listen on. Valid values are HTTP,
	// HTTPS, TCP, or SSL
	// +kubebuilder:validation:Required
	LBProtocol *string `json:"lbProtocol" tf:"lb_protocol,omitempty"`

	// The ARN of an SSL certificate you have
	// uploaded to AWS IAM. Note ECDSA-specific restrictions below.  Only valid when
	// +kubebuilder:validation:Optional
	SSLCertificateID *string `json:"sslCertificateId,omitempty" tf:"ssl_certificate_id,omitempty"`
}

// ELBSpec defines the desired state of ELB
type ELBSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ELBParameters `json:"forProvider"`
}

// ELBStatus defines the observed state of ELB.
type ELBStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ELBObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ELB is the Schema for the ELBs API. Provides an Elastic Load Balancer resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ELB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.listener)",message="listener is a required parameter"
	Spec   ELBSpec   `json:"spec"`
	Status ELBStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ELBList contains a list of ELBs
type ELBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ELB `json:"items"`
}

// Repository type metadata.
var (
	ELB_Kind             = "ELB"
	ELB_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ELB_Kind}.String()
	ELB_KindAPIVersion   = ELB_Kind + "." + CRDGroupVersion.String()
	ELB_GroupVersionKind = CRDGroupVersion.WithKind(ELB_Kind)
)

func init() {
	SchemeBuilder.Register(&ELB{}, &ELBList{})
}
