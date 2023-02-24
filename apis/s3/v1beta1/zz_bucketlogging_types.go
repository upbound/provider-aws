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

type BucketLoggingObservation struct {

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The name of the bucket where you want Amazon S3 to store server access logs.
	TargetBucket *string `json:"targetBucket,omitempty" tf:"target_bucket,omitempty"`

	// Set of configuration blocks with information for granting permissions documented below.
	TargetGrant []TargetGrantObservation `json:"targetGrant,omitempty" tf:"target_grant,omitempty"`

	// A prefix for all log object keys.
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type BucketLoggingParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the bucket where you want Amazon S3 to store server access logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetBucket *string `json:"targetBucket,omitempty" tf:"target_bucket,omitempty"`

	// Reference to a Bucket in s3 to populate targetBucket.
	// +kubebuilder:validation:Optional
	TargetBucketRef *v1.Reference `json:"targetBucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate targetBucket.
	// +kubebuilder:validation:Optional
	TargetBucketSelector *v1.Selector `json:"targetBucketSelector,omitempty" tf:"-"`

	// Set of configuration blocks with information for granting permissions documented below.
	// +kubebuilder:validation:Optional
	TargetGrant []TargetGrantParameters `json:"targetGrant,omitempty" tf:"target_grant,omitempty"`

	// A prefix for all log object keys.
	// +kubebuilder:validation:Required
	TargetPrefix *string `json:"targetPrefix" tf:"target_prefix,omitempty"`
}

type TargetGrantGranteeObservation struct {
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Email address of the grantee. See Regions and Endpoints for supported AWS regions where this argument can be specified.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// The canonical user ID of the grantee.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of grantee. Valid values: CanonicalUser, AmazonCustomerByEmail, Group.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// URI of the grantee group.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type TargetGrantGranteeParameters struct {

	// Email address of the grantee. See Regions and Endpoints for supported AWS regions where this argument can be specified.
	// +kubebuilder:validation:Optional
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// The canonical user ID of the grantee.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of grantee. Valid values: CanonicalUser, AmazonCustomerByEmail, Group.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// URI of the grantee group.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type TargetGrantObservation struct {

	// A configuration block for the person being granted permissions documented below.
	Grantee []TargetGrantGranteeObservation `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Logging permissions assigned to the grantee for the bucket. Valid values: FULL_CONTROL, READ, WRITE.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type TargetGrantParameters struct {

	// A configuration block for the person being granted permissions documented below.
	// +kubebuilder:validation:Required
	Grantee []TargetGrantGranteeParameters `json:"grantee" tf:"grantee,omitempty"`

	// Logging permissions assigned to the grantee for the bucket. Valid values: FULL_CONTROL, READ, WRITE.
	// +kubebuilder:validation:Required
	Permission *string `json:"permission" tf:"permission,omitempty"`
}

// BucketLoggingSpec defines the desired state of BucketLogging
type BucketLoggingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketLoggingParameters `json:"forProvider"`
}

// BucketLoggingStatus defines the observed state of BucketLogging.
type BucketLoggingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketLoggingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLogging is the Schema for the BucketLoggings API. Provides an S3 bucket (server access) logging resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketLogging struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketLoggingSpec   `json:"spec"`
	Status            BucketLoggingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLoggingList contains a list of BucketLoggings
type BucketLoggingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLogging `json:"items"`
}

// Repository type metadata.
var (
	BucketLogging_Kind             = "BucketLogging"
	BucketLogging_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketLogging_Kind}.String()
	BucketLogging_KindAPIVersion   = BucketLogging_Kind + "." + CRDGroupVersion.String()
	BucketLogging_GroupVersionKind = CRDGroupVersion.WithKind(BucketLogging_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketLogging{}, &BucketLoggingList{})
}
