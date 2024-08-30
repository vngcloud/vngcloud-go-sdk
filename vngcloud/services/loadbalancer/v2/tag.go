package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *LoadBalancerServiceV2) ListTags(popts IListTagsRequest) (*lsentity.ListTags, lserr.IError) {
	url := listTagsUrl(s.VServerClient, popts)
	resp := new(ListTagsResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListTags(), nil
}

func (s *LoadBalancerServiceV2) CreateTags(popts ICreateTagsRequest) lserr.IError {
	url := createTagsUrl(s.VServerClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp)
	}

	return nil
}

func (s *LoadBalancerServiceV2) UpdateTags(popts IUpdateTagsRequest) lserr.IError {
	tmpTags, sdkErr := s.ListTags(NewListTagsRequest(popts.GetLoadBalancerId()))
	if sdkErr != nil {
		return sdkErr
	}

	// Do not update system tags
	tags := new(lsentity.ListTags)
	for _, tag := range tmpTags.Items {
		if !tag.SystemTag {
			tags.Items = append(tags.Items, tag)
		}
	}

	url := updateTagsUrl(s.VServerClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody(tags)).
		WithJsonError(errResp)

	if _, sdkErr = s.VServerClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorTagKeyInvalid(errResp)).WithParameters(popts.ToMap())
	}

	return nil

}
