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

type ControlInitParameters struct {

	// One or more input parameter blocks. An example of a control with two parameters is: "backup plan frequency is at least daily and the retention period is at least 1 year". The first parameter is daily. The second parameter is 1 year. Detailed below.
	InputParameter []InputParameterInitParameters `json:"inputParameter,omitempty" tf:"input_parameter,omitempty"`

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans. Detailed below.
	Scope []ScopeInitParameters `json:"scope,omitempty" tf:"scope,omitempty"`
}

type ControlObservation struct {

	// One or more input parameter blocks. An example of a control with two parameters is: "backup plan frequency is at least daily and the retention period is at least 1 year". The first parameter is daily. The second parameter is 1 year. Detailed below.
	InputParameter []InputParameterObservation `json:"inputParameter,omitempty" tf:"input_parameter,omitempty"`

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans. Detailed below.
	Scope []ScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`
}

type ControlParameters struct {

	// One or more input parameter blocks. An example of a control with two parameters is: "backup plan frequency is at least daily and the retention period is at least 1 year". The first parameter is daily. The second parameter is 1 year. Detailed below.
	// +kubebuilder:validation:Optional
	InputParameter []InputParameterParameters `json:"inputParameter,omitempty" tf:"input_parameter,omitempty"`

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans. Detailed below.
	// +kubebuilder:validation:Optional
	Scope []ScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`
}

type FrameworkInitParameters struct {

	// One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
	Control []ControlInitParameters `json:"control,omitempty" tf:"control,omitempty"`

	// The description of the framework with a maximum of 1,024 characters
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FrameworkObservation struct {

	// The ARN of the backup framework.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
	Control []ControlObservation `json:"control,omitempty" tf:"control,omitempty"`

	// The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC).
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The deployment status of a framework. The statuses are: CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED.
	DeploymentStatus *string `json:"deploymentStatus,omitempty" tf:"deployment_status,omitempty"`

	// The description of the framework with a maximum of 1,024 characters
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of the backup framework.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. For more information refer to the AWS documentation for Framework Status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType:granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FrameworkParameters struct {

	// One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
	// +kubebuilder:validation:Optional
	Control []ControlParameters `json:"control,omitempty" tf:"control,omitempty"`

	// The description of the framework with a maximum of 1,024 characters
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type InputParameterInitParameters struct {

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of parameter, for example, hourly.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type InputParameterObservation struct {

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of parameter, for example, hourly.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type InputParameterParameters struct {

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of parameter, for example, hourly.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ScopeInitParameters struct {

	// The ID of the only AWS resource that you want your control scope to contain. Minimum number of 1 item. Maximum number of 100 items.
	// +listType:set
	ComplianceResourceIds []*string `json:"complianceResourceIds,omitempty" tf:"compliance_resource_ids,omitempty"`

	// Describes whether the control scope includes one or more types of resources, such as EFS or RDS.
	// +listType:set
	ComplianceResourceTypes []*string `json:"complianceResourceTypes,omitempty" tf:"compliance_resource_types,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ScopeObservation struct {

	// The ID of the only AWS resource that you want your control scope to contain. Minimum number of 1 item. Maximum number of 100 items.
	// +listType:set
	ComplianceResourceIds []*string `json:"complianceResourceIds,omitempty" tf:"compliance_resource_ids,omitempty"`

	// Describes whether the control scope includes one or more types of resources, such as EFS or RDS.
	// +listType:set
	ComplianceResourceTypes []*string `json:"complianceResourceTypes,omitempty" tf:"compliance_resource_types,omitempty"`

	// Key-value map of resource tags.
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ScopeParameters struct {

	// The ID of the only AWS resource that you want your control scope to contain. Minimum number of 1 item. Maximum number of 100 items.
	// +kubebuilder:validation:Optional
	// +listType:set
	ComplianceResourceIds []*string `json:"complianceResourceIds,omitempty" tf:"compliance_resource_ids,omitempty"`

	// Describes whether the control scope includes one or more types of resources, such as EFS or RDS.
	// +kubebuilder:validation:Optional
	// +listType:set
	ComplianceResourceTypes []*string `json:"complianceResourceTypes,omitempty" tf:"compliance_resource_types,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType:granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// FrameworkSpec defines the desired state of Framework
type FrameworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrameworkParameters `json:"forProvider"`
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
	InitProvider FrameworkInitParameters `json:"initProvider,omitempty"`
}

// FrameworkStatus defines the observed state of Framework.
type FrameworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrameworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Framework is the Schema for the Frameworks API. Provides an AWS Backup Framework resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Framework struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.control) || (has(self.initProvider) && has(self.initProvider.control))",message="spec.forProvider.control is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   FrameworkSpec   `json:"spec"`
	Status FrameworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrameworkList contains a list of Frameworks
type FrameworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Framework `json:"items"`
}

// Repository type metadata.
var (
	Framework_Kind             = "Framework"
	Framework_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Framework_Kind}.String()
	Framework_KindAPIVersion   = Framework_Kind + "." + CRDGroupVersion.String()
	Framework_GroupVersionKind = CRDGroupVersion.WithKind(Framework_Kind)
)

func init() {
	SchemeBuilder.Register(&Framework{}, &FrameworkList{})
}
