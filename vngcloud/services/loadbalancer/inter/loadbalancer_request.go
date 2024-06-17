package inter

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

func NewCreateLoadBalancerRequest(puserId, pname, ppackageId, pbeSubnetId, psubnetId string) ICreateLoadBalancerRequest {
	opt := new(CreateLoadBalancerRequest)
	opt.SetPortalUserId(puserId)
	opt.Name = pname
	opt.PackageID = ppackageId
	opt.Scheme = InterVpcLoadBalancerScheme
	opt.BackEndSubnetId = pbeSubnetId
	opt.SubnetID = psubnetId
	opt.Type = CreateOptsTypeOptLayer4
	return opt
}

type CreateLoadBalancerRequest struct {
	Name            string                 `json:"name"`
	PackageID       string                 `json:"packageId"`
	Scheme          LoadBalancerScheme     `json:"scheme"`
	SubnetID        string                 `json:"subnetId"`
	BackEndSubnetId string                 `json:"backendSubnetId,omitempty"`
	ProjectId       string                 `json:"projectId,omitempty"`
	Type            LoadBalancerType       `json:"type"`
	Listener        ICreateListenerRequest `json:"listener,omitempty"`
	Pool            ICreatePoolRequest     `json:"pool,omitempty"`

	lscommon.PortalUser
	lscommon.UserAgent
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

func (s *CreateLoadBalancerRequest) WithProjectId(pprojectId string) ICreateLoadBalancerRequest {
	s.ProjectId = pprojectId
	return s
}

func (s *CreateLoadBalancerRequest) AddUserAgent(pagent ...string) ICreateLoadBalancerRequest {
	s.UserAgent.Agent = append(s.UserAgent.Agent, pagent...)
	return s
}

func (s *CreateLoadBalancerRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *CreateLoadBalancerRequest) WithListener(plistener ICreateListenerRequest) ICreateLoadBalancerRequest {
	s.Listener = plistener
	return s
}

func (s *CreateLoadBalancerRequest) WithPool(ppool ICreatePoolRequest) ICreateLoadBalancerRequest {
	s.Pool = ppool
	return s
}
