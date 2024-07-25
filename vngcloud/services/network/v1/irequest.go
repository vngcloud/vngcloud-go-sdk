package v1

type IGetEndpointByIdRequest interface {
	GetEndpointId() string
	AddUserAgent(pagent ...string) IGetEndpointByIdRequest
}
