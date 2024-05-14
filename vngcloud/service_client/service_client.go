package service_client

import (
	lstr "strings"

	lshttp "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/http"
)

type serviceClient struct {
	name        string
	endpoint    string
	moreHeaders map[string]string
	client      lshttp.IHttpClient
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

func (s *serviceClient) WithKVheader(pkey string, pvalue string) IServiceClient {
	s.moreHeaders[pkey] = pvalue
	return s
}

func (s *serviceClient) WithClient(pclient lshttp.IHttpClient) IServiceClient {
	s.client = pclient
	return s
}

func (s *serviceClient) ServiceURL(pparts ...string) string {
	return s.endpoint + lstr.Join(pparts, "/")
}
