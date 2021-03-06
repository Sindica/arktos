/*
Copyright 2018 The Kubernetes Authors.
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

// File modified by cherrypick from kubernetes on 03/04/2021
package internal

import (
	"bytes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/structured-merge-diff/v3/fieldpath"
)

// EmptyFields represents a set with no paths
// It looks like metav1.Fields{Raw: []byte("{}")}
var EmptyFields metav1.Fields = func() metav1.Fields {
	f, err := SetToFields(*fieldpath.NewSet())
	if err != nil {
		panic("should never happen")
	}
	return f
}()

// FieldsToSet creates a set paths from an input trie of fields
func FieldsToSet(f metav1.Fields) (s fieldpath.Set, err error) {
	err = s.FromJSON(bytes.NewReader(f.Raw))
	return s, err
}

// SetToFields creates a trie of fields from an input set of paths
func SetToFields(s fieldpath.Set) (f metav1.Fields, err error) {
	f.Raw, err = s.ToJSON()
	return f, err
}
