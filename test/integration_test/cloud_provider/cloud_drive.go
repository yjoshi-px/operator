package cloud_provider

import (
	"github.com/libopenstorage/cloudops"
	"github.com/libopenstorage/operator/test/integration_test/utils"
	"log"
	"strings"
)

type Provider interface {
	GetDefaultDataDrives() *[]string
	GetDefaultMetadataDrive() *string
	GetDefaultKvdbDrive() *string
	GetDefaultJournalDrive() *string
}

func stringSlicePtr(slice []string) *[]string {
	return &slice
}

func stringPtr(str string) *string {
	return &str
}

func isVsphere() bool {
	return strings.Contains(utils.PxEnvVars, "VSPHERE_VCENTER") ||
		utils.CloudProvider == cloudops.Vsphere
}

func GetCloudProvider() Provider {
	if isVsphere() {
		return newVsphereProvider()
	}

	log.Fatalf("CloudProvider Unknown %v", utils.CloudProvider)

	return nil
}
