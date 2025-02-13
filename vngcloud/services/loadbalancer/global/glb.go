package global

import (
	"fmt"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *LoadBalancerServiceGlobal) ListGlobalPools(popts IListGlobalPoolsRequest) (*lsentity.ListGlobalPools, lserr.IError) {
	url := listGlobalPoolsUrl(s.VLBClient, popts)
	resp := new(ListGlobalPoolsResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPools(), nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) CreateGlobalPool(popts ICreateGlobalPoolRequest) (*lsentity.GlobalPool, lserr.IError) {
	url := createGlobalPoolUrl(s.VLBClient, popts)
	resp := new(CreateGlobalPoolResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		fmt.Println("sdkErr: ", sdkErr)
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) UpdateGlobalPool(popts IUpdateGlobalPoolRequest) (*lsentity.GlobalPool, lserr.IError) {
	url := updateGlobalPoolUrl(s.VLBClient, popts)
	resp := new(UpdateGlobalPoolResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) DeleteGlobalPool(popts IDeleteGlobalPoolRequest) lserr.IError {
	url := deleteGlobalPoolUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) ListGlobalPoolMembers(popts IListGlobalPoolMembersRequest) (*lsentity.ListGlobalPoolMembers, lserr.IError) {
	url := listGlobalPoolMembersUrl(s.VLBClient, popts)
	resp := new(ListGlobalPoolMembersResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPoolMembers(), nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) PatchGlobalPoolMember(popts IPatchGlobalPoolMemberRequest) lserr.IError {
	url := patchGlobalPoolMemberUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Patch(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) ListGlobalListeners(popts IListGlobalListenersRequest) (*lsentity.ListGlobalListeners, lserr.IError) {
	url := listGlobalListenersUrl(s.VLBClient, popts)
	resp := new(ListGlobalListenersResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalListeners(), nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) CreateGlobalListener(popts ICreateGlobalListenerRequest) (*lsentity.GlobalListener, lserr.IError) {
	url := createGlobalListenerUrl(s.VLBClient, popts)
	resp := new(CreateGlobalListenerResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalListener(), nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) UpdateGlobalListener(popts IUpdateGlobalListenerRequest) (*lsentity.GlobalListener, lserr.IError) {
	url := updateGlobalListenerUrl(s.VLBClient, popts)
	resp := new(UpdateGlobalListenerResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)

	}

	return resp.ToEntityGlobalListener(), nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) DeleteGlobalListener(popts IDeleteGlobalListenerRequest) lserr.IError {
	url := deleteGlobalListenerUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) ListGlobalLoadBalancers(popts IListGlobalLoadBalancersRequest) (*lsentity.ListGlobalLoadBalancers, lserr.IError) {
	url := listGlobalLoadBalancersUrl(s.VLBClient, popts)
	resp := new(ListGlobalLoadBalancersResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
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

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) CreateGlobalLoadBalancer(popts ICreateGlobalLoadBalancerRequest) (*lsentity.GlobalLoadBalancer, lserr.IError) {
	url := createGlobalLoadBalancerUrl(s.VLBClient, popts)
	resp := new(CreateGlobalLoadBalancerResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) DeleteGlobalLoadBalancer(popts IDeleteGlobalLoadBalancerRequest) lserr.IError {
	url := deleteGlobalLoadBalancerUrl(s.VLBClient, popts)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *LoadBalancerServiceGlobal) GetGlobalLoadBalancerById(popts IGetGlobalLoadBalancerByIdRequest) (*lsentity.GlobalLoadBalancer, lserr.IError) {
	url := getGlobalLoadBalancerByIdUrl(s.VLBClient, popts)
	resp := new(GetGlobalLoadBalancerByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.GlobalLoadBalancerErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(lserr.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}
