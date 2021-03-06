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

package testsuites

import (
	"context"
	"flag"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/onsi/ginkgo"

	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	clientset "k8s.io/client-go/kubernetes"
	csitrans "k8s.io/csi-translation-lib"
	"k8s.io/kubernetes/test/e2e/framework"
	e2elog "k8s.io/kubernetes/test/e2e/framework/log"
	"k8s.io/kubernetes/test/e2e/framework/metrics"
	"k8s.io/kubernetes/test/e2e/framework/podlogs"
	"k8s.io/kubernetes/test/e2e/framework/volume"
	"k8s.io/kubernetes/test/e2e/storage/testpatterns"
)

var (
	migratedPlugins *string
)

func init() {
	migratedPlugins = flag.String("storage.migratedPlugins", "", "comma separated list of in-tree plugin names of form 'kubernetes.io/{pluginName}' migrated to CSI")
}

type opCounts map[string]int64

// TestSuite represents an interface for a set of tests which works with TestDriver
type TestSuite interface {
	// getTestSuiteInfo returns the TestSuiteInfo for this TestSuite
	getTestSuiteInfo() TestSuiteInfo
	// defineTest defines tests of the testpattern for the driver.
	// Called inside a Ginkgo context that reflects the current driver and test pattern,
	// so the test suite can define tests directly with ginkgo.It.
	defineTests(TestDriver, testpatterns.TestPattern)
}

// TestSuiteInfo represents a set of parameters for TestSuite
type TestSuiteInfo struct {
	name         string                     // name of the TestSuite
	featureTag   string                     // featureTag for the TestSuite
	testPatterns []testpatterns.TestPattern // Slice of TestPattern for the TestSuite
}

// TestResource represents an interface for resources that is used by TestSuite
type TestResource interface {
	// cleanupResource cleans up the test resources created when setting up the resource
	cleanupResource()
}

func getTestNameStr(suite TestSuite, pattern testpatterns.TestPattern) string {
	tsInfo := suite.getTestSuiteInfo()
	return fmt.Sprintf("[Testpattern: %s]%s %s%s", pattern.Name, pattern.FeatureTag, tsInfo.name, tsInfo.featureTag)
}

// DefineTestSuite defines tests for all testpatterns and all testSuites for a driver
func DefineTestSuite(driver TestDriver, tsInits []func() TestSuite) {
	for _, testSuiteInit := range tsInits {
		suite := testSuiteInit()
		for _, pattern := range suite.getTestSuiteInfo().testPatterns {
			p := pattern
			ginkgo.Context(getTestNameStr(suite, p), func() {
				ginkgo.BeforeEach(func() {
					// Skip unsupported tests to avoid unnecessary resource initialization
					skipUnsupportedTest(driver, p)
				})
				suite.defineTests(driver, p)
			})
		}
	}
}

// skipUnsupportedTest will skip tests if the combination of driver,  and testpattern
// is not suitable to be tested.
// Whether it needs to be skipped is checked by following steps:
// 1. Check if Whether SnapshotType is supported by driver from its interface
// 2. Check if Whether volType is supported by driver from its interface
// 3. Check if fsType is supported
// 4. Check with driver specific logic
//
// Test suites can also skip tests inside their own defineTests function or in
// individual tests.
func skipUnsupportedTest(driver TestDriver, pattern testpatterns.TestPattern) {
	dInfo := driver.GetDriverInfo()
	var isSupported bool

	// 1. Check if Whether SnapshotType is supported by driver from its interface
	// if isSupported, we still execute the driver and suite tests
	if len(pattern.SnapshotType) > 0 {
		switch pattern.SnapshotType {
		case testpatterns.DynamicCreatedSnapshot:
			_, isSupported = driver.(SnapshottableTestDriver)
		default:
			isSupported = false
		}
		if !isSupported {
			framework.Skipf("Driver %s doesn't support snapshot type %v -- skipping", dInfo.Name, pattern.SnapshotType)
		}
	} else {
		// 2. Check if Whether volType is supported by driver from its interface
		switch pattern.VolType {
		case testpatterns.InlineVolume:
			_, isSupported = driver.(InlineVolumeTestDriver)
		case testpatterns.PreprovisionedPV:
			_, isSupported = driver.(PreprovisionedPVTestDriver)
		case testpatterns.DynamicPV:
			_, isSupported = driver.(DynamicPVTestDriver)
		default:
			isSupported = false
		}

		if !isSupported {
			framework.Skipf("Driver %s doesn't support %v -- skipping", dInfo.Name, pattern.VolType)
		}

		// 3. Check if fsType is supported
		if !dInfo.SupportedFsType.Has(pattern.FsType) {
			framework.Skipf("Driver %s doesn't support %v -- skipping", dInfo.Name, pattern.FsType)
		}
		if pattern.FsType == "xfs" && framework.NodeOSDistroIs("gci", "cos", "windows") {
			framework.Skipf("Distro doesn't support xfs -- skipping")
		}
		if pattern.FsType == "ntfs" && !framework.NodeOSDistroIs("windows") {
			framework.Skipf("Distro %s doesn't support ntfs -- skipping", framework.TestContext.NodeOSDistro)
		}
	}

	// 4. Check with driver specific logic
	driver.SkipUnsupportedTest(pattern)
}

// genericVolumeTestResource is a generic implementation of TestResource that wil be able to
// be used in most of TestSuites.
// See volume_io.go or volumes.go in test/e2e/storage/testsuites/ for how to use this resource.
// Also, see subpath.go in the same directory for how to extend and use it.
type genericVolumeTestResource struct {
	driver    TestDriver
	config    *PerTestConfig
	pattern   testpatterns.TestPattern
	volType   string
	volSource *v1.VolumeSource
	pvc       *v1.PersistentVolumeClaim
	pv        *v1.PersistentVolume
	sc        *storagev1.StorageClass

	volume TestVolume
}

var _ TestResource = &genericVolumeTestResource{}

func createGenericVolumeTestResource(driver TestDriver, config *PerTestConfig, pattern testpatterns.TestPattern) *genericVolumeTestResource {
	r := genericVolumeTestResource{
		driver:  driver,
		config:  config,
		pattern: pattern,
	}
	dInfo := driver.GetDriverInfo()
	f := config.Framework
	cs := f.ClientSet
	fsType := pattern.FsType
	volType := pattern.VolType

	// Create volume for pre-provisioned volume tests
	r.volume = CreateVolume(driver, config, volType)

	switch volType {
	case testpatterns.InlineVolume:
		e2elog.Logf("Creating resource for inline volume")
		if iDriver, ok := driver.(InlineVolumeTestDriver); ok {
			r.volSource = iDriver.GetVolumeSource(false, fsType, r.volume)
			r.volType = dInfo.Name
		}
	case testpatterns.PreprovisionedPV:
		e2elog.Logf("Creating resource for pre-provisioned PV")
		if pDriver, ok := driver.(PreprovisionedPVTestDriver); ok {
			pvSource, volumeNodeAffinity := pDriver.GetPersistentVolumeSource(false, fsType, r.volume)
			if pvSource != nil {
				r.volSource, r.pv, r.pvc = createVolumeSourceWithPVCPV(f, dInfo.Name, pvSource, volumeNodeAffinity, false, pattern.VolMode)
			}
			r.volType = fmt.Sprintf("%s-preprovisionedPV", dInfo.Name)
		}
	case testpatterns.DynamicPV:
		e2elog.Logf("Creating resource for dynamic PV")
		if dDriver, ok := driver.(DynamicPVTestDriver); ok {
			claimSize := dDriver.GetClaimSize()
			r.sc = dDriver.GetDynamicProvisionStorageClass(r.config, fsType)

			ginkgo.By("creating a StorageClass " + r.sc.Name)
			var err error
			r.sc, err = cs.StorageV1().StorageClasses().Create(r.sc)
			framework.ExpectNoError(err)

			if r.sc != nil {
				r.volSource, r.pv, r.pvc = createVolumeSourceWithPVCPVFromDynamicProvisionSC(
					f, dInfo.Name, claimSize, r.sc, false, pattern.VolMode)
			}
			r.volType = fmt.Sprintf("%s-dynamicPV", dInfo.Name)
		}
	default:
		framework.Failf("genericVolumeTestResource doesn't support: %s", volType)
	}

	if r.volSource == nil {
		framework.Skipf("Driver %s doesn't support %v -- skipping", dInfo.Name, volType)
	}

	return &r
}

// cleanupResource cleans up genericVolumeTestResource
func (r *genericVolumeTestResource) cleanupResource() {
	f := r.config.Framework
	volType := r.pattern.VolType

	if r.pvc != nil || r.pv != nil {
		switch volType {
		case testpatterns.PreprovisionedPV:
			ginkgo.By("Deleting pv and pvc")
			if errs := framework.PVPVCCleanup(f.ClientSet, f.Namespace.Name, r.pv, r.pvc); len(errs) != 0 {
				framework.Failf("Failed to delete PVC or PV: %v", utilerrors.NewAggregate(errs))
			}
		case testpatterns.DynamicPV:
			ginkgo.By("Deleting pvc")
			// We only delete the PVC so that PV (and disk) can be cleaned up by dynamic provisioner
			if r.pv != nil && r.pv.Spec.PersistentVolumeReclaimPolicy != v1.PersistentVolumeReclaimDelete {
				framework.Failf("Test framework does not currently support Dynamically Provisioned Persistent Volume %v specified with reclaim policy that isnt %v",
					r.pv.Name, v1.PersistentVolumeReclaimDelete)
			}
			if r.pvc != nil {
				err := framework.DeletePersistentVolumeClaim(f.ClientSet, r.pvc.Name, f.Namespace.Name)
				framework.ExpectNoError(err, "Failed to delete PVC %v", r.pvc.Name)
				if r.pv != nil {
					err = framework.WaitForPersistentVolumeDeleted(f.ClientSet, r.pv.Name, 5*time.Second, 5*time.Minute)
					framework.ExpectNoError(err, "Persistent Volume %v not deleted by dynamic provisioner", r.pv.Name)
				}
			}
		default:
			framework.Failf("Found PVC (%v) or PV (%v) but not running Preprovisioned or Dynamic test pattern", r.pvc, r.pv)
		}
	}

	if r.sc != nil {
		ginkgo.By("Deleting sc")
		deleteStorageClass(f.ClientSet, r.sc.Name)
	}

	// Cleanup volume for pre-provisioned volume tests
	if r.volume != nil {
		r.volume.DeleteVolume()
	}
}

func createVolumeSourceWithPVCPV(
	f *framework.Framework,
	name string,
	pvSource *v1.PersistentVolumeSource,
	volumeNodeAffinity *v1.VolumeNodeAffinity,
	readOnly bool,
	volMode v1.PersistentVolumeMode,
) (*v1.VolumeSource, *v1.PersistentVolume, *v1.PersistentVolumeClaim) {
	pvConfig := framework.PersistentVolumeConfig{
		NamePrefix:       fmt.Sprintf("%s-", name),
		StorageClassName: f.Namespace.Name,
		PVSource:         *pvSource,
		NodeAffinity:     volumeNodeAffinity,
	}

	pvcConfig := framework.PersistentVolumeClaimConfig{
		StorageClassName: &f.Namespace.Name,
	}

	if volMode != "" {
		pvConfig.VolumeMode = &volMode
		pvcConfig.VolumeMode = &volMode
	}

	e2elog.Logf("Creating PVC and PV")
	pv, pvc, err := framework.CreatePVCPV(f.ClientSet, pvConfig, pvcConfig, f.Namespace.Name, false)
	framework.ExpectNoError(err, "PVC, PV creation failed")

	err = framework.WaitOnPVandPVC(f.ClientSet, f.Namespace.Name, pv, pvc)
	framework.ExpectNoError(err, "PVC, PV failed to bind")

	volSource := &v1.VolumeSource{
		PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
			ClaimName: pvc.Name,
			ReadOnly:  readOnly,
		},
	}
	return volSource, pv, pvc
}

func createVolumeSourceWithPVCPVFromDynamicProvisionSC(
	f *framework.Framework,
	name string,
	claimSize string,
	sc *storagev1.StorageClass,
	readOnly bool,
	volMode v1.PersistentVolumeMode,
) (*v1.VolumeSource, *v1.PersistentVolume, *v1.PersistentVolumeClaim) {
	cs := f.ClientSet
	ns := f.Namespace.Name

	ginkgo.By("creating a claim")
	pvc := getClaim(claimSize, ns)
	pvc.Spec.StorageClassName = &sc.Name
	if volMode != "" {
		pvc.Spec.VolumeMode = &volMode
	}

	var err error
	pvc, err = cs.CoreV1().PersistentVolumeClaims(ns).Create(pvc)
	framework.ExpectNoError(err)

	if !isDelayedBinding(sc) {
		err = framework.WaitForPersistentVolumeClaimPhase(v1.ClaimBound, cs, pvc.Namespace, pvc.Name, framework.Poll, framework.ClaimProvisionTimeout)
		framework.ExpectNoError(err)
	}

	pvc, err = cs.CoreV1().PersistentVolumeClaims(pvc.Namespace).Get(pvc.Name, metav1.GetOptions{})
	framework.ExpectNoError(err)

	var pv *v1.PersistentVolume
	if !isDelayedBinding(sc) {
		pv, err = cs.CoreV1().PersistentVolumes().Get(pvc.Spec.VolumeName, metav1.GetOptions{})
		framework.ExpectNoError(err)
	}

	volSource := &v1.VolumeSource{
		PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
			ClaimName: pvc.Name,
			ReadOnly:  readOnly,
		},
	}
	return volSource, pv, pvc
}

func isDelayedBinding(sc *storagev1.StorageClass) bool {
	if sc.VolumeBindingMode != nil {
		return *sc.VolumeBindingMode == storagev1.VolumeBindingWaitForFirstConsumer
	}
	return false
}

func getClaim(claimSize string, ns string) *v1.PersistentVolumeClaim {
	claim := v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "pvc-",
			Namespace:    ns,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{
				v1.ReadWriteOnce,
			},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceName(v1.ResourceStorage): resource.MustParse(claimSize),
				},
			},
		},
	}

	return &claim
}

// deleteStorageClass deletes the passed in StorageClass and catches errors other than "Not Found"
func deleteStorageClass(cs clientset.Interface, className string) {
	err := cs.StorageV1().StorageClasses().Delete(className, nil)
	if err != nil && !apierrs.IsNotFound(err) {
		framework.ExpectNoError(err)
	}
}

// convertTestConfig returns a framework test config with the
// parameters specified for the testsuite or (if available) the
// dynamically created config for the volume server.
//
// This is done because TestConfig is the public API for
// the testsuites package whereas volume.TestConfig is merely
// an implementation detail. It contains fields that have no effect,
// which makes it unsuitable for use in the testsuits public API.
func convertTestConfig(in *PerTestConfig) volume.TestConfig {
	if in.ServerConfig != nil {
		return *in.ServerConfig
	}

	return volume.TestConfig{
		Namespace:      in.Framework.Namespace.Name,
		Prefix:         in.Prefix,
		ClientNodeName: in.ClientNodeName,
		NodeSelector:   in.ClientNodeSelector,
	}
}

func getSnapshot(claimName string, ns, snapshotClassName string) *unstructured.Unstructured {
	snapshot := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "VolumeSnapshot",
			"apiVersion": snapshotAPIVersion,
			"metadata": map[string]interface{}{
				"generateName": "snapshot-",
				"namespace":    ns,
			},
			"spec": map[string]interface{}{
				"snapshotClassName": snapshotClassName,
				"source": map[string]interface{}{
					"name": claimName,
					"kind": "PersistentVolumeClaim",
				},
			},
		},
	}

	return snapshot
}

// StartPodLogs begins capturing log output and events from current
// and future pods running in the namespace of the framework. That
// ends when the returned cleanup function is called.
//
// The output goes to log files (when using --report-dir, as in the
// CI) or the output stream (otherwise).
func StartPodLogs(f *framework.Framework) func() {
	ctx, cancel := context.WithCancel(context.Background())
	cs := f.ClientSet
	ns := f.Namespace

	to := podlogs.LogOutput{
		StatusWriter: ginkgo.GinkgoWriter,
	}
	if framework.TestContext.ReportDir == "" {
		to.LogWriter = ginkgo.GinkgoWriter
	} else {
		test := ginkgo.CurrentGinkgoTestDescription()
		reg := regexp.MustCompile("[^a-zA-Z0-9_-]+")
		// We end the prefix with a slash to ensure that all logs
		// end up in a directory named after the current test.
		//
		// TODO: use a deeper directory hierarchy once gubernator
		// supports that (https://github.com/kubernetes/test-infra/issues/10289).
		to.LogPathPrefix = framework.TestContext.ReportDir + "/" +
			reg.ReplaceAllString(test.FullTestText, "_") + "/"
	}
	podlogs.CopyAllLogs(ctx, cs, ns.Name, to)

	// pod events are something that the framework already collects itself
	// after a failed test. Logging them live is only useful for interactive
	// debugging, not when we collect reports.
	if framework.TestContext.ReportDir == "" {
		podlogs.WatchPods(ctx, cs, ns.Name, ginkgo.GinkgoWriter)
	}

	return cancel
}

func getVolumeOpsFromMetricsForPlugin(ms metrics.Metrics, pluginName string) opCounts {
	totOps := opCounts{}

	for method, samples := range ms {
		switch method {
		case "storage_operation_status_count":
			for _, sample := range samples {
				plugin := string(sample.Metric["volume_plugin"])
				if pluginName != plugin {
					continue
				}
				opName := string(sample.Metric["operation_name"])
				if opName == "verify_controller_attached_volume" {
					// We ignore verify_controller_attached_volume because it does not call into
					// the plugin. It only watches Node API and updates Actual State of World cache
					continue
				}
				totOps[opName] = totOps[opName] + int64(sample.Value)
			}
		}
	}
	return totOps
}

func getVolumeOpCounts(c clientset.Interface, pluginName string) opCounts {
	if !framework.ProviderIs("gce", "gke", "aws") {
		return opCounts{}
	}

	nodeLimit := 25

	metricsGrabber, err := metrics.NewMetricsGrabber(c, nil, true, false, true, false, false)

	if err != nil {
		framework.ExpectNoError(err, "Error creating metrics grabber: %v", err)
	}

	if !metricsGrabber.HasRegisteredMaster() {
		e2elog.Logf("Warning: Environment does not support getting controller-manager metrics")
		return opCounts{}
	}

	controllerMetrics, err := metricsGrabber.GrabFromControllerManager()
	framework.ExpectNoError(err, "Error getting c-m metrics : %v", err)
	totOps := getVolumeOpsFromMetricsForPlugin(metrics.Metrics(controllerMetrics), pluginName)

	e2elog.Logf("Node name not specified for getVolumeOpCounts, falling back to listing nodes from API Server")
	nodes, err := c.CoreV1().Nodes().List(metav1.ListOptions{})
	framework.ExpectNoError(err, "Error listing nodes: %v", err)
	if len(nodes.Items) <= nodeLimit {
		// For large clusters with > nodeLimit nodes it is too time consuming to
		// gather metrics from all nodes. We just ignore the node metrics
		// for those clusters
		for _, node := range nodes.Items {
			nodeMetrics, err := metricsGrabber.GrabFromKubelet(node.GetName())
			framework.ExpectNoError(err, "Error getting Kubelet %v metrics: %v", node.GetName(), err)
			totOps = addOpCounts(totOps, getVolumeOpsFromMetricsForPlugin(metrics.Metrics(nodeMetrics), pluginName))
		}
	} else {
		e2elog.Logf("Skipping operation metrics gathering from nodes in getVolumeOpCounts, greater than %v nodes", nodeLimit)
	}

	return totOps
}

func addOpCounts(o1 opCounts, o2 opCounts) opCounts {
	totOps := opCounts{}
	seen := sets.NewString()
	for op, count := range o1 {
		seen.Insert(op)
		totOps[op] = totOps[op] + count + o2[op]
	}
	for op, count := range o2 {
		if !seen.Has(op) {
			totOps[op] = totOps[op] + count
		}
	}
	return totOps
}

func getMigrationVolumeOpCounts(cs clientset.Interface, pluginName string) (opCounts, opCounts) {
	if len(pluginName) > 0 {
		var migratedOps opCounts
		l := csitrans.New()
		csiName, err := l.GetCSINameFromInTreeName(pluginName)
		if err != nil {
			e2elog.Logf("Could not find CSI Name for in-tree plugin %v", pluginName)
			migratedOps = opCounts{}
		} else {
			csiName = "kubernetes.io/csi:" + csiName
			migratedOps = getVolumeOpCounts(cs, csiName)
		}
		return getVolumeOpCounts(cs, pluginName), migratedOps
	} else {
		// Not an in-tree driver
		e2elog.Logf("Test running for native CSI Driver, not checking metrics")
		return opCounts{}, opCounts{}
	}
}

func getTotOps(ops opCounts) int64 {
	var tot int64 = 0
	for _, count := range ops {
		tot += count
	}
	return tot
}

func validateMigrationVolumeOpCounts(cs clientset.Interface, pluginName string, oldInTreeOps, oldMigratedOps opCounts) {
	if len(pluginName) == 0 {
		// This is a native CSI Driver and we don't check ops
		return
	}

	if sets.NewString(strings.Split(*migratedPlugins, ",")...).Has(pluginName) {
		// If this plugin is migrated based on the test flag storage.migratedPlugins
		newInTreeOps, _ := getMigrationVolumeOpCounts(cs, pluginName)

		for op, count := range newInTreeOps {
			if count != oldInTreeOps[op] {
				framework.Failf("In-tree plugin %v migrated to CSI Driver, however found %v %v metrics for in-tree plugin", pluginName, count-oldInTreeOps[op], op)
			}
		}
		// We don't check for migrated metrics because some negative test cases
		// may not do any volume operations and therefore not emit any metrics
	} else {
		// In-tree plugin is not migrated
		e2elog.Logf("In-tree plugin %v is not migrated, not validating any metrics", pluginName)

		// We don't check in-tree plugin metrics because some negative test
		// cases may not do any volume operations and therefore not emit any
		// metrics

		// We don't check counts for the Migrated version of the driver because
		// if tests are running in parallel a test could be using the CSI Driver
		// natively and increase the metrics count

		// TODO(dyzz): Add a dimension to OperationGenerator metrics for
		// "migrated"->true/false so that we can disambiguate migrated metrics
		// and native CSI Driver metrics. This way we can check the counts for
		// migrated version of the driver for stronger negative test case
		// guarantees (as well as more informative metrics).
	}
}
