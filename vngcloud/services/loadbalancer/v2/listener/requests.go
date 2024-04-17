package listener

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lbCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2"
)

// **************************************************** CreateOpts *****************************************************

type (
	CreateOptsListenerProtocolOpt string
)

const (
	CreateOptsListenerProtocolOptTCP   CreateOptsListenerProtocolOpt = "TCP"
	CreateOptsListenerProtocolOptUDP   CreateOptsListenerProtocolOpt = "UDP"
	CreateOptsListenerProtocolOptHTTP  CreateOptsListenerProtocolOpt = "HTTP"
	CreateOptsListenerProtocolOptHTTPS CreateOptsListenerProtocolOpt = "HTTPS"
)

type CreateOpts struct {
	AllowedCidrs                string                        `json:"allowedCidrs"`
	ListenerName                string                        `json:"listenerName"`
	ListenerProtocol            CreateOptsListenerProtocolOpt `json:"listenerProtocol"`
	ListenerProtocolPort        int                           `json:"listenerProtocolPort"`
	TimeoutClient               int                           `json:"timeoutClient"`
	TimeoutConnection           int                           `json:"timeoutConnection"`
	TimeoutMember               int                           `json:"timeoutMember"`
	DefaultPoolId               *string                       `json:"defaultPoolId"`
	CertificateAuthorities      *[]string                     `json:"certificateAuthorities"`
	ClientCertificate           *string                       `json:"clientCertificate"`
	DefaultCertificateAuthority *string                       `json:"defaultCertificateAuthority"`

	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

func (s *CreateOpts) ToRequestBody() interface{} {
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

// ****************************************************** GetOpts ******************************************************

type GetBasedLoadBalancerOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

// ***************************************************** DeleteOpts ****************************************************

type DeleteOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.ListenerV2Common
}

type UpdateOpts struct {
	AllowedCidrs                string   `json:"allowedCidrs"`
	DefaultPoolId               string   `json:"defaultPoolId"`
	TimeoutClient               int      `json:"timeoutClient"`
	TimeoutConnection           int      `json:"timeoutConnection"`
	TimeoutMember               int      `json:"timeoutMember"`
	Headers                     []string `json:"headers"`
	ClientCertificate           *string  `json:"clientCertificate"`
	DefaultCertificateAuthority *string  `json:"defaultCertificateAuthority"`
	// CertificateAuthorities      *[]string `json:"certificateAuthorities"`

	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.ListenerV2Common
}

func (s *UpdateOpts) ToRequestBody() interface{} {
	return s
}
