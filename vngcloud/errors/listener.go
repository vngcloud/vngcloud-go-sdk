package errors

import (
	lstr "strings"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

const (
	patternListenerNotFound = "cannot get listener with id"
)

var (
	ErrCodeListenerNotFound lssdkErr.ErrorCode = "ErrorListenerNotFound"
)

func WithErrorListenerNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternListenerNotFound) {
			sdkError.Code = ErrCodeListenerNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
