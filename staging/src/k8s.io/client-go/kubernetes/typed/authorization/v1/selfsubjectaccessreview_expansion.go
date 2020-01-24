/*
Copyright 2017 The Kubernetes Authors.
Copyright 2020 Authors of Alkaid - file modified.

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

package v1

import (
	authorizationapi "k8s.io/api/authorization/v1"
)

type SelfSubjectAccessReviewExpansion interface {
	Create(sar *authorizationapi.SelfSubjectAccessReview) (result *authorizationapi.SelfSubjectAccessReview, err error)
}

func (c *selfSubjectAccessReviews) Create(sar *authorizationapi.SelfSubjectAccessReview) (result *authorizationapi.SelfSubjectAccessReview, err error) {
	result = &authorizationapi.SelfSubjectAccessReview{}
	err = c.client.Post().
		Tenant(c.te).
		Resource("selfsubjectaccessreviews").
		Body(sar).
		Do().
		Into(result)
	return
}
