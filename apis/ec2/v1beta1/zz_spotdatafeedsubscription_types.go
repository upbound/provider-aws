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

type SpotDatafeedSubscriptionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SpotDatafeedSubscriptionParameters struct {

	// The Amazon S3 bucket in which to store the Spot instance data feed.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Path of folder inside bucket to place spot pricing data.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// SpotDatafeedSubscriptionSpec defines the desired state of SpotDatafeedSubscription
type SpotDatafeedSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpotDatafeedSubscriptionParameters `json:"forProvider"`
}

// SpotDatafeedSubscriptionStatus defines the observed state of SpotDatafeedSubscription.
type SpotDatafeedSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpotDatafeedSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpotDatafeedSubscription is the Schema for the SpotDatafeedSubscriptions API. Provides a Spot Datafeed Subscription resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SpotDatafeedSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.bucket)",message="bucket is a required parameter"
	Spec   SpotDatafeedSubscriptionSpec   `json:"spec"`
	Status SpotDatafeedSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotDatafeedSubscriptionList contains a list of SpotDatafeedSubscriptions
type SpotDatafeedSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotDatafeedSubscription `json:"items"`
}

// Repository type metadata.
var (
	SpotDatafeedSubscription_Kind             = "SpotDatafeedSubscription"
	SpotDatafeedSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpotDatafeedSubscription_Kind}.String()
	SpotDatafeedSubscription_KindAPIVersion   = SpotDatafeedSubscription_Kind + "." + CRDGroupVersion.String()
	SpotDatafeedSubscription_GroupVersionKind = CRDGroupVersion.WithKind(SpotDatafeedSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&SpotDatafeedSubscription{}, &SpotDatafeedSubscriptionList{})
}
