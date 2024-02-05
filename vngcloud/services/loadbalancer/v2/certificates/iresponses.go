package certificates

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type IListResponse interface {
	ToListCertificateObjects() []*objects.Certificate
}

type IGetResponse interface {
	ToCertificateObject() *objects.Certificate
}

type IImportResponse interface {
	ToCertificateObject() *objects.Certificate
}

type IDeleteResponse interface{}
