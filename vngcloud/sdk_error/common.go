package sdk_error

import (
	lfmt "fmt"
	lstr "strings"
)

const (
	patternOutOfPoc      = "you do not have sufficient credits to complete the purchase"
	patternPagingInvalid = "page or size invalid"
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
	if psdkErr == nil && perrResp == nil {
		return nil
	}

	if psdkErr != nil && psdkErr.GetErrorCode() != EcUnknownError {
		return psdkErr
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

func WithErrorPermissionDenied() func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcPermissionDenied).
			WithMessage("Permission Denied").
			WithErrors(lfmt.Errorf("permission denied when making request to external service"))
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

func WithErrorPagingInvalid(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternPagingInvalid) {
			sdkError.WithErrorCode(EcPagingInvalid).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorUnexpected() func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcUnexpectedError).
			WithMessage("Unexpected Error").
			WithErrors(lfmt.Errorf("unexpected error from making request to external service"))
	}
}
