package sdk_error

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
