package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *ComputeServiceV2) CreateServer(popts ICreateServerRequest) (*lsentity.Server, lserr.ISdkError) {
	url := createServerUrl(s.VServerClient)
	resp := new(CreateServerResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorOutOfPoc(errResp),
			lserr.WithErrorSubnetNotFound(errResp),
			lserr.WithErrorServerExceedQuota(errResp),
			lserr.WithErrorVolumeTypeNotFound(errResp),
			lserr.WithErrorNetworkNotFound(errResp),
			lserr.WithErrorServerCreateBillingPaymentMethodNotAllowed(errResp)).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) GetServerById(popts IGetServerByIdRequest) (*lsentity.Server, lserr.ISdkError) {
	url := getServerByIdUrl(s.VServerClient, popts)
	resp := new(GetServerByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorServerNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"serverId", popts.GetServerId())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) DeleteServerById(popts IDeleteServerByIdRequest) lserr.ISdkError {
	url := deleteServerByIdUrl(s.VServerClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorServerNotFound(errResp),
			lserr.WithErrorServerDeleteDeletingServer(errResp),
			lserr.WithErrorServerDeleteBillingServer(errResp),
			lserr.WithErrorServerDeleteCreatingServer(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"serverId", popts.GetServerId())
	}

	return nil
}

func (s *ComputeServiceV2) UpdateServerSecgroupsByServerId(popts IUpdateServerSecgroupsByServerIdRequest) (*lsentity.Server, lserr.ISdkError) {
	url := updateServerSecgroupsByServerIdUrl(s.VServerClient, popts)
	resp := new(UpdateServerSecgroupsByServerIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorServerNotFound(errResp),
			lserr.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"serverId", popts.GetServerId(),
				"secgroupIds", popts.GetListSecgroupsIds())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) AttachBlockVolume(popts IAttachBlockVolumeRequest) lserr.ISdkError {
	url := attachBlockVolumeUrl(s.VServerClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(map[string]interface{}{}).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeNotFound(errResp),
			lserr.WithErrorServerNotFound(errResp),
			lserr.WithErrorVolumeAvailable(errResp),
			lserr.WithErrorVolumeInProcess(errResp),
			lserr.WithErrorVolumeAlreadyAttached(errResp),
			lserr.WithErrorServerAttachEncryptedVolume(errResp),
			lserr.WithErrorVolumeAlreadyAttachedThisServer(errResp),
			lserr.WithErrorServerAttachVolumeQuotaExceeded(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId(),
				"serverId", popts.GetServerId())
	}

	return nil
}

func (s *ComputeServiceV2) DetachBlockVolume(popts IDetachBlockVolumeRequest) lserr.ISdkError {
	url := detachBlockVolumeUrl(s.VServerClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(202).
		WithJsonBody(map[string]interface{}{}).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorVolumeNotFound(errResp),
			lserr.WithErrorVolumeInProcess(errResp),
			lserr.WithErrorVolumeAvailable(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId(),
				"serverId", popts.GetServerId())
	}

	return nil
}
