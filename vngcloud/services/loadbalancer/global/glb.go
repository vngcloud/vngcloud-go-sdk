package global

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

// func (s *GlobalLoadBalancerService) CreateLoadBalancer(popts ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError) {
// 	url := createLoadBalancerUrl(s.VLBClient)
// 	resp := new(CreateLoadBalancerResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerExceedQuota(errResp)).
// 			WithParameters(popts.ToMap()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return resp.ToEntityLoadBalancer(), nil
// }

// func (s *GlobalLoadBalancerService) ResizeLoadBalancer(popts IResizeLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError) {
// 	url := resizeLoadBalancerUrl(s.VLBClient, popts)
// 	resp := new(ResizeLoadBalancerResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp))
// 	}

// 	return resp.ToEntityLoadBalancer(), nil
// }

// func (s *GlobalLoadBalancerService) ListLoadBalancerPackages(popts IListLoadBalancerPackagesRequest) (*lsentity.ListLoadBalancerPackages, lserr.IError) {
// 	url := listLoadBalancerPackagesUrl(s.VLBClient)
// 	resp := new(ListLoadBalancerPackagesResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
// 	}

// 	return resp.ToEntityListLoadBalancerPackages(), nil
// }

// func (s *GlobalLoadBalancerService) GetLoadBalancerById(popts IGetLoadBalancerByIdRequest) (*lsentity.LoadBalancer, lserr.IError) {
// 	url := getLoadBalancerByIdUrl(s.VLBClient, popts)
// 	resp := new(GetLoadBalancerByIdResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp)).
// 			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return resp.ToEntityLoadBalancer(), nil
// }

func (s *LoadBalancerServiceGlobal) ListGlobalLoadBalancers(popts IListGlobalLoadBalancersRequest) (*lsentity.ListGlobalLoadBalancers, lserr.IError) {
	url := listGlobalLoadBalancersUrl(s.VLBClient, popts)
	resp := new(ListGlobalLoadBalancersResponse)
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

	return resp.ToEntityListGlobalLoadBalancers(), nil
}

// func (s *GlobalLoadBalancerService) GetPoolHealthMonitorById(popts IGetPoolHealthMonitorByIdRequest) (*lsentity.HealthMonitor, lserr.IError) {
// 	url := getPoolHealthMonitorByIdUrl(s.VLBClient, popts)
// 	resp := new(GetPoolHealthMonitorByIdResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorPoolNotFound(errResp))
// 	}

// 	return resp.ToEntityHealthMonitor(), nil
// }

// func (s *GlobalLoadBalancerService) CreatePool(popts ICreatePoolRequest) (*lsentity.Pool, lserr.IError) {
// 	url := createPoolUrl(s.VLBClient, popts)
// 	resp := new(CreatePoolResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorLoadBalancerNotFound2(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp),
// 			lserr.WithErrorLoadBalancerDuplicatePoolName(errResp)).
// 			WithParameters(popts.ToMap()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return resp.ToEntityPool(), nil
// }

// func (s *GlobalLoadBalancerService) UpdatePool(popts IUpdatePoolRequest) lserr.IError {
// 	url := updatePoolUrl(s.VLBClient, popts)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
// 		return lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound2(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp),
// 			lserr.WithErrorListenerNotFound(errResp)).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}
// 	return nil
// }

// func (s *GlobalLoadBalancerService) CreateListener(popts ICreateListenerRequest) (*lsentity.Listener, lserr.IError) {
// 	url := createListenerUrl(s.VLBClient, popts)
// 	resp := new(CreateListenerResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound2(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp),
// 			lserr.WithErrorListenerDuplicateName(errResp),
// 			lserr.WithErrorPoolNotFound(errResp),
// 			lserr.WithErrorListenerDuplicateProtocolOrPort(errResp)).
// 			WithParameters(popts.ToMap()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return resp.ToEntityListener(), nil
// }

// func (s *GlobalLoadBalancerService) UpdateListener(popts IUpdateListenerRequest) lserr.IError {
// 	url := updateListenerUrl(s.VLBClient, popts)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
// 		return lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound2(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp),
// 			lserr.WithErrorListenerNotFound(errResp)).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return nil
// }

// func (s *GlobalLoadBalancerService) ListListenersByLoadBalancerId(popts IListListenersByLoadBalancerIdRequest) (*lsentity.ListListeners, lserr.IError) {
// 	url := listListenersByLoadBalancerIdUrl(s.VLBClient, popts)
// 	resp := new(ListListenersByLoadBalancerIdResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp)).
// 			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return resp.ToEntityListListeners(), nil
// }

// func (s *GlobalLoadBalancerService) ListPoolsByLoadBalancerId(popts IListPoolsByLoadBalancerIdRequest) (*lsentity.ListPools, lserr.IError) {
// 	url := listPoolsByLoadBalancerIdUrl(s.VLBClient, popts)
// 	resp := new(ListPoolsByLoadBalancerIdResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp)).
// 			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return resp.ToEntityListPools(), nil
// }

// func (s *GlobalLoadBalancerService) UpdatePoolMembers(popts IUpdatePoolMembersRequest) lserr.IError {
// 	url := updatePoolMembersUrl(s.VLBClient, popts)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
// 		return lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound2(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp),
// 			lserr.WithErrorPoolNotFound(errResp),
// 			lserr.WithErrorMemberMustIdentical(errResp)).AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return nil
// }

// func (s *GlobalLoadBalancerService) ListPoolMembers(popts IListPoolMembersRequest) (*lsentity.ListMembers, lserr.IError) {
// 	url := listPoolMembersUrl(s.VLBClient, popts)
// 	resp := new(ListPoolMembersResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorPoolNotFound(errResp)).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return resp.ToEntityListMembers(), nil
// }

// func (s *GlobalLoadBalancerService) DeletePoolById(popts IDeletePoolByIdRequest) lserr.IError {
// 	url := deletePoolByIdUrl(s.VLBClient, popts)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
// 		return lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorPoolInUse(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp),
// 			lserr.WithErrorPoolNotFound(errResp)).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return nil
// }

// func (s *GlobalLoadBalancerService) DeleteListenerById(popts IDeleteListenerByIdRequest) lserr.IError {
// 	url := deleteListenerByIdUrl(s.VLBClient, popts)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
// 		return lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound2(errResp),
// 			lserr.WithErrorListenerNotFound(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp)).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return nil
// }

// func (s *GlobalLoadBalancerService) DeleteLoadBalancerById(popts IDeleteLoadBalancerByIdRequest) lserr.IError {
// 	url := deleteLoadBalancerByIdUrl(s.VLBClient, popts)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
// 		return lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp),
// 			lserr.WithErrorLoadBalancerIsCreating(errResp),
// 			lserr.WithErrorLoadBalancerIsDeleting(errResp)).
// 			WithKVparameters(
// 				"loadBalancerId", popts.GetLoadBalancerId(),
// 				"projectId", s.getProjectId()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return nil
// }

// func (s *GlobalLoadBalancerService) GetPoolById(popts IGetPoolByIdRequest) (*lsentity.Pool, lserr.IError) {
// 	url := getPoolByIdUrl(s.VLBClient, popts)
// 	resp := new(GetPoolByIdResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorPoolNotFound(errResp)).
// 			WithKVparameters(
// 				"loadBalancerId", popts.GetLoadBalancerId(),
// 				"poolId", popts.GetPoolId()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return resp.ToEntityPool(), nil
// }

// func (s *GlobalLoadBalancerService) GetListenerById(popts IGetListenerByIdRequest) (*lsentity.Listener, lserr.IError) {
// 	url := getListenerByIdUrl(s.VLBClient, popts)
// 	resp := new(GetListenerByIdResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorListenerNotFound(errResp)).
// 			WithKVparameters(
// 				"loadBalancerId", popts.GetLoadBalancerId(),
// 				"listenerId", popts.GetListenerId()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return resp.ToEntityListener(), nil
// }

// func (s *GlobalLoadBalancerService) ResizeLoadBalancerById(popts IResizeLoadBalancerByIdRequest) lserr.IError {
// 	url := resizeLoadBalancerByIdUrl(s.VLBClient, popts)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
// 		return lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerPackageNotFound(errResp),
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorLoadBalancerNotFound2(errResp),
// 			lserr.WithErrorLoadBalancerNotReady(errResp),
// 			lserr.WithErrorLoadBalancerResizeSamePackage(errResp)).
// 			WithParameters(popts.ToMap()).
// 			AppendCategories(lserr.ErrCatProductVlb)
// 	}

// 	return nil
// }

// // policy

// func (s *GlobalLoadBalancerService) ListPolicies(popts IListPoliciesRequest) (*lsentity.ListPolicies, lserr.IError) {
// 	url := listPoliciesUrl(s.VLBClient, popts)
// 	resp := new(ListPoliciesResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
// 	}

// 	return resp.ToEntityListPolicies(), nil
// }

// func (s *GlobalLoadBalancerService) CreatePolicy(popts ICreatePolicyRequest) (*lsentity.Policy, lserr.IError) {
// 	url := createPolicyUrl(s.VLBClient, popts)
// 	resp := new(CreatePolicyResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorListenerNotFound(errResp),
// 		).WithParameters(popts.ToMap())
// 	}

// 	return resp.ToEntityPolicy(), nil
// }

// func (s *GlobalLoadBalancerService) GetPolicyById(popts IGetPolicyByIdRequest) (*lsentity.Policy, lserr.IError) {
// 	url := getPolicyByIdUrl(s.VLBClient, popts)
// 	resp := new(GetPolicyResponse)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(200).
// 		WithJsonResponse(resp).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
// 		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorListenerNotFound(errResp),
// 		).WithKVparameters("policyId", popts.GetPolicyId())
// 	}

// 	return resp.ToEntityPolicy(), nil
// }

// func (s *GlobalLoadBalancerService) UpdatePolicy(popts IUpdatePolicyRequest) lserr.IError {
// 	url := updatePolicyUrl(s.VLBClient, popts)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonBody(popts.ToRequestBody()).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
// 		return lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorListenerNotFound(errResp),
// 		)
// 	}

// 	return nil
// }

// func (s *GlobalLoadBalancerService) DeletePolicyById(popts IDeletePolicyByIdRequest) lserr.IError {
// 	url := deletePolicyByIdUrl(s.VLBClient, popts)
// 	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
// 	req := lsclient.NewRequest().
// 		WithHeader("User-Agent", popts.ParseUserAgent()).
// 		WithOkCodes(202).
// 		WithJsonError(errResp)

// 	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
// 		return lserr.SdkErrorHandler(sdkErr, errResp,
// 			lserr.WithErrorLoadBalancerNotFound(errResp),
// 			lserr.WithErrorListenerNotFound(errResp),
// 		)
// 	}

// 	return nil
// }
