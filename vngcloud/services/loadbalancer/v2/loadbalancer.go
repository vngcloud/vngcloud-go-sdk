package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *LoadBalancerServiceV2) CreateLoadBalancer(popts ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.ISdkError) {
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
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) GetLoadBalancerById(popts IGetLoadBalancerByIdRequest) (*lsentity.LoadBalancer, lserr.ISdkError) {
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

func (s *LoadBalancerServiceV2) ListLoadBalancers(popts IListLoadBalancersRequest) (*lsentity.ListLoadBalancers, lserr.ISdkError) {
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

func (s *LoadBalancerServiceV2) CreatePool(popts ICreatePoolRequest) (*lsentity.Pool, lserr.ISdkError) {
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
			lserr.WithErrorLoadBalancerNotFound2(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp),
			lserr.WithErrorLoadBalancerDuplicatePoolName(errResp)).
			WithParameters(popts.ToMap())
	}

	return resp.ToEntityPool(), nil
}

func (s *LoadBalancerServiceV2) CreateListener(popts ICreateListenerRequest) (*lsentity.Listener, lserr.ISdkError) {
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

func (s *LoadBalancerServiceV2) UpdateListener(popts IUpdateListenerRequest) lserr.ISdkError {
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

func (s *LoadBalancerServiceV2) ListListenersByLoadBalancerId(popts IListListenersByLoadBalancerIdRequest) (*lsentity.ListListeners, lserr.ISdkError) {
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

func (s *LoadBalancerServiceV2) ListPoolsByLoadBalancerId(popts IListPoolsByLoadBalancerIdRequest) (*lsentity.ListPools, lserr.ISdkError) {
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

func (s *LoadBalancerServiceV2) UpdatePoolMembers(popts IUpdatePoolMembersRequest) lserr.ISdkError {
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
