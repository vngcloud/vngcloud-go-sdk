package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VolumeServiceV2) ListSnapshotsByBlockVolumeId(popts IListSnapshotsByBlockVolumeIdRequest) (*lsentity.ListSnapshots, lserr.IError) {
	url := listSnapshotsByBlockVolumeIdUrl(s.VServerClient, popts)
	resp := new(ListSnapshotsByBlockVolumeIdResponse)
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

func (s *VolumeServiceV2) CreateSnapshotByBlockVolumeId(popts ICreateSnapshotByBlockVolumeIdRequest) (*lsentity.Snapshot, lserr.IError) {
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

func (s *VolumeServiceV2) DeleteSnapshotById(popts IDeleteSnapshotByIdRequest) lserr.IError {
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
