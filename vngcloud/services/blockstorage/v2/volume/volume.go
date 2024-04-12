package volume

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	lsdkError "github.com/vngcloud/vngcloud-go-sdk/error"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/pagination"
)

func List(pSc *client.ServiceClient, pOpts IListOptsBuilder) *pagination.Pager {
	qp, err := pOpts.ToListQuery()
	url := listURL(pSc, pOpts)
	if err == nil {
		url = url + qp
	}
	return pagination.NewPager(pSc, url, pOpts,
		func() interface{} {
			return NewListResponse()
		},
		func(r interface{}) pagination.IPage {
			resp := r.(*ListResponse)
			return resp
		})
}

func ListAll(pSc *client.ServiceClient, pOpts IListAllOptsBuilder) ([]*objects.Volume, error) {
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

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*objects.Volume, error) {
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

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.Volume, *lsdkError.SdkError) {
	errResp := lsdkError.NewErrorResponse()
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
