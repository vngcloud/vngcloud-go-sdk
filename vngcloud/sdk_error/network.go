package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternNetworkNotFound = "is not found"
	patternSubnetNotFound  = `subnet with id [^.]+ is not found`
)

var (
	regexErrorSubnetNotFound = lregexp.MustCompile(patternSubnetNotFound)
)

func WithErrorNetworkNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternNetworkNotFound) {
			sdkError.WithErrorCode(EcVServerNetworkNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSubnetNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorSubnetNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerSubnetNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
