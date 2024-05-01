package errors

import (
	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
	lstr "strings"
)

const (
	patternOutOfPoc = "you do not have sufficient credits to complete the purchase"
)

var (
	ErrCodeOutOfPoc lssdkErr.ErrorCode = "ErrorOutOfPoc"
)

func WithErrorOutOfPoc(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternOutOfPoc) {
			sdkError.Code = ErrCodeOutOfPoc
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
