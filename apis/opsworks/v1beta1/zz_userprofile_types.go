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

type UserProfileInitParameters struct {

	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement *bool `json:"allowSelfManagement,omitempty" tf:"allow_self_management,omitempty"`

	// The users public key
	SSHPublicKey *string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`

	// The ssh username, with witch this user wants to log in
	SSHUsername *string `json:"sshUsername,omitempty" tf:"ssh_username,omitempty"`
}

type UserProfileObservation struct {

	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement *bool `json:"allowSelfManagement,omitempty" tf:"allow_self_management,omitempty"`

	// Same value as user_arn
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The users public key
	SSHPublicKey *string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`

	// The ssh username, with witch this user wants to log in
	SSHUsername *string `json:"sshUsername,omitempty" tf:"ssh_username,omitempty"`

	// The user's IAM ARN
	UserArn *string `json:"userArn,omitempty" tf:"user_arn,omitempty"`
}

type UserProfileParameters struct {

	// Whether users can specify their own SSH public key through the My Settings page
	// +kubebuilder:validation:Optional
	AllowSelfManagement *bool `json:"allowSelfManagement,omitempty" tf:"allow_self_management,omitempty"`

	// The users public key
	// +kubebuilder:validation:Optional
	SSHPublicKey *string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`

	// The ssh username, with witch this user wants to log in
	// +kubebuilder:validation:Optional
	SSHUsername *string `json:"sshUsername,omitempty" tf:"ssh_username,omitempty"`

	// The user's IAM ARN
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	UserArn *string `json:"userArn,omitempty" tf:"user_arn,omitempty"`

	// Reference to a User in iam to populate userArn.
	// +kubebuilder:validation:Optional
	UserArnRef *v1.Reference `json:"userArnRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate userArn.
	// +kubebuilder:validation:Optional
	UserArnSelector *v1.Selector `json:"userArnSelector,omitempty" tf:"-"`
}

// UserProfileSpec defines the desired state of UserProfile
type UserProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserProfileParameters `json:"forProvider"`
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
	InitProvider UserProfileInitParameters `json:"initProvider,omitempty"`
}

// UserProfileStatus defines the observed state of UserProfile.
type UserProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserProfile is the Schema for the UserProfiles API. Provides an OpsWorks User Profile resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sshUsername) || (has(self.initProvider) && has(self.initProvider.sshUsername))",message="spec.forProvider.sshUsername is a required parameter"
	Spec   UserProfileSpec   `json:"spec"`
	Status UserProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserProfileList contains a list of UserProfiles
type UserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserProfile `json:"items"`
}

// Repository type metadata.
var (
	UserProfile_Kind             = "UserProfile"
	UserProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserProfile_Kind}.String()
	UserProfile_KindAPIVersion   = UserProfile_Kind + "." + CRDGroupVersion.String()
	UserProfile_GroupVersionKind = CRDGroupVersion.WithKind(UserProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&UserProfile{}, &UserProfileList{})
}
