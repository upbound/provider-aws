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

type HSMClientCertificateInitParameters struct {

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HSMClientCertificateObservation struct {

	// Amazon Resource Name (ARN) of the Hsm Client Certificate.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
	HSMClientCertificatePublicKey *string `json:"hsmClientCertificatePublicKey,omitempty" tf:"hsm_client_certificate_public_key,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type HSMClientCertificateParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// HSMClientCertificateSpec defines the desired state of HSMClientCertificate
type HSMClientCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HSMClientCertificateParameters `json:"forProvider"`
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
	InitProvider HSMClientCertificateInitParameters `json:"initProvider,omitempty"`
}

// HSMClientCertificateStatus defines the observed state of HSMClientCertificate.
type HSMClientCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HSMClientCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HSMClientCertificate is the Schema for the HSMClientCertificates API. Creates an HSM client certificate that an Amazon Redshift cluster will use to connect to the client's HSM in order to store and retrieve the keys used to encrypt the cluster databases.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HSMClientCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HSMClientCertificateSpec   `json:"spec"`
	Status            HSMClientCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HSMClientCertificateList contains a list of HSMClientCertificates
type HSMClientCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HSMClientCertificate `json:"items"`
}

// Repository type metadata.
var (
	HSMClientCertificate_Kind             = "HSMClientCertificate"
	HSMClientCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HSMClientCertificate_Kind}.String()
	HSMClientCertificate_KindAPIVersion   = HSMClientCertificate_Kind + "." + CRDGroupVersion.String()
	HSMClientCertificate_GroupVersionKind = CRDGroupVersion.WithKind(HSMClientCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&HSMClientCertificate{}, &HSMClientCertificateList{})
}
