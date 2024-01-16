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

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*objects.Policy, error) {
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

	return response.ToPolicyObject(), nil
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

func Update(pSc *client.ServiceClient, pOpts IUpdateOptsBuilder) error {
	body := pOpts.ToRequestBody()
	_, err := pSc.Put(updateURL(pSc, pOpts), &client.RequestOpts{
		JSONBody: body,
		OkCodes:  []int{202},
	})

	if err != nil {
		return err
	}

	return nil
}
