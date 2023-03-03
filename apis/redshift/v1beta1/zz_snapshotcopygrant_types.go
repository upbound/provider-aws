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

type SnapshotCopyGrantObservation struct {

	// Amazon Resource Name (ARN) of snapshot copy grant
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A friendly name for identifying the grant.
	SnapshotCopyGrantName *string `json:"snapshotCopyGrantName,omitempty" tf:"snapshot_copy_grant_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SnapshotCopyGrantParameters struct {

	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A friendly name for identifying the grant.
	// +kubebuilder:validation:Optional
	SnapshotCopyGrantName *string `json:"snapshotCopyGrantName,omitempty" tf:"snapshot_copy_grant_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SnapshotCopyGrantSpec defines the desired state of SnapshotCopyGrant
type SnapshotCopyGrantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotCopyGrantParameters `json:"forProvider"`
}

// SnapshotCopyGrantStatus defines the observed state of SnapshotCopyGrant.
type SnapshotCopyGrantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotCopyGrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotCopyGrant is the Schema for the SnapshotCopyGrants API. Creates a snapshot copy grant that allows AWS Redshift to encrypt copied snapshots with a customer master key from AWS KMS in a destination region.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SnapshotCopyGrant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.snapshotCopyGrantName)",message="snapshotCopyGrantName is a required parameter"
	Spec   SnapshotCopyGrantSpec   `json:"spec"`
	Status SnapshotCopyGrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotCopyGrantList contains a list of SnapshotCopyGrants
type SnapshotCopyGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotCopyGrant `json:"items"`
}

// Repository type metadata.
var (
	SnapshotCopyGrant_Kind             = "SnapshotCopyGrant"
	SnapshotCopyGrant_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotCopyGrant_Kind}.String()
	SnapshotCopyGrant_KindAPIVersion   = SnapshotCopyGrant_Kind + "." + CRDGroupVersion.String()
	SnapshotCopyGrant_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotCopyGrant_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotCopyGrant{}, &SnapshotCopyGrantList{})
}
