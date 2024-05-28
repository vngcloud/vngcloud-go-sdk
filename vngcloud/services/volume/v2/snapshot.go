package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VolumeServiceV2) ListSnapshotByVolumeId(popts IListSnapshotsByBlockVolumeIdRequest) (*lsentity.ListSnapshots, lserr.ISdkError) {
	url := listSnapshotsByBlockVolumeIdUrl(s.VServerClient, popts)
	resp := new(ListSnapshotByBlockVolumeIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId())
	}

	return resp.ToEntityListSnapshots(), nil
}

func (s *VolumeServiceV2) CreateSnapshotByVolumeId(popts ICreateSnapshotByVolumeIdRequest) (*lsentity.Snapshot, lserr.ISdkError) {
	url := createSnapshotByBlockVolumeIdUrl(s.VServerClient, popts)
	resp := new(CreateSnapshotByBlockVolumeIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp).
		WithJsonBody(popts.ToRequestBody())

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeNotFound(errResp),
			lserr.WithErrorSnapshotNameNotValid(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId())
	}

	return resp.ToEntitySnapshot(), nil
}

func (s *VolumeServiceV2) DeleteSnapshotById(popts IDeleteSnapshotByIdRequest) lserr.ISdkError {
	url := deleteSnapshotByIdUrl(s.VServerClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSnapshotNameNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"snapshotId", popts.GetSnapshotId())
	}

	return nil
}
