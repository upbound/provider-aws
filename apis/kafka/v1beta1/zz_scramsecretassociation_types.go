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

type ScramSecretAssociationInitParameters struct {
}

type ScramSecretAssociationObservation struct {

	// Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn *string `json:"clusterArn,omitempty" tf:"cluster_arn,omitempty"`

	// Amazon Resource Name (ARN) of the MSK cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of all AWS Secrets Manager secret ARNs to associate with the cluster. Secrets not referenced, selected or listed here will be disassociated from the cluster.
	SecretArnList []*string `json:"secretArnList,omitempty" tf:"secret_arn_list,omitempty"`
}

type ScramSecretAssociationParameters struct {

	// Amazon Resource Name (ARN) of the MSK cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kafka/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ClusterArn *string `json:"clusterArn,omitempty" tf:"cluster_arn,omitempty"`

	// Reference to a Cluster in kafka to populate clusterArn.
	// +kubebuilder:validation:Optional
	ClusterArnRef *v1.Reference `json:"clusterArnRef,omitempty" tf:"-"`

	// Selector for a Cluster in kafka to populate clusterArn.
	// +kubebuilder:validation:Optional
	ClusterArnSelector *v1.Selector `json:"clusterArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of all AWS Secrets Manager secret ARNs to associate with the cluster. Secrets not referenced, selected or listed here will be disassociated from the cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:refFieldName=SecretArnRefs
	// +crossplane:generate:reference:selectorFieldName=SecretArnSelector
	// +kubebuilder:validation:Optional
	SecretArnList []*string `json:"secretArnList,omitempty" tf:"secret_arn_list,omitempty"`

	// References to Secret in secretsmanager to populate secretArnList.
	// +kubebuilder:validation:Optional
	SecretArnRefs []v1.Reference `json:"secretArnRefs,omitempty" tf:"-"`

	// Selector for a list of Secret in secretsmanager to populate secretArnList.
	// +kubebuilder:validation:Optional
	SecretArnSelector *v1.Selector `json:"secretArnSelector,omitempty" tf:"-"`
}

// ScramSecretAssociationSpec defines the desired state of ScramSecretAssociation
type ScramSecretAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScramSecretAssociationParameters `json:"forProvider"`
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
	InitProvider ScramSecretAssociationInitParameters `json:"initProvider,omitempty"`
}

// ScramSecretAssociationStatus defines the observed state of ScramSecretAssociation.
type ScramSecretAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScramSecretAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScramSecretAssociation is the Schema for the ScramSecretAssociations API. Associates SCRAM secrets with a Managed Streaming for Kafka (MSK) cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ScramSecretAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScramSecretAssociationSpec   `json:"spec"`
	Status            ScramSecretAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScramSecretAssociationList contains a list of ScramSecretAssociations
type ScramSecretAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScramSecretAssociation `json:"items"`
}

// Repository type metadata.
var (
	ScramSecretAssociation_Kind             = "ScramSecretAssociation"
	ScramSecretAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScramSecretAssociation_Kind}.String()
	ScramSecretAssociation_KindAPIVersion   = ScramSecretAssociation_Kind + "." + CRDGroupVersion.String()
	ScramSecretAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ScramSecretAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ScramSecretAssociation{}, &ScramSecretAssociationList{})
}
