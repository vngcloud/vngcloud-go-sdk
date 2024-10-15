package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *NetworkServiceV1) GetEndpointById(popts IGetEndpointByIdRequest) (*lsentity.Endpoint, lserr.IError) {
	url := getEndpointByIdUrl(s.VNetworkClient, popts)
	resp := new(GetEndpointByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		//WithUserId(s.getUserId()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"endpointId", popts.GetEndpointId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) CreateEndpoint(popts ICreateEndpointRequest) (*lsentity.Endpoint, lserr.IError) {
	url := createEndpointUrl(s.VNetworkClient)
	resp := new(CreateEndpointResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(201).
		//WithUserId(s.getUserId()).
		WithJsonBody(popts.ToRequestBody(s.VNetworkClient)).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorEndpointOfVpcExists(errResp),
			lserr.WithErrorLockOnProcess(errResp),
			lserr.WithErrorNetworkNotFound(errResp),
			lserr.WithErrorSubnetNotFound(errResp),
			lserr.WithErrorEndpointPackageNotBelongToEndpointService(errResp),
			lserr.WithErrorContainInvalidCharacter(errResp)).
			WithParameters(popts.GetParameters())
	}

	return resp.ToEntityEndpoint(), nil
}

func (s *NetworkServiceV1) DeleteEndpointById(popts IDeleteEndpointByIdRequest) lserr.IError {
	url := deleteEndpointByIdUrl(s.VNetworkClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody(s.VNetworkClient)).
		//WithUserId(s.getUserId()).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorEndpointStatusInvalid(errResp),
			lserr.WithErrorNetworkNotFound(errResp),
			lserr.WithErrorSubnetNotFound(errResp)).
			WithParameters(popts.GetParameters())
	}

	return nil
}

func (s *NetworkServiceV1) ListEndpoints(popts IListEndpointsRequest) (*lsentity.ListEndpoints, lserr.IError) {
	url := listEndpointsUrl(s.VNetworkClient, popts)
	resp := new(ListEndpointsResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithOkCodes(200).
		//WithUserId(s.getUserId()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.GetParameters())
	}

	return resp.ToEntityListEndpoints(), nil
}

// ________________________________________________________________________ NetworkServiceInternalV1

func (s *NetworkServiceInternalV1) ListTagsByEndpointId(popts IListTagsByEndpointIdRequest) (*lsentity.ListTags, lserr.IError) {
	url := listTagsByEndpointIdUrl(s.VNetworkClient, popts)
	resp := new(ListTagsByEndpointIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithMapHeaders(popts.GetMapHeaders()).
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VNetworkClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.GetParameters())
	}

	return resp.ToEntityListTags(), nil
}
