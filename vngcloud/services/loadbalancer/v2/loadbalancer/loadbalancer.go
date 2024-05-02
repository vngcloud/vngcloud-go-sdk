package loadbalancer

import (
	"encoding/json"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/client"
	lserr "github.com/vngcloud/vngcloud-go-sdk/error"
	lserrHandler "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	lsobj "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Create(pSc *lsclient.ServiceClient, pOpts ICreateOptsBuilder) (*lsobj.LoadBalancer, error) {
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	_, err := pSc.Post(createURL(pSc, pOpts), &lsclient.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{202},
	})

	if err != nil {
		return nil, err
	}

	return response.ToLoadBalancerObject(), nil
}

func Get(pSc *lsclient.ServiceClient, pOpts IGetOptsBuilder) (*lsobj.LoadBalancer, *lserr.SdkError) {
	response := NewGetResponse()
	errResp := lserr.NewErrorResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &lsclient.RequestOpts{
		JSONResponse: response,
		JSONError:    errResp,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, lserrHandler.ErrorHandler(err,
			lserrHandler.WithErrorLoadBalancerNotFound(errResp, err))
	}

	return response.ToLoadBalancerObject(), nil
}

func Delete(pSc *lsclient.ServiceClient, pOpts IDeleteOptsBuilder) error {
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &lsclient.RequestOpts{
		OkCodes: []int{202},
	})

	if err != nil {
		return err
	}

	return nil
}

func ListBySubnetID(pSc *lsclient.ServiceClient, pOpts IListBySubnetIDOptsBuilder) ([]*lsobj.LoadBalancer, error) {
	response := NewListBySubnetIDResponse()
	resp, err := pSc.Get(listBySubnetIDURL(pSc, pOpts), &lsclient.RequestOpts{
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

	var lstLbObjects []*lsobj.LoadBalancer
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

func List(pSc *lsclient.ServiceClient, pOpts IListOptsBuilder) ([]*lsobj.LoadBalancer, error) {
	response := NewListResponse()
	_, err := pSc.Get(listURL(pSc, pOpts), &lsclient.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListLoadBalancerObjects(), nil
}

func Update(pSc *lsclient.ServiceClient, pOpts IUpdateOptsBuilder) (*lsobj.LoadBalancer, error) {
	response := NewUpdateResponse()
	body := pOpts.ToRequestBody()
	_, err := pSc.Put(updateURL(pSc, pOpts), &lsclient.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{202},
	})

	if err != nil {
		return nil, err
	}

	return response.ToLoadBalancerObject(), nil
}

func CreateTag(psc *lsclient.ServiceClient, popts ICreateTagOptsBuilder) *lserr.SdkError {
	body := popts.ToRequestBody()
	errResp := lserr.NewErrorResponse()
	_, err := psc.Put(createTagUrl(psc, popts), &lsclient.RequestOpts{
		JSONBody:  body,
		JSONError: errResp,
		OkCodes:   []int{200},
	})

	if err != nil {
		return lserrHandler.ErrorHandler(err)
	}

	return nil
}
