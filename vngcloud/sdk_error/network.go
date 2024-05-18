package sdk_error

import lstr "strings"

const (
	patternNetworkNotFound = "is not found"
)

func WithErrorNetworkNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternNetworkNotFound) {
			sdkError.WithErrorCode(EcVServerNetworkNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
