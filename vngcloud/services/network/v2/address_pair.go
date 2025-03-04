package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV2) GetAllAddressPairByVirtualSubnetId(popts IGetAllAddressPairByVirtualSubnetIdRequest) ([]*lsentity.AddressPair, lserr.IError) {
	url := getAllAddressPairByVirtualSubnetIdUrl(s.VserverClient, popts)
	resp := new(GetAllAddressPairByVirtualSubnetIdResponse)
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

func (s *NetworkServiceV2) DeleteAddressPair(popts IDeleteAddressPairRequest) lserr.IError {
	url := deleteAddressPairUrl(s.VserverClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200, 201, 202, 203, 204).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorAddressPairNotFound(errResp)).
			WithKVparameters("addressPairId", popts.GetAddressPairID())
	}
	return nil
}

func (s *NetworkServiceV2) CreateAddressPair(popts ICreateAddressPairRequest) (*lsentity.AddressPair, lserr.IError) {
	url := createAddressPairUrl(s.VserverClient, popts)
	resp := new(CreateAddressPairResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorInternalNetworkInterfaceNotFound(errResp),
			lserr.WithErrorAddressPairExisted(errResp)).
			WithErrorCategories(lserr.ErrCatVServer, lserr.ErrCatVirtualAddress)
	}

	return resp.ToAddressPair(), nil
}
