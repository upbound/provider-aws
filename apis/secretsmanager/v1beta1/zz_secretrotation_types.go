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

type SecretRotationInitParameters struct {

	// A structure that defines the rotation configuration for this secret. Defined below.
	RotationRules []SecretRotationRotationRulesInitParameters `json:"rotationRules,omitempty" tf:"rotation_rules,omitempty"`
}

type SecretRotationObservation struct {

	// Amazon Resource Name (ARN) of the secret.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether automatic rotation is enabled for this secret.
	RotationEnabled *bool `json:"rotationEnabled,omitempty" tf:"rotation_enabled,omitempty"`

	// Specifies the ARN of the Lambda function that can rotate the secret.
	RotationLambdaArn *string `json:"rotationLambdaArn,omitempty" tf:"rotation_lambda_arn,omitempty"`

	// A structure that defines the rotation configuration for this secret. Defined below.
	RotationRules []SecretRotationRotationRulesObservation `json:"rotationRules,omitempty" tf:"rotation_rules,omitempty"`

	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`
}

type SecretRotationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the ARN of the Lambda function that can rotate the secret.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RotationLambdaArn *string `json:"rotationLambdaArn,omitempty" tf:"rotation_lambda_arn,omitempty"`

	// Reference to a Function in lambda to populate rotationLambdaArn.
	// +kubebuilder:validation:Optional
	RotationLambdaArnRef *v1.Reference `json:"rotationLambdaArnRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate rotationLambdaArn.
	// +kubebuilder:validation:Optional
	RotationLambdaArnSelector *v1.Selector `json:"rotationLambdaArnSelector,omitempty" tf:"-"`

	// A structure that defines the rotation configuration for this secret. Defined below.
	// +kubebuilder:validation:Optional
	RotationRules []SecretRotationRotationRulesParameters `json:"rotationRules,omitempty" tf:"rotation_rules,omitempty"`

	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`

	// Reference to a Secret in secretsmanager to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDRef *v1.Reference `json:"secretIdRef,omitempty" tf:"-"`

	// Selector for a Secret in secretsmanager to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDSelector *v1.Selector `json:"secretIdSelector,omitempty" tf:"-"`
}

type SecretRotationRotationRulesInitParameters struct {

	// Specifies the number of days between automatic scheduled rotations of the secret. Either automatically_after_days or schedule_expression must be specified.
	AutomaticallyAfterDays *float64 `json:"automaticallyAfterDays,omitempty" tf:"automatically_after_days,omitempty"`

	// - The length of the rotation window in hours. For example, 3h for a three hour window.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// A cron() or rate() expression that defines the schedule for rotating your secret. Either automatically_after_days or schedule_expression must be specified.
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`
}

type SecretRotationRotationRulesObservation struct {

	// Specifies the number of days between automatic scheduled rotations of the secret. Either automatically_after_days or schedule_expression must be specified.
	AutomaticallyAfterDays *float64 `json:"automaticallyAfterDays,omitempty" tf:"automatically_after_days,omitempty"`

	// - The length of the rotation window in hours. For example, 3h for a three hour window.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// A cron() or rate() expression that defines the schedule for rotating your secret. Either automatically_after_days or schedule_expression must be specified.
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`
}

type SecretRotationRotationRulesParameters struct {

	// Specifies the number of days between automatic scheduled rotations of the secret. Either automatically_after_days or schedule_expression must be specified.
	// +kubebuilder:validation:Optional
	AutomaticallyAfterDays *float64 `json:"automaticallyAfterDays,omitempty" tf:"automatically_after_days,omitempty"`

	// - The length of the rotation window in hours. For example, 3h for a three hour window.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// A cron() or rate() expression that defines the schedule for rotating your secret. Either automatically_after_days or schedule_expression must be specified.
	// +kubebuilder:validation:Optional
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`
}

// SecretRotationSpec defines the desired state of SecretRotation
type SecretRotationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretRotationParameters `json:"forProvider"`
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
	InitProvider SecretRotationInitParameters `json:"initProvider,omitempty"`
}

// SecretRotationStatus defines the observed state of SecretRotation.
type SecretRotationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretRotationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretRotation is the Schema for the SecretRotations API. Provides a resource to manage AWS Secrets Manager secret rotation
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecretRotation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rotationRules) || (has(self.initProvider) && has(self.initProvider.rotationRules))",message="spec.forProvider.rotationRules is a required parameter"
	Spec   SecretRotationSpec   `json:"spec"`
	Status SecretRotationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretRotationList contains a list of SecretRotations
type SecretRotationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretRotation `json:"items"`
}

// Repository type metadata.
var (
	SecretRotation_Kind             = "SecretRotation"
	SecretRotation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretRotation_Kind}.String()
	SecretRotation_KindAPIVersion   = SecretRotation_Kind + "." + CRDGroupVersion.String()
	SecretRotation_GroupVersionKind = CRDGroupVersion.WithKind(SecretRotation_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretRotation{}, &SecretRotationList{})
}
