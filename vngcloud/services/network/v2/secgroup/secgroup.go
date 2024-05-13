package secgroup

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/client"
	lserrors "github.com/vngcloud/vngcloud-go-sdk/error"
	lserror "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	lsobj "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Create(psc *lsclient.ServiceClient, popts ICreateOptsBuilder) (*lsobj.Secgroup, *lserrors.SdkError) {
	response := NewCreateResponse()
	body := popts.ToRequestBody()
	errResp := lserrors.NewErrorResponse()

	if _, err := psc.Post(createURL(psc, popts), &lsclient.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		JSONError:    errResp,
		OkCodes:      []int{201},
	}); err != nil {
		return nil, lserror.ErrorHandler(err,
			lserror.WithErrorSecgroupNameAlreadyExists(errResp, err),
			lserror.WithErrorSecgroupExceedQuota(errResp, err))
	}

	return response.ToSecgroupObject(), nil
}

func Delete(pSc *lsclient.ServiceClient, pOpts IDeleteOptsBuilder) *lserrors.SdkError {
	errResp := lserrors.NewErrorResponse()

	if _, err := pSc.Delete(deleteURL(pSc, pOpts), &lsclient.RequestOpts{
		OkCodes:   []int{204},
		JSONError: errResp,
	}); err != nil {
		return lserror.ErrorHandler(err,
			lserror.WithErrorSecgroupInUse(errResp, err),
			lserror.WithErrorSecgroupNotFound(errResp, err))
	}

	return nil
}

func Get(pSc *lsclient.ServiceClient, pOpts IGetOptsBuilder) (*lsobj.Secgroup, *lserrors.SdkError) {
	response := NewGetResponse()
	errResp := lserrors.NewErrorResponse()

	if _, err := pSc.Get(getURL(pSc, pOpts), &lsclient.RequestOpts{
		JSONResponse: response,
		JSONError:    errResp,
		OkCodes:      []int{200},
	}); err != nil {
		return nil, lserror.ErrorHandler(err,
			lserror.WithErrorSecgroupNotFound(errResp, err))
	}

	return response.ToSecgroupObject(), nil
}

func List(pSc *lsclient.ServiceClient, pOpts IListOptsBuilder) ([]*lsobj.Secgroup, error) {
	response := NewListResponse()
	if _, err := pSc.Get(listURL(pSc, pOpts), &lsclient.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	}); err != nil {
		return nil, err
	}

	return response.ToListSecgroupObjects(), nil
}
