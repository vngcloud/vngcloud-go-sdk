package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VolumeServiceV2) CreateBlockVolume(popts ICreateBlockVolumeRequest) (*lsentity.Volume, lserr.IError) {
	url := createBlockVolumeUrl(s.VServerClient)
	resp := new(CreateBlockVolumeResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeTypeNotFound(errResp),
			lserr.WithErrorVolumeSizeOutOfRange(errResp),
			lserr.WithErrorVolumeSizeExceedGlobalQuota(errResp),
			lserr.WithErrorVolumeNameNotValid(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.ToMap())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) DeleteBlockVolumeById(popts IDeleteBlockVolumeByIdRequest) lserr.IError {
	url := deleteBlockVolumeByIdUrl(s.VServerClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId())
	}

	return nil
}

func (s *VolumeServiceV2) ListBlockVolumes(popts IListBlockVolumesRequest) (*lsentity.ListVolumes, lserr.IError) {
	url := listBlockVolumesUrl(s.VServerClient, popts)
	resp := new(ListBlockVolumesResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorPagingInvalid(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.ToMap())
	}

	return resp.ToEntityListVolumes(), nil
}

func (s *VolumeServiceV2) GetBlockVolumeById(popts IGetBlockVolumeByIdRequest) (*lsentity.Volume, lserr.IError) {
	url := getBlockVolumeByIdUrl(s.VServerClient, popts)
	resp := new(GetBlockVolumeByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) ResizeBlockVolumeById(popts IResizeBlockVolumeByIdRequest) (*lsentity.Volume, lserr.IError) {
	url := resizeBlockVolumeByIdUrl(s.VServerClient, popts)
	resp := new(ResizeBlockVolumeByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeNotFound(errResp),
			lserr.WithErrorVolumeSizeOutOfRange(errResp),
			lserr.WithErrorVolumeMustSameZone(errResp),
			lserr.WithErrorVolumeUnchanged(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId(),
				"size", popts.GetSize())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) GetUnderBlockVolumeId(popts IGetUnderBlockVolumeIdRequest) (*lsentity.Volume, lserr.IError) {
	url := getUnderBlockVolumeIdUrl(s.VServerClient, popts)
	resp := new(GetUnderBlockVolumeIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) MigrateBlockVolumeById(popts IMigrateBlockVolumeByIdRequest) lserr.IError {
	url := migrateBlockVolumeByIdUrl(s.VServerClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	resp := map[string]interface{}{}
	req := lsclient.NewRequest().
		WithOkCodes(204).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp).
		WithJsonResponse(&resp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		sdkErr = lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeMigrateInSameZone(errResp),
			lserr.WithErrorVolumeMigrateMissingInit(errResp),
			lserr.WithErrorVolumeMigrateNeedProcess(errResp),
			lserr.WithErrorVolumeMigrateNeedConfirm(errResp),
			lserr.WithErrorVolumeMigrateBeingProcess(errResp),
			lserr.WithErrorVolumeMigrateProcessingConfirm(errResp),
			lserr.WithErrorVolumeMigrateBeingMigrating(errResp), // should under WithErrorVolumeMigrateBeingProcess
			lserr.WithErrorVolumeMigrateBeingFinish(errResp),
			lserr.WithErrorVolumeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId())

		if popts.IsConfirm() {
			switch sdkErr.GetErrorCode() {
			case lserr.EcVServerVolumeMigrateMissingInit:
				popts = popts.WithAction(InitMigrateAction)
				return s.MigrateBlockVolumeById(popts)
			case lserr.EcVServerVolumeMigrateNeedProcess:
				popts = popts.WithAction(ProcessMigrateAction)
				return s.MigrateBlockVolumeById(popts)
			case lserr.EcVServerVolumeMigrateNeedConfirm:
				popts = popts.WithAction(ConfirmMigrateAction)
				return s.MigrateBlockVolumeById(popts)
			}
		}

		return sdkErr
	}

	return nil
}
