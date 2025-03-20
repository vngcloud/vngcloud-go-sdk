package v2

import (
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

var _ IListCertificatesRequest = &ListCertificatesRequest{}

type ListCertificatesRequest struct {
	lscommon.UserAgent
}

func (s *ListCertificatesRequest) AddUserAgent(pagent ...string) IListCertificatesRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewListCertificatesRequest() *ListCertificatesRequest {
	return &ListCertificatesRequest{}
}

// --------------------------------------------------------

var _ IGetCertificateByIdRequest = &GetCertificateByIdRequest{}

type GetCertificateByIdRequest struct {
	lscommon.UserAgent
	CertificateId string
}

func (r *GetCertificateByIdRequest) GetCertificateId() string {
	return r.CertificateId
}

func (s *GetCertificateByIdRequest) AddUserAgent(pagent ...string) IGetCertificateByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewGetCertificateByIdRequest(pcertificateId string) *GetCertificateByIdRequest {
	return &GetCertificateByIdRequest{
		CertificateId: pcertificateId,
	}
}

// --------------------------------------------------------

type (
	ImportOptsTypeOpt string
)

const (
	ImportOptsTypeOptTLS ImportOptsTypeOpt = "TLS/SSL"
	ImportOptsTypeOptCA  ImportOptsTypeOpt = "CA"
)

var _ ICreateCertificateRequest = &CreateCertificateRequest{}

type CreateCertificateRequest struct {
	lscommon.UserAgent
	Name        string            `json:"name"`
	Type        ImportOptsTypeOpt `json:"type"`
	Certificate string            `json:"certificate"`

	CertificateChain *string `json:"certificateChain"`
	Passphrase       *string `json:"passphrase"`
	PrivateKey       *string `json:"privateKey"`
}

func (r *CreateCertificateRequest) WithCertificateChain(pchain string) ICreateCertificateRequest {
	r.CertificateChain = &pchain
	return r
}

func (r *CreateCertificateRequest) WithPassphrase(ppassphrase string) ICreateCertificateRequest {
	r.Passphrase = &ppassphrase
	return r
}

func (r *CreateCertificateRequest) WithPrivateKey(pprivateKey string) ICreateCertificateRequest {
	r.PrivateKey = &pprivateKey
	return r
}

func (r *CreateCertificateRequest) ToRequestBody() interface{} {
	return r
}

func (r *CreateCertificateRequest) ToMap() map[string]interface{} {
	re := map[string]interface{}{
		"name":        r.Name,
		"type":        r.Type,
		"certificate": r.Certificate,
	}
	if r.Type == ImportOptsTypeOptTLS {
		re["certificateChain"] = r.CertificateChain
		re["passphrase"] = r.Passphrase
		re["privateKey"] = r.PrivateKey
	}
	return re
}

func NewCreateCertificateRequest(name, cert string, pType ImportOptsTypeOpt) ICreateCertificateRequest {
	return &CreateCertificateRequest{
		Name:             name,
		Type:             pType,
		Certificate:      cert,
		CertificateChain: nil,
		Passphrase:       nil,
		PrivateKey:       nil,
	}
}

func (s *CreateCertificateRequest) AddUserAgent(pagent ...string) ICreateCertificateRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

// --------------------------------------------------------

var _ IDeleteCertificateByIdRequest = &DeleteCertificateByIdRequest{}

type DeleteCertificateByIdRequest struct {
	lscommon.UserAgent
	CertificateId string
}

func (r *DeleteCertificateByIdRequest) GetCertificateId() string {
	return r.CertificateId
}

func (s *DeleteCertificateByIdRequest) AddUserAgent(pagent ...string) IDeleteCertificateByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewDeleteCertificateByIdRequest(pcertificateId string) *DeleteCertificateByIdRequest {
	return &DeleteCertificateByIdRequest{
		CertificateId: pcertificateId,
	}
}
