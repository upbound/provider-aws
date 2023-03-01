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

type AppSourceObservation struct {
}

type AppSourceParameters struct {

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

	// The URL where the app resource can be found.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Username to use when authenticating to the source.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ApplicationObservation struct {

	// The id of the application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ApplicationParameters struct {

	// SCM configuration of the app as described below.
	// +kubebuilder:validation:Optional
	AppSource []AppSourceParameters `json:"appSource,omitempty" tf:"app_source,omitempty"`

	// Run bundle install when deploying for application of type rails.
	// +kubebuilder:validation:Optional
	AutoBundleOnDeploy *string `json:"autoBundleOnDeploy,omitempty" tf:"auto_bundle_on_deploy,omitempty"`

	// Specify activity and workflow workers for your app using the aws-flow gem.
	// +kubebuilder:validation:Optional
	AwsFlowRubySettings *string `json:"awsFlowRubySettings,omitempty" tf:"aws_flow_ruby_settings,omitempty"`

	// The data source's ARN.
	// +kubebuilder:validation:Optional
	DataSourceArn *string `json:"dataSourceArn,omitempty" tf:"data_source_arn,omitempty"`

	// The database name.
	// +kubebuilder:validation:Optional
	DataSourceDatabaseName *string `json:"dataSourceDatabaseName,omitempty" tf:"data_source_database_name,omitempty"`

	// The data source's type one of AutoSelectOpsworksMysqlInstance, OpsworksMysqlInstance, or RdsDbInstance.
	// +kubebuilder:validation:Optional
	DataSourceType *string `json:"dataSourceType,omitempty" tf:"data_source_type,omitempty"`

	// A description of the app.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Subfolder for the document root for application of type rails.
	// +kubebuilder:validation:Optional
	DocumentRoot *string `json:"documentRoot,omitempty" tf:"document_root,omitempty"`

	// A list of virtual host alias.
	// +kubebuilder:validation:Optional
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// Whether to enable SSL for the app. This must be set in order to let ssl_configuration.private_key, ssl_configuration.certificate and ssl_configuration.chain take effect.
	// +kubebuilder:validation:Optional
	EnableSSL *bool `json:"enableSsl,omitempty" tf:"enable_ssl,omitempty"`

	// Object to define environment variables.  Object is described below.
	// +kubebuilder:validation:Optional
	Environment []EnvironmentParameters `json:"environment,omitempty" tf:"environment,omitempty"`

	// A human-readable name for the application.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Rails environment for application of type rails.
	// +kubebuilder:validation:Optional
	RailsEnv *string `json:"railsEnv,omitempty" tf:"rails_env,omitempty"`

	// The SSL configuration of the app. Object is described below.
	// +kubebuilder:validation:Optional
	SSLConfiguration []SSLConfigurationParameters `json:"sslConfiguration,omitempty" tf:"ssl_configuration,omitempty"`

	// A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
	// +kubebuilder:validation:Optional
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`

	// ID of the stack the application will belong to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opsworks/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Reference to a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDRef *v1.Reference `json:"stackIdRef,omitempty" tf:"-"`

	// Selector for a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDSelector *v1.Selector `json:"stackIdSelector,omitempty" tf:"-"`

	// Opsworks application type. One of aws-flow-ruby, java, rails, php, nodejs, static or other.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EnvironmentObservation struct {
}

type EnvironmentParameters struct {

	// Variable name.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Set visibility of the variable value to true or false.
	// +kubebuilder:validation:Optional
	Secure *bool `json:"secure,omitempty" tf:"secure,omitempty"`

	// Variable value.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type SSLConfigurationObservation struct {
}

type SSLConfigurationParameters struct {

	// The contents of the certificate's domain.crt file.
	// +kubebuilder:validation:Required
	Certificate *string `json:"certificate" tf:"certificate,omitempty"`

	// Can be used to specify an intermediate certificate authority key or client authentication.
	// +kubebuilder:validation:Optional
	Chain *string `json:"chain,omitempty" tf:"chain,omitempty"`

	// The private key; the contents of the certificate's domain.key file.
	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. Provides an OpsWorks application resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.type)",message="type is a required parameter"
	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
