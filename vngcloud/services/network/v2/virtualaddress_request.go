package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

/**
 * The group of Virtual Address APIs
 */

// API Create virtual address cross project
type CreateVirtualAddressCrossProjectRequest struct {
	// Name is the name of the virtual address.
	Name string `json:"name"` // required

	// Description is the description of the virtual address.
	Description string `json:"description"`

	// Contains the required information to create a virtual address cross project.
	CrossProjectRequest struct {
		// The project ID whose virtual address will be created.
		ProjectId string `json:"projectId"` // required

		// The subnet ID where the virtual address will be created.
		SubnetId string `json:"subnetId"` // required

		// The IP address of the virtual address.
		IpAddress string `json:"ipAddress"`
	} `json:"crossProjectRequest"` // required

	// others...
	lscommon.UserAgent
}

func (s *CreateVirtualAddressCrossProjectRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateVirtualAddressCrossProjectRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":        s.Name,
		"description": s.Description,
		"crossProjectRequest": map[string]interface{}{
			"projectId": s.CrossProjectRequest.ProjectId,
			"subnetId":  s.CrossProjectRequest.SubnetId,
			"ipAddress": s.CrossProjectRequest.IpAddress,
		},
	}
}

func (s *CreateVirtualAddressCrossProjectRequest) AddUserAgent(pagent ...string) ICreateVirtualAddressCrossProjectRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreateVirtualAddressCrossProjectRequest) WithDescription(pdescription string) ICreateVirtualAddressCrossProjectRequest {
	s.Description = pdescription
	return s
}

// API Delete virtual address by ID
type DeleteVirtualAddressByIdRequest struct {
	lscommon.VirtualAddressCommon
	lscommon.UserAgent
}

func (s *DeleteVirtualAddressByIdRequest) AddUserAgent(pagent ...string) IDeleteVirtualAddressByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *DeleteVirtualAddressByIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"virtualAddressId": s.VirtualAddressId,
	}
}

// Api Get virtual address by ID

type GetVirtualAddressByIdRequest struct {
	lscommon.VirtualAddressCommon
	lscommon.UserAgent
}

func (s *GetVirtualAddressByIdRequest) AddUserAgent(pagent ...string) IGetVirtualAddressByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *GetVirtualAddressByIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"virtualAddressId": s.VirtualAddressId,
	}
}

// Api list address pairs by virtual address ID

type ListAddressPairsByVirtualAddressIdRequest struct {
	lscommon.VirtualAddressCommon
	lscommon.UserAgent
}

func (s *ListAddressPairsByVirtualAddressIdRequest) AddUserAgent(pagent ...string) IListAddressPairsByVirtualAddressIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *ListAddressPairsByVirtualAddressIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"virtualAddressId": s.VirtualAddressId,
	}
}
