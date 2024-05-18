package v2

import (
	"fmt"
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV2) CreateSecgroupRule(popts ICreateSecgroupRuleRequest) (*lsentity.SecgroupRule, lserr.ISdkError) {
	url := createSecgroupRuleUrl(s.VserverClient, popts)
	resp := new(CreateSecgroupRuleResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		fmt.Println("sdkErr: ", sdkErr)
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntitySecgroupRule(), nil
}
