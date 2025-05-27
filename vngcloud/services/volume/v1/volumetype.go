package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VolumeServiceV1) GetVolumeTypeById(popts IGetVolumeTypeByIdRequest) (*lsentity.VolumeType, lserr.IError) {
	url := getVolumeTypeByIdUrl(s.VServerClient, popts)
	resp := new(GetVolumeTypeByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeTypeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeTypeId", popts.GetVolumeTypeId())
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetDefaultVolumeType() (*lsentity.VolumeType, lserr.IError) {
	url := getDefaultVolumeTypeUrl(s.VServerClient)
	resp := new(GetDefaultVolumeTypeResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.getProjectId())
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetVolumeTypeZones(popts IGetVolumeTypeZonesRequest) (*lsentity.ListVolumeTypeZones, lserr.IError) {
	url := getVolumeTypeZonesUrl(s.VServerClient, popts)
	resp := new(ListVolumeTypeZonesResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeTypeNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityListVolumeTypeZones(), nil

}

func (s *VolumeServiceV1) GetListVolumeTypes(popts IGetListVolumeTypeRequest) (*lsentity.ListVolumeType, lserr.IError) {
	url := getVolumeTypesUrl(s.VServerClient, popts)
	resp := new(ListVolumeTypeResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeTypeNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"volumeTypeZoneId", popts.GetVolumeTypeZoneId())
	}

	return resp.ToEntityListVolumeType(), nil

}
