package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV2) GetAllAddressPairsByVirtualSubnetId(popts IGetAllAddressPairByVirtualSubnetIdRequest) ([]*lsentity.AddressPair, lserr.IError) {
	url := getAllAddressPairsByVirtualSubnetIdUrl(s.VserverClient, popts)
	resp := new(GetAllAddressPairsByVirtualSubnetIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
		// lserr.WithErrorSubnetNotBelongNetwork(sdkErr),
		// lserr.WithErrorSubnetNotFound(errResp)).
		// WithKVparameters(
		// 	"virtualSubnetId", popts.GetVirtualSubnetId(),
		// 	"projectId", s.getProjectId(),

	}
	return resp.ToListAddressPair(), nil
}

func (s *NetworkServiceV2) SetAddressPairInVirtualSubnet(popts ISetAddressPairInVirtualSubnetRequest) (*lsentity.AddressPair, lserr.IError) {
	url := setAddressPairInVirtualSubnetUrl(s.VserverClient, popts)
	resp := new(SetAddressPairInVirtualSubnetResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200, 201, 202, 203, 204).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}
	return resp.ToAddressPair(), nil
}
