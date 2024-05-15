package sdk_error

import lfmt "fmt"

type (
	ErrorResponse struct {
		Errors []ErrorResponseDTO `json:"errors,omitempty"`
	}

	ErrorResponseDTO struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}
)

func NewErrorResponse() IErrorRespone {
	return &ErrorResponse{}
}

func (s *ErrorResponse) GetErrors() []ErrorResponseDTO {
	return s.Errors
}

func (s *ErrorResponse) GetMessage() string {
	if len(s.Errors) < 1 {
		return ""
	}

	return s.Errors[0].Message
}

func (s *ErrorResponse) GetError() error {
	if len(s.Errors) < 1 {
		return nil
	}

	return lfmt.Errorf("%s: %s", s.Errors[0].Code, s.Errors[0].Message)
}
