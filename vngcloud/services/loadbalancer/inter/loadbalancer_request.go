package inter

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

func NewCreateLoadBalancerRequest(pname, ppackageId, pbeSubnetId, psubnetId string) ICreateLoadBalancerRequest {
	return &CreateLoadBalancerRequest{
		Name:            pname,
		PackageID:       ppackageId,
		Scheme:          InterVpcLoadBalancerScheme,
		BackEndSubnetId: pbeSubnetId,
		SubnetID:        psubnetId,
		Type:            CreateOptsTypeOptLayer4,
	}
}

type CreateLoadBalancerRequest struct {
	Name            string                 `json:"name"`
	PackageID       string                 `json:"packageId"`
	Scheme          LoadBalancerScheme     `json:"scheme"`
	SubnetID        string                 `json:"subnetId"`
	BackEndSubnetId string                 `json:"backendSubnetId,omitempty"`
	ProjectId       string                 `json:"projectId,omitempty"`
	Type            LoadBalancerType       `json:"type"`
	Listener        *CreateListenerRequest `json:"listener,omitempty"`
	Pool            *CreatePoolRequest     `json:"pool,omitempty"`
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
