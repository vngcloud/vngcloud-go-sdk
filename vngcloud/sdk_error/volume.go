package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const (
	patternVolumeTypeNotFound              = "cannot get volume type with id"                                                                                       // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
	patternVolumeNameNotValid              = "only letters (a-z, a-z, 0-9, '.', '@', '_', '-', space) are allowed. your input data length must be between 5 and 50" // "Volume name is not valid"
	patternVolumeSizeOutOfRange            = "field volume_size must from"
	patternVolumeNewSizeOutOfRange         = "field new_volume_size must from"
	patternVolumeNotFound                  = `volume with id [^.]+ is not found`
	patternVolumeAvailable                 = "this volume is available"
	patternVolumeAlreadyAttached           = "already attached to instance"
	patternVolumeAlreadyAttachedThisServer = "this volume has been attached"
	patternVolumeInProcess                 = "is in-process"
	patternVolumeUnchaged                  = "volume size or volume type must be changed"
	patternVolumeMustSameZone              = "new volume type must be same zone"
)

var (
	regexErrorVolumeNotFound = lregexp.MustCompile(patternVolumeNotFound)
)

func WithErrorVolumeTypeNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeTypeNotFound) {
			sdkError.WithErrorCode(EcVServerVolumeTypeNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeNameNotValid(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeNameNotValid) {
			sdkError.WithErrorCode(EcVServerVolumeNameNotValid).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeSizeOutOfRange(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeSizeOutOfRange) ||
			lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeNewSizeOutOfRange) {
			sdkError.WithErrorCode(EcVServerVolumeSizeOutOfRange).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeNotFound(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorVolumeNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVolumeNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeAvailable(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeAvailable) {
			sdkError.WithErrorCode(EcVServerVolumeAvailable).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeAlreadyAttached(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeAlreadyAttached) {
			sdkError.WithErrorCode(EcVServerVolumeAlreadyAttached).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeAlreadyAttachedThisServer(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeAlreadyAttachedThisServer) {
			sdkError.WithErrorCode(EcVServerVolumeAlreadyAttachedThisServer).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeInProcess(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeInProcess) {
			sdkError.WithErrorCode(EcVServerVolumeInProcess).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeUnchanged(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeUnchaged) {
			sdkError.WithErrorCode(EcVServerVolumeUnchanged).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeMustSameZone(perrResp IErrorRespone) func(sdkError ISdkError) {
	return func(sdkError ISdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeMustSameZone) {
			sdkError.WithErrorCode(EcVServerVolumeMustSameZone).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
