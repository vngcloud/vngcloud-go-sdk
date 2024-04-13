package snapshot

import (
	lsc "github.com/vngcloud/vngcloud-go-sdk/client"
	lse "github.com/vngcloud/vngcloud-go-sdk/error"
	lserrHandler "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	lso "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Create(pSc *lsc.ServiceClient, pOpts ICreateOptsBuilder) (*lso.Snapshot, error) {
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	_, err := pSc.Post(createURL(pSc, pOpts), &lsc.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToSnapshotObject(), nil
}

func Delete(pSc *lsc.ServiceClient, pOpts IDeleteOptsBuilder) error {
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &lsc.RequestOpts{
		OkCodes: []int{200},
	})

	return err
}

func ListVolumeSnapshot(psc *lsc.ServiceClient, popts IListVolumeOptsBuilder) (*lso.SnapshotList, *lse.SdkError) {
	resp := NewListVolumeResponse()
	errResp := lse.NewErrorResponse()
	url := listVolumeSnapshotURL(psc, popts)

	_, err := psc.Get(url, &lsc.RequestOpts{
		JSONResponse: resp,
		JSONError:    errResp,
		OkCodes:      []int{200},
	})

	if err != nil {
		sdkErr := lserrHandler.ErrorHandler(err)
		return nil, sdkErr
	}

	return resp.ToSnapshotListObject(), nil
}
