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

type CloudwatchLogsObservation struct {

	// The name of the log group where the project stores evaluation events.
	LogGroup *string `json:"logGroup,omitempty" tf:"log_group,omitempty"`
}

type CloudwatchLogsParameters struct {

	// The name of the log group where the project stores evaluation events.
	// +kubebuilder:validation:Optional
	LogGroup *string `json:"logGroup,omitempty" tf:"log_group,omitempty"`
}

type DataDeliveryObservation struct {

	// A block that defines the CloudWatch Log Group that stores the evaluation events. See below.
	CloudwatchLogs []CloudwatchLogsObservation `json:"cloudwatchLogs,omitempty" tf:"cloudwatch_logs,omitempty"`

	// A block that defines the S3 bucket and prefix that stores the evaluation events. See below.
	S3Destination []S3DestinationObservation `json:"s3Destination,omitempty" tf:"s3_destination,omitempty"`
}

type DataDeliveryParameters struct {

	// A block that defines the CloudWatch Log Group that stores the evaluation events. See below.
	// +kubebuilder:validation:Optional
	CloudwatchLogs []CloudwatchLogsParameters `json:"cloudwatchLogs,omitempty" tf:"cloudwatch_logs,omitempty"`

	// A block that defines the S3 bucket and prefix that stores the evaluation events. See below.
	// +kubebuilder:validation:Optional
	S3Destination []S3DestinationParameters `json:"s3Destination,omitempty" tf:"s3_destination,omitempty"`
}

type ProjectObservation struct {

	// The number of ongoing experiments currently in the project.
	ActiveExperimentCount *float64 `json:"activeExperimentCount,omitempty" tf:"active_experiment_count,omitempty"`

	// The number of ongoing launches currently in the project.
	ActiveLaunchCount *float64 `json:"activeLaunchCount,omitempty" tf:"active_launch_count,omitempty"`

	// The ARN of the project.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date and time that the project is created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// A block that contains information about where Evidently is to store evaluation events for longer term storage, if you choose to do so. If you choose not to store these events, Evidently deletes them after using them to produce metrics and other experiment results that you can view. See below.
	DataDelivery []DataDeliveryObservation `json:"dataDelivery,omitempty" tf:"data_delivery,omitempty"`

	// Specifies the description of the project.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The number of experiments currently in the project. This includes all experiments that have been created and not deleted, whether they are ongoing or not.
	ExperimentCount *float64 `json:"experimentCount,omitempty" tf:"experiment_count,omitempty"`

	// The number of features currently in the project.
	FeatureCount *float64 `json:"featureCount,omitempty" tf:"feature_count,omitempty"`

	// The ID has the same value as the arn of the project.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time that the project was most recently updated.
	LastUpdatedTime *string `json:"lastUpdatedTime,omitempty" tf:"last_updated_time,omitempty"`

	// The number of launches currently in the project. This includes all launches that have been created and not deleted, whether they are ongoing or not.
	LaunchCount *float64 `json:"launchCount,omitempty" tf:"launch_count,omitempty"`

	// A name for the project.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The current state of the project. Valid values are AVAILABLE and UPDATING.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ProjectParameters struct {

	// A block that contains information about where Evidently is to store evaluation events for longer term storage, if you choose to do so. If you choose not to store these events, Evidently deletes them after using them to produce metrics and other experiment results that you can view. See below.
	// +kubebuilder:validation:Optional
	DataDelivery []DataDeliveryParameters `json:"dataDelivery,omitempty" tf:"data_delivery,omitempty"`

	// Specifies the description of the project.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A name for the project.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type S3DestinationObservation struct {

	// The name of the bucket in which Evidently stores evaluation events.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The bucket prefix in which Evidently stores evaluation events.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3DestinationParameters struct {

	// The name of the bucket in which Evidently stores evaluation events.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The bucket prefix in which Evidently stores evaluation events.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Project is the Schema for the Projects API. Provides a CloudWatch Evidently Project resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   ProjectSpec   `json:"spec"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
