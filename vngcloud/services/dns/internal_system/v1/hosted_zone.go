package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VDnsServiceInternal) GetHostedZoneById(popts IGetHostedZoneByIdRequest, portalUserId string) (*lsentity.HostedZone, lserr.IError) {
	url := getHostedZoneByIdUrl(s.DnsClient, popts)
	resp := new(GetHostedZoneByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"hostedZoneId", popts.GetHostedZoneId()).
			WithErrorCategories(lserr.ErrCatProductVdns)
	}

	return resp.ToEntityHostedZone(), nil
}

func (s *VDnsServiceInternal) ListHostedZones(popts IListHostedZonesRequest, portalUserId string) (*lsentity.ListHostedZone, lserr.IError) {
	url := listHostedZonesUrl(s.DnsClient, popts)
	resp := new(ListHostedZonesResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatProductVdns)
	}

	return resp.ToEntityListHostedZones(), nil
}

func (s *VDnsServiceInternal) CreateHostedZone(popts ICreateHostedZoneRequest, portalUserId string) (*lsentity.HostedZone, lserr.IError) {
	url := createHostedZoneUrl(s.DnsClient)
	resp := new(CreateHostedZoneResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody(s.DnsClient)).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatProductVdns)
	}

	return resp.ToEntityHostedZone(), nil
}

func (s *VDnsServiceInternal) DeleteHostedZone(popts IDeleteHostedZoneRequest, portalUserId string) lserr.IError {
	url := deleteHostedZoneUrl(s.DnsClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceInternal) UpdateHostedZone(popts IUpdateHostedZoneRequest, portalUserId string) lserr.IError {
	url := updateHostedZoneUrl(s.DnsClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(portalUserId).
		WithOkCodes(204).
		WithJsonBody(popts.ToRequestBody(s.DnsClient)).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Put(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatProductVdns)
	}

	return nil
}
