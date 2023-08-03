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

type AssociationInitParameters struct {

	// By default, when you create a new or update associations, the system runs it immediately and then according to the schedule you specified. Enable this option if you do not want an association to run immediately after you create or update it. This parameter is not supported for rate expressions. Default: false.
	ApplyOnlyAtCronInterval *bool `json:"applyOnlyAtCronInterval,omitempty" tf:"apply_only_at_cron_interval,omitempty"`

	// The descriptive name for the association.
	AssociationName *string `json:"associationName,omitempty" tf:"association_name,omitempty"`

	// Specify the target for the association. This target is required for associations that use an Automation document and target resources by using rate controls. This should be set to the SSM document parameter that will define how your automation will branch out.
	AutomationTargetParameterName *string `json:"automationTargetParameterName,omitempty" tf:"automation_target_parameter_name,omitempty"`

	// The compliance severity for the association. Can be one of the following: UNSPECIFIED, LOW, MEDIUM, HIGH or CRITICAL
	ComplianceSeverity *string `json:"complianceSeverity,omitempty" tf:"compliance_severity,omitempty"`

	// The document version you want to associate with the target(s). Can be a specific version or the default version.
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// The instance ID to apply an SSM document to. Use targets with key InstanceIds for document schema versions 2.0 and above. Use the targets attribute instead.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxConcurrency *string `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxErrors *string `json:"maxErrors,omitempty" tf:"max_errors,omitempty"`

	// An output location block. Output Location is documented below.
	OutputLocation []OutputLocationInitParameters `json:"outputLocation,omitempty" tf:"output_location,omitempty"`

	// A block of arbitrary string parameters to pass to the SSM document.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A cron or rate expression that specifies when the association runs.
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`

	// A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
	Targets []TargetsInitParameters `json:"targets,omitempty" tf:"targets,omitempty"`

	// The number of seconds to wait for the association status to be Success. If Success status is not reached within the given time, create opration will fail.
	WaitForSuccessTimeoutSeconds *float64 `json:"waitForSuccessTimeoutSeconds,omitempty" tf:"wait_for_success_timeout_seconds,omitempty"`
}

type AssociationObservation struct {

	// By default, when you create a new or update associations, the system runs it immediately and then according to the schedule you specified. Enable this option if you do not want an association to run immediately after you create or update it. This parameter is not supported for rate expressions. Default: false.
	ApplyOnlyAtCronInterval *bool `json:"applyOnlyAtCronInterval,omitempty" tf:"apply_only_at_cron_interval,omitempty"`

	// The ARN of the SSM association
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the SSM association.
	AssociationID *string `json:"associationId,omitempty" tf:"association_id,omitempty"`

	// The descriptive name for the association.
	AssociationName *string `json:"associationName,omitempty" tf:"association_name,omitempty"`

	// Specify the target for the association. This target is required for associations that use an Automation document and target resources by using rate controls. This should be set to the SSM document parameter that will define how your automation will branch out.
	AutomationTargetParameterName *string `json:"automationTargetParameterName,omitempty" tf:"automation_target_parameter_name,omitempty"`

	// The compliance severity for the association. Can be one of the following: UNSPECIFIED, LOW, MEDIUM, HIGH or CRITICAL
	ComplianceSeverity *string `json:"complianceSeverity,omitempty" tf:"compliance_severity,omitempty"`

	// The document version you want to associate with the target(s). Can be a specific version or the default version.
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance ID to apply an SSM document to. Use targets with key InstanceIds for document schema versions 2.0 and above. Use the targets attribute instead.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxConcurrency *string `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxErrors *string `json:"maxErrors,omitempty" tf:"max_errors,omitempty"`

	// The name of the SSM document to apply.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An output location block. Output Location is documented below.
	OutputLocation []OutputLocationObservation `json:"outputLocation,omitempty" tf:"output_location,omitempty"`

	// A block of arbitrary string parameters to pass to the SSM document.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A cron or rate expression that specifies when the association runs.
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`

	// A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
	Targets []TargetsObservation `json:"targets,omitempty" tf:"targets,omitempty"`

	// The number of seconds to wait for the association status to be Success. If Success status is not reached within the given time, create opration will fail.
	WaitForSuccessTimeoutSeconds *float64 `json:"waitForSuccessTimeoutSeconds,omitempty" tf:"wait_for_success_timeout_seconds,omitempty"`
}

type AssociationParameters struct {

	// By default, when you create a new or update associations, the system runs it immediately and then according to the schedule you specified. Enable this option if you do not want an association to run immediately after you create or update it. This parameter is not supported for rate expressions. Default: false.
	// +kubebuilder:validation:Optional
	ApplyOnlyAtCronInterval *bool `json:"applyOnlyAtCronInterval,omitempty" tf:"apply_only_at_cron_interval,omitempty"`

	// The descriptive name for the association.
	// +kubebuilder:validation:Optional
	AssociationName *string `json:"associationName,omitempty" tf:"association_name,omitempty"`

	// Specify the target for the association. This target is required for associations that use an Automation document and target resources by using rate controls. This should be set to the SSM document parameter that will define how your automation will branch out.
	// +kubebuilder:validation:Optional
	AutomationTargetParameterName *string `json:"automationTargetParameterName,omitempty" tf:"automation_target_parameter_name,omitempty"`

	// The compliance severity for the association. Can be one of the following: UNSPECIFIED, LOW, MEDIUM, HIGH or CRITICAL
	// +kubebuilder:validation:Optional
	ComplianceSeverity *string `json:"complianceSeverity,omitempty" tf:"compliance_severity,omitempty"`

	// The document version you want to associate with the target(s). Can be a specific version or the default version.
	// +kubebuilder:validation:Optional
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// The instance ID to apply an SSM document to. Use targets with key InstanceIds for document schema versions 2.0 and above. Use the targets attribute instead.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	// +kubebuilder:validation:Optional
	MaxConcurrency *string `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	// +kubebuilder:validation:Optional
	MaxErrors *string `json:"maxErrors,omitempty" tf:"max_errors,omitempty"`

	// The name of the SSM document to apply.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssm/v1beta1.Document
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Document in ssm to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Document in ssm to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// An output location block. Output Location is documented below.
	// +kubebuilder:validation:Optional
	OutputLocation []OutputLocationParameters `json:"outputLocation,omitempty" tf:"output_location,omitempty"`

	// A block of arbitrary string parameters to pass to the SSM document.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A cron or rate expression that specifies when the association runs.
	// +kubebuilder:validation:Optional
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`

	// A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
	// +kubebuilder:validation:Optional
	Targets []TargetsParameters `json:"targets,omitempty" tf:"targets,omitempty"`

	// The number of seconds to wait for the association status to be Success. If Success status is not reached within the given time, create opration will fail.
	// +kubebuilder:validation:Optional
	WaitForSuccessTimeoutSeconds *float64 `json:"waitForSuccessTimeoutSeconds,omitempty" tf:"wait_for_success_timeout_seconds,omitempty"`
}

type OutputLocationInitParameters struct {

	// The S3 bucket name.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// The S3 bucket prefix. Results stored in the root if not configured.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`

	// The S3 bucket region.
	S3Region *string `json:"s3Region,omitempty" tf:"s3_region,omitempty"`
}

type OutputLocationObservation struct {

	// The S3 bucket name.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// The S3 bucket prefix. Results stored in the root if not configured.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`

	// The S3 bucket region.
	S3Region *string `json:"s3Region,omitempty" tf:"s3_region,omitempty"`
}

type OutputLocationParameters struct {

	// The S3 bucket name.
	// +kubebuilder:validation:Optional
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// The S3 bucket prefix. Results stored in the root if not configured.
	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`

	// The S3 bucket region.
	// +kubebuilder:validation:Optional
	S3Region *string `json:"s3Region,omitempty" tf:"s3_region,omitempty"`
}

type TargetsInitParameters struct {

	// Either InstanceIds or tag:Tag Name to specify an EC2 tag.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A list of instance IDs or tag values. AWS currently limits this list size to one value.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TargetsObservation struct {

	// Either InstanceIds or tag:Tag Name to specify an EC2 tag.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A list of instance IDs or tag values. AWS currently limits this list size to one value.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TargetsParameters struct {

	// Either InstanceIds or tag:Tag Name to specify an EC2 tag.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A list of instance IDs or tag values. AWS currently limits this list size to one value.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

// AssociationSpec defines the desired state of Association
type AssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AssociationParameters `json:"forProvider"`
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
	InitProvider AssociationInitParameters `json:"initProvider,omitempty"`
}

// AssociationStatus defines the observed state of Association.
type AssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Association is the Schema for the Associations API. Associates an SSM Document to an instance or EC2 tag.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Association struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AssociationSpec   `json:"spec"`
	Status            AssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssociationList contains a list of Associations
type AssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Association `json:"items"`
}

// Repository type metadata.
var (
	Association_Kind             = "Association"
	Association_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Association_Kind}.String()
	Association_KindAPIVersion   = Association_Kind + "." + CRDGroupVersion.String()
	Association_GroupVersionKind = CRDGroupVersion.WithKind(Association_Kind)
)

func init() {
	SchemeBuilder.Register(&Association{}, &AssociationList{})
}
