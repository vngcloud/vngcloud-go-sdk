package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV2) GetSecgroupById(popts IGetSecgroupByIdRequest) (*lsentity.Secgroup, lserr.IError) {
	url := getSecgroupByIdUrl(s.VserverClient, popts)
	resp := new(GetSecgroupByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", popts.GetSecgroupId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) CreateSecgroup(popts ICreateSecgroupRequest) (*lsentity.Secgroup, lserr.IError) {
	url := createSecgroupUrl(s.VserverClient)
	resp := new(CreateSecgroupResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSecgroupNameAlreadyExists(errResp),
			lserr.WithErrorSecgroupRuleExceedQuota(errResp),
			lserr.WithErrorSecgroupExceedQuota(errResp)).
			WithKVparameters(
				"secgroupName", popts.GetSecgroupName(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) ListSecgroup(popts IListSecgroupRequest) (*lsentity.ListSecgroups, lserr.IError) {
	url := listSecgroupUrl(s.VserverClient, popts)
	resp := new(ListSecgroupResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToListEntitySecgroups(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupById(popts IDeleteSecgroupByIdRequest) lserr.IError {
	url := deleteSecgroupByIdUrl(s.VserverClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSecgroupInUse(errResp),
			lserr.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", popts.GetSecgroupId(),
				"projectId", s.getProjectId())
	}

	return nil
}
