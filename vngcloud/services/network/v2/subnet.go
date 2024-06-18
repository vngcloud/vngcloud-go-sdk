package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV2) GetSubnetById(popts IGetSubnetByIdRequest) (*lsentity.Subnet, lserr.ISdkError) {
	url := getSubnetByIdUrl(s.VserverClient, popts)
	resp := new(GetSubnetByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSubnetNotBelongNetwork(sdkErr),
			lserr.WithErrorSubnetNotFound(errResp)).
			WithKVparameters(
				"subnetId", popts.GetSubnetId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntitySubnet(), nil
}
