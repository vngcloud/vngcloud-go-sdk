package sdk_error

import (
	lfmt "fmt"
)

func WithErrorAuthenticationFailed(perrResp IErrorRespone) func(ISdkError) {
	return func(sdkErr ISdkError) {
		if perrResp == nil {
			return
		}

		if perrResp.GetError() == nil {
			return
		}

		if perrResp.GetError().Error() == "AUTHENTICATION_FAILED" {
			sdkErr.WithErrorCode(EcAuthenticationFailed).
				WithErrors(perrResp.GetError()).
				WithMessage(perrResp.GetMessage())
		}
	}
}

func WithErrorReauthFuncNotSet() func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcReauthFuncNotSet).
			WithMessage("Reauthentication function is not configured").
			WithErrors(lfmt.Errorf("reauthentication function is not configured"))
	}
}

func WithErrorTooManyFailedLogin(perrResp IErrorRespone) func(ISdkError) {
	return func(sdkErr ISdkError) {
		if perrResp == nil {
			return
		}

		if perrResp.GetError() == nil {
			return
		}

		if perrResp.GetError().Error() == "TOO_MANY_FAILED_LOGINS" {
			sdkErr.WithErrorCode(EcTooManyFailedLogins).
				WithErrors(perrResp.GetError()).
				WithMessage(perrResp.GetMessage())
		}
	}
}

func WithErrorUnknownAuthFailure(perrResp IErrorRespone) func(ISdkError) {
	return func(sdkErr ISdkError) {
		if perrResp == nil {
			return
		}

		if perrResp.GetError() == nil {
			return
		}

		if sdkErr.GetErrorCode() == EcUnknownError {
			sdkErr.WithErrorCode(EcUnknownAuthFailure).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}
