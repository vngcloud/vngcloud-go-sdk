package loadbalancer

import (
	lsCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lbCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2"
	lsListener "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2/listener"
	lsPool "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2/pool"
)

const (
	// For internet facing
	CreateOptsSchemeOptInternet CreateOptsSchemeOpt = "Internet"
	CreateOptsSchemeOptInternal CreateOptsSchemeOpt = "Internal"

	// For lb type
	CreateOptsTypeOptLayer4 CreateOptsTypeOpt = "Layer 4"
	CreateOptsTypeOptLayer7 CreateOptsTypeOpt = "Layer 7"
)

type (
	CreateOptsSchemeOpt string
	CreateOptsTypeOpt   string
)

type CreateOpts struct {
	Name      string                 `json:"name"`
	PackageID string                 `json:"packageId"`
	Scheme    CreateOptsSchemeOpt    `json:"scheme"`
	SubnetID  string                 `json:"subnetId"`
	Type      CreateOptsTypeOpt      `json:"type"`
	Listener  *lsListener.CreateOpts `json:"listener"`
	Pool      *lsPool.CreateOpts     `json:"pool"`

	lsCm.CommonOpts
}

func (s *CreateOpts) ToRequestBody() interface{} {
	if s.Pool != nil {
		s.Pool = s.Pool.ToRequestBody().(*lsPool.CreateOpts)
	}

	if s.Listener != nil {
		s.Listener = s.Listener.ToRequestBody().(*lsListener.CreateOpts)
	}
	return s
}

type GetOpts struct {
	lsCm.CommonOpts
	lbCm.LoadBalancerV2Common
}

type DeleteOpts struct {
	lsCm.CommonOpts
	lbCm.LoadBalancerV2Common
}

type ListBySubnetIDOpts struct {
	SubnetID string
	lsCm.CommonOpts
}

func (s *ListBySubnetIDOpts) GetSubnetID() string {
	return s.SubnetID
}

type ListOpts struct {
	lsCm.CommonOpts
}

type UpdateOpts struct {
	PackageID string `json:"packageId"`

	lsCm.CommonOpts
	lbCm.LoadBalancerV2Common
}

func (s *UpdateOpts) ToRequestBody() interface{} {
	return s
}

type (
	CreateTagOpts struct {
		ResourceID     string             `json:"resourceId"`
		ResourceType   string             `json:"resourceType"`
		TagRequestList []CreateTagOptsTag `json:"tagRequestList"`
		lsCm.CommonOpts
		lbCm.LoadBalancerV2Common
	}

	CreateTagOptsTag struct {
		IsEdited bool   `json:"isEdited"`
		Key      string `json:"key"`
		Value    string `json:"value"`
	}
)

func (s *CreateTagOpts) ToRequestBody() interface{} {
	return s
}
