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

type CustomCookbooksSourceObservation struct {
}

type CustomCookbooksSourceParameters struct {

	// Password to use when authenticating to the source.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// For sources that are version-aware, the revision to use.
	// +kubebuilder:validation:Optional
	Revision *string `json:"revision,omitempty" tf:"revision,omitempty"`

	// SSH key to use when authenticating to the source.
	// +kubebuilder:validation:Optional
	SSHKeySecretRef *v1.SecretKeySelector `json:"sshKeySecretRef,omitempty" tf:"-"`

	// The type of source to use. For example, "archive".
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The URL where the cookbooks resource can be found.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// Username to use when authenticating to the source.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type StackObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The id of the stack.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	StackEndpoint *string `json:"stackEndpoint,omitempty" tf:"stack_endpoint,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type StackParameters struct {

	// If set to "LATEST", OpsWorks will automatically install the latest version.
	// +kubebuilder:validation:Optional
	AgentVersion *string `json:"agentVersion,omitempty" tf:"agent_version,omitempty"`

	// If manage_berkshelf is enabled, the version of Berkshelf to use.
	// +kubebuilder:validation:Optional
	BerkshelfVersion *string `json:"berkshelfVersion,omitempty" tf:"berkshelf_version,omitempty"`

	// Color to paint next to the stack's resources in the OpsWorks console.
	// +kubebuilder:validation:Optional
	Color *string `json:"color,omitempty" tf:"color,omitempty"`

	// Name of the configuration manager to use. Defaults to "Chef".
	// +kubebuilder:validation:Optional
	ConfigurationManagerName *string `json:"configurationManagerName,omitempty" tf:"configuration_manager_name,omitempty"`

	// Version of the configuration manager to use. Defaults to "11.4".
	// +kubebuilder:validation:Optional
	ConfigurationManagerVersion *string `json:"configurationManagerVersion,omitempty" tf:"configuration_manager_version,omitempty"`

	// When use_custom_cookbooks is set, provide this sub-object as described below.
	// +kubebuilder:validation:Optional
	CustomCookbooksSource []CustomCookbooksSourceParameters `json:"customCookbooksSource,omitempty" tf:"custom_cookbooks_source,omitempty"`

	// User defined JSON passed to "Chef". Use a "here doc" for multiline JSON.
	// +kubebuilder:validation:Optional
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json,omitempty"`

	// Name of the availability zone where instances will be created by default.
	// Cannot be set when vpc_id is set.
	// +kubebuilder:validation:Optional
	DefaultAvailabilityZone *string `json:"defaultAvailabilityZone,omitempty" tf:"default_availability_zone,omitempty"`

	// The ARN of an IAM Instance Profile that created instances will have by default.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.InstanceProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	DefaultInstanceProfileArn *string `json:"defaultInstanceProfileArn,omitempty" tf:"default_instance_profile_arn,omitempty"`

	// Reference to a InstanceProfile in iam to populate defaultInstanceProfileArn.
	// +kubebuilder:validation:Optional
	DefaultInstanceProfileArnRef *v1.Reference `json:"defaultInstanceProfileArnRef,omitempty" tf:"-"`

	// Selector for a InstanceProfile in iam to populate defaultInstanceProfileArn.
	// +kubebuilder:validation:Optional
	DefaultInstanceProfileArnSelector *v1.Selector `json:"defaultInstanceProfileArnSelector,omitempty" tf:"-"`

	// Name of OS that will be installed on instances by default.
	// +kubebuilder:validation:Optional
	DefaultOs *string `json:"defaultOs,omitempty" tf:"default_os,omitempty"`

	// Name of the type of root device instances will have by default.
	// +kubebuilder:validation:Optional
	DefaultRootDeviceType *string `json:"defaultRootDeviceType,omitempty" tf:"default_root_device_type,omitempty"`

	// Name of the SSH keypair that instances will have by default.
	// +kubebuilder:validation:Optional
	DefaultSSHKeyName *string `json:"defaultSshKeyName,omitempty" tf:"default_ssh_key_name,omitempty"`

	// ID of the subnet in which instances will be created by default.
	// Required if vpc_id is set to a VPC other than the default VPC, and forbidden if it isn't.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	DefaultSubnetID *string `json:"defaultSubnetId,omitempty" tf:"default_subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate defaultSubnetId.
	// +kubebuilder:validation:Optional
	DefaultSubnetIDRef *v1.Reference `json:"defaultSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate defaultSubnetId.
	// +kubebuilder:validation:Optional
	DefaultSubnetIDSelector *v1.Selector `json:"defaultSubnetIdSelector,omitempty" tf:"-"`

	// Keyword representing the naming scheme that will be used for instance hostnames within this stack.
	// +kubebuilder:validation:Optional
	HostnameTheme *string `json:"hostnameTheme,omitempty" tf:"hostname_theme,omitempty"`

	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	// +kubebuilder:validation:Optional
	ManageBerkshelf *bool `json:"manageBerkshelf,omitempty" tf:"manage_berkshelf,omitempty"`

	// The name of the stack.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the region where the stack will exist.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The ARN of an IAM role that the OpsWorks service will act as.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnRef *v1.Reference `json:"serviceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnSelector *v1.Selector `json:"serviceRoleArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Boolean value controlling whether the custom cookbook settings are enabled.
	// +kubebuilder:validation:Optional
	UseCustomCookbooks *bool `json:"useCustomCookbooks,omitempty" tf:"use_custom_cookbooks,omitempty"`

	// Boolean value controlling whether the standard OpsWorks security groups apply to created instances.
	// +kubebuilder:validation:Optional
	UseOpsworksSecurityGroups *bool `json:"useOpsworksSecurityGroups,omitempty" tf:"use_opsworks_security_groups,omitempty"`

	// ID of the VPC that this stack belongs to.
	// Defaults to the region's default VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// StackSpec defines the desired state of Stack
type StackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackParameters `json:"forProvider"`
}

// StackStatus defines the observed state of Stack.
type StackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stack is the Schema for the Stacks API. Provides an OpsWorks stack resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   StackSpec   `json:"spec"`
	Status StackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackList contains a list of Stacks
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stack `json:"items"`
}

// Repository type metadata.
var (
	Stack_Kind             = "Stack"
	Stack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stack_Kind}.String()
	Stack_KindAPIVersion   = Stack_Kind + "." + CRDGroupVersion.String()
	Stack_GroupVersionKind = CRDGroupVersion.WithKind(Stack_Kind)
)

func init() {
	SchemeBuilder.Register(&Stack{}, &StackList{})
}
