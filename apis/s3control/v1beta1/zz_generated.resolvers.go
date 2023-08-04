/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/lambda/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AccessPoint.
func (mg *AccessPoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &v1beta1.BucketList{},
			Managed: &v1beta1.Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCConfiguration[i3].VPCID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.VPCConfiguration[i3].VPCIDRef,
			Selector:     mg.Spec.ForProvider.VPCConfiguration[i3].VPCIDSelector,
			To: reference.To{
				List:    &v1beta11.VPCList{},
				Managed: &v1beta11.VPC{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfiguration[i3].VPCID")
		}
		mg.Spec.ForProvider.VPCConfiguration[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPCConfiguration[i3].VPCIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this AccessPointPolicy.
func (mg *AccessPointPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessPointArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.AccessPointArnRef,
		Selector:     mg.Spec.ForProvider.AccessPointArnSelector,
		To: reference.To{
			List:    &AccessPointList{},
			Managed: &AccessPoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccessPointArn")
	}
	mg.Spec.ForProvider.AccessPointArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessPointArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Details); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Details[i3].Region); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Details[i3].Region[i4].Bucket),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Details[i3].Region[i4].BucketRef,
				Selector:     mg.Spec.ForProvider.Details[i3].Region[i4].BucketSelector,
				To: reference.To{
					List:    &v1beta1.BucketList{},
					Managed: &v1beta1.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Details[i3].Region[i4].Bucket")
			}
			mg.Spec.ForProvider.Details[i3].Region[i4].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Details[i3].Region[i4].BucketRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].SupportingAccessPoint),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Configuration[i3].SupportingAccessPointRef,
			Selector:     mg.Spec.ForProvider.Configuration[i3].SupportingAccessPointSelector,
			To: reference.To{
				List:    &AccessPointList{},
				Managed: &AccessPoint{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].SupportingAccessPoint")
		}
		mg.Spec.ForProvider.Configuration[i3].SupportingAccessPoint = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Configuration[i3].SupportingAccessPointRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArn),
						Extract:      common.ARNExtractor(),
						Reference:    mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArnRef,
						Selector:     mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArnSelector,
						To: reference.To{
							List:    &v1beta12.FunctionList{},
							Managed: &v1beta12.Function{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArn")
					}
					mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArnRef = rsp.ResolvedReference

				}
			}
		}
	}

	return nil
}

// ResolveReferences of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Name),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.NameRef,
		Selector:     mg.Spec.ForProvider.NameSelector,
		To: reference.To{
			List:    &ObjectLambdaAccessPointList{},
			Managed: &ObjectLambdaAccessPoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Name")
	}
	mg.Spec.ForProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLensConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].Arn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].ArnRef,
					Selector:     mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].ArnSelector,
					To: reference.To{
						List:    &v1beta1.BucketList{},
						Managed: &v1beta1.Bucket{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].Arn")
				}
				mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}
