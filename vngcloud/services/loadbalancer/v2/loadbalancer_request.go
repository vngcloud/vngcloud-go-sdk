package v2

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

type CreateLoadBalancerRequest struct {
	Name      string                 `json:"name"`
	PackageID string                 `json:"packageId"`
	Scheme    LoadBalancerScheme     `json:"scheme"`
	SubnetID  string                 `json:"subnetId"`
	Type      LoadBalancerType       `json:"type"`
	Listener  *CreateListenerRequest `json:"listener"`
	Pool      *CreatePoolRequest     `json:"pool"`
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
