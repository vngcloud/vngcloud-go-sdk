package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *ComputeServiceV2) CreateServer(popts ICreateServerRequest) (*lsentity.Server, lserr.ISdkError) {
	url := createServerUrl(s.VserverClient)
	resp := new(CreateServerResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorOutOfPoc(errResp),
			lserr.WithErrorSubnetNotFound(errResp),
			lserr.WithErrorServerExceedQuota(errResp),
			lserr.WithErrorVolumeTypeNotFound(errResp),
			lserr.WithErrorNetworkNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) GetServerById(popts IGetServerByIdRequest) (*lsentity.Server, lserr.ISdkError) {
	url := getServerByIdUrl(s.VserverClient, popts)
	resp := new(GetServerByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorServerNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"serverId", popts.GetServerId())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) DeleteServerById(popts IDeleteServerByIdRequest) lserr.ISdkError {
	url := deleteServerByIdUrl(s.VserverClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorServerNotFound(errResp),
			lserr.WithErrorServerDeleteCreatingServer(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"serverId", popts.GetServerId())
	}

	return nil
}

func (s *ComputeServiceV2) UpdateServerSecgroupsByServerId(popts IUpdateServerSecgroupsByServerIdRequest) (*lsentity.Server, lserr.ISdkError) {
	url := updateServerSecgroupsByServerIdUrl(s.VserverClient, popts)
	resp := new(UpdateServerSecgroupsByServerIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Put(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorServerNotFound(errResp),
			lserr.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"serverId", popts.GetServerId(),
				"secgroupIds", popts.GetListSecgroupsIds())
	}

	return resp.ToEntityServer(), nil
}
