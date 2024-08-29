package sdk_error

import (
	lerrors "errors"
	lfmt "fmt"
)

var (
	_ IError = new(SdkError)
)

type (
	SdkError struct {
		error      error
		errorCode  ErrorCode
		message    string
		parameters map[string]interface{}
	}

	ErrorCode string
)

func (s *SdkError) IsError(perrCode ErrorCode) bool {
	return s.errorCode == perrCode
}

func (s *SdkError) IsErrorAny(perrCodes ...ErrorCode) bool {
	for _, perrCode := range perrCodes {
		if s.errorCode == perrCode {
			return true
		}
	}

	return false
}

func (s *SdkError) WithErrorCode(perrCode ErrorCode) IError {
	s.errorCode = perrCode
	return s
}

func (s *SdkError) WithMessage(pmsg string) IError {
	s.message = pmsg
	return s
}

func (s *SdkError) WithErrors(perrs ...error) IError {
	if len(perrs) == 0 {
		return s
	}

	if len(perrs) == 1 {
		s.error = perrs[0]
		return s
	}

	for _, err := range perrs {
		s.error = lerrors.Join(s.error, err)
	}

	return s
}

func (s *SdkError) WithParameters(pparams map[string]interface{}) IError {
	if s.parameters == nil {
		s.parameters = pparams
		return s
	}

	for key, val := range pparams {
		s.parameters[key] = val
	}

	return s
}

func (s *SdkError) WithKVparameters(pparams ...interface{}) IError {
	if s.parameters == nil {
		s.parameters = make(map[string]interface{})
	}

	// Always make sure that the length of pparams is even
	if len(pparams)%2 != 0 {
		pparams = append(pparams, nil)
	}

	for i := 0; i < len(pparams); i += 2 {
		key, ok := pparams[i].(string)
		if !ok {
			continue
		}

		s.parameters[key] = pparams[i+1]
	}

	return s
}

func (s *SdkError) GetError() error {
	return s.error
}

func (s *SdkError) GetMessage() string {
	return s.message
}

func (s *SdkError) GetErrorCode() ErrorCode {
	return s.errorCode
}

func (s *SdkError) GetStringErrorCode() string {
	return string(s.errorCode)
}

func (s *SdkError) GetParameters() map[string]interface{} {
	return s.parameters
}

func (s *SdkError) GetErrorMessages() string {
	if s.error == nil {
		return s.message
	}

	return lfmt.Sprintf("%s: %s", s.message, s.error.Error())
}

func (s *SdkError) GetListParameters() []interface{} {
	var result []interface{}
	if s.parameters == nil || len(s.parameters) < 1 {
		return result
	}

	for key, val := range s.parameters {
		result = append(result, key, val)
	}

	return result
}
