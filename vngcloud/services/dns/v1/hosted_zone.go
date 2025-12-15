package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VDnsServiceV1) GetHostedZoneById(popts IGetHostedZoneByIdRequest) (*lsentity.HostedZone, lserr.IError) {
	url := getHostedZoneByIdUrl(s.DnsClient, popts)
	resp := new(GetHostedZoneByIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NetworkGatewayErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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
