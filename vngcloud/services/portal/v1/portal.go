package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *PortalServiceV1) GetPortalInfo(popts IGetPortalInfoRequest) (*lsentity.Portal, lserr.IError) {
	url := getPortalInfoUrl(s.PortalClient, popts)
	resp := new(GetPortalInfoResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.PortalClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("backendProjectId", popts.GetBackEndProjectId())
	}

	return resp.ToEntityPortal(), nil
}

func (s *PortalServiceV1) ListProjects(popts IListProjectsRequest) (*lsentity.ListPortals, lserr.IError) {
	url := listProjectsUrl(s.PortalClient)
	resp := new(ListProjectsResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.PortalClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListPortals(), nil
}
