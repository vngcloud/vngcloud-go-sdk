package sdk_error

import (
	lfmt "fmt"
	lstr "strings"
)

const (
	patternOutOfPoc = "you do not have sufficient credits to complete the purchase"
)

func ErrorHandler(perr error, popts ...func(psdkErr ISdkError)) ISdkError {
	sdkErr := &SdkError{
		error:     perr,
		errorCode: EcUnknownError,
		message:   "Unknown error",
	}

	for _, opt := range popts {
		opt(sdkErr)
		if sdkErr.errorCode != EcUnknownError {
			return sdkErr
		}
	}

	sdkErr.error = perr
	sdkErr.message = ""

	return sdkErr
}

func SdkErrorHandler(psdkErr ISdkError, perrResp IErrorRespone, popts ...func(psdkErr ISdkError)) ISdkError {
	if psdkErr == nil {
		return nil
	}

	// Fill the default error
	if perrResp != nil {
		psdkErr.WithErrorCode(EcUnknownError).WithMessage(perrResp.GetMessage()).WithErrors(perrResp.GetError())
	}

	for _, opt := range popts {
		opt(psdkErr)
		if psdkErr.GetErrorCode() != EcUnknownError {
			return psdkErr
		}
	}

	return psdkErr
}

func WithErrorInternalServerError() func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcInternalServerError).
			WithMessage("Internal Server Error").
			WithErrors(lfmt.Errorf("internal server error from making request to external service"))
	}
}

func WithErrorOutOfPoc(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternOutOfPoc) {
			sdkError.WithErrorCode(EcBillingOutOfPoc).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
