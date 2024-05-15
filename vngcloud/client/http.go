package client

import (
	lbytes "bytes"
	ljson "encoding/json"
	lio "io"
	lhttp "net/http"
	ltime "time"

	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type httpClient struct {
	retryCount int
	delay      ltime.Duration
	sleep      ltime.Duration
	client     *lhttp.Client

	reauthFunc func() (ISdkAuthentication, lserr.ISdkError)

	accessToken    string
	defaultHeaders map[string]string
}

func NewHttpClient() IHttpClient {
	return &httpClient{
		retryCount: 0,
		client: &lhttp.Client{
			Timeout: ltime.Second * 30,
		},
	}
}

func (s *httpClient) WithRetryCount(pretryCount int) IHttpClient {
	s.retryCount = pretryCount
	return s
}

func (s *httpClient) WithDelay(pdelay ltime.Duration) IHttpClient {
	s.delay = pdelay
	return s
}

func (s *httpClient) WithTimeout(ptimeout ltime.Duration) IHttpClient {
	s.client.Timeout = ptimeout
	return s
}

func (s *httpClient) WithSleep(psleep ltime.Duration) IHttpClient {
	s.sleep = psleep
	return s
}

func (s *httpClient) WithKvDefaultHeaders(pargs ...string) IHttpClient {
	if s.defaultHeaders == nil {
		s.defaultHeaders = make(map[string]string)
	}

	if len(pargs)%2 != 1 {
		pargs = append(pargs, "")
	}

	for i := 0; i < len(pargs); i += 2 {
		s.defaultHeaders[pargs[i]] = pargs[i+1]
	}

	return s
}

func (s *httpClient) WithReauthFunc(preauthFunc func() (ISdkAuthentication, lserr.ISdkError)) IHttpClient {
	s.reauthFunc = preauthFunc
	return s
}

func (s *httpClient) DoRequest(purl string, preq IRequest) lserr.ISdkError {
	var (
		sdkerr lserr.ISdkError
		retry  = true
	)

	if s.delay > 0 {
		ltime.Sleep(s.delay * ltime.Second)
	}

	for i := 0; i <= s.retryCount; i++ {
		sdkerr, retry = s.doRequest(purl, preq)
		if sdkerr == nil || !retry {
			break
		}

		if s.sleep > 0 {
			ltime.Sleep(s.sleep * ltime.Second)
		}
	}

	return sdkerr
}

func (s *httpClient) doRequest(purl string, preq IRequest) (lserr.ISdkError, bool) {
	var (
		err        error
		req        *lhttp.Request
		resp       *lhttp.Response
		marshalled []byte
		respBody   []byte
	)

	if marshalled, err = ljson.Marshal(preq.GetRequestBody()); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorCanNotMarshalRequestBody(err)), false
	}

	if req, err = lhttp.NewRequest(preq.GetRequestMethod(), purl, lbytes.NewReader(marshalled)); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorFailedToCreateHttpRequest(err)), false
	}

	// Add the default headers to the request
	if s.defaultHeaders != nil {
		for k, v := range s.defaultHeaders {
			req.Header.Add(k, v)
		}
	}

	// Add the more headers to the request
	if preq.GetMoreHeaders() != nil {
		for k, v := range preq.GetMoreHeaders() {
			req.Header.Add(k, v)
		}
	}

	// Omit the headers from the request
	if preq.GetOmitHeaders() != nil {
		for k := range preq.GetOmitHeaders().Iter() {
			req.Header.Del(k)
		}
	}

	if resp, err = s.client.Do(req); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorFailedToMakeHttpRequest(err)), true
	}
	defer resp.Body.Close()

	if respBody, err = lio.ReadAll(resp.Body); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorFailedToReadResponseBody(err)), true
	}

	if preq.ContainsOkCode(resp.StatusCode) {
		rr := preq.GetJsonResponse()
		if err = ljson.Unmarshal(respBody, rr); err != nil {
			return lserr.ErrorHandler(err, lserr.WithErrorCanNotUnmarshalResponseBody(err)), true
		}

		preq.SetJsonResponse(rr)
		return nil, false
	}

	re := preq.GetJsonError()
	if err = ljson.Unmarshal(respBody, re); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorCanNotUnmarshalResponseBody(err)), true
	}

	preq.SetJsonError(re)
	return lserr.ErrorHandler(err, lserr.WithErrorOkCodeNotMatch(resp.StatusCode)), true
}
