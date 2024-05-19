package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *ComputeServiceV2) CreateServer(popts ICreateServerRequest) (*lsentity.Server, lserr.ISdkError) {
	url := createServerUrl(s.VserverClient)
	resp := new(CreateServerResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorOutOfPoc(errResp),
			lserr.WithErrorSubnetNotFound(errResp),
			lserr.WithErrorVolumeTypeNotFound(errResp),
			lserr.WithErrorNetworkNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityServer(), nil
}
