package sdk_error

import lstr "strings"

const (
	patternVolumeTypeNotFound   = "cannot get volume type with id"                                                                                       // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
	patternVolumeNameNotValid   = "only letters (a-z, a-z, 0-9, '.', '@', '_', '-', space) are allowed. your input data length must be between 5 and 50" // "Volume name is not valid"
	patternVolumeSizeOutOfRange = "field volume_size must from"
)

func WithErrorVolumeTypeNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeTypeNotFound) {
			sdkError.WithErrorCode(EcVServerVolumeTypeNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeNameNotValid(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeNameNotValid) {
			sdkError.WithErrorCode(EcVServerVolumeNameNotValid).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeSizeOutOfRange(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeSizeOutOfRange) {
			sdkError.WithErrorCode(EcVServerVolumeSizeOutOfRange).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
