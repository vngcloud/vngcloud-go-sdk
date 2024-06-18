package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternNetworkNotFound        = "is not found"
	patternSubnetNotFound         = `subnet with id [^.]+ is not found`
	patternSubnetNotBelongNetwork = `subnet id: [^.]+ belong to network id: [^.]+ not found`
)

var (
	regexErrorSubnetNotFound         = lregexp.MustCompile(patternSubnetNotFound)
	regexErrorSubnetNotBelongNetwork = lregexp.MustCompile(patternSubnetNotBelongNetwork)
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

func WithErrorSubnetNotBelongNetwork(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
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
