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

type SubscriptionFilterObservation struct {

	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution *string `json:"distribution,omitempty" tf:"distribution,omitempty"`

	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events. Use empty string "" to match everything. For more information, see the Amazon CloudWatch Logs User Guide.
	FilterPattern *string `json:"filterPattern,omitempty" tf:"filter_pattern,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the log group to associate the subscription filter with
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// A name for the subscription filter
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use aws_lambda_permission resource for granting access from CloudWatch logs to the destination Lambda function.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type SubscriptionFilterParameters struct {

	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// Reference to a Stream in kinesis to populate destinationArn.
	// +kubebuilder:validation:Optional
	DestinationArnRef *v1.Reference `json:"destinationArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate destinationArn.
	// +kubebuilder:validation:Optional
	DestinationArnSelector *v1.Selector `json:"destinationArnSelector,omitempty" tf:"-"`

	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	// +kubebuilder:validation:Optional
	Distribution *string `json:"distribution,omitempty" tf:"distribution,omitempty"`

	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events. Use empty string "" to match everything. For more information, see the Amazon CloudWatch Logs User Guide.
	// +kubebuilder:validation:Optional
	FilterPattern *string `json:"filterPattern,omitempty" tf:"filter_pattern,omitempty"`

	// The name of the log group to associate the subscription filter with
	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// A name for the subscription filter
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use aws_lambda_permission resource for granting access from CloudWatch logs to the destination Lambda function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

// SubscriptionFilterSpec defines the desired state of SubscriptionFilter
type SubscriptionFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionFilterParameters `json:"forProvider"`
}

// SubscriptionFilterStatus defines the observed state of SubscriptionFilter.
type SubscriptionFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionFilter is the Schema for the SubscriptionFilters API. Provides a CloudWatch Logs subscription filter.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SubscriptionFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.filterPattern)",message="filterPattern is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.logGroupName)",message="logGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   SubscriptionFilterSpec   `json:"spec"`
	Status SubscriptionFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionFilterList contains a list of SubscriptionFilters
type SubscriptionFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubscriptionFilter `json:"items"`
}

// Repository type metadata.
var (
	SubscriptionFilter_Kind             = "SubscriptionFilter"
	SubscriptionFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubscriptionFilter_Kind}.String()
	SubscriptionFilter_KindAPIVersion   = SubscriptionFilter_Kind + "." + CRDGroupVersion.String()
	SubscriptionFilter_GroupVersionKind = CRDGroupVersion.WithKind(SubscriptionFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&SubscriptionFilter{}, &SubscriptionFilterList{})
}
