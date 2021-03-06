// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package rediscaches

import (
	"testing"

	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

type TestContext struct {
	ResourceGroupName     string
	ResourceGroupLocation string
	ResourceGroupManager  resourcegroupsresourcemanager.ResourceGroupManager
	RedisCacheManager     *AzureRedisCacheManager
}

var tc TestContext

func TestAPIs(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping Resource Manager Eventhubs Suite")
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisCache Suite")
}

var _ = BeforeSuite(func() {
	zaplogger := zap.LoggerTo(GinkgoWriter, true)
	logf.SetLogger(zaplogger)

	By("bootstrapping test environment")

	err := resourcemanagerconfig.ParseEnvironment()
	Expect(err).ToNot(HaveOccurred())
	Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "t-rg-dev-rm-eh-" + helpers.RandomString(10)
	resourceGroupLocation := resourcemanagerconfig.DefaultLocation()
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager()

	//create resourcegroup for this suite
	_, err = resourceGroupManager.CreateGroup(context.Background(), resourceGroupName, resourceGroupLocation)
	Expect(err).ToNot(HaveOccurred())

	tc = TestContext{
		ResourceGroupName:     resourceGroupName,
		ResourceGroupLocation: resourceGroupLocation,
		RedisCacheManager:     &AzureRedisCacheManager{},
		ResourceGroupManager:  resourceGroupManager,
	}
})

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	// delete the resource group and contained resources
	_, _ = tc.ResourceGroupManager.DeleteGroup(context.Background(), tc.ResourceGroupName)
})
