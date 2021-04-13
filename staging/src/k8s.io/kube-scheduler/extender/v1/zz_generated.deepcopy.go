// +build !ignore_autogenerated

/*
Copyright 2020 Authors of Arktos.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderArgs) DeepCopyInto(out *ExtenderArgs) {
	*out = *in
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(corev1.Pod)
		(*in).DeepCopyInto(*out)
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = new(corev1.NodeList)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderArgs.
func (in *ExtenderArgs) DeepCopy() *ExtenderArgs {
	if in == nil {
		return nil
	}
	out := new(ExtenderArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderBindingArgs) DeepCopyInto(out *ExtenderBindingArgs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderBindingArgs.
func (in *ExtenderBindingArgs) DeepCopy() *ExtenderBindingArgs {
	if in == nil {
		return nil
	}
	out := new(ExtenderBindingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderBindingResult) DeepCopyInto(out *ExtenderBindingResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderBindingResult.
func (in *ExtenderBindingResult) DeepCopy() *ExtenderBindingResult {
	if in == nil {
		return nil
	}
	out := new(ExtenderBindingResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderFilterResult) DeepCopyInto(out *ExtenderFilterResult) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = new(corev1.NodeList)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.FailedNodes != nil {
		in, out := &in.FailedNodes, &out.FailedNodes
		*out = make(FailedNodesMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderFilterResult.
func (in *ExtenderFilterResult) DeepCopy() *ExtenderFilterResult {
	if in == nil {
		return nil
	}
	out := new(ExtenderFilterResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderPreemptionArgs) DeepCopyInto(out *ExtenderPreemptionArgs) {
	*out = *in
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(corev1.Pod)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeNameToVictims != nil {
		in, out := &in.NodeNameToVictims, &out.NodeNameToVictims
		*out = make(map[string]*Victims, len(*in))
		for key, val := range *in {
			var outVal *Victims
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Victims)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.NodeNameToMetaVictims != nil {
		in, out := &in.NodeNameToMetaVictims, &out.NodeNameToMetaVictims
		*out = make(map[string]*MetaVictims, len(*in))
		for key, val := range *in {
			var outVal *MetaVictims
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(MetaVictims)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderPreemptionArgs.
func (in *ExtenderPreemptionArgs) DeepCopy() *ExtenderPreemptionArgs {
	if in == nil {
		return nil
	}
	out := new(ExtenderPreemptionArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderPreemptionResult) DeepCopyInto(out *ExtenderPreemptionResult) {
	*out = *in
	if in.NodeNameToMetaVictims != nil {
		in, out := &in.NodeNameToMetaVictims, &out.NodeNameToMetaVictims
		*out = make(map[string]*MetaVictims, len(*in))
		for key, val := range *in {
			var outVal *MetaVictims
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(MetaVictims)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderPreemptionResult.
func (in *ExtenderPreemptionResult) DeepCopy() *ExtenderPreemptionResult {
	if in == nil {
		return nil
	}
	out := new(ExtenderPreemptionResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FailedNodesMap) DeepCopyInto(out *FailedNodesMap) {
	{
		in := &in
		*out = make(FailedNodesMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailedNodesMap.
func (in FailedNodesMap) DeepCopy() FailedNodesMap {
	if in == nil {
		return nil
	}
	out := new(FailedNodesMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPriority) DeepCopyInto(out *HostPriority) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPriority.
func (in *HostPriority) DeepCopy() *HostPriority {
	if in == nil {
		return nil
	}
	out := new(HostPriority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in HostPriorityList) DeepCopyInto(out *HostPriorityList) {
	{
		in := &in
		*out = make(HostPriorityList, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPriorityList.
func (in HostPriorityList) DeepCopy() HostPriorityList {
	if in == nil {
		return nil
	}
	out := new(HostPriorityList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaPod) DeepCopyInto(out *MetaPod) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaPod.
func (in *MetaPod) DeepCopy() *MetaPod {
	if in == nil {
		return nil
	}
	out := new(MetaPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaVictims) DeepCopyInto(out *MetaVictims) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]*MetaPod, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MetaPod)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaVictims.
func (in *MetaVictims) DeepCopy() *MetaVictims {
	if in == nil {
		return nil
	}
	out := new(MetaVictims)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Victims) DeepCopyInto(out *Victims) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]*corev1.Pod, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.Pod)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Victims.
func (in *Victims) DeepCopy() *Victims {
	if in == nil {
		return nil
	}
	out := new(Victims)
	in.DeepCopyInto(out)
	return out
}
