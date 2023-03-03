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

type AccountObservation struct {

	// The ARN for this account.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// If true, a deletion event will close the account. Otherwise, it will only remove from the organization. This is not supported for GovCloud accounts.
	CloseOnDeletion *bool `json:"closeOnDeletion,omitempty" tf:"close_on_deletion,omitempty"`

	// Whether to also create a GovCloud account. The GovCloud account is tied to the main (commercial) account this resource creates. If true, the GovCloud account ID is available in the govcloud_id attribute.
	CreateGovcloud *bool `json:"createGovcloud,omitempty" tf:"create_govcloud,omitempty"`

	// Email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// ID for a GovCloud account created with the account.
	GovcloudID *string `json:"govcloudId,omitempty" tf:"govcloud_id,omitempty"`

	// If set to ALLOW, the new account enables IAM users and roles to access account billing information if they have the required permissions. If set to DENY, then only the root user (and no roles) of the new account can access account billing information. If this is unset, the AWS API will default this to ALLOW. If the resource is created and this option is changed, it will try to recreate the account.
	IAMUserAccessToBilling *string `json:"iamUserAccessToBilling,omitempty" tf:"iam_user_access_to_billing,omitempty"`

	// The AWS account id
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	JoinedMethod *string `json:"joinedMethod,omitempty" tf:"joined_method,omitempty"`

	JoinedTimestamp *string `json:"joinedTimestamp,omitempty" tf:"joined_timestamp,omitempty"`

	// Friendly name for the member account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AccountParameters struct {

	// If true, a deletion event will close the account. Otherwise, it will only remove from the organization. This is not supported for GovCloud accounts.
	// +kubebuilder:validation:Optional
	CloseOnDeletion *bool `json:"closeOnDeletion,omitempty" tf:"close_on_deletion,omitempty"`

	// Whether to also create a GovCloud account. The GovCloud account is tied to the main (commercial) account this resource creates. If true, the GovCloud account ID is available in the govcloud_id attribute.
	// +kubebuilder:validation:Optional
	CreateGovcloud *bool `json:"createGovcloud,omitempty" tf:"create_govcloud,omitempty"`

	// Email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// If set to ALLOW, the new account enables IAM users and roles to access account billing information if they have the required permissions. If set to DENY, then only the root user (and no roles) of the new account can access account billing information. If this is unset, the AWS API will default this to ALLOW. If the resource is created and this option is changed, it will try to recreate the account.
	// +kubebuilder:validation:Optional
	IAMUserAccessToBilling *string `json:"iamUserAccessToBilling,omitempty" tf:"iam_user_access_to_billing,omitempty"`

	// Friendly name for the member account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	// +kubebuilder:validation:Optional
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Account is the Schema for the Accounts API. Provides a resource to create a member account in the current AWS Organization.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.email)",message="email is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   AccountSpec   `json:"spec"`
	Status AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
