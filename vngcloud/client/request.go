package client

import ljset "github.com/cuongpiger/joat/data-structure/set"

type request struct {
	JsonBody     interface{}
	JsonResponse interface{}
	JsonError    interface{}
	MoreHeaders  map[string]string
	OmitHeaders  ljset.Set[string]
	OkCodes      ljset.Set[int]
	Method       requestMethod
	SkipAuth     bool
}

type requestMethod string

const (
	MethodGet    = requestMethod("GET")
	MethodPost   = requestMethod("POST")
	MethodPut    = requestMethod("PUT")
	MethodPatch  = requestMethod("PATCH")
	MethodDelete = requestMethod("DELETE")
)

func NewRequest() IRequest {
	return &request{
		OkCodes: ljset.NewSet[int](),
	}
}

func (s *request) WithOkCodes(pokCodes ...int) IRequest {
	s.OkCodes.Append(pokCodes...)
	return s
}

func (s *request) WithUserId(puserId string) IRequest {
	return s.WithHeader("portal-user-id", puserId)
}

func (s *request) WithJsonBody(pjsonBody interface{}) IRequest {
	s.JsonBody = pjsonBody
	return s
}

func (s *request) WithJsonResponse(pjsonResponse interface{}) IRequest {
	s.JsonResponse = pjsonResponse
	return s
}

func (s *request) WithJsonError(pjsonError interface{}) IRequest {
	s.JsonError = pjsonError
	return s
}

func (s *request) WithRequestMethod(pmethod requestMethod) IRequest {
	s.Method = pmethod
	return s
}

func (s *request) WithSkipAuth(pskipAuth bool) IRequest {
	s.SkipAuth = pskipAuth
	return s
}

func (s *request) GetRequestBody() interface{} {
	return s.JsonBody
}

func (s *request) GetJsonError() interface{} {
	return s.JsonError
}

func (s *request) GetRequestMethod() string {
	return string(s.Method)
}

func (s *request) GetMoreHeaders() map[string]string {
	return s.MoreHeaders
}

func (s *request) GetOmitHeaders() ljset.Set[string] {
	return s.OmitHeaders
}

func (s *request) GetJsonResponse() interface{} {
	return s.JsonResponse
}

func (s *request) SetJsonResponse(pjsonResponse interface{}) {
	s.JsonResponse = pjsonResponse
}

func (s *request) SetJsonError(pjsonError interface{}) {
	s.JsonError = pjsonError
}

func (s *request) ContainsOkCode(pcode ...int) bool {
	return s.OkCodes.ContainsAny(pcode...)
}

func (s *request) WithHeader(pkey, pvalue string) IRequest {
	if pkey == "" || pvalue == "" {
		return s
	}

	if s.MoreHeaders == nil {
		s.MoreHeaders = make(map[string]string)
	}

	s.MoreHeaders[pkey] = pvalue
	return s
}

func (s *request) WithMapHeaders(pheaders map[string]string) IRequest {
	if s.MoreHeaders == nil {
		s.MoreHeaders = make(map[string]string)
	}

	for k, v := range pheaders {
		s.MoreHeaders[k] = v
	}

	return s
}

func (s *request) SkipAuthentication() bool {
	return s.SkipAuth
}
