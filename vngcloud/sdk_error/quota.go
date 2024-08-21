package sdk_error

import lfmt "fmt"

func WithErrorQuotaNotFound(_ IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		sdkError.WithErrorCode(EcVServerQuotaNotFound).
			WithMessage("Quota not found").
			WithErrors(lfmt.Errorf("quota not found"))
	}
}
