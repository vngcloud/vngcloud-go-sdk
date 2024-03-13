package loadbalancer

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lbCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2"
)

// ************************************************* CreateOptsBuilder *************************************************

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
	Name      string              `json:"name"`
	PackageID string              `json:"packageId"`
	Scheme    CreateOptsSchemeOpt `json:"scheme"`
	SubnetID  string              `json:"subnetId"`
	Type      CreateOptsTypeOpt   `json:"type"`
	common.CommonOpts
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

// ************************************************** GetOptsBuilder ***************************************************

type GetOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

// ************************************************* DeleteOptsBuilder *************************************************

type DeleteOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

// ********************************************* GetBySubnetIDOptsBuilder **********************************************

type ListBySubnetIDOptsBuilder struct {
	SubnetID string
	common.CommonOpts
}

func (s *ListBySubnetIDOptsBuilder) GetSubnetID() string {
	return s.SubnetID
}

// ************************************************** ListOptsBuilder **************************************************

type ListOptsBuilder struct {
	common.CommonOpts
}

// ************************************************* ResizeOptsBuilder *************************************************
type ResizeOpts struct {
	PackageId string `json:"packageId"`

	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

func (s *ResizeOpts) ToRequestBody() interface{} {
	return s
}
