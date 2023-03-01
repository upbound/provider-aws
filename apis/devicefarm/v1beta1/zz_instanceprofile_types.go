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

	// The Amazon Resource Name of this instance profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type InstanceProfileParameters struct {

	// The description of the instance profile.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run.
	// +kubebuilder:validation:Optional
	ExcludeAppPackagesFromCleanup []*string `json:"excludeAppPackagesFromCleanup,omitempty" tf:"exclude_app_packages_from_cleanup,omitempty"`

	// The name for the instance profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// When set to true, Device Farm removes app packages after a test run. The default value is false for private devices.
	// +kubebuilder:validation:Optional
	PackageCleanup *bool `json:"packageCleanup,omitempty" tf:"package_cleanup,omitempty"`

	// When set to true, Device Farm reboots the instance after a test run. The default value is true.
	// +kubebuilder:validation:Optional
	RebootAfterUse *bool `json:"rebootAfterUse,omitempty" tf:"reboot_after_use,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

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

// InstanceProfile is the Schema for the InstanceProfiles API. Provides a Devicefarm instance profile
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstanceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   InstanceProfileSpec   `json:"spec"`
	Status InstanceProfileStatus `json:"status,omitempty"`
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
