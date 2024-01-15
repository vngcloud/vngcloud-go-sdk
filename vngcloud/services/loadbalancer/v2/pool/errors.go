package pool

import (
	"fmt"
	lStr "strings"

	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
)

const (
	errPoolUnchangeIdentifiers = "the members provided are identical to the existing members in the pool"
)

// *********************************************** ErrorResult for pool ************************************************

type ErrorResolver struct {
	Message string `json:"message"`
}

func (s *ErrorResolver) ToError() error {
	msg := lStr.TrimSpace(lStr.ToLower(s.Message))
	switch msg {
	case errPoolUnchangeIdentifiers:
		return NewErrPoolMemberUnchanged()
	}

	return fmt.Errorf(s.Message)
}

// ************************************************* ErrPoolNotFound ***************************************************

func NewErrPoolInUse(pPoolID, pInfo string) vconError.IErrorBuilder {
	err := new(ErrPoolInUse)
	err.PoolID = pPoolID
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrPoolInUse(pErr error) bool {
	_, ok := pErr.(*ErrPoolInUse)
	return ok
}

type ErrPoolInUse struct {
	PoolID string
	vconError.BaseError
}

func (s *ErrPoolInUse) Error() string {
	s.DefaultError = fmt.Sprintf("pool %s is in use by listener", s.PoolID)
	return s.ChoseErrString()
}

// ********************************************** ErrPoolMemberUnchanged ***********************************************

func NewErrPoolMemberUnchanged() vconError.IErrorBuilder {
	err := new(ErrPoolMemberUnchanged)
	return err
}

func IsErrPoolMemberUnchanged(pErr error) bool {
	_, ok := pErr.(*ErrPoolMemberUnchanged)
	return ok
}

type ErrPoolMemberUnchanged struct {
	vconError.BaseError
}

func (s *ErrPoolMemberUnchanged) Error() string {
	s.DefaultError = errPoolUnchangeIdentifiers
	return s.ChoseErrString()
}
