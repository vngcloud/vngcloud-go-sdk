package errors

import (
	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
	lstr "strings"
)

const (
	patternOutOfPoc                        = "you do not have sufficient credits to complete the purchase"
	patternServerAttachVolumeQuotaExceeded = "exceeded volume_per_server quota" // "Exceeded VOLUME_PER_SERVER quota. Current used: 10, max: 10."
)

var (
	ErrCodeOutOfPoc                        lssdkErr.ErrorCode = "ErrorOutOfPoc"
	ErrCodeServerAttachVolumeQuotaExceeded lssdkErr.ErrorCode = "ErrorServerAttachVolumeQuotaExceeded"
)

func WithErrorOutOfPoc(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternOutOfPoc) {
			sdkError.Code = ErrCodeOutOfPoc
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}

func WithErrorServerAttachVolumeQuotaExceeded(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerAttachVolumeQuotaExceeded) {
			sdkError.Code = ErrCodeServerAttachVolumeQuotaExceeded
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
