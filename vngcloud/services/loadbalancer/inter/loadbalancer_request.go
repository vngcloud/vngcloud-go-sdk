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
	Tags            []lscommon.Tag         `json:"tags,omitempty"`
	ZoneId          *lscommon.Zone         `json:"zoneId,omitempty"`

	lscommon.PortalUser
	lscommon.UserAgent
}

func (s *CreateLoadBalancerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":            s.Name,
		"packageId":       s.PackageID,
		"scheme":          s.Scheme,
		"subnetId":        s.SubnetID,
		"backendSubnetId": s.BackEndSubnetId,
		"projectId":       s.ProjectId,
		"type":            s.Type,
		"tags":            s.Tags,
		"zoneId":          s.ZoneId,
	}
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
	s.Agent = append(s.Agent, pagent...)
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

func (s *CreateLoadBalancerRequest) WithZoneId(pzoneId lscommon.Zone) ICreateLoadBalancerRequest {
	s.ZoneId = &pzoneId
	return s
}
