package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternNetworkNotFound                  = "is not found"
	patternSubnetNotFound                   = `subnet with id [^.]+ is not found`
	patternSubnetNotBelongNetwork           = `subnet id: [^.]+ belong to network id: [^.]+ not found`
	patternInternalNetworkInterfaceNotFound = `interface network interface with id [^.]+ is not found`
	patternWanIpAvailable                   = "wan ip is available"
	pattermWapIdNotFound                    = "cannot get wan ip with id"
)

var (
	regexErrorSubnetNotFound                   = lregexp.MustCompile(patternSubnetNotFound)
	regexErrorSubnetNotBelongNetwork           = lregexp.MustCompile(patternSubnetNotBelongNetwork)
	regexErrorInternalNetworkInterfaceNotFound = lregexp.MustCompile(patternInternalNetworkInterfaceNotFound)
)

func WithErrorNetworkNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternNetworkNotFound) ||
			lstr.ToUpper(lstr.TrimSpace(perrResp.GetError().Error())) == "VPC_IS_NOT_FOUND" {
			sdkError.WithErrorCode(EcVServerNetworkNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSubnetNotBelongNetwork(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorSubnetNotBelongNetwork.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerSubnetNotBelongNetwork).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSubnetNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorSubnetNotFound.FindString(errMsg) != "" ||
			lstr.ToUpper(lstr.TrimSpace(perrResp.GetError().Error())) == "SUBNET_IS_NOT_FOUND" {
			sdkError.WithErrorCode(EcVServerSubnetNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorInternalNetworkInterfaceNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorInternalNetworkInterfaceNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerInternalNetworkInterfaceNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorWanIpAvailable(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternWanIpAvailable) {
			sdkError.WithErrorCode(EcVServerWanIpAvailable).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorWanIdNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), pattermWapIdNotFound) {
			sdkError.WithErrorCode(EcVServerWanIdNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
