package loadbalancer

import (
	"encoding/json"
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*objects.LoadBalancer, error) {
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	_, err := pSc.Post(createURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{202},
	})

	if err != nil {
		return nil, err
	}

	return response.ToLoadBalancerObject(), nil
}

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.LoadBalancer, error) {
	response := NewGetResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToLoadBalancerObject(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) error {
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes: []int{202},
	})

	if err != nil {
		return err
	}

	return nil
}

func ListBySubnetID(pSc *client.ServiceClient, pOpts IListBySubnetIDOptsBuilder) ([]*objects.LoadBalancer, error) {
	response := NewListBySubnetIDResponse()
	resp, err := pSc.Get(listBySubnetIDURL(pSc, pOpts), &client.RequestOpts{
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

	var lstLbObjects []*objects.LoadBalancer
	if response == nil || len(response) < 1 {
		return lstLbObjects, nil
	}

	for _, itemLb := range response {
		if itemLb != nil {
			lstLbObjects = append(lstLbObjects, itemLb.ToLoadBalancerObject())
		}
	}

	return lstLbObjects, nil
}

func List(pSc *client.ServiceClient, pOpts IListOptsBuilder) ([]*objects.LoadBalancer, error) {
	response := NewListResponse()
	_, err := pSc.Get(listURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListLoadBalancerObjects(), nil
}
