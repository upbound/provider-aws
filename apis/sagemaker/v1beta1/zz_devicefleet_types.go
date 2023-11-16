// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type DeviceFleetInitParameters struct {

	// A description of the fleet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
	EnableIotRoleAlias *bool `json:"enableIotRoleAlias,omitempty" tf:"enable_iot_role_alias,omitempty"`

	// Specifies details about the repository. see Output Config details below.
	OutputConfig []OutputConfigInitParameters `json:"outputConfig,omitempty" tf:"output_config,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DeviceFleetObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this Device Fleet.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A description of the fleet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
	EnableIotRoleAlias *bool `json:"enableIotRoleAlias,omitempty" tf:"enable_iot_role_alias,omitempty"`

	// The name of the Device Fleet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IotRoleAlias *string `json:"iotRoleAlias,omitempty" tf:"iot_role_alias,omitempty"`

	// Specifies details about the repository. see Output Config details below.
	OutputConfig []OutputConfigObservation `json:"outputConfig,omitempty" tf:"output_config,omitempty"`

	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DeviceFleetParameters struct {

	// A description of the fleet.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
	// +kubebuilder:validation:Optional
	EnableIotRoleAlias *bool `json:"enableIotRoleAlias,omitempty" tf:"enable_iot_role_alias,omitempty"`

	// Specifies details about the repository. see Output Config details below.
	// +kubebuilder:validation:Optional
	OutputConfig []OutputConfigParameters `json:"outputConfig,omitempty" tf:"output_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
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

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type OutputConfigInitParameters struct {

	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume after compilation job. If you don't provide a KMS key ID, Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The Amazon Simple Storage (S3) bucker URI.
	S3OutputLocation *string `json:"s3OutputLocation,omitempty" tf:"s3_output_location,omitempty"`
}

type OutputConfigObservation struct {

	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume after compilation job. If you don't provide a KMS key ID, Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The Amazon Simple Storage (S3) bucker URI.
	S3OutputLocation *string `json:"s3OutputLocation,omitempty" tf:"s3_output_location,omitempty"`
}

type OutputConfigParameters struct {

	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume after compilation job. If you don't provide a KMS key ID, Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The Amazon Simple Storage (S3) bucker URI.
	// +kubebuilder:validation:Optional
	S3OutputLocation *string `json:"s3OutputLocation" tf:"s3_output_location,omitempty"`
}

// DeviceFleetSpec defines the desired state of DeviceFleet
type DeviceFleetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceFleetParameters `json:"forProvider"`
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
	InitProvider DeviceFleetInitParameters `json:"initProvider,omitempty"`
}

// DeviceFleetStatus defines the observed state of DeviceFleet.
type DeviceFleetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceFleetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceFleet is the Schema for the DeviceFleets API. Provides a SageMaker Device Fleet resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws},path=devicefleet
type DeviceFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.outputConfig) || (has(self.initProvider) && has(self.initProvider.outputConfig))",message="spec.forProvider.outputConfig is a required parameter"
	Spec   DeviceFleetSpec   `json:"spec"`
	Status DeviceFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceFleetList contains a list of DeviceFleets
type DeviceFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeviceFleet `json:"items"`
}

// Repository type metadata.
var (
	DeviceFleet_Kind             = "DeviceFleet"
	DeviceFleet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeviceFleet_Kind}.String()
	DeviceFleet_KindAPIVersion   = DeviceFleet_Kind + "." + CRDGroupVersion.String()
	DeviceFleet_GroupVersionKind = CRDGroupVersion.WithKind(DeviceFleet_Kind)
)

func init() {
	SchemeBuilder.Register(&DeviceFleet{}, &DeviceFleetList{})
}
