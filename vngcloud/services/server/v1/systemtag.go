package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *ServerServiceInternalV1) CreateSystemTags(popts ICreateSystemTagRequest) (*[]lsentity.SystemTag, lserr.IError) {

	url := createSystemTagUrl(s.VServerClient)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)

	rawResp := new([]SystemTagResponse)

	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(rawResp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	result := make([]lsentity.SystemTag, 0, len(*rawResp))

	for _, r := range *rawResp {
		result = append(result, r.toSystemTag())
	}

	return &result, nil
}
