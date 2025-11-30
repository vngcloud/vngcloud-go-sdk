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

	NetworkGatewayErrorResponse struct {
		Message   string `json:"message,omitempty"`
		Code      int    `json:"code,omitempty"`
		ErrorCode string `json:"errorCode,omitempty"`
		Success   bool   `json:"success,omitempty"`
	}

	// {"code":"global_load_balancer_not_found","error":"Global load balancer is not found"}
	GlobalLoadBalancerErrorResponse struct {
		Code  string `json:"code,omitempty"`
		Error string `json:"error,omitempty"`
	}
)

const (
	NormalErrorType = iota
	IamErrorType
	NetworkGatewayErrorType
	GlobalLoadBalancerErrorType
)

func NewErrorResponse(ptype int) IErrorRespone {
	switch ptype {
	case IamErrorType:
		return new(IamErrorResponse)
	case NetworkGatewayErrorType:
		return new(NetworkGatewayErrorResponse)
	case GlobalLoadBalancerErrorType:
		return new(GlobalLoadBalancerErrorResponse)
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

func (s *NetworkGatewayErrorResponse) GetMessage() string {
	return lfmt.Sprintf("%s/%d/%s", s.ErrorCode, s.Code, s.Message)
}

func (s *NetworkGatewayErrorResponse) GetError() error {
	return lfmt.Errorf("%s", s.ErrorCode)
}

func (s *GlobalLoadBalancerErrorResponse) GetMessage() string {
	return s.Error
}

func (s *GlobalLoadBalancerErrorResponse) GetError() error {
	return lfmt.Errorf("%s", s.Code)
}
