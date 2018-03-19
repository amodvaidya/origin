// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	build_v1 "github.com/openshift/origin/pkg/build/apis/build/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildDefaultsConfig, InType: reflect.TypeOf(&BuildDefaultsConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SourceStrategyDefaultsConfig, InType: reflect.TypeOf(&SourceStrategyDefaultsConfig{})},
	)
}

// DeepCopy_v1_BuildDefaultsConfig is an autogenerated deepcopy function.
func DeepCopy_v1_BuildDefaultsConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildDefaultsConfig)
		out := out.(*BuildDefaultsConfig)
		*out = *in
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.SourceStrategyDefaults != nil {
			in, out := &in.SourceStrategyDefaults, &out.SourceStrategyDefaults
			*out = new(SourceStrategyDefaultsConfig)
			if err := DeepCopy_v1_SourceStrategyDefaultsConfig(*in, *out, c); err != nil {
				return err
			}
		}
		if in.ImageLabels != nil {
			in, out := &in.ImageLabels, &out.ImageLabels
			*out = make([]build_v1.ImageLabel, len(*in))
			copy(*out, *in)
		}
		if in.NodeSelector != nil {
			in, out := &in.NodeSelector, &out.NodeSelector
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.Annotations != nil {
			in, out := &in.Annotations, &out.Annotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if err := api_v1.DeepCopy_v1_ResourceRequirements(&in.Resources, &out.Resources, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_SourceStrategyDefaultsConfig is an autogenerated deepcopy function.
func DeepCopy_v1_SourceStrategyDefaultsConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SourceStrategyDefaultsConfig)
		out := out.(*SourceStrategyDefaultsConfig)
		*out = *in
		if in.Incremental != nil {
			in, out := &in.Incremental, &out.Incremental
			*out = new(bool)
			**out = **in
		}
		return nil
	}
}