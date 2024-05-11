package listener

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	lsdkError "github.com/vngcloud/vngcloud-go-sdk/error"
	lserrHandler "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*objects.Listener, error) {
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

	return response.ToListenerObject(), nil
}

func GetBasedLoadBalancer(pSc *client.ServiceClient, pOpts IGetBasedLoadBalancerOptsBuilder) ([]*objects.Listener, error) {
	response := NewGetBasedLoadBalancerResponse()
	_, err := pSc.Get(getBasedLoadBalancerURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListListenerObject(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) *lsdkError.SdkError {
	errResp := lsdkError.NewErrorResponse()
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes:   []int{202},
		JSONError: errResp,
	})

	if err != nil {
		return lserrHandler.ErrorHandler(err,
			lserrHandler.WithErrorListenerNotFound(errResp, err))
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
