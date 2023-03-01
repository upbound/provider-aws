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

type AccountRecoverySettingObservation struct {
}

type AccountRecoverySettingParameters struct {

	// List of Account Recovery Options of the following structure:
	// +kubebuilder:validation:Required
	RecoveryMechanism []RecoveryMechanismParameters `json:"recoveryMechanism" tf:"recovery_mechanism,omitempty"`
}

type AdminCreateUserConfigObservation struct {
}

type AdminCreateUserConfigParameters struct {

	// Set to True if only the administrator is allowed to create user profiles. Set to False if users can sign themselves up via an app.
	// +kubebuilder:validation:Optional
	AllowAdminCreateUserOnly *bool `json:"allowAdminCreateUserOnly,omitempty" tf:"allow_admin_create_user_only,omitempty"`

	// Invite message template structure. Detailed below.
	// +kubebuilder:validation:Optional
	InviteMessageTemplate []InviteMessageTemplateParameters `json:"inviteMessageTemplate,omitempty" tf:"invite_message_template,omitempty"`
}

type CustomEmailSenderObservation struct {
}

type CustomEmailSenderParameters struct {

	// The Lambda Amazon Resource Name of the Lambda function that Amazon Cognito triggers to send email notifications to users.
	// +kubebuilder:validation:Required
	LambdaArn *string `json:"lambdaArn" tf:"lambda_arn,omitempty"`

	// The Lambda version represents the signature of the "request" attribute in the "event" information Amazon Cognito passes to your custom email Lambda function. The only supported value is V1_0.
	// +kubebuilder:validation:Required
	LambdaVersion *string `json:"lambdaVersion" tf:"lambda_version,omitempty"`
}

type CustomSMSSenderObservation struct {
}

type CustomSMSSenderParameters struct {

	// The Lambda Amazon Resource Name of the Lambda function that Amazon Cognito triggers to send SMS notifications to users.
	// +kubebuilder:validation:Required
	LambdaArn *string `json:"lambdaArn" tf:"lambda_arn,omitempty"`

	// The Lambda version represents the signature of the "request" attribute in the "event" information Amazon Cognito passes to your custom SMS Lambda function. The only supported value is V1_0.
	// +kubebuilder:validation:Required
	LambdaVersion *string `json:"lambdaVersion" tf:"lambda_version,omitempty"`
}

type DeviceConfigurationObservation struct {
}

type DeviceConfigurationParameters struct {

	// Whether a challenge is required on a new device. Only applicable to a new device.
	// +kubebuilder:validation:Optional
	ChallengeRequiredOnNewDevice *bool `json:"challengeRequiredOnNewDevice,omitempty" tf:"challenge_required_on_new_device,omitempty"`

	// Whether a device is only remembered on user prompt. false equates to "Always" remember, true is "User Opt In," and not using a device_configuration block is "No."
	// +kubebuilder:validation:Optional
	DeviceOnlyRememberedOnUserPrompt *bool `json:"deviceOnlyRememberedOnUserPrompt,omitempty" tf:"device_only_remembered_on_user_prompt,omitempty"`
}

type EmailConfigurationObservation struct {
}

type EmailConfigurationParameters struct {

	// Email configuration set name from SES.
	// +kubebuilder:validation:Optional
	ConfigurationSet *string `json:"configurationSet,omitempty" tf:"configuration_set,omitempty"`

	// Email delivery method to use. COGNITO_DEFAULT for the default email functionality built into Cognito or DEVELOPER to use your Amazon SES configuration.
	// +kubebuilder:validation:Optional
	EmailSendingAccount *string `json:"emailSendingAccount,omitempty" tf:"email_sending_account,omitempty"`

	// Sender’s email address or sender’s display name with their email address (e.g., john@example.com, John Smith <john@example.com> or \"John Smith Ph.D.\" <john@example.com>). Escaped double quotes are required around display names that contain certain characters as specified in RFC 5322.
	// +kubebuilder:validation:Optional
	FromEmailAddress *string `json:"fromEmailAddress,omitempty" tf:"from_email_address,omitempty"`

	// REPLY-TO email address.
	// +kubebuilder:validation:Optional
	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty" tf:"reply_to_email_address,omitempty"`

	// ARN of the SES verified email identity to use. Required if email_sending_account is set to DEVELOPER.
	// +kubebuilder:validation:Optional
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`
}

type InviteMessageTemplateObservation struct {
}

type InviteMessageTemplateParameters struct {

	// Message template for email messages. Must contain {username} and {####} placeholders, for username and temporary password, respectively.
	// +kubebuilder:validation:Optional
	EmailMessage *string `json:"emailMessage,omitempty" tf:"email_message,omitempty"`

	// Subject line for email messages.
	// +kubebuilder:validation:Optional
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`

	// Message template for SMS messages. Must contain {username} and {####} placeholders, for username and temporary password, respectively.
	// +kubebuilder:validation:Optional
	SMSMessage *string `json:"smsMessage,omitempty" tf:"sms_message,omitempty"`
}

type LambdaConfigObservation struct {
}

type LambdaConfigParameters struct {

	// ARN of the lambda creating an authentication challenge.
	// +kubebuilder:validation:Optional
	CreateAuthChallenge *string `json:"createAuthChallenge,omitempty" tf:"create_auth_challenge,omitempty"`

	// A custom email sender AWS Lambda trigger. See custom_email_sender Below.
	// +kubebuilder:validation:Optional
	CustomEmailSender []CustomEmailSenderParameters `json:"customEmailSender,omitempty" tf:"custom_email_sender,omitempty"`

	// Custom Message AWS Lambda trigger.
	// +kubebuilder:validation:Optional
	CustomMessage *string `json:"customMessage,omitempty" tf:"custom_message,omitempty"`

	// A custom SMS sender AWS Lambda trigger. See custom_sms_sender Below.
	// +kubebuilder:validation:Optional
	CustomSMSSender []CustomSMSSenderParameters `json:"customSmsSender,omitempty" tf:"custom_sms_sender,omitempty"`

	// Defines the authentication challenge.
	// +kubebuilder:validation:Optional
	DefineAuthChallenge *string `json:"defineAuthChallenge,omitempty" tf:"define_auth_challenge,omitempty"`

	// The Amazon Resource Name of Key Management Service Customer master keys. Amazon Cognito uses the key to encrypt codes and temporary passwords sent to CustomEmailSender and CustomSMSSender.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Post-authentication AWS Lambda trigger.
	// +kubebuilder:validation:Optional
	PostAuthentication *string `json:"postAuthentication,omitempty" tf:"post_authentication,omitempty"`

	// Post-confirmation AWS Lambda trigger.
	// +kubebuilder:validation:Optional
	PostConfirmation *string `json:"postConfirmation,omitempty" tf:"post_confirmation,omitempty"`

	// Pre-authentication AWS Lambda trigger.
	// +kubebuilder:validation:Optional
	PreAuthentication *string `json:"preAuthentication,omitempty" tf:"pre_authentication,omitempty"`

	// Pre-registration AWS Lambda trigger.
	// +kubebuilder:validation:Optional
	PreSignUp *string `json:"preSignUp,omitempty" tf:"pre_sign_up,omitempty"`

	// Allow to customize identity token claims before token generation.
	// +kubebuilder:validation:Optional
	PreTokenGeneration *string `json:"preTokenGeneration,omitempty" tf:"pre_token_generation,omitempty"`

	// User migration Lambda config type.
	// +kubebuilder:validation:Optional
	UserMigration *string `json:"userMigration,omitempty" tf:"user_migration,omitempty"`

	// Verifies the authentication challenge response.
	// +kubebuilder:validation:Optional
	VerifyAuthChallengeResponse *string `json:"verifyAuthChallengeResponse,omitempty" tf:"verify_auth_challenge_response,omitempty"`
}

type NumberAttributeConstraintsObservation struct {
}

type NumberAttributeConstraintsParameters struct {

	// Maximum value of an attribute that is of the number data type.
	// +kubebuilder:validation:Optional
	MaxValue *string `json:"maxValue,omitempty" tf:"max_value,omitempty"`

	// Minimum value of an attribute that is of the number data type.
	// +kubebuilder:validation:Optional
	MinValue *string `json:"minValue,omitempty" tf:"min_value,omitempty"`
}

type PasswordPolicyObservation struct {
}

type PasswordPolicyParameters struct {

	// Minimum length of the password policy that you have set.
	// +kubebuilder:validation:Optional
	MinimumLength *float64 `json:"minimumLength,omitempty" tf:"minimum_length,omitempty"`

	// Whether you have required users to use at least one lowercase letter in their password.
	// +kubebuilder:validation:Optional
	RequireLowercase *bool `json:"requireLowercase,omitempty" tf:"require_lowercase,omitempty"`

	// Whether you have required users to use at least one number in their password.
	// +kubebuilder:validation:Optional
	RequireNumbers *bool `json:"requireNumbers,omitempty" tf:"require_numbers,omitempty"`

	// Whether you have required users to use at least one symbol in their password.
	// +kubebuilder:validation:Optional
	RequireSymbols *bool `json:"requireSymbols,omitempty" tf:"require_symbols,omitempty"`

	// Whether you have required users to use at least one uppercase letter in their password.
	// +kubebuilder:validation:Optional
	RequireUppercase *bool `json:"requireUppercase,omitempty" tf:"require_uppercase,omitempty"`

	// In the password policy you have set, refers to the number of days a temporary password is valid. If the user does not sign-in during this time, their password will need to be reset by an administrator.
	// +kubebuilder:validation:Optional
	TemporaryPasswordValidityDays *float64 `json:"temporaryPasswordValidityDays,omitempty" tf:"temporary_password_validity_days,omitempty"`
}

type RecoveryMechanismObservation struct {
}

type RecoveryMechanismParameters struct {

	// Name of the user pool.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Positive integer specifying priority of a method with 1 being the highest priority.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`
}

type SMSConfigurationObservation struct {
}

type SMSConfigurationParameters struct {

	// External ID used in IAM role trust relationships. For more information about using external IDs, see How to Use an External ID When Granting Access to Your AWS Resources to a Third Party.
	// +kubebuilder:validation:Required
	ExternalID *string `json:"externalId" tf:"external_id,omitempty"`

	// ARN of the Amazon SNS caller. This is usually the IAM role that you've given Cognito permission to assume.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SnsCallerArn *string `json:"snsCallerArn,omitempty" tf:"sns_caller_arn,omitempty"`

	// Reference to a Role in iam to populate snsCallerArn.
	// +kubebuilder:validation:Optional
	SnsCallerArnRef *v1.Reference `json:"snsCallerArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate snsCallerArn.
	// +kubebuilder:validation:Optional
	SnsCallerArnSelector *v1.Selector `json:"snsCallerArnSelector,omitempty" tf:"-"`

	// The AWS Region to use with Amazon SNS integration. You can choose the same Region as your user pool, or a supported Legacy Amazon SNS alternate Region. Amazon Cognito resources in the Asia Pacific (Seoul) AWS Region must use your Amazon SNS configuration in the Asia Pacific (Tokyo) Region. For more information, see SMS message settings for Amazon Cognito user pools.
	// +kubebuilder:validation:Optional
	SnsRegion *string `json:"snsRegion,omitempty" tf:"sns_region,omitempty"`
}

type SchemaObservation struct {
}

type SchemaParameters struct {

	// Attribute data type. Must be one of Boolean, Number, String, DateTime.
	// +kubebuilder:validation:Required
	AttributeDataType *string `json:"attributeDataType" tf:"attribute_data_type,omitempty"`

	// Whether the attribute type is developer only.
	// +kubebuilder:validation:Optional
	DeveloperOnlyAttribute *bool `json:"developerOnlyAttribute,omitempty" tf:"developer_only_attribute,omitempty"`

	// Whether the attribute can be changed once it has been created.
	// +kubebuilder:validation:Optional
	Mutable *bool `json:"mutable,omitempty" tf:"mutable,omitempty"`

	// Name of the user pool.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Configuration block for the constraints for an attribute of the number type. Detailed below.
	// +kubebuilder:validation:Optional
	NumberAttributeConstraints []NumberAttributeConstraintsParameters `json:"numberAttributeConstraints,omitempty" tf:"number_attribute_constraints,omitempty"`

	// Whether a user pool attribute is required. If the attribute is required and the user does not provide a value, registration or sign-in will fail.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// Constraints for an attribute of the string type. Detailed below.
	// +kubebuilder:validation:Optional
	StringAttributeConstraints []StringAttributeConstraintsParameters `json:"stringAttributeConstraints,omitempty" tf:"string_attribute_constraints,omitempty"`
}

type SoftwareTokenMfaConfigurationObservation struct {
}

type SoftwareTokenMfaConfigurationParameters struct {

	// Boolean whether to enable software token Multi-Factor (MFA) tokens, such as Time-based One-Time Password (TOTP). To disable software token MFA When sms_configuration is not present, the mfa_configuration argument must be set to OFF and the software_token_mfa_configuration configuration block must be fully removed.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type StringAttributeConstraintsObservation struct {
}

type StringAttributeConstraintsParameters struct {

	// Maximum length of an attribute value of the string type.
	// +kubebuilder:validation:Optional
	MaxLength *string `json:"maxLength,omitempty" tf:"max_length,omitempty"`

	// Minimum length of an attribute value of the string type.
	// +kubebuilder:validation:Optional
	MinLength *string `json:"minLength,omitempty" tf:"min_length,omitempty"`
}

type UserAttributeUpdateSettingsObservation struct {
}

type UserAttributeUpdateSettingsParameters struct {

	// A list of attributes requiring verification before update. If set, the provided value(s) must also be set in auto_verified_attributes. Valid values: email, phone_number.
	// +kubebuilder:validation:Required
	AttributesRequireVerificationBeforeUpdate []*string `json:"attributesRequireVerificationBeforeUpdate" tf:"attributes_require_verification_before_update,omitempty"`
}

type UserPoolAddOnsObservation struct {
}

type UserPoolAddOnsParameters struct {

	// Mode for advanced security, must be one of OFF, AUDIT or ENFORCED.
	// +kubebuilder:validation:Required
	AdvancedSecurityMode *string `json:"advancedSecurityMode" tf:"advanced_security_mode,omitempty"`
}

type UserPoolObservation struct {

	// ARN of the user pool.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Date the user pool was created.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// A custom domain name that you provide to Amazon Cognito. This parameter applies only if you use a custom domain to host the sign-up and sign-in pages for your application. For example: auth.example.com.
	CustomDomain *string `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`

	// Holds the domain prefix if the user pool has a domain associated with it.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// A number estimating the size of the user pool.
	EstimatedNumberOfUsers *float64 `json:"estimatedNumberOfUsers,omitempty" tf:"estimated_number_of_users,omitempty"`

	// ID of the user pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Date the user pool was last modified.
	LastModifiedDate *string `json:"lastModifiedDate,omitempty" tf:"last_modified_date,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type UserPoolParameters struct {

	// Configuration block to define which verified available method a user can use to recover their forgotten password. Detailed below.
	// +kubebuilder:validation:Optional
	AccountRecoverySetting []AccountRecoverySettingParameters `json:"accountRecoverySetting,omitempty" tf:"account_recovery_setting,omitempty"`

	// Configuration block for creating a new user profile. Detailed below.
	// +kubebuilder:validation:Optional
	AdminCreateUserConfig []AdminCreateUserConfigParameters `json:"adminCreateUserConfig,omitempty" tf:"admin_create_user_config,omitempty"`

	// Attributes supported as an alias for this user pool. Valid values: phone_number, email, or preferred_username. Conflicts with username_attributes.
	// +kubebuilder:validation:Optional
	AliasAttributes []*string `json:"aliasAttributes,omitempty" tf:"alias_attributes,omitempty"`

	// Attributes to be auto-verified. Valid values: email, phone_number.
	// +kubebuilder:validation:Optional
	AutoVerifiedAttributes []*string `json:"autoVerifiedAttributes,omitempty" tf:"auto_verified_attributes,omitempty"`

	// When active, DeletionProtection prevents accidental deletion of your user pool. Before you can delete a user pool that you have protected against deletion, you must deactivate this feature. Valid values are ACTIVE and INACTIVE, Default value is INACTIVE.
	// +kubebuilder:validation:Optional
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Configuration block for the user pool's device tracking. Detailed below.
	// +kubebuilder:validation:Optional
	DeviceConfiguration []DeviceConfigurationParameters `json:"deviceConfiguration,omitempty" tf:"device_configuration,omitempty"`

	// Configuration block for configuring email. Detailed below.
	// +kubebuilder:validation:Optional
	EmailConfiguration []EmailConfigurationParameters `json:"emailConfiguration,omitempty" tf:"email_configuration,omitempty"`

	// String representing the email verification message. Conflicts with verification_message_template configuration block email_message argument.
	// +kubebuilder:validation:Optional
	EmailVerificationMessage *string `json:"emailVerificationMessage,omitempty" tf:"email_verification_message,omitempty"`

	// String representing the email verification subject. Conflicts with verification_message_template configuration block email_subject argument.
	// +kubebuilder:validation:Optional
	EmailVerificationSubject *string `json:"emailVerificationSubject,omitempty" tf:"email_verification_subject,omitempty"`

	// Configuration block for the AWS Lambda triggers associated with the user pool. Detailed below.
	// +kubebuilder:validation:Optional
	LambdaConfig []LambdaConfigParameters `json:"lambdaConfig,omitempty" tf:"lambda_config,omitempty"`

	// Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of OFF. Valid values are OFF (MFA Tokens are not required), ON (MFA is required for all users to sign in; requires at least one of sms_configuration or software_token_mfa_configuration to be configured), or OPTIONAL (MFA Will be required only for individual users who have MFA Enabled; requires at least one of sms_configuration or software_token_mfa_configuration to be configured).
	// +kubebuilder:validation:Optional
	MfaConfiguration *string `json:"mfaConfiguration,omitempty" tf:"mfa_configuration,omitempty"`

	// Name of the user pool.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration blocked for information about the user pool password policy. Detailed below.
	// +kubebuilder:validation:Optional
	PasswordPolicy []PasswordPolicyParameters `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// String representing the SMS authentication message. The Message must contain the {####} placeholder, which will be replaced with the code.
	// +kubebuilder:validation:Optional
	SMSAuthenticationMessage *string `json:"smsAuthenticationMessage,omitempty" tf:"sms_authentication_message,omitempty"`

	// Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the taint command.
	// +kubebuilder:validation:Optional
	SMSConfiguration []SMSConfigurationParameters `json:"smsConfiguration,omitempty" tf:"sms_configuration,omitempty"`

	// String representing the SMS verification message. Conflicts with verification_message_template configuration block sms_message argument.
	// +kubebuilder:validation:Optional
	SMSVerificationMessage *string `json:"smsVerificationMessage,omitempty" tf:"sms_verification_message,omitempty"`

	// Configuration block for the schema attributes of a user pool. Detailed below. Schema attributes from the standard attribute set only need to be specified if they are different from the default configuration. Attributes can be added, but not modified or removed. Maximum of 50 attributes.
	// +kubebuilder:validation:Optional
	Schema []SchemaParameters `json:"schema,omitempty" tf:"schema,omitempty"`

	// Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
	// +kubebuilder:validation:Optional
	SoftwareTokenMfaConfiguration []SoftwareTokenMfaConfigurationParameters `json:"softwareTokenMfaConfiguration,omitempty" tf:"software_token_mfa_configuration,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for user attribute update settings. Detailed below.
	// +kubebuilder:validation:Optional
	UserAttributeUpdateSettings []UserAttributeUpdateSettingsParameters `json:"userAttributeUpdateSettings,omitempty" tf:"user_attribute_update_settings,omitempty"`

	// Configuration block for user pool add-ons to enable user pool advanced security mode features. Detailed below.
	// +kubebuilder:validation:Optional
	UserPoolAddOns []UserPoolAddOnsParameters `json:"userPoolAddOns,omitempty" tf:"user_pool_add_ons,omitempty"`

	// Whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with alias_attributes.
	// +kubebuilder:validation:Optional
	UsernameAttributes []*string `json:"usernameAttributes,omitempty" tf:"username_attributes,omitempty"`

	// Configuration block for username configuration. Detailed below.
	// +kubebuilder:validation:Optional
	UsernameConfiguration []UsernameConfigurationParameters `json:"usernameConfiguration,omitempty" tf:"username_configuration,omitempty"`

	// Configuration block for verification message templates. Detailed below.
	// +kubebuilder:validation:Optional
	VerificationMessageTemplate []VerificationMessageTemplateParameters `json:"verificationMessageTemplate,omitempty" tf:"verification_message_template,omitempty"`
}

type UsernameConfigurationObservation struct {
}

type UsernameConfigurationParameters struct {

	// Whether username case sensitivity will be applied for all users in the user pool through Cognito APIs.
	// +kubebuilder:validation:Required
	CaseSensitive *bool `json:"caseSensitive" tf:"case_sensitive,omitempty"`
}

type VerificationMessageTemplateObservation struct {
}

type VerificationMessageTemplateParameters struct {

	// Default email option. Must be either CONFIRM_WITH_CODE or CONFIRM_WITH_LINK. Defaults to CONFIRM_WITH_CODE.
	// +kubebuilder:validation:Optional
	DefaultEmailOption *string `json:"defaultEmailOption,omitempty" tf:"default_email_option,omitempty"`

	// Email message template. Must contain the {####} placeholder. Conflicts with email_verification_message argument.
	// +kubebuilder:validation:Optional
	EmailMessage *string `json:"emailMessage,omitempty" tf:"email_message,omitempty"`

	// Email message template for sending a confirmation link to the user, it must contain the {##Click Here##} placeholder.
	// +kubebuilder:validation:Optional
	EmailMessageByLink *string `json:"emailMessageByLink,omitempty" tf:"email_message_by_link,omitempty"`

	// Subject line for the email message template. Conflicts with email_verification_subject argument.
	// +kubebuilder:validation:Optional
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`

	// Subject line for the email message template for sending a confirmation link to the user.
	// +kubebuilder:validation:Optional
	EmailSubjectByLink *string `json:"emailSubjectByLink,omitempty" tf:"email_subject_by_link,omitempty"`

	// SMS message template. Must contain the {####} placeholder. Conflicts with sms_verification_message argument.
	// +kubebuilder:validation:Optional
	SMSMessage *string `json:"smsMessage,omitempty" tf:"sms_message,omitempty"`
}

// UserPoolSpec defines the desired state of UserPool
type UserPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPoolParameters `json:"forProvider"`
}

// UserPoolStatus defines the observed state of UserPool.
type UserPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPool is the Schema for the UserPools API. Provides a Cognito User Pool resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   UserPoolSpec   `json:"spec"`
	Status UserPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolList contains a list of UserPools
type UserPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPool `json:"items"`
}

// Repository type metadata.
var (
	UserPool_Kind             = "UserPool"
	UserPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPool_Kind}.String()
	UserPool_KindAPIVersion   = UserPool_Kind + "." + CRDGroupVersion.String()
	UserPool_GroupVersionKind = CRDGroupVersion.WithKind(UserPool_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPool{}, &UserPoolList{})
}
