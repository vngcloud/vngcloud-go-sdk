package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VDnsServiceV1) ListRecords(popts IListRecordsRequest) (*lsentity.ListDnsRecords, lserr.IError) {
	url := listRecordsUrl(s.DnsClient, popts)
	resp := new(ListRecordsResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatProductVdns)
	}

	return resp.ToEntityListRecords(), nil
}

func (s *VDnsServiceV1) GetRecord(popts IGetRecordRequest) (*lsentity.DnsRecord, lserr.IError) {
	url := getRecordUrl(s.DnsClient, popts)
	resp := new(GetRecordResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}

func (s *VDnsServiceV1) UpdateRecord(popts IUpdateRecordRequest) lserr.IError {
	url := updateRecordUrl(s.DnsClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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

func (s *VDnsServiceV1) DeleteRecord(popts IDeleteRecordRequest) lserr.IError {
	url := deleteRecordUrl(s.DnsClient, popts)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Delete(url, req); sdkErr != nil {
		return lserr.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatProductVdns)
	}

	return nil
}

func (s *VDnsServiceV1) CreateDnsRecord(popts ICreateDnsRecordRequest) (*lsentity.DnsRecord, lserr.IError) {
	url := createDnsRecordUrl(s.DnsClient, popts)
	resp := new(CreateDnsRecordResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonBody(popts.ToRequestBody(s.DnsClient)).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.DnsClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			WithErrorCategories(lserr.ErrCatProductVdns)
	}

	return resp.ToEntityDnsRecord(), nil
}
