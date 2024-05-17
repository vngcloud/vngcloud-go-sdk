package sdk_error

import lfmt "fmt"

type (
	IamErrorResponse struct {
		Errors []struct {
			Code    string `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		} `json:"errors,omitempty"`
	}

	NormalErrorResponse struct {
		Message string `json:"message,omitempty"`
	}
)

const (
	IamErrorType = iota
	NormalErrorType
)

func NewErrorResponse(ptype int) IErrorRespone {
	switch {
	case ptype == IamErrorType:
		return new(IamErrorResponse)
	default:
		return new(NormalErrorResponse)
	}
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

	return lfmt.Errorf("%s", s.Errors[0].Code)
}

func (s *NormalErrorResponse) GetMessage() string {
	return s.Message
}

func (s *NormalErrorResponse) GetError() error {
	return lfmt.Errorf("%s", s.Message)
}
