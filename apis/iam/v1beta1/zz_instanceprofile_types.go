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

type InstanceProfileObservation struct {

	// ARN assigned by AWS to the instance profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Creation timestamp of the instance profile.
	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	// Instance profile's ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Path to the instance profile. For more information about paths, see IAM Identifiers in the IAM User Guide. Can be a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. Can include any ASCII character from the ! (\u0021) through the DEL character (\u007F), including most punctuation characters, digits, and upper and lowercase letters.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Name of the role to add to the profile.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Unique ID assigned by AWS.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type InstanceProfileParameters struct {

	// Path to the instance profile. For more information about paths, see IAM Identifiers in the IAM User Guide. Can be a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. Can include any ASCII character from the ! (\u0021) through the DEL character (\u007F), including most punctuation characters, digits, and upper and lowercase letters.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Name of the role to add to the profile.
	// +crossplane:generate:reference:type=Role
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// InstanceProfileSpec defines the desired state of InstanceProfile
type InstanceProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceProfileParameters `json:"forProvider"`
}

// InstanceProfileStatus defines the observed state of InstanceProfile.
type InstanceProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceProfile is the Schema for the InstanceProfiles API. Provides an IAM instance profile.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstanceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceProfileSpec   `json:"spec"`
	Status            InstanceProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceProfileList contains a list of InstanceProfiles
type InstanceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceProfile `json:"items"`
}

// Repository type metadata.
var (
	InstanceProfile_Kind             = "InstanceProfile"
	InstanceProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceProfile_Kind}.String()
	InstanceProfile_KindAPIVersion   = InstanceProfile_Kind + "." + CRDGroupVersion.String()
	InstanceProfile_GroupVersionKind = CRDGroupVersion.WithKind(InstanceProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceProfile{}, &InstanceProfileList{})
}
