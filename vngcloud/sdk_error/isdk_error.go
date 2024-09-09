package sdk_error

type IError interface {
	IsError(perrCode ErrorCode) bool
	IsErrorAny(perrCodes ...ErrorCode) bool

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
	GetErrorMessages() string
	GetListParameters() []interface{}
}

type IErrorRespone interface {
	GetMessage() string
	GetError() error
}
