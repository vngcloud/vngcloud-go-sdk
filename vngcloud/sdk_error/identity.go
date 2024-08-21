package sdk_error

import (
	lfmt "fmt"
)

func WithErrorAuthenticationFailed(perrResp IErrorRespone) func(IError) {
	return func(sdkErr IError) {
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

func WithErrorReauthFuncNotSet() func(IError) {
	return func(sdkErr IError) {
		sdkErr.WithErrorCode(EcReauthFuncNotSet).
			WithMessage("Reauthentication function is not configured").
			WithErrors(lfmt.Errorf("reauthentication function is not configured"))
	}
}

func WithErrorTooManyFailedLogin(perrResp IErrorRespone) func(IError) {
	return func(sdkErr IError) {
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

func WithErrorUnknownAuthFailure(perrResp IErrorRespone) func(IError) {
	return func(sdkErr IError) {
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
