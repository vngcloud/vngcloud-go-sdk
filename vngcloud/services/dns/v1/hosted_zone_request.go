package v1

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

type IGetHostedZoneByIdRequest interface {
	GetHostedZoneId() string
	WithHostedZoneId(hostedZoneId string) IGetHostedZoneByIdRequest
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type GetHostedZoneByIdRequest struct {
	HostedZoneId string

	lscommon.UserAgent
}

func (r *GetHostedZoneByIdRequest) GetHostedZoneId() string {
	return r.HostedZoneId
}

func (r *GetHostedZoneByIdRequest) WithHostedZoneId(hostedZoneId string) IGetHostedZoneByIdRequest {
	r.HostedZoneId = hostedZoneId
	return r
}

func (r *GetHostedZoneByIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"hostedZoneId": r.HostedZoneId,
	}
}

func NewGetHostedZoneByIdRequest(hostedZoneId string) IGetHostedZoneByIdRequest {
	return &GetHostedZoneByIdRequest{
		HostedZoneId: hostedZoneId,
	}
}
