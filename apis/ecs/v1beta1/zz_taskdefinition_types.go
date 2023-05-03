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

type AuthorizationConfigObservation struct {

	// Access point ID to use. If an access point is specified, the root directory value will be relative to the directory set for the access point. If specified, transit encryption must be enabled in the EFSVolumeConfiguration.
	AccessPointID *string `json:"accessPointId,omitempty" tf:"access_point_id,omitempty"`

	// Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system. If enabled, transit encryption must be enabled in the EFSVolumeConfiguration. Valid values: ENABLED, DISABLED. If this parameter is omitted, the default value of DISABLED is used.
	IAM *string `json:"iam,omitempty" tf:"iam,omitempty"`
}

type AuthorizationConfigParameters struct {

	// Access point ID to use. If an access point is specified, the root directory value will be relative to the directory set for the access point. If specified, transit encryption must be enabled in the EFSVolumeConfiguration.
	// +kubebuilder:validation:Optional
	AccessPointID *string `json:"accessPointId,omitempty" tf:"access_point_id,omitempty"`

	// Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system. If enabled, transit encryption must be enabled in the EFSVolumeConfiguration. Valid values: ENABLED, DISABLED. If this parameter is omitted, the default value of DISABLED is used.
	// +kubebuilder:validation:Optional
	IAM *string `json:"iam,omitempty" tf:"iam,omitempty"`
}

type DockerVolumeConfigurationObservation struct {

	// If this value is true, the Docker volume is created if it does not already exist. Note: This field is only used if the scope is shared.
	Autoprovision *bool `json:"autoprovision,omitempty" tf:"autoprovision,omitempty"`

	// Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
	Driver *string `json:"driver,omitempty" tf:"driver,omitempty"`

	// Map of Docker driver specific options.
	DriverOpts map[string]*string `json:"driverOpts,omitempty" tf:"driver_opts,omitempty"`

	// Map of custom metadata to add to your Docker volume.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Scope for the Docker volume, which determines its lifecycle, either task or shared.  Docker volumes that are scoped to a task are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as shared persist after the task stops.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type DockerVolumeConfigurationParameters struct {

	// If this value is true, the Docker volume is created if it does not already exist. Note: This field is only used if the scope is shared.
	// +kubebuilder:validation:Optional
	Autoprovision *bool `json:"autoprovision,omitempty" tf:"autoprovision,omitempty"`

	// Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
	// +kubebuilder:validation:Optional
	Driver *string `json:"driver,omitempty" tf:"driver,omitempty"`

	// Map of Docker driver specific options.
	// +kubebuilder:validation:Optional
	DriverOpts map[string]*string `json:"driverOpts,omitempty" tf:"driver_opts,omitempty"`

	// Map of custom metadata to add to your Docker volume.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Scope for the Docker volume, which determines its lifecycle, either task or shared.  Docker volumes that are scoped to a task are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as shared persist after the task stops.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type EFSVolumeConfigurationObservation struct {

	// Configuration block for authorization for the Amazon EFS file system. Detailed below.
	AuthorizationConfig []AuthorizationConfigObservation `json:"authorizationConfig,omitempty" tf:"authorization_config,omitempty"`

	// ID of the EFS File System.
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Directory within the Amazon EFS file system to mount as the root directory inside the host. If this parameter is omitted, the root of the Amazon EFS volume will be used. Specifying / will have the same effect as omitting this parameter. This argument is ignored when using authorization_config.
	RootDirectory *string `json:"rootDirectory,omitempty" tf:"root_directory,omitempty"`

	// Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server. Transit encryption must be enabled if Amazon EFS IAM authorization is used. Valid values: ENABLED, DISABLED. If this parameter is omitted, the default value of DISABLED is used.
	TransitEncryption *string `json:"transitEncryption,omitempty" tf:"transit_encryption,omitempty"`

	// Port to use for transit encryption. If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon EFS mount helper uses.
	TransitEncryptionPort *float64 `json:"transitEncryptionPort,omitempty" tf:"transit_encryption_port,omitempty"`
}

type EFSVolumeConfigurationParameters struct {

	// Configuration block for authorization for the Amazon EFS file system. Detailed below.
	// +kubebuilder:validation:Optional
	AuthorizationConfig []AuthorizationConfigParameters `json:"authorizationConfig,omitempty" tf:"authorization_config,omitempty"`

	// ID of the EFS File System.
	// +kubebuilder:validation:Required
	FileSystemID *string `json:"fileSystemId" tf:"file_system_id,omitempty"`

	// Directory within the Amazon EFS file system to mount as the root directory inside the host. If this parameter is omitted, the root of the Amazon EFS volume will be used. Specifying / will have the same effect as omitting this parameter. This argument is ignored when using authorization_config.
	// +kubebuilder:validation:Optional
	RootDirectory *string `json:"rootDirectory,omitempty" tf:"root_directory,omitempty"`

	// Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server. Transit encryption must be enabled if Amazon EFS IAM authorization is used. Valid values: ENABLED, DISABLED. If this parameter is omitted, the default value of DISABLED is used.
	// +kubebuilder:validation:Optional
	TransitEncryption *string `json:"transitEncryption,omitempty" tf:"transit_encryption,omitempty"`

	// Port to use for transit encryption. If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon EFS mount helper uses.
	// +kubebuilder:validation:Optional
	TransitEncryptionPort *float64 `json:"transitEncryptionPort,omitempty" tf:"transit_encryption_port,omitempty"`
}

type EphemeralStorageObservation struct {

	// The total amount, in GiB, of ephemeral storage to set for the task. The minimum supported value is 21 GiB and the maximum supported value is 200 GiB.
	SizeInGib *float64 `json:"sizeInGib,omitempty" tf:"size_in_gib,omitempty"`
}

type EphemeralStorageParameters struct {

	// The total amount, in GiB, of ephemeral storage to set for the task. The minimum supported value is 21 GiB and the maximum supported value is 200 GiB.
	// +kubebuilder:validation:Required
	SizeInGib *float64 `json:"sizeInGib" tf:"size_in_gib,omitempty"`
}

type FSXWindowsFileServerVolumeConfigurationAuthorizationConfigObservation struct {

	// The authorization credential option to use. The authorization credential options can be provided using either the Amazon Resource Name (ARN) of an AWS Secrets Manager secret or AWS Systems Manager Parameter Store parameter. The ARNs refer to the stored credentials.
	CredentialsParameter *string `json:"credentialsParameter,omitempty" tf:"credentials_parameter,omitempty"`

	// A fully qualified domain name hosted by an AWS Directory Service Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`
}

type FSXWindowsFileServerVolumeConfigurationAuthorizationConfigParameters struct {

	// The authorization credential option to use. The authorization credential options can be provided using either the Amazon Resource Name (ARN) of an AWS Secrets Manager secret or AWS Systems Manager Parameter Store parameter. The ARNs refer to the stored credentials.
	// +kubebuilder:validation:Required
	CredentialsParameter *string `json:"credentialsParameter" tf:"credentials_parameter,omitempty"`

	// A fully qualified domain name hosted by an AWS Directory Service Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`
}

type FSXWindowsFileServerVolumeConfigurationObservation struct {

	// Configuration block for authorization for the Amazon FSx for Windows File Server file system detailed below.
	AuthorizationConfig []FSXWindowsFileServerVolumeConfigurationAuthorizationConfigObservation `json:"authorizationConfig,omitempty" tf:"authorization_config,omitempty"`

	// The Amazon FSx for Windows File Server file system ID to use.
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// The directory within the Amazon FSx for Windows File Server file system to mount as the root directory inside the host.
	RootDirectory *string `json:"rootDirectory,omitempty" tf:"root_directory,omitempty"`
}

type FSXWindowsFileServerVolumeConfigurationParameters struct {

	// Configuration block for authorization for the Amazon FSx for Windows File Server file system detailed below.
	// +kubebuilder:validation:Required
	AuthorizationConfig []FSXWindowsFileServerVolumeConfigurationAuthorizationConfigParameters `json:"authorizationConfig" tf:"authorization_config,omitempty"`

	// The Amazon FSx for Windows File Server file system ID to use.
	// +kubebuilder:validation:Required
	FileSystemID *string `json:"fileSystemId" tf:"file_system_id,omitempty"`

	// The directory within the Amazon FSx for Windows File Server file system to mount as the root directory inside the host.
	// +kubebuilder:validation:Required
	RootDirectory *string `json:"rootDirectory" tf:"root_directory,omitempty"`
}

type InferenceAcceleratorObservation struct {

	// Elastic Inference accelerator device name. The deviceName must also be referenced in a container definition as a ResourceRequirement.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Elastic Inference accelerator type to use.
	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`
}

type InferenceAcceleratorParameters struct {

	// Elastic Inference accelerator device name. The deviceName must also be referenced in a container definition as a ResourceRequirement.
	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// Elastic Inference accelerator type to use.
	// +kubebuilder:validation:Required
	DeviceType *string `json:"deviceType" tf:"device_type,omitempty"`
}

type ProxyConfigurationObservation struct {

	// Name of the container that will serve as the App Mesh proxy.
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Proxy type. The default value is APPMESH. The only supported value is APPMESH.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProxyConfigurationParameters struct {

	// Name of the container that will serve as the App Mesh proxy.
	// +kubebuilder:validation:Required
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`

	// Set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Proxy type. The default value is APPMESH. The only supported value is APPMESH.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuntimePlatformObservation struct {

	// Must be set to either X86_64 or ARM64; see cpu architecture
	CPUArchitecture *string `json:"cpuArchitecture,omitempty" tf:"cpu_architecture,omitempty"`

	// If the requires_compatibilities is FARGATE this field is required; must be set to a valid option from the operating system family in the runtime platform setting
	OperatingSystemFamily *string `json:"operatingSystemFamily,omitempty" tf:"operating_system_family,omitempty"`
}

type RuntimePlatformParameters struct {

	// Must be set to either X86_64 or ARM64; see cpu architecture
	// +kubebuilder:validation:Optional
	CPUArchitecture *string `json:"cpuArchitecture,omitempty" tf:"cpu_architecture,omitempty"`

	// If the requires_compatibilities is FARGATE this field is required; must be set to a valid option from the operating system family in the runtime platform setting
	// +kubebuilder:validation:Optional
	OperatingSystemFamily *string `json:"operatingSystemFamily,omitempty" tf:"operating_system_family,omitempty"`
}

type TaskDefinitionObservation struct {

	// Full ARN of the Task Definition (including both family and revision).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the Task Definition with the trailing revision removed. This may be useful for situations where the latest task definition is always desired. If a revision isn't specified, the latest ACTIVE revision is used. See the AWS documentation for details.
	ArnWithoutRevision *string `json:"arnWithoutRevision,omitempty" tf:"arn_without_revision,omitempty"`

	// Number of cpu units used by the task. If the requires_compatibilities is FARGATE this field is required.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// A list of valid container definitions provided as a single valid JSON document. Please note that you should only provide values that are part of the container definition document. For a detailed description of what parameters are available, see the Task Definition Parameters section from the official Developer Guide.
	ContainerDefinitions *string `json:"containerDefinitions,omitempty" tf:"container_definitions,omitempty"`

	// The amount of ephemeral storage to allocate for the task. This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate. See Ephemeral Storage.
	EphemeralStorage []EphemeralStorageObservation `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage,omitempty"`

	// ARN of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	// A unique name for your task definition.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block(s) with Inference Accelerators settings. Detailed below.
	InferenceAccelerator []InferenceAcceleratorObservation `json:"inferenceAccelerator,omitempty" tf:"inference_accelerator,omitempty"`

	// IPC resource namespace to be used for the containers in the task The valid values are host, task, and none.
	IpcMode *string `json:"ipcMode,omitempty" tf:"ipc_mode,omitempty"`

	// Amount (in MiB) of memory used by the task. If the requires_compatibilities is FARGATE this field is required.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`

	// Docker networking mode to use for the containers in the task. Valid values are none, bridge, awsvpc, and host.
	NetworkMode *string `json:"networkMode,omitempty" tf:"network_mode,omitempty"`

	// Process namespace to use for the containers in the task. The valid values are host and task.
	PidMode *string `json:"pidMode,omitempty" tf:"pid_mode,omitempty"`

	// Configuration block for rules that are taken into consideration during task placement. Maximum number of placement_constraints is 10. Detailed below.
	PlacementConstraints []TaskDefinitionPlacementConstraintsObservation `json:"placementConstraints,omitempty" tf:"placement_constraints,omitempty"`

	// Configuration block for the App Mesh proxy. Detailed below.
	ProxyConfiguration []ProxyConfigurationObservation `json:"proxyConfiguration,omitempty" tf:"proxy_configuration,omitempty"`

	// Set of launch types required by the task. The valid values are EC2 and FARGATE.
	RequiresCompatibilities []*string `json:"requiresCompatibilities,omitempty" tf:"requires_compatibilities,omitempty"`

	// Revision of the task in a particular family.
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

	// Configuration block for runtime_platform that containers in your task may use.
	RuntimePlatform []RuntimePlatformObservation `json:"runtimePlatform,omitempty" tf:"runtime_platform,omitempty"`

	// Whether to retain the old revision when the resource is destroyed or replacement is necessary. Default is false.
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn *string `json:"taskRoleArn,omitempty" tf:"task_role_arn,omitempty"`

	// Configuration block for volumes that containers in your task may use. Detailed below.
	Volume []VolumeObservation `json:"volume,omitempty" tf:"volume,omitempty"`
}

type TaskDefinitionParameters struct {

	// Number of cpu units used by the task. If the requires_compatibilities is FARGATE this field is required.
	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// A list of valid container definitions provided as a single valid JSON document. Please note that you should only provide values that are part of the container definition document. For a detailed description of what parameters are available, see the Task Definition Parameters section from the official Developer Guide.
	// +kubebuilder:validation:Optional
	ContainerDefinitions *string `json:"containerDefinitions,omitempty" tf:"container_definitions,omitempty"`

	// The amount of ephemeral storage to allocate for the task. This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate. See Ephemeral Storage.
	// +kubebuilder:validation:Optional
	EphemeralStorage []EphemeralStorageParameters `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage,omitempty"`

	// ARN of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	// Reference to a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnRef *v1.Reference `json:"executionRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnSelector *v1.Selector `json:"executionRoleArnSelector,omitempty" tf:"-"`

	// A unique name for your task definition.
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Configuration block(s) with Inference Accelerators settings. Detailed below.
	// +kubebuilder:validation:Optional
	InferenceAccelerator []InferenceAcceleratorParameters `json:"inferenceAccelerator,omitempty" tf:"inference_accelerator,omitempty"`

	// IPC resource namespace to be used for the containers in the task The valid values are host, task, and none.
	// +kubebuilder:validation:Optional
	IpcMode *string `json:"ipcMode,omitempty" tf:"ipc_mode,omitempty"`

	// Amount (in MiB) of memory used by the task. If the requires_compatibilities is FARGATE this field is required.
	// +kubebuilder:validation:Optional
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`

	// Docker networking mode to use for the containers in the task. Valid values are none, bridge, awsvpc, and host.
	// +kubebuilder:validation:Optional
	NetworkMode *string `json:"networkMode,omitempty" tf:"network_mode,omitempty"`

	// Process namespace to use for the containers in the task. The valid values are host and task.
	// +kubebuilder:validation:Optional
	PidMode *string `json:"pidMode,omitempty" tf:"pid_mode,omitempty"`

	// Configuration block for rules that are taken into consideration during task placement. Maximum number of placement_constraints is 10. Detailed below.
	// +kubebuilder:validation:Optional
	PlacementConstraints []TaskDefinitionPlacementConstraintsParameters `json:"placementConstraints,omitempty" tf:"placement_constraints,omitempty"`

	// Configuration block for the App Mesh proxy. Detailed below.
	// +kubebuilder:validation:Optional
	ProxyConfiguration []ProxyConfigurationParameters `json:"proxyConfiguration,omitempty" tf:"proxy_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Set of launch types required by the task. The valid values are EC2 and FARGATE.
	// +kubebuilder:validation:Optional
	RequiresCompatibilities []*string `json:"requiresCompatibilities,omitempty" tf:"requires_compatibilities,omitempty"`

	// Configuration block for runtime_platform that containers in your task may use.
	// +kubebuilder:validation:Optional
	RuntimePlatform []RuntimePlatformParameters `json:"runtimePlatform,omitempty" tf:"runtime_platform,omitempty"`

	// Whether to retain the old revision when the resource is destroyed or replacement is necessary. Default is false.
	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	// +kubebuilder:validation:Optional
	TaskRoleArn *string `json:"taskRoleArn,omitempty" tf:"task_role_arn,omitempty"`

	// Configuration block for volumes that containers in your task may use. Detailed below.
	// +kubebuilder:validation:Optional
	Volume []VolumeParameters `json:"volume,omitempty" tf:"volume,omitempty"`
}

type TaskDefinitionPlacementConstraintsObservation struct {

	// Cluster Query Language expression to apply to the constraint. For more information, see Cluster Query Language in the Amazon EC2 Container Service Developer Guide.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Type of constraint. Use memberOf to restrict selection to a group of valid candidates. Note that distinctInstance is not supported in task definitions.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TaskDefinitionPlacementConstraintsParameters struct {

	// Cluster Query Language expression to apply to the constraint. For more information, see Cluster Query Language in the Amazon EC2 Container Service Developer Guide.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Type of constraint. Use memberOf to restrict selection to a group of valid candidates. Note that distinctInstance is not supported in task definitions.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type VolumeObservation struct {

	// Configuration block to configure a docker volume. Detailed below.
	DockerVolumeConfiguration []DockerVolumeConfigurationObservation `json:"dockerVolumeConfiguration,omitempty" tf:"docker_volume_configuration,omitempty"`

	// Configuration block for an EFS volume. Detailed below.
	EFSVolumeConfiguration []EFSVolumeConfigurationObservation `json:"efsVolumeConfiguration,omitempty" tf:"efs_volume_configuration,omitempty"`

	// Configuration block for an FSX Windows File Server volume. Detailed below.
	FSXWindowsFileServerVolumeConfiguration []FSXWindowsFileServerVolumeConfigurationObservation `json:"fsxWindowsFileServerVolumeConfiguration,omitempty" tf:"fsx_windows_file_server_volume_configuration,omitempty"`

	// Path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
	HostPath *string `json:"hostPath,omitempty" tf:"host_path,omitempty"`

	// Name of the volume. This name is referenced in the sourceVolume
	// parameter of container definition in the mountPoints section.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type VolumeParameters struct {

	// Configuration block to configure a docker volume. Detailed below.
	// +kubebuilder:validation:Optional
	DockerVolumeConfiguration []DockerVolumeConfigurationParameters `json:"dockerVolumeConfiguration,omitempty" tf:"docker_volume_configuration,omitempty"`

	// Configuration block for an EFS volume. Detailed below.
	// +kubebuilder:validation:Optional
	EFSVolumeConfiguration []EFSVolumeConfigurationParameters `json:"efsVolumeConfiguration,omitempty" tf:"efs_volume_configuration,omitempty"`

	// Configuration block for an FSX Windows File Server volume. Detailed below.
	// +kubebuilder:validation:Optional
	FSXWindowsFileServerVolumeConfiguration []FSXWindowsFileServerVolumeConfigurationParameters `json:"fsxWindowsFileServerVolumeConfiguration,omitempty" tf:"fsx_windows_file_server_volume_configuration,omitempty"`

	// Path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
	// +kubebuilder:validation:Optional
	HostPath *string `json:"hostPath,omitempty" tf:"host_path,omitempty"`

	// Name of the volume. This name is referenced in the sourceVolume
	// parameter of container definition in the mountPoints section.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// TaskDefinitionSpec defines the desired state of TaskDefinition
type TaskDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TaskDefinitionParameters `json:"forProvider"`
}

// TaskDefinitionStatus defines the observed state of TaskDefinition.
type TaskDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TaskDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TaskDefinition is the Schema for the TaskDefinitions API. Manages a revision of an ECS task definition.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TaskDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.containerDefinitions)",message="containerDefinitions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.family)",message="family is a required parameter"
	Spec   TaskDefinitionSpec   `json:"spec"`
	Status TaskDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TaskDefinitionList contains a list of TaskDefinitions
type TaskDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TaskDefinition `json:"items"`
}

// Repository type metadata.
var (
	TaskDefinition_Kind             = "TaskDefinition"
	TaskDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TaskDefinition_Kind}.String()
	TaskDefinition_KindAPIVersion   = TaskDefinition_Kind + "." + CRDGroupVersion.String()
	TaskDefinition_GroupVersionKind = CRDGroupVersion.WithKind(TaskDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&TaskDefinition{}, &TaskDefinitionList{})
}
