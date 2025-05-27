package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func getVolumeTypeByIdUrl(psc lsclient.IServiceClient, popts IGetVolumeTypeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volume_types",
		popts.GetVolumeTypeId())
}

func getDefaultVolumeTypeUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volume_default_id")
}

func getVolumeTypeZonesUrl(psc lsclient.IServiceClient, popts IGetVolumeTypeZonesRequest) string {
	query, err := popts.ToQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volume_type_zones",
	) + query
}

func getVolumeTypesUrl(psc lsclient.IServiceClient, popts IGetListVolumeTypeRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		popts.GetVolumeTypeZoneId(),
		"volume_types",
	)
}
