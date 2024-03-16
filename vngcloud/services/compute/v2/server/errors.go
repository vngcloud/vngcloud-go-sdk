package server

import (
	"fmt"
	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
	"strings"
)

const (
	patternErrNotFound = "Cannot get server with id ins-"
)

var (
	errNotFound = new(ErrNotFound)
)

type ErrNotFound struct {
	vconError.BaseError
}

func (s *ErrNotFound) Error() string {
	s.DefaultError = "server not found"
	return s.ChoseErrString()
}

func IsErrNotFound(pErr error) bool {
	_, ok := pErr.(*ErrNotFound)
	return ok
}

type UnknowError struct {
	vconError.BaseError
}

func (s *UnknowError) Error() string {
	s.DefaultError = "unknown error"
	return s.ChoseErrString()
}

// ********************************************* UTILS *********************************************

func resolveError(perrResp *vconError.ErrorResponse, perr error) error {
	if perr == nil {
		return nil
	}

	if perrResp != nil {
		if strings.Contains(perrResp.Message, patternErrNotFound) {
			return errNotFound
		}

		unknowErr := new(UnknowError)
		unknowErr.Info = fmt.Sprintf("unknow error: %v", perrResp.Message)
		return unknowErr
	}

	return perr
}
