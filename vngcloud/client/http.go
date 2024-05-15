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

	accessToken string
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

func (s *httpClient) WithReauthFunc(preauthFunc func() (ISdkAuthentication, lserr.ISdkError)) IHttpClient {
	s.reauthFunc = preauthFunc
	return s
}

func (s *httpClient) DoRequest(purl string, preq IRequest) lserr.ISdkError {
	var (
		err        error
		req        *lhttp.Request
		resp       *lhttp.Response
		marshalled []byte
		respBody   []byte
	)

	if marshalled, err = ljson.Marshal(preq.GetRequestBody()); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorCanNotMarshalRequestBody(err))
	}

	if req, err = lhttp.NewRequest(preq.GetRequestMethod(), purl, lbytes.NewReader(marshalled)); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorFailedToCreateHttpRequest(err))
	}

	if preq.GetMoreHeaders() != nil {
		for k, v := range preq.GetMoreHeaders() {
			req.Header.Add(k, v)
		}
	}

	if preq.GetOmitHeaders() != nil {
		for k := range preq.GetOmitHeaders().Iter() {
			req.Header.Del(k)
		}
	}

	if resp, err = s.client.Do(req); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorFailedToMakeHttpRequest(err))
	}
	defer resp.Body.Close()

	if respBody, err = lio.ReadAll(resp.Body); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorFailedToReadResponseBody(err))
	}

	if preq.ContainsOkCode(resp.StatusCode) {
		rr := preq.GetJsonResponse()
		if err = ljson.Unmarshal(respBody, rr); err != nil {
			return lserr.ErrorHandler(err, lserr.WithErrorCanNotUnmarshalResponseBody(err))
		}

		preq.SetJsonResponse(rr)
		return nil
	}

	re := preq.GetJsonError()
	if err = ljson.Unmarshal(respBody, re); err != nil {
		return lserr.ErrorHandler(err, lserr.WithErrorCanNotUnmarshalResponseBody(err))
	}
	preq.SetJsonError(re)

	return lserr.ErrorHandler(err, lserr.WithErrorOkCodeNotMatch(resp.StatusCode))
}
