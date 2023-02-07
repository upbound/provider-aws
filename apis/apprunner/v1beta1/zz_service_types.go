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

type AuthenticationConfigurationObservation struct {
}

type AuthenticationConfigurationParameters struct {

	// ARN of the IAM role that grants the App Runner service access to a source repository. Required for ECR image repositories (but not for ECR Public)
	// +kubebuilder:validation:Optional
	AccessRoleArn *string `json:"accessRoleArn,omitempty" tf:"access_role_arn,omitempty"`

	// ARN of the App Runner connection that enables the App Runner service to connect to a source repository. Required for GitHub code repositories.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apprunner/v1beta1.Connection
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ConnectionArn *string `json:"connectionArn,omitempty" tf:"connection_arn,omitempty"`

	// Reference to a Connection in apprunner to populate connectionArn.
	// +kubebuilder:validation:Optional
	ConnectionArnRef *v1.Reference `json:"connectionArnRef,omitempty" tf:"-"`

	// Selector for a Connection in apprunner to populate connectionArn.
	// +kubebuilder:validation:Optional
	ConnectionArnSelector *v1.Selector `json:"connectionArnSelector,omitempty" tf:"-"`
}

type CodeConfigurationObservation struct {
}

type CodeConfigurationParameters struct {

	// Basic configuration for building and running the App Runner service. Use this parameter to quickly launch an App Runner service without providing an apprunner.yaml file in the source code repository (or ignoring the file if it exists). See Code Configuration Values below for more details.
	// +kubebuilder:validation:Optional
	CodeConfigurationValues []CodeConfigurationValuesParameters `json:"codeConfigurationValues,omitempty" tf:"code_configuration_values,omitempty"`

	// Source of the App Runner configuration. Valid values: REPOSITORY, API. Values are interpreted as follows:
	// +kubebuilder:validation:Required
	ConfigurationSource *string `json:"configurationSource" tf:"configuration_source,omitempty"`
}

type CodeConfigurationValuesObservation struct {
}

type CodeConfigurationValuesParameters struct {

	// Command App Runner runs to build your application.
	// +kubebuilder:validation:Optional
	BuildCommand *string `json:"buildCommand,omitempty" tf:"build_command,omitempty"`

	// Port that your application listens to in the container. Defaults to "8080".
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Runtime environment type for building and running an App Runner service. Represents a programming language runtime. Valid values: PYTHON_3, NODEJS_12, NODEJS_14, NODEJS_16, CORRETTO_8, CORRETTO_11, GO_1, DOTNET_6, PHP_81, RUBY_31.
	// +kubebuilder:validation:Required
	Runtime *string `json:"runtime" tf:"runtime,omitempty"`

	// Secrets and parameters available to your service as environment variables. A map of key/value pairs.
	// +kubebuilder:validation:Optional
	RuntimeEnvironmentSecrets map[string]*string `json:"runtimeEnvironmentSecrets,omitempty" tf:"runtime_environment_secrets,omitempty"`

	// Environment variables available to your running App Runner service. A map of key/value pairs. Keys with a prefix of AWSAPPRUNNER are reserved for system use and aren't valid.
	// +kubebuilder:validation:Optional
	RuntimeEnvironmentVariables map[string]*string `json:"runtimeEnvironmentVariables,omitempty" tf:"runtime_environment_variables,omitempty"`

	// Command App Runner runs to start the application in the source image. If specified, this command overrides the Docker image’s default start command.
	// +kubebuilder:validation:Optional
	StartCommand *string `json:"startCommand,omitempty" tf:"start_command,omitempty"`
}

type CodeRepositoryObservation struct {
}

type CodeRepositoryParameters struct {

	// Configuration for building and running the service from a source code repository. See Code Configuration below for more details.
	// +kubebuilder:validation:Optional
	CodeConfiguration []CodeConfigurationParameters `json:"codeConfiguration,omitempty" tf:"code_configuration,omitempty"`

	// Location of the repository that contains the source code.
	// +kubebuilder:validation:Required
	RepositoryURL *string `json:"repositoryUrl" tf:"repository_url,omitempty"`

	// Version that should be used within the source code repository. See Source Code Version below for more details.
	// +kubebuilder:validation:Required
	SourceCodeVersion []SourceCodeVersionParameters `json:"sourceCodeVersion" tf:"source_code_version,omitempty"`
}

type EgressConfigurationObservation struct {
}

type EgressConfigurationParameters struct {

	// Type of egress configuration.Set to DEFAULT for access to resources hosted on public networks.Set to VPC to associate your service to a custom VPC specified by VpcConnectorArn.
	// +kubebuilder:validation:Optional
	EgressType *string `json:"egressType,omitempty" tf:"egress_type,omitempty"`

	// ARN of the App Runner VPC connector that you want to associate with your App Runner service. Only valid when EgressType = VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apprunner/v1beta1.VPCConnector
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	VPCConnectorArn *string `json:"vpcConnectorArn,omitempty" tf:"vpc_connector_arn,omitempty"`

	// Reference to a VPCConnector in apprunner to populate vpcConnectorArn.
	// +kubebuilder:validation:Optional
	VPCConnectorArnRef *v1.Reference `json:"vpcConnectorArnRef,omitempty" tf:"-"`

	// Selector for a VPCConnector in apprunner to populate vpcConnectorArn.
	// +kubebuilder:validation:Optional
	VPCConnectorArnSelector *v1.Selector `json:"vpcConnectorArnSelector,omitempty" tf:"-"`
}

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {

	// ARN of the KMS key used for encryption.
	// +kubebuilder:validation:Required
	KMSKey *string `json:"kmsKey" tf:"kms_key,omitempty"`
}

type HealthCheckConfigurationObservation struct {
}

type HealthCheckConfigurationParameters struct {

	// Number of consecutive checks that must succeed before App Runner decides that the service is healthy. Defaults to 1. Minimum value of 1. Maximum value of 20.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// Time interval, in seconds, between health checks. Defaults to 5. Minimum value of 1. Maximum value of 20.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// URL to send requests to for health checks. Defaults to /. Minimum length of 0. Maximum length of 51200.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// IP protocol that App Runner uses to perform health checks for your service. Valid values: TCP, HTTP. Defaults to TCP. If you set protocol to HTTP, App Runner sends health check requests to the HTTP path specified by path.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Time, in seconds, to wait for a health check response before deciding it failed. Defaults to 2. Minimum value of  1. Maximum value of 20.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Number of consecutive checks that must fail before App Runner decides that the service is unhealthy. Defaults to 5. Minimum value of  1. Maximum value of 20.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type ImageConfigurationObservation struct {
}

type ImageConfigurationParameters struct {

	// Port that your application listens to in the container. Defaults to "8080".
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Secrets and parameters available to your service as environment variables. A map of key/value pairs.
	// +kubebuilder:validation:Optional
	RuntimeEnvironmentSecrets map[string]*string `json:"runtimeEnvironmentSecrets,omitempty" tf:"runtime_environment_secrets,omitempty"`

	// Environment variables available to your running App Runner service. A map of key/value pairs. Keys with a prefix of AWSAPPRUNNER are reserved for system use and aren't valid.
	// +kubebuilder:validation:Optional
	RuntimeEnvironmentVariables map[string]*string `json:"runtimeEnvironmentVariables,omitempty" tf:"runtime_environment_variables,omitempty"`

	// Command App Runner runs to start the application in the source image. If specified, this command overrides the Docker image’s default start command.
	// +kubebuilder:validation:Optional
	StartCommand *string `json:"startCommand,omitempty" tf:"start_command,omitempty"`
}

type ImageRepositoryObservation struct {
}

type ImageRepositoryParameters struct {

	// Configuration for running the identified image. See Image Configuration below for more details.
	// +kubebuilder:validation:Optional
	ImageConfiguration []ImageConfigurationParameters `json:"imageConfiguration,omitempty" tf:"image_configuration,omitempty"`

	// Identifier of an image. For an image in Amazon Elastic Container Registry (Amazon ECR), this is an image name. For the
	// image name format, see Pulling an image in the Amazon ECR User Guide.
	// +kubebuilder:validation:Required
	ImageIdentifier *string `json:"imageIdentifier" tf:"image_identifier,omitempty"`

	// Type of the image repository. This reflects the repository provider and whether the repository is private or public. Valid values: ECR , ECR_PUBLIC.
	// +kubebuilder:validation:Required
	ImageRepositoryType *string `json:"imageRepositoryType" tf:"image_repository_type,omitempty"`
}

type IngressConfigurationObservation struct {
}

type IngressConfigurationParameters struct {

	// Specifies whether your App Runner service is publicly accessible. To make the service publicly accessible set it to True. To make the service privately accessible, from only within an Amazon VPC set it to False.
	// +kubebuilder:validation:Optional
	IsPubliclyAccessible *bool `json:"isPubliclyAccessible,omitempty" tf:"is_publicly_accessible,omitempty"`
}

type InstanceConfigurationObservation struct {
}

type InstanceConfigurationParameters struct {

	// Number of CPU units reserved for each instance of your App Runner service represented as a String. Defaults to 1024. Valid values: 1024|2048|(1|2) vCPU.
	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// ARN of an IAM role that provides permissions to your App Runner service. These are permissions that your code needs when it calls any AWS APIs.
	// +kubebuilder:validation:Optional
	InstanceRoleArn *string `json:"instanceRoleArn,omitempty" tf:"instance_role_arn,omitempty"`

	// Amount of memory, in MB or GB, reserved for each instance of your App Runner service. Defaults to 2048. Valid values: 2048|3072|4096|(2|3|4) GB.
	// +kubebuilder:validation:Optional
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type NetworkConfigurationObservation struct {
}

type NetworkConfigurationParameters struct {

	// Network configuration settings for outbound message traffic. See Egress Configuration below for more details.
	// +kubebuilder:validation:Optional
	EgressConfiguration []EgressConfigurationParameters `json:"egressConfiguration,omitempty" tf:"egress_configuration,omitempty"`

	// Network configuration settings for inbound network traffic. See Ingress Configuration below for more details.
	// +kubebuilder:validation:Optional
	IngressConfiguration []IngressConfigurationParameters `json:"ingressConfiguration,omitempty" tf:"ingress_configuration,omitempty"`
}

type ObservabilityConfigurationObservation struct {
}

type ObservabilityConfigurationParameters struct {

	// ARN of the observability configuration that is associated with the service. Specified only when observability_enabled is true.
	// +kubebuilder:validation:Optional
	ObservabilityConfigurationArn *string `json:"observabilityConfigurationArn,omitempty" tf:"observability_configuration_arn,omitempty"`

	// When true, an observability configuration resource is associated with the service.
	// +kubebuilder:validation:Required
	ObservabilityEnabled *bool `json:"observabilityEnabled" tf:"observability_enabled,omitempty"`
}

type ServiceObservation struct {

	// ARN of the App Runner service.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An alphanumeric ID that App Runner generated for this service. Unique within the AWS Region.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.
	ServiceURL *string `json:"serviceUrl,omitempty" tf:"service_url,omitempty"`

	// Current state of the App Runner service.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ServiceParameters struct {

	// ARN of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner associates the latest revision of a default auto scaling configuration.
	// +kubebuilder:validation:Optional
	AutoScalingConfigurationArn *string `json:"autoScalingConfigurationArn,omitempty" tf:"auto_scaling_configuration_arn,omitempty"`

	// (Forces new resource) An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default, App Runner uses an AWS managed CMK. See Encryption Configuration below for more details.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// (Forces new resource) Settings of the health check that AWS App Runner performs to monitor the health of your service. See Health Check Configuration below for more details.
	// +kubebuilder:validation:Optional
	HealthCheckConfiguration []HealthCheckConfigurationParameters `json:"healthCheckConfiguration,omitempty" tf:"health_check_configuration,omitempty"`

	// The runtime configuration of instances (scaling units) of the App Runner service. See Instance Configuration below for more details.
	// +kubebuilder:validation:Optional
	InstanceConfiguration []InstanceConfigurationParameters `json:"instanceConfiguration,omitempty" tf:"instance_configuration,omitempty"`

	// Configuration settings related to network traffic of the web application that the App Runner service runs. See Network Configuration below for more details.
	// +kubebuilder:validation:Optional
	NetworkConfiguration []NetworkConfigurationParameters `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// The observability configuration of your service. See Observability Configuration below for more details.
	// +kubebuilder:validation:Optional
	ObservabilityConfiguration []ObservabilityConfigurationParameters `json:"observabilityConfiguration,omitempty" tf:"observability_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// (Forces new resource) Name of the service.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// The source to deploy to the App Runner service. Can be a code or an image repository. See Source Configuration below for more details.
	// +kubebuilder:validation:Required
	SourceConfiguration []SourceConfigurationParameters `json:"sourceConfiguration" tf:"source_configuration,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SourceCodeVersionObservation struct {
}

type SourceCodeVersionParameters struct {

	// Type of version identifier. For a git-based repository, branches represent versions. Valid values: BRANCH.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Source code version. For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type SourceConfigurationObservation struct {
}

type SourceConfigurationParameters struct {

	// Describes resources needed to authenticate access to some source repositories. See Authentication Configuration below for more details.
	// +kubebuilder:validation:Optional
	AuthenticationConfiguration []AuthenticationConfigurationParameters `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`

	// Whether continuous integration from the source repository is enabled for the App Runner service. If set to true, each repository change (source code commit or new image version) starts a deployment. Defaults to true.
	// +kubebuilder:validation:Optional
	AutoDeploymentsEnabled *bool `json:"autoDeploymentsEnabled,omitempty" tf:"auto_deployments_enabled,omitempty"`

	// Description of a source code repository. See Code Repository below for more details.
	// +kubebuilder:validation:Optional
	CodeRepository []CodeRepositoryParameters `json:"codeRepository,omitempty" tf:"code_repository,omitempty"`

	// Description of a source image repository. See Image Repository below for more details.
	// +kubebuilder:validation:Optional
	ImageRepository []ImageRepositoryParameters `json:"imageRepository,omitempty" tf:"image_repository,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API. Manages an App Runner Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec"`
	Status            ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}