package errors

import (
	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
	lstr "strings"
)

const (
	patternSecgroupNameAlreadyExists = "name of security group already exist"
	patternSecgroupExceedQuota       = "exceeded secgroup quota"
	patternSecgroupInUse             = "securitygroupinuse"
	patternSecgroupNotFound          = "cannot get security group with id"
)

var (
	ErrCodeSecgroupNameAlreadyExists lssdkErr.ErrorCode = "ErrorSecgroupNameAlreadyExists"
	ErrCodeSecgroupExceedQuota       lssdkErr.ErrorCode = "ErrorSecgroupExceedQuota"
	ErrCodeSecgroupInUse             lssdkErr.ErrorCode = "ErrorSecgroupInUse"
	ErrCodeSecgroupNotFound          lssdkErr.ErrorCode = "ErrorSecgroupNotFound"
)

func WithErrorSecgroupNameAlreadyExists(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupNameAlreadyExists) {
			sdkError.Code = ErrCodeSecgroupNameAlreadyExists
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}

func WithErrorSecgroupExceedQuota(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupExceedQuota) {
			sdkError.Code = ErrCodeSecgroupExceedQuota
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}

func WithErrorSecgroupInUse(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupInUse) {
			sdkError.Code = ErrCodeSecgroupInUse
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}

func WithErrorSecgroupNotFound(perrResp *lssdkErr.ErrorResponse, perr error) func(*lssdkErr.SdkError) {
	return func(sdkError *lssdkErr.SdkError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.Message
		if lstr.Contains(lstr.ToLower(lstr.TrimSpace(errMsg)), patternSecgroupNotFound) {
			sdkError.Code = ErrCodeSecgroupNotFound
			sdkError.Message = errMsg
			sdkError.Error = perr
		}
	}
}
