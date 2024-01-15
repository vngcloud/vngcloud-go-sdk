package volume_actions

import "github.com/vngcloud/vngcloud-go-sdk/client"

func resizeURL(pSc *client.ServiceClient, pOpts IResizeOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes",
		pOpts.GetVolumeID(),
		"resize")
}

func mappingURL(pSc *client.ServiceClient, pOpts IMappingOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes",
		pOpts.GetVolumeID(),
		"mapping",
	)
}
