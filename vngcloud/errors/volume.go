package errors

import (
	lstr "strings"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

var (
	ErrCodeVolumeAvailable       lssdkErr.ErrorCode = "ErrorVolumeAvailable"
	ErrCodeVolumeNotFound        lssdkErr.ErrorCode = "ErrorVolumeNotFound"
	ErrCodeVolumeAlreadyAttached lssdkErr.ErrorCode = "ErrorVolumeAlreadyAttached"
	ErrCodeVolumeUnchanged       lssdkErr.ErrorCode = "ErrorVolumeUnchanged"
)

const (
	patternVolumeAvailable       = "this volume is available"
	patternVolumeNotFound        = "is not found"
	patternVolumeAlreadyAttached = "already attached to instance"
	patternVolumeUnchaged        = "volume size or volume type must be changed"
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

func WithErrorVolumeUnchanged(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeUnchaged) {
			sdkError.Code = ErrCodeVolumeUnchanged
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
