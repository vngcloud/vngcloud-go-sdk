package client

type serviceClient struct {
	name        string
	moreHeaders map[string]string

	IHttpClient
}

func NewServiceClient() IServiceClient {
	return &serviceClient{}
}
