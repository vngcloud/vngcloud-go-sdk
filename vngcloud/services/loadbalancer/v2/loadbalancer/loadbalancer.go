package loadbalancer

import (
	ljson "encoding/json"

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

func Delete(pSc *lsclient.ServiceClient, pOpts IDeleteOptsBuilder) *lserr.SdkError {
	errResp := lserr.NewErrorResponse()
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &lsclient.RequestOpts{
		OkCodes:   []int{202},
		JSONError: errResp,
	})

	if err != nil {
		return lserrHandler.ErrorHandler(err,
			lserrHandler.WithErrorLoadBalancerIsDeleting(errResp, err))
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
		err = ljson.Unmarshal(resp.Bytes(), &response)
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

func List(psc *lsclient.ServiceClient, popts IListOptsBuilder) ([]*lsobj.LoadBalancer, *lserr.SdkError) {
	response := NewListResponse()
	errResp := lserr.NewErrorResponse()
	url := listURL(psc, popts)
	_, err := psc.Get(url, &lsclient.RequestOpts{
		JSONResponse: response,
		JSONError:    errResp,
		OkCodes:      []int{200},
	})

	if err != nil {
		sdkErr := lserrHandler.ErrorHandler(err)
		return nil, sdkErr
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

func ListTags(psc *lsclient.ServiceClient, popts IListTagsOptsBuilder) ([]*lsobj.LoadBalancerTag, *lserr.SdkError) {
	response := NewListTagsResponse()
	errResp := lserr.NewErrorResponse()
	resp, err := psc.Get(listTagsURL(psc, popts), &lsclient.RequestOpts{
		JSONError: errResp,
		OkCodes:   []int{200},
	})

	if err != nil {
		sdkErr := lserrHandler.ErrorHandler(err)
		return nil, sdkErr
	}

	if err := ljson.Unmarshal(resp.Bytes(), &response); err != nil {
		return nil, lserrHandler.ErrorHandler(err)
	}

	var lstTags []*lsobj.LoadBalancerTag
	for _, item := range response {
		if item != nil {
			lstTags = append(lstTags, item.ToLoadBalancerTagObject())
		}
	}

	return lstTags, nil
}

func UpdateTag(psc *lsclient.ServiceClient, popts IUpdateTagOptsBuilder) *lserr.SdkError {
	tags, sdkErr := ListTags(psc, NewListTagsOpts(popts.GetProjectID(), popts.GetLoadBalancerID()))
	if sdkErr != nil {
		return sdkErr
	}

	body := popts.ToRequestBody(tags)
	errResp := lserr.NewErrorResponse()
	_, err := psc.Put(updateTagURL(psc, popts), &lsclient.RequestOpts{
		JSONBody:  body,
		JSONError: errResp,
		OkCodes:   []int{200},
	})

	if err != nil {
		return lserrHandler.ErrorHandler(err)
	}

	return nil
}
