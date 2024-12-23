package global

import (
	lfmt "fmt"
	ljparser "github.com/cuongpiger/joat/parser"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
	lstr "strings"
)

func NewListGlobalLoadBalancersRequest(offset, limit int) IListGlobalLoadBalancersRequest {
	opts := new(ListGlobalLoadBalancersRequest)
	opts.Offset = offset
	opts.Limit = limit
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

	var tuples []string
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

// const (
// 	InternalLoadBalancerScheme LoadBalancerScheme = "Internal"
// 	InternetLoadBalancerScheme LoadBalancerScheme = "Internet"
// 	InterVpcLoadBalancerScheme LoadBalancerScheme = "InterVPC"
// )

// const (
// 	LoadBalancerTypeLayer4 LoadBalancerType = "Layer 4"
// 	LoadBalancerTypeLayer7 LoadBalancerType = "Layer 7"
// )

// type (
// 	LoadBalancerScheme string
// 	LoadBalancerType   string
// )

// func NewCreateLoadBalancerRequest(pname, ppackageId, psubnetId string) ICreateLoadBalancerRequest {
// 	return &CreateLoadBalancerRequest{
// 		Name:      pname,
// 		PackageID: ppackageId,
// 		Scheme:    InternetLoadBalancerScheme,
// 		SubnetID:  psubnetId,
// 		Type:      LoadBalancerTypeLayer4,
// 	}
// }

// func NewResizeLoadBalancerRequest(plbId, packageID string) IResizeLoadBalancerRequest {
// 	return &ResizeLoadBalancerRequest{
// 		LoadBalancerCommon: lscommon.LoadBalancerCommon{
// 			LoadBalancerId: plbId,
// 		},
// 		PackageID: packageID,
// 	}
// }

// func NewListGlobalLoadBalancerPackagesRequest() IListGlobalLoadBalancerPackagesRequest {
// 	return new(ListGlobalLoadBalancerPackagesRequest)
// }

// func NewGetGlobalLoadBalancerByIdRequest(plbId string) IGetGlobalLoadBalancerByIdRequest {
// 	opts := new(GetGlobalLoadBalancerByIdRequest)
// 	opts.LoadBalancerId = plbId
// 	return opts
// }

// func (s *ListGlobalLoadBalancersRequest) WithName(pname string) IListGlobalLoadBalancersRequest {
// 	s.Name = pname
// 	return s
// }

// func NewDeleteLoadBalancerByIdRequest(plbId string) IDeleteLoadBalancerByIdRequest {
// 	opts := new(DeleteLoadBalancerByIdRequest)
// 	opts.LoadBalancerId = plbId
// 	return opts
// }

// type CreateLoadBalancerRequest struct {
// 	Name         string                 `json:"name"`
// 	PackageID    string                 `json:"packageId"`
// 	Scheme       LoadBalancerScheme     `json:"scheme"`
// 	AutoScalable bool                   `json:"autoScalable"`
// 	SubnetID     string                 `json:"subnetId"`
// 	Type         LoadBalancerType       `json:"type"`
// 	Listener     ICreateListenerRequest `json:"listener"`
// 	Pool         ICreatePoolRequest     `json:"pool"`
// 	Tags         []lscommon.Tag         `json:"tags,omitempty"`
// 	IsPoc        bool                   `json:"isPoc"`

// 	lscommon.UserAgent
// }

// type ResizeLoadBalancerRequest struct {
// 	PackageID string `json:"packageId"`
// 	lscommon.UserAgent
// 	lscommon.LoadBalancerCommon
// }

// type ListGlobalLoadBalancerPackagesRequest struct {
// 	lscommon.UserAgent
// }

// type GetGlobalLoadBalancerByIdRequest struct {
// 	lscommon.UserAgent
// 	lscommon.LoadBalancerCommon
// }

// type DeleteLoadBalancerByIdRequest struct {
// 	lscommon.UserAgent
// 	lscommon.LoadBalancerCommon
// }

// type ResizeLoadBalancerByIdRequest struct {
// 	lscommon.UserAgent
// 	lscommon.LoadBalancerCommon

// 	PackageId string `json:"packageId"`
// }

// func (s *CreateLoadBalancerRequest) ToMap() map[string]interface{} {
// 	err := map[string]interface{}{
// 		"name":         s.Name,
// 		"packageId":    s.PackageID,
// 		"scheme":       s.Scheme,
// 		"autoScalable": s.AutoScalable,
// 		"subnetId":     s.SubnetID,
// 		"type":         s.Type,
// 		"tags":         s.Tags,
// 	}

// 	if s.Listener != nil {
// 		err["listener"] = s.Listener.ToMap()
// 	}

// 	if s.Pool != nil {
// 		err["pool"] = s.Pool.ToMap()
// 	}

// 	return err
// }

// func (s *CreateLoadBalancerRequest) ToRequestBody() interface{} {
// 	if s.Pool != nil {
// 		s.Pool = s.Pool.ToRequestBody().(*CreatePoolRequest)
// 	}

// 	if s.Listener != nil {
// 		s.Listener = s.Listener.ToRequestBody().(*CreateListenerRequest)
// 	}

// 	return s
// }

// func (s *CreateLoadBalancerRequest) AddUserAgent(pagent ...string) ICreateLoadBalancerRequest {
// 	s.UserAgent.AddUserAgent(pagent...)
// 	return s
// }
// func (s *CreateLoadBalancerRequest) WithListener(plistener ICreateListenerRequest) ICreateLoadBalancerRequest {
// 	s.Listener = plistener
// 	return s
// }

// func (s *CreateLoadBalancerRequest) WithPool(ppool ICreatePoolRequest) ICreateLoadBalancerRequest {
// 	s.Pool = ppool
// 	return s
// }

// func (s *CreateLoadBalancerRequest) WithTags(ptags ...string) ICreateLoadBalancerRequest {
// 	if s.Tags == nil {
// 		s.Tags = make([]lscommon.Tag, 0)
// 	}

// 	if len(ptags)%2 != 0 {
// 		ptags = append(ptags, "none")
// 	}

// 	for i := 0; i < len(ptags); i += 2 {
// 		s.Tags = append(s.Tags, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
// 	}

// 	return s
// }

// func (s *CreateLoadBalancerRequest) WithScheme(pscheme LoadBalancerScheme) ICreateLoadBalancerRequest {
// 	s.Scheme = pscheme
// 	return s
// }

// func (s *CreateLoadBalancerRequest) WithAutoScalable(pautoScalable bool) ICreateLoadBalancerRequest {
// 	s.AutoScalable = pautoScalable
// 	return s
// }

// func (s *CreateLoadBalancerRequest) WithPackageId(ppackageId string) ICreateLoadBalancerRequest {
// 	s.PackageID = ppackageId
// 	return s
// }

// func (s *CreateLoadBalancerRequest) WithSubnetId(psubnetId string) ICreateLoadBalancerRequest {
// 	s.SubnetID = psubnetId
// 	return s
// }

// func (s *CreateLoadBalancerRequest) WithType(ptype LoadBalancerType) ICreateLoadBalancerRequest {
// 	s.Type = ptype
// 	return s
// }

// func (s *CreateLoadBalancerRequest) WithPoc(isPoc bool) ICreateLoadBalancerRequest {
// 	s.IsPoc = isPoc
// 	return s
// }

// func (s *ResizeLoadBalancerRequest) ToRequestBody() interface{} {
// 	return s
// }

// func (s *ResizeLoadBalancerRequest) AddUserAgent(pagent ...string) IResizeLoadBalancerRequest {
// 	s.UserAgent.AddUserAgent(pagent...)
// 	return s
// }

// func (s *ResizeLoadBalancerRequest) WithPackageId(ppackageId string) IResizeLoadBalancerRequest {
// 	s.PackageID = ppackageId
// 	return s
// }

// func (s *ListGlobalLoadBalancerPackagesRequest) AddUserAgent(pagent ...string) IListGlobalLoadBalancerPackagesRequest {
// 	s.UserAgent.AddUserAgent(pagent...)
// 	return s
// }

// func (s *ListGlobalLoadBalancerPackagesRequest) ToMap() map[string]interface{} {
// 	return map[string]interface{}{}
// }

// func (s *GetGlobalLoadBalancerByIdRequest) AddUserAgent(pagent ...string) IGetGlobalLoadBalancerByIdRequest {
// 	s.UserAgent.AddUserAgent(pagent...)
// 	return s
// }

// 	if len(ptags)%2 != 0 {
// 		ptags = append(ptags, "")
// 	}

// 	for i := 0; i < len(ptags); i += 2 {
// 		s.Tags = append(s.Tags, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
// 	}

// 	return s
// }

// 	var tuples []string
// 	for _, tag := range s.Tags {
// 		if tag.Key == "" {
// 			continue
// 		}

// 		tuple := "tags=key:" + tag.Key
// 		if tag.Value != "" {
// 			tuple += ",value:" + tag.Value
// 		}
// 		tuples = append(tuples, tuple)
// 	}

// 	if len(tuples) > 0 {
// 		return url.String() + "&" + lstr.Join(tuples, "&"), nil
// 	}

// 	return url.String(), err
// }

// func (s *ResizeLoadBalancerByIdRequest) ToMap() map[string]interface{} {
// 	return map[string]interface{}{
// 		"packageId":      s.PackageId,
// 		"loadBalancerId": s.LoadBalancerId,
// 	}
// }

// func (s *ResizeLoadBalancerByIdRequest) ToRequestBody() interface{} {
// 	return s
// }
