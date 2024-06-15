package inter

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

const (
	CreateOptsListenerProtocolOptTCP   ListenerProtocol = "TCP"
	CreateOptsListenerProtocolOptUDP   ListenerProtocol = "UDP"
	CreateOptsListenerProtocolOptHTTP  ListenerProtocol = "HTTP"
	CreateOptsListenerProtocolOptHTTPS ListenerProtocol = "HTTPS"
)

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
