package sdk_error

import lfmt "fmt"

func WithErrorQuotaNotFound(_ IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		sdkError.WithErrorCode(EcVServerQuotaNotFound).
			WithMessage("Quota not found").
			WithErrors(lfmt.Errorf("quota not found"))
	}
}
