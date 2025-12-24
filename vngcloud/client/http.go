package client

import (
	lctx "context"
	lhttp "net/http"
	lstr "strings"
	lsync "sync"
	ltime "time"

	ljtime "github.com/cuongpiger/joat/timer"
	lreq "github.com/imroc/req/v3"

	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

const (
	IamOauth2 AuthOpts = "IamOauth2"
)

type (
	httpClient struct {
		context    lctx.Context
		retryCount int
		client     *lreq.Client

		reauthFunc   func() (ISdkAuthentication, lserr.IError)
		reauthOption AuthOpts

		accessToken    ISdkAuthentication
		defaultHeaders map[string]string

		mut       *lsync.RWMutex
		reauthmut *reauthlock
	}

	reauthlock struct {
		lsync.RWMutex
		ongoing *reauthFuture
	}

	reauthFuture struct {
		done chan struct{}
		err  lserr.IError
	}

	AuthOpts string
)

func NewHttpClient(pctx lctx.Context) IHttpClient {
	return &httpClient{
		context:    pctx,
		retryCount: 0,
		client: lreq.NewClient().
			SetCommonRetryCount(3).
			SetCommonRetryFixedInterval(10).
			SetTimeout(ljtime.Second(120)),
		mut:       new(lsync.RWMutex),
		reauthmut: new(reauthlock),
	}
}

func (s *httpClient) WithRetryCount(pretryCount int) IHttpClient {
	s.client.SetCommonRetryCount(pretryCount)
	return s
}

func (s *httpClient) WithTimeout(ptimeout ltime.Duration) IHttpClient {
	s.client.SetTimeout(ptimeout)
	return s
}

func (s *httpClient) WithSleep(psleep ltime.Duration) IHttpClient {
	s.client.SetCommonRetryFixedInterval(psleep)
	return s
}

func (s *httpClient) WithKvDefaultHeaders(pargs ...string) IHttpClient {
	if s.defaultHeaders == nil {
		s.defaultHeaders = make(map[string]string)
	}

	if len(pargs)%2 != 0 {
		pargs = append(pargs, "")
	}

	for i := 0; i < len(pargs); i += 2 {
		s.defaultHeaders[pargs[i]] = pargs[i+1]
	}

	return s
}

func (s *httpClient) WithReauthFunc(pauthOpt AuthOpts, preauthFunc func() (ISdkAuthentication, lserr.IError)) IHttpClient {
	s.reauthFunc = preauthFunc
	s.reauthOption = pauthOpt
	return s
}

func (s *httpClient) DoRequest(purl string, preq IRequest) (*lreq.Response, lserr.IError) {
	req := s.client.R().SetContext(s.context).SetHeaders(s.getDefaultHeaders()).SetHeaders(preq.GetMoreHeaders())
	if opt := preq.GetRequestBody(); opt != nil {
		req.SetBodyJsonMarshal(opt)
	}

	if opt := preq.GetJsonResponse(); opt != nil {
		req.SetSuccessResult(opt)
	}

	if opt := preq.GetJsonError(); opt != nil {
		req.SetErrorResult(opt)
	}

	var resp *lreq.Response
	if !s.needReauth(preq) {
		var err error

		switch lstr.ToUpper(preq.GetRequestMethod()) {
		case "POST":
			resp, err = req.Post(purl)
		case "GET":
			resp, err = req.Get(purl)
		case "DELETE":
			resp, err = req.Delete(purl)
		case "PUT":
			resp, err = req.Put(purl)
		case "PATCH":
			resp, err = req.Patch(purl)
		}

		if err != nil && resp == nil {
			return resp, lserr.ErrorHandler(err)
		}
	} else {
		if !preq.SkipAuthentication() && s.reauthFunc != nil {
			if sdkErr := s.reauthenticate(); sdkErr != nil {
				return nil, sdkErr
			}

			return s.DoRequest(purl, preq)
		}
	}

	if resp.Response != nil {
		switch resp.Response.StatusCode {
		case lhttp.StatusUnauthorized:
			if !preq.SkipAuthentication() && s.reauthFunc != nil {
				if sdkErr := s.reauthenticate(); sdkErr != nil {
					return nil, sdkErr
				}

				return s.DoRequest(purl, preq)
			} else {
				return nil, defaultErrorResponse(resp.Err, purl, preq, resp)
			}
		case lhttp.StatusTooManyRequests:
			return nil, lserr.SdkErrorHandler(
				defaultErrorResponse(resp.Err, purl, preq, resp), nil,
				lserr.WithErrorPermissionDenied())
		case lhttp.StatusInternalServerError:
			return nil, lserr.SdkErrorHandler(
				defaultErrorResponse(resp.Err, purl, preq, resp), nil,
				lserr.WithErrorInternalServerError())
		case lhttp.StatusServiceUnavailable:
			return nil, lserr.SdkErrorHandler(
				defaultErrorResponse(resp.Err, purl, preq, resp), nil,
				lserr.WithErrorServiceMaintenance())
		case lhttp.StatusForbidden:
			return nil, lserr.SdkErrorHandler(
				defaultErrorResponse(resp.Err, purl, preq, resp), nil,
				lserr.WithErrorPermissionDenied())
		}

		if preq.ContainsOkCode(resp.StatusCode) {
			return resp, nil
		}

		return resp, lserr.ErrorHandler(resp.Err)
	}

	return nil, lserr.ErrorHandler(nil, lserr.WithErrorUnexpected(resp))
}

func (s *httpClient) needReauth(preq IRequest) bool {
	if preq.SkipAuthentication() {
		return false
	}

	if s.accessToken == nil {
		return true
	}

	return s.accessToken.NeedReauth()
}

func (s *httpClient) reauthenticate() lserr.IError {
	if s.reauthFunc == nil {
		return lserr.ErrorHandler(nil, lserr.WithErrorReauthFuncNotSet())
	}

	s.reauthmut.Lock()
	ongoing := s.reauthmut.ongoing
	if ongoing == nil {
		s.reauthmut.ongoing = newReauthFuture()
	}
	s.reauthmut.Unlock()

	if ongoing != nil {
		return ongoing.get()
	}

	auth, sdkerr := s.reauthFunc()
	s.reauthmut.Lock()
	s.reauthmut.ongoing.set(sdkerr)
	s.reauthmut.ongoing = nil
	s.reauthmut.Unlock()

	s.setAccessToken(auth)

	return sdkerr
}

func (s *httpClient) setAccessToken(pnewToken ISdkAuthentication) IHttpClient {
	s.mut.Lock()
	defer s.mut.Unlock()
	if pnewToken != nil {
		s.accessToken = pnewToken
		s.WithKvDefaultHeaders("Authorization", "Bearer "+s.accessToken.GetAccessToken())
	}

	return s
}

func (s *httpClient) getDefaultHeaders() map[string]string {
	s.mut.RLock()
	defer s.mut.RUnlock()
	if s.defaultHeaders == nil {
		s.defaultHeaders = make(map[string]string)
	}

	return s.defaultHeaders
}

func newReauthFuture() *reauthFuture {
	return &reauthFuture{
		done: make(chan struct{}),
		err:  nil,
	}
}

func (s *reauthFuture) get() lserr.IError {
	<-s.done
	return s.err
}

func (s *reauthFuture) set(err lserr.IError) {
	s.err = err
	close(s.done)
}

func defaultErrorResponse(perr error, purl string, preq IRequest, resp *lreq.Response) lserr.IError {
	headers := preq.GetMoreHeaders()

	// Remove sensitive information
	if headers != nil {
		delete(headers, "Authorization")
	}

	return lserr.ErrorHandler(perr).WithKVparameters(
		"statusCode", resp.StatusCode,
		"url", purl,
		"method", preq.GetRequestMethod(),
		"requestHeaders", headers,
		"responseHeaders", resp.Header,
	)
}
