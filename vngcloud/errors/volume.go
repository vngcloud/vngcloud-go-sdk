package errors

import (
	lstr "strings"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

var (
	ErrCodeVolumeAvailable       lssdkErr.ErrorCode = "ErrorVolumeAvailable"
	ErrCodeVolumeNotFound        lssdkErr.ErrorCode = "ErrorVolumeNotFound"
	ErrCodeVolumeAlreadyAttached lssdkErr.ErrorCode = "ErrorVolumeAlreadyAttached"
)

const (
	patternVolumeAvailable       = "this volume is available"
	patternVolumeNotFound        = "is not found"
	patternVolumeAlreadyAttached = "already attached to instance"
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

func WithErrorVolumeNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
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

func WithErrorVolumeAlreadyAttached(perrResp *lssdkErr.ErrorResponse, pvolID, pinstanceID string, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeAlreadyAttached) {
			sdkError.Code = ErrCodeVolumeAlreadyAttached
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
