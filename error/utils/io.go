package utils

import (
	"fmt"

	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
)

// ************************************************** ErrMissingInput **************************************************

func NewErrMissingInput(pArgument, pInfo string) vconError.IErrorBuilder {
	err := new(ErrMissingInput)
	err.Argument = pArgument
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrMissingInput(pErr error) bool {
	_, ok := pErr.(*ErrMissingInput)
	return ok
}

type ErrMissingInput struct {
	Argument string

	vconError.BaseError
}

func (s *ErrMissingInput) Error() string {
	s.DefaultError = fmt.Sprintf("missing input for argument [%s]", s.Argument)
	return s.ChoseErrString()
}

// *************************************************** ErrEmptyConfig **************************************************

func NewErrEmptyConfig(pArgument, pInfo string) vconError.IErrorBuilder {
	err := new(ErrMissingInput)
	err.Argument = pArgument
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrEmptyConfig(pErr error) bool {
	_, ok := pErr.(*ErrEmptyConfig)
	return ok
}

type ErrEmptyConfig struct {
	Argument string

	vconError.BaseError
}

func (s *ErrEmptyConfig) Error() string {
	s.DefaultError = fmt.Sprintf("missing input for argument [%s]", s.Argument)
	return s.ChoseErrString()
}
