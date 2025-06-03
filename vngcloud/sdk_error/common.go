package sdk_error

import (
	lfmt "fmt"
	lregexp "regexp"
	lstr "strings"

	lreq "github.com/imroc/req/v3"
)

const (
	patternPurchaseIssue      = "you do not have sufficient credits to complete the purchase"
	patternPagingInvalid      = "page or size invalid"
	patternTagKeyInvalid      = "the value for the tag key contains illegal characters"
	patternServiceMaintenance = "this service is in maintenance"
	patternProjectConflict    = `project [^.]+ is not belong to user`
)

var (
	regexErrorProjectConflict = lregexp.MustCompile(patternProjectConflict)
)

func ErrorHandler(perr error, popts ...func(psdkErr IError)) IError {
	sdkErr := &SdkError{
		error:     perr,
		errorCode: EcUnknownError,
		message:   "Unknown error",
	}

	if perr != nil && lstr.Contains(lstr.ToLower(lstr.TrimSpace(perr.Error())), patternServiceMaintenance) {
		sdkErr.errorCode = EcServiceMaintenance
		sdkErr.message = "Service Maintenance"
		sdkErr.error = lfmt.Errorf("service is under maintenance")

		return sdkErr
	}

	for _, opt := range popts {
		opt(sdkErr)
		if sdkErr.errorCode != EcUnknownError {
			return sdkErr
		}
	}

	sdkErr.error = perr
	sdkErr.message = ""

	return sdkErr
}

func SdkErrorHandler(psdkErr IError, perrResp IErrorRespone, popts ...func(psdkErr IError)) IError {
	if psdkErr == nil && perrResp == nil {
		return nil
	}

	if psdkErr != nil && psdkErr.GetErrorCode() != EcUnknownError {
		return psdkErr
	}

	// Fill the default error
	if perrResp != nil {
		psdkErr.WithErrorCode(EcUnknownError).WithMessage(perrResp.GetMessage()).WithErrors(perrResp.GetError())
	}

	for _, opt := range popts {
		opt(psdkErr)
		if psdkErr.GetErrorCode() != EcUnknownError {
			return psdkErr
		}
	}

	return psdkErr
}

func WithErrorInternalServerError() func(IError) {
	return func(sdkErr IError) {
		sdkErr.WithErrorCode(EcInternalServerError).
			WithMessage("Internal Server Error").
			WithErrors(lfmt.Errorf("internal server error from making request to external service"))
	}
}

func WithErrorServiceMaintenance() func(IError) {
	return func(sdkErr IError) {
		sdkErr.WithErrorCode(EcServiceMaintenance).
			WithMessage("Service Maintenance").
			WithErrors(lfmt.Errorf("service is under maintenance"))
	}
}

func WithErrorPermissionDenied() func(IError) {
	return func(sdkErr IError) {
		sdkErr.WithErrorCode(EcPermissionDenied).
			WithMessage("Permission Denied").
			WithErrors(lfmt.Errorf("permission denied when making request to external service"))
	}
}

func WithErrorPurchaseIssue(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternPurchaseIssue) {
			sdkError.WithErrorCode(EcPurchaseIssue).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatPurchase)
		}
	}
}

func WithErrorTagKeyInvalid(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternTagKeyInvalid) {
			sdkError.WithErrorCode(EcTagKeyInvalid).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorPagingInvalid(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternPagingInvalid) {
			sdkError.WithErrorCode(EcPagingInvalid).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorUnexpected(presponse *lreq.Response) func(IError) {
	statusCode := 0
	url := ""
	err := lfmt.Errorf("unexpected error from making request to external service")
	if presponse != nil {
		if presponse.Response != nil && presponse.StatusCode != 0 {
			statusCode = presponse.StatusCode
		}

		if presponse.Request != nil && presponse.Request.URL != nil {
			url = presponse.Request.URL.String()
		}

		if presponse.Err != nil {
			err = presponse.Err
		}
	}

	return func(sdkErr IError) {
		sdkErr.WithErrorCode(EcUnexpectedError).
			WithMessage("Unexpected Error").
			WithErrors(err).
			WithParameters(map[string]interface{}{
				"statusCode": statusCode,
				"url":        url,
			})
	}
}

func WithErrorPaymentMethodNotAllow(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if lstr.Contains(errMsg, "ext_pm_payment_method_not_allow") {
			sdkError.WithErrorCode(EcPaymentMethodNotAllow).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorCreditNotEnough(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if lstr.Contains(errMsg, "ext_pm_credit_not_enough") {
			sdkError.WithErrorCode(EcCreditNotEnough).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorProjectConflict(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorProjectConflict.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcProjectConflict).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
