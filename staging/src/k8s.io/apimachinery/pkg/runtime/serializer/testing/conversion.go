/*
Copyright 2018 The Kubernetes Authors.

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

// File modified by cherrypick from kubernetes on 05/06/2021
package testing

import (
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
)

func convertTestType1ToExternalTestType1(in *TestType1, out *ExternalTestType1, s conversion.Scope) error {
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	out.A = in.A
	out.B = in.B
	out.C = in.C
	out.D = in.D
	out.E = in.E
	out.F = in.F
	out.G = in.G
	out.H = in.H
	out.I = in.I
	out.J = in.J
	out.K = in.K
	out.L = in.L
	out.M = in.M
	if in.N != nil {
		out.N = make(map[string]ExternalTestType2)
		for key := range in.N {
			in, tmp := in.N[key], ExternalTestType2{}
			if err := convertTestType2ToExternalTestType2(&in, &tmp, s); err != nil {
				return err
			}
			out.N[key] = tmp
		}
	} else {
		out.N = nil
	}
	if in.O != nil {
		out.O = new(ExternalTestType2)
		if err := convertTestType2ToExternalTestType2(in.O, out.O, s); err != nil {
			return err
		}
	} else {
		out.O = nil
	}
	if in.P != nil {
		out.P = make([]ExternalTestType2, len(in.P))
		for i := range in.P {
			if err := convertTestType2ToExternalTestType2(&in.P[i], &out.P[i], s); err != nil {
				return err
			}
		}
	}
	return nil
}

func convertExternalTestType1ToTestType1(in *ExternalTestType1, out *TestType1, s conversion.Scope) error {
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	out.A = in.A
	out.B = in.B
	out.C = in.C
	out.D = in.D
	out.E = in.E
	out.F = in.F
	out.G = in.G
	out.H = in.H
	out.I = in.I
	out.J = in.J
	out.K = in.K
	out.L = in.L
	out.M = in.M
	if in.N != nil {
		out.N = make(map[string]TestType2)
		for key := range in.N {
			in, tmp := in.N[key], TestType2{}
			if err := convertExternalTestType2ToTestType2(&in, &tmp, s); err != nil {
				return err
			}
			out.N[key] = tmp
		}
	} else {
		out.N = nil
	}
	if in.O != nil {
		out.O = new(TestType2)
		if err := convertExternalTestType2ToTestType2(in.O, out.O, s); err != nil {
			return err
		}
	} else {
		out.O = nil
	}
	if in.P != nil {
		out.P = make([]TestType2, len(in.P))
		for i := range in.P {
			if err := convertExternalTestType2ToTestType2(&in.P[i], &out.P[i], s); err != nil {
				return err
			}
		}
	}
	return nil
}

func convertTestType2ToExternalTestType2(in *TestType2, out *ExternalTestType2, s conversion.Scope) error {
	out.A = in.A
	out.B = in.B
	return nil
}

func convertExternalTestType2ToTestType2(in *ExternalTestType2, out *TestType2, s conversion.Scope) error {
	out.A = in.A
	out.B = in.B
	return nil
}

func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddConversionFunc((*TestType1)(nil), (*ExternalTestType1)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return convertTestType1ToExternalTestType1(a.(*TestType1), b.(*ExternalTestType1), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ExternalTestType1)(nil), (*TestType1)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return convertExternalTestType1ToTestType1(a.(*ExternalTestType1), b.(*TestType1), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*TestType2)(nil), (*ExternalTestType2)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return convertTestType2ToExternalTestType2(a.(*TestType2), b.(*ExternalTestType2), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ExternalTestType2)(nil), (*TestType2)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return convertExternalTestType2ToTestType2(a.(*ExternalTestType2), b.(*TestType2), scope)
	}); err != nil {
		return err
	}
	return nil
}
