package sdk_error

import ljset "github.com/cuongpiger/joat/data-structure/set"

type IError interface {
	IsError(perrCode ErrorCode) bool
	IsErrorAny(perrCodes ...ErrorCode) bool
	IsCategory(pcategory ErrorCategory) bool
	IsCategories(pcategories ...ErrorCategory) bool

	WithErrorCode(perrCode ErrorCode) IError
	WithMessage(pmsg string) IError
	WithErrors(perrs ...error) IError
	WithErrorCategories(pcategories ...ErrorCategory) IError
	WithParameters(pparams map[string]interface{}) IError
	WithKVparameters(pparams ...interface{}) IError

	GetError() error
	GetMessage() string
	GetErrorCode() ErrorCode
	GetStringErrorCode() string
	GetParameters() map[string]interface{}
	GetErrorCategories() ljset.Set[ErrorCategory]
	GetErrorMessages() string
	GetListParameters() []interface{}

	RemoveCategories(pcategories ...ErrorCategory) IError

	AppendCategories(pcategories ...ErrorCategory) IError
}

type IErrorRespone interface {
	GetMessage() string
	GetError() error
}
