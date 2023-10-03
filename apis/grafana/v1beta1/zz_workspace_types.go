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

type NetworkAccessControlInitParameters struct {

	// - An array of prefix list IDs.
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids,omitempty"`

	// - An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
	VpceIds []*string `json:"vpceIds,omitempty" tf:"vpce_ids,omitempty"`
}

type NetworkAccessControlObservation struct {

	// - An array of prefix list IDs.
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids,omitempty"`

	// - An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
	VpceIds []*string `json:"vpceIds,omitempty" tf:"vpce_ids,omitempty"`
}

type NetworkAccessControlParameters struct {

	// - An array of prefix list IDs.
	// +kubebuilder:validation:Optional
	PrefixListIds []*string `json:"prefixListIds" tf:"prefix_list_ids,omitempty"`

	// - An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
	// +kubebuilder:validation:Optional
	VpceIds []*string `json:"vpceIds" tf:"vpce_ids,omitempty"`
}

type VPCConfigurationInitParameters struct {

	// - The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// - The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type VPCConfigurationObservation struct {

	// - The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// - The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type VPCConfigurationParameters struct {

	// - The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// - The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`
}

type WorkspaceInitParameters struct {

	// The type of account access for the workspace. Valid values are CURRENT_ACCOUNT and ORGANIZATION. If ORGANIZATION is specified, then organizational_units must also be present.
	AccountAccessType *string `json:"accountAccessType,omitempty" tf:"account_access_type,omitempty"`

	// The authentication providers for the workspace. Valid values are AWS_SSO, SAML, or both.
	AuthenticationProviders []*string `json:"authenticationProviders,omitempty" tf:"authentication_providers,omitempty"`

	// The configuration string for the workspace that you create. For more information about the format and configuration options available, see Working in your Grafana workspace.
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// The data sources for the workspace. Valid values are AMAZON_OPENSEARCH_SERVICE, ATHENA, CLOUDWATCH, PROMETHEUS, REDSHIFT, SITEWISE, TIMESTREAM, XRAY
	DataSources []*string `json:"dataSources,omitempty" tf:"data_sources,omitempty"`

	// The workspace description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the version of Grafana to support in the new workspace. Supported values are 8.4 and 9.4. If not specified, defaults to 8.4. Upgrading the workspace version isn't supported, however it's possible to copy content from the old version to the new one using AWS official migration tool.
	GrafanaVersion *string `json:"grafanaVersion,omitempty" tf:"grafana_version,omitempty"`

	// The Grafana workspace name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for network access to your workspace.See Network Access Control below.
	NetworkAccessControl []NetworkAccessControlInitParameters `json:"networkAccessControl,omitempty" tf:"network_access_control,omitempty"`

	// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to SNS.
	NotificationDestinations []*string `json:"notificationDestinations,omitempty" tf:"notification_destinations,omitempty"`

	// The role name that the workspace uses to access resources through Amazon Organizations.
	OrganizationRoleName *string `json:"organizationRoleName,omitempty" tf:"organization_role_name,omitempty"`

	// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
	OrganizationalUnits []*string `json:"organizationalUnits,omitempty" tf:"organizational_units,omitempty"`

	// The permission type of the workspace. If SERVICE_MANAGED is specified, the IAM roles and IAM policy attachments are generated automatically. If CUSTOMER_MANAGED is specified, the IAM roles and IAM policy attachments will not be created.
	PermissionType *string `json:"permissionType,omitempty" tf:"permission_type,omitempty"`

	// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
	StackSetName *string `json:"stackSetName,omitempty" tf:"stack_set_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
	VPCConfiguration []VPCConfigurationInitParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type WorkspaceObservation struct {

	// The type of account access for the workspace. Valid values are CURRENT_ACCOUNT and ORGANIZATION. If ORGANIZATION is specified, then organizational_units must also be present.
	AccountAccessType *string `json:"accountAccessType,omitempty" tf:"account_access_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Grafana workspace.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The authentication providers for the workspace. Valid values are AWS_SSO, SAML, or both.
	AuthenticationProviders []*string `json:"authenticationProviders,omitempty" tf:"authentication_providers,omitempty"`

	// The configuration string for the workspace that you create. For more information about the format and configuration options available, see Working in your Grafana workspace.
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// The data sources for the workspace. Valid values are AMAZON_OPENSEARCH_SERVICE, ATHENA, CLOUDWATCH, PROMETHEUS, REDSHIFT, SITEWISE, TIMESTREAM, XRAY
	DataSources []*string `json:"dataSources,omitempty" tf:"data_sources,omitempty"`

	// The workspace description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The endpoint of the Grafana workspace.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Specifies the version of Grafana to support in the new workspace. Supported values are 8.4 and 9.4. If not specified, defaults to 8.4. Upgrading the workspace version isn't supported, however it's possible to copy content from the old version to the new one using AWS official migration tool.
	GrafanaVersion *string `json:"grafanaVersion,omitempty" tf:"grafana_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Grafana workspace name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for network access to your workspace.See Network Access Control below.
	NetworkAccessControl []NetworkAccessControlObservation `json:"networkAccessControl,omitempty" tf:"network_access_control,omitempty"`

	// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to SNS.
	NotificationDestinations []*string `json:"notificationDestinations,omitempty" tf:"notification_destinations,omitempty"`

	// The role name that the workspace uses to access resources through Amazon Organizations.
	OrganizationRoleName *string `json:"organizationRoleName,omitempty" tf:"organization_role_name,omitempty"`

	// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
	OrganizationalUnits []*string `json:"organizationalUnits,omitempty" tf:"organizational_units,omitempty"`

	// The permission type of the workspace. If SERVICE_MANAGED is specified, the IAM roles and IAM policy attachments are generated automatically. If CUSTOMER_MANAGED is specified, the IAM roles and IAM policy attachments will not be created.
	PermissionType *string `json:"permissionType,omitempty" tf:"permission_type,omitempty"`

	// The IAM role ARN that the workspace assumes.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	SAMLConfigurationStatus *string `json:"samlConfigurationStatus,omitempty" tf:"saml_configuration_status,omitempty"`

	// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
	StackSetName *string `json:"stackSetName,omitempty" tf:"stack_set_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
	VPCConfiguration []VPCConfigurationObservation `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type WorkspaceParameters struct {

	// The type of account access for the workspace. Valid values are CURRENT_ACCOUNT and ORGANIZATION. If ORGANIZATION is specified, then organizational_units must also be present.
	// +kubebuilder:validation:Optional
	AccountAccessType *string `json:"accountAccessType,omitempty" tf:"account_access_type,omitempty"`

	// The authentication providers for the workspace. Valid values are AWS_SSO, SAML, or both.
	// +kubebuilder:validation:Optional
	AuthenticationProviders []*string `json:"authenticationProviders,omitempty" tf:"authentication_providers,omitempty"`

	// The configuration string for the workspace that you create. For more information about the format and configuration options available, see Working in your Grafana workspace.
	// +kubebuilder:validation:Optional
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// The data sources for the workspace. Valid values are AMAZON_OPENSEARCH_SERVICE, ATHENA, CLOUDWATCH, PROMETHEUS, REDSHIFT, SITEWISE, TIMESTREAM, XRAY
	// +kubebuilder:validation:Optional
	DataSources []*string `json:"dataSources,omitempty" tf:"data_sources,omitempty"`

	// The workspace description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the version of Grafana to support in the new workspace. Supported values are 8.4 and 9.4. If not specified, defaults to 8.4. Upgrading the workspace version isn't supported, however it's possible to copy content from the old version to the new one using AWS official migration tool.
	// +kubebuilder:validation:Optional
	GrafanaVersion *string `json:"grafanaVersion,omitempty" tf:"grafana_version,omitempty"`

	// The Grafana workspace name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for network access to your workspace.See Network Access Control below.
	// +kubebuilder:validation:Optional
	NetworkAccessControl []NetworkAccessControlParameters `json:"networkAccessControl,omitempty" tf:"network_access_control,omitempty"`

	// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to SNS.
	// +kubebuilder:validation:Optional
	NotificationDestinations []*string `json:"notificationDestinations,omitempty" tf:"notification_destinations,omitempty"`

	// The role name that the workspace uses to access resources through Amazon Organizations.
	// +kubebuilder:validation:Optional
	OrganizationRoleName *string `json:"organizationRoleName,omitempty" tf:"organization_role_name,omitempty"`

	// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
	// +kubebuilder:validation:Optional
	OrganizationalUnits []*string `json:"organizationalUnits,omitempty" tf:"organizational_units,omitempty"`

	// The permission type of the workspace. If SERVICE_MANAGED is specified, the IAM roles and IAM policy attachments are generated automatically. If CUSTOMER_MANAGED is specified, the IAM roles and IAM policy attachments will not be created.
	// +kubebuilder:validation:Optional
	PermissionType *string `json:"permissionType,omitempty" tf:"permission_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The IAM role ARN that the workspace assumes.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
	// +kubebuilder:validation:Optional
	StackSetName *string `json:"stackSetName,omitempty" tf:"stack_set_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
	// +kubebuilder:validation:Optional
	VPCConfiguration []VPCConfigurationParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
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
	InitProvider WorkspaceInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workspace is the Schema for the Workspaces API. Provides an Amazon Managed Grafana workspace resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountAccessType) || (has(self.initProvider) && has(self.initProvider.accountAccessType))",message="spec.forProvider.accountAccessType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authenticationProviders) || (has(self.initProvider) && has(self.initProvider.authenticationProviders))",message="spec.forProvider.authenticationProviders is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permissionType) || (has(self.initProvider) && has(self.initProvider.permissionType))",message="spec.forProvider.permissionType is a required parameter"
	Spec   WorkspaceSpec   `json:"spec"`
	Status WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Repository type metadata.
var (
	Workspace_Kind             = "Workspace"
	Workspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workspace_Kind}.String()
	Workspace_KindAPIVersion   = Workspace_Kind + "." + CRDGroupVersion.String()
	Workspace_GroupVersionKind = CRDGroupVersion.WithKind(Workspace_Kind)
)

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
