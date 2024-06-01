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

func getBlockVolumeByIdUrl(psc lsclient.IServiceClient, popts IGetBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId())
}

func resizeBlockVolumeByIdUrl(psc lsclient.IServiceClient, popts IResizeBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"resize")
}

func listSnapshotsByBlockVolumeIdUrl(psc lsclient.IServiceClient, popts IListSnapshotsByBlockVolumeIdRequest) string {
	query, err := popts.ToQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"snapshots",
	) + query
}

func createSnapshotByBlockVolumeIdUrl(psc lsclient.IServiceClient, popts ICreateSnapshotByBlockVolumeIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"snapshots")
}

func deleteSnapshotByIdUrl(psc lsclient.IServiceClient, popts IDeleteSnapshotByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"snapshots",
		popts.GetSnapshotId(),
	)
}

func getUnderBlockVolumeIdUrl(psc lsclient.IServiceClient, popts IGetUnderBlockVolumeIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"mapping",
	)
}

func migrateBlockVolumeByIdUrl(psc lsclient.IServiceClient, popts IMigrateBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"change-device-type",
	)
}
