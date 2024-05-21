package sdk_error

import lstr "strings"

const (
	patternServerNotFound    = "cannot get server with id"                 // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
	patternServerCreating    = "cannot delete server with status creating" // "Server is creating"
	patternServerExceedQuota = "exceeded vm quota"                         // "The number of servers exceeds the quota"
	patternServerDeleting    = "cannot delete server with status deleting" // "Server is deleting"
	patternServerBilling     = "cannot delete server with status creating-billing"
)

func WithErrorServerNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
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

func WithErrorServerDeleteCreatingServer(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
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

func WithErrorServerExceedQuota(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternServerExceedQuota) {
			sdkError.WithErrorCode(EcVServerServerExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorServerDeleteDeletingServer(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
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

func WithErrorServerDeleteBillingServer(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
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
