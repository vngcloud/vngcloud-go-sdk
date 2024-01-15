package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/imroc/req/v3"

	"github.com/vngcloud/vngcloud-go-sdk/error/utils"
)

type (
	ProviderClient struct {
		IdentityEndpoint string
		AccessToken      string
		HTTPClient       *req.Client
		ReauthFunc       func() error
		ThrowAway        bool
		Context          context.Context

		mut       *sync.RWMutex
		reauthmut *reauthlock
	}

	RequestOpts struct {
		JSONBody     interface{}
		JSONResponse interface{}
		JSONError    interface{}
		OkCodes      []int
		MoreHeaders  map[string]string
		OmitHeaders  []string
	}

	reauthlock struct {
		sync.RWMutex
		ongoing *reauthFuture
	}

	reauthFuture struct {
		done chan struct{}
		err  error
	}
)

func (s *ProviderClient) UseTokenLock() {
	s.mut = new(sync.RWMutex)
	s.reauthmut = new(reauthlock)
}

func (s *ProviderClient) UseHTTPClient() {
	s.HTTPClient = req.NewClient().
		SetCommonRetryCount(noRetries).
		SetCommonRetryFixedInterval(retryInterval)
}

func (s *ProviderClient) SetThrowaway(pIsThrowAway bool) {
	if s.reauthmut != nil {
		s.reauthmut.Lock()
		defer s.reauthmut.Unlock()
	}
	s.ThrowAway = pIsThrowAway
}

func (s *ProviderClient) IsThrowAway() bool {
	if s.reauthmut != nil {
		s.reauthmut.RLock()
		defer s.reauthmut.RUnlock()
	}

	return s.ThrowAway
}

func (s *ProviderClient) Reauthenticate() error {
	s.SetThrowaway(true)
	defer s.SetThrowaway(false)

	if s.ReauthFunc == nil {
		return NewErrEmptyReauthFunction("")
	}

	s.reauthmut.Lock()
	ongoing := s.reauthmut.ongoing
	if ongoing == nil {
		s.reauthmut.ongoing = newReauthFuture()
	}
	s.reauthmut.Unlock()

	if ongoing != nil {
		return ongoing.Get()
	}

	err := s.ReauthFunc()
	s.reauthmut.Lock()
	s.reauthmut.ongoing.Set(err)
	s.reauthmut.ongoing = nil
	s.reauthmut.Unlock()

	return err
}

func (s *ProviderClient) SetAccessToken(pNewAccessToken string) {
	if s.mut != nil {
		s.mut.Lock()
		defer s.mut.Unlock()
	}

	s.AccessToken = pNewAccessToken
}

func (s *ProviderClient) getDefaultHeaders() map[string]string {
	authorization := "Bearer " + s.AccessToken
	contentType := "application/json"
	return map[string]string{
		"Authorization": authorization,
		"Content-Type":  contentType,
	}
}

func (s *ProviderClient) doRequest(pMethod, pUrl string, pOpts *RequestOpts) (*req.Response, error) {
	request := s.HTTPClient.R().SetContext(s.Context).SetHeaders(s.getDefaultHeaders()).SetHeaders(pOpts.MoreHeaders)

	if pOpts.JSONBody != nil {
		request.SetBodyJsonMarshal(pOpts.JSONBody)
	}

	if pOpts.JSONResponse != nil {
		request.SetSuccessResult(pOpts.JSONResponse)
	}

	if pOpts.JSONError != nil {
		request.SetErrorResult(pOpts.JSONError)
	}

	var resp *req.Response
	var err error
	switch strings.ToUpper(pMethod) {
	case "POST":
		resp, err = request.Post(pUrl)
	case "GET":
		resp, err = request.Get(pUrl)
	case "DELETE":
		resp, err = request.Delete(pUrl)
	case "PUT":
		resp, err = request.Put(pUrl)
	}

	if err != nil {
		return resp, err
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		if s.ReauthFunc != nil {
			err = s.Reauthenticate()
			if err != nil {
				return nil, NewErrReauthFailed(err.Error())
			}

			return s.doRequest(pMethod, pUrl, pOpts)
		}
	case http.StatusTooManyRequests:
		return nil, utils.NewErrTooManyRequests(fmt.Sprintf("too many requests [%s %s] are made to server", pMethod, pUrl))
	}

	for _, sttc := range pOpts.OkCodes {
		if sttc == resp.StatusCode {
			return resp, nil
		}
	}

	return resp, utils.NewErrUnknown(strings.TrimSpace(fmt.Sprintf("%+v", resp)))
}

func (s *ProviderClient) Request(pMethod, pUrl string, pOpts *RequestOpts) (*req.Response, error) {
	return s.doRequest(pMethod, pUrl, pOpts)
}

func newReauthFuture() *reauthFuture {
	return &reauthFuture{
		done: make(chan struct{}),
		err:  nil,
	}
}

func (f *reauthFuture) Get() error {
	<-f.done
	return f.err
}

func (f *reauthFuture) Set(err error) {
	f.err = err
	close(f.done)
}
