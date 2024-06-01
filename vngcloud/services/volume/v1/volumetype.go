package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VolumeServiceV1) GetVolumeTypeById(popts IGetVolumeTypeByIdRequest) (*lsentity.VolumeType, lserr.ISdkError) {
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
