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

type AbortIncompleteMultipartUploadObservation struct {

	// The number of days after which Amazon S3 aborts an incomplete multipart upload.
	DaysAfterInitiation *float64 `json:"daysAfterInitiation,omitempty" tf:"days_after_initiation,omitempty"`
}

type AbortIncompleteMultipartUploadParameters struct {

	// The number of days after which Amazon S3 aborts an incomplete multipart upload.
	// +kubebuilder:validation:Optional
	DaysAfterInitiation *float64 `json:"daysAfterInitiation,omitempty" tf:"days_after_initiation,omitempty"`
}

type AndObservation struct {

	// Minimum object size (in bytes) to which the rule applies.
	ObjectSizeGreaterThan *float64 `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	ObjectSizeLessThan *float64 `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags. All of these tags must exist in the object's tag set in order for the rule to apply.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AndParameters struct {

	// Minimum object size (in bytes) to which the rule applies.
	// +kubebuilder:validation:Optional
	ObjectSizeGreaterThan *float64 `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	// +kubebuilder:validation:Optional
	ObjectSizeLessThan *float64 `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags. All of these tags must exist in the object's tag set in order for the rule to apply.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BucketLifecycleConfigurationObservation struct {

	// The name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// and status)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// List of configuration blocks describing the rules managing the replication documented below.
	Rule []BucketLifecycleConfigurationRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketLifecycleConfigurationParameters struct {

	// The name of the source S3 bucket you want Amazon S3 to monitor.
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

	// The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of configuration blocks describing the rules managing the replication documented below.
	// +kubebuilder:validation:Optional
	Rule []BucketLifecycleConfigurationRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketLifecycleConfigurationRuleObservation struct {

	// Configuration block that specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload documented below.
	AbortIncompleteMultipartUpload []AbortIncompleteMultipartUploadObservation `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload,omitempty"`

	// Configuration block that specifies the expiration for the lifecycle of the object in the form of date, days and, whether the object has a delete marker documented below.
	Expiration []RuleExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Configuration block used to identify objects that a Lifecycle Rule applies to documented below. If not specified, the rule will default to using prefix.
	Filter []RuleFilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block that specifies when noncurrent object versions expire documented below.
	NoncurrentVersionExpiration []RuleNoncurrentVersionExpirationObservation `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Set of configuration blocks that specify the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class documented below.
	NoncurrentVersionTransition []RuleNoncurrentVersionTransitionObservation `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Whether the rule is currently being applied. Valid values: Enabled or Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Set of configuration blocks that specify when an Amazon S3 object transitions to a specified storage class documented below.
	Transition []RuleTransitionObservation `json:"transition,omitempty" tf:"transition,omitempty"`
}

type BucketLifecycleConfigurationRuleParameters struct {

	// Configuration block that specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload documented below.
	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUpload []AbortIncompleteMultipartUploadParameters `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload,omitempty"`

	// Configuration block that specifies the expiration for the lifecycle of the object in the form of date, days and, whether the object has a delete marker documented below.
	// +kubebuilder:validation:Optional
	Expiration []RuleExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Configuration block used to identify objects that a Lifecycle Rule applies to documented below. If not specified, the rule will default to using prefix.
	// +kubebuilder:validation:Optional
	Filter []RuleFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Configuration block that specifies when noncurrent object versions expire documented below.
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []RuleNoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Set of configuration blocks that specify the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class documented below.
	// +kubebuilder:validation:Optional
	NoncurrentVersionTransition []RuleNoncurrentVersionTransitionParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Whether the rule is currently being applied. Valid values: Enabled or Disabled.
	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`

	// Set of configuration blocks that specify when an Amazon S3 object transitions to a specified storage class documented below.
	// +kubebuilder:validation:Optional
	Transition []RuleTransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type RuleExpirationObservation struct {

	// The date objects are transitioned to the specified storage class. The date value must be in RFC3339 format and set to midnight UTC e.g. 2023-01-13T00:00:00Z.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// The number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no action.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type RuleExpirationParameters struct {

	// The date objects are transitioned to the specified storage class. The date value must be in RFC3339 format and set to midnight UTC e.g. 2023-01-13T00:00:00Z.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// The number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no action.
	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type RuleFilterObservation struct {

	// Configuration block used to apply a logical AND to two or more predicates documented below. The Lifecycle Rule will apply to any object matching all the predicates configured inside the and block.
	And []AndObservation `json:"and,omitempty" tf:"and,omitempty"`

	// Minimum object size (in bytes) to which the rule applies.
	ObjectSizeGreaterThan *string `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	ObjectSizeLessThan *string `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// A configuration block for specifying a tag key and value documented below.
	Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type RuleFilterParameters struct {

	// Configuration block used to apply a logical AND to two or more predicates documented below. The Lifecycle Rule will apply to any object matching all the predicates configured inside the and block.
	// +kubebuilder:validation:Optional
	And []AndParameters `json:"and,omitempty" tf:"and,omitempty"`

	// Minimum object size (in bytes) to which the rule applies.
	// +kubebuilder:validation:Optional
	ObjectSizeGreaterThan *string `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	// +kubebuilder:validation:Optional
	ObjectSizeLessThan *string `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// A configuration block for specifying a tag key and value documented below.
	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type RuleNoncurrentVersionExpirationObservation struct {

	// The number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// The number of days an object is noncurrent before Amazon S3 can perform the associated action.
	NoncurrentDays *float64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`
}

type RuleNoncurrentVersionExpirationParameters struct {

	// The number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	// +kubebuilder:validation:Optional
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// The number of days an object is noncurrent before Amazon S3 can perform the associated action.
	// +kubebuilder:validation:Optional
	NoncurrentDays *float64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`
}

type RuleNoncurrentVersionTransitionObservation struct {

	// The number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// The number of days an object is noncurrent before Amazon S3 can perform the associated action.
	NoncurrentDays *float64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`

	// The class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleNoncurrentVersionTransitionParameters struct {

	// The number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	// +kubebuilder:validation:Optional
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// The number of days an object is noncurrent before Amazon S3 can perform the associated action.
	// +kubebuilder:validation:Optional
	NoncurrentDays *float64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`

	// The class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type RuleTransitionObservation struct {

	// The date objects are transitioned to the specified storage class. The date value must be in RFC3339 format and set to midnight UTC e.g. 2023-01-13T00:00:00Z.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// The number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// The class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleTransitionParameters struct {

	// The date objects are transitioned to the specified storage class. The date value must be in RFC3339 format and set to midnight UTC e.g. 2023-01-13T00:00:00Z.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// The number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// The class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type TagObservation struct {

	// Name of the object key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of the tag.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagParameters struct {

	// Name of the object key.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Value of the tag.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// BucketLifecycleConfigurationSpec defines the desired state of BucketLifecycleConfiguration
type BucketLifecycleConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketLifecycleConfigurationParameters `json:"forProvider"`
}

// BucketLifecycleConfigurationStatus defines the observed state of BucketLifecycleConfiguration.
type BucketLifecycleConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketLifecycleConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleConfiguration is the Schema for the BucketLifecycleConfigurations API. Provides a S3 bucket lifecycle configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.rule)",message="rule is a required parameter"
	Spec   BucketLifecycleConfigurationSpec   `json:"spec"`
	Status BucketLifecycleConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleConfigurationList contains a list of BucketLifecycleConfigurations
type BucketLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLifecycleConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketLifecycleConfiguration_Kind             = "BucketLifecycleConfiguration"
	BucketLifecycleConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketLifecycleConfiguration_Kind}.String()
	BucketLifecycleConfiguration_KindAPIVersion   = BucketLifecycleConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketLifecycleConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketLifecycleConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketLifecycleConfiguration{}, &BucketLifecycleConfigurationList{})
}
