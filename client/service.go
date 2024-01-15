package client

import (
	"strings"

	"github.com/imroc/req/v3"

	"github.com/vngcloud/vngcloud-go-sdk/error/utils"
)

type ServiceClient struct {
	Endpoint    string
	Type        string
	MoreHeaders map[string]string

	*ProviderClient
}

func (s *ServiceClient) Post(pUrl string, pOpts *RequestOpts) (*req.Response, error) {
	if pOpts == nil {
		return nil, utils.NewErrEmptyRequestOption("")
	}

	return s.Request("POST", pUrl, pOpts)
}

func (s *ServiceClient) Get(pUrl string, pOpts *RequestOpts) (*req.Response, error) {
	if pOpts == nil {
		return nil, utils.NewErrEmptyRequestOption("")
	}

	return s.Request("GET", pUrl, pOpts)
}

func (s *ServiceClient) Delete(pUrl string, pOpts *RequestOpts) (*req.Response, error) {
	if pOpts == nil {
		return nil, utils.NewErrEmptyRequestOption("")
	}

	return s.Request("DELETE", pUrl, pOpts)
}

func (s *ServiceClient) Put(pUrl string, pOpts *RequestOpts) (*req.Response, error) {
	if pOpts == nil {
		return nil, utils.NewErrEmptyRequestOption("")
	}

	return s.Request("PUT", pUrl, pOpts)
}

func (s *ServiceClient) Request(pMethod, pUrl string, pOpts *RequestOpts) (*req.Response, error) {
	if len(s.MoreHeaders) > 0 {
		for k, v := range s.MoreHeaders {
			pOpts.MoreHeaders[k] = v
		}
	}

	return s.ProviderClient.Request(pMethod, pUrl, pOpts)
}

func (s *ServiceClient) ServiceURL(pParts ...string) string {
	return s.Endpoint + strings.Join(pParts, "/")
}
