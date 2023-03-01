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

type CopyStepDetailsDestinationFileLocationObservation struct {
}

type CopyStepDetailsDestinationFileLocationParameters struct {

	// Specifies the details for the EFS file being copied.
	// +kubebuilder:validation:Optional
	EFSFileLocation []DestinationFileLocationEFSFileLocationParameters `json:"efsFileLocation,omitempty" tf:"efs_file_location,omitempty"`

	// Specifies the details for the S3 file being copied.
	// +kubebuilder:validation:Optional
	S3FileLocation []DestinationFileLocationS3FileLocationParameters `json:"s3FileLocation,omitempty" tf:"s3_file_location,omitempty"`
}

type CopyStepDetailsObservation struct {
}

type CopyStepDetailsParameters struct {

	// Specifies the location for the file being copied. Use ${Transfer:username} in this field to parametrize the destination prefix by username.
	// +kubebuilder:validation:Optional
	DestinationFileLocation []DestinationFileLocationParameters `json:"destinationFileLocation,omitempty" tf:"destination_file_location,omitempty"`

	// The name of the step, used as an identifier.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE. Valid values are TRUE and FALSE.
	// +kubebuilder:validation:Optional
	OverwriteExisting *string `json:"overwriteExisting,omitempty" tf:"overwrite_existing,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	// +kubebuilder:validation:Optional
	SourceFileLocation *string `json:"sourceFileLocation,omitempty" tf:"source_file_location,omitempty"`
}

type CustomStepDetailsObservation struct {
}

type CustomStepDetailsParameters struct {

	// The name of the step, used as an identifier.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	// +kubebuilder:validation:Optional
	SourceFileLocation *string `json:"sourceFileLocation,omitempty" tf:"source_file_location,omitempty"`

	// The ARN for the lambda function that is being called.
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Timeout, in seconds, for the step.
	// +kubebuilder:validation:Optional
	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`
}

type DeleteStepDetailsObservation struct {
}

type DeleteStepDetailsParameters struct {

	// The name of the step, used as an identifier.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	// +kubebuilder:validation:Optional
	SourceFileLocation *string `json:"sourceFileLocation,omitempty" tf:"source_file_location,omitempty"`
}

type DestinationFileLocationEFSFileLocationObservation struct {
}

type DestinationFileLocationEFSFileLocationParameters struct {

	// The ID of the file system, assigned by Amazon EFS.
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// The pathname for the folder being used by a workflow.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type DestinationFileLocationObservation struct {
}

type DestinationFileLocationParameters struct {

	// Specifies the details for the EFS file being copied.
	// +kubebuilder:validation:Optional
	EFSFileLocation []EFSFileLocationParameters `json:"efsFileLocation,omitempty" tf:"efs_file_location,omitempty"`

	// Specifies the details for the S3 file being copied.
	// +kubebuilder:validation:Optional
	S3FileLocation []S3FileLocationParameters `json:"s3FileLocation,omitempty" tf:"s3_file_location,omitempty"`
}

type DestinationFileLocationS3FileLocationObservation struct {
}

type DestinationFileLocationS3FileLocationParameters struct {

	// Specifies the S3 bucket for the customer input file.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The name assigned to the file when it was created in S3. You use the object key to retrieve the object.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type EFSFileLocationObservation struct {
}

type EFSFileLocationParameters struct {

	// The ID of the file system, assigned by Amazon EFS.
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// The pathname for the folder being used by a workflow.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type OnExceptionStepsObservation struct {
}

type OnExceptionStepsParameters struct {

	// Details for a step that performs a file copy. See Copy Step Details below.
	// +kubebuilder:validation:Optional
	CopyStepDetails []CopyStepDetailsParameters `json:"copyStepDetails,omitempty" tf:"copy_step_details,omitempty"`

	// Details for a step that invokes a lambda function.
	// +kubebuilder:validation:Optional
	CustomStepDetails []CustomStepDetailsParameters `json:"customStepDetails,omitempty" tf:"custom_step_details,omitempty"`

	// Details for a step that deletes the file.
	// +kubebuilder:validation:Optional
	DeleteStepDetails []DeleteStepDetailsParameters `json:"deleteStepDetails,omitempty" tf:"delete_step_details,omitempty"`

	// Details for a step that creates one or more tags.
	// +kubebuilder:validation:Optional
	TagStepDetails []TagStepDetailsParameters `json:"tagStepDetails,omitempty" tf:"tag_step_details,omitempty"`

	// One of the following step types are supported. COPY, CUSTOM, DELETE, and TAG.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type S3FileLocationObservation struct {
}

type S3FileLocationParameters struct {

	// Specifies the S3 bucket for the customer input file.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The name assigned to the file when it was created in S3. You use the object key to retrieve the object.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type StepsCopyStepDetailsObservation struct {
}

type StepsCopyStepDetailsParameters struct {

	// Specifies the location for the file being copied. Use ${Transfer:username} in this field to parametrize the destination prefix by username.
	// +kubebuilder:validation:Optional
	DestinationFileLocation []CopyStepDetailsDestinationFileLocationParameters `json:"destinationFileLocation,omitempty" tf:"destination_file_location,omitempty"`

	// The name of the step, used as an identifier.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE. Valid values are TRUE and FALSE.
	// +kubebuilder:validation:Optional
	OverwriteExisting *string `json:"overwriteExisting,omitempty" tf:"overwrite_existing,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	// +kubebuilder:validation:Optional
	SourceFileLocation *string `json:"sourceFileLocation,omitempty" tf:"source_file_location,omitempty"`
}

type StepsCustomStepDetailsObservation struct {
}

type StepsCustomStepDetailsParameters struct {

	// The name of the step, used as an identifier.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	// +kubebuilder:validation:Optional
	SourceFileLocation *string `json:"sourceFileLocation,omitempty" tf:"source_file_location,omitempty"`

	// The ARN for the lambda function that is being called.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Reference to a Function in lambda to populate target.
	// +kubebuilder:validation:Optional
	TargetRef *v1.Reference `json:"targetRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate target.
	// +kubebuilder:validation:Optional
	TargetSelector *v1.Selector `json:"targetSelector,omitempty" tf:"-"`

	// Timeout, in seconds, for the step.
	// +kubebuilder:validation:Optional
	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`
}

type StepsDeleteStepDetailsObservation struct {
}

type StepsDeleteStepDetailsParameters struct {

	// The name of the step, used as an identifier.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	// +kubebuilder:validation:Optional
	SourceFileLocation *string `json:"sourceFileLocation,omitempty" tf:"source_file_location,omitempty"`
}

type StepsObservation struct {
}

type StepsParameters struct {

	// Details for a step that performs a file copy. See Copy Step Details below.
	// +kubebuilder:validation:Optional
	CopyStepDetails []StepsCopyStepDetailsParameters `json:"copyStepDetails,omitempty" tf:"copy_step_details,omitempty"`

	// Details for a step that invokes a lambda function.
	// +kubebuilder:validation:Optional
	CustomStepDetails []StepsCustomStepDetailsParameters `json:"customStepDetails,omitempty" tf:"custom_step_details,omitempty"`

	// Details for a step that deletes the file.
	// +kubebuilder:validation:Optional
	DeleteStepDetails []StepsDeleteStepDetailsParameters `json:"deleteStepDetails,omitempty" tf:"delete_step_details,omitempty"`

	// Details for a step that creates one or more tags.
	// +kubebuilder:validation:Optional
	TagStepDetails []StepsTagStepDetailsParameters `json:"tagStepDetails,omitempty" tf:"tag_step_details,omitempty"`

	// One of the following step types are supported. COPY, CUSTOM, DELETE, and TAG.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type StepsTagStepDetailsObservation struct {
}

type StepsTagStepDetailsParameters struct {

	// The name of the step, used as an identifier.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	// +kubebuilder:validation:Optional
	SourceFileLocation *string `json:"sourceFileLocation,omitempty" tf:"source_file_location,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags []TagStepDetailsTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TagStepDetailsObservation struct {
}

type TagStepDetailsParameters struct {

	// The name of the step, used as an identifier.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	// +kubebuilder:validation:Optional
	SourceFileLocation *string `json:"sourceFileLocation,omitempty" tf:"source_file_location,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TagStepDetailsTagsObservation struct {
}

type TagStepDetailsTagsParameters struct {

	// The name assigned to the file when it was created in S3. You use the object key to retrieve the object.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The value that corresponds to the key.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type TagsObservation struct {
}

type TagsParameters struct {

	// The name assigned to the file when it was created in S3. You use the object key to retrieve the object.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The value that corresponds to the key.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type WorkflowObservation struct {

	// The Workflow ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Workflow id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkflowParameters struct {

	// A textual description for the workflow.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the steps (actions) to take if errors are encountered during execution of the workflow. See Workflow Steps below.
	// +kubebuilder:validation:Optional
	OnExceptionSteps []OnExceptionStepsParameters `json:"onExceptionSteps,omitempty" tf:"on_exception_steps,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the details for the steps that are in the specified workflow. See Workflow Steps below.
	// +kubebuilder:validation:Optional
	Steps []StepsParameters `json:"steps,omitempty" tf:"steps,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkflowSpec defines the desired state of Workflow
type WorkflowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkflowParameters `json:"forProvider"`
}

// WorkflowStatus defines the observed state of Workflow.
type WorkflowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkflowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workflow is the Schema for the Workflows API. Provides a AWS Transfer Workflow resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.steps)",message="steps is a required parameter"
	Spec   WorkflowSpec   `json:"spec"`
	Status WorkflowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkflowList contains a list of Workflows
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workflow `json:"items"`
}

// Repository type metadata.
var (
	Workflow_Kind             = "Workflow"
	Workflow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workflow_Kind}.String()
	Workflow_KindAPIVersion   = Workflow_Kind + "." + CRDGroupVersion.String()
	Workflow_GroupVersionKind = CRDGroupVersion.WithKind(Workflow_Kind)
)

func init() {
	SchemeBuilder.Register(&Workflow{}, &WorkflowList{})
}
