package errors

import (
	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
	lstr "strings"
)

const (
	patternLoadBalancerNotBelongToProject = "does not belong to project"
	patternLoadBalancerNotFound           = "cannot get load balancer with id" // "Cannot get load balancer with id"
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
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternLoadBalancerNotBelongToProject) ||
			lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternLoadBalancerNotFound) {
			sdkError.Code = ErrCodeLoadBalancerNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
