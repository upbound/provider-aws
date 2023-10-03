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

type ScheduleInitParameters struct {

	// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to -1 if you don't want to change the desired capacity at the scheduled time. Defaults to 0.
	DesiredCapacity *float64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// The date and time for the recurring schedule to end, in UTC with the format "YYYY-MM-DDThh:mm:ssZ" (e.g. "2021-06-01T00:00:00Z").
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The maximum size of the Auto Scaling group. Set to -1 if you don't want to change the maximum size at the scheduled time. Defaults to 0.
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// The minimum size of the Auto Scaling group. Set to -1 if you don't want to change the minimum size at the scheduled time. Defaults to 0.
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// The recurring schedule for this action specified using the Unix cron syntax format.
	Recurrence *string `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// The date and time for the recurring schedule to start, in UTC with the format "YYYY-MM-DDThh:mm:ssZ" (e.g. "2021-06-01T00:00:00Z").
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as Etc/GMT+9 or Pacific/Tahiti).
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type ScheduleObservation struct {

	// ARN assigned by AWS to the autoscaling schedule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the Auto Scaling group.
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to -1 if you don't want to change the desired capacity at the scheduled time. Defaults to 0.
	DesiredCapacity *float64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// The date and time for the recurring schedule to end, in UTC with the format "YYYY-MM-DDThh:mm:ssZ" (e.g. "2021-06-01T00:00:00Z").
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum size of the Auto Scaling group. Set to -1 if you don't want to change the maximum size at the scheduled time. Defaults to 0.
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// The minimum size of the Auto Scaling group. Set to -1 if you don't want to change the minimum size at the scheduled time. Defaults to 0.
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// The recurring schedule for this action specified using the Unix cron syntax format.
	Recurrence *string `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// The date and time for the recurring schedule to start, in UTC with the format "YYYY-MM-DDThh:mm:ssZ" (e.g. "2021-06-01T00:00:00Z").
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as Etc/GMT+9 or Pacific/Tahiti).
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type ScheduleParameters struct {

	// The name of the Auto Scaling group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/autoscaling/v1beta1.AutoscalingGroup
	// +kubebuilder:validation:Optional
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// Reference to a AutoscalingGroup in autoscaling to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameRef *v1.Reference `json:"autoscalingGroupNameRef,omitempty" tf:"-"`

	// Selector for a AutoscalingGroup in autoscaling to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameSelector *v1.Selector `json:"autoscalingGroupNameSelector,omitempty" tf:"-"`

	// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to -1 if you don't want to change the desired capacity at the scheduled time. Defaults to 0.
	// +kubebuilder:validation:Optional
	DesiredCapacity *float64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// The date and time for the recurring schedule to end, in UTC with the format "YYYY-MM-DDThh:mm:ssZ" (e.g. "2021-06-01T00:00:00Z").
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The maximum size of the Auto Scaling group. Set to -1 if you don't want to change the maximum size at the scheduled time. Defaults to 0.
	// +kubebuilder:validation:Optional
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// The minimum size of the Auto Scaling group. Set to -1 if you don't want to change the minimum size at the scheduled time. Defaults to 0.
	// +kubebuilder:validation:Optional
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// The recurring schedule for this action specified using the Unix cron syntax format.
	// +kubebuilder:validation:Optional
	Recurrence *string `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The date and time for the recurring schedule to start, in UTC with the format "YYYY-MM-DDThh:mm:ssZ" (e.g. "2021-06-01T00:00:00Z").
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as Etc/GMT+9 or Pacific/Tahiti).
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

// ScheduleSpec defines the desired state of Schedule
type ScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduleParameters `json:"forProvider"`
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
	InitProvider ScheduleInitParameters `json:"initProvider,omitempty"`
}

// ScheduleStatus defines the observed state of Schedule.
type ScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schedule is the Schema for the Schedules API. Provides an AutoScaling Schedule resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Schedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduleSpec   `json:"spec"`
	Status            ScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduleList contains a list of Schedules
type ScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schedule `json:"items"`
}

// Repository type metadata.
var (
	Schedule_Kind             = "Schedule"
	Schedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Schedule_Kind}.String()
	Schedule_KindAPIVersion   = Schedule_Kind + "." + CRDGroupVersion.String()
	Schedule_GroupVersionKind = CRDGroupVersion.WithKind(Schedule_Kind)
)

func init() {
	SchemeBuilder.Register(&Schedule{}, &ScheduleList{})
}
