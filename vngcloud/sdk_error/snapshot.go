package sdk_error

import lstr "strings"

const (
	patternSnapshotNameNotValid = "only letters (a-z, a-z, 0-9, '.', '@', '_', '-', space) are allowed. your input data length must be between 5 and 50" // "Volume name is not valid"
	patternSnapshotNotFound     = "not found snapshot-volume-point"
)

func WithErrorSnapshotNameNotValid(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSnapshotNameNotValid) {
			sdkError.WithErrorCode(EcVServerSnapshotNameNotValid).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSnapshotNameNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSnapshotNotFound) {
			sdkError.WithErrorCode(EcVServerSnapshotNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
