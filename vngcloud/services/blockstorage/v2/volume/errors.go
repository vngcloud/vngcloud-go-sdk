package volume

import vconError "github.com/vngcloud/vngcloud-go-sdk/error"

const (
	patternVolumeNotFound = "is not found"
)

var (
	ErrVolumeNotFound vconError.ErrorCode = "VOLUME_NOT_FOUND"
	ErrVolumeUnknown  vconError.ErrorCode = "VOLUME_UNKNOWN_ERROR"
)
