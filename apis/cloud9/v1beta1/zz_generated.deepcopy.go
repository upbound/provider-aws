//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2) DeepCopyInto(out *EnvironmentEC2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2.
func (in *EnvironmentEC2) DeepCopy() *EnvironmentEC2 {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentEC2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2List) DeepCopyInto(out *EnvironmentEC2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvironmentEC2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2List.
func (in *EnvironmentEC2List) DeepCopy() *EnvironmentEC2List {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentEC2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2Observation) DeepCopyInto(out *EnvironmentEC2Observation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutomaticStopTimeMinutes != nil {
		in, out := &in.AutomaticStopTimeMinutes, &out.AutomaticStopTimeMinutes
		*out = new(float64)
		**out = **in
	}
	if in.ConnectionType != nil {
		in, out := &in.ConnectionType, &out.ConnectionType
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerArn != nil {
		in, out := &in.OwnerArn, &out.OwnerArn
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2Observation.
func (in *EnvironmentEC2Observation) DeepCopy() *EnvironmentEC2Observation {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2Parameters) DeepCopyInto(out *EnvironmentEC2Parameters) {
	*out = *in
	if in.AutomaticStopTimeMinutes != nil {
		in, out := &in.AutomaticStopTimeMinutes, &out.AutomaticStopTimeMinutes
		*out = new(float64)
		**out = **in
	}
	if in.ConnectionType != nil {
		in, out := &in.ConnectionType, &out.ConnectionType
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerArn != nil {
		in, out := &in.OwnerArn, &out.OwnerArn
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2Parameters.
func (in *EnvironmentEC2Parameters) DeepCopy() *EnvironmentEC2Parameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2Spec) DeepCopyInto(out *EnvironmentEC2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2Spec.
func (in *EnvironmentEC2Spec) DeepCopy() *EnvironmentEC2Spec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2Status) DeepCopyInto(out *EnvironmentEC2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2Status.
func (in *EnvironmentEC2Status) DeepCopy() *EnvironmentEC2Status {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentMembership) DeepCopyInto(out *EnvironmentMembership) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentMembership.
func (in *EnvironmentMembership) DeepCopy() *EnvironmentMembership {
	if in == nil {
		return nil
	}
	out := new(EnvironmentMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentMembership) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentMembershipList) DeepCopyInto(out *EnvironmentMembershipList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvironmentMembership, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentMembershipList.
func (in *EnvironmentMembershipList) DeepCopy() *EnvironmentMembershipList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentMembershipList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentMembershipList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentMembershipObservation) DeepCopyInto(out *EnvironmentMembershipObservation) {
	*out = *in
	if in.EnvironmentID != nil {
		in, out := &in.EnvironmentID, &out.EnvironmentID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.UserArn != nil {
		in, out := &in.UserArn, &out.UserArn
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentMembershipObservation.
func (in *EnvironmentMembershipObservation) DeepCopy() *EnvironmentMembershipObservation {
	if in == nil {
		return nil
	}
	out := new(EnvironmentMembershipObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentMembershipParameters) DeepCopyInto(out *EnvironmentMembershipParameters) {
	*out = *in
	if in.EnvironmentID != nil {
		in, out := &in.EnvironmentID, &out.EnvironmentID
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentIDRef != nil {
		in, out := &in.EnvironmentIDRef, &out.EnvironmentIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvironmentIDSelector != nil {
		in, out := &in.EnvironmentIDSelector, &out.EnvironmentIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.UserArn != nil {
		in, out := &in.UserArn, &out.UserArn
		*out = new(string)
		**out = **in
	}
	if in.UserArnRef != nil {
		in, out := &in.UserArnRef, &out.UserArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.UserArnSelector != nil {
		in, out := &in.UserArnSelector, &out.UserArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentMembershipParameters.
func (in *EnvironmentMembershipParameters) DeepCopy() *EnvironmentMembershipParameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentMembershipParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentMembershipSpec) DeepCopyInto(out *EnvironmentMembershipSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentMembershipSpec.
func (in *EnvironmentMembershipSpec) DeepCopy() *EnvironmentMembershipSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentMembershipSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentMembershipStatus) DeepCopyInto(out *EnvironmentMembershipStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentMembershipStatus.
func (in *EnvironmentMembershipStatus) DeepCopy() *EnvironmentMembershipStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentMembershipStatus)
	in.DeepCopyInto(out)
	return out
}
