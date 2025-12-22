package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

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

// ----------------------------------------------------------------------

type IListHostedZonesRequest interface {
	GetName() string
	WithName(name string) IListHostedZonesRequest
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type ListHostedZonesRequest struct {
	Name string

	lscommon.UserAgent
}

func (r *ListHostedZonesRequest) GetName() string {
	return r.Name
}

func (r *ListHostedZonesRequest) WithName(name string) IListHostedZonesRequest {
	r.Name = name
	return r
}

func (r *ListHostedZonesRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name": r.Name,
	}
}

func NewListHostedZonesRequest() IListHostedZonesRequest {
	return &ListHostedZonesRequest{}
}

// ----------------------------------------------------------------------

type ICreateHostedZoneRequest interface {
	WithDomainName(domainName string) ICreateHostedZoneRequest
	WithAssocVpcIds(assocVpcIds []string) ICreateHostedZoneRequest
	WithType(zoneType HostedZoneType) ICreateHostedZoneRequest
	WithDescription(description string) ICreateHostedZoneRequest
	ToRequestBody(psc lsclient.IServiceClient) map[string]interface{}
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type CreateHostedZoneRequest struct {
	DomainName  string         `json:"domainName"`
	AssocVpcIds []string       `json:"assocVpcIds"`
	Type        HostedZoneType `json:"type"`
	Description string         `json:"description"`

	lscommon.UserAgent
}

func (r *CreateHostedZoneRequest) WithDomainName(domainName string) ICreateHostedZoneRequest {
	r.DomainName = domainName
	return r
}

func (r *CreateHostedZoneRequest) WithAssocVpcIds(assocVpcIds []string) ICreateHostedZoneRequest {
	r.AssocVpcIds = assocVpcIds
	return r
}

func (r *CreateHostedZoneRequest) WithType(zoneType HostedZoneType) ICreateHostedZoneRequest {
	r.Type = zoneType
	return r
}

func (r *CreateHostedZoneRequest) WithDescription(description string) ICreateHostedZoneRequest {
	r.Description = description
	return r
}

func (r *CreateHostedZoneRequest) ToRequestBody(psc lsclient.IServiceClient) map[string]interface{} {
	return map[string]interface{}{
		"domainName":  r.DomainName,
		"assocVpcIds": r.AssocVpcIds,
		"type":        r.Type,
		"description": r.Description,
	}
}

func (r *CreateHostedZoneRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"domainName":  r.DomainName,
		"assocVpcIds": r.AssocVpcIds,
		"type":        r.Type,
		"description": r.Description,
	}
}

func NewCreateHostedZoneRequest(domainName string, assocVpcIds []string, zoneType HostedZoneType) ICreateHostedZoneRequest {
	return &CreateHostedZoneRequest{
		DomainName:  domainName,
		AssocVpcIds: assocVpcIds,
		Type:        zoneType,
	}
}

// ----------------------------------------------------------------------

type IDeleteHostedZoneRequest interface {
	GetHostedZoneId() string
	WithHostedZoneId(hostedZoneId string) IDeleteHostedZoneRequest
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type DeleteHostedZoneRequest struct {
	HostedZoneId string

	lscommon.UserAgent
}

func (r *DeleteHostedZoneRequest) GetHostedZoneId() string {
	return r.HostedZoneId
}

func (r *DeleteHostedZoneRequest) WithHostedZoneId(hostedZoneId string) IDeleteHostedZoneRequest {
	r.HostedZoneId = hostedZoneId
	return r
}

func (r *DeleteHostedZoneRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"hostedZoneId": r.HostedZoneId,
	}
}

func NewDeleteHostedZoneRequest(hostedZoneId string) IDeleteHostedZoneRequest {
	return &DeleteHostedZoneRequest{
		HostedZoneId: hostedZoneId,
	}
}

// ----------------------------------------------------------------------

type IUpdateHostedZoneRequest interface {
	GetHostedZoneId() string
	WithHostedZoneId(hostedZoneId string) IUpdateHostedZoneRequest
	WithAssocVpcIds(assocVpcIds []string) IUpdateHostedZoneRequest
	WithDescription(description string) IUpdateHostedZoneRequest
	ToRequestBody(psc lsclient.IServiceClient) map[string]interface{}
	ToMap() map[string]interface{}

	ParseUserAgent() string
}

type UpdateHostedZoneRequest struct {
	HostedZoneId string   `json:"-"`
	AssocVpcIds  []string `json:"assocVpcIds"`
	Description  string   `json:"description"`

	lscommon.UserAgent
}

func (r *UpdateHostedZoneRequest) GetHostedZoneId() string {
	return r.HostedZoneId
}

func (r *UpdateHostedZoneRequest) WithHostedZoneId(hostedZoneId string) IUpdateHostedZoneRequest {
	r.HostedZoneId = hostedZoneId
	return r
}

func (r *UpdateHostedZoneRequest) WithAssocVpcIds(assocVpcIds []string) IUpdateHostedZoneRequest {
	r.AssocVpcIds = assocVpcIds
	return r
}

func (r *UpdateHostedZoneRequest) WithDescription(description string) IUpdateHostedZoneRequest {
	r.Description = description
	return r
}

func (r *UpdateHostedZoneRequest) ToRequestBody(psc lsclient.IServiceClient) map[string]interface{} {
	return map[string]interface{}{
		"assocVpcIds": r.AssocVpcIds,
		"description": r.Description,
	}
}

func (r *UpdateHostedZoneRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"hostedZoneId": r.HostedZoneId,
		"assocVpcIds":  r.AssocVpcIds,
		"description":  r.Description,
	}
}

func NewUpdateHostedZoneRequest(hostedZoneId string) IUpdateHostedZoneRequest {
	return &UpdateHostedZoneRequest{
		HostedZoneId: hostedZoneId,
	}
}
