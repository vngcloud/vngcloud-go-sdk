package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV1) GetEndpointById(popts IGetEndpointByIdRequest) (*lsentity.Endpoint, lserr.ISdkError) {
	url := getEndpointByIdUrl(s.VNetworkClient, popts)
	resp := new(GetEndpointByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithUserId(s.getUserId()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"endpointId", popts.GetEndpointId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntityEndpoint(), nil
}
