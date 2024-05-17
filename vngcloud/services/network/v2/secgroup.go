package v2

import (
	"fmt"
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
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("secgroupId", popts.GetSecgroupId())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) CreateSecgroup(popts ICreateSecgroupRequest) (*lsentity.Secgroup, lserr.ISdkError) {
	url := createSecgroupUrl(s.VserverClient)
	resp := new(CreateSecgroupResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		fmt.Println("sdkErr: ", sdkErr)
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSecgroupNameAlreadyExists(errResp),
			lserr.WithErrorSecgroupExceedQuota(errResp)).
			WithKVparameters("secgroupName", popts.GetSecgroupName())
	}

	return resp.ToEntitySecgroup(), nil
}
