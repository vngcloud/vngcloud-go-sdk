package certificates

type IListOptsBuilder interface {
	GetProjectID() string
}

type IGetOptsBuilder interface {
	GetProjectID() string
	GetCertificateID() string
}

type IImportOptsBuilder interface {
	GetProjectID() string
	ToRequestBody() interface{}
}

type IDeleteOptsBuilder interface {
	GetProjectID() string
	GetCertificateID() string
}
