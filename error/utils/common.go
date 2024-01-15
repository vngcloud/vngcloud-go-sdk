package utils

import (
	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
	"strings"
)

// **************************************************** ErrUnknown *****************************************************

func NewErrUnknown(pInfo string) vconError.IErrorBuilder {
	err := new(ErrUnknown)
	if pInfo != "" {
		err.Info = strings.TrimSpace(pInfo)
	}
	return err
}

func IsErrUnknown(pErr error) bool {
	_, ok := pErr.(*ErrUnknown)
	return ok
}

type ErrUnknown struct {
	vconError.BaseError
}

func (s *ErrUnknown) Error() string {
	s.DefaultError = "unknown error, something went wrong"
	return s.ChoseErrString()
}

// ************************************************ ErrInvalidEndpoint *************************************************

func NewErrInvalidEndpoint(pInfo string) vconError.IErrorBuilder {
	err := new(ErrInvalidEndpoint)
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrInvalidEndpoint(pErr error) bool {
	_, ok := pErr.(*ErrInvalidEndpoint)
	return ok
}

type ErrInvalidEndpoint struct {
	vconError.BaseError
}

func (s *ErrInvalidEndpoint) Error() string {
	s.DefaultError = "invalid endpoint URL"
	return s.ChoseErrString()
}
