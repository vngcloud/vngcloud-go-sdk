package certificates

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type ResponseData struct {
	UUID               string `json:"uuid"`
	Name               string `json:"name"`
	CertificateType    string `json:"certificateType"`
	ExpiredAt          string `json:"expiredAt"`
	ImportedAt         string `json:"importedAt"`
	NotAfter           int64  `json:"notAfter"`
	KeyAlgorithm       string `json:"keyAlgorithm"`
	Serial             string `json:"serial"`
	Subject            string `json:"subject"`
	DomainName         string `json:"domainName"`
	InUse              bool   `json:"inUse"`
	Issuer             string `json:"issuer"`
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	NotBefore          int64  `json:"notBefore"`
	// SubjectAlternativeNames
}

func (s *ResponseData) ToCertificateObject() *objects.Certificate {
	if s == nil {
		return nil
	}

	return &objects.Certificate{
		UUID:               s.UUID,
		Name:               s.Name,
		CertificateType:    s.CertificateType,
		ExpiredAt:          s.ExpiredAt,
		ImportedAt:         s.ImportedAt,
		NotAfter:           s.NotAfter,
		KeyAlgorithm:       s.KeyAlgorithm,
		Serial:             s.Serial,
		Subject:            s.Subject,
		DomainName:         s.DomainName,
		InUse:              s.InUse,
		Issuer:             s.Issuer,
		SignatureAlgorithm: s.SignatureAlgorithm,
		NotBefore:          s.NotBefore,
		// SubjectAlternativeNames
	}
}

// ******************************************* Response of CreateVolume API ********************************************

type ImportResponse struct {
	Data ResponseData `json:"data"`
}

func (s *ImportResponse) ToCertificateObject() *objects.Certificate {
	if s == nil {
		return nil
	}

	return s.Data.ToCertificateObject()
}

type GetResponse struct {
	ResponseData
}

type ListResponse struct {
	ListData  []ResponseData `json:"listData"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

func (s *ListResponse) ToListCertificateObjects() []*objects.Certificate {
	if s == nil || s.ListData == nil || len(s.ListData) < 1 {
		return nil
	}

	var result []*objects.Certificate
	for _, itemLb := range s.ListData {
		result = append(result, itemLb.ToCertificateObject())
	}

	return result
}
