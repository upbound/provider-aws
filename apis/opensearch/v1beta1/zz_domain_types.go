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

type AdvancedSecurityOptionsObservation struct {
}

type AdvancedSecurityOptionsParameters struct {

	// Whether Anonymous auth is enabled. Enables fine-grained access control on an existing domain. Ignored unless advanced_security_options are enabled. Can only be enabled on an existing domain.
	// +kubebuilder:validation:Optional
	AnonymousAuthEnabled *bool `json:"anonymousAuthEnabled,omitempty" tf:"anonymous_auth_enabled,omitempty"`

	// Whether advanced security is enabled.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Whether the internal user database is enabled. Default is false.
	// +kubebuilder:validation:Optional
	InternalUserDatabaseEnabled *bool `json:"internalUserDatabaseEnabled,omitempty" tf:"internal_user_database_enabled,omitempty"`

	// Configuration block for the main user. Detailed below.
	// +kubebuilder:validation:Optional
	MasterUserOptions []MasterUserOptionsParameters `json:"masterUserOptions,omitempty" tf:"master_user_options,omitempty"`
}

type AutoTuneOptionsObservation struct {
}

type AutoTuneOptionsParameters struct {

	// Auto-Tune desired state for the domain. Valid values: ENABLED or DISABLED.
	// +kubebuilder:validation:Required
	DesiredState *string `json:"desiredState" tf:"desired_state,omitempty"`

	// Configuration block for Auto-Tune maintenance windows. Can be specified multiple times for each maintenance window. Detailed below.
	// +kubebuilder:validation:Optional
	MaintenanceSchedule []MaintenanceScheduleParameters `json:"maintenanceSchedule,omitempty" tf:"maintenance_schedule,omitempty"`

	// Whether to roll back to default Auto-Tune settings when disabling Auto-Tune. Valid values: DEFAULT_ROLLBACK or NO_ROLLBACK.
	// +kubebuilder:validation:Optional
	RollbackOnDisable *string `json:"rollbackOnDisable,omitempty" tf:"rollback_on_disable,omitempty"`
}

type ClusterConfigObservation struct {
}

type ClusterConfigParameters struct {

	// Configuration block containing cold storage configuration. Detailed below.
	// +kubebuilder:validation:Optional
	ColdStorageOptions []ColdStorageOptionsParameters `json:"coldStorageOptions,omitempty" tf:"cold_storage_options,omitempty"`

	// Number of dedicated main nodes in the cluster.
	// +kubebuilder:validation:Optional
	DedicatedMasterCount *float64 `json:"dedicatedMasterCount,omitempty" tf:"dedicated_master_count,omitempty"`

	// Whether dedicated main nodes are enabled for the cluster.
	// +kubebuilder:validation:Optional
	DedicatedMasterEnabled *bool `json:"dedicatedMasterEnabled,omitempty" tf:"dedicated_master_enabled,omitempty"`

	// Instance type of the dedicated main nodes in the cluster.
	// +kubebuilder:validation:Optional
	DedicatedMasterType *string `json:"dedicatedMasterType,omitempty" tf:"dedicated_master_type,omitempty"`

	// Number of instances in the cluster.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Instance type of data nodes in the cluster.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Number of warm nodes in the cluster. Valid values are between 2 and 150. warm_count can be only and must be set when warm_enabled is set to true.
	// +kubebuilder:validation:Optional
	WarmCount *float64 `json:"warmCount,omitempty" tf:"warm_count,omitempty"`

	// Whether to enable warm storage.
	// +kubebuilder:validation:Optional
	WarmEnabled *bool `json:"warmEnabled,omitempty" tf:"warm_enabled,omitempty"`

	// Instance type for the OpenSearch cluster's warm nodes. Valid values are ultrawarm1.medium.search, ultrawarm1.large.search and ultrawarm1.xlarge.search. warm_type can be only and must be set when warm_enabled is set to true.
	// +kubebuilder:validation:Optional
	WarmType *string `json:"warmType,omitempty" tf:"warm_type,omitempty"`

	// Configuration block containing zone awareness settings. Detailed below.
	// +kubebuilder:validation:Optional
	ZoneAwarenessConfig []ZoneAwarenessConfigParameters `json:"zoneAwarenessConfig,omitempty" tf:"zone_awareness_config,omitempty"`

	// Whether zone awareness is enabled, set to true for multi-az deployment. To enable awareness with three Availability Zones, the availability_zone_count within the zone_awareness_config must be set to 3.
	// +kubebuilder:validation:Optional
	ZoneAwarenessEnabled *bool `json:"zoneAwarenessEnabled,omitempty" tf:"zone_awareness_enabled,omitempty"`
}

type CognitoOptionsObservation struct {
}

type CognitoOptionsParameters struct {

	// Whether Amazon Cognito authentication with Kibana is enabled or not. Default is false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// ID of the Cognito Identity Pool to use.
	// +kubebuilder:validation:Required
	IdentityPoolID *string `json:"identityPoolId" tf:"identity_pool_id,omitempty"`

	// ARN of the IAM role that has the AmazonOpenSearchServiceCognitoAccess policy attached.
	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`

	// ID of the Cognito User Pool to use.
	// +kubebuilder:validation:Required
	UserPoolID *string `json:"userPoolId" tf:"user_pool_id,omitempty"`
}

type ColdStorageOptionsObservation struct {
}

type ColdStorageOptionsParameters struct {

	// Boolean to enable cold storage for an OpenSearch domain. Defaults to false. Master and ultrawarm nodes must be enabled for cold storage.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DomainEndpointOptionsObservation struct {
}

type DomainEndpointOptionsParameters struct {

	// Fully qualified domain for your custom endpoint.
	// +kubebuilder:validation:Optional
	CustomEndpoint *string `json:"customEndpoint,omitempty" tf:"custom_endpoint,omitempty"`

	// ACM certificate ARN for your custom endpoint.
	// +kubebuilder:validation:Optional
	CustomEndpointCertificateArn *string `json:"customEndpointCertificateArn,omitempty" tf:"custom_endpoint_certificate_arn,omitempty"`

	// Whether to enable custom endpoint for the OpenSearch domain.
	// +kubebuilder:validation:Optional
	CustomEndpointEnabled *bool `json:"customEndpointEnabled,omitempty" tf:"custom_endpoint_enabled,omitempty"`

	// Whether or not to require HTTPS. Defaults to true.
	// +kubebuilder:validation:Optional
	EnforceHTTPS *bool `json:"enforceHttps,omitempty" tf:"enforce_https,omitempty"`

	// Name of the TLS security policy that needs to be applied to the HTTPS endpoint. Valid values:  Policy-Min-TLS-1-0-2019-07 and Policy-Min-TLS-1-2-2019-07.
	// +kubebuilder:validation:Optional
	TLSSecurityPolicy *string `json:"tlsSecurityPolicy,omitempty" tf:"tls_security_policy,omitempty"`
}

type DomainObservation struct {

	// , are prefaced with es: for both.
	AccessPolicies *string `json:"accessPolicies,omitempty" tf:"access_policies,omitempty"`

	// ARN of the domain.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Unique identifier for the domain.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// Domain-specific endpoint used to submit index, search, and data upload requests.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Domain-specific endpoint for kibana without https scheme.
	KibanaEndpoint *string `json:"kibanaEndpoint,omitempty" tf:"kibana_endpoint,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Configuration block for VPC related options. Adding or removing this configuration forces a new resource (documentation). Detailed below.
	// +kubebuilder:validation:Optional
	VPCOptions []VPCOptionsObservation `json:"vpcOptions,omitempty" tf:"vpc_options,omitempty"`
}

type DomainParameters struct {

	// Key-value string pairs to specify advanced configuration options.
	// +kubebuilder:validation:Optional
	AdvancedOptions map[string]*string `json:"advancedOptions,omitempty" tf:"advanced_options,omitempty"`

	// Configuration block for fine-grained access control. Detailed below.
	// +kubebuilder:validation:Optional
	AdvancedSecurityOptions []AdvancedSecurityOptionsParameters `json:"advancedSecurityOptions,omitempty" tf:"advanced_security_options,omitempty"`

	// Configuration block for the Auto-Tune options of the domain. Detailed below.
	// +kubebuilder:validation:Optional
	AutoTuneOptions []AutoTuneOptionsParameters `json:"autoTuneOptions,omitempty" tf:"auto_tune_options,omitempty"`

	// Configuration block for the cluster of the domain. Detailed below.
	// +kubebuilder:validation:Optional
	ClusterConfig []ClusterConfigParameters `json:"clusterConfig,omitempty" tf:"cluster_config,omitempty"`

	// Configuration block for authenticating Kibana with Cognito. Detailed below.
	// +kubebuilder:validation:Optional
	CognitoOptions []CognitoOptionsParameters `json:"cognitoOptions,omitempty" tf:"cognito_options,omitempty"`

	// Configuration block for domain endpoint HTTP(S) related options. Detailed below.
	// +kubebuilder:validation:Optional
	DomainEndpointOptions []DomainEndpointOptionsParameters `json:"domainEndpointOptions,omitempty" tf:"domain_endpoint_options,omitempty"`

	// Name of the domain.
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// Configuration block for EBS related options, may be required based on chosen instance size. Detailed below.
	// +kubebuilder:validation:Optional
	EBSOptions []EBSOptionsParameters `json:"ebsOptions,omitempty" tf:"ebs_options,omitempty"`

	// Configuration block for encrypt at rest options. Only available for certain instance types. Detailed below.
	// +kubebuilder:validation:Optional
	EncryptAtRest []EncryptAtRestParameters `json:"encryptAtRest,omitempty" tf:"encrypt_at_rest,omitempty"`

	// while Elasticsearch has elasticsearch_version
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
	// +kubebuilder:validation:Optional
	LogPublishingOptions []LogPublishingOptionsParameters `json:"logPublishingOptions,omitempty" tf:"log_publishing_options,omitempty"`

	// Configuration block for node-to-node encryption options. Detailed below.
	// +kubebuilder:validation:Optional
	NodeToNodeEncryption []NodeToNodeEncryptionParameters `json:"nodeToNodeEncryption,omitempty" tf:"node_to_node_encryption,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running OpenSearch 5.3 and later, Amazon OpenSearch takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions, OpenSearch takes daily automated snapshots.
	// +kubebuilder:validation:Optional
	SnapshotOptions []SnapshotOptionsParameters `json:"snapshotOptions,omitempty" tf:"snapshot_options,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for VPC related options. Adding or removing this configuration forces a new resource (documentation). Detailed below.
	// +kubebuilder:validation:Optional
	VPCOptions []VPCOptionsParameters `json:"vpcOptions,omitempty" tf:"vpc_options,omitempty"`
}

type DurationObservation struct {
}

type DurationParameters struct {

	// Unit of time specifying the duration of an Auto-Tune maintenance window. Valid values: HOURS.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// An integer specifying the value of the duration of an Auto-Tune maintenance window.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type EBSOptionsObservation struct {
}

type EBSOptionsParameters struct {

	// Whether EBS volumes are attached to data nodes in the domain.
	// +kubebuilder:validation:Required
	EBSEnabled *bool `json:"ebsEnabled" tf:"ebs_enabled,omitempty"`

	// Baseline input/output (I/O) performance of EBS volumes attached to data nodes. Applicable only for the GP3 and Provisioned IOPS EBS volume types.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Specifies the throughput (in MiB/s) of the EBS volumes attached to data nodes. Applicable only for the gp3 volume type. Valid values are between 125 and 1000.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Size of EBS volumes attached to data nodes (in GiB).
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of EBS volumes attached to data nodes.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EncryptAtRestObservation struct {
}

type EncryptAtRestParameters struct {

	// Whether to enable encryption at rest. If the encrypt_at_rest block is not provided then this defaults to false. Enabling encryption on new domains requires an engine_version of OpenSearch_X.Y or Elasticsearch_5.1 or greater.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// KMS key ARN to encrypt the Elasticsearch domain with. If not specified then it defaults to using the aws/es service KMS key. Note that KMS will accept a KMS key ID but will return the key ARN.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type LogPublishingOptionsObservation struct {
}

type LogPublishingOptionsParameters struct {

	// ARN of the Cloudwatch log group to which log needs to be published.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1.Group
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn,omitempty" tf:"cloudwatch_log_group_arn,omitempty"`

	// Reference to a Group in cloudwatchlogs to populate cloudwatchLogGroupArn.
	// +kubebuilder:validation:Optional
	CloudwatchLogGroupArnRef *v1.Reference `json:"cloudwatchLogGroupArnRef,omitempty" tf:"-"`

	// Selector for a Group in cloudwatchlogs to populate cloudwatchLogGroupArn.
	// +kubebuilder:validation:Optional
	CloudwatchLogGroupArnSelector *v1.Selector `json:"cloudwatchLogGroupArnSelector,omitempty" tf:"-"`

	// Whether given log publishing option is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Type of OpenSearch log. Valid values: INDEX_SLOW_LOGS, SEARCH_SLOW_LOGS, ES_APPLICATION_LOGS, AUDIT_LOGS.
	// +kubebuilder:validation:Required
	LogType *string `json:"logType" tf:"log_type,omitempty"`
}

type MaintenanceScheduleObservation struct {
}

type MaintenanceScheduleParameters struct {

	// A cron expression specifying the recurrence pattern for an Auto-Tune maintenance schedule.
	// +kubebuilder:validation:Required
	CronExpressionForRecurrence *string `json:"cronExpressionForRecurrence" tf:"cron_expression_for_recurrence,omitempty"`

	// Configuration block for the duration of the Auto-Tune maintenance window. Detailed below.
	// +kubebuilder:validation:Required
	Duration []DurationParameters `json:"duration" tf:"duration,omitempty"`

	// Date and time at which to start the Auto-Tune maintenance schedule in RFC3339 format.
	// +kubebuilder:validation:Required
	StartAt *string `json:"startAt" tf:"start_at,omitempty"`
}

type MasterUserOptionsObservation struct {
}

type MasterUserOptionsParameters struct {

	// ARN for the main user. Only specify if internal_user_database_enabled is not set or set to false.
	// +kubebuilder:validation:Optional
	MasterUserArn *string `json:"masterUserArn,omitempty" tf:"master_user_arn,omitempty"`

	// Main user's username, which is stored in the Amazon OpenSearch Service domain's internal database. Only specify if internal_user_database_enabled is set to true.
	// +kubebuilder:validation:Optional
	MasterUserName *string `json:"masterUserName,omitempty" tf:"master_user_name,omitempty"`

	// Main user's password, which is stored in the Amazon OpenSearch Service domain's internal database. Only specify if internal_user_database_enabled is set to true.
	// +kubebuilder:validation:Optional
	MasterUserPasswordSecretRef *v1.SecretKeySelector `json:"masterUserPasswordSecretRef,omitempty" tf:"-"`
}

type NodeToNodeEncryptionObservation struct {
}

type NodeToNodeEncryptionParameters struct {

	// Whether to enable node-to-node encryption. If the node_to_node_encryption block is not provided then this defaults to false. Enabling node-to-node encryption of a new domain requires an engine_version of OpenSearch_X.Y or Elasticsearch_6.0 or greater.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type SnapshotOptionsObservation struct {
}

type SnapshotOptionsParameters struct {

	// Hour during which the service takes an automated daily snapshot of the indices in the domain.
	// +kubebuilder:validation:Required
	AutomatedSnapshotStartHour *float64 `json:"automatedSnapshotStartHour" tf:"automated_snapshot_start_hour,omitempty"`
}

type VPCOptionsObservation struct {

	// If the domain was created inside a VPC, the names of the availability zones the configured subnet_ids were created inside.
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// If the domain was created inside a VPC, the ID of the VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCOptionsParameters struct {

	// List of VPC Security Group IDs to be applied to the OpenSearch domain endpoints. If omitted, the default Security Group for the VPC will be used.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// List of VPC Subnet IDs for the OpenSearch domain endpoints to be created in.
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type ZoneAwarenessConfigObservation struct {
}

type ZoneAwarenessConfigParameters struct {

	// Number of Availability Zones for the domain to use with zone_awareness_enabled. Defaults to 2. Valid values: 2 or 3.
	// +kubebuilder:validation:Optional
	AvailabilityZoneCount *float64 `json:"availabilityZoneCount,omitempty" tf:"availability_zone_count,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Domain is the Schema for the Domains API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec"`
	Status            DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
