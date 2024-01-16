package policy

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func List(pSc *client.ServiceClient, pOpts IListOptsBuilder) ([]*objects.Policy, error) {
	response := NewListResponse()
	_, err := pSc.Get(listURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})


	if err != nil {
		return nil, err
	}

	return response.ToListPolicyObjects(), nil
}
