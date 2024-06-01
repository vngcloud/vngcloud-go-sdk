package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func getVolumeTypeByIdUrl(psc lsclient.IServiceClient, popts IGetVolumeTypeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volume_types",
		popts.GetVolumeTypeId())
}
