package client

import (
	lbytes "bytes"
	ljson "encoding/json"
	lio "io"
	lhttp "net/http"
	lsync "sync"
	ltime "time"

	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

const (
	IamOauth2 AuthOpts = "IamOauth2"
)

type (
	httpClient struct {
		retryCount int
		delay      ltime.Duration
		sleep      ltime.Duration
		client     *lhttp.Client

		reauthFunc   func() (ISdkAuthentication, lserr.ISdkError)
		reauthOption AuthOpts

		accessToken    ISdkAuthentication
		defaultHeaders map[string]string

		mut       *lsync.RWMutex
		reauthMut *reauthLock
		throwAway bool
	}

	reauthLock struct {
		ongoing *reauthFuture

		lsync.RWMutex
	}

	reauthFuture struct {
		done   chan struct{}
		sdkErr lserr.ISdkError
	}

	AuthOpts string
)

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

func (s *httpClient) WithReauthFunc(pauthOpt AuthOpts, preauthFunc func() (ISdkAuthentication, lserr.ISdkError)) IHttpClient {
	s.reauthFunc = preauthFunc
	s.reauthOption = pauthOpt
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

	// So I need to reauth here
	if s.needReauth(preq) {
		auth, sdkErr := s.reauthenticate()
		if sdkErr != nil {
			return sdkErr, true
		}

		s.accessToken = auth
		s.WithKvDefaultHeaders("Authorization", "Bearer "+s.accessToken.GetAccessToken())
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
	} else {
		switch resp.StatusCode {
		case lhttp.StatusUnauthorized:

		}
	}

	re := preq.GetJsonError()
	if err = ljson.Unmarshal(respBody, re); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorCanNotUnmarshalResponseBody(err)), true
	}

	preq.SetJsonError(re)
	return lserr.ErrorHandler(err, lserr.WithErrorOkCodeNotMatch(resp.StatusCode)), true
}

func (s *httpClient) needReauth(preq IRequest) bool {
	if preq.SkipAuthentication() {
		return false
	}

	return s.accessToken.NeedReauth()
}

func (s *httpClient) reauthenticate() (ISdkAuthentication, lserr.ISdkError) {
	s.setThrowaway(true)
	defer s.setThrowaway(false)

	if s.reauthFunc == nil {
		return nil, lserr.ErrorHandler(nil, lserr.WithErrorReauthFuncNotSet())
	}

	s.reauthMut.Lock()
	ongoing := s.reauthMut.ongoing
	if ongoing == nil {
		s.reauthMut.ongoing = newReauthFuture()
	}
	s.reauthMut.Unlock()

	if ongoing != nil {
		return nil, ongoing.Get()
	}

	auth, sdkerr := s.reauthFunc()
	s.reauthMut.Lock()
	s.reauthMut.ongoing.Set(sdkerr)
	s.reauthMut.ongoing = nil
	s.reauthMut.Unlock()

	return auth, sdkerr
}

func (s *httpClient) setThrowaway(pisThrowAway bool) {
	if s.reauthMut != nil {
		s.reauthMut.Lock()
		defer s.reauthMut.Unlock()
	}
	s.throwAway = pisThrowAway
}

func newReauthFuture() *reauthFuture {
	return &reauthFuture{
		done:   make(chan struct{}),
		sdkErr: nil,
	}
}

func (s *reauthFuture) Get() lserr.ISdkError {
	<-s.done
	return s.sdkErr
}

func (s *reauthFuture) Set(perr lserr.ISdkError) {
	s.sdkErr = perr
	close(s.done)
}
