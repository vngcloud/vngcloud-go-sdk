package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VolumeServiceV2) CreateBlockVolume(popts ICreateBlockVolumeRequest) (*lsentity.Volume, lserr.ISdkError) {
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
			lserr.WithErrorVolumeNameNotValid(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.ToMap())
	}

	return resp.ToEntityVolume(), nil
}
