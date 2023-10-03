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

type AccountAliasInitParameters struct {
}

type AccountAliasObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AccountAliasParameters struct {
}

// AccountAliasSpec defines the desired state of AccountAlias
type AccountAliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountAliasParameters `json:"forProvider"`
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
	InitProvider AccountAliasInitParameters `json:"initProvider,omitempty"`
}

// AccountAliasStatus defines the observed state of AccountAlias.
type AccountAliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountAlias is the Schema for the AccountAliass API. Manages the account alias for the AWS Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AccountAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountAliasSpec   `json:"spec"`
	Status            AccountAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountAliasList contains a list of AccountAliass
type AccountAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountAlias `json:"items"`
}

// Repository type metadata.
var (
	AccountAlias_Kind             = "AccountAlias"
	AccountAlias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountAlias_Kind}.String()
	AccountAlias_KindAPIVersion   = AccountAlias_Kind + "." + CRDGroupVersion.String()
	AccountAlias_GroupVersionKind = CRDGroupVersion.WithKind(AccountAlias_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountAlias{}, &AccountAliasList{})
}
