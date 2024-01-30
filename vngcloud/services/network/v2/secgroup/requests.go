package secgroup

import (
	lParser "github.com/cuongpiger/joat/parser"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lSecgroupCommonV2 "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/network/v2"
)

// ****************************************************** CreateOpts ***************************************************

type CreateOpts struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	common.CommonOpts
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

type DeleteOpts struct {
	lSecgroupCommonV2.SecgroupV2Common
	common.CommonOpts
}

// ****************************************************** GetOpts ******************************************************

type GetOpts struct {
	common.CommonOpts
	lSecgroupCommonV2.SecgroupV2Common
}

// ***************************************************** ListOpts ******************************************************

type ListOpts struct {
	Name string `q:"name,beempty"`
	common.CommonOpts
}

func (s *ListOpts) ToListQuery() (string, error) {
	parser, _ := lParser.GetParser()
	url, err := parser.UrlMe(s)

	if err != nil {
		return "", err
	}

	return url.String(), err
}
