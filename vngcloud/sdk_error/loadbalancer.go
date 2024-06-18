package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternLoadBalancerNotBelongToProject  = "does not belong to project"
	patternLoadBalancerNotFound            = "cannot get load balancer with id" // "Cannot get load balancer with id"
	patternLoadBalancerDuplicatePoolName   = "duplicated pool name"
	patternLoadBalancerNotFound2           = "could not find resource"
	patternListenerDuplicateName           = "duplicated listener name"
	patternListenerDuplicateProtocolOrPort = "duplicated listener protocol port"
	patternPoolNotFound                    = "cannot get pool with id"
	patternLoadBalancerNotReady            = `the load balancer id [^.]+ is not ready`
)

var (
	regexErrorLoadBalancerNotReady = lregexp.MustCompile(patternLoadBalancerNotReady)
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

func WithErrorListenerDuplicateProtocolOrPort(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternListenerDuplicateProtocolOrPort) {
			sdkError.WithErrorCode(EcVLBListenerDuplicateProtocolOrPort).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorListenerDuplicateName(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternListenerDuplicateName) {
			sdkError.WithErrorCode(EcVLBListenerDuplicateName).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorPoolNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternPoolNotFound) {
			sdkError.WithErrorCode(EcVLBPoolNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorLoadBalancerNotReady(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorLoadBalancerNotReady.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVLBLoadBalancerNotReady).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
