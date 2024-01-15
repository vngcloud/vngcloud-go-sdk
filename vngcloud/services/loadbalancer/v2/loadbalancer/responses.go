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

// ******************************************* Response of CreateVolume API ********************************************

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

// ********************************************* Response of GetVolume API *********************************************

type GetResponse struct {
	Data ResponseData `json:"data"`
}

func (s *GetResponse) ToLoadBalancerObject() *objects.LoadBalancer {
	if s == nil {
		return nil
	}

	internal := lStr.ToUpper(lStr.TrimSpace(s.Data.LoadBalancerSchema)) == "INTERNAL"
	return &objects.LoadBalancer{
		UUID:     s.Data.UUID,
		Status:   s.Data.DisplayStatus,
		Address:  s.Data.Address,
		Name:     s.Data.Name,
		SubnetID: s.Data.PrivateSubnetID,
		Internal: internal,
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
		UUID:     s.UUID,
		Status:   s.DisplayStatus,
		Address:  s.Address,
		Name:     s.Name,
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
			UUID:     itemLb.UUID,
			Status:   itemLb.DisplayStatus,
			Address:  itemLb.Address,
			Name:     itemLb.Name,
			SubnetID: itemLb.PrivateSubnetID,
		})
	}

	return result
}
