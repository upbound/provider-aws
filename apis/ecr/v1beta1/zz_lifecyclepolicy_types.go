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

type LifecyclePolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The policy document. This is a JSON formatted string. See more details about Policy Parameters in the official AWS docs.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The registry ID where the repository was created.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// Name of the repository to apply the policy.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`
}

type LifecyclePolicyParameters struct {

	// The policy document. This is a JSON formatted string. See more details about Policy Parameters in the official AWS docs.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the repository to apply the policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ecr/v1beta1.Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in ecr to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in ecr to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`
}

// LifecyclePolicySpec defines the desired state of LifecyclePolicy
type LifecyclePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LifecyclePolicyParameters `json:"forProvider"`
}

// LifecyclePolicyStatus defines the observed state of LifecyclePolicy.
type LifecyclePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LifecyclePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LifecyclePolicy is the Schema for the LifecyclePolicys API. Manages an ECR repository lifecycle policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.policy)",message="policy is a required parameter"
	Spec   LifecyclePolicySpec   `json:"spec"`
	Status LifecyclePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LifecyclePolicyList contains a list of LifecyclePolicys
type LifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LifecyclePolicy `json:"items"`
}

// Repository type metadata.
var (
	LifecyclePolicy_Kind             = "LifecyclePolicy"
	LifecyclePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LifecyclePolicy_Kind}.String()
	LifecyclePolicy_KindAPIVersion   = LifecyclePolicy_Kind + "." + CRDGroupVersion.String()
	LifecyclePolicy_GroupVersionKind = CRDGroupVersion.WithKind(LifecyclePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LifecyclePolicy{}, &LifecyclePolicyList{})
}
