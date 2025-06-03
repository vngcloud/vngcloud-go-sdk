package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV2) GetNetworkById(popts IGetNetworkByIdRequest) (*lsentity.Network, lserr.IError) {
	url := getNetworkByIdUrl(s.VserverClient, popts)
	resp := new(GetNetworkByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorNetworkNotFound(errResp)).
			WithKVparameters(
				"networkId", popts.GetNetworkId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntityNetwork(), nil
}
