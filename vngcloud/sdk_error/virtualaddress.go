package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternVirtualAddressNotFound    = `virtual ip address with id [^.]+ is not found`
	patternAddressPairNotFound       = `address pair with uuid: [^.]+ was not existing`
	patternVirtualAddressExceedQuota = "exceeded virtual_ip_address quota"
)

var (
	regexErrorVirtualAddressNotFound = lregexp.MustCompile(patternVirtualAddressNotFound)
	regexErrorAddressPairNotFound    = lregexp.MustCompile(patternAddressPairNotFound)
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

func WithErrorAddressPairNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorAddressPairNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVirtualAddressNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVirtualAddressExceedQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVirtualAddressExceedQuota) {
			sdkError.WithErrorCode(EcVServerVirtualAddressExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}