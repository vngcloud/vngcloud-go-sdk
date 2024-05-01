package snapshot

import (
	lsc "github.com/vngcloud/vngcloud-go-sdk/client"
	lse "github.com/vngcloud/vngcloud-go-sdk/error"
	lserrHandler "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	lso "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

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
