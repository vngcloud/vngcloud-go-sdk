package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternServerNotFound                      = "cannot get server with id"                 // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
	patternServerCreating                      = "cannot delete server with status creating" // "Server is creating"
	patternServerExceedQuota                   = "exceeded vm quota"                         // "The number of servers exceeds the quota"
	patternServerDeleting                      = "cannot delete server with status deleting" // "Server is deleting"
	patternServerBilling                       = "cannot delete server with status creating-billing"
	patternBillingPaymentMethodNotAllowed      = "payment method is not allowed for the user"
	patternServerAttachVolumeQuotaExceeded     = "exceeded volume_per_server quota"
	patternServerAttachEncryptedVolume         = "cannot attach encryption volume"
	patternServerExpired                       = "server is expired"
	patternServerFlavorSystemExceedQuota       = "there are no more remaining flavor with id"
	patternServerUpdatingSecgroups             = "cannot change security group of server with status changing-security-group"
	patternServerExceedCpuQuota                = "exceeded vcpu quota. current used"
	patternServerImageNotSupported             = "doesn't support image with id"
	patternImageNotSupport                     = "don't support image"
	patternServerCanNotAttachFloatingIp        = "the server only allows attaching 1 floating ip"
	patternServerFlavorNotSupported            = `flavor [^.]+ don't support image [^.]+`
	patternServerDeleteServerUpdatingSecgroups = "cannot delete server with status changing-security-group"
	patternServerExceedFloatingIpQuota         = "exceeded floating_ip quota"
	patternImageNotFound                       = "cannot get image with id"
	patternServerGroupNotFound                 = "not found server group"
	patternServerGroupInUse                    = "server group is in use"
	patternServerGroupNameMustBeUnique         = "name must be unique"
)

var (
	regexErrorServerFlavorNotSupported = lregexp.MustCompile(patternServerFlavorNotSupported)
)

func WithErrorServerFlavorNotSupported(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorServerFlavorNotSupported.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerFlavorNotSupported).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerNotFound) {
			sdkError.WithErrorCode(EcVServerServerNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorImageNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if lstr.Contains(errMsg, patternImageNotFound) {
			sdkError.WithErrorCode(EcVServerImageNotFound).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerFlavorSystemExceedQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerFlavorSystemExceedQuota) {
			sdkError.WithErrorCode(EcVServerServerFlavorSystemExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatInfra)
		}
	}
}

func WithErrorServerDeleteCreatingServer(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerCreating) {
			sdkError.WithErrorCode(EcVServerServerDeleteCreatingServer).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerExpired(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerExpired) {
			sdkError.WithErrorCode(EcVServerServerExpired).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerUpdatingSecgroups(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		stdErrMsg := lstr.ToLower(lstr.TrimSpace(errMsg))
		if lstr.Contains(stdErrMsg, patternServerUpdatingSecgroups) ||
			lstr.Contains(stdErrMsg, patternServerDeleteServerUpdatingSecgroups) {
			sdkError.WithErrorCode(EcVServerServerUpdatingSecgroups).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerExceedCpuQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerExceedCpuQuota) {
			sdkError.WithErrorCode(EcVServerServerExceedCpuQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorServerExceedFloatingIpQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerExceedFloatingIpQuota) {
			sdkError.WithErrorCode(EcVServerServerExceedFloatingIpQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorServerImageNotSupported(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		lowerErrMsg := lstr.ToLower(lstr.TrimSpace(errMsg))
		if lstr.Contains(lowerErrMsg, patternServerImageNotSupported) ||
			lstr.Contains(lowerErrMsg, patternImageNotSupport) {
			sdkError.WithErrorCode(EcVServerServerImageNotSupported).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerExceedQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerExceedQuota) {
			sdkError.WithErrorCode(EcVServerServerExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorServerDeleteDeletingServer(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerDeleting) {
			sdkError.WithErrorCode(EcVServerServerDeleteDeletingServer).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerDeleteBillingServer(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerBilling) {
			sdkError.WithErrorCode(EcVServerServerDeleteBillingServer).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerCreateBillingPaymentMethodNotAllowed(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if lstr.Contains(errMsg, patternBillingPaymentMethodNotAllowed) {
			sdkError.WithErrorCode(EcVServerCreateBillingPaymentMethodNotAllowed).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerAttachVolumeQuotaExceeded(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerAttachVolumeQuotaExceeded) {
			sdkError.WithErrorCode(EcVServerServerVolumeAttachQuotaExceeded).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorServerAttachEncryptedVolume(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerAttachEncryptedVolume) {
			sdkError.WithErrorCode(EcVServerServerAttachEncryptedVolume).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerCanNotAttachFloatingIp(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerCanNotAttachFloatingIp) {
			sdkError.WithErrorCode(EcVServerServerCanNotAttachFloatingIp).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerGroupNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerGroupNotFound) {
			sdkError.WithErrorCode(EcVServerServerGroupNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerGroupInUse(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerGroupInUse) {
			sdkError.WithErrorCode(EcVServerServerGroupInUse).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerGroupNameMustBeUnique(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerGroupNameMustBeUnique) {
			sdkError.WithErrorCode(EcVServerServerGroupNameMustBeUnique).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
