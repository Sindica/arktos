/*
Copyright 2014 The Kubernetes Authors.
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

package fake

import (
	"k8s.io/api/extensions/v1beta1"
	core "k8s.io/client-go/testing"
)

func (c *FakeDeployments) Rollback(deploymentRollback *v1beta1.DeploymentRollback) error {
	action := core.CreateActionImpl{}
	action.Verb = "create"
	action.Tenant = c.te
	action.Namespace = c.ns
	action.Resource = deploymentsResource
	action.Subresource = "rollback"
	action.Object = deploymentRollback

	_, err := c.Fake.Invokes(action, deploymentRollback)
	return err
}
