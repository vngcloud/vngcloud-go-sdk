package errors

import (
	lstr "strings"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

var (
	ErrCodeVolumeAvailable lssdkErr.ErrorCode = "ErrorVolumeAvailable"
	ErrCodeVolumeNotFound  lssdkErr.ErrorCode = "ErrorVolumeNotFound"
)

const (
	patternVolumeAvailable = "this volume is available"
	patternVolumeNotFound  = "is not found"
)

func WithErrorVolumeAvailable(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeAvailable) {
			sdkError.Code = ErrCodeVolumeAvailable
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}

func WithErrorNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeNotFound) {
			sdkError.Code = ErrCodeVolumeNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
