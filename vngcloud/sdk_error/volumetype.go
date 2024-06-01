package sdk_error

import lstr "strings"

const (
	patternVolumeTypeNotFound = "cannot get volume type with id"
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
