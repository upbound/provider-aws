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

type ProxyEndpointInitParameters struct {

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is READ_WRITE. Valid values are READ_WRITE and READ_ONLY.
	TargetRole *string `json:"targetRole,omitempty" tf:"target_role,omitempty"`

	// One or more VPC subnet IDs to associate with the new proxy.
	VPCSubnetIds []*string `json:"vpcSubnetIds,omitempty" tf:"vpc_subnet_ids,omitempty"`
}

type ProxyEndpointObservation struct {

	// The Amazon Resource Name (ARN) for the proxy endpoint.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	DBProxyName *string `json:"dbProxyName,omitempty" tf:"db_proxy_name,omitempty"`

	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The name of the proxy and proxy endpoint separated by /, DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether this endpoint is the default endpoint for the associated DB proxy.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is READ_WRITE. Valid values are READ_WRITE and READ_ONLY.
	TargetRole *string `json:"targetRole,omitempty" tf:"target_role,omitempty"`

	// The VPC ID of the DB proxy endpoint.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// One or more VPC security group IDs to associate with the new proxy.
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// One or more VPC subnet IDs to associate with the new proxy.
	VPCSubnetIds []*string `json:"vpcSubnetIds,omitempty" tf:"vpc_subnet_ids,omitempty"`
}

type ProxyEndpointParameters struct {

	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Proxy
	// +kubebuilder:validation:Optional
	DBProxyName *string `json:"dbProxyName,omitempty" tf:"db_proxy_name,omitempty"`

	// Reference to a Proxy in rds to populate dbProxyName.
	// +kubebuilder:validation:Optional
	DBProxyNameRef *v1.Reference `json:"dbProxyNameRef,omitempty" tf:"-"`

	// Selector for a Proxy in rds to populate dbProxyName.
	// +kubebuilder:validation:Optional
	DBProxyNameSelector *v1.Selector `json:"dbProxyNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is READ_WRITE. Valid values are READ_WRITE and READ_ONLY.
	// +kubebuilder:validation:Optional
	TargetRole *string `json:"targetRole,omitempty" tf:"target_role,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// One or more VPC security group IDs to associate with the new proxy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// One or more VPC subnet IDs to associate with the new proxy.
	// +kubebuilder:validation:Optional
	VPCSubnetIds []*string `json:"vpcSubnetIds,omitempty" tf:"vpc_subnet_ids,omitempty"`
}

// ProxyEndpointSpec defines the desired state of ProxyEndpoint
type ProxyEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProxyEndpointParameters `json:"forProvider"`
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
	InitProvider ProxyEndpointInitParameters `json:"initProvider,omitempty"`
}

// ProxyEndpointStatus defines the observed state of ProxyEndpoint.
type ProxyEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProxyEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyEndpoint is the Schema for the ProxyEndpoints API. Provides an RDS DB proxy endpoint resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ProxyEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcSubnetIds) || (has(self.initProvider) && has(self.initProvider.vpcSubnetIds))",message="spec.forProvider.vpcSubnetIds is a required parameter"
	Spec   ProxyEndpointSpec   `json:"spec"`
	Status ProxyEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyEndpointList contains a list of ProxyEndpoints
type ProxyEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProxyEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ProxyEndpoint_Kind             = "ProxyEndpoint"
	ProxyEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProxyEndpoint_Kind}.String()
	ProxyEndpoint_KindAPIVersion   = ProxyEndpoint_Kind + "." + CRDGroupVersion.String()
	ProxyEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(ProxyEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ProxyEndpoint{}, &ProxyEndpointList{})
}
