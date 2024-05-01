package snapshot

import (
	lsc "github.com/vngcloud/vngcloud-go-sdk/client"
	lse "github.com/vngcloud/vngcloud-go-sdk/error"
	lserrHandler "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	lso "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Create(psc *lsc.ServiceClient, ptops ICreateOptsBuilder) (*lso.Snapshot, error) {
	response := NewCreateResponse()
	body := ptops.ToRequestBody()
	_, err := psc.Post(createURL(psc, ptops), &lsc.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToSnapshotObject(), nil
}

func Delete(psc *lsc.ServiceClient, ptops IDeleteOptsBuilder) *lse.SdkError {
	errResp := lse.NewErrorResponse()
	_, err := psc.Delete(deleteURL(psc, ptops), &lsc.RequestOpts{
		JSONError: errResp,
		OkCodes:   []int{200},
	})

	if err != nil {
		return lserrHandler.ErrorHandler(err,
			lserrHandler.WithErrorSnapshotNotFound(errResp, err))
	}

	return nil
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
