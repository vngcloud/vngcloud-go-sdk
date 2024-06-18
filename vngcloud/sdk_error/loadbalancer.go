package sdk_error

import (
	lstr "strings"
)

const (
	patternLoadBalancerNotBelongToProject = "does not belong to project"
	patternLoadBalancerNotFound           = "cannot get load balancer with id" // "Cannot get load balancer with id"
	patternLoadBalancerDuplicatePoolName  = "duplicated pool name"
	patternLoadBalancerNotFound2          = "could not find resource"
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

// WithErrorLoadBalancerNotFound2 indicate the issue creating Pool with non-existed LoadBalancer
func WithErrorLoadBalancerNotFound2(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternLoadBalancerNotFound2) {
			sdkError.WithErrorCode(EcVLBLoadBalancerNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorLoadBalancerDuplicatePoolName(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternLoadBalancerDuplicatePoolName) {
			sdkError.WithErrorCode(EcVLBLoadBalancerDuplicatePoolName).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
