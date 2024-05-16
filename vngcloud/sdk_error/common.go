package sdk_error

const (
	EcUnknownError = ErrorCode("UnknownError")
	EmUnknownError = "Unknown error"

	EcCanNotMarshalRequestBody    = ErrorCode("CanNotMarshalRequestBody")
	EcFailedToCreateHttpRequest   = ErrorCode("FailedToCreateHttpRequest")
	EcFailedToMakeHttpRequest     = ErrorCode("FailedToMakeHttpRequest")
	EcFailedToReadResponseBody    = ErrorCode("FailedToReadResponseBody")
	EcCanNotUnmarshalResponseBody = ErrorCode("CanNotUnmarshalResponseBody")
	EcOkCodeNotMatch              = ErrorCode("OkCodeNotMatch")
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

func WithErrorCanNotMarshalRequestBody(perr error) func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcCanNotMarshalRequestBody).
			WithMessage(perr.Error()).
			WithErrors(perr)
	}
}
func WithErrorFailedToCreateHttpRequest(perr error) func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcFailedToCreateHttpRequest).
			WithMessage(perr.Error()).
			WithErrors(perr)
	}
}

func WithErrorFailedToMakeHttpRequest(perr error) func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcFailedToMakeHttpRequest).
			WithMessage(perr.Error()).
			WithErrors(perr)
	}
}

func WithErrorFailedToReadResponseBody(perr error) func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcFailedToReadResponseBody).
			WithMessage(perr.Error()).
			WithErrors(perr)
	}
}

func WithErrorCanNotUnmarshalResponseBody(perr error) func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcCanNotUnmarshalResponseBody).
			WithMessage(perr.Error()).
			WithErrors(perr)
	}
}

func WithErrorOkCodeNotMatch(pcode int) func(ISdkError) {
	return func(sdkErr ISdkError) {
		sdkErr.WithErrorCode(EcOkCodeNotMatch).
			WithMessage("Ok code not match").
			WithParameters(map[string]interface{}{
				"code": pcode,
			})
	}
}
