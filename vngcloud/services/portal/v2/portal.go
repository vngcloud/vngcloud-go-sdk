package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *PortalServiceV2) ListAllQuotaUsed() (*lsentity.ListQuotas, lserr.IError) {
	url := listAllQuotaUsedUrl(s.PortalClient)
	resp := new(ListAllQuotaUsedResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.PortalClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityListQuotas(), nil
}

func (s *PortalServiceV2) GetQuotaByName(popts IGetQuotaByNameRequest) (*lsentity.Quota, lserr.IError) {
	listQuotas, sdkErr := s.ListAllQuotaUsed()
	if sdkErr != nil {
		return nil, sdkErr
	}

	quota := listQuotas.FindQuotaByName(string(popts.GetName()))
	if quota == nil {
		return nil, lserr.ErrorHandler(nil, lserr.WithErrorQuotaNotFound(nil)).WithKVparameters("quotaName", popts.GetName())
	}

	return quota, nil
}

func (s *PortalServiceV2) getProjectId() string {
	return s.PortalClient.GetProjectId()
}
