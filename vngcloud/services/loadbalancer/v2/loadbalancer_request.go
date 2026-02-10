package v2

import (
	lfmt "fmt"
	lstr "strings"

	ljparser "github.com/cuongpiger/joat/parser"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

const (
	InternalLoadBalancerScheme LoadBalancerScheme = "Internal"
	InternetLoadBalancerScheme LoadBalancerScheme = "Internet"
	InterVpcLoadBalancerScheme LoadBalancerScheme = "InterVPC"
)

const (
	LoadBalancerTypeLayer4 LoadBalancerType = "Layer 4"
	LoadBalancerTypeLayer7 LoadBalancerType = "Layer 7"
)

type (
	LoadBalancerScheme string
	LoadBalancerType   string
)

// ScalingConfig defines scaling policies (min/max nodes)
type ScalingConfig struct {
	MinNodes int `json:"minSize"`
	MaxNodes int `json:"maxSize"`
}

// NetworkingConfig defines networking topology (subnets)
type NetworkingConfig struct {
	Subnets []string `json:"subnets,omitempty"`
}

func NewCreateLoadBalancerRequest(pname, ppackageId, psubnetId string) ICreateLoadBalancerRequest {
	return &CreateLoadBalancerRequest{
		Name:      pname,
		PackageID: ppackageId,
		Scheme:    InternetLoadBalancerScheme,
		SubnetID:  psubnetId,
		Type:      LoadBalancerTypeLayer4,
	}
}

func NewResizeLoadBalancerRequest(plbId, packageID string) IResizeLoadBalancerRequest {
	return &ResizeLoadBalancerRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		PackageID: packageID,
	}
}

func NewListLoadBalancerPackagesRequest() IListLoadBalancerPackagesRequest {
	return &ListLoadBalancerPackagesRequest{}
}

func NewGetLoadBalancerByIdRequest(plbId string) IGetLoadBalancerByIdRequest {
	opts := new(GetLoadBalancerByIdRequest)
	opts.LoadBalancerId = plbId
	return opts
}

func NewListLoadBalancersRequest(ppage, psize int) IListLoadBalancersRequest {
	opts := new(ListLoadBalancersRequest)
	opts.Page = ppage
	opts.Size = psize
	return opts
}

func (s *ListLoadBalancersRequest) AddUserAgent(pagent ...string) IListLoadBalancersRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewDeleteLoadBalancerByIdRequest(plbId string) IDeleteLoadBalancerByIdRequest {
	opts := new(DeleteLoadBalancerByIdRequest)
	opts.LoadBalancerId = plbId
	return opts
}

func NewScaleLoadBalancerRequest(plbId string) IScaleLoadBalancerRequest {
	return &ScaleLoadBalancerRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
	}
}

type CreateLoadBalancerRequest struct {
	Name         string                 `json:"name"`
	PackageID    string                 `json:"packageId"`
	Scheme       LoadBalancerScheme     `json:"scheme"`
	AutoScalable bool                   `json:"autoScalable"`
	SubnetID     string                 `json:"subnetId"`
	Type         LoadBalancerType       `json:"type"`
	Listener     ICreateListenerRequest `json:"listener"`
	Pool         ICreatePoolRequest     `json:"pool"`
	Tags         []lscommon.Tag         `json:"tags,omitempty"`
	IsPoc        bool                   `json:"isPoc"`
	ZoneId       *lscommon.Zone         `json:"zoneId"`

	lscommon.UserAgent
}

type ResizeLoadBalancerRequest struct {
	PackageID string `json:"packageId"`
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

type ListLoadBalancerPackagesRequest struct {
	lscommon.UserAgent
	ZoneId lscommon.Zone `q:"zoneId,beempty"`
}

type GetLoadBalancerByIdRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

type ListLoadBalancersRequest struct {
	Name string `q:"name,beempty"`
	Page int    `q:"page"`
	Size int    `q:"size"`

	Tags []lscommon.Tag
	lscommon.UserAgent
}

type DeleteLoadBalancerByIdRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

func (s *DeleteLoadBalancerByIdRequest) AddUserAgent(pagent ...string) IDeleteLoadBalancerByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type ResizeLoadBalancerByIdRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon

	PackageId string `json:"packageId"`
}

func (s *ResizeLoadBalancerByIdRequest) AddUserAgent(pagent ...string) IResizeLoadBalancerByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type ScaleLoadBalancerRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon

	Scaling    *ScalingConfig    `json:"scaling"`
	Networking *NetworkingConfig `json:"networking"`
}

func (s *ScaleLoadBalancerRequest) AddUserAgent(pagent ...string) IScaleLoadBalancerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *ScaleLoadBalancerRequest) WithScaling(pscaling *ScalingConfig) IScaleLoadBalancerRequest {
	s.Scaling = pscaling
	return s
}

func (s *ScaleLoadBalancerRequest) WithNetworking(pnetworking *NetworkingConfig) IScaleLoadBalancerRequest {
	s.Networking = pnetworking
	return s
}

func (s *ScaleLoadBalancerRequest) ToRequestBody() interface{} {
	return s
}

func (s *ScaleLoadBalancerRequest) ToMap() map[string]interface{} {
	result := map[string]interface{}{}
	if s.Scaling != nil {
		result["scaling"] = map[string]interface{}{
			"minSize": s.Scaling.MinNodes,
			"maxSize": s.Scaling.MaxNodes,
		}
	}
	if s.Networking != nil {
		result["networking"] = map[string]interface{}{
			"subnets": s.Networking.Subnets,
		}
	}
	return result
}

func (s *CreateLoadBalancerRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"name":         s.Name,
		"packageId":    s.PackageID,
		"scheme":       s.Scheme,
		"autoScalable": s.AutoScalable,
		"subnetId":     s.SubnetID,
		"type":         s.Type,
		"tags":         s.Tags,
	}

	if s.Listener != nil {
		err["listener"] = s.Listener.ToMap()
	}

	if s.Pool != nil {
		err["pool"] = s.Pool.ToMap()
	}

	return err
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

func (s *CreateLoadBalancerRequest) WithScheme(pscheme LoadBalancerScheme) ICreateLoadBalancerRequest {
	s.Scheme = pscheme
	return s
}

func (s *CreateLoadBalancerRequest) WithAutoScalable(pautoScalable bool) ICreateLoadBalancerRequest {
	s.AutoScalable = pautoScalable
	return s
}

func (s *CreateLoadBalancerRequest) WithPackageId(ppackageId string) ICreateLoadBalancerRequest {
	s.PackageID = ppackageId
	return s
}

func (s *CreateLoadBalancerRequest) WithSubnetId(psubnetId string) ICreateLoadBalancerRequest {
	s.SubnetID = psubnetId
	return s
}

func (s *CreateLoadBalancerRequest) WithType(ptype LoadBalancerType) ICreateLoadBalancerRequest {
	s.Type = ptype
	return s
}

func (s *CreateLoadBalancerRequest) WithPoc(isPoc bool) ICreateLoadBalancerRequest {
	s.IsPoc = isPoc
	return s
}

func (s *CreateLoadBalancerRequest) WithZoneId(pzoneId lscommon.Zone) ICreateLoadBalancerRequest {
	s.ZoneId = &pzoneId
	return s
}

func (s *ResizeLoadBalancerRequest) ToRequestBody() interface{} {
	return s
}

func (s *ResizeLoadBalancerRequest) AddUserAgent(pagent ...string) IResizeLoadBalancerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *ResizeLoadBalancerRequest) WithPackageId(ppackageId string) IResizeLoadBalancerRequest {
	s.PackageID = ppackageId
	return s
}

func (s *ListLoadBalancerPackagesRequest) AddUserAgent(pagent ...string) IListLoadBalancerPackagesRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *ListLoadBalancerPackagesRequest) WithZoneId(pzoneId lscommon.Zone) IListLoadBalancerPackagesRequest {
	s.ZoneId = pzoneId
	return s
}

func (s *ListLoadBalancerPackagesRequest) GetZoneId() string {
	return string(s.ZoneId)
}

func (s *ListLoadBalancerPackagesRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{}
}

func (s *GetLoadBalancerByIdRequest) AddUserAgent(pagent ...string) IGetLoadBalancerByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *ListLoadBalancersRequest) WithName(pname string) IListLoadBalancersRequest {
	s.Name = pname
	return s
}

func (s *ListLoadBalancersRequest) WithTags(ptags ...string) IListLoadBalancersRequest {
	if s.Tags == nil {
		s.Tags = make([]lscommon.Tag, 0)
	}

	if len(ptags)%2 != 0 {
		ptags = append(ptags, "")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.Tags = append(s.Tags, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
	}

	return s
}

func (s *ListLoadBalancersRequest) ToListQuery() (string, error) {
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

func (s *ListLoadBalancersRequest) GetDefaultQuery() string {
	return lfmt.Sprintf("name=&page=%d&size=%d", defaultPageListLoadBalancer, defaultSizeListLoadBalancer)
}

func (s *ResizeLoadBalancerByIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"packageId":      s.PackageId,
		"loadBalancerId": s.LoadBalancerId,
	}
}

func (s *ResizeLoadBalancerByIdRequest) ToRequestBody() interface{} {
	return s
}
