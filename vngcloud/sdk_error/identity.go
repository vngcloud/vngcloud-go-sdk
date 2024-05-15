package sdk_error

import (
	lfmt "fmt"
)

const (
	EcAuthenticationFailed        = ErrorCode("AuthenticationFailed")
	EcUnknownAuthenticationFailed = ErrorCode("UnknownAuthenticationFailed")
	EcReauthFuncNotSet            = ErrorCode("ReauthFuncNotSet")
)

func WithErrorAuthenticationFailed(perr error, perrResp IErrorRespone) func(ISdkError) {
	return func(sdkErr ISdkError) {
		if perrResp == nil {
			return
		}

		if perrResp != nil && len(perrResp.GetErrors()) < 1 {
			return
		}

		for _, err := range perrResp.GetErrors() {
			if err.Code == "AUTHENTICATION_FAILED" {
				sdkErr.WithErrorCode(EcAuthenticationFailed).
					WithMessage(err.Message).
					WithErrors(lfmt.Errorf("%s: %s", err.Code, err.Message))
				return
			}
		}

		sdkErr.WithErrorCode(EcUnknownAuthenticationFailed).
			WithMessage(perr.Error()).
			WithErrors(perr)
	}
}

func WithErrorReauthFuncNotSet() func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcReauthFuncNotSet).
			WithMessage("Reauthentication function is not configured").
			WithErrors(lfmt.Errorf("Reauthentication function is not configured"))
	}
}
