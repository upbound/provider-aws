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

type InstanceObservation struct {

	// Amazon Resource Name (ARN) of the instance.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// When the instance was created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// The identifier of the instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The service role of the instance.
	ServiceRole *string `json:"serviceRole,omitempty" tf:"service_role,omitempty"`

	// The state of the instance.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type InstanceParameters struct {

	// Specifies whether auto resolve best voices is enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	AutoResolveBestVoicesEnabled *bool `json:"autoResolveBestVoicesEnabled,omitempty" tf:"auto_resolve_best_voices_enabled,omitempty"`

	// Specifies whether contact flow logs are enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	ContactFlowLogsEnabled *bool `json:"contactFlowLogsEnabled,omitempty" tf:"contact_flow_logs_enabled,omitempty"`

	// Specifies whether contact lens is enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	ContactLensEnabled *bool `json:"contactLensEnabled,omitempty" tf:"contact_lens_enabled,omitempty"`

	// The identifier for the directory if identity_management_type is EXISTING_DIRECTORY.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ds/v1beta1.Directory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Reference to a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDRef *v1.Reference `json:"directoryIdRef,omitempty" tf:"-"`

	// Selector for a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDSelector *v1.Selector `json:"directoryIdSelector,omitempty" tf:"-"`

	// Specifies whether early media for outbound calls is enabled . Defaults to true if outbound calls is enabled.
	// +kubebuilder:validation:Optional
	EarlyMediaEnabled *bool `json:"earlyMediaEnabled,omitempty" tf:"early_media_enabled,omitempty"`

	// Specifies the identity management type attached to the instance. Allowed Values are: SAML, CONNECT_MANAGED, EXISTING_DIRECTORY.
	// +kubebuilder:validation:Optional
	IdentityManagementType *string `json:"identityManagementType,omitempty" tf:"identity_management_type,omitempty"`

	// Specifies whether inbound calls are enabled.
	// +kubebuilder:validation:Optional
	InboundCallsEnabled *bool `json:"inboundCallsEnabled,omitempty" tf:"inbound_calls_enabled,omitempty"`

	// Specifies the name of the instance. Required if directory_id not specified.
	// +kubebuilder:validation:Optional
	InstanceAlias *string `json:"instanceAlias,omitempty" tf:"instance_alias,omitempty"`

	// Specifies whether multi-party calls/conference is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	MultiPartyConferenceEnabled *bool `json:"multiPartyConferenceEnabled,omitempty" tf:"multi_party_conference_enabled,omitempty"`

	// Specifies whether outbound calls are enabled.
	// +kubebuilder:validation:Optional
	OutboundCallsEnabled *bool `json:"outboundCallsEnabled,omitempty" tf:"outbound_calls_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. Provides details about a specific Connect Instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.identityManagementType)",message="identityManagementType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.inboundCallsEnabled)",message="inboundCallsEnabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.outboundCallsEnabled)",message="outboundCallsEnabled is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
