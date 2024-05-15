package http

import (
	ljset "github.com/cuongpiger/joat/data-structure/set"
	ltime "time"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type IHttpClient interface {
	WithRetryCount(pretryCount int) IHttpClient
	WithDelay(pdelay ltime.Duration) IHttpClient
	WithTimeout(ptimeout ltime.Duration) IHttpClient
	WithSleep(psleep ltime.Duration) IHttpClient
	WithReauthFunc(preauthFunc func() (lsclient.ISdkAuthentication, lserr.ISdkError)) IHttpClient

	DoRequest(purl string, preq IRequest) lserr.ISdkError
}

type IRequest interface {
	WithOkCodes(pokCodes ...int) IRequest
	WithJsonBody(pjsonBody interface{}) IRequest
	WithJsonResponse(pjsonResponse interface{}) IRequest
	WithJsonError(pjsonError interface{}) IRequest
	WithRequestMethod(pmethod requestMethod) IRequest
	WithHeader(pkey, pvalue string) IRequest

	GetRequestBody() interface{}
	GetRequestMethod() string
	GetMoreHeaders() map[string]string
	GetOmitHeaders() ljset.Set[string]
	GetJsonResponse() interface{}

	SetJsonResponse(pjsonResponse interface{})

	ContainsOkCode(pcode ...int) bool
}
