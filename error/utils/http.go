package utils

import (
	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
)

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
