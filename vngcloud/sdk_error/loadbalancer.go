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
	patternListenerNotFound                = "cannot get listener with id"
	patternListenerDuplicateProtocolOrPort = "duplicated listener protocol port"
	patternPoolNotFound                    = "cannot get pool with id"
	patternPoolInUse                       = "is used in listener"
	patternLoadBalancerNotReady            = `the load balancer id [^.]+ is not ready`
	patternListenerNotReady                = `listener id [^.]+ is not ready`
	patternMemberMustIdentical             = "the members provided are identical to the existing members in the pool"
	patternPoolIsUpdating                  = `pool id [^.]+ is updating`
	patternLoadBalancerExceedQuota         = "exceeded load_balancer quota. current used"
	patternLoadBalancerIsDeleting          = `load balancer id [^.]+ is deleting`
)

var (
	regexErrorLoadBalancerNotReady   = lregexp.MustCompile(patternLoadBalancerNotReady)
	regexErrorListenerNotReady       = lregexp.MustCompile(patternListenerNotReady)
	regexErrorPoolIsUpdating         = lregexp.MustCompile(patternPoolIsUpdating)
	regexErrorLoadBalancerIsDeleting = lregexp.MustCompile(patternLoadBalancerIsDeleting)
)

func WithErrorLoadBalancerNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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
func WithErrorLoadBalancerNotFound2(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorLoadBalancerExceedQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternLoadBalancerExceedQuota) {
			sdkError.WithErrorCode(EcVLBLoadBalancerExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorLoadBalancerDuplicatePoolName(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorListenerDuplicateProtocolOrPort(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorListenerDuplicateName(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorPoolNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorPoolInUse(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternPoolInUse) {
			sdkError.WithErrorCode(EcVLBPoolInUse).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorLoadBalancerNotReady(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorLoadBalancerNotReady.FindString(errMsg) != "" ||
			regexErrorListenerNotReady.FindString(errMsg) != "" ||
			regexErrorPoolIsUpdating.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVLBLoadBalancerNotReady).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorLoadBalancerIsDeleting(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorLoadBalancerIsDeleting.FindString(errMsg) != ""  {
			sdkError.WithErrorCode(EcVLBLoadBalancerIsDeleting).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorListenerNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternListenerNotFound) {
			sdkError.WithErrorCode(EcVLBListenerNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorMemberMustIdentical(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternMemberMustIdentical) {
			sdkError.WithErrorCode(EcVLBMemberMustIdentical).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
