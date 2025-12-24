package sdk_error

import lstr "strings"

func WithErrorEndpointStatusInvalid(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorEndpointOfVpcExists(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorEndpointPackageNotBelongToEndpointService(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorContainInvalidCharacter(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		if lstr.ToUpper(lstr.TrimSpace(perrResp.GetError().Error())) == "CONTAIN_INVALID_CHARACTER" {
			sdkError.WithErrorCode(EcVNetworkContainInvalidCharacter).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorLockOnProcess(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		if lstr.ToUpper(lstr.TrimSpace(perrResp.GetError().Error())) == "LOCK_ON_PROCESS" {
			sdkError.WithErrorCode(EcVNetworkLockOnProcess).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorEndpointTagNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		if lstr.ToUpper(lstr.TrimSpace(perrResp.GetError().Error())) == "TAG_RESOURCE_WAS_DELETED" {
			sdkError.WithErrorCode(EcVNetworkEndpointTagNotFound).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorEndpointTagExisted(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		if lstr.ToUpper(lstr.TrimSpace(perrResp.GetError().Error())) == "TAG_EXISTED" {
			sdkError.WithErrorCode(EcVNetworkEndpointTagExisted).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}
