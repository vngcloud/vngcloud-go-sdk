package subnet

import (
	"encoding/json"
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.Subnet, error) {
	response := NewGetResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: response,
	})

	if err != nil {
		return nil, err
	}

	return response.ToSubnetObject(), nil
}

func ListByNetworkID(pSc *client.ServiceClient, pOpts IListByNetworkIDOptsBuilder) ([]*objects.Subnet, error) {
	response := NewListByNetworkIDResponse()
	resp, err := pSc.Get(listByNetworkIDURL(pSc, pOpts), &client.RequestOpts{
		OkCodes: []int{200},
	})

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		err = json.Unmarshal(resp.Bytes(), &response)
		if err != nil {
			return nil, err
		}
	}

	var lstSubnetObjects []*objects.Subnet
	if response == nil || len(response) < 1 {
		return lstSubnetObjects, nil
	}

	for _, itemLb := range response {
		if itemLb != nil {
			lstSubnetObjects = append(lstSubnetObjects, itemLb.ToSubnetObject())
		}
	}

	return lstSubnetObjects, nil
}
