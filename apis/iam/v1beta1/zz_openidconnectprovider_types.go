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

type OpenIDConnectProviderInitParameters struct {

	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
	ClientIDList []*string `json:"clientIdList,omitempty" tf:"client_id_list,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintList []*string `json:"thumbprintList,omitempty" tf:"thumbprint_list,omitempty"`

	// The URL of the identity provider. Corresponds to the iss claim.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type OpenIDConnectProviderObservation struct {

	// The ARN assigned by AWS for this provider.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
	ClientIDList []*string `json:"clientIdList,omitempty" tf:"client_id_list,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintList []*string `json:"thumbprintList,omitempty" tf:"thumbprint_list,omitempty"`

	// The URL of the identity provider. Corresponds to the iss claim.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type OpenIDConnectProviderParameters struct {

	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
	// +kubebuilder:validation:Optional
	ClientIDList []*string `json:"clientIdList,omitempty" tf:"client_id_list,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	// +kubebuilder:validation:Optional
	ThumbprintList []*string `json:"thumbprintList,omitempty" tf:"thumbprint_list,omitempty"`

	// The URL of the identity provider. Corresponds to the iss claim.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// OpenIDConnectProviderSpec defines the desired state of OpenIDConnectProvider
type OpenIDConnectProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OpenIDConnectProviderParameters `json:"forProvider"`
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
	InitProvider OpenIDConnectProviderInitParameters `json:"initProvider,omitempty"`
}

// OpenIDConnectProviderStatus defines the observed state of OpenIDConnectProvider.
type OpenIDConnectProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OpenIDConnectProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OpenIDConnectProvider is the Schema for the OpenIDConnectProviders API. Provides an IAM OpenID Connect provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OpenIDConnectProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientIdList) || (has(self.initProvider) && has(self.initProvider.clientIdList))",message="spec.forProvider.clientIdList is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.thumbprintList) || (has(self.initProvider) && has(self.initProvider.thumbprintList))",message="spec.forProvider.thumbprintList is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   OpenIDConnectProviderSpec   `json:"spec"`
	Status OpenIDConnectProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpenIDConnectProviderList contains a list of OpenIDConnectProviders
type OpenIDConnectProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenIDConnectProvider `json:"items"`
}

// Repository type metadata.
var (
	OpenIDConnectProvider_Kind             = "OpenIDConnectProvider"
	OpenIDConnectProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OpenIDConnectProvider_Kind}.String()
	OpenIDConnectProvider_KindAPIVersion   = OpenIDConnectProvider_Kind + "." + CRDGroupVersion.String()
	OpenIDConnectProvider_GroupVersionKind = CRDGroupVersion.WithKind(OpenIDConnectProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&OpenIDConnectProvider{}, &OpenIDConnectProviderList{})
}
