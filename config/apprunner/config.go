/*
Copyright 2022 Upbound Inc.
*/

package apprunner

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for apprunner group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_apprunner_vpc_connector", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"subnets": {
				Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"security_groups": {
				Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
		}
		r.UseAsync = true
	})
}