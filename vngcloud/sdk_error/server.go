package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternServerNotFound    = "cannot get server with id" // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
	patternServerExceedQuota = "exceeded vm quota"         // "The number of servers exceeds the quota"
	patternNotAllowDelete    = `cannot delete server with status (creating|deleting|creating-billing)`
)

var (
	regexErrorServerNotAllowDeleteServer = lregexp.MustCompile(patternNotAllowDelete)
)

func WithErrorServerNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerNotFound) {
			sdkError.WithErrorCode(EcVServerServerNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerNotAllowDeleteServer(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if regexErrorServerNotAllowDeleteServer.FindString(lstr.ToLower(lstr.TrimSpace(errMsg))) != "" {
			sdkError.WithErrorCode(EcVServerServerNotAllowDeleteServer).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerExceedQuota(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerExceedQuota) {
			sdkError.WithErrorCode(EcVServerServerExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
