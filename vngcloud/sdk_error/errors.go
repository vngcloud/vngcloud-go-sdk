package sdk_error

import lfmt "fmt"

type (
	IamErrorResponse struct {
		Errors []struct {
			Code    string `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		} `json:"errors,omitempty"`
	}
)

func NewErrorResponse() IErrorRespone {
	return &IamErrorResponse{}
}

func (s *IamErrorResponse) GetMessage() string {
	if len(s.Errors) < 1 {
		return ""
	}

	return s.Errors[0].Message
}

func (s *IamErrorResponse) GetError() error {
	if len(s.Errors) < 1 {
		return nil
	}

	return lfmt.Errorf("%s: %s", s.Errors[0].Code, s.Errors[0].Message)
}
