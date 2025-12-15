package v1

import (
	lfmt "fmt"
	lstr "strings"

	ljparser "github.com/cuongpiger/joat/parser"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

type (
	GlobalLoadBalancerType        string
	GlobalLoadBalancerPaymentFlow string
)

const (
	GlobalLoadBalancerTypeLayer4           GlobalLoadBalancerType        = "Layer 4"
	GlobalLoadBalancerPaymentFlowAutomated GlobalLoadBalancerPaymentFlow = "automated"
)

// --------------------------------------------------------------------------

func NewListGlobalLoadBalancersRequest(offset, limit int) IListGlobalLoadBalancersRequest {
	opts := &ListGlobalLoadBalancersRequest{
		Name:   "",
		Offset: offset,
		Limit:  limit,
		Tags:   make([]lscommon.Tag, 0),
	}
	return opts
}

type ListGlobalLoadBalancersRequest struct {
	Name   string `q:"name,beempty"`
	Offset int    `q:"offset"`
	Limit  int    `q:"limit"`

	Tags []lscommon.Tag
	lscommon.UserAgent
}

func (s *ListGlobalLoadBalancersRequest) WithName(pname string) IListGlobalLoadBalancersRequest {
	s.Name = pname
	return s
}

func (s *ListGlobalLoadBalancersRequest) WithTags(ptags ...string) IListGlobalLoadBalancersRequest {
	if s.Tags == nil {
		s.Tags = make([]lscommon.Tag, 0)
	}

	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.Tags = append(s.Tags, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
	}

	return s
}

func (s *ListGlobalLoadBalancersRequest) ToListQuery() (string, error) {
	parser, _ := ljparser.GetParser()
	url, err := parser.UrlMe(s)
	if err != nil {
		return "", err
	}

	tuples := make([]string, 0, len(s.Tags))
	for _, tag := range s.Tags {
		if tag.Key == "" {
			continue
		}

		tuple := "tags=key:" + tag.Key
		if tag.Value != "" {
			tuple += ",value:" + tag.Value
		}
		tuples = append(tuples, tuple)
	}

	if len(tuples) > 0 {
		return url.String() + "&" + lstr.Join(tuples, "&"), nil
	}

	return url.String(), err
}

func (s *ListGlobalLoadBalancersRequest) GetDefaultQuery() string {
	return lfmt.Sprintf("offset=%d&limit=%d", defaultOffsetListGlobalLoadBalancer, defaultLimitListGlobalLoadBalancer)
}

func (s *ListGlobalLoadBalancersRequest) AddUserAgent(pagent ...string) IListGlobalLoadBalancersRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

// --------------------------------------------------------------------------

var _ ICreateGlobalLoadBalancerRequest = &CreateGlobalLoadBalancerRequest{}

type CreateGlobalLoadBalancerRequest struct {
	Description    string                        `json:"description"`
	Name           string                        `json:"name"`
	Type           GlobalLoadBalancerType        `json:"type"`
	Package        string                        `json:"package"`
	PaymentFlow    GlobalLoadBalancerPaymentFlow `json:"paymentFlow"`
	GlobalListener ICreateGlobalListenerRequest  `json:"globalListener"`
	GlobalPool     ICreateGlobalPoolRequest      `json:"globalPool"`

	lscommon.UserAgent
}

func (s *CreateGlobalLoadBalancerRequest) WithDescription(pdesc string) ICreateGlobalLoadBalancerRequest {
	s.Description = pdesc
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithName(pname string) ICreateGlobalLoadBalancerRequest {
	s.Name = pname
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithType(ptype GlobalLoadBalancerType) ICreateGlobalLoadBalancerRequest {
	s.Type = ptype
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithPackage(ppackageId string) ICreateGlobalLoadBalancerRequest {
	s.Package = ppackageId
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithPaymentFlow(ppaymentFlow GlobalLoadBalancerPaymentFlow) ICreateGlobalLoadBalancerRequest {
	s.PaymentFlow = ppaymentFlow
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithGlobalListener(plistener ICreateGlobalListenerRequest) ICreateGlobalLoadBalancerRequest {
	s.GlobalListener = plistener
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithGlobalPool(ppool ICreateGlobalPoolRequest) ICreateGlobalLoadBalancerRequest {
	s.GlobalPool = ppool
	return s
}

func (s *CreateGlobalLoadBalancerRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateGlobalLoadBalancerRequest) AddUserAgent(pagent ...string) ICreateGlobalLoadBalancerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreateGlobalLoadBalancerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"description":    s.Description,
		"name":           s.Name,
		"type":           s.Type,
		"globalListener": s.GlobalListener,
		"globalPool":     s.GlobalPool,
	}
}

func NewCreateGlobalLoadBalancerRequest(name string) ICreateGlobalLoadBalancerRequest {
	opts := &CreateGlobalLoadBalancerRequest{
		Description:    "",
		Name:           name,
		Type:           GlobalLoadBalancerTypeLayer4,
		Package:        "",
		PaymentFlow:    GlobalLoadBalancerPaymentFlowAutomated,
		GlobalListener: nil,
		GlobalPool:     nil,
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IDeleteGlobalLoadBalancerRequest = &DeleteGlobalLoadBalancerRequest{}

type DeleteGlobalLoadBalancerRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

func (s *DeleteGlobalLoadBalancerRequest) WithLoadBalancerId(plbId string) IDeleteGlobalLoadBalancerRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *DeleteGlobalLoadBalancerRequest) AddUserAgent(pagent ...string) IDeleteGlobalLoadBalancerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewDeleteGlobalLoadBalancerRequest(lbId string) IDeleteGlobalLoadBalancerRequest {
	opts := &DeleteGlobalLoadBalancerRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IListGlobalPackagesRequest = &ListGlobalPackagesRequest{}

type ListGlobalPackagesRequest struct {
	lscommon.UserAgent
}

func (s *ListGlobalPackagesRequest) AddUserAgent(pagent ...string) IListGlobalPackagesRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewListGlobalPackagesRequest() IListGlobalPackagesRequest {
	opts := &ListGlobalPackagesRequest{}
	return opts
}

// --------------------------------------------------------------------------

var _ IListGlobalRegionsRequest = &ListGlobalRegionsRequest{}

type ListGlobalRegionsRequest struct {
	lscommon.UserAgent
}

func (s *ListGlobalRegionsRequest) AddUserAgent(pagent ...string) IListGlobalRegionsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewListGlobalRegionsRequest() IListGlobalRegionsRequest {
	opts := &ListGlobalRegionsRequest{}
	return opts
}

// --------------------------------------------------------------------------

var _ IGetGlobalLoadBalancerUsageHistoriesRequest = &GetGlobalLoadBalancerUsageHistoriesRequest{}

type GetGlobalLoadBalancerUsageHistoriesRequest struct {
	From string `q:"from"`
	To   string `q:"to"`
	Type string `q:"type"`

	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) WithLoadBalancerId(plbId string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) WithFrom(pfrom string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.From = pfrom
	return s
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) WithTo(pto string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.To = pto
	return s
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) WithType(ptype string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.Type = ptype
	return s
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) ToListQuery() (string, error) {
	parser, _ := ljparser.GetParser()
	url, err := parser.UrlMe(s)
	return url.String(), err
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) GetDefaultQuery() string {
	return ""
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) AddUserAgent(pagent ...string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewGetGlobalLoadBalancerUsageHistoriesRequest(lbId, from, to, usageType string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	opts := &GetGlobalLoadBalancerUsageHistoriesRequest{
		From: from,
		To:   to,
		Type: usageType,
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IGetGlobalLoadBalancerByIdRequest = &GetGlobalLoadBalancerByIdRequest{}

type GetGlobalLoadBalancerByIdRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

func (s *GetGlobalLoadBalancerByIdRequest) WithLoadBalancerId(plbId string) IGetGlobalLoadBalancerByIdRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *GetGlobalLoadBalancerByIdRequest) AddUserAgent(pagent ...string) IGetGlobalLoadBalancerByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewGetGlobalLoadBalancerByIdRequest(lbId string) IGetGlobalLoadBalancerByIdRequest {
	opts := &GetGlobalLoadBalancerByIdRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
	}
	return opts
}
