package pool

import (
	"encoding/json"
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
	"strings"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*objects.Pool, error) {
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

	return response.ToPoolObject(), nil
}

func ListPoolsBasedLoadBalancer(pSc *client.ServiceClient, pOpts IListPoolsBasedLoadBalancerOptsBuilder) ([]*objects.Pool, error) {
	response := NewListPoolsBasedLoadBalancerResponse()
	_, err := pSc.Get(listPoolsBasedLoadBalancerURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListPoolObjects(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) error {
	reqRes, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes: []int{202},
	})

	if err != nil {
		result := make(map[string]interface{})
		err2 := json.Unmarshal(reqRes.Bytes(), &result)
		if err2 == nil {
			if message, _ := result["message"].(string); strings.Contains(strings.ToLower(strings.TrimSpace(message)), "is used in listener") {
				return NewErrPoolInUse(pOpts.GetPoolID(), "")
			}
		}

		return err
	}

	return nil
}

func UpdatePoolMembers(pSc *client.ServiceClient, pOpts IUpdatePoolMembersOptsBuilder) error {
	errorResolver := new(ErrorResolver)
	_, err := pSc.Put(updatePoolMembersURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:  pOpts.ToRequestBody(),
		JSONError: errorResolver,
		OkCodes:   []int{202},
	})

	if err != nil {
		return errorResolver.ToError()
	}

	return nil
}
