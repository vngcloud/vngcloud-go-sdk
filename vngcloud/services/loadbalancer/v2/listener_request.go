package v2

import (
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
	lstr "strings"
)

const (
	CreateOptsListenerProtocolOptTCP   ListenerProtocol = "TCP"
	CreateOptsListenerProtocolOptUDP   ListenerProtocol = "UDP"
	CreateOptsListenerProtocolOptHTTP  ListenerProtocol = "HTTP"
	CreateOptsListenerProtocolOptHTTPS ListenerProtocol = "HTTPS"
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

type ListenerProtocol string

type CreateListenerRequest struct {
	AllowedCidrs                string           `json:"allowedCidrs"`
	ListenerName                string           `json:"listenerName"`
	ListenerProtocol            ListenerProtocol `json:"listenerProtocol"`
	ListenerProtocolPort        int              `json:"listenerProtocolPort"`
	TimeoutClient               int              `json:"timeoutClient"`
	TimeoutConnection           int              `json:"timeoutConnection"`
	TimeoutMember               int              `json:"timeoutMember"`
	DefaultPoolId               *string          `json:"defaultPoolId"`
	CertificateAuthorities      *[]string        `json:"certificateAuthorities"`
	ClientCertificate           *string          `json:"clientCertificate"`
	DefaultCertificateAuthority *string          `json:"defaultCertificateAuthority"`

	lscommon.LoadBalancerCommon
	lscommon.UserAgent
}

type UpdateListenerRequest struct {
	AllowedCidrs                string   `json:"allowedCidrs"`
	DefaultPoolId               string   `json:"defaultPoolId"`
	TimeoutClient               int      `json:"timeoutClient"`
	TimeoutConnection           int      `json:"timeoutConnection"`
	TimeoutMember               int      `json:"timeoutMember"`
	Headers                     []string `json:"headers"`
	ClientCertificate           *string  `json:"clientCertificate"`
	DefaultCertificateAuthority *string  `json:"defaultCertificateAuthority"`

	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
	lscommon.UserAgent
}

func (s *CreateListenerRequest) ToRequestBody() interface{} {
	if s == nil {
		return nil
	}

	if s.ListenerProtocol == CreateOptsListenerProtocolOptHTTPS {
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

func (s *CreateListenerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"listenerName":         s.ListenerName,
		"listenerProtocol":     s.ListenerProtocol,
		"listenerProtocolPort": s.ListenerProtocolPort,
		"timeoutClient":        s.TimeoutClient,
		"timeoutConnection":    s.TimeoutConnection,
		"timeoutMember":        s.TimeoutMember,
		"allowedCidrs":         s.AllowedCidrs,
		"defaultPoolId":        s.DefaultPoolId,
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

func (s *UpdateListenerRequest) WithHeaders(pheaders ...string) IUpdateListenerRequest {
	if len(pheaders) < 1 {
		return s
	}

	s.Headers = pheaders
	return s
}
