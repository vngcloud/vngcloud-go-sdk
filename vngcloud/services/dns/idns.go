package dns

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsdnsSvcInternal "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/dns/internal_system/v1"
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

type IVDnsServiceInternal interface {
	GetHostedZoneById(popts lsdnsSvcInternal.IGetHostedZoneByIdRequest, portalUserId string) (*lsentity.HostedZone, lserr.IError)
	ListHostedZones(popts lsdnsSvcInternal.IListHostedZonesRequest, portalUserId string) (*lsentity.ListHostedZone, lserr.IError)
	CreateHostedZone(popts lsdnsSvcInternal.ICreateHostedZoneRequest, portalUserId string) (*lsentity.HostedZone, lserr.IError)
	UpdateHostedZone(popts lsdnsSvcInternal.IUpdateHostedZoneRequest, portalUserId string) lserr.IError
	DeleteHostedZone(popts lsdnsSvcInternal.IDeleteHostedZoneRequest, portalUserId string) lserr.IError

	ListRecords(popts lsdnsSvcInternal.IListRecordsRequest, portalUserId string) (*lsentity.ListDnsRecords, lserr.IError)
	GetRecord(popts lsdnsSvcInternal.IGetRecordRequest, portalUserId string) (*lsentity.DnsRecord, lserr.IError)
	CreateDnsRecord(popts lsdnsSvcInternal.ICreateDnsRecordRequest, portalUserId string) (*lsentity.DnsRecord, lserr.IError)
	UpdateRecord(popts lsdnsSvcInternal.IUpdateRecordRequest, portalUserId string) lserr.IError
	DeleteRecord(popts lsdnsSvcInternal.IDeleteRecordRequest, portalUserId string) lserr.IError
}
