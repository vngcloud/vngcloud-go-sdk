package sdk_error

var (
	EcUnknownError = ErrorCode("UnknownError")
	EmUnknownError = "Unknown error"
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
	sdkErr.message = perr.Error()

	return sdkErr
}
