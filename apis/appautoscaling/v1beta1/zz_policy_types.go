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

type CustomizedMetricSpecificationObservation struct {

	// Configuration block(s) with the dimensions of the metric if the metric was published with dimensions. Detailed below.
	Dimensions []DimensionsObservation `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Statistic of the metric. Valid values: Average, Minimum, Maximum, SampleCount, and Sum.
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Unit of the metric.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CustomizedMetricSpecificationParameters struct {

	// Configuration block(s) with the dimensions of the metric if the metric was published with dimensions. Detailed below.
	// +kubebuilder:validation:Optional
	Dimensions []DimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// Statistic of the metric. Valid values: Average, Minimum, Maximum, SampleCount, and Sum.
	// +kubebuilder:validation:Required
	Statistic *string `json:"statistic" tf:"statistic,omitempty"`

	// Unit of the metric.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type DimensionsObservation struct {

	// Name of the dimension.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Value of the dimension.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DimensionsParameters struct {

	// Name of the dimension.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Value of the dimension.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PolicyObservation struct {

	// List of CloudWatch alarm ARNs associated with the scaling policy.
	AlarmArns []*string `json:"alarmArns,omitempty" tf:"alarm_arns,omitempty"`

	// ARN assigned by AWS to the scaling policy.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Policy type. Valid values are StepScaling and TargetTrackingScaling. Defaults to StepScaling. Certain services only support only one policy type. For more information see the Target Tracking Scaling Policies and Step Scaling Policies documentation.
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the ResourceId parameter at: AWS Application Auto Scaling API Reference
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Scalable dimension of the scalable target. Documentation can be found in the ScalableDimension parameter at: AWS Application Auto Scaling API Reference
	ScalableDimension *string `json:"scalableDimension,omitempty" tf:"scalable_dimension,omitempty"`

	// AWS service namespace of the scalable target. Documentation can be found in the ServiceNamespace parameter at: AWS Application Auto Scaling API Reference
	ServiceNamespace *string `json:"serviceNamespace,omitempty" tf:"service_namespace,omitempty"`

	// Step scaling policy configuration, requires policy_type = "StepScaling" (default). See supported fields below.
	StepScalingPolicyConfiguration []StepScalingPolicyConfigurationObservation `json:"stepScalingPolicyConfiguration,omitempty" tf:"step_scaling_policy_configuration,omitempty"`

	// Target tracking policy, requires policy_type = "TargetTrackingScaling". See supported fields below.
	TargetTrackingScalingPolicyConfiguration []TargetTrackingScalingPolicyConfigurationObservation `json:"targetTrackingScalingPolicyConfiguration,omitempty" tf:"target_tracking_scaling_policy_configuration,omitempty"`
}

type PolicyParameters struct {

	// Policy type. Valid values are StepScaling and TargetTrackingScaling. Defaults to StepScaling. Certain services only support only one policy type. For more information see the Target Tracking Scaling Policies and Step Scaling Policies documentation.
	// +kubebuilder:validation:Optional
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the ResourceId parameter at: AWS Application Auto Scaling API Reference
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appautoscaling/v1beta1.Target
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("resource_id",false)
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Target in appautoscaling to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Target in appautoscaling to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// Scalable dimension of the scalable target. Documentation can be found in the ScalableDimension parameter at: AWS Application Auto Scaling API Reference
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appautoscaling/v1beta1.Target
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("scalable_dimension",false)
	// +kubebuilder:validation:Optional
	ScalableDimension *string `json:"scalableDimension,omitempty" tf:"scalable_dimension,omitempty"`

	// Reference to a Target in appautoscaling to populate scalableDimension.
	// +kubebuilder:validation:Optional
	ScalableDimensionRef *v1.Reference `json:"scalableDimensionRef,omitempty" tf:"-"`

	// Selector for a Target in appautoscaling to populate scalableDimension.
	// +kubebuilder:validation:Optional
	ScalableDimensionSelector *v1.Selector `json:"scalableDimensionSelector,omitempty" tf:"-"`

	// AWS service namespace of the scalable target. Documentation can be found in the ServiceNamespace parameter at: AWS Application Auto Scaling API Reference
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appautoscaling/v1beta1.Target
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("service_namespace",false)
	// +kubebuilder:validation:Optional
	ServiceNamespace *string `json:"serviceNamespace,omitempty" tf:"service_namespace,omitempty"`

	// Reference to a Target in appautoscaling to populate serviceNamespace.
	// +kubebuilder:validation:Optional
	ServiceNamespaceRef *v1.Reference `json:"serviceNamespaceRef,omitempty" tf:"-"`

	// Selector for a Target in appautoscaling to populate serviceNamespace.
	// +kubebuilder:validation:Optional
	ServiceNamespaceSelector *v1.Selector `json:"serviceNamespaceSelector,omitempty" tf:"-"`

	// Step scaling policy configuration, requires policy_type = "StepScaling" (default). See supported fields below.
	// +kubebuilder:validation:Optional
	StepScalingPolicyConfiguration []StepScalingPolicyConfigurationParameters `json:"stepScalingPolicyConfiguration,omitempty" tf:"step_scaling_policy_configuration,omitempty"`

	// Target tracking policy, requires policy_type = "TargetTrackingScaling". See supported fields below.
	// +kubebuilder:validation:Optional
	TargetTrackingScalingPolicyConfiguration []TargetTrackingScalingPolicyConfigurationParameters `json:"targetTrackingScalingPolicyConfiguration,omitempty" tf:"target_tracking_scaling_policy_configuration,omitempty"`
}

type PredefinedMetricSpecificationObservation struct {

	// Metric type.
	PredefinedMetricType *string `json:"predefinedMetricType,omitempty" tf:"predefined_metric_type,omitempty"`

	// Reserved for future use if the predefined_metric_type is not ALBRequestCountPerTarget. If the predefined_metric_type is ALBRequestCountPerTarget, you must specify this argument. Documentation can be found at: AWS Predefined Scaling Metric Specification. Must be less than or equal to 1023 characters in length.
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type PredefinedMetricSpecificationParameters struct {

	// Metric type.
	// +kubebuilder:validation:Required
	PredefinedMetricType *string `json:"predefinedMetricType" tf:"predefined_metric_type,omitempty"`

	// Reserved for future use if the predefined_metric_type is not ALBRequestCountPerTarget. If the predefined_metric_type is ALBRequestCountPerTarget, you must specify this argument. Documentation can be found at: AWS Predefined Scaling Metric Specification. Must be less than or equal to 1023 characters in length.
	// +kubebuilder:validation:Optional
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type StepAdjustmentObservation struct {

	// Lower bound for the difference between the alarm threshold and the CloudWatch metric. Without a value, AWS will treat this bound as negative infinity.
	MetricIntervalLowerBound *string `json:"metricIntervalLowerBound,omitempty" tf:"metric_interval_lower_bound,omitempty"`

	// Upper bound for the difference between the alarm threshold and the CloudWatch metric. Without a value, AWS will treat this bound as infinity. The upper bound must be greater than the lower bound.
	MetricIntervalUpperBound *string `json:"metricIntervalUpperBound,omitempty" tf:"metric_interval_upper_bound,omitempty"`

	// Number of members by which to scale, when the adjustment bounds are breached. A positive value scales up. A negative value scales down.
	ScalingAdjustment *float64 `json:"scalingAdjustment,omitempty" tf:"scaling_adjustment,omitempty"`
}

type StepAdjustmentParameters struct {

	// Lower bound for the difference between the alarm threshold and the CloudWatch metric. Without a value, AWS will treat this bound as negative infinity.
	// +kubebuilder:validation:Optional
	MetricIntervalLowerBound *string `json:"metricIntervalLowerBound,omitempty" tf:"metric_interval_lower_bound,omitempty"`

	// Upper bound for the difference between the alarm threshold and the CloudWatch metric. Without a value, AWS will treat this bound as infinity. The upper bound must be greater than the lower bound.
	// +kubebuilder:validation:Optional
	MetricIntervalUpperBound *string `json:"metricIntervalUpperBound,omitempty" tf:"metric_interval_upper_bound,omitempty"`

	// Number of members by which to scale, when the adjustment bounds are breached. A positive value scales up. A negative value scales down.
	// +kubebuilder:validation:Required
	ScalingAdjustment *float64 `json:"scalingAdjustment" tf:"scaling_adjustment,omitempty"`
}

type StepScalingPolicyConfigurationObservation struct {

	// Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are ChangeInCapacity, ExactCapacity, and PercentChangeInCapacity.
	AdjustmentType *string `json:"adjustmentType,omitempty" tf:"adjustment_type,omitempty"`

	// Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown *float64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// Aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType *string `json:"metricAggregationType,omitempty" tf:"metric_aggregation_type,omitempty"`

	// Minimum number to adjust your scalable dimension as a result of a scaling activity. If the adjustment type is PercentChangeInCapacity, the scaling policy changes the scalable dimension of the scalable target by this amount.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude,omitempty" tf:"min_adjustment_magnitude,omitempty"`

	// Set of adjustments that manage scaling. These have the following structure:
	StepAdjustment []StepAdjustmentObservation `json:"stepAdjustment,omitempty" tf:"step_adjustment,omitempty"`
}

type StepScalingPolicyConfigurationParameters struct {

	// Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are ChangeInCapacity, ExactCapacity, and PercentChangeInCapacity.
	// +kubebuilder:validation:Optional
	AdjustmentType *string `json:"adjustmentType,omitempty" tf:"adjustment_type,omitempty"`

	// Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	// +kubebuilder:validation:Optional
	Cooldown *float64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// Aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	// +kubebuilder:validation:Optional
	MetricAggregationType *string `json:"metricAggregationType,omitempty" tf:"metric_aggregation_type,omitempty"`

	// Minimum number to adjust your scalable dimension as a result of a scaling activity. If the adjustment type is PercentChangeInCapacity, the scaling policy changes the scalable dimension of the scalable target by this amount.
	// +kubebuilder:validation:Optional
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude,omitempty" tf:"min_adjustment_magnitude,omitempty"`

	// Set of adjustments that manage scaling. These have the following structure:
	// +kubebuilder:validation:Optional
	StepAdjustment []StepAdjustmentParameters `json:"stepAdjustment,omitempty" tf:"step_adjustment,omitempty"`
}

type TargetTrackingScalingPolicyConfigurationObservation struct {

	// Custom CloudWatch metric. Documentation can be found  at: AWS Customized Metric Specification. See supported fields below.
	CustomizedMetricSpecification []CustomizedMetricSpecificationObservation `json:"customizedMetricSpecification,omitempty" tf:"customized_metric_specification,omitempty"`

	// Whether scale in by the target tracking policy is disabled. If the value is true, scale in is disabled and the target tracking policy won't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking policy can remove capacity from the scalable resource. The default value is false.
	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`

	// Predefined metric. See supported fields below.
	PredefinedMetricSpecification []PredefinedMetricSpecificationObservation `json:"predefinedMetricSpecification,omitempty" tf:"predefined_metric_specification,omitempty"`

	// Amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown *float64 `json:"scaleInCooldown,omitempty" tf:"scale_in_cooldown,omitempty"`

	// Amount of time, in seconds, after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown *float64 `json:"scaleOutCooldown,omitempty" tf:"scale_out_cooldown,omitempty"`

	// Target value for the metric.
	TargetValue *float64 `json:"targetValue,omitempty" tf:"target_value,omitempty"`
}

type TargetTrackingScalingPolicyConfigurationParameters struct {

	// Custom CloudWatch metric. Documentation can be found  at: AWS Customized Metric Specification. See supported fields below.
	// +kubebuilder:validation:Optional
	CustomizedMetricSpecification []CustomizedMetricSpecificationParameters `json:"customizedMetricSpecification,omitempty" tf:"customized_metric_specification,omitempty"`

	// Whether scale in by the target tracking policy is disabled. If the value is true, scale in is disabled and the target tracking policy won't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking policy can remove capacity from the scalable resource. The default value is false.
	// +kubebuilder:validation:Optional
	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`

	// Predefined metric. See supported fields below.
	// +kubebuilder:validation:Optional
	PredefinedMetricSpecification []PredefinedMetricSpecificationParameters `json:"predefinedMetricSpecification,omitempty" tf:"predefined_metric_specification,omitempty"`

	// Amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
	// +kubebuilder:validation:Optional
	ScaleInCooldown *float64 `json:"scaleInCooldown,omitempty" tf:"scale_in_cooldown,omitempty"`

	// Amount of time, in seconds, after a scale out activity completes before another scale out activity can start.
	// +kubebuilder:validation:Optional
	ScaleOutCooldown *float64 `json:"scaleOutCooldown,omitempty" tf:"scale_out_cooldown,omitempty"`

	// Target value for the metric.
	// +kubebuilder:validation:Required
	TargetValue *float64 `json:"targetValue" tf:"target_value,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. Provides an Application AutoScaling Policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
