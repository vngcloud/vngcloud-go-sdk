package sdk_error

import (
	lfmt "fmt"
)

const (
	// Internal SDK error
	EcReauthFuncNotSet = ErrorCode("SdkReauthFunctionNotSet")

	// Error from VngCloudApi
	EcAuthenticationFailed = ErrorCode("VngCloudIamAuthenticationFailed")
	EcTooManyFailedLogins  = ErrorCode("VngCloudIamTooManyFailedLogins")
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
