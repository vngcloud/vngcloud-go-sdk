package volume_type

import "github.com/vngcloud/vngcloud-go-sdk/client"

func getVolumeTypeURL(pSc *client.ServiceClient, pOpts IGetVolumeTypeOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volume_types",
		pOpts.GetVolumeTypeID())
}
