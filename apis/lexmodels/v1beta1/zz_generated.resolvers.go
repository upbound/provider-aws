/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BotAlias.
func (mg *BotAlias) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ConversationLogs); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConversationLogs[i3].IAMRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ConversationLogs[i3].IAMRoleArnRef,
			Selector:     mg.Spec.ForProvider.ConversationLogs[i3].IAMRoleArnSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ConversationLogs[i3].IAMRoleArn")
		}
		mg.Spec.ForProvider.ConversationLogs[i3].IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ConversationLogs[i3].IAMRoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ConversationLogs); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ConversationLogs[i3].LogSettings); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConversationLogs[i3].LogSettings[i4].KMSKeyArn),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.ConversationLogs[i3].LogSettings[i4].KMSKeyArnRef,
				Selector:     mg.Spec.ForProvider.ConversationLogs[i3].LogSettings[i4].KMSKeyArnSelector,
				To: reference.To{
					List:    &v1beta11.KeyList{},
					Managed: &v1beta11.Key{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ConversationLogs[i3].LogSettings[i4].KMSKeyArn")
			}
			mg.Spec.ForProvider.ConversationLogs[i3].LogSettings[i4].KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ConversationLogs[i3].LogSettings[i4].KMSKeyArnRef = rsp.ResolvedReference

		}
	}

	return nil
}
