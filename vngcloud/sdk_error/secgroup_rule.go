package sdk_error

import lstr "strings"

const (
	patternSecgroupRuleNotFound = "cannot get security group rule with id"
)

func WithErrorSecgroupRuleNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupRuleNotFound) {
			sdkError.WithErrorCode(EcVServerSecgroupRuleNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
