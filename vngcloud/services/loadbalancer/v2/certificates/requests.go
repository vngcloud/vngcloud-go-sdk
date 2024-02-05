package certificates

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lbCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2"
)

type (
	ImportOptsTypeOpt string
)

const (
	ImportOptsTypeOptTLS ImportOptsTypeOpt = "TLS/SSL"
	ImportOptsTypeOptCA  ImportOptsTypeOpt = "CA"
)

type ImportOpts struct {
	Name        string            `json:"name"`
	Type        ImportOptsTypeOpt `json:"type"`
	Certificate string            `json:"certificate"`

	CertificateChain *string `json:"certificateChain"`
	Passphrase       *string `json:"passphrase"`
	PrivateKey       *string `json:"privateKey"`
	common.CommonOpts
}

func (s *ImportOpts) ToRequestBody() interface{} {
	if s == nil {
		return nil
	}
	if s.Type != ImportOptsTypeOptTLS {
		s.CertificateChain = nil
		s.Passphrase = nil
		s.PrivateKey = nil
	}
	return s
}

type GetOpts struct {
	common.CommonOpts
	lbCm.CertificateV2Common
}

type DeleteOpts struct {
	common.CommonOpts
	lbCm.CertificateV2Common
}

type ListOpts struct {
	common.CommonOpts
}
