package sdk_error

const (
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
