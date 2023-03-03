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
func (in *Accelerator) DeepCopyInto(out *Accelerator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Accelerator.
func (in *Accelerator) DeepCopy() *Accelerator {
	if in == nil {
		return nil
	}
	out := new(Accelerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Accelerator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorList) DeepCopyInto(out *AcceleratorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Accelerator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorList.
func (in *AcceleratorList) DeepCopy() *AcceleratorList {
	if in == nil {
		return nil
	}
	out := new(AcceleratorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcceleratorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorObservation) DeepCopyInto(out *AcceleratorObservation) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]AttributesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DNSName != nil {
		in, out := &in.DNSName, &out.DNSName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.HostedZoneID != nil {
		in, out := &in.HostedZoneID, &out.HostedZoneID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPAddressType != nil {
		in, out := &in.IPAddressType, &out.IPAddressType
		*out = new(string)
		**out = **in
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IPSets != nil {
		in, out := &in.IPSets, &out.IPSets
		*out = make([]IPSetsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorObservation.
func (in *AcceleratorObservation) DeepCopy() *AcceleratorObservation {
	if in == nil {
		return nil
	}
	out := new(AcceleratorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorParameters) DeepCopyInto(out *AcceleratorParameters) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]AttributesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IPAddressType != nil {
		in, out := &in.IPAddressType, &out.IPAddressType
		*out = new(string)
		**out = **in
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorParameters.
func (in *AcceleratorParameters) DeepCopy() *AcceleratorParameters {
	if in == nil {
		return nil
	}
	out := new(AcceleratorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorSpec) DeepCopyInto(out *AcceleratorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorSpec.
func (in *AcceleratorSpec) DeepCopy() *AcceleratorSpec {
	if in == nil {
		return nil
	}
	out := new(AcceleratorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorStatus) DeepCopyInto(out *AcceleratorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorStatus.
func (in *AcceleratorStatus) DeepCopy() *AcceleratorStatus {
	if in == nil {
		return nil
	}
	out := new(AcceleratorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttributesObservation) DeepCopyInto(out *AttributesObservation) {
	*out = *in
	if in.FlowLogsEnabled != nil {
		in, out := &in.FlowLogsEnabled, &out.FlowLogsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.FlowLogsS3Bucket != nil {
		in, out := &in.FlowLogsS3Bucket, &out.FlowLogsS3Bucket
		*out = new(string)
		**out = **in
	}
	if in.FlowLogsS3Prefix != nil {
		in, out := &in.FlowLogsS3Prefix, &out.FlowLogsS3Prefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttributesObservation.
func (in *AttributesObservation) DeepCopy() *AttributesObservation {
	if in == nil {
		return nil
	}
	out := new(AttributesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttributesParameters) DeepCopyInto(out *AttributesParameters) {
	*out = *in
	if in.FlowLogsEnabled != nil {
		in, out := &in.FlowLogsEnabled, &out.FlowLogsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.FlowLogsS3Bucket != nil {
		in, out := &in.FlowLogsS3Bucket, &out.FlowLogsS3Bucket
		*out = new(string)
		**out = **in
	}
	if in.FlowLogsS3Prefix != nil {
		in, out := &in.FlowLogsS3Prefix, &out.FlowLogsS3Prefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttributesParameters.
func (in *AttributesParameters) DeepCopy() *AttributesParameters {
	if in == nil {
		return nil
	}
	out := new(AttributesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointConfigurationObservation) DeepCopyInto(out *EndpointConfigurationObservation) {
	*out = *in
	if in.ClientIPPreservationEnabled != nil {
		in, out := &in.ClientIPPreservationEnabled, &out.ClientIPPreservationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EndpointID != nil {
		in, out := &in.EndpointID, &out.EndpointID
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointConfigurationObservation.
func (in *EndpointConfigurationObservation) DeepCopy() *EndpointConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(EndpointConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointConfigurationParameters) DeepCopyInto(out *EndpointConfigurationParameters) {
	*out = *in
	if in.ClientIPPreservationEnabled != nil {
		in, out := &in.ClientIPPreservationEnabled, &out.ClientIPPreservationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EndpointID != nil {
		in, out := &in.EndpointID, &out.EndpointID
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointConfigurationParameters.
func (in *EndpointConfigurationParameters) DeepCopy() *EndpointConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(EndpointConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroup) DeepCopyInto(out *EndpointGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroup.
func (in *EndpointGroup) DeepCopy() *EndpointGroup {
	if in == nil {
		return nil
	}
	out := new(EndpointGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EndpointGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupList) DeepCopyInto(out *EndpointGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EndpointGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupList.
func (in *EndpointGroupList) DeepCopy() *EndpointGroupList {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EndpointGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupObservation) DeepCopyInto(out *EndpointGroupObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.EndpointConfiguration != nil {
		in, out := &in.EndpointConfiguration, &out.EndpointConfiguration
		*out = make([]EndpointConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndpointGroupRegion != nil {
		in, out := &in.EndpointGroupRegion, &out.EndpointGroupRegion
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckIntervalSeconds != nil {
		in, out := &in.HealthCheckIntervalSeconds, &out.HealthCheckIntervalSeconds
		*out = new(float64)
		**out = **in
	}
	if in.HealthCheckPath != nil {
		in, out := &in.HealthCheckPath, &out.HealthCheckPath
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckPort != nil {
		in, out := &in.HealthCheckPort, &out.HealthCheckPort
		*out = new(float64)
		**out = **in
	}
	if in.HealthCheckProtocol != nil {
		in, out := &in.HealthCheckProtocol, &out.HealthCheckProtocol
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ListenerArn != nil {
		in, out := &in.ListenerArn, &out.ListenerArn
		*out = new(string)
		**out = **in
	}
	if in.PortOverride != nil {
		in, out := &in.PortOverride, &out.PortOverride
		*out = make([]PortOverrideObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ThresholdCount != nil {
		in, out := &in.ThresholdCount, &out.ThresholdCount
		*out = new(float64)
		**out = **in
	}
	if in.TrafficDialPercentage != nil {
		in, out := &in.TrafficDialPercentage, &out.TrafficDialPercentage
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupObservation.
func (in *EndpointGroupObservation) DeepCopy() *EndpointGroupObservation {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupParameters) DeepCopyInto(out *EndpointGroupParameters) {
	*out = *in
	if in.EndpointConfiguration != nil {
		in, out := &in.EndpointConfiguration, &out.EndpointConfiguration
		*out = make([]EndpointConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndpointGroupRegion != nil {
		in, out := &in.EndpointGroupRegion, &out.EndpointGroupRegion
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckIntervalSeconds != nil {
		in, out := &in.HealthCheckIntervalSeconds, &out.HealthCheckIntervalSeconds
		*out = new(float64)
		**out = **in
	}
	if in.HealthCheckPath != nil {
		in, out := &in.HealthCheckPath, &out.HealthCheckPath
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckPort != nil {
		in, out := &in.HealthCheckPort, &out.HealthCheckPort
		*out = new(float64)
		**out = **in
	}
	if in.HealthCheckProtocol != nil {
		in, out := &in.HealthCheckProtocol, &out.HealthCheckProtocol
		*out = new(string)
		**out = **in
	}
	if in.ListenerArn != nil {
		in, out := &in.ListenerArn, &out.ListenerArn
		*out = new(string)
		**out = **in
	}
	if in.ListenerArnRef != nil {
		in, out := &in.ListenerArnRef, &out.ListenerArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ListenerArnSelector != nil {
		in, out := &in.ListenerArnSelector, &out.ListenerArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PortOverride != nil {
		in, out := &in.PortOverride, &out.PortOverride
		*out = make([]PortOverrideParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ThresholdCount != nil {
		in, out := &in.ThresholdCount, &out.ThresholdCount
		*out = new(float64)
		**out = **in
	}
	if in.TrafficDialPercentage != nil {
		in, out := &in.TrafficDialPercentage, &out.TrafficDialPercentage
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupParameters.
func (in *EndpointGroupParameters) DeepCopy() *EndpointGroupParameters {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupSpec) DeepCopyInto(out *EndpointGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupSpec.
func (in *EndpointGroupSpec) DeepCopy() *EndpointGroupSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupStatus) DeepCopyInto(out *EndpointGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupStatus.
func (in *EndpointGroupStatus) DeepCopy() *EndpointGroupStatus {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPSetsObservation) DeepCopyInto(out *IPSetsObservation) {
	*out = *in
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IPFamily != nil {
		in, out := &in.IPFamily, &out.IPFamily
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPSetsObservation.
func (in *IPSetsObservation) DeepCopy() *IPSetsObservation {
	if in == nil {
		return nil
	}
	out := new(IPSetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPSetsParameters) DeepCopyInto(out *IPSetsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPSetsParameters.
func (in *IPSetsParameters) DeepCopy() *IPSetsParameters {
	if in == nil {
		return nil
	}
	out := new(IPSetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listener) DeepCopyInto(out *Listener) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listener.
func (in *Listener) DeepCopy() *Listener {
	if in == nil {
		return nil
	}
	out := new(Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Listener) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerList) DeepCopyInto(out *ListenerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Listener, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerList.
func (in *ListenerList) DeepCopy() *ListenerList {
	if in == nil {
		return nil
	}
	out := new(ListenerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ListenerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerObservation) DeepCopyInto(out *ListenerObservation) {
	*out = *in
	if in.AcceleratorArn != nil {
		in, out := &in.AcceleratorArn, &out.AcceleratorArn
		*out = new(string)
		**out = **in
	}
	if in.ClientAffinity != nil {
		in, out := &in.ClientAffinity, &out.ClientAffinity
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PortRange != nil {
		in, out := &in.PortRange, &out.PortRange
		*out = make([]PortRangeObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerObservation.
func (in *ListenerObservation) DeepCopy() *ListenerObservation {
	if in == nil {
		return nil
	}
	out := new(ListenerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerParameters) DeepCopyInto(out *ListenerParameters) {
	*out = *in
	if in.AcceleratorArn != nil {
		in, out := &in.AcceleratorArn, &out.AcceleratorArn
		*out = new(string)
		**out = **in
	}
	if in.AcceleratorArnRef != nil {
		in, out := &in.AcceleratorArnRef, &out.AcceleratorArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AcceleratorArnSelector != nil {
		in, out := &in.AcceleratorArnSelector, &out.AcceleratorArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientAffinity != nil {
		in, out := &in.ClientAffinity, &out.ClientAffinity
		*out = new(string)
		**out = **in
	}
	if in.PortRange != nil {
		in, out := &in.PortRange, &out.PortRange
		*out = make([]PortRangeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerParameters.
func (in *ListenerParameters) DeepCopy() *ListenerParameters {
	if in == nil {
		return nil
	}
	out := new(ListenerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerSpec) DeepCopyInto(out *ListenerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerSpec.
func (in *ListenerSpec) DeepCopy() *ListenerSpec {
	if in == nil {
		return nil
	}
	out := new(ListenerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerStatus) DeepCopyInto(out *ListenerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerStatus.
func (in *ListenerStatus) DeepCopy() *ListenerStatus {
	if in == nil {
		return nil
	}
	out := new(ListenerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortOverrideObservation) DeepCopyInto(out *PortOverrideObservation) {
	*out = *in
	if in.EndpointPort != nil {
		in, out := &in.EndpointPort, &out.EndpointPort
		*out = new(float64)
		**out = **in
	}
	if in.ListenerPort != nil {
		in, out := &in.ListenerPort, &out.ListenerPort
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortOverrideObservation.
func (in *PortOverrideObservation) DeepCopy() *PortOverrideObservation {
	if in == nil {
		return nil
	}
	out := new(PortOverrideObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortOverrideParameters) DeepCopyInto(out *PortOverrideParameters) {
	*out = *in
	if in.EndpointPort != nil {
		in, out := &in.EndpointPort, &out.EndpointPort
		*out = new(float64)
		**out = **in
	}
	if in.ListenerPort != nil {
		in, out := &in.ListenerPort, &out.ListenerPort
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortOverrideParameters.
func (in *PortOverrideParameters) DeepCopy() *PortOverrideParameters {
	if in == nil {
		return nil
	}
	out := new(PortOverrideParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRangeObservation) DeepCopyInto(out *PortRangeObservation) {
	*out = *in
	if in.FromPort != nil {
		in, out := &in.FromPort, &out.FromPort
		*out = new(float64)
		**out = **in
	}
	if in.ToPort != nil {
		in, out := &in.ToPort, &out.ToPort
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRangeObservation.
func (in *PortRangeObservation) DeepCopy() *PortRangeObservation {
	if in == nil {
		return nil
	}
	out := new(PortRangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRangeParameters) DeepCopyInto(out *PortRangeParameters) {
	*out = *in
	if in.FromPort != nil {
		in, out := &in.FromPort, &out.FromPort
		*out = new(float64)
		**out = **in
	}
	if in.ToPort != nil {
		in, out := &in.ToPort, &out.ToPort
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRangeParameters.
func (in *PortRangeParameters) DeepCopy() *PortRangeParameters {
	if in == nil {
		return nil
	}
	out := new(PortRangeParameters)
	in.DeepCopyInto(out)
	return out
}
