package v1

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

type IListRecordsRequest interface {
	GetHostedZoneId() string
	GetName() string
	WithHostedZoneId(hostedZoneId string) IListRecordsRequest
	WithName(name string) IListRecordsRequest
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type ListRecordsRequest struct {
	HostedZoneId string
	Name         string

	lscommon.UserAgent
}

func (r *ListRecordsRequest) GetHostedZoneId() string {
	return r.HostedZoneId
}

func (r *ListRecordsRequest) GetName() string {
	return r.Name
}

func (r *ListRecordsRequest) WithHostedZoneId(hostedZoneId string) IListRecordsRequest {
	r.HostedZoneId = hostedZoneId
	return r
}

func (r *ListRecordsRequest) WithName(name string) IListRecordsRequest {
	r.Name = name
	return r
}

func (r *ListRecordsRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"hostedZoneId": r.HostedZoneId,
		"name":         r.Name,
	}
}

func NewListRecordsRequest(hostedZoneId string) IListRecordsRequest {
	return &ListRecordsRequest{
		HostedZoneId: hostedZoneId,
	}
}