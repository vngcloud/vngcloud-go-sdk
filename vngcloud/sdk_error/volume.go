package sdk_error

import lstr "strings"

const (
	patternVolumeTypeNotFound = "cannot get volume type with id" // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
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
