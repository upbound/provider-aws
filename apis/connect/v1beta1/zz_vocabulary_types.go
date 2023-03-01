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

type VocabularyObservation struct {

	// The Amazon Resource Name (ARN) of the vocabulary.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The reason why the custom vocabulary was not created.
	FailureReason *string `json:"failureReason,omitempty" tf:"failure_reason,omitempty"`

	// The identifier of the hosting Amazon Connect Instance and identifier of the vocabulary
	// separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp when the custom vocabulary was last modified.
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`

	// The current state of the custom vocabulary. Valid values are CREATION_IN_PROGRESS, ACTIVE, CREATION_FAILED, DELETE_IN_PROGRESS.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The identifier of the custom vocabulary.
	VocabularyID *string `json:"vocabularyId,omitempty" tf:"vocabulary_id,omitempty"`
}

type VocabularyParameters struct {

	// The content of the custom vocabulary in plain-text format with a table of values. Each row in the table represents a word or a phrase, described with Phrase, IPA, SoundsLike, and DisplayAs fields. Separate the fields with TAB characters. For more information, see Create a custom vocabulary using a table. Minimum length of 1. Maximum length of 60000.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The language code of the vocabulary entries. For a list of languages and their corresponding language codes, see What is Amazon Transcribe?. Valid Values are ar-AE, de-CH, de-DE, en-AB, en-AU, en-GB, en-IE, en-IN, en-US, en-WL, es-ES, es-US, fr-CA, fr-FR, hi-IN, it-IT, ja-JP, ko-KR, pt-BR, pt-PT, zh-CN.
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// A unique name of the custom vocabulary. Must not be more than 140 characters.
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

// VocabularySpec defines the desired state of Vocabulary
type VocabularySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VocabularyParameters `json:"forProvider"`
}

// VocabularyStatus defines the observed state of Vocabulary.
type VocabularyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VocabularyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vocabulary is the Schema for the Vocabularys API. Provides details about a specific Amazon Connect Vocabulary
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Vocabulary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.content)",message="content is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.languageCode)",message="languageCode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   VocabularySpec   `json:"spec"`
	Status VocabularyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VocabularyList contains a list of Vocabularys
type VocabularyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vocabulary `json:"items"`
}

// Repository type metadata.
var (
	Vocabulary_Kind             = "Vocabulary"
	Vocabulary_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vocabulary_Kind}.String()
	Vocabulary_KindAPIVersion   = Vocabulary_Kind + "." + CRDGroupVersion.String()
	Vocabulary_GroupVersionKind = CRDGroupVersion.WithKind(Vocabulary_Kind)
)

func init() {
	SchemeBuilder.Register(&Vocabulary{}, &VocabularyList{})
}
