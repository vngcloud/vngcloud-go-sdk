package server

import error2 "github.com/vngcloud/vngcloud-go-sdk/error"

const (
	patternErrNotFound          = "Cannot get server with id ins-"
	patternDeleteCreatingServer = "Cannot delete server with status CREATING"
)

const (
	ErrNotFound             error2.ErrorCode = "SERVER_NOT_FOUND"
	ErrUnknown              error2.ErrorCode = "SERVER_UNKNOWN_ERROR"
	ErrDeleteCreatingServer error2.ErrorCode = "DELETE_CREATING_SERVER"
)
