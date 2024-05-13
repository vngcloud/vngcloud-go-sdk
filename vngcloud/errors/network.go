package errors

import (
	lregexp "regexp"
	lstr "strings"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

const (
	patternNetworkNotFound = `network with id [^.]+ is not found`
)

var (
	ErrCodeNetworkNotFound lssdkErr.ErrorCode = "ErrorNetworkNotFound"
)

var (
	regexErrorNetworkNotFound = lregexp.MustCompile(patternNetworkNotFound)
)

func WithErrorNetworkNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.Message))
		if regexErrorNetworkNotFound.FindString(errMsg) != "" {
			sdkError.Code = ErrCodeNetworkNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
