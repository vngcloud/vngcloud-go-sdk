package client

import (
	ltime "time"

	lreq "github.com/imroc/req/v3"

	ljset "github.com/cuongpiger/joat/data-structure/set"

	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type IHttpClient interface {
	WithRetryCount(pretryCount int) IHttpClient
	WithTimeout(ptimeout ltime.Duration) IHttpClient
	WithSleep(psleep ltime.Duration) IHttpClient
	WithKvDefaultHeaders(pargs ...string) IHttpClient
	WithReauthFunc(pauthOpt AuthOpts, preauthFunc func() (ISdkAuthentication, lserr.IError)) IHttpClient

	DoRequest(purl string, preq IRequest) (*lreq.Response, lserr.IError)
}

type IRequest interface {
	WithOkCodes(pokCodes ...int) IRequest
	WithJsonBody(pjsonBody interface{}) IRequest
	WithJsonResponse(pjsonResponse interface{}) IRequest
	WithJsonError(pjsonError interface{}) IRequest
	WithRequestMethod(pmethod requestMethod) IRequest
	WithSkipAuth(pskipAuth bool) IRequest
	WithHeader(pkey, pvalue string) IRequest
	WithMapHeaders(pheaders map[string]string) IRequest
	WithUserId(puserId string) IRequest

	GetRequestBody() interface{}
	GetRequestMethod() string
	GetMoreHeaders() map[string]string
	GetOmitHeaders() ljset.Set[string]
	GetJsonResponse() interface{}
	GetJsonError() interface{}

	SetJsonResponse(pjsonResponse interface{})
	SetJsonError(pjsonError interface{})

	ContainsOkCode(pcode ...int) bool
	SkipAuthentication() bool
}
