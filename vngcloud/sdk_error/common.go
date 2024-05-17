package sdk_error

import lfmt "fmt"

const (
	EcUnknownError = ErrorCode("UnknownError")
	EmUnknownError = "Unknown error"

	EcInternalServerError = ErrorCode("VngCloudApiInternalServerError")
)

func ErrorHandler(perr error, popts ...func(psdkErr ISdkError)) ISdkError {
	sdkErr := &SdkError{
		error:     perr,
		errorCode: EcUnknownError,
		message:   EmUnknownError,
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

func SdkErrorHandler(psdkErr ISdkError, popts ...func(psdkErr ISdkError)) ISdkError {
	if psdkErr == nil {
		return nil
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
