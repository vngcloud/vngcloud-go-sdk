package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func createBlockVolumeUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes")
}

func deleteBlockVolumeByIdUrl(psc lsclient.IServiceClient, popts IDeleteBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId())
}

func listBlockVolumesUrl(psc lsclient.IServiceClient, popts IListBlockVolumesRequest) string {
	query, err := popts.ToQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes") + query
}
