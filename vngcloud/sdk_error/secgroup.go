package sdk_error

import lstr "strings"

const (
	patternSecgroupNotFound = "cannot get security group with id"
)

func WithErrorSecgroupNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupNotFound) {
			sdkError.WithErrorCode(EcVServerSecgroupNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
