package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternVirtualAddressNotFound = `virtual ip address with id [^.]+ is not found`
)

var (
	regexErrorVirtualAddressNotFound = lregexp.MustCompile(patternVirtualAddressNotFound)
)

func WithErrorVirtualAddressNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorVirtualAddressNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVirtualAddressNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

// Virtual Ip Address with id vip-0d2402cf-49e8-43bf-abbe-b707597320e0 is not found
