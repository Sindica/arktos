/*
Copyright 2017 The Kubernetes Authors.
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

package vsphere

import (
	"fmt"
	"hash/fnv"
	"time"

	"strings"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/test/e2e/framework"
	e2elog "k8s.io/kubernetes/test/e2e/framework/log"
	e2enode "k8s.io/kubernetes/test/e2e/framework/node"
	"k8s.io/kubernetes/test/e2e/storage/utils"
)

const (
	VmfsDatastore                              = "sharedVmfs-0"
	VsanDatastore                              = "vsanDatastore"
	Datastore                                  = "datastore"
	Policy_DiskStripes                         = "diskStripes"
	Policy_HostFailuresToTolerate              = "hostFailuresToTolerate"
	Policy_CacheReservation                    = "cacheReservation"
	Policy_ObjectSpaceReservation              = "objectSpaceReservation"
	Policy_IopsLimit                           = "iopsLimit"
	DiskFormat                                 = "diskformat"
	ThinDisk                                   = "thin"
	SpbmStoragePolicy                          = "storagepolicyname"
	BronzeStoragePolicy                        = "bronze"
	HostFailuresToTolerateCapabilityVal        = "0"
	CacheReservationCapabilityVal              = "20"
	DiskStripesCapabilityVal                   = "1"
	ObjectSpaceReservationCapabilityVal        = "30"
	IopsLimitCapabilityVal                     = "100"
	StripeWidthCapabilityVal                   = "2"
	DiskStripesCapabilityInvalidVal            = "14"
	HostFailuresToTolerateCapabilityInvalidVal = "4"
	DummyVMPrefixName                          = "vsphere-k8s"
	DiskStripesCapabilityMaxVal                = "11"
)

/*
   Test to verify the storage policy based management for dynamic volume provisioning inside kubernetes.
   There are 2 ways to achieve it:
   1. Specify VSAN storage capabilities in the storage-class.
   2. Use existing vCenter SPBM storage policies.

   Valid VSAN storage capabilities are mentioned below:
   1. hostFailuresToTolerate
   2. forceProvisioning
   3. cacheReservation
   4. diskStripes
   5. objectSpaceReservation
   6. iopsLimit

   Steps
   1. Create StorageClass with.
   		a. VSAN storage capabilities set to valid/invalid values (or)
		b. Use existing vCenter SPBM storage policies.
   2. Create PVC which uses the StorageClass created in step 1.
   3. Wait for PV to be provisioned.
   4. Wait for PVC's status to become Bound
   5. Create pod using PVC on specific node.
   6. Wait for Disk to be attached to the node.
   7. Delete pod and Wait for Volume Disk to be detached from the Node.
   8. Delete PVC, PV and Storage Class


*/

var _ = utils.SIGDescribe("Storage Policy Based Volume Provisioning [Feature:vsphere]", func() {
	f := framework.NewDefaultFramework("volume-vsan-policy")
	var (
		client       clientset.Interface
		namespace    string
		scParameters map[string]string
		policyName   string
		tagPolicy    string
		masterNode   string
	)
	ginkgo.BeforeEach(func() {
		framework.SkipUnlessProviderIs("vsphere")
		Bootstrap(f)
		client = f.ClientSet
		namespace = f.Namespace.Name
		policyName = GetAndExpectStringEnvVar(SPBMPolicyName)
		tagPolicy = GetAndExpectStringEnvVar(SPBMTagPolicy)
		e2elog.Logf("framework: %+v", f)
		scParameters = make(map[string]string)
		nodeList := framework.GetReadySchedulableNodesOrDie(f.ClientSet)
		if !(len(nodeList.Items) > 0) {
			framework.Failf("Unable to find ready and schedulable Node")
		}
		masternodes, _, err := e2enode.GetMasterAndWorkerNodes(client)
		if err != nil {
			e2elog.Logf("Unexpected error occurred: %v", err)
		}
		// TODO: write a wrapper for ExpectNoErrorWithOffset()
		framework.ExpectNoErrorWithOffset(0, err)
		gomega.Expect(masternodes).NotTo(gomega.BeEmpty())
		masterNode = masternodes.List()[0]
	})

	// Valid policy.
	ginkgo.It("verify VSAN storage capability with valid hostFailuresToTolerate and cacheReservation values is honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for VSAN policy hostFailuresToTolerate: %s, cacheReservation: %s", HostFailuresToTolerateCapabilityVal, CacheReservationCapabilityVal))
		scParameters[Policy_HostFailuresToTolerate] = HostFailuresToTolerateCapabilityVal
		scParameters[Policy_CacheReservation] = CacheReservationCapabilityVal
		e2elog.Logf("Invoking test for VSAN storage capabilities: %+v", scParameters)
		invokeValidPolicyTest(f, client, namespace, scParameters)
	})

	// Valid policy.
	ginkgo.It("verify VSAN storage capability with valid diskStripes and objectSpaceReservation values is honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for VSAN policy diskStripes: %s, objectSpaceReservation: %s", DiskStripesCapabilityVal, ObjectSpaceReservationCapabilityVal))
		scParameters[Policy_DiskStripes] = "1"
		scParameters[Policy_ObjectSpaceReservation] = "30"
		e2elog.Logf("Invoking test for VSAN storage capabilities: %+v", scParameters)
		invokeValidPolicyTest(f, client, namespace, scParameters)
	})

	// Valid policy.
	ginkgo.It("verify VSAN storage capability with valid diskStripes and objectSpaceReservation values and a VSAN datastore is honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for VSAN policy diskStripes: %s, objectSpaceReservation: %s", DiskStripesCapabilityVal, ObjectSpaceReservationCapabilityVal))
		scParameters[Policy_DiskStripes] = DiskStripesCapabilityVal
		scParameters[Policy_ObjectSpaceReservation] = ObjectSpaceReservationCapabilityVal
		scParameters[Datastore] = VsanDatastore
		e2elog.Logf("Invoking test for VSAN storage capabilities: %+v", scParameters)
		invokeValidPolicyTest(f, client, namespace, scParameters)
	})

	// Valid policy.
	ginkgo.It("verify VSAN storage capability with valid objectSpaceReservation and iopsLimit values is honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for VSAN policy objectSpaceReservation: %s, iopsLimit: %s", ObjectSpaceReservationCapabilityVal, IopsLimitCapabilityVal))
		scParameters[Policy_ObjectSpaceReservation] = ObjectSpaceReservationCapabilityVal
		scParameters[Policy_IopsLimit] = IopsLimitCapabilityVal
		e2elog.Logf("Invoking test for VSAN storage capabilities: %+v", scParameters)
		invokeValidPolicyTest(f, client, namespace, scParameters)
	})

	// Invalid VSAN storage capabilities parameters.
	ginkgo.It("verify VSAN storage capability with invalid capability name objectSpaceReserve is not honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for VSAN policy objectSpaceReserve: %s, stripeWidth: %s", ObjectSpaceReservationCapabilityVal, StripeWidthCapabilityVal))
		scParameters["objectSpaceReserve"] = ObjectSpaceReservationCapabilityVal
		scParameters[Policy_DiskStripes] = StripeWidthCapabilityVal
		e2elog.Logf("Invoking test for VSAN storage capabilities: %+v", scParameters)
		err := invokeInvalidPolicyTestNeg(client, namespace, scParameters)
		framework.ExpectError(err)
		errorMsg := "invalid option \\\"objectSpaceReserve\\\" for volume plugin kubernetes.io/vsphere-volume"
		if !strings.Contains(err.Error(), errorMsg) {
			framework.ExpectNoError(err, errorMsg)
		}
	})

	// Invalid policy on a VSAN test bed.
	// diskStripes value has to be between 1 and 12.
	ginkgo.It("verify VSAN storage capability with invalid diskStripes value is not honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for VSAN policy diskStripes: %s, cacheReservation: %s", DiskStripesCapabilityInvalidVal, CacheReservationCapabilityVal))
		scParameters[Policy_DiskStripes] = DiskStripesCapabilityInvalidVal
		scParameters[Policy_CacheReservation] = CacheReservationCapabilityVal
		e2elog.Logf("Invoking test for VSAN storage capabilities: %+v", scParameters)
		err := invokeInvalidPolicyTestNeg(client, namespace, scParameters)
		framework.ExpectError(err)
		errorMsg := "Invalid value for " + Policy_DiskStripes + "."
		if !strings.Contains(err.Error(), errorMsg) {
			framework.ExpectNoError(err, errorMsg)
		}
	})

	// Invalid policy on a VSAN test bed.
	// hostFailuresToTolerate value has to be between 0 and 3 including.
	ginkgo.It("verify VSAN storage capability with invalid hostFailuresToTolerate value is not honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for VSAN policy hostFailuresToTolerate: %s", HostFailuresToTolerateCapabilityInvalidVal))
		scParameters[Policy_HostFailuresToTolerate] = HostFailuresToTolerateCapabilityInvalidVal
		e2elog.Logf("Invoking test for VSAN storage capabilities: %+v", scParameters)
		err := invokeInvalidPolicyTestNeg(client, namespace, scParameters)
		framework.ExpectError(err)
		errorMsg := "Invalid value for " + Policy_HostFailuresToTolerate + "."
		if !strings.Contains(err.Error(), errorMsg) {
			framework.ExpectNoError(err, errorMsg)
		}
	})

	// Specify a valid VSAN policy on a non-VSAN test bed.
	// The test should fail.
	ginkgo.It("verify VSAN storage capability with non-vsan datastore is not honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for VSAN policy diskStripes: %s, objectSpaceReservation: %s and a non-VSAN datastore: %s", DiskStripesCapabilityVal, ObjectSpaceReservationCapabilityVal, VmfsDatastore))
		scParameters[Policy_DiskStripes] = DiskStripesCapabilityVal
		scParameters[Policy_ObjectSpaceReservation] = ObjectSpaceReservationCapabilityVal
		scParameters[Datastore] = VmfsDatastore
		e2elog.Logf("Invoking test for VSAN storage capabilities: %+v", scParameters)
		err := invokeInvalidPolicyTestNeg(client, namespace, scParameters)
		framework.ExpectError(err)
		errorMsg := "The specified datastore: \\\"" + VmfsDatastore + "\\\" is not a VSAN datastore. " +
			"The policy parameters will work only with VSAN Datastore."
		if !strings.Contains(err.Error(), errorMsg) {
			framework.ExpectNoError(err, errorMsg)
		}
	})

	ginkgo.It("verify an existing and compatible SPBM policy is honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for SPBM policy: %s", policyName))
		scParameters[SpbmStoragePolicy] = policyName
		scParameters[DiskFormat] = ThinDisk
		e2elog.Logf("Invoking test for SPBM storage policy: %+v", scParameters)
		invokeValidPolicyTest(f, client, namespace, scParameters)
	})

	ginkgo.It("verify clean up of stale dummy VM for dynamically provisioned pvc using SPBM policy", func() {
		scParameters[Policy_DiskStripes] = DiskStripesCapabilityMaxVal
		scParameters[Policy_ObjectSpaceReservation] = ObjectSpaceReservationCapabilityVal
		scParameters[Datastore] = VsanDatastore
		e2elog.Logf("Invoking test for SPBM storage policy: %+v", scParameters)
		kubernetesClusterName := GetAndExpectStringEnvVar(KubernetesClusterName)
		invokeStaleDummyVMTestWithStoragePolicy(client, masterNode, namespace, kubernetesClusterName, scParameters)
	})

	ginkgo.It("verify if a SPBM policy is not honored on a non-compatible datastore for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for SPBM policy: %s and datastore: %s", tagPolicy, VsanDatastore))
		scParameters[SpbmStoragePolicy] = tagPolicy
		scParameters[Datastore] = VsanDatastore
		scParameters[DiskFormat] = ThinDisk
		e2elog.Logf("Invoking test for SPBM storage policy on a non-compatible datastore: %+v", scParameters)
		err := invokeInvalidPolicyTestNeg(client, namespace, scParameters)
		framework.ExpectError(err)
		errorMsg := "User specified datastore is not compatible with the storagePolicy: \\\"" + tagPolicy + "\\\""
		if !strings.Contains(err.Error(), errorMsg) {
			framework.ExpectNoError(err, errorMsg)
		}
	})

	ginkgo.It("verify if a non-existing SPBM policy is not honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for SPBM policy: %s", BronzeStoragePolicy))
		scParameters[SpbmStoragePolicy] = BronzeStoragePolicy
		scParameters[DiskFormat] = ThinDisk
		e2elog.Logf("Invoking test for non-existing SPBM storage policy: %+v", scParameters)
		err := invokeInvalidPolicyTestNeg(client, namespace, scParameters)
		framework.ExpectError(err)
		errorMsg := "no pbm profile found with name: \\\"" + BronzeStoragePolicy + "\\"
		if !strings.Contains(err.Error(), errorMsg) {
			framework.ExpectNoError(err, errorMsg)
		}
	})

	ginkgo.It("verify an if a SPBM policy and VSAN capabilities cannot be honored for dynamically provisioned pvc using storageclass", func() {
		ginkgo.By(fmt.Sprintf("Invoking test for SPBM policy: %s with VSAN storage capabilities", policyName))
		scParameters[SpbmStoragePolicy] = policyName
		gomega.Expect(scParameters[SpbmStoragePolicy]).NotTo(gomega.BeEmpty())
		scParameters[Policy_DiskStripes] = DiskStripesCapabilityVal
		scParameters[DiskFormat] = ThinDisk
		e2elog.Logf("Invoking test for SPBM storage policy and VSAN capabilities together: %+v", scParameters)
		err := invokeInvalidPolicyTestNeg(client, namespace, scParameters)
		framework.ExpectError(err)
		errorMsg := "Cannot specify storage policy capabilities along with storage policy name. Please specify only one"
		if !strings.Contains(err.Error(), errorMsg) {
			framework.ExpectNoError(err, errorMsg)
		}
	})
})

func invokeValidPolicyTest(f *framework.Framework, client clientset.Interface, namespace string, scParameters map[string]string) {
	ginkgo.By("Creating Storage Class With storage policy params")
	storageclass, err := client.StorageV1().StorageClasses().Create(getVSphereStorageClassSpec("storagepolicysc", scParameters, nil))
	framework.ExpectNoError(err, fmt.Sprintf("Failed to create storage class with err: %v", err))
	defer client.StorageV1().StorageClasses().Delete(storageclass.Name, nil)

	ginkgo.By("Creating PVC using the Storage Class")
	pvclaim, err := framework.CreatePVC(client, namespace, getVSphereClaimSpecWithStorageClass(namespace, "2Gi", storageclass))
	framework.ExpectNoError(err)
	defer framework.DeletePersistentVolumeClaim(client, pvclaim.Name, namespace)

	var pvclaims []*v1.PersistentVolumeClaim
	pvclaims = append(pvclaims, pvclaim)
	ginkgo.By("Waiting for claim to be in bound phase")
	persistentvolumes, err := framework.WaitForPVClaimBoundPhase(client, pvclaims, framework.ClaimProvisionTimeout)
	framework.ExpectNoError(err)

	ginkgo.By("Creating pod to attach PV to the node")
	// Create pod to attach Volume to Node
	pod, err := framework.CreatePod(client, namespace, nil, pvclaims, false, "")
	framework.ExpectNoError(err)

	ginkgo.By("Verify the volume is accessible and available in the pod")
	verifyVSphereVolumesAccessible(client, pod, persistentvolumes)

	ginkgo.By("Deleting pod")
	framework.DeletePodWithWait(f, client, pod)

	ginkgo.By("Waiting for volumes to be detached from the node")
	waitForVSphereDiskToDetach(persistentvolumes[0].Spec.VsphereVolume.VolumePath, pod.Spec.NodeName)
}

func invokeInvalidPolicyTestNeg(client clientset.Interface, namespace string, scParameters map[string]string) error {
	ginkgo.By("Creating Storage Class With storage policy params")
	storageclass, err := client.StorageV1().StorageClasses().Create(getVSphereStorageClassSpec("storagepolicysc", scParameters, nil))
	framework.ExpectNoError(err, fmt.Sprintf("Failed to create storage class with err: %v", err))
	defer client.StorageV1().StorageClasses().Delete(storageclass.Name, nil)

	ginkgo.By("Creating PVC using the Storage Class")
	pvclaim, err := framework.CreatePVC(client, namespace, getVSphereClaimSpecWithStorageClass(namespace, "2Gi", storageclass))
	framework.ExpectNoError(err)
	defer framework.DeletePersistentVolumeClaim(client, pvclaim.Name, namespace)

	ginkgo.By("Waiting for claim to be in bound phase")
	err = framework.WaitForPersistentVolumeClaimPhase(v1.ClaimBound, client, pvclaim.Namespace, pvclaim.Name, framework.Poll, 2*time.Minute)
	framework.ExpectError(err)

	eventList, err := client.CoreV1().Events(pvclaim.Namespace).List(metav1.ListOptions{})
	return fmt.Errorf("Failure message: %+q", eventList.Items[0].Message)
}

func invokeStaleDummyVMTestWithStoragePolicy(client clientset.Interface, masterNode string, namespace string, clusterName string, scParameters map[string]string) {
	ginkgo.By("Creating Storage Class With storage policy params")
	storageclass, err := client.StorageV1().StorageClasses().Create(getVSphereStorageClassSpec("storagepolicysc", scParameters, nil))
	framework.ExpectNoError(err, fmt.Sprintf("Failed to create storage class with err: %v", err))
	defer client.StorageV1().StorageClasses().Delete(storageclass.Name, nil)

	ginkgo.By("Creating PVC using the Storage Class")
	pvclaim, err := framework.CreatePVC(client, namespace, getVSphereClaimSpecWithStorageClass(namespace, "2Gi", storageclass))
	framework.ExpectNoError(err)

	var pvclaims []*v1.PersistentVolumeClaim
	pvclaims = append(pvclaims, pvclaim)
	ginkgo.By("Expect claim to fail provisioning volume")
	_, err = framework.WaitForPVClaimBoundPhase(client, pvclaims, 2*time.Minute)
	framework.ExpectError(err)

	updatedClaim, err := client.CoreV1().PersistentVolumeClaims(namespace).Get(pvclaim.Name, metav1.GetOptions{})
	framework.ExpectNoError(err)
	vmName := clusterName + "-dynamic-pvc-" + string(updatedClaim.UID)
	framework.DeletePersistentVolumeClaim(client, pvclaim.Name, namespace)
	// Wait for 6 minutes to let the vSphere Cloud Provider clean up routine delete the dummy VM
	time.Sleep(6 * time.Minute)

	fnvHash := fnv.New32a()
	fnvHash.Write([]byte(vmName))
	dummyVMFullName := DummyVMPrefixName + "-" + fmt.Sprint(fnvHash.Sum32())
	errorMsg := "Dummy VM - " + vmName + "is still present. Failing the test.."
	nodeInfo := TestContext.NodeMapper.GetNodeInfo(masterNode)
	gomega.Expect(nodeInfo.VSphere.IsVMPresent(dummyVMFullName, nodeInfo.DataCenterRef)).NotTo(gomega.BeTrue(), errorMsg)
}
