package inter

import (
	lstr "strings"

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

func (s *CreateListenerRequest) WithLoadBalancerId(plbid string) ICreateListenerRequest {
	s.LoadBalancerId = plbid
	return s
}

func (s *CreateListenerRequest) WithDefaultPoolId(ppoolId string) ICreateListenerRequest {
	s.DefaultPoolId = &ppoolId
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
