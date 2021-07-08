#!/usr/bin/env bash

# Copyright 2014 The Kubernetes Authors.
# Copyright 2020 Authors of Arktos - file modified.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
source "${KUBE_ROOT}/hack/lib/init.sh"
# Lists of API Versions of each groups that should be tested, groups are
# separated by comma, lists are separated by semicolon. e.g.,
# "v1,compute/v1alpha1,experimental/v1alpha2;v1,compute/v2,experimental/v1alpha3"
# TODO: It's going to be:
# KUBE_TEST_API_VERSIONS=${KUBE_TEST_API_VERSIONS:-"v1,extensions/v1beta1"}
# FIXME: due to current implementation of a test client (see: pkg/api/testapi/testapi.go)
# ONLY the last version is tested in each group.
ALL_VERSIONS_CSV=$(IFS=',';echo "${KUBE_AVAILABLE_GROUP_VERSIONS[*]// /,}";IFS=$)
KUBE_TEST_API_VERSIONS="${KUBE_TEST_API_VERSIONS:-${ALL_VERSIONS_CSV}}"

# start the cache mutation detector by default so that cache mutators will be found
KUBE_CACHE_MUTATION_DETECTOR="${KUBE_CACHE_MUTATION_DETECTOR:-true}"
export KUBE_CACHE_MUTATION_DETECTOR

# panic the server on watch decode errors since they are considered coder mistakes
KUBE_PANIC_WATCH_DECODE_ERROR="${KUBE_PANIC_WATCH_DECODE_ERROR:-true}"
export KUBE_PANIC_WATCH_DECODE_ERROR

# Give integration tests longer to run by default.
KUBE_TIMEOUT=${KUBE_TIMEOUT:--timeout=600s}
KUBE_INTEGRATION_TEST_MAX_CONCURRENCY=${KUBE_INTEGRATION_TEST_MAX_CONCURRENCY:-"-1"}
LOG_LEVEL=${LOG_LEVEL:-2}
KUBE_TEST_ARGS=${KUBE_TEST_ARGS:-}
# Default glog module settings.
KUBE_TEST_VMODULE=${KUBE_TEST_VMODULE:-"garbagecollector*=6,graph_builder*=6"}

kube::test::find_integration_test_dirs() {
  (
    cd "${KUBE_ROOT}"
    find test/integration/ -name '*_test.go' -print0 \
      | xargs -0n1 dirname | sed "s|^|${KUBE_GO_PACKAGE}/|" \
      | LC_ALL=C sort -u
    find vendor/k8s.io/apiextensions-apiserver/test/integration/ -name '*_test.go' -print0 \
      | xargs -0n1 dirname | sed "s|^|${KUBE_GO_PACKAGE}/|" \
      | LC_ALL=C sort -u
  )
}

CLEANUP_REQUIRED=
cleanup() {
  if [[ -z "${CLEANUP_REQUIRED}" ]]; then
    return
  fi
  kube::log::status "Cleaning up etcd"
  kube::etcd::cleanup
  CLEANUP_REQUIRED=
  kube::log::status "Integration test cleanup complete"
}

runTests() {
  kube::log::status "Starting etcd instance"
  CLEANUP_REQUIRED=1
  kube::etcd::start
  kube::log::status "Running integration test cases"

  # export KUBE_RACE
  #
  # Enable the Go race detector.
  export KUBE_RACE="-race"
  make -C "${KUBE_ROOT}" test \
      WHAT="${WHAT:-$(kube::test::find_integration_test_dirs | paste -sd' ' -)}" \
      GOFLAGS="${GOFLAGS:-}" \
      KUBE_TEST_ARGS="--alsologtostderr=true ${KUBE_TEST_ARGS:-} ${SHORT:--short=true} --vmodule=${KUBE_TEST_VMODULE}" \
      KUBE_RACE="" \
      KUBE_TIMEOUT="${KUBE_TIMEOUT}" \
      KUBE_TEST_API_VERSIONS="$1"

  cleanup
}

checkEtcdOnPath() {
  kube::log::status "Checking etcd is on PATH"
  which etcd && return
  kube::log::status "Cannot find etcd, cannot run integration tests."
  kube::log::status "Please see https://git.k8s.io/community/contributors/devel/sig-testing/integration-tests.md#install-etcd-dependency for instructions."
  kube::log::usage "You can use 'hack/install-etcd.sh' to install a copy in third_party/."
  return 1
}

checkEtcdOnPath

# Run cleanup to stop etcd on interrupt or other kill signal.
trap cleanup EXIT

# Convert the CSV to an array of API versions to test
IFS=';' read -ra apiVersions <<< "${KUBE_TEST_API_VERSIONS}"
for apiVersion in "${apiVersions[@]}"; do
  runTests "${apiVersion}"
done
