package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *LoadBalancerServiceV2) CreateLoadBalancer(popts ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError) {
	url := createLoadBalancerUrl(s.VLBClient)
	resp := new(CreateLoadBalancerResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerExceedQuota(errResp))
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) GetLoadBalancerById(popts IGetLoadBalancerByIdRequest) (*lsentity.LoadBalancer, lserr.IError) {
	url := getLoadBalancerByIdUrl(s.VLBClient, popts)
	resp := new(GetLoadBalancerByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId())
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ListLoadBalancers(popts IListLoadBalancersRequest) (*lsentity.ListLoadBalancers, lserr.IError) {
	url := listLoadBalancersUrl(s.VLBClient, popts)
	resp := new(ListLoadBalancersResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListLoadBalancers(), nil
}

func (s *LoadBalancerServiceV2) CreatePool(popts ICreatePoolRequest) (*lsentity.Pool, lserr.IError) {
	url := createPoolUrl(s.VLBClient, popts)
	resp := new(CreatePoolResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorLoadBalancerNotFound2(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp),
			lserr.WithErrorLoadBalancerDuplicatePoolName(errResp)).
			WithParameters(popts.ToMap())
	}

	return resp.ToEntityPool(), nil
}

func (s *LoadBalancerServiceV2) UpdatePool(popts IUpdatePoolRequest) lserr.IError {
	url := updatePoolUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound2(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp),
			lserr.WithErrorListenerNotFound(errResp))
	}
	return nil
}

func (s *LoadBalancerServiceV2) CreateListener(popts ICreateListenerRequest) (*lsentity.Listener, lserr.IError) {
	url := createListenerUrl(s.VLBClient, popts)
	resp := new(CreateListenerResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound2(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp),
			lserr.WithErrorListenerDuplicateName(errResp),
			lserr.WithErrorPoolNotFound(errResp),
			lserr.WithErrorListenerDuplicateProtocolOrPort(errResp)).
			WithParameters(popts.ToMap())
	}

	return resp.ToEntityListener(), nil
}

func (s *LoadBalancerServiceV2) UpdateListener(popts IUpdateListenerRequest) lserr.IError {
	url := updateListenerUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound2(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp),
			lserr.WithErrorListenerNotFound(errResp))
	}

	return nil
}

func (s *LoadBalancerServiceV2) ListListenersByLoadBalancerId(popts IListListenersByLoadBalancerIdRequest) (*lsentity.ListListeners, lserr.IError) {
	url := listListenersByLoadBalancerIdUrl(s.VLBClient, popts)
	resp := new(ListListenersByLoadBalancerIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId())
	}

	return resp.ToEntityListListeners(), nil
}

func (s *LoadBalancerServiceV2) ListPoolsByLoadBalancerId(popts IListPoolsByLoadBalancerIdRequest) (*lsentity.ListPools, lserr.IError) {
	url := listPoolsByLoadBalancerIdUrl(s.VLBClient, popts)
	resp := new(ListPoolsByLoadBalancerIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId())
	}

	return resp.ToEntityListPools(), nil
}

func (s *LoadBalancerServiceV2) UpdatePoolMembers(popts IUpdatePoolMembersRequest) lserr.IError {
	url := updatePoolMembersUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound2(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp),
			lserr.WithErrorPoolNotFound(errResp),
			lserr.WithErrorMemberMustIdentical(errResp))
	}

	return nil
}

func (s *LoadBalancerServiceV2) ListPoolMembers(popts IListPoolMembersRequest) (*lsentity.ListMembers, lserr.IError) {
	url := listPoolMembersUrl(s.VLBClient, popts)
	resp := new(ListPoolMembersResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorPoolNotFound(errResp))
	}

	return resp.ToEntityListMembers(), nil
}

func (s *LoadBalancerServiceV2) DeletePoolById(popts IDeletePoolByIdRequest) lserr.IError {
	url := deletePoolByIdUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorPoolInUse(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp),
			lserr.WithErrorPoolNotFound(errResp))
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeleteListenerById(popts IDeleteListenerByIdRequest) lserr.IError {
	url := deleteListenerByIdUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound2(errResp),
			lserr.WithErrorListenerNotFound(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp))
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeleteLoadBalancerById(popts IDeleteLoadBalancerByIdRequest) lserr.IError {
	url := deleteLoadBalancerByIdUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp),
			lserr.WithErrorLoadBalancerIsDeleting(errResp)).
			WithKVparameters(
				"loadBalancerId", popts.GetLoadBalancerId(),
				"projectId", s.getProjectId())
	}

	return nil
}
