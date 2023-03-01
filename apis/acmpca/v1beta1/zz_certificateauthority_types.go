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

type CertificateAuthorityConfigurationObservation struct {
}

type CertificateAuthorityConfigurationParameters struct {

	// Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the ACM PCA Documentation.
	// +kubebuilder:validation:Required
	KeyAlgorithm *string `json:"keyAlgorithm" tf:"key_algorithm,omitempty"`

	// Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the ACM PCA Documentation.
	// +kubebuilder:validation:Required
	SigningAlgorithm *string `json:"signingAlgorithm" tf:"signing_algorithm,omitempty"`

	// Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
	// +kubebuilder:validation:Required
	Subject []SubjectParameters `json:"subject" tf:"subject,omitempty"`
}

type CertificateAuthorityObservation struct {

	// ARN of the certificate authority.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
	CertificateSigningRequest *string `json:"certificateSigningRequest,omitempty" tf:"certificate_signing_request,omitempty"`

	// ARN of the certificate authority.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotAfter *string `json:"notAfter,omitempty" tf:"not_after,omitempty"`

	// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotBefore *string `json:"notBefore,omitempty" tf:"not_before,omitempty"`

	// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
	Serial *string `json:"serial,omitempty" tf:"serial,omitempty"`

	// (Deprecated use the enabled attribute instead) Status of the certificate authority.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CertificateAuthorityParameters struct {

	// Nested argument containing algorithms and certificate subject information. Defined below.
	// +kubebuilder:validation:Optional
	CertificateAuthorityConfiguration []CertificateAuthorityConfigurationParameters `json:"certificateAuthorityConfiguration,omitempty" tf:"certificate_authority_configuration,omitempty"`

	// Whether the certificate authority is enabled or disabled. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	// +kubebuilder:validation:Optional
	PermanentDeletionTimeInDays *float64 `json:"permanentDeletionTimeInDays,omitempty" tf:"permanent_deletion_time_in_days,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Nested argument containing revocation configuration. Defined below.
	// +kubebuilder:validation:Optional
	RevocationConfiguration []RevocationConfigurationParameters `json:"revocationConfiguration,omitempty" tf:"revocation_configuration,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of the certificate authority. Defaults to SUBORDINATE. Valid values: ROOT and SUBORDINATE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. Defaults to GENERAL_PURPOSE. Valid values: GENERAL_PURPOSE and SHORT_LIVED_CERTIFICATE.
	// +kubebuilder:validation:Optional
	UsageMode *string `json:"usageMode,omitempty" tf:"usage_mode,omitempty"`
}

type CrlConfigurationObservation struct {
}

type CrlConfigurationParameters struct {

	// Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point. Use this value if you don't want the name of your S3 bucket to be public. Must be less than or equal to 253 characters in length.
	// +kubebuilder:validation:Optional
	CustomCname *string `json:"customCname,omitempty" tf:"custom_cname,omitempty"`

	// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Number of days until a certificate expires. Must be between 1 and 5000.
	// +kubebuilder:validation:Required
	ExpirationInDays *float64 `json:"expirationInDays" tf:"expiration_in_days,omitempty"`

	// Name of the S3 bucket that contains the CRL. If you do not provide a value for the custom_cname argument, the name of your S3 bucket is placed into the CRL Distribution Points extension of the issued certificate. You must specify a bucket policy that allows ACM PCA to write the CRL to your bucket. Must be less than or equal to 255 characters in length.
	// +kubebuilder:validation:Optional
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// Determines whether the CRL will be publicly readable or privately held in the CRL Amazon S3 bucket. Defaults to PUBLIC_READ.
	// +kubebuilder:validation:Optional
	S3ObjectACL *string `json:"s3ObjectAcl,omitempty" tf:"s3_object_acl,omitempty"`
}

type OcspConfigurationObservation struct {
}

type OcspConfigurationParameters struct {

	// Boolean value that specifies whether a custom OCSP responder is enabled.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// CNAME specifying a customized OCSP domain. Note: The value of the CNAME must not include a protocol prefix such as "http://" or "https://".
	// +kubebuilder:validation:Optional
	OcspCustomCname *string `json:"ocspCustomCname,omitempty" tf:"ocsp_custom_cname,omitempty"`
}

type RevocationConfigurationObservation struct {
}

type RevocationConfigurationParameters struct {

	// Nested argument containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority. Defined below.
	// +kubebuilder:validation:Optional
	CrlConfiguration []CrlConfigurationParameters `json:"crlConfiguration,omitempty" tf:"crl_configuration,omitempty"`

	// Nested argument containing configuration of
	// the custom OCSP responder endpoint. Defined below.
	// +kubebuilder:validation:Optional
	OcspConfiguration []OcspConfigurationParameters `json:"ocspConfiguration,omitempty" tf:"ocsp_configuration,omitempty"`
}

type SubjectObservation struct {
}

type SubjectParameters struct {

	// Fully qualified domain name (FQDN) associated with the certificate subject. Must be less than or equal to 64 characters in length.
	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// Two digit code that specifies the country in which the certificate subject located. Must be less than or equal to 2 characters in length.
	// +kubebuilder:validation:Optional
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// Disambiguating information for the certificate subject. Must be less than or equal to 64 characters in length.
	// +kubebuilder:validation:Optional
	DistinguishedNameQualifier *string `json:"distinguishedNameQualifier,omitempty" tf:"distinguished_name_qualifier,omitempty"`

	// Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third. Must be less than or equal to 3 characters in length.
	// +kubebuilder:validation:Optional
	GenerationQualifier *string `json:"generationQualifier,omitempty" tf:"generation_qualifier,omitempty"`

	// First name. Must be less than or equal to 16 characters in length.
	// +kubebuilder:validation:Optional
	GivenName *string `json:"givenName,omitempty" tf:"given_name,omitempty"`

	// Concatenation that typically contains the first letter of the given_name, the first letter of the middle name if one exists, and the first letter of the surname. Must be less than or equal to 5 characters in length.
	// +kubebuilder:validation:Optional
	Initials *string `json:"initials,omitempty" tf:"initials,omitempty"`

	// Locality (such as a city or town) in which the certificate subject is located. Must be less than or equal to 128 characters in length.
	// +kubebuilder:validation:Optional
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// Legal name of the organization with which the certificate subject is affiliated. Must be less than or equal to 64 characters in length.
	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated. Must be less than or equal to 64 characters in length.
	// +kubebuilder:validation:Optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit,omitempty"`

	// Typically a shortened version of a longer given_name. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza. Must be less than or equal to 128 characters in length.
	// +kubebuilder:validation:Optional
	Pseudonym *string `json:"pseudonym,omitempty" tf:"pseudonym,omitempty"`

	// State in which the subject of the certificate is located. Must be less than or equal to 128 characters in length.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Family name. In the US and the UK for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first. Must be less than or equal to 40 characters in length.
	// +kubebuilder:validation:Optional
	Surname *string `json:"surname,omitempty" tf:"surname,omitempty"`

	// Title such as Mr. or Ms. which is pre-pended to the name to refer formally to the certificate subject. Must be less than or equal to 64 characters in length.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// CertificateAuthoritySpec defines the desired state of CertificateAuthority
type CertificateAuthoritySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateAuthorityParameters `json:"forProvider"`
}

// CertificateAuthorityStatus defines the observed state of CertificateAuthority.
type CertificateAuthorityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateAuthorityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateAuthority is the Schema for the CertificateAuthoritys API. Provides a resource to manage AWS Certificate Manager Private Certificate Authorities
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CertificateAuthority struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.certificateAuthorityConfiguration)",message="certificateAuthorityConfiguration is a required parameter"
	Spec   CertificateAuthoritySpec   `json:"spec"`
	Status CertificateAuthorityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateAuthorityList contains a list of CertificateAuthoritys
type CertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateAuthority `json:"items"`
}

// Repository type metadata.
var (
	CertificateAuthority_Kind             = "CertificateAuthority"
	CertificateAuthority_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateAuthority_Kind}.String()
	CertificateAuthority_KindAPIVersion   = CertificateAuthority_Kind + "." + CRDGroupVersion.String()
	CertificateAuthority_GroupVersionKind = CRDGroupVersion.WithKind(CertificateAuthority_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateAuthority{}, &CertificateAuthorityList{})
}
