package v2

import (
	lstr "strings"

	"github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

const (
	ListenerProtocolTCP   ListenerProtocol = "TCP"
	ListenerProtocolUDP   ListenerProtocol = "UDP"
	ListenerProtocolHTTP  ListenerProtocol = "HTTP"
	ListenerProtocolHTTPS ListenerProtocol = "HTTPS"
)

func NewCreateListenerRequest(pname string, pprotocol ListenerProtocol, pport int) ICreateListenerRequest {
	opts := new(CreateListenerRequest)
	opts.ListenerName = pname
	opts.ListenerProtocol = pprotocol
	opts.ListenerProtocolPort = pport
	opts.AllowedCidrs = "0.0.0.0/0"
	opts.TimeoutClient = 50
	opts.TimeoutMember = 50
	opts.TimeoutConnection = 5

	return opts
}

func NewUpdateListenerRequest(plbId, plistenerId string) IUpdateListenerRequest {
	opts := new(UpdateListenerRequest)
	opts.LoadBalancerId = plbId
	opts.ListenerId = plistenerId

	return opts
}

func NewListListenersByLoadBalancerIdRequest(plbId string) IListListenersByLoadBalancerIdRequest {
	opts := new(ListListenersByLoadBalancerIdRequest)
	opts.LoadBalancerId = plbId

	return opts
}

func NewDeleteListenerByIdRequest(plbId, plistenerId string) IDeleteListenerByIdRequest {
	opts := new(DeleteListenerByIdRequest)
	opts.LoadBalancerId = plbId
	opts.ListenerId = plistenerId

	return opts
}

func NewGetListenerByIdRequest(plbId, plistenerId string) IGetListenerByIdRequest {
	opts := new(GetListenerByIdRequest)
	opts.LoadBalancerId = plbId
	opts.ListenerId = plistenerId

	return opts
}

type ListenerProtocol string

type CreateListenerRequest struct {
	AllowedCidrs                string                         `json:"allowedCidrs"`
	ListenerName                string                         `json:"listenerName"`
	ListenerProtocol            ListenerProtocol               `json:"listenerProtocol"`
	ListenerProtocolPort        int                            `json:"listenerProtocolPort"`
	TimeoutClient               int                            `json:"timeoutClient"`
	TimeoutConnection           int                            `json:"timeoutConnection"`
	TimeoutMember               int                            `json:"timeoutMember"`
	DefaultPoolId               *string                        `json:"defaultPoolId"`
	CertificateAuthorities      *[]string                      `json:"certificateAuthorities"`
	ClientCertificate           *string                        `json:"clientCertificate"`
	DefaultCertificateAuthority *string                        `json:"defaultCertificateAuthority"`
	InsertHeaders               *[]entity.ListenerInsertHeader `json:"insertHeaders"`

	lscommon.LoadBalancerCommon
	lscommon.UserAgent
}

func (s *CreateListenerRequest) AddUserAgent(pagent ...string) ICreateListenerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type UpdateListenerRequest struct {
	AllowedCidrs                string                         `json:"allowedCidrs"`
	DefaultPoolId               string                         `json:"defaultPoolId"`
	TimeoutClient               int                            `json:"timeoutClient"`
	TimeoutConnection           int                            `json:"timeoutConnection"`
	TimeoutMember               int                            `json:"timeoutMember"`
	CertificateAuthorities      *[]string                      `json:"certificateAuthorities"`
	ClientCertificate           *string                        `json:"clientCertificate"`
	DefaultCertificateAuthority *string                        `json:"defaultCertificateAuthority"`
	InsertHeaders               *[]entity.ListenerInsertHeader `json:"insertHeaders"`

	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
	lscommon.UserAgent
}

func (s *UpdateListenerRequest) AddUserAgent(pagent ...string) IUpdateListenerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type ListListenersByLoadBalancerIdRequest struct {
	lscommon.LoadBalancerCommon
	lscommon.UserAgent
}

func (s *ListListenersByLoadBalancerIdRequest) AddUserAgent(pagent ...string) IListListenersByLoadBalancerIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type DeleteListenerByIdRequest struct {
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
	lscommon.UserAgent
}

func (s *DeleteListenerByIdRequest) AddUserAgent(pagent ...string) IDeleteListenerByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type GetListenerByIdRequest struct {
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
	lscommon.UserAgent
}

func (s *GetListenerByIdRequest) AddUserAgent(pagent ...string) IGetListenerByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreateListenerRequest) ToRequestBody() interface{} {
	if s == nil {
		return nil
	}

	if s.ListenerProtocol == ListenerProtocolHTTPS {
		return s
	}

	s.CertificateAuthorities = nil
	s.ClientCertificate = nil
	s.DefaultCertificateAuthority = nil

	return s
}

func (s *CreateListenerRequest) WithAllowedCidrs(pcidrs ...string) ICreateListenerRequest {
	if len(pcidrs) < 1 {
		return s
	}

	s.AllowedCidrs = lstr.Join(pcidrs, ",")
	return s
}

func (s *CreateListenerRequest) WithTimeoutClient(ptoc int) ICreateListenerRequest {
	s.TimeoutClient = ptoc
	return s
}

func (s *CreateListenerRequest) WithTimeoutConnection(ptoc int) ICreateListenerRequest {
	s.TimeoutConnection = ptoc
	return s
}

func (s *CreateListenerRequest) WithTimeoutMember(ptom int) ICreateListenerRequest {
	s.TimeoutMember = ptom
	return s
}

func (s *CreateListenerRequest) AddCidrs(pcidrs ...string) ICreateListenerRequest {
	if len(pcidrs) < 1 {
		return s
	}

	if s.AllowedCidrs == "" {
		return s.WithAllowedCidrs(pcidrs...)
	} else {
		s.AllowedCidrs = s.AllowedCidrs + "," + lstr.Join(pcidrs, ",")
	}

	return s
}

func (s *CreateListenerRequest) WithLoadBalancerId(plbid string) ICreateListenerRequest {
	s.LoadBalancerId = plbid
	return s
}

func (s *CreateListenerRequest) WithDefaultPoolId(ppoolId string) ICreateListenerRequest {
	s.DefaultPoolId = &ppoolId
	return s
}

func (s *CreateListenerRequest) WithCertificateAuthorities(pca *[]string) ICreateListenerRequest {
	s.CertificateAuthorities = pca
	return s
}

func (s *CreateListenerRequest) WithClientCertificate(pclientCert *string) ICreateListenerRequest {
	s.ClientCertificate = pclientCert
	return s
}

func (s *CreateListenerRequest) WithDefaultCertificateAuthority(pdefaultCA *string) ICreateListenerRequest {
	s.DefaultCertificateAuthority = pdefaultCA
	return s
}

func (s *CreateListenerRequest) WithInsertHeaders(pheaders ...string) ICreateListenerRequest {
	if len(pheaders) < 1 {
		s.InsertHeaders = nil
		return s
	}

	headers := make([]entity.ListenerInsertHeader, 0)
	for i := 0; i < len(pheaders); i += 2 {
		if i+1 >= len(pheaders) {
			break
		}
		headers = append(headers, entity.ListenerInsertHeader{
			HeaderName:  pheaders[i],
			HeaderValue: pheaders[i+1],
		})
	}
	s.InsertHeaders = &headers
	return s
}

func (s *CreateListenerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"listenerName":                s.ListenerName,
		"listenerProtocol":            s.ListenerProtocol,
		"listenerProtocolPort":        s.ListenerProtocolPort,
		"timeoutClient":               s.TimeoutClient,
		"timeoutConnection":           s.TimeoutConnection,
		"timeoutMember":               s.TimeoutMember,
		"allowedCidrs":                s.AllowedCidrs,
		"defaultPoolId":               s.DefaultPoolId,
		"certificateAuthorities":      s.CertificateAuthorities,
		"clientCertificate":           s.ClientCertificate,
		"defaultCertificateAuthority": s.DefaultCertificateAuthority,
		"insertHeaders":               s.InsertHeaders,
	}
}

func (s *UpdateListenerRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateListenerRequest) WithCidrs(pcidrs ...string) IUpdateListenerRequest {
	if len(pcidrs) < 1 {
		return s
	}

	s.AllowedCidrs = lstr.Join(pcidrs, ",")
	return s
}

func (s *UpdateListenerRequest) WithTimeoutClient(ptoc int) IUpdateListenerRequest {
	s.TimeoutClient = ptoc
	return s
}

func (s *UpdateListenerRequest) WithTimeoutConnection(ptoc int) IUpdateListenerRequest {
	s.TimeoutConnection = ptoc
	return s
}

func (s *UpdateListenerRequest) WithTimeoutMember(ptom int) IUpdateListenerRequest {
	s.TimeoutMember = ptom
	return s
}

func (s *UpdateListenerRequest) WithDefaultPoolId(ppoolId string) IUpdateListenerRequest {
	s.DefaultPoolId = ppoolId
	return s
}

func (s *UpdateListenerRequest) WithInsertHeaders(pheaders ...string) IUpdateListenerRequest {
	if len(pheaders) < 1 {
		s.InsertHeaders = nil
		return s
	}

	headers := make([]entity.ListenerInsertHeader, 0)
	for i := 0; i < len(pheaders); i += 2 {
		if i+1 >= len(pheaders) {
			break
		}
		headers = append(headers, entity.ListenerInsertHeader{
			HeaderName:  pheaders[i],
			HeaderValue: pheaders[i+1],
		})
	}
	s.InsertHeaders = &headers
	return s
}

func (s *UpdateListenerRequest) WithCertificateAuthorities(pca ...string) IUpdateListenerRequest {
	s.CertificateAuthorities = &pca
	return s
}

func (s *UpdateListenerRequest) WithClientCertificate(pclientCert string) IUpdateListenerRequest {
	s.ClientCertificate = &pclientCert
	return s
}

func (s *UpdateListenerRequest) WithDefaultCertificateAuthority(pdefaultCA string) IUpdateListenerRequest {
	s.DefaultCertificateAuthority = &pdefaultCA
	return s
}
