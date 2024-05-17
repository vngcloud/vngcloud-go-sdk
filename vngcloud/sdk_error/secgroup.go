package sdk_error

import lstr "strings"

const (
	patternSecgroupNotFound          = "cannot get security group with id"
	patternSecgroupNameAlreadyExists = "name of security group already exist"
	patternSecgroupExceedQuota       = "exceeded secgroup quota"
	patternSecgroupInUse             = "securitygroupinuse"
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

func WithErrorSecgroupNameAlreadyExists(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupNameAlreadyExists) {
			sdkError.WithErrorCode(EcVServerSecgroupNameAlreadyExists).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSecgroupExceedQuota(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupExceedQuota) {
			sdkError.WithErrorCode(EcVServerSecgroupExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSecgroupInUse(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupInUse) {
			sdkError.WithErrorCode(EcVServerSecgroupInUse).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
