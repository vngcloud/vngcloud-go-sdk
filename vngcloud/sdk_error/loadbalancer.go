package sdk_error

import (
	lstr "strings"
)

const (
	patternLoadBalancerNotBelongToProject = "does not belong to project"
	patternLoadBalancerNotFound           = "cannot get load balancer with id" // "Cannot get load balancer with id"
)

func WithErrorLoadBalancerNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternLoadBalancerNotBelongToProject) ||
			lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternLoadBalancerNotFound) {
			sdkError.WithErrorCode(EcVLBLoadBalancerNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
