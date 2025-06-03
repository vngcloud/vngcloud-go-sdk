package sdk_error

import lstr "strings"

const (
	patternSecgroupRuleNotFound    = "cannot get security group rule with id"
	patternSecgroupRuleExists      = "securitygroupruleexists"
	patternSecgroupRuleExceedQuota = "exceeded secgroup_rule quota"
)

func WithErrorSecgroupRuleNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorSecgroupRuleAlreadyExists(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupRuleExists) {
			sdkError.WithErrorCode(EcVServerSecgroupRuleAlreadyExists).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSecgroupRuleExceedQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupRuleExceedQuota) {
			sdkError.WithErrorCode(EcVServerSecgroupRuleExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}
