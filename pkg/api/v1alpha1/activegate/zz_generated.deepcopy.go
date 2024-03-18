//go:build !ignore_autogenerated

/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package activegate

import (
	"github.com/Dynatrace/dynatrace-operator/pkg/api/v1beta1/dynakube"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveGate) DeepCopyInto(out *ActiveGate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveGate.
func (in *ActiveGate) DeepCopy() *ActiveGate {
	if in == nil {
		return nil
	}

	out := new(ActiveGate)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveGate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveGateList) DeepCopyInto(out *ActiveGateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)

	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveGate, len(*in))

		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveGateList.
func (in *ActiveGateList) DeepCopy() *ActiveGateList {
	if in == nil {
		return nil
	}

	out := new(ActiveGateList)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveGateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveGateStatus) DeepCopyInto(out *ActiveGateStatus) {
	*out = *in
	in.UpdatedTimestamp.DeepCopyInto(&out.UpdatedTimestamp)
	in.VersionStatus.DeepCopyInto(&out.VersionStatus)
	in.ConnectionInfoStatus.DeepCopyInto(&out.ConnectionInfoStatus)

	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))

		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveGateStatus.
func (in *ActiveGateStatus) DeepCopy() *ActiveGateStatus {
	if in == nil {
		return nil
	}

	out := new(ActiveGateStatus)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Capability) DeepCopyInto(out *Capability) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Capability.
func (in *Capability) DeepCopy() *Capability {
	if in == nil {
		return nil
	}

	out := new(Capability)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapabilityProperties) DeepCopyInto(out *CapabilityProperties) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}

	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = new(ValueSource)
		**out = **in
	}

	in.Resources.DeepCopyInto(&out.Resources)

	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))

		for key, val := range *in {
			(*out)[key] = val
		}
	}

	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))

		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}

	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))

		for key, val := range *in {
			(*out)[key] = val
		}
	}

	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))

		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}

	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]corev1.TopologySpreadConstraint, len(*in))

		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapabilityProperties.
func (in *CapabilityProperties) DeepCopy() *CapabilityProperties {
	if in == nil {
		return nil
	}

	out := new(CapabilityProperties)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSpec) DeepCopyInto(out *CommonSpec) {
	*out = *in
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(dynakube.DynaKubeProxy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSpec.
func (in *CommonSpec) DeepCopy() *CommonSpec {
	if in == nil {
		return nil
	}

	out := new(CommonSpec)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionInfoStatus) DeepCopyInto(out *ConnectionInfoStatus) {
	*out = *in
	in.LastRequest.DeepCopyInto(&out.LastRequest)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionInfoStatus.
func (in *ConnectionInfoStatus) DeepCopy() *ConnectionInfoStatus {
	if in == nil {
		return nil
	}

	out := new(ConnectionInfoStatus)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecBasis) DeepCopyInto(out *SpecBasis) {
	*out = *in
	in.Common.DeepCopyInto(&out.Common)
	in.Specific.DeepCopyInto(&out.Specific)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecBasis.
func (in *SpecBasis) DeepCopy() *SpecBasis {
	if in == nil {
		return nil
	}

	out := new(SpecBasis)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecificSpec) DeepCopyInto(out *SpecificSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))

		for key, val := range *in {
			(*out)[key] = val
		}
	}

	in.CapabilityProperties.DeepCopyInto(&out.CapabilityProperties)

	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]CapabilityDisplayName, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecificSpec.
func (in *SpecificSpec) DeepCopy() *SpecificSpec {
	if in == nil {
		return nil
	}

	out := new(SpecificSpec)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueSource) DeepCopyInto(out *ValueSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueSource.
func (in *ValueSource) DeepCopy() *ValueSource {
	if in == nil {
		return nil
	}

	out := new(ValueSource)
	in.DeepCopyInto(out)

	return out
}
