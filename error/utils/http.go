package utils

import (
	"fmt"

	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
)

// ********************************************** ErrUnexpectedStatusCode **********************************************

func NewErrUnexpectedStatusCode(pURL, pMethod string, pExpected []int, pActual int, pInfo string) vconError.IErrorBuilder {
	err := new(ErrUnexpectedStatusCode)
	err.URL = pURL
	err.Method = pMethod
	err.Expected = pExpected
	err.Actual = pActual
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrUnexpectedStatusCode(pErr error) bool {
	_, ok := pErr.(*ErrUnexpectedStatusCode)
	return ok
}

type ErrUnexpectedStatusCode struct {
	URL      string
	Method   string
	Expected []int
	Actual   int

	vconError.BaseError
}

func (s *ErrUnexpectedStatusCode) Error() string {
	s.DefaultError = fmt.Sprintf("expected HTTP status code [%v] when accessing [%s %s], but got %d instead\n[%s]",
		s.Expected, s.Method, s.URL, s.Actual, s.Info)
	return s.ChoseErrString()
}

// ********************************************** ErrEmptyRequestOption ************************************************

func NewErrEmptyRequestOption(pInfo string) vconError.IErrorBuilder {
	err := new(ErrEmptyRequestOption)
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrEmptyRequestOption(pErr error) bool {
	_, ok := pErr.(*ErrEmptyRequestOption)
	return ok
}

type ErrEmptyRequestOption struct {
	vconError.BaseError
}

func (s *ErrEmptyRequestOption) Error() string {
	s.DefaultError = "request options SHOULD NOT be nil"
	return s.ChoseErrString()
}

// ************************************************* ErrTooManyRequests ************************************************

func NewErrTooManyRequests(pInfo string) vconError.IErrorBuilder {
	err := new(ErrTooManyRequests)
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrTooManyRequests(pErr error) bool {
	_, ok := pErr.(*ErrTooManyRequests)
	return ok
}

type ErrTooManyRequests struct {
	vconError.BaseError
}

func (s *ErrTooManyRequests) Error() string {
	s.DefaultError = "too many requests are being made to the server"
	return s.ChoseErrString()
}
