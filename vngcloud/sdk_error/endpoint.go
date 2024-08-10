package sdk_error

import lstr "strings"

func WithErrorEndpointStatusInvalid(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		if lstr.ToUpper(lstr.TrimSpace(perrResp.GetError().Error())) == "ENDPOINT_STATUS_INVALID" {
			sdkError.WithErrorCode(EcVNetworkEndpointStatusInvalid).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorEndpointOfVpcExists(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		if lstr.ToUpper(lstr.TrimSpace(perrResp.GetError().Error())) == "ENDPOINT_OF_VPC_IS_EXISTS" {
			sdkError.WithErrorCode(EcVNetworkEndpointOfVpcExists).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorEndpointPackageNotBelongToEndpointService(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		if lstr.ToUpper(lstr.TrimSpace(perrResp.GetError().Error())) == "ENDPOINT_PACKAGE_NOT_BELONG_TO_ENDPOINT_SERVICE" {
			sdkError.WithErrorCode(EcVNetworkEndpointPackageNotBelongToEndpointService).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}
