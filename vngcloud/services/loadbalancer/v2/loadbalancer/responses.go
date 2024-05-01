package loadbalancer

import (
	lStr "strings"

	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

// ******************************************************* MISC ********************************************************

type ResponseData struct {
	UUID               string `json:"uuid"`
	Name               string `json:"name"`
	DisplayStatus      string `json:"displayStatus"`
	Address            string `json:"address"`
	PrivateSubnetID    string `json:"privateSubnetId"`
	PrivateSubnetCidr  string `json:"privateSubnetCidr"`
	Type               string `json:"type"`
	DisplayType        string `json:"displayType"`
	LoadBalancerSchema string `json:"loadBalancerSchema"`
	PackageID          string `json:"packageId"`
	Description        string `json:"description"`
	Location           string `json:"location"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
	PackageInfo        struct {
		PackageID        string `json:"packageId"`
		ConnectionNumber int    `json:"connectionNumber"`
		DataTransfer     int    `json:"dataTransfer"`
		Name             string `json:"name"`
	} `json:"packageInfo"`
	ProgressStatus string `json:"progressStatus"`
}

// **************************************** Response of CreateLoadBalancer API *****************************************

type CreateResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreateResponse) ToLoadBalancerObject() *objects.LoadBalancer {
	if s == nil {
		return nil
	}

	return &objects.LoadBalancer{
		UUID: s.UUID,
	}
}

// ****************************************** Response of GetLoadbalancer API ******************************************

type GetResponse struct {
	Data ResponseData `json:"data"`
}

func (s *GetResponse) ToLoadBalancerObject() *objects.LoadBalancer {
	if s == nil {
		return nil
	}

	internal := lStr.ToUpper(lStr.TrimSpace(s.Data.LoadBalancerSchema)) == "INTERNAL"
	return &objects.LoadBalancer{
		UUID:               s.Data.UUID,
		Name:               s.Data.Name,
		Address:            s.Data.Address,
		DisplayStatus:      s.Data.DisplayStatus,
		PrivateSubnetID:    s.Data.PrivateSubnetID,
		PrivateSubnetCidr:  s.Data.PrivateSubnetCidr,
		Type:               s.Data.Type,
		DisplayType:        s.Data.DisplayType,
		LoadBalancerSchema: s.Data.LoadBalancerSchema,
		PackageID:          s.Data.PackageID,
		Description:        s.Data.Description,
		Location:           s.Data.Location,
		CreatedAt:          s.Data.CreatedAt,
		UpdatedAt:          s.Data.UpdatedAt,
		ProgressStatus:     s.Data.ProgressStatus,

		// will be removed
		Status:   s.Data.DisplayStatus,
		Internal: internal,
		SubnetID: s.Data.PrivateSubnetID,
	}
}

// ******************************************* Response of ListBySubnetID API ******************************************

type ListBySubnetIDResponse struct {
	Address            string `json:"address"`
	CreatedAt          string `json:"createdAt"`
	Description        string `json:"description"`
	DisplayStatus      string `json:"displayStatus"`
	DisplayType        string `json:"displayType"`
	LoadBalancerSchema string `json:"loadBalancerSchema"`
	Name               string `json:"name"`
	PackageId          string `json:"packageId"`
	PrivateSubnetCidr  string `json:"privateSubnetCidr"`
	PrivateSubnetId    string `json:"privateSubnetId"`
	ProgressStatus     string `json:"progressStatus"`
	Type               string `json:"type"`
	UUID               string `json:"uuid"`
	UpdatedAt          string `json:"updatedAt"`
}

func (s *ListBySubnetIDResponse) ToLoadBalancerObject() *objects.LoadBalancer {
	if s == nil {
		return nil
	}

	return &objects.LoadBalancer{
		UUID:               s.UUID,
		Name:               s.Name,
		Address:            s.Address,
		Description:        s.Description,
		DisplayStatus:      s.DisplayStatus,
		DisplayType:        s.DisplayType,
		LoadBalancerSchema: s.LoadBalancerSchema,
		PackageID:          s.PackageId,
		PrivateSubnetCidr:  s.PrivateSubnetCidr,
		PrivateSubnetID:    s.PrivateSubnetId,
		ProgressStatus:     s.ProgressStatus,
		Type:               s.Type,
		CreatedAt:          s.CreatedAt,
		UpdatedAt:          s.UpdatedAt,

		Status:   s.DisplayStatus,
		Internal: lStr.ToUpper(lStr.TrimSpace(s.LoadBalancerSchema)) == "INTERNAL",
		SubnetID: s.PrivateSubnetId,
	}
}

// *********************************************** Response of List API ************************************************

type ListResponse struct {
	ListData  []ResponseData `json:"listData"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

func (s *ListResponse) ToListLoadBalancerObjects() []*objects.LoadBalancer {
	if s == nil || s.ListData == nil || len(s.ListData) < 1 {
		return nil
	}

	var result []*objects.LoadBalancer
	for _, itemLb := range s.ListData {
		result = append(result, &objects.LoadBalancer{
			UUID:               itemLb.UUID,
			Name:               itemLb.Name,
			Address:            itemLb.Address,
			Description:        itemLb.Description,
			DisplayStatus:      itemLb.DisplayStatus,
			DisplayType:        itemLb.DisplayType,
			LoadBalancerSchema: itemLb.LoadBalancerSchema,
			PackageID:          itemLb.PackageID,
			PrivateSubnetCidr:  itemLb.PrivateSubnetCidr,
			PrivateSubnetID:    itemLb.PrivateSubnetID,
			ProgressStatus:     itemLb.ProgressStatus,
			Type:               itemLb.Type,
			CreatedAt:          itemLb.CreatedAt,
			UpdatedAt:          itemLb.UpdatedAt,

			Status:   itemLb.DisplayStatus,
			Internal: lStr.ToUpper(lStr.TrimSpace(itemLb.LoadBalancerSchema)) == "INTERNAL",
			SubnetID: itemLb.PrivateSubnetID,
		})
	}

	return result
}

// ********************************************** Response of Update API ***********************************************

type UpdateResponse struct {
	UUID string `json:"uuid"`
}

func (s *UpdateResponse) ToLoadBalancerObject() *objects.LoadBalancer {
	if s == nil {
		return nil
	}

	return &objects.LoadBalancer{
		UUID: s.UUID,
	}
}
