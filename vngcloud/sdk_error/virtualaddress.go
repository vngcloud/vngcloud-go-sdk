package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternVirtualAddressNotFound    = `virtual ip address with id [^.]+ is not found`
	patternAddressPairNotFound       = `address pair with uuid: [^.]+ was not existing`
	patternVirtualAddressExceedQuota = "exceeded virtual_ip_address quota"
	patternVirtualAddressInUse       = "ip address is already in use"
	patternVirtualAddressInUse2      = `virtual ip address id [^.]+ is used. please remove address pairs first`
)

var (
	regexErrorVirtualAddressNotFound = lregexp.MustCompile(patternVirtualAddressNotFound)
	regexErrorAddressPairNotFound    = lregexp.MustCompile(patternAddressPairNotFound)
	regexErrorVirtualAddressInUse    = lregexp.MustCompile(patternVirtualAddressInUse)
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

func WithErrorVirtualAddressInUse(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if lstr.Contains(errMsg, patternVirtualAddressInUse) ||
			regexErrorVirtualAddressInUse.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVirtualAddressInUse).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}
