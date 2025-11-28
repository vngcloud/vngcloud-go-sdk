package v1

import (
	"strings"

	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
	// lstr "strings"
)

type GlobalListenerProtocol string

const (
	GlobalListenerProtocolTCP GlobalListenerProtocol = "TCP"
)

var _ IListGlobalListenersRequest = &ListGlobalListenersRequest{}

type ListGlobalListenersRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

func (s *ListGlobalListenersRequest) WithLoadBalancerId(plbId string) IListGlobalListenersRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *ListGlobalListenersRequest) AddUserAgent(pagent ...string) IListGlobalListenersRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewListGlobalListenersRequest(plbId string) IListGlobalListenersRequest {
	opts := &ListGlobalListenersRequest{}
	opts.LoadBalancerId = plbId
	return opts
}

// --------------------------------------------------

var _ ICreateGlobalListenerRequest = &CreateGlobalListenerRequest{}

// WithAllowedCidrs(pcidrs ...string) ICreateGlobalListenerRequest
// WithDescription(pdesc string) ICreateGlobalListenerRequest
// WithHeaders(pheaders ...string) ICreateGlobalListenerRequest
// WithName(pname string) ICreateGlobalListenerRequest
// WithPort(pport int) ICreateGlobalListenerRequest
// WithProtocol(pprotocol GlobalListenerProtocol) ICreateGlobalListenerRequest
// WithTimeoutClient(ptoc int) ICreateGlobalListenerRequest
// WithTimeoutConnection(ptoc int) ICreateGlobalListenerRequest
// WithTimeoutMember(ptom int) ICreateGlobalListenerRequest
// WithDefaultPoolId(ppoolId string) ICreateGlobalListenerRequest
type CreateGlobalListenerRequest struct {
	AllowedCidrs      string                 `json:"allowedCidrs"`
	Description       string                 `json:"description"`
	Headers           []string               `json:"headers"`
	Name              string                 `json:"name"`
	Port              int                    `json:"port"`
	Protocol          GlobalListenerProtocol `json:"protocol"`
	TimeoutClient     int                    `json:"timeoutClient"`
	TimeoutConnection int                    `json:"timeoutConnection"`
	TimeoutMember     int                    `json:"timeoutMember"`
	GlobalPoolId      string                 `json:"globalPoolId"`

	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

func (s *CreateGlobalListenerRequest) WithAllowedCidrs(pcidrs ...string) ICreateGlobalListenerRequest {
	s.AllowedCidrs = strings.Join(pcidrs, ",")
	return s
}

func (s *CreateGlobalListenerRequest) WithDescription(pdesc string) ICreateGlobalListenerRequest {
	s.Description = pdesc
	return s
}

func (s *CreateGlobalListenerRequest) WithHeaders(pheaders ...string) ICreateGlobalListenerRequest {
	s.Headers = pheaders
	return s
}

func (s *CreateGlobalListenerRequest) WithName(pname string) ICreateGlobalListenerRequest {
	s.Name = pname
	return s
}

func (s *CreateGlobalListenerRequest) WithPort(pport int) ICreateGlobalListenerRequest {
	s.Port = pport
	return s
}

func (s *CreateGlobalListenerRequest) WithProtocol(pprotocol GlobalListenerProtocol) ICreateGlobalListenerRequest {
	s.Protocol = pprotocol
	return s
}

func (s *CreateGlobalListenerRequest) WithTimeoutClient(ptoc int) ICreateGlobalListenerRequest {
	s.TimeoutClient = ptoc
	return s
}

func (s *CreateGlobalListenerRequest) WithTimeoutConnection(ptoc int) ICreateGlobalListenerRequest {
	s.TimeoutConnection = ptoc
	return s
}

func (s *CreateGlobalListenerRequest) WithTimeoutMember(ptom int) ICreateGlobalListenerRequest {
	s.TimeoutMember = ptom
	return s
}

func (s *CreateGlobalListenerRequest) WithGlobalPoolId(ppoolId string) ICreateGlobalListenerRequest {
	s.GlobalPoolId = ppoolId
	return s
}

func (s *CreateGlobalListenerRequest) WithLoadBalancerId(plbid string) ICreateGlobalListenerRequest {
	s.LoadBalancerId = plbid
	return s
}

func (s *CreateGlobalListenerRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateGlobalListenerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"allowedCidrs":      s.AllowedCidrs,
		"description":       s.Description,
		"headers":           s.Headers,
		"name":              s.Name,
		"port":              s.Port,
		"protocol":          s.Protocol,
		"timeoutClient":     s.TimeoutClient,
		"timeoutConnection": s.TimeoutConnection,
		"timeoutMember":     s.TimeoutMember,
		"globalPoolId":      s.GlobalPoolId,
	}
}

func (s *CreateGlobalListenerRequest) AddUserAgent(pagent ...string) ICreateGlobalListenerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewCreateGlobalListenerRequest(plbId, name string) ICreateGlobalListenerRequest {
	opts := &CreateGlobalListenerRequest{
		AllowedCidrs:      "0.0.0.0/0",
		Description:       "",
		Headers:           nil,
		Name:              name,
		Port:              80,
		Protocol:          GlobalListenerProtocolTCP,
		TimeoutClient:     50,
		TimeoutConnection: 5,
		TimeoutMember:     50,
		GlobalPoolId:      "",
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
	}
	return opts
}

// --------------------------------------------------

var _ IUpdateGlobalListenerRequest = &UpdateGlobalListenerRequest{}

type UpdateGlobalListenerRequest struct {
	AllowedCidrs      string  `json:"allowedCidrs"`
	TimeoutClient     int     `json:"timeoutClient"`
	TimeoutMember     int     `json:"timeoutMember"`
	TimeoutConnection int     `json:"timeoutConnection"`
	Headers           *string `json:"headers"`
	GlobalPoolId      string  `json:"globalPoolId"`

	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
}

func (s *UpdateGlobalListenerRequest) WithAllowedCidrs(pcidrs ...string) IUpdateGlobalListenerRequest {
	s.AllowedCidrs = strings.Join(pcidrs, ",")
	return s
}

func (s *UpdateGlobalListenerRequest) WithTimeoutClient(ptoc int) IUpdateGlobalListenerRequest {
	s.TimeoutClient = ptoc
	return s
}

func (s *UpdateGlobalListenerRequest) WithTimeoutMember(ptom int) IUpdateGlobalListenerRequest {
	s.TimeoutMember = ptom
	return s
}

func (s *UpdateGlobalListenerRequest) WithTimeoutConnection(ptoc int) IUpdateGlobalListenerRequest {
	s.TimeoutConnection = ptoc
	return s
}

func (s *UpdateGlobalListenerRequest) WithHeaders(pheaders ...string) IUpdateGlobalListenerRequest {
	h := strings.Join(pheaders, ",")
	s.Headers = &h
	return s
}

func (s *UpdateGlobalListenerRequest) WithGlobalPoolId(ppoolId string) IUpdateGlobalListenerRequest {
	s.GlobalPoolId = ppoolId
	return s
}

func (s *UpdateGlobalListenerRequest) WithLoadBalancerId(plbid string) IUpdateGlobalListenerRequest {
	s.LoadBalancerId = plbid
	return s
}

func (s *UpdateGlobalListenerRequest) WithListenerId(plid string) IUpdateGlobalListenerRequest {
	s.ListenerId = plid
	return s
}

func (s *UpdateGlobalListenerRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateGlobalListenerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"allowedCidrs":      s.AllowedCidrs,
		"timeoutClient":     s.TimeoutClient,
		"timeoutMember":     s.TimeoutMember,
		"timeoutConnection": s.TimeoutConnection,
		"headers":           s.Headers,
		"globalPoolId":      s.GlobalPoolId,
	}
}

func (s *UpdateGlobalListenerRequest) AddUserAgent(pagent ...string) IUpdateGlobalListenerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewUpdateGlobalListenerRequest(plbId, plId string) IUpdateGlobalListenerRequest {
	opts := &UpdateGlobalListenerRequest{
		AllowedCidrs:      "0.0.0.0/0",
		TimeoutClient:     50,
		TimeoutMember:     50,
		TimeoutConnection: 5,
		Headers:           nil,
		GlobalPoolId:      "",
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		ListenerCommon: lscommon.ListenerCommon{
			ListenerId: plId,
		},
	}
	return opts
}

// --------------------------------------------------

var _ IDeleteGlobalListenerRequest = &DeleteGlobalListenerRequest{}

type DeleteGlobalListenerRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
}

func (s *DeleteGlobalListenerRequest) WithLoadBalancerId(plbid string) IDeleteGlobalListenerRequest {
	s.LoadBalancerId = plbid
	return s
}

func (s *DeleteGlobalListenerRequest) WithListenerId(plid string) IDeleteGlobalListenerRequest {
	s.ListenerId = plid
	return s
}

func (s *DeleteGlobalListenerRequest) AddUserAgent(pagent ...string) IDeleteGlobalListenerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewDeleteGlobalListenerRequest(plbId, plId string) IDeleteGlobalListenerRequest {
	opts := &DeleteGlobalListenerRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		ListenerCommon: lscommon.ListenerCommon{
			ListenerId: plId,
		},
	}
	return opts
}

// --------------------------------------------------

var _ IGetGlobalListenerRequest = &GetGlobalListenerRequest{}

type GetGlobalListenerRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
}

func (s *GetGlobalListenerRequest) WithLoadBalancerId(plbid string) IGetGlobalListenerRequest {
	s.LoadBalancerId = plbid
	return s
}

func (s *GetGlobalListenerRequest) WithListenerId(plid string) IGetGlobalListenerRequest {
	s.ListenerId = plid
	return s
}

func (s *GetGlobalListenerRequest) AddUserAgent(pagent ...string) IGetGlobalListenerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewGetGlobalListenerRequest(plbId, plId string) IGetGlobalListenerRequest {
	opts := &GetGlobalListenerRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		ListenerCommon: lscommon.ListenerCommon{
			ListenerId: plId,
		},
	}
	return opts
}
