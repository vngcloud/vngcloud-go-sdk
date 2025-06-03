package sdk_error

import (
	lregexp "regexp"
	lstr "strings"
)

const ( // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
	patternVolumeNameNotValid              = "only letters (a-z, a-z, 0-9, '.', '@', '_', '-', space) are allowed. your input data length must be between 5 and 50" // "Volume name is not valid"
	patternVolumeSizeOutOfRange            = "field volume_size must from"
	patternVolumeNewSizeOutOfRange         = "field new_volume_size must from"
	patternVolumeNotFound                  = `volume with id [^.]+ is not found`
	patternVolumeNotFound2                 = "cannot get volume with id"
	patternVolumeAvailable                 = "this volume is available"
	patternVolumeAlreadyAttached           = "already attached to instance"
	patternVolumeAlreadyAttachedThisServer = "this volume has been attached"
	patternVolumeInProcess                 = "is in-process"
	patternVolumeUnchaged                  = "volume size or volume type must be changed"
	patternVolumeMustSameZone              = "new volume type must be same zone"
	patternVolumeMigrateMissingInit        = "the action must be init-migrate or migrate or confirm-migrate"
	patternVolumeMigrateNeedProcess        = "this volume cannot initialize migration because state is ready to migrate difference"
	patternVolumeMigrateNeedConfirm        = "this volume cannot initialize migration because state is confirm final migration"
	patternVolumeMigrateBeingProcess       = "this volume cannot initialize migration because state is migrating difference"
	patternVolumeMigrateBeingMigrating     = "this volume cannot initialize migration because state is migrating"
	patternVolumeMigrateBeingFinish        = "this volume cannot migrate difference data because state is confirm final migration"
	patternVolumeMigrateProcessingConfirm  = "this volume cannot initialize migration because state is processing to confirm"
	patternVolumeMigrateInSameZone         = "new volume type must be different zone"
	patternVolumeIsMigrating               = "is migrating"
	patternVolumeSizeExceedGlobalQuota     = "exceeded volume_size quota"
	patternVolumeExceedQuota               = "exceeded volume quota. current used"
)

var (
	regexErrorVolumeNotFound = lregexp.MustCompile(patternVolumeNotFound)
)

func WithErrorVolumeNameNotValid(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorVolumeSizeOutOfRange(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorVolumeSizeExceedGlobalQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeSizeExceedGlobalQuota) {
			sdkError.WithErrorCode(EcVServerVolumeSizeExceedGlobalQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorVolumeExceedQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeExceedQuota) {
			sdkError.WithErrorCode(EcVServerVolumeExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorVolumeNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := lstr.ToLower(lstr.TrimSpace(perrResp.GetMessage()))
		if regexErrorVolumeNotFound.FindString(errMsg) != "" ||
			lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeNotFound2) {
			sdkError.WithErrorCode(EcVServerVolumeNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

// WithErrorVolumeAvailable indicates that the volume is AVAILABLE state but try to make detach this volume out of server
func WithErrorVolumeAvailable(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorVolumeAlreadyAttached(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorVolumeAlreadyAttachedThisServer(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorVolumeInProcess(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorVolumeUnchanged(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorVolumeMustSameZone(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
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

func WithErrorVolumeMigrateMissingInit(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeMigrateMissingInit) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateMissingInit).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateNeedProcess(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeMigrateNeedProcess) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateNeedProcess).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateNeedConfirm(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeMigrateNeedConfirm) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateNeedConfirm).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateBeingProcess(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeMigrateBeingProcess) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateBeingProcess).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateBeingFinish(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeMigrateBeingFinish) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateBeingFinish).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateProcessingConfirm(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeMigrateProcessingConfirm) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateProcessingConfirm).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

//

func WithErrorVolumeMigrateBeingMigrating(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeMigrateBeingMigrating) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateBeingMigrating).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateInSameZone(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeMigrateInSameZone) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateInSameZone).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVolumeIsMigrating(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternVolumeIsMigrating) {
			sdkError.WithErrorCode(EcVServerVolumeIsMigrating).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

// patternVolumeIsMigrating
