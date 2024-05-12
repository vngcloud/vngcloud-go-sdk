package errors

import (
	lstr "strings"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

const (
	patternPoolNotFound     = "cannot get pool with id" // "Cannot get pool with id lb-af85f586-5d71-435c-8fe4-ea0c0dc26655"
	patternPoolInUse        = "is used in listener"
	patternResourceNotFound = "could not find resource" // "Could not find resource"

)

var (
	ErrCodePoolNotFound lssdkErr.ErrorCode = "ErrorPoolNotFound"
	ErrCodePoolInUse    lssdkErr.ErrorCode = "ErrorPoolInUse"
)

func WithErrorPoolNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternPoolNotFound) ||
			lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternResourceNotFound) {
			sdkError.Code = ErrCodePoolNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}

func WithErrorPoolInUse(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternPoolInUse) {
			sdkError.Code = ErrCodePoolInUse
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
