package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV2) GetSecgroupById(popts IGetSecgroupByIdRequest) (*lsentity.Secgroup, lserr.ISdkError) {
	url := getSecgroupUrl(s.VserverClient, popts)
	resp := new(GetSecgroupByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr,
			lserr.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("secgroupId", popts.GetSecgroupId())
	}

	return resp.ToEntitySecgroup(), nil
}
