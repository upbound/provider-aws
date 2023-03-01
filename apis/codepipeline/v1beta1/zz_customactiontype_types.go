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

type ConfigurationPropertyObservation struct {
}

type ConfigurationPropertyParameters struct {

	// The description of the action configuration property.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the configuration property is a key.
	// +kubebuilder:validation:Required
	Key *bool `json:"key" tf:"key,omitempty"`

	// The name of the action configuration property.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Indicates that the property will be used in conjunction with PollForJobs.
	// +kubebuilder:validation:Optional
	Queryable *bool `json:"queryable,omitempty" tf:"queryable,omitempty"`

	// Whether the configuration property is a required value.
	// +kubebuilder:validation:Required
	Required *bool `json:"required" tf:"required,omitempty"`

	// Whether the configuration property is secret.
	// +kubebuilder:validation:Required
	Secret *bool `json:"secret" tf:"secret,omitempty"`

	// The type of the configuration property. Valid values: String, Number, Boolean
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CustomActionTypeObservation struct {

	// The action ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Composed of category, provider and version
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The creator of the action being called.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CustomActionTypeParameters struct {

	// The category of the custom action. Valid values: Source, Build, Deploy, Test, Invoke, Approval
	// +kubebuilder:validation:Optional
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The configuration properties for the custom action. Max 10 items.
	// +kubebuilder:validation:Optional
	ConfigurationProperty []ConfigurationPropertyParameters `json:"configurationProperty,omitempty" tf:"configuration_property,omitempty"`

	// The details of the input artifact for the action.
	// +kubebuilder:validation:Optional
	InputArtifactDetails []InputArtifactDetailsParameters `json:"inputArtifactDetails,omitempty" tf:"input_artifact_details,omitempty"`

	// The details of the output artifact of the action.
	// +kubebuilder:validation:Optional
	OutputArtifactDetails []OutputArtifactDetailsParameters `json:"outputArtifactDetails,omitempty" tf:"output_artifact_details,omitempty"`

	// The provider of the service used in the custom action
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The settings for an action type.
	// +kubebuilder:validation:Optional
	Settings []SettingsParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version identifier of the custom action.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type InputArtifactDetailsObservation struct {
}

type InputArtifactDetailsParameters struct {

	// The maximum number of artifacts allowed for the action type. Min: 0, Max: 5
	// +kubebuilder:validation:Required
	MaximumCount *float64 `json:"maximumCount" tf:"maximum_count,omitempty"`

	// The minimum number of artifacts allowed for the action type. Min: 0, Max: 5
	// +kubebuilder:validation:Required
	MinimumCount *float64 `json:"minimumCount" tf:"minimum_count,omitempty"`
}

type OutputArtifactDetailsObservation struct {
}

type OutputArtifactDetailsParameters struct {

	// The maximum number of artifacts allowed for the action type. Min: 0, Max: 5
	// +kubebuilder:validation:Required
	MaximumCount *float64 `json:"maximumCount" tf:"maximum_count,omitempty"`

	// The minimum number of artifacts allowed for the action type. Min: 0, Max: 5
	// +kubebuilder:validation:Required
	MinimumCount *float64 `json:"minimumCount" tf:"minimum_count,omitempty"`
}

type SettingsObservation struct {
}

type SettingsParameters struct {

	// The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system.
	// +kubebuilder:validation:Optional
	EntityURLTemplate *string `json:"entityUrlTemplate,omitempty" tf:"entity_url_template,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system.
	// +kubebuilder:validation:Optional
	ExecutionURLTemplate *string `json:"executionUrlTemplate,omitempty" tf:"execution_url_template,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action.
	// +kubebuilder:validation:Optional
	RevisionURLTemplate *string `json:"revisionUrlTemplate,omitempty" tf:"revision_url_template,omitempty"`

	// The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
	// +kubebuilder:validation:Optional
	ThirdPartyConfigurationURL *string `json:"thirdPartyConfigurationUrl,omitempty" tf:"third_party_configuration_url,omitempty"`
}

// CustomActionTypeSpec defines the desired state of CustomActionType
type CustomActionTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomActionTypeParameters `json:"forProvider"`
}

// CustomActionTypeStatus defines the observed state of CustomActionType.
type CustomActionTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomActionTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomActionType is the Schema for the CustomActionTypes API. Provides a CodePipeline CustomActionType.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CustomActionType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.category)",message="category is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.inputArtifactDetails)",message="inputArtifactDetails is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.outputArtifactDetails)",message="outputArtifactDetails is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.providerName)",message="providerName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.version)",message="version is a required parameter"
	Spec   CustomActionTypeSpec   `json:"spec"`
	Status CustomActionTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomActionTypeList contains a list of CustomActionTypes
type CustomActionTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomActionType `json:"items"`
}

// Repository type metadata.
var (
	CustomActionType_Kind             = "CustomActionType"
	CustomActionType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomActionType_Kind}.String()
	CustomActionType_KindAPIVersion   = CustomActionType_Kind + "." + CRDGroupVersion.String()
	CustomActionType_GroupVersionKind = CRDGroupVersion.WithKind(CustomActionType_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomActionType{}, &CustomActionTypeList{})
}
