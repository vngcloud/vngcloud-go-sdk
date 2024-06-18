package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

const (
	InterVpcLoadBalancerScheme LoadBalancerScheme = "InterVPC"
	InternalLoadBalancerScheme LoadBalancerScheme = "Internal"
	InternetLoadBalancerScheme LoadBalancerScheme = "Internet"
)

const (
	CreateOptsTypeOptLayer4 LoadBalancerType = "Layer 4"
	CreateOptsTypeOptLayer7 LoadBalancerType = "Layer 7"
)

type (
	LoadBalancerScheme string
	LoadBalancerType   string
)

func NewCreateLoadBalancerRequest(pname, ppackageId, psubnetId string) ICreateLoadBalancerRequest {
	return &CreateLoadBalancerRequest{
		Name:      pname,
		PackageID: ppackageId,
		Scheme:    InternetLoadBalancerScheme,
		SubnetID:  psubnetId,
		Type:      CreateOptsTypeOptLayer4,
	}
}

func NewGetLoadBalancerByIdRequest(plbId string) IGetLoadBalancerByIdRequest {
	opts := new(GetLoadBalancerByIdRequest)
	opts.LoadBalancerId = plbId
	return opts
}

type CreateLoadBalancerRequest struct {
	Name      string                 `json:"name"`
	PackageID string                 `json:"packageId"`
	Scheme    LoadBalancerScheme     `json:"scheme"`
	SubnetID  string                 `json:"subnetId"`
	Type      LoadBalancerType       `json:"type"`
	Listener  ICreateListenerRequest `json:"listener"`
	Pool      ICreatePoolRequest     `json:"pool"`
	Tags      []lscommon.Tag         `json:"tags,omitempty"`

	lscommon.UserAgent
}

type GetLoadBalancerByIdRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

func (s *CreateLoadBalancerRequest) ToRequestBody() interface{} {
	if s.Pool != nil {
		s.Pool = s.Pool.ToRequestBody().(*CreatePoolRequest)
	}

	if s.Listener != nil {
		s.Listener = s.Listener.ToRequestBody().(*CreateListenerRequest)
	}

	return s
}

func (s *CreateLoadBalancerRequest) AddUserAgent(pagent ...string) ICreateLoadBalancerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
func (s *CreateLoadBalancerRequest) WithListener(plistener ICreateListenerRequest) ICreateLoadBalancerRequest {
	s.Listener = plistener
	return s
}

func (s *CreateLoadBalancerRequest) WithPool(ppool ICreatePoolRequest) ICreateLoadBalancerRequest {
	s.Pool = ppool
	return s
}

func (s *CreateLoadBalancerRequest) WithTags(ptags ...string) ICreateLoadBalancerRequest {
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

func (s *GetLoadBalancerByIdRequest) AddUserAgent(pagent ...string) IGetLoadBalancerByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
