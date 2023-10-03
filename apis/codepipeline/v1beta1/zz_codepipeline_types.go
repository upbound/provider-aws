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

type ActionInitParameters struct {

	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are Approval, Build, Deploy, Invoke, Source and Test.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// A map of the action declaration's configuration. Configurations options for action types and providers can be found in the Pipeline Structure Reference and Action Structure Reference documentation.
	Configuration map[string]*string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of artifact names to be worked on.
	InputArtifacts []*string `json:"inputArtifacts,omitempty" tf:"input_artifacts,omitempty"`

	// The action declaration's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace all output variables will be accessed from.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A list of artifact names to output. Output artifact names must be unique within a pipeline.
	OutputArtifacts []*string `json:"outputArtifacts,omitempty" tf:"output_artifacts,omitempty"`

	// The creator of the action being called. Possible values are AWS, Custom and ThirdParty.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The provider of the service being called by the action. Valid providers are determined by the action category. Provider names are listed in the Action Structure Reference documentation.
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The order in which actions are run.
	RunOrder *float64 `json:"runOrder,omitempty" tf:"run_order,omitempty"`

	// A string that identifies the action type.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ActionObservation struct {

	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are Approval, Build, Deploy, Invoke, Source and Test.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// A map of the action declaration's configuration. Configurations options for action types and providers can be found in the Pipeline Structure Reference and Action Structure Reference documentation.
	Configuration map[string]*string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of artifact names to be worked on.
	InputArtifacts []*string `json:"inputArtifacts,omitempty" tf:"input_artifacts,omitempty"`

	// The action declaration's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace all output variables will be accessed from.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A list of artifact names to output. Output artifact names must be unique within a pipeline.
	OutputArtifacts []*string `json:"outputArtifacts,omitempty" tf:"output_artifacts,omitempty"`

	// The creator of the action being called. Possible values are AWS, Custom and ThirdParty.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The provider of the service being called by the action. Valid providers are determined by the action category. Provider names are listed in the Action Structure Reference documentation.
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// The region in which to run the action.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The order in which actions are run.
	RunOrder *float64 `json:"runOrder,omitempty" tf:"run_order,omitempty"`

	// A string that identifies the action type.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ActionParameters struct {

	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are Approval, Build, Deploy, Invoke, Source and Test.
	// +kubebuilder:validation:Optional
	Category *string `json:"category" tf:"category,omitempty"`

	// A map of the action declaration's configuration. Configurations options for action types and providers can be found in the Pipeline Structure Reference and Action Structure Reference documentation.
	// +kubebuilder:validation:Optional
	Configuration map[string]*string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of artifact names to be worked on.
	// +kubebuilder:validation:Optional
	InputArtifacts []*string `json:"inputArtifacts,omitempty" tf:"input_artifacts,omitempty"`

	// The action declaration's name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The namespace all output variables will be accessed from.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A list of artifact names to output. Output artifact names must be unique within a pipeline.
	// +kubebuilder:validation:Optional
	OutputArtifacts []*string `json:"outputArtifacts,omitempty" tf:"output_artifacts,omitempty"`

	// The creator of the action being called. Possible values are AWS, Custom and ThirdParty.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// The provider of the service being called by the action. Valid providers are determined by the action category. Provider names are listed in the Action Structure Reference documentation.
	// +kubebuilder:validation:Optional
	Provider *string `json:"provider" tf:"provider,omitempty"`

	// The region in which to run the action.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The order in which actions are run.
	// +kubebuilder:validation:Optional
	RunOrder *float64 `json:"runOrder,omitempty" tf:"run_order,omitempty"`

	// A string that identifies the action type.
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

type ArtifactStoreInitParameters struct {

	// The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An encryption_key block is documented below.
	EncryptionKey []EncryptionKeyInitParameters `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// The type of the artifact store, such as Amazon S3
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ArtifactStoreObservation struct {

	// The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An encryption_key block is documented below.
	EncryptionKey []EncryptionKeyObservation `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// The location where AWS CodePipeline stores artifacts for a pipeline; currently only S3 is supported.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of the artifact store, such as Amazon S3
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ArtifactStoreParameters struct {

	// The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An encryption_key block is documented below.
	// +kubebuilder:validation:Optional
	EncryptionKey []EncryptionKeyParameters `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// The location where AWS CodePipeline stores artifacts for a pipeline; currently only S3 is supported.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Reference to a Bucket in s3 to populate location.
	// +kubebuilder:validation:Optional
	LocationRef *v1.Reference `json:"locationRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate location.
	// +kubebuilder:validation:Optional
	LocationSelector *v1.Selector `json:"locationSelector,omitempty" tf:"-"`

	// The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of the artifact store, such as Amazon S3
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type CodepipelineInitParameters struct {

	// One or more artifact_store blocks. Artifact stores are documented below.
	ArtifactStore []ArtifactStoreInitParameters `json:"artifactStore,omitempty" tf:"artifact_store,omitempty"`

	// (Minimum of at least two stage blocks is required) A stage block. Stages are documented below.
	Stage []StageInitParameters `json:"stage,omitempty" tf:"stage,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CodepipelineObservation struct {

	// The codepipeline ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// One or more artifact_store blocks. Artifact stores are documented below.
	ArtifactStore []ArtifactStoreObservation `json:"artifactStore,omitempty" tf:"artifact_store,omitempty"`

	// The codepipeline ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// (Minimum of at least two stage blocks is required) A stage block. Stages are documented below.
	Stage []StageObservation `json:"stage,omitempty" tf:"stage,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CodepipelineParameters struct {

	// One or more artifact_store blocks. Artifact stores are documented below.
	// +kubebuilder:validation:Optional
	ArtifactStore []ArtifactStoreParameters `json:"artifactStore,omitempty" tf:"artifact_store,omitempty"`

	// The region in which to run the action.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
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

	// (Minimum of at least two stage blocks is required) A stage block. Stages are documented below.
	// +kubebuilder:validation:Optional
	Stage []StageParameters `json:"stage,omitempty" tf:"stage,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EncryptionKeyInitParameters struct {

	// The KMS key ARN or ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of key; currently only KMS is supported
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EncryptionKeyObservation struct {

	// The KMS key ARN or ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of key; currently only KMS is supported
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EncryptionKeyParameters struct {

	// The KMS key ARN or ID
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// The type of key; currently only KMS is supported
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type StageInitParameters struct {

	// The action(s) to include in the stage. Defined as an action block below
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// The name of the stage.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StageObservation struct {

	// The action(s) to include in the stage. Defined as an action block below
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// The name of the stage.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StageParameters struct {

	// The action(s) to include in the stage. Defined as an action block below
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action" tf:"action,omitempty"`

	// The name of the stage.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// CodepipelineSpec defines the desired state of Codepipeline
type CodepipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CodepipelineParameters `json:"forProvider"`
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
	InitProvider CodepipelineInitParameters `json:"initProvider,omitempty"`
}

// CodepipelineStatus defines the observed state of Codepipeline.
type CodepipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CodepipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Codepipeline is the Schema for the Codepipelines API. Provides a CodePipeline
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Codepipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.artifactStore) || (has(self.initProvider) && has(self.initProvider.artifactStore))",message="spec.forProvider.artifactStore is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stage) || (has(self.initProvider) && has(self.initProvider.stage))",message="spec.forProvider.stage is a required parameter"
	Spec   CodepipelineSpec   `json:"spec"`
	Status CodepipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodepipelineList contains a list of Codepipelines
type CodepipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Codepipeline `json:"items"`
}

// Repository type metadata.
var (
	Codepipeline_Kind             = "Codepipeline"
	Codepipeline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Codepipeline_Kind}.String()
	Codepipeline_KindAPIVersion   = Codepipeline_Kind + "." + CRDGroupVersion.String()
	Codepipeline_GroupVersionKind = CRDGroupVersion.WithKind(Codepipeline_Kind)
)

func init() {
	SchemeBuilder.Register(&Codepipeline{}, &CodepipelineList{})
}
