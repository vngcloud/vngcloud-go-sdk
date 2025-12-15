package dns

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsdnsSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/dns/v1"
)

type IVDnsServiceV1 interface {
	GetHostedZoneById(popts lsdnsSvcV1.IGetHostedZoneByIdRequest) (*lsentity.HostedZone, lserr.IError)
	ListHostedZones(popts lsdnsSvcV1.IListHostedZonesRequest) (*lsentity.ListHostedZone, lserr.IError)
	CreateHostedZone(popts lsdnsSvcV1.ICreateHostedZoneRequest) (*lsentity.HostedZone, lserr.IError)

	ListRecords(popts lsdnsSvcV1.IListRecordsRequest) (*lsentity.ListDnsRecords, lserr.IError)
}
