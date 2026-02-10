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
			lserr.WithErrorLoadBalancerExceedQuota(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ResizeLoadBalancer(popts IResizeLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError) {
	url := resizeLoadBalancerUrl(s.VLBClient, popts)
	resp := new(ResizeLoadBalancerResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp))
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ListLoadBalancerPackages(popts IListLoadBalancerPackagesRequest) (*lsentity.ListLoadBalancerPackages, lserr.IError) {
	url := listLoadBalancerPackagesUrl(s.VLBClient, popts)
	resp := new(ListLoadBalancerPackagesResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListLoadBalancerPackages(), nil
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
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			AppendCategories(lserr.ErrCatProductVlb)
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
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityListLoadBalancers(), nil
}

func (s *LoadBalancerServiceV2) GetPoolHealthMonitorById(popts IGetPoolHealthMonitorByIdRequest) (*lsentity.HealthMonitor, lserr.IError) {
	url := getPoolHealthMonitorByIdUrl(s.VLBClient, popts)
	resp := new(GetPoolHealthMonitorByIdResponse)
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

	return resp.ToEntityHealthMonitor(), nil
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
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
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
			lserr.WithErrorListenerNotFound(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
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
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
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
			lserr.WithErrorListenerNotFound(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
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
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			AppendCategories(lserr.ErrCatProductVlb)
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
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			AppendCategories(lserr.ErrCatProductVlb)
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
			lserr.WithErrorMemberMustIdentical(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
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
			lserr.WithErrorPoolNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId(), "poolId", popts.GetPoolId()).
			AppendCategories(lserr.ErrCatProductVlb)
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
			lserr.WithErrorPoolNotFound(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
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
			lserr.WithErrorLoadBalancerNotReady(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
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
			lserr.WithErrorLoadBalancerIsCreating(errResp),
			lserr.WithErrorLoadBalancerIsDeleting(errResp)).
			WithKVparameters(
				"loadBalancerId", popts.GetLoadBalancerId(),
				"projectId", s.getProjectId()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) GetPoolById(popts IGetPoolByIdRequest) (*lsentity.Pool, lserr.IError) {
	url := getPoolByIdUrl(s.VLBClient, popts)
	resp := new(GetPoolByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorPoolNotFound(errResp)).
			WithKVparameters(
				"loadBalancerId", popts.GetLoadBalancerId(),
				"poolId", popts.GetPoolId()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *LoadBalancerServiceV2) GetListenerById(popts IGetListenerByIdRequest) (*lsentity.Listener, lserr.IError) {
	url := getListenerByIdUrl(s.VLBClient, popts)
	resp := new(GetListenerByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorListenerNotFound(errResp)).
			WithKVparameters(
				"loadBalancerId", popts.GetLoadBalancerId(),
				"listenerId", popts.GetListenerId()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityListener(), nil
}

func (s *LoadBalancerServiceV2) ResizeLoadBalancerById(popts IResizeLoadBalancerByIdRequest) lserr.IError {
	url := resizeLoadBalancerByIdUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerPackageNotFound(errResp),
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorLoadBalancerNotFound2(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp),
			lserr.WithErrorLoadBalancerResizeSamePackage(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ScaleLoadBalancer(popts IScaleLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError) {
	url := scaleLoadBalancerUrl(s.VLBClient, popts)
	resp := new(ScaleLoadBalancerResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorLoadBalancerNotFound2(errResp),
			lserr.WithErrorLoadBalancerNotReady(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

// policy

func (s *LoadBalancerServiceV2) ListPolicies(popts IListPoliciesRequest) (*lsentity.ListPolicies, lserr.IError) {
	url := listPoliciesUrl(s.VLBClient, popts)
	resp := new(ListPoliciesResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListPolicies(), nil
}

func (s *LoadBalancerServiceV2) CreatePolicy(popts ICreatePolicyRequest) (*lsentity.Policy, lserr.IError) {
	url := createPolicyUrl(s.VLBClient, popts)
	resp := new(CreatePolicyResponse)
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
			lserr.WithErrorListenerNotFound(errResp),
		).WithParameters(popts.ToMap())
	}

	return resp.ToEntityPolicy(), nil
}

func (s *LoadBalancerServiceV2) GetPolicyById(popts IGetPolicyByIdRequest) (*lsentity.Policy, lserr.IError) {
	url := getPolicyByIdUrl(s.VLBClient, popts)
	resp := new(GetPolicyResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorListenerNotFound(errResp),
		).WithKVparameters("policyId", popts.GetPolicyId())
	}

	return resp.ToEntityPolicy(), nil
}

func (s *LoadBalancerServiceV2) UpdatePolicy(popts IUpdatePolicyRequest) lserr.IError {
	url := updatePolicyUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorListenerNotFound(errResp),
		)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeletePolicyById(popts IDeletePolicyByIdRequest) lserr.IError {
	url := deletePolicyByIdUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorListenerNotFound(errResp),
		)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ReorderPolicies(popts IReorderPoliciesRequest) lserr.IError {
	url := reorderPoliciesUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)
	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorLoadBalancerNotFound(errResp),
			lserr.WithErrorListenerNotFound(errResp),
		)
	}
	return nil
}

// --------------------------------------------------------

func (s *LoadBalancerServiceV2) ListCertificates(popts IListCertificatesRequest) (*lsentity.ListCertificates, lserr.IError) {
	url := listCertificatesUrl(s.VLBClient)
	resp := new(ListCertificatesResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListCertificates(), nil
}

func (s *LoadBalancerServiceV2) GetCertificateById(popts IGetCertificateByIdRequest) (*lsentity.Certificate, lserr.IError) {
	url := getCertificateByIdUrl(s.VLBClient, popts)
	resp := new(GetCertificateByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityCertificate(), nil
}

func (s *LoadBalancerServiceV2) CreateCertificate(popts ICreateCertificateRequest) (*lsentity.Certificate, lserr.IError) {
	url := createCertificateUrl(s.VLBClient)
	resp := new(CreateCertificateResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityCertificate(), nil
}

func (s *LoadBalancerServiceV2) DeleteCertificateById(popts IDeleteCertificateByIdRequest) lserr.IError {
	url := deleteCertificateByIdUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return nil
}
