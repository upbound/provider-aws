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

type UserGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UserGroupParameters struct {

	// The description of the user group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the user group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The precedence of the user group.
	// +kubebuilder:validation:Optional
	Precedence *float64 `json:"precedence,omitempty" tf:"precedence,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of the IAM role to be associated with the user group.
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

	// The user pool ID.
	// +crossplane:generate:reference:type=UserPool
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

// UserGroupSpec defines the desired state of UserGroup
type UserGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserGroupParameters `json:"forProvider"`
}

// UserGroupStatus defines the observed state of UserGroup.
type UserGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroup is the Schema for the UserGroups API. Provides a Cognito User Group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   UserGroupSpec   `json:"spec"`
	Status UserGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroupList contains a list of UserGroups
type UserGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserGroup `json:"items"`
}

// Repository type metadata.
var (
	UserGroup_Kind             = "UserGroup"
	UserGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserGroup_Kind}.String()
	UserGroup_KindAPIVersion   = UserGroup_Kind + "." + CRDGroupVersion.String()
	UserGroup_GroupVersionKind = CRDGroupVersion.WithKind(UserGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&UserGroup{}, &UserGroupList{})
}
