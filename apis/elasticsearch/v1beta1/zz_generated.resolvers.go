/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1"
	v1beta13 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Domain.
func (mg *Domain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CognitoOptions); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CognitoOptions[i3].RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.CognitoOptions[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.CognitoOptions[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CognitoOptions[i3].RoleArn")
		}
		mg.Spec.ForProvider.CognitoOptions[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CognitoOptions[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptAtRest); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptAtRest[i3].KMSKeyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EncryptAtRest[i3].KMSKeyIDRef,
			Selector:     mg.Spec.ForProvider.EncryptAtRest[i3].KMSKeyIDSelector,
			To: reference.To{
				List:    &v1beta11.KeyList{},
				Managed: &v1beta11.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EncryptAtRest[i3].KMSKeyID")
		}
		mg.Spec.ForProvider.EncryptAtRest[i3].KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EncryptAtRest[i3].KMSKeyIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LogPublishingOptions); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnRef,
			Selector:     mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnSelector,
			To: reference.To{
				List:    &v1beta12.GroupList{},
				Managed: &v1beta12.Group{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn")
		}
		mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCOptions); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCOptions[i3].SecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.VPCOptions[i3].SecurityGroupIDRefs,
			Selector:      mg.Spec.ForProvider.VPCOptions[i3].SecurityGroupIDSelector,
			To: reference.To{
				List:    &v1beta13.SecurityGroupList{},
				Managed: &v1beta13.SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCOptions[i3].SecurityGroupIds")
		}
		mg.Spec.ForProvider.VPCOptions[i3].SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCOptions[i3].SecurityGroupIDRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCOptions); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCOptions[i3].SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.VPCOptions[i3].SubnetIDRefs,
			Selector:      mg.Spec.ForProvider.VPCOptions[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta13.SubnetList{},
				Managed: &v1beta13.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCOptions[i3].SubnetIds")
		}
		mg.Spec.ForProvider.VPCOptions[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCOptions[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this DomainPolicy.
func (mg *DomainPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DomainNameRef,
		Selector:     mg.Spec.ForProvider.DomainNameSelector,
		To: reference.To{
			List:    &DomainList{},
			Managed: &Domain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DomainName")
	}
	mg.Spec.ForProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainNameRef = rsp.ResolvedReference

	return nil
}
