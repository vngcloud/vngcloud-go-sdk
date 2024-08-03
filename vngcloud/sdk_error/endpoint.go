package sdk_error

func WithErrorEndpointStatusInvalid(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		if perrResp.GetError().Error() == "ENDPOINT_STATUS_INVALID" {
			sdkError.WithErrorCode(EcVNetworkEndpointStatusInvalid).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}
