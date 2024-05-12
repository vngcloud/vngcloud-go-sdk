package pool

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	lserror "github.com/vngcloud/vngcloud-go-sdk/error"
	lserrors "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
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

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) *lserror.SdkError {
	errResp := lserror.NewErrorResponse()
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes:   []int{202},
		JSONError: errResp,
	})

	if err != nil {
		return lserrors.ErrorHandler(err,
			lserrors.WithErrorPoolNotFound(errResp, err),
			lserrors.WithErrorPoolInUse(errResp, err))
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

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.Pool, error) {
	response := NewGetResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToPoolObject(), nil
}

func GetMember(pSc *client.ServiceClient, pOpts IGetMemberOptsBuilder) ([]*objects.Member, error) {
	response := NewGetMemberResponse()
	_, err := pSc.Get(getMemberURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListMemberObject(), nil
}

func Update(pSc *client.ServiceClient, pOpts IUpdateOptsBuilder) error {
	_, err := pSc.Put(updateURL(pSc, pOpts), &client.RequestOpts{
		JSONBody: pOpts.ToRequestBody(),
		OkCodes:  []int{202},
	})

	if err != nil {
		return err
	}

	return nil
}

func GetHealthMonitor(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.PoolHealthMonitor, error) {
	response := NewGetHealthMonitorResponse()
	_, err := pSc.Get(getHealthMonitorURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToPoolHealthMonitorObject(), nil
}

func GetTotal(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.Pool, error) {
	p, err := Get(pSc, pOpts)
	if err != nil {
		return nil, err
	}

	h, err := GetHealthMonitor(pSc, pOpts)
	if err != nil {
		return nil, err
	}

	p.HealthMonitor = h
	return p, nil
}
