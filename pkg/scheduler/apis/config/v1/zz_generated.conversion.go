// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	time "time"
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kube-scheduler/config/v1"
	config "k8s.io/kubernetes/pkg/scheduler/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1.Extender)(nil), (*config.Extender)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Extender_To_config_Extender(a.(*v1.Extender), b.(*config.Extender), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Extender)(nil), (*v1.Extender)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Extender_To_v1_Extender(a.(*config.Extender), b.(*v1.Extender), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ExtenderManagedResource)(nil), (*config.ExtenderManagedResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ExtenderManagedResource_To_config_ExtenderManagedResource(a.(*v1.ExtenderManagedResource), b.(*config.ExtenderManagedResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ExtenderManagedResource)(nil), (*v1.ExtenderManagedResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ExtenderManagedResource_To_v1_ExtenderManagedResource(a.(*config.ExtenderManagedResource), b.(*v1.ExtenderManagedResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ExtenderTLSConfig)(nil), (*config.ExtenderTLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ExtenderTLSConfig_To_config_ExtenderTLSConfig(a.(*v1.ExtenderTLSConfig), b.(*config.ExtenderTLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ExtenderTLSConfig)(nil), (*v1.ExtenderTLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ExtenderTLSConfig_To_v1_ExtenderTLSConfig(a.(*config.ExtenderTLSConfig), b.(*v1.ExtenderTLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LabelPreference)(nil), (*config.LabelPreference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LabelPreference_To_config_LabelPreference(a.(*v1.LabelPreference), b.(*config.LabelPreference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LabelPreference)(nil), (*v1.LabelPreference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LabelPreference_To_v1_LabelPreference(a.(*config.LabelPreference), b.(*v1.LabelPreference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LabelsPresence)(nil), (*config.LabelsPresence)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LabelsPresence_To_config_LabelsPresence(a.(*v1.LabelsPresence), b.(*config.LabelsPresence), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LabelsPresence)(nil), (*v1.LabelsPresence)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LabelsPresence_To_v1_LabelsPresence(a.(*config.LabelsPresence), b.(*v1.LabelsPresence), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.Policy)(nil), (*config.Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Policy_To_config_Policy(a.(*v1.Policy), b.(*config.Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Policy)(nil), (*v1.Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Policy_To_v1_Policy(a.(*config.Policy), b.(*v1.Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PredicateArgument)(nil), (*config.PredicateArgument)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PredicateArgument_To_config_PredicateArgument(a.(*v1.PredicateArgument), b.(*config.PredicateArgument), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PredicateArgument)(nil), (*v1.PredicateArgument)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PredicateArgument_To_v1_PredicateArgument(a.(*config.PredicateArgument), b.(*v1.PredicateArgument), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PredicatePolicy)(nil), (*config.PredicatePolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PredicatePolicy_To_config_PredicatePolicy(a.(*v1.PredicatePolicy), b.(*config.PredicatePolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PredicatePolicy)(nil), (*v1.PredicatePolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PredicatePolicy_To_v1_PredicatePolicy(a.(*config.PredicatePolicy), b.(*v1.PredicatePolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PriorityArgument)(nil), (*config.PriorityArgument)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PriorityArgument_To_config_PriorityArgument(a.(*v1.PriorityArgument), b.(*config.PriorityArgument), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PriorityArgument)(nil), (*v1.PriorityArgument)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PriorityArgument_To_v1_PriorityArgument(a.(*config.PriorityArgument), b.(*v1.PriorityArgument), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PriorityPolicy)(nil), (*config.PriorityPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PriorityPolicy_To_config_PriorityPolicy(a.(*v1.PriorityPolicy), b.(*config.PriorityPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PriorityPolicy)(nil), (*v1.PriorityPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PriorityPolicy_To_v1_PriorityPolicy(a.(*config.PriorityPolicy), b.(*v1.PriorityPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RequestedToCapacityRatioArguments)(nil), (*config.RequestedToCapacityRatioArguments)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RequestedToCapacityRatioArguments_To_config_RequestedToCapacityRatioArguments(a.(*v1.RequestedToCapacityRatioArguments), b.(*config.RequestedToCapacityRatioArguments), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RequestedToCapacityRatioArguments)(nil), (*v1.RequestedToCapacityRatioArguments)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RequestedToCapacityRatioArguments_To_v1_RequestedToCapacityRatioArguments(a.(*config.RequestedToCapacityRatioArguments), b.(*v1.RequestedToCapacityRatioArguments), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ResourceSpec)(nil), (*config.ResourceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ResourceSpec_To_config_ResourceSpec(a.(*v1.ResourceSpec), b.(*config.ResourceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ResourceSpec)(nil), (*v1.ResourceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ResourceSpec_To_v1_ResourceSpec(a.(*config.ResourceSpec), b.(*v1.ResourceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceAffinity)(nil), (*config.ServiceAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceAffinity_To_config_ServiceAffinity(a.(*v1.ServiceAffinity), b.(*config.ServiceAffinity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ServiceAffinity)(nil), (*v1.ServiceAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServiceAffinity_To_v1_ServiceAffinity(a.(*config.ServiceAffinity), b.(*v1.ServiceAffinity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceAntiAffinity)(nil), (*config.ServiceAntiAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceAntiAffinity_To_config_ServiceAntiAffinity(a.(*v1.ServiceAntiAffinity), b.(*config.ServiceAntiAffinity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ServiceAntiAffinity)(nil), (*v1.ServiceAntiAffinity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServiceAntiAffinity_To_v1_ServiceAntiAffinity(a.(*config.ServiceAntiAffinity), b.(*v1.ServiceAntiAffinity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.UtilizationShapePoint)(nil), (*config.UtilizationShapePoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_UtilizationShapePoint_To_config_UtilizationShapePoint(a.(*v1.UtilizationShapePoint), b.(*config.UtilizationShapePoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.UtilizationShapePoint)(nil), (*v1.UtilizationShapePoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_UtilizationShapePoint_To_v1_UtilizationShapePoint(a.(*config.UtilizationShapePoint), b.(*v1.UtilizationShapePoint), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Extender_To_config_Extender(in *v1.Extender, out *config.Extender, s conversion.Scope) error {
	out.URLPrefix = in.URLPrefix
	out.FilterVerb = in.FilterVerb
	out.PreemptVerb = in.PreemptVerb
	out.PrioritizeVerb = in.PrioritizeVerb
	out.Weight = in.Weight
	out.BindVerb = in.BindVerb
	out.EnableHTTPS = in.EnableHTTPS
	out.TLSConfig = (*config.ExtenderTLSConfig)(unsafe.Pointer(in.TLSConfig))
	out.HTTPTimeout = time.Duration(in.HTTPTimeout)
	out.NodeCacheCapable = in.NodeCacheCapable
	out.ManagedResources = *(*[]config.ExtenderManagedResource)(unsafe.Pointer(&in.ManagedResources))
	out.Ignorable = in.Ignorable
	return nil
}

// Convert_v1_Extender_To_config_Extender is an autogenerated conversion function.
func Convert_v1_Extender_To_config_Extender(in *v1.Extender, out *config.Extender, s conversion.Scope) error {
	return autoConvert_v1_Extender_To_config_Extender(in, out, s)
}

func autoConvert_config_Extender_To_v1_Extender(in *config.Extender, out *v1.Extender, s conversion.Scope) error {
	out.URLPrefix = in.URLPrefix
	out.FilterVerb = in.FilterVerb
	out.PreemptVerb = in.PreemptVerb
	out.PrioritizeVerb = in.PrioritizeVerb
	out.Weight = in.Weight
	out.BindVerb = in.BindVerb
	out.EnableHTTPS = in.EnableHTTPS
	out.TLSConfig = (*v1.ExtenderTLSConfig)(unsafe.Pointer(in.TLSConfig))
	out.HTTPTimeout = time.Duration(in.HTTPTimeout)
	out.NodeCacheCapable = in.NodeCacheCapable
	out.ManagedResources = *(*[]v1.ExtenderManagedResource)(unsafe.Pointer(&in.ManagedResources))
	out.Ignorable = in.Ignorable
	return nil
}

// Convert_config_Extender_To_v1_Extender is an autogenerated conversion function.
func Convert_config_Extender_To_v1_Extender(in *config.Extender, out *v1.Extender, s conversion.Scope) error {
	return autoConvert_config_Extender_To_v1_Extender(in, out, s)
}

func autoConvert_v1_ExtenderManagedResource_To_config_ExtenderManagedResource(in *v1.ExtenderManagedResource, out *config.ExtenderManagedResource, s conversion.Scope) error {
	out.Name = in.Name
	out.IgnoredByScheduler = in.IgnoredByScheduler
	return nil
}

// Convert_v1_ExtenderManagedResource_To_config_ExtenderManagedResource is an autogenerated conversion function.
func Convert_v1_ExtenderManagedResource_To_config_ExtenderManagedResource(in *v1.ExtenderManagedResource, out *config.ExtenderManagedResource, s conversion.Scope) error {
	return autoConvert_v1_ExtenderManagedResource_To_config_ExtenderManagedResource(in, out, s)
}

func autoConvert_config_ExtenderManagedResource_To_v1_ExtenderManagedResource(in *config.ExtenderManagedResource, out *v1.ExtenderManagedResource, s conversion.Scope) error {
	out.Name = in.Name
	out.IgnoredByScheduler = in.IgnoredByScheduler
	return nil
}

// Convert_config_ExtenderManagedResource_To_v1_ExtenderManagedResource is an autogenerated conversion function.
func Convert_config_ExtenderManagedResource_To_v1_ExtenderManagedResource(in *config.ExtenderManagedResource, out *v1.ExtenderManagedResource, s conversion.Scope) error {
	return autoConvert_config_ExtenderManagedResource_To_v1_ExtenderManagedResource(in, out, s)
}

func autoConvert_v1_ExtenderTLSConfig_To_config_ExtenderTLSConfig(in *v1.ExtenderTLSConfig, out *config.ExtenderTLSConfig, s conversion.Scope) error {
	out.Insecure = in.Insecure
	out.ServerName = in.ServerName
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	out.CAFile = in.CAFile
	out.CertData = *(*[]byte)(unsafe.Pointer(&in.CertData))
	out.KeyData = *(*[]byte)(unsafe.Pointer(&in.KeyData))
	out.CAData = *(*[]byte)(unsafe.Pointer(&in.CAData))
	return nil
}

// Convert_v1_ExtenderTLSConfig_To_config_ExtenderTLSConfig is an autogenerated conversion function.
func Convert_v1_ExtenderTLSConfig_To_config_ExtenderTLSConfig(in *v1.ExtenderTLSConfig, out *config.ExtenderTLSConfig, s conversion.Scope) error {
	return autoConvert_v1_ExtenderTLSConfig_To_config_ExtenderTLSConfig(in, out, s)
}

func autoConvert_config_ExtenderTLSConfig_To_v1_ExtenderTLSConfig(in *config.ExtenderTLSConfig, out *v1.ExtenderTLSConfig, s conversion.Scope) error {
	out.Insecure = in.Insecure
	out.ServerName = in.ServerName
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	out.CAFile = in.CAFile
	out.CertData = *(*[]byte)(unsafe.Pointer(&in.CertData))
	out.KeyData = *(*[]byte)(unsafe.Pointer(&in.KeyData))
	out.CAData = *(*[]byte)(unsafe.Pointer(&in.CAData))
	return nil
}

// Convert_config_ExtenderTLSConfig_To_v1_ExtenderTLSConfig is an autogenerated conversion function.
func Convert_config_ExtenderTLSConfig_To_v1_ExtenderTLSConfig(in *config.ExtenderTLSConfig, out *v1.ExtenderTLSConfig, s conversion.Scope) error {
	return autoConvert_config_ExtenderTLSConfig_To_v1_ExtenderTLSConfig(in, out, s)
}

func autoConvert_v1_LabelPreference_To_config_LabelPreference(in *v1.LabelPreference, out *config.LabelPreference, s conversion.Scope) error {
	out.Label = in.Label
	out.Presence = in.Presence
	return nil
}

// Convert_v1_LabelPreference_To_config_LabelPreference is an autogenerated conversion function.
func Convert_v1_LabelPreference_To_config_LabelPreference(in *v1.LabelPreference, out *config.LabelPreference, s conversion.Scope) error {
	return autoConvert_v1_LabelPreference_To_config_LabelPreference(in, out, s)
}

func autoConvert_config_LabelPreference_To_v1_LabelPreference(in *config.LabelPreference, out *v1.LabelPreference, s conversion.Scope) error {
	out.Label = in.Label
	out.Presence = in.Presence
	return nil
}

// Convert_config_LabelPreference_To_v1_LabelPreference is an autogenerated conversion function.
func Convert_config_LabelPreference_To_v1_LabelPreference(in *config.LabelPreference, out *v1.LabelPreference, s conversion.Scope) error {
	return autoConvert_config_LabelPreference_To_v1_LabelPreference(in, out, s)
}

func autoConvert_v1_LabelsPresence_To_config_LabelsPresence(in *v1.LabelsPresence, out *config.LabelsPresence, s conversion.Scope) error {
	out.Labels = *(*[]string)(unsafe.Pointer(&in.Labels))
	out.Presence = in.Presence
	return nil
}

// Convert_v1_LabelsPresence_To_config_LabelsPresence is an autogenerated conversion function.
func Convert_v1_LabelsPresence_To_config_LabelsPresence(in *v1.LabelsPresence, out *config.LabelsPresence, s conversion.Scope) error {
	return autoConvert_v1_LabelsPresence_To_config_LabelsPresence(in, out, s)
}

func autoConvert_config_LabelsPresence_To_v1_LabelsPresence(in *config.LabelsPresence, out *v1.LabelsPresence, s conversion.Scope) error {
	out.Labels = *(*[]string)(unsafe.Pointer(&in.Labels))
	out.Presence = in.Presence
	return nil
}

// Convert_config_LabelsPresence_To_v1_LabelsPresence is an autogenerated conversion function.
func Convert_config_LabelsPresence_To_v1_LabelsPresence(in *config.LabelsPresence, out *v1.LabelsPresence, s conversion.Scope) error {
	return autoConvert_config_LabelsPresence_To_v1_LabelsPresence(in, out, s)
}

func autoConvert_v1_Policy_To_config_Policy(in *v1.Policy, out *config.Policy, s conversion.Scope) error {
	out.Predicates = *(*[]config.PredicatePolicy)(unsafe.Pointer(&in.Predicates))
	out.Priorities = *(*[]config.PriorityPolicy)(unsafe.Pointer(&in.Priorities))
	out.Extenders = *(*[]config.Extender)(unsafe.Pointer(&in.Extenders))
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.AlwaysCheckAllPredicates = in.AlwaysCheckAllPredicates
	return nil
}

// Convert_v1_Policy_To_config_Policy is an autogenerated conversion function.
func Convert_v1_Policy_To_config_Policy(in *v1.Policy, out *config.Policy, s conversion.Scope) error {
	return autoConvert_v1_Policy_To_config_Policy(in, out, s)
}

func autoConvert_config_Policy_To_v1_Policy(in *config.Policy, out *v1.Policy, s conversion.Scope) error {
	out.Predicates = *(*[]v1.PredicatePolicy)(unsafe.Pointer(&in.Predicates))
	out.Priorities = *(*[]v1.PriorityPolicy)(unsafe.Pointer(&in.Priorities))
	out.Extenders = *(*[]v1.Extender)(unsafe.Pointer(&in.Extenders))
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.AlwaysCheckAllPredicates = in.AlwaysCheckAllPredicates
	return nil
}

// Convert_config_Policy_To_v1_Policy is an autogenerated conversion function.
func Convert_config_Policy_To_v1_Policy(in *config.Policy, out *v1.Policy, s conversion.Scope) error {
	return autoConvert_config_Policy_To_v1_Policy(in, out, s)
}

func autoConvert_v1_PredicateArgument_To_config_PredicateArgument(in *v1.PredicateArgument, out *config.PredicateArgument, s conversion.Scope) error {
	out.ServiceAffinity = (*config.ServiceAffinity)(unsafe.Pointer(in.ServiceAffinity))
	out.LabelsPresence = (*config.LabelsPresence)(unsafe.Pointer(in.LabelsPresence))
	return nil
}

// Convert_v1_PredicateArgument_To_config_PredicateArgument is an autogenerated conversion function.
func Convert_v1_PredicateArgument_To_config_PredicateArgument(in *v1.PredicateArgument, out *config.PredicateArgument, s conversion.Scope) error {
	return autoConvert_v1_PredicateArgument_To_config_PredicateArgument(in, out, s)
}

func autoConvert_config_PredicateArgument_To_v1_PredicateArgument(in *config.PredicateArgument, out *v1.PredicateArgument, s conversion.Scope) error {
	out.ServiceAffinity = (*v1.ServiceAffinity)(unsafe.Pointer(in.ServiceAffinity))
	out.LabelsPresence = (*v1.LabelsPresence)(unsafe.Pointer(in.LabelsPresence))
	return nil
}

// Convert_config_PredicateArgument_To_v1_PredicateArgument is an autogenerated conversion function.
func Convert_config_PredicateArgument_To_v1_PredicateArgument(in *config.PredicateArgument, out *v1.PredicateArgument, s conversion.Scope) error {
	return autoConvert_config_PredicateArgument_To_v1_PredicateArgument(in, out, s)
}

func autoConvert_v1_PredicatePolicy_To_config_PredicatePolicy(in *v1.PredicatePolicy, out *config.PredicatePolicy, s conversion.Scope) error {
	out.Name = in.Name
	out.Argument = (*config.PredicateArgument)(unsafe.Pointer(in.Argument))
	return nil
}

// Convert_v1_PredicatePolicy_To_config_PredicatePolicy is an autogenerated conversion function.
func Convert_v1_PredicatePolicy_To_config_PredicatePolicy(in *v1.PredicatePolicy, out *config.PredicatePolicy, s conversion.Scope) error {
	return autoConvert_v1_PredicatePolicy_To_config_PredicatePolicy(in, out, s)
}

func autoConvert_config_PredicatePolicy_To_v1_PredicatePolicy(in *config.PredicatePolicy, out *v1.PredicatePolicy, s conversion.Scope) error {
	out.Name = in.Name
	out.Argument = (*v1.PredicateArgument)(unsafe.Pointer(in.Argument))
	return nil
}

// Convert_config_PredicatePolicy_To_v1_PredicatePolicy is an autogenerated conversion function.
func Convert_config_PredicatePolicy_To_v1_PredicatePolicy(in *config.PredicatePolicy, out *v1.PredicatePolicy, s conversion.Scope) error {
	return autoConvert_config_PredicatePolicy_To_v1_PredicatePolicy(in, out, s)
}

func autoConvert_v1_PriorityArgument_To_config_PriorityArgument(in *v1.PriorityArgument, out *config.PriorityArgument, s conversion.Scope) error {
	out.ServiceAntiAffinity = (*config.ServiceAntiAffinity)(unsafe.Pointer(in.ServiceAntiAffinity))
	out.LabelPreference = (*config.LabelPreference)(unsafe.Pointer(in.LabelPreference))
	out.RequestedToCapacityRatioArguments = (*config.RequestedToCapacityRatioArguments)(unsafe.Pointer(in.RequestedToCapacityRatioArguments))
	return nil
}

// Convert_v1_PriorityArgument_To_config_PriorityArgument is an autogenerated conversion function.
func Convert_v1_PriorityArgument_To_config_PriorityArgument(in *v1.PriorityArgument, out *config.PriorityArgument, s conversion.Scope) error {
	return autoConvert_v1_PriorityArgument_To_config_PriorityArgument(in, out, s)
}

func autoConvert_config_PriorityArgument_To_v1_PriorityArgument(in *config.PriorityArgument, out *v1.PriorityArgument, s conversion.Scope) error {
	out.ServiceAntiAffinity = (*v1.ServiceAntiAffinity)(unsafe.Pointer(in.ServiceAntiAffinity))
	out.LabelPreference = (*v1.LabelPreference)(unsafe.Pointer(in.LabelPreference))
	out.RequestedToCapacityRatioArguments = (*v1.RequestedToCapacityRatioArguments)(unsafe.Pointer(in.RequestedToCapacityRatioArguments))
	return nil
}

// Convert_config_PriorityArgument_To_v1_PriorityArgument is an autogenerated conversion function.
func Convert_config_PriorityArgument_To_v1_PriorityArgument(in *config.PriorityArgument, out *v1.PriorityArgument, s conversion.Scope) error {
	return autoConvert_config_PriorityArgument_To_v1_PriorityArgument(in, out, s)
}

func autoConvert_v1_PriorityPolicy_To_config_PriorityPolicy(in *v1.PriorityPolicy, out *config.PriorityPolicy, s conversion.Scope) error {
	out.Name = in.Name
	out.Weight = in.Weight
	out.Argument = (*config.PriorityArgument)(unsafe.Pointer(in.Argument))
	return nil
}

// Convert_v1_PriorityPolicy_To_config_PriorityPolicy is an autogenerated conversion function.
func Convert_v1_PriorityPolicy_To_config_PriorityPolicy(in *v1.PriorityPolicy, out *config.PriorityPolicy, s conversion.Scope) error {
	return autoConvert_v1_PriorityPolicy_To_config_PriorityPolicy(in, out, s)
}

func autoConvert_config_PriorityPolicy_To_v1_PriorityPolicy(in *config.PriorityPolicy, out *v1.PriorityPolicy, s conversion.Scope) error {
	out.Name = in.Name
	out.Weight = in.Weight
	out.Argument = (*v1.PriorityArgument)(unsafe.Pointer(in.Argument))
	return nil
}

// Convert_config_PriorityPolicy_To_v1_PriorityPolicy is an autogenerated conversion function.
func Convert_config_PriorityPolicy_To_v1_PriorityPolicy(in *config.PriorityPolicy, out *v1.PriorityPolicy, s conversion.Scope) error {
	return autoConvert_config_PriorityPolicy_To_v1_PriorityPolicy(in, out, s)
}

func autoConvert_v1_RequestedToCapacityRatioArguments_To_config_RequestedToCapacityRatioArguments(in *v1.RequestedToCapacityRatioArguments, out *config.RequestedToCapacityRatioArguments, s conversion.Scope) error {
	out.Shape = *(*[]config.UtilizationShapePoint)(unsafe.Pointer(&in.Shape))
	out.Resources = *(*[]config.ResourceSpec)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1_RequestedToCapacityRatioArguments_To_config_RequestedToCapacityRatioArguments is an autogenerated conversion function.
func Convert_v1_RequestedToCapacityRatioArguments_To_config_RequestedToCapacityRatioArguments(in *v1.RequestedToCapacityRatioArguments, out *config.RequestedToCapacityRatioArguments, s conversion.Scope) error {
	return autoConvert_v1_RequestedToCapacityRatioArguments_To_config_RequestedToCapacityRatioArguments(in, out, s)
}

func autoConvert_config_RequestedToCapacityRatioArguments_To_v1_RequestedToCapacityRatioArguments(in *config.RequestedToCapacityRatioArguments, out *v1.RequestedToCapacityRatioArguments, s conversion.Scope) error {
	out.Shape = *(*[]v1.UtilizationShapePoint)(unsafe.Pointer(&in.Shape))
	out.Resources = *(*[]v1.ResourceSpec)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_config_RequestedToCapacityRatioArguments_To_v1_RequestedToCapacityRatioArguments is an autogenerated conversion function.
func Convert_config_RequestedToCapacityRatioArguments_To_v1_RequestedToCapacityRatioArguments(in *config.RequestedToCapacityRatioArguments, out *v1.RequestedToCapacityRatioArguments, s conversion.Scope) error {
	return autoConvert_config_RequestedToCapacityRatioArguments_To_v1_RequestedToCapacityRatioArguments(in, out, s)
}

func autoConvert_v1_ResourceSpec_To_config_ResourceSpec(in *v1.ResourceSpec, out *config.ResourceSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.Weight = in.Weight
	return nil
}

// Convert_v1_ResourceSpec_To_config_ResourceSpec is an autogenerated conversion function.
func Convert_v1_ResourceSpec_To_config_ResourceSpec(in *v1.ResourceSpec, out *config.ResourceSpec, s conversion.Scope) error {
	return autoConvert_v1_ResourceSpec_To_config_ResourceSpec(in, out, s)
}

func autoConvert_config_ResourceSpec_To_v1_ResourceSpec(in *config.ResourceSpec, out *v1.ResourceSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.Weight = in.Weight
	return nil
}

// Convert_config_ResourceSpec_To_v1_ResourceSpec is an autogenerated conversion function.
func Convert_config_ResourceSpec_To_v1_ResourceSpec(in *config.ResourceSpec, out *v1.ResourceSpec, s conversion.Scope) error {
	return autoConvert_config_ResourceSpec_To_v1_ResourceSpec(in, out, s)
}

func autoConvert_v1_ServiceAffinity_To_config_ServiceAffinity(in *v1.ServiceAffinity, out *config.ServiceAffinity, s conversion.Scope) error {
	out.Labels = *(*[]string)(unsafe.Pointer(&in.Labels))
	return nil
}

// Convert_v1_ServiceAffinity_To_config_ServiceAffinity is an autogenerated conversion function.
func Convert_v1_ServiceAffinity_To_config_ServiceAffinity(in *v1.ServiceAffinity, out *config.ServiceAffinity, s conversion.Scope) error {
	return autoConvert_v1_ServiceAffinity_To_config_ServiceAffinity(in, out, s)
}

func autoConvert_config_ServiceAffinity_To_v1_ServiceAffinity(in *config.ServiceAffinity, out *v1.ServiceAffinity, s conversion.Scope) error {
	out.Labels = *(*[]string)(unsafe.Pointer(&in.Labels))
	return nil
}

// Convert_config_ServiceAffinity_To_v1_ServiceAffinity is an autogenerated conversion function.
func Convert_config_ServiceAffinity_To_v1_ServiceAffinity(in *config.ServiceAffinity, out *v1.ServiceAffinity, s conversion.Scope) error {
	return autoConvert_config_ServiceAffinity_To_v1_ServiceAffinity(in, out, s)
}

func autoConvert_v1_ServiceAntiAffinity_To_config_ServiceAntiAffinity(in *v1.ServiceAntiAffinity, out *config.ServiceAntiAffinity, s conversion.Scope) error {
	out.Label = in.Label
	return nil
}

// Convert_v1_ServiceAntiAffinity_To_config_ServiceAntiAffinity is an autogenerated conversion function.
func Convert_v1_ServiceAntiAffinity_To_config_ServiceAntiAffinity(in *v1.ServiceAntiAffinity, out *config.ServiceAntiAffinity, s conversion.Scope) error {
	return autoConvert_v1_ServiceAntiAffinity_To_config_ServiceAntiAffinity(in, out, s)
}

func autoConvert_config_ServiceAntiAffinity_To_v1_ServiceAntiAffinity(in *config.ServiceAntiAffinity, out *v1.ServiceAntiAffinity, s conversion.Scope) error {
	out.Label = in.Label
	return nil
}

// Convert_config_ServiceAntiAffinity_To_v1_ServiceAntiAffinity is an autogenerated conversion function.
func Convert_config_ServiceAntiAffinity_To_v1_ServiceAntiAffinity(in *config.ServiceAntiAffinity, out *v1.ServiceAntiAffinity, s conversion.Scope) error {
	return autoConvert_config_ServiceAntiAffinity_To_v1_ServiceAntiAffinity(in, out, s)
}

func autoConvert_v1_UtilizationShapePoint_To_config_UtilizationShapePoint(in *v1.UtilizationShapePoint, out *config.UtilizationShapePoint, s conversion.Scope) error {
	out.Utilization = in.Utilization
	out.Score = in.Score
	return nil
}

// Convert_v1_UtilizationShapePoint_To_config_UtilizationShapePoint is an autogenerated conversion function.
func Convert_v1_UtilizationShapePoint_To_config_UtilizationShapePoint(in *v1.UtilizationShapePoint, out *config.UtilizationShapePoint, s conversion.Scope) error {
	return autoConvert_v1_UtilizationShapePoint_To_config_UtilizationShapePoint(in, out, s)
}

func autoConvert_config_UtilizationShapePoint_To_v1_UtilizationShapePoint(in *config.UtilizationShapePoint, out *v1.UtilizationShapePoint, s conversion.Scope) error {
	out.Utilization = in.Utilization
	out.Score = in.Score
	return nil
}

// Convert_config_UtilizationShapePoint_To_v1_UtilizationShapePoint is an autogenerated conversion function.
func Convert_config_UtilizationShapePoint_To_v1_UtilizationShapePoint(in *config.UtilizationShapePoint, out *v1.UtilizationShapePoint, s conversion.Scope) error {
	return autoConvert_config_UtilizationShapePoint_To_v1_UtilizationShapePoint(in, out, s)
}
