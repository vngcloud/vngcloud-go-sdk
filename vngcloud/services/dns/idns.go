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
	UpdateHostedZone(popts lsdnsSvcV1.IUpdateHostedZoneRequest) lserr.IError
	DeleteHostedZone(popts lsdnsSvcV1.IDeleteHostedZoneRequest) lserr.IError

	ListRecords(popts lsdnsSvcV1.IListRecordsRequest) (*lsentity.ListDnsRecords, lserr.IError)
	GetRecord(popts lsdnsSvcV1.IGetRecordRequest) (*lsentity.DnsRecord, lserr.IError)
	CreateDnsRecord(popts lsdnsSvcV1.ICreateDnsRecordRequest) (*lsentity.DnsRecord, lserr.IError)
	UpdateRecord(popts lsdnsSvcV1.IUpdateRecordRequest) lserr.IError
	DeleteRecord(popts lsdnsSvcV1.IDeleteRecordRequest) lserr.IError
}
