package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV2) CreateVirtualAddressCrossProject(popts ICreateVirtualAddressCrossProjectRequest) (*lsentity.VirtualAddress, lserr.IError) {
	url := createVirtualAddressCrossProjectUrl(s.VserverClient)
	resp := new(CreateVirtualAddressCrossProjectResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSubnetNotFound(errResp),
			lserr.WithErrorVirtualAddressExceedQuota(errResp)).
			WithKVparameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatVServer, lserr.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) DeleteVirtualAddressById(popts IDeleteVirtualAddressByIdRequest) lserr.IError {
	url := deleteVirtualAddressByIdUrl(s.VserverClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVirtualAddressNotFound(errResp),
			lserr.WithErrorVirtualAddressInUse(errResp)).
			WithKVparameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatVServer, lserr.ErrCatVirtualAddress)
	}

	return nil
}

func (s *NetworkServiceV2) GetVirtualAddressById(popts IGetVirtualAddressByIdRequest) (*lsentity.VirtualAddress, lserr.IError) {
	url := getVirtualAddressByIdUrl(s.VserverClient, popts)
	resp := new(GetVirtualAddressByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithJsonResponse(resp).
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVirtualAddressNotFound(errResp)).
			WithKVparameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatVServer, lserr.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) ListAddressPairsByVirtualAddressId(popts IListAddressPairsByVirtualAddressIdRequest) (*lsentity.ListAddressPairs, lserr.IError) {
	url := listAddressPairsByVirtualAddressIdUrl(s.VserverClient, popts)
	resp := new(ListAddressPairsByVirtualAddressIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithJsonResponse(resp).
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVirtualAddressNotFound(errResp)).
			WithKVparameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatVServer, lserr.ErrCatVirtualAddress)
	}

	return resp.ToEntityListAddressPairs(), nil
}
