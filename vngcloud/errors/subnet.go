package errors

import (
	lregexp "regexp"
	lstr "strings"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

const (
	patternSubnetNotFound = `subnet with id [^.]+ is not found`
)

var (
	ErrCodeSubnetNotFound lssdkErr.ErrorCode = "ErrorSubnetNotFound"
)

var (
	regexErrorSubnetNotFound = lregexp.MustCompile(patternSubnetNotFound)
)

func WithErrorSubnetNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.Message))
		if regexErrorSubnetNotFound.FindString(errMsg) != "" {
			sdkError.Code = ErrCodeSubnetNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
