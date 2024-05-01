package errors

import (
	lfmt "fmt"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

var (
	ErrCodeBaseUnknown lssdkErr.ErrorCode = "ErrorBaseUnknown"
	ErrUnknown                            = lfmt.Errorf("unknown error")
	ErrUnknownMessage                     = "Unknown error"
)

func ErrorHandler(perr error, popts ...func(response *lssdkErr.SdkError)) *lssdkErr.SdkError {
	errRes := &lssdkErr.SdkError{
		Code:    ErrCodeBaseUnknown,
		Error:   ErrUnknown,
		Message: ErrUnknownMessage,
	}

	for _, opt := range popts {
		opt(errRes)
		if errRes.Code != ErrCodeBaseUnknown {
			return errRes
		}
	}

	errRes.Error = perr
	errRes.Message = perr.Error()

	return errRes
}
