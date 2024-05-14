package client

type serviceClient struct {
	name        string
	endpoint    string
	moreHeaders map[string]string

	IHttpClient
}

func NewServiceClient() IServiceClient {
	return &serviceClient{}
}

func (s *serviceClient) WithEndpoint(pendpoint string) IServiceClient {
	s.endpoint = pendpoint
	return s
}

func (s *serviceClient) WithName(pname string) IServiceClient {
	s.name = pname
	return s
}

func (s *serviceClient) WithMoreHeaders(pmoreHeaders map[string]string) IServiceClient {
	s.moreHeaders = pmoreHeaders
	return s
}

func (s *serviceClient) WithKVheaders(pkey string, pvalue string) IServiceClient {
	s.moreHeaders[pkey] = pvalue
	return s
}
