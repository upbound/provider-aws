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

type InfrastructureConfigurationObservation struct {

	// Amazon Resource Name (ARN) of the configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Date when the configuration was created.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// Date when the configuration was updated.
	DateUpdated *string `json:"dateUpdated,omitempty" tf:"date_updated,omitempty"`

	// Description for the configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amazon Resource Name (ARN) of the configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block with instance metadata options for the HTTP requests that pipeline builds use to launch EC2 build and test instances. Detailed below.
	InstanceMetadataOptions []InstanceMetadataOptionsObservation `json:"instanceMetadataOptions,omitempty" tf:"instance_metadata_options,omitempty"`

	// Name of IAM Instance Profile.
	InstanceProfileName *string `json:"instanceProfileName,omitempty" tf:"instance_profile_name,omitempty"`

	// Set of EC2 Instance Types.
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// Name of EC2 Key Pair.
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Configuration block with logging settings. Detailed below.
	Logging []LoggingObservation `json:"logging,omitempty" tf:"logging,omitempty"`

	// Name for the configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags to assign to infrastructure created by the configuration.
	ResourceTags map[string]*string `json:"resourceTags,omitempty" tf:"resource_tags,omitempty"`

	// Set of EC2 Security Group identifiers.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Amazon Resource Name (ARN) of SNS Topic.
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// EC2 Subnet identifier. Also requires security_group_ids argument.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Enable if the instance should be terminated when the pipeline fails. Defaults to false.
	TerminateInstanceOnFailure *bool `json:"terminateInstanceOnFailure,omitempty" tf:"terminate_instance_on_failure,omitempty"`
}

type InfrastructureConfigurationParameters struct {

	// Description for the configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block with instance metadata options for the HTTP requests that pipeline builds use to launch EC2 build and test instances. Detailed below.
	// +kubebuilder:validation:Optional
	InstanceMetadataOptions []InstanceMetadataOptionsParameters `json:"instanceMetadataOptions,omitempty" tf:"instance_metadata_options,omitempty"`

	// Name of IAM Instance Profile.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.InstanceProfile
	// +kubebuilder:validation:Optional
	InstanceProfileName *string `json:"instanceProfileName,omitempty" tf:"instance_profile_name,omitempty"`

	// Reference to a InstanceProfile in iam to populate instanceProfileName.
	// +kubebuilder:validation:Optional
	InstanceProfileNameRef *v1.Reference `json:"instanceProfileNameRef,omitempty" tf:"-"`

	// Selector for a InstanceProfile in iam to populate instanceProfileName.
	// +kubebuilder:validation:Optional
	InstanceProfileNameSelector *v1.Selector `json:"instanceProfileNameSelector,omitempty" tf:"-"`

	// Set of EC2 Instance Types.
	// +kubebuilder:validation:Optional
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// Name of EC2 Key Pair.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.KeyPair
	// +kubebuilder:validation:Optional
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Reference to a KeyPair in ec2 to populate keyPair.
	// +kubebuilder:validation:Optional
	KeyPairRef *v1.Reference `json:"keyPairRef,omitempty" tf:"-"`

	// Selector for a KeyPair in ec2 to populate keyPair.
	// +kubebuilder:validation:Optional
	KeyPairSelector *v1.Selector `json:"keyPairSelector,omitempty" tf:"-"`

	// Configuration block with logging settings. Detailed below.
	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// Name for the configuration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags to assign to infrastructure created by the configuration.
	// +kubebuilder:validation:Optional
	ResourceTags map[string]*string `json:"resourceTags,omitempty" tf:"resource_tags,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Set of EC2 Security Group identifiers.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Amazon Resource Name (ARN) of SNS Topic.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// Reference to a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnRef *v1.Reference `json:"snsTopicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnSelector *v1.Selector `json:"snsTopicArnSelector,omitempty" tf:"-"`

	// EC2 Subnet identifier. Also requires security_group_ids argument.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Enable if the instance should be terminated when the pipeline fails. Defaults to false.
	// +kubebuilder:validation:Optional
	TerminateInstanceOnFailure *bool `json:"terminateInstanceOnFailure,omitempty" tf:"terminate_instance_on_failure,omitempty"`
}

type InstanceMetadataOptionsObservation struct {

	// The number of hops that an instance can traverse to reach its destonation.
	HTTPPutResponseHopLimit *float64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit,omitempty"`

	// Whether a signed token is required for instance metadata retrieval requests. Valid values: required, optional.
	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens,omitempty"`
}

type InstanceMetadataOptionsParameters struct {

	// The number of hops that an instance can traverse to reach its destonation.
	// +kubebuilder:validation:Optional
	HTTPPutResponseHopLimit *float64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit,omitempty"`

	// Whether a signed token is required for instance metadata retrieval requests. Valid values: required, optional.
	// +kubebuilder:validation:Optional
	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens,omitempty"`
}

type LoggingObservation struct {

	// Configuration block with S3 logging settings. Detailed below.
	S3Logs []S3LogsObservation `json:"s3Logs,omitempty" tf:"s3_logs,omitempty"`
}

type LoggingParameters struct {

	// Configuration block with S3 logging settings. Detailed below.
	// +kubebuilder:validation:Required
	S3Logs []S3LogsParameters `json:"s3Logs" tf:"s3_logs,omitempty"`
}

type S3LogsObservation struct {

	// Name of the S3 Bucket.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// Prefix to use for S3 logs. Defaults to /.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type S3LogsParameters struct {

	// Name of the S3 Bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// Reference to a Bucket in s3 to populate s3BucketName.
	// +kubebuilder:validation:Optional
	S3BucketNameRef *v1.Reference `json:"s3BucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate s3BucketName.
	// +kubebuilder:validation:Optional
	S3BucketNameSelector *v1.Selector `json:"s3BucketNameSelector,omitempty" tf:"-"`

	// Prefix to use for S3 logs. Defaults to /.
	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

// InfrastructureConfigurationSpec defines the desired state of InfrastructureConfiguration
type InfrastructureConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InfrastructureConfigurationParameters `json:"forProvider"`
}

// InfrastructureConfigurationStatus defines the observed state of InfrastructureConfiguration.
type InfrastructureConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InfrastructureConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InfrastructureConfiguration is the Schema for the InfrastructureConfigurations API. Manages an Image Builder Infrastructure Configuration
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InfrastructureConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InfrastructureConfigurationSpec   `json:"spec"`
	Status            InfrastructureConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InfrastructureConfigurationList contains a list of InfrastructureConfigurations
type InfrastructureConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InfrastructureConfiguration `json:"items"`
}

// Repository type metadata.
var (
	InfrastructureConfiguration_Kind             = "InfrastructureConfiguration"
	InfrastructureConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InfrastructureConfiguration_Kind}.String()
	InfrastructureConfiguration_KindAPIVersion   = InfrastructureConfiguration_Kind + "." + CRDGroupVersion.String()
	InfrastructureConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(InfrastructureConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&InfrastructureConfiguration{}, &InfrastructureConfigurationList{})
}
