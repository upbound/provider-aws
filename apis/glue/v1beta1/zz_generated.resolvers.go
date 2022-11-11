/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CatalogTable.
func (mg *CatalogTable) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatabaseNameRef,
		Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
		To: reference.To{
			List:    &CatalogDatabaseList{},
			Managed: &CatalogDatabase{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Connection.
func (mg *Connection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PhysicalConnectionRequirements); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].AvailabilityZone),
			Extract:      resource.ExtractParamPath("availability_zone", false),
			Reference:    mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].AvailabilityZoneRef,
			Selector:     mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].AvailabilityZoneSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].AvailabilityZone")
		}
		mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].AvailabilityZone = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].AvailabilityZoneRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.PhysicalConnectionRequirements); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].SubnetID")
		}
		mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PhysicalConnectionRequirements[i3].SubnetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this DataCatalogEncryptionSettings.
func (mg *DataCatalogEncryptionSettings) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DataCatalogEncryptionSettings); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].ConnectionPasswordEncryption); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].ConnectionPasswordEncryption[i4].AwsKMSKeyID),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].ConnectionPasswordEncryption[i4].AwsKMSKeyIDRef,
				Selector:     mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].ConnectionPasswordEncryption[i4].AwsKMSKeyIDSelector,
				To: reference.To{
					List:    &v1beta11.KeyList{},
					Managed: &v1beta11.Key{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].ConnectionPasswordEncryption[i4].AwsKMSKeyID")
			}
			mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].ConnectionPasswordEncryption[i4].AwsKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].ConnectionPasswordEncryption[i4].AwsKMSKeyIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.DataCatalogEncryptionSettings); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].EncryptionAtRest); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].EncryptionAtRest[i4].SseAwsKMSKeyID),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].EncryptionAtRest[i4].SseAwsKMSKeyIDRef,
				Selector:     mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].EncryptionAtRest[i4].SseAwsKMSKeyIDSelector,
				To: reference.To{
					List:    &v1beta11.KeyList{},
					Managed: &v1beta11.Key{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].EncryptionAtRest[i4].SseAwsKMSKeyID")
			}
			mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].EncryptionAtRest[i4].SseAwsKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.DataCatalogEncryptionSettings[i3].EncryptionAtRest[i4].SseAwsKMSKeyIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Job.
func (mg *Job) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityConfiguration.
func (mg *SecurityConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptionConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.EncryptionConfiguration[i3].CloudwatchEncryption); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionConfiguration[i3].CloudwatchEncryption[i4].KMSKeyArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.EncryptionConfiguration[i3].CloudwatchEncryption[i4].KMSKeyArnRef,
				Selector:     mg.Spec.ForProvider.EncryptionConfiguration[i3].CloudwatchEncryption[i4].KMSKeyArnSelector,
				To: reference.To{
					List:    &v1beta11.KeyList{},
					Managed: &v1beta11.Key{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionConfiguration[i3].CloudwatchEncryption[i4].KMSKeyArn")
			}
			mg.Spec.ForProvider.EncryptionConfiguration[i3].CloudwatchEncryption[i4].KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EncryptionConfiguration[i3].CloudwatchEncryption[i4].KMSKeyArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptionConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.EncryptionConfiguration[i3].JobBookmarksEncryption); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionConfiguration[i3].JobBookmarksEncryption[i4].KMSKeyArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.EncryptionConfiguration[i3].JobBookmarksEncryption[i4].KMSKeyArnRef,
				Selector:     mg.Spec.ForProvider.EncryptionConfiguration[i3].JobBookmarksEncryption[i4].KMSKeyArnSelector,
				To: reference.To{
					List:    &v1beta11.KeyList{},
					Managed: &v1beta11.Key{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionConfiguration[i3].JobBookmarksEncryption[i4].KMSKeyArn")
			}
			mg.Spec.ForProvider.EncryptionConfiguration[i3].JobBookmarksEncryption[i4].KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EncryptionConfiguration[i3].JobBookmarksEncryption[i4].KMSKeyArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptionConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.EncryptionConfiguration[i3].S3Encryption); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionConfiguration[i3].S3Encryption[i4].KMSKeyArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.EncryptionConfiguration[i3].S3Encryption[i4].KMSKeyArnRef,
				Selector:     mg.Spec.ForProvider.EncryptionConfiguration[i3].S3Encryption[i4].KMSKeyArnSelector,
				To: reference.To{
					List:    &v1beta11.KeyList{},
					Managed: &v1beta11.Key{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionConfiguration[i3].S3Encryption[i4].KMSKeyArn")
			}
			mg.Spec.ForProvider.EncryptionConfiguration[i3].S3Encryption[i4].KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EncryptionConfiguration[i3].S3Encryption[i4].KMSKeyArnRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Trigger.
func (mg *Trigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Actions); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Actions[i3].JobName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Actions[i3].JobNameRef,
			Selector:     mg.Spec.ForProvider.Actions[i3].JobNameSelector,
			To: reference.To{
				List:    &JobList{},
				Managed: &Job{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Actions[i3].JobName")
		}
		mg.Spec.ForProvider.Actions[i3].JobName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Actions[i3].JobNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Predicate); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Predicate[i3].Conditions); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Predicate[i3].Conditions[i4].JobName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Predicate[i3].Conditions[i4].JobNameRef,
				Selector:     mg.Spec.ForProvider.Predicate[i3].Conditions[i4].JobNameSelector,
				To: reference.To{
					List:    &JobList{},
					Managed: &Job{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Predicate[i3].Conditions[i4].JobName")
			}
			mg.Spec.ForProvider.Predicate[i3].Conditions[i4].JobName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Predicate[i3].Conditions[i4].JobNameRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this UserDefinedFunction.
func (mg *UserDefinedFunction) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatabaseNameRef,
		Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
		To: reference.To{
			List:    &CatalogDatabaseList{},
			Managed: &CatalogDatabase{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference

	return nil
}
