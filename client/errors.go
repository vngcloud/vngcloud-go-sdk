package client

import vconError "github.com/vngcloud/vngcloud-go-sdk/error"

// ********************************************** ErrEmptyReauthFunction ***********************************************

func NewErrEmptyReauthFunction(pInfo string) vconError.IErrorBuilder {
	err := new(ErrEmptyReauthFunction)
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrEmptyRequestOption(pErr error) bool {
	_, ok := pErr.(*ErrEmptyReauthFunction)
	return ok
}

type ErrEmptyReauthFunction struct {
	vconError.BaseError
}

func (s *ErrEmptyReauthFunction) Error() string {
	s.DefaultError = "re-authentication function SHOULD NOT be nil"
	return s.ChoseErrString()
}

// ************************************************** ErrReauthFailed **************************************************

func NewErrReauthFailed(pInfo string) vconError.IErrorBuilder {
	err := new(ErrReauthFailed)
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrReauthFailed(pErr error) bool {
	_, ok := pErr.(*ErrReauthFailed)
	return ok
}

type ErrReauthFailed struct {
	vconError.BaseError
}

func (s *ErrReauthFailed) Error() string {
	s.DefaultError = "re-authentication failed"
	return s.ChoseErrString()
}
