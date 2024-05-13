package errors

import (
	lstr "strings"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

var (
	ErrCodeVolumeTypeNotFound lssdkErr.ErrorCode = "ErrorVolumeTypeNotFound"
)

const (
	patternVolumeTypeNotFound = "cannot get volume type with id" // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
)

func WithErrorVolumeTypeNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeTypeNotFound) {
			sdkError.Code = ErrCodeVolumeTypeNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
