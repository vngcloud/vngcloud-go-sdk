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
