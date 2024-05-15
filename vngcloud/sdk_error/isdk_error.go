package sdk_error

type ISdkError interface {
	IsError(perrCode ErrorCode) bool
	IsErrorAny(perrCodes ...ErrorCode) bool

	WithErrorCode(perrCode ErrorCode) ISdkError
	WithMessage(pmsg string) ISdkError
	WithErrors(perrs ...error) ISdkError
	WithParameters(pparams map[string]interface{}) ISdkError
	WithKVparameters(pparams ...interface{}) ISdkError

	GetError() error
	GetMessage() string
	GetParameters() map[string]interface{}
	GetErrorMessages() string
}

type IErrorRespone interface {
	GetErrors() []ErrorResponseDTO
}
