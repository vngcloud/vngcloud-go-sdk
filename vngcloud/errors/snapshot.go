package errors

import (
	lstr "strings"

	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
)

var (
	ErrCodeSnapshotNotFound lssdkErr.ErrorCode = "ErrorVolumeNotFound"
)

const (
	patternSnapshotNotFound = "not found snapshot-volume-point"
)

func WithErrorSnapshotNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSnapshotNotFound) {
			sdkError.Code = ErrCodeSnapshotNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
