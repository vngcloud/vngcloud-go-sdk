package volume

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	lssdkError "github.com/vngcloud/vngcloud-go-sdk/error"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	lsobj "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func List(psc *client.ServiceClient, popts IListOptsBuilder) (*lsobj.VolumeList, *lssdkError.SdkError) {
	resp := NewListResponse()
	errResp := lssdkError.NewErrorResponse()
	url := listURL(psc, popts)

	_, err := psc.Get(url, &client.RequestOpts{
		JSONResponse: resp,
		JSONError:    errResp,
		OkCodes:      []int{200},
	})

	if err != nil {
		sdkErr := errors.ErrorHandler(err)
		return nil, sdkErr

	}

	return resp.ToVolumeListObject(), nil
}

func ListAll(pSc *client.ServiceClient, pOpts IListAllOptsBuilder) ([]*lsobj.Volume, error) {
	resp := NewListAllResponse()
	url := listAllURL(pSc, pOpts)
	_, err := pSc.Get(url, &client.RequestOpts{
		JSONResponse: resp,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return resp.ToListVolumeObjects(), nil
}

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*lsobj.Volume, error) {
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

	return response.ToVolumeObject(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) error {
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes: []int{202},
	})

	return err
}

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*lsobj.Volume, *lssdkError.SdkError) {
	errResp := lssdkError.NewErrorResponse()
	response := NewGetResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
		JSONError:    errResp,
	})

	if err != nil {
		sdkErr := errors.ErrorHandler(err,
			errors.WithErrorVolumeNotFound(errResp, err))

		return nil, sdkErr
	}

	return response.ToVolumeObject(), nil
}
