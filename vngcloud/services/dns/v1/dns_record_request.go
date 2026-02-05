package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

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

// ----------------------------------------------------------------------

type IGetRecordRequest interface {
	GetHostedZoneId() string
	GetRecordId() string
	WithHostedZoneId(hostedZoneId string) IGetRecordRequest
	WithRecordId(recordId string) IGetRecordRequest
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type GetRecordRequest struct {
	HostedZoneId string
	RecordId     string

	lscommon.UserAgent
}

func (r *GetRecordRequest) GetHostedZoneId() string {
	return r.HostedZoneId
}

func (r *GetRecordRequest) GetRecordId() string {
	return r.RecordId
}

func (r *GetRecordRequest) WithHostedZoneId(hostedZoneId string) IGetRecordRequest {
	r.HostedZoneId = hostedZoneId
	return r
}

func (r *GetRecordRequest) WithRecordId(recordId string) IGetRecordRequest {
	r.RecordId = recordId
	return r
}

func (r *GetRecordRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"hostedZoneId": r.HostedZoneId,
		"recordId":     r.RecordId,
	}
}

func NewGetRecordRequest(hostedZoneId, recordId string) IGetRecordRequest {
	return &GetRecordRequest{
		HostedZoneId: hostedZoneId,
		RecordId:     recordId,
	}
}

// ----------------------------------------------------------------------

type IUpdateRecordRequest interface {
	GetHostedZoneId() string
	GetRecordId() string
	WithHostedZoneId(hostedZoneId string) IUpdateRecordRequest
	WithRecordId(recordId string) IUpdateRecordRequest
	WithSubDomain(subDomain string) IUpdateRecordRequest
	WithTTL(ttl int) IUpdateRecordRequest
	WithType(recordType DnsRecordType) IUpdateRecordRequest
	WithRoutingPolicy(routingPolicy RoutingPolicy) IUpdateRecordRequest
	WithEnableStickySession(enable bool) IUpdateRecordRequest
	WithValue(value []RecordValueRequest) IUpdateRecordRequest
	ToRequestBody(psc lsclient.IServiceClient) map[string]interface{}
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type UpdateRecordRequest struct {
	HostedZoneId        string               `json:"-"`
	RecordId            string               `json:"-"`
	SubDomain           string               `json:"subDomain"`
	TTL                 int                  `json:"ttl"`
	Type                DnsRecordType        `json:"type"`
	RoutingPolicy       RoutingPolicy        `json:"routingPolicy"`
	EnableStickySession *bool                `json:"enableStickySession,omitempty"`
	Value               []RecordValueRequest `json:"value"`

	lscommon.UserAgent
}

func (r *UpdateRecordRequest) GetHostedZoneId() string {
	return r.HostedZoneId
}

func (r *UpdateRecordRequest) GetRecordId() string {
	return r.RecordId
}

func (r *UpdateRecordRequest) WithHostedZoneId(hostedZoneId string) IUpdateRecordRequest {
	r.HostedZoneId = hostedZoneId
	return r
}

func (r *UpdateRecordRequest) WithRecordId(recordId string) IUpdateRecordRequest {
	r.RecordId = recordId
	return r
}

func (r *UpdateRecordRequest) WithSubDomain(subDomain string) IUpdateRecordRequest {
	r.SubDomain = subDomain
	return r
}

func (r *UpdateRecordRequest) WithTTL(ttl int) IUpdateRecordRequest {
	r.TTL = ttl
	return r
}

func (r *UpdateRecordRequest) WithType(recordType DnsRecordType) IUpdateRecordRequest {
	r.Type = recordType
	return r
}

func (r *UpdateRecordRequest) WithRoutingPolicy(routingPolicy RoutingPolicy) IUpdateRecordRequest {
	r.RoutingPolicy = routingPolicy
	return r
}

func (r *UpdateRecordRequest) WithEnableStickySession(enable bool) IUpdateRecordRequest {
	r.EnableStickySession = &enable
	return r
}

func (r *UpdateRecordRequest) WithValue(value []RecordValueRequest) IUpdateRecordRequest {
	r.Value = value
	return r
}

func (r *UpdateRecordRequest) ToRequestBody(psc lsclient.IServiceClient) map[string]interface{} {
	body := map[string]interface{}{
		"subDomain":     r.SubDomain,
		"ttl":           r.TTL,
		"type":          r.Type,
		"routingPolicy": r.RoutingPolicy,
		"value":         r.Value,
	}
	if r.EnableStickySession != nil {
		body["enableStickySession"] = *r.EnableStickySession
	}
	return body
}

func (r *UpdateRecordRequest) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"hostedZoneId":  r.HostedZoneId,
		"recordId":      r.RecordId,
		"subDomain":     r.SubDomain,
		"ttl":           r.TTL,
		"type":          r.Type,
		"routingPolicy": r.RoutingPolicy,
		"value":         r.Value,
	}
	if r.EnableStickySession != nil {
		m["enableStickySession"] = *r.EnableStickySession
	}
	return m
}

func NewUpdateRecordRequest(hostedZoneId, recordId string) IUpdateRecordRequest {
	return &UpdateRecordRequest{
		HostedZoneId: hostedZoneId,
		RecordId:     recordId,
	}
}

// ----------------------------------------------------------------------

type IDeleteRecordRequest interface {
	GetHostedZoneId() string
	GetRecordId() string
	WithHostedZoneId(hostedZoneId string) IDeleteRecordRequest
	WithRecordId(recordId string) IDeleteRecordRequest
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type DeleteRecordRequest struct {
	HostedZoneId string
	RecordId     string

	lscommon.UserAgent
}

func (r *DeleteRecordRequest) GetHostedZoneId() string {
	return r.HostedZoneId
}

func (r *DeleteRecordRequest) GetRecordId() string {
	return r.RecordId
}

func (r *DeleteRecordRequest) WithHostedZoneId(hostedZoneId string) IDeleteRecordRequest {
	r.HostedZoneId = hostedZoneId
	return r
}

func (r *DeleteRecordRequest) WithRecordId(recordId string) IDeleteRecordRequest {
	r.RecordId = recordId
	return r
}

func (r *DeleteRecordRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"hostedZoneId": r.HostedZoneId,
		"recordId":     r.RecordId,
	}
}

func NewDeleteRecordRequest(hostedZoneId, recordId string) IDeleteRecordRequest {
	return &DeleteRecordRequest{
		HostedZoneId: hostedZoneId,
		RecordId:     recordId,
	}
}

// ----------------------------------------------------------------------

type ICreateDnsRecordRequest interface {
	GetHostedZoneId() string
	WithHostedZoneId(hostedZoneId string) ICreateDnsRecordRequest
	WithSubDomain(subDomain string) ICreateDnsRecordRequest
	WithTTL(ttl int) ICreateDnsRecordRequest
	WithType(recordType DnsRecordType) ICreateDnsRecordRequest
	WithRoutingPolicy(routingPolicy RoutingPolicy) ICreateDnsRecordRequest
	WithEnableStickySession(enable bool) ICreateDnsRecordRequest
	WithValue(value []RecordValueRequest) ICreateDnsRecordRequest
	ToRequestBody(psc lsclient.IServiceClient) map[string]interface{}
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type RecordValueRequest struct {
	Value    string  `json:"value"`
	Location *string `json:"location,omitempty"`
	Weight   *int    `json:"weight,omitempty"`
}

type CreateDnsRecordRequest struct {
	HostedZoneId        string               `json:"-"`
	SubDomain           string               `json:"subDomain"`
	TTL                 int                  `json:"ttl"`
	Type                DnsRecordType        `json:"type"`
	RoutingPolicy       RoutingPolicy        `json:"routingPolicy"`
	EnableStickySession *bool                `json:"enableStickySession,omitempty"`
	Value               []RecordValueRequest `json:"value"`

	lscommon.UserAgent
}

func (r *CreateDnsRecordRequest) GetHostedZoneId() string {
	return r.HostedZoneId
}

func (r *CreateDnsRecordRequest) WithHostedZoneId(hostedZoneId string) ICreateDnsRecordRequest {
	r.HostedZoneId = hostedZoneId
	return r
}

func (r *CreateDnsRecordRequest) WithSubDomain(subDomain string) ICreateDnsRecordRequest {
	r.SubDomain = subDomain
	return r
}

func (r *CreateDnsRecordRequest) WithTTL(ttl int) ICreateDnsRecordRequest {
	r.TTL = ttl
	return r
}

func (r *CreateDnsRecordRequest) WithType(recordType DnsRecordType) ICreateDnsRecordRequest {
	r.Type = recordType
	return r
}

func (r *CreateDnsRecordRequest) WithRoutingPolicy(routingPolicy RoutingPolicy) ICreateDnsRecordRequest {
	r.RoutingPolicy = routingPolicy
	return r
}

func (r *CreateDnsRecordRequest) WithEnableStickySession(enable bool) ICreateDnsRecordRequest {
	r.EnableStickySession = &enable
	return r
}

func (r *CreateDnsRecordRequest) WithValue(value []RecordValueRequest) ICreateDnsRecordRequest {
	r.Value = value
	return r
}

func (r *CreateDnsRecordRequest) ToRequestBody(psc lsclient.IServiceClient) map[string]interface{} {
	body := map[string]interface{}{
		"subDomain":     r.SubDomain,
		"ttl":           r.TTL,
		"type":          r.Type,
		"routingPolicy": r.RoutingPolicy,
		"value":         r.Value,
	}
	if r.EnableStickySession != nil {
		body["enableStickySession"] = *r.EnableStickySession
	}
	return body
}

func (r *CreateDnsRecordRequest) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"hostedZoneId":  r.HostedZoneId,
		"subDomain":     r.SubDomain,
		"ttl":           r.TTL,
		"type":          r.Type,
		"routingPolicy": r.RoutingPolicy,
		"value":         r.Value,
	}
	if r.EnableStickySession != nil {
		m["enableStickySession"] = *r.EnableStickySession
	}
	return m
}

func NewCreateDnsRecordRequest(hostedZoneId, subDomain string, ttl int, recordType DnsRecordType, routingPolicy RoutingPolicy, value []RecordValueRequest) ICreateDnsRecordRequest {
	return &CreateDnsRecordRequest{
		HostedZoneId:  hostedZoneId,
		SubDomain:     subDomain,
		TTL:           ttl,
		Type:          recordType,
		RoutingPolicy: routingPolicy,
		Value:         value,
	}
}

// NewRecordValueRequest creates a record value with optional location and weight pointers
func NewRecordValueRequest(value string, location *string, weight *int) RecordValueRequest {
	return RecordValueRequest{
		Value:    value,
		Location: location,
		Weight:   weight,
	}
}
