package errors

import (
	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
	lstr "strings"
)

const (
	patternLoadBalancerNotBelongToProject = "does not belong to project"
)

var (
	ErrCodeLoadBalancerNotFound lssdkErr.ErrorCode = "ErrorLoadBalancerNotFound"
)

func WithErrorLoadBalancerNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternLoadBalancerNotBelongToProject) {
			sdkError.Code = ErrCodeLoadBalancerNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
