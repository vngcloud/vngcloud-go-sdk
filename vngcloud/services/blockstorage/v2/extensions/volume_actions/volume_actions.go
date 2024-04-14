package volume_actions

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/client"
	lssdkErr "github.com/vngcloud/vngcloud-go-sdk/error"
	lserr "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	lsobj "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Resize(psc *lsclient.ServiceClient, popts IResizeOptsBuilder) (*lsobj.ResizeVolume, *lssdkErr.SdkError) {
	response := NewResizeResponse()
	errResp := lssdkErr.NewErrorResponse()
	_, err := psc.Put(resizeURL(psc, popts), &lsclient.RequestOpts{
		JSONResponse: response,
		JSONBody:     popts.ToRequestBody(),
		JSONError:    errResp,
		OkCodes:      []int{202},
	})

	if err != nil {
		sdkErr := lserr.ErrorHandler(err,
			lserr.WithErrorVolumeUnchanged(errResp, err))

		return nil, sdkErr
	}

	return response.ToResizeVolumeObject(), nil
}

func GetMappingVolume(psc *lsclient.ServiceClient, popts IMappingOptsBuilder) (*lsobj.MappingVolume, error) {
	response := NewBackendVolumeResponse()
	_, err := psc.Get(mappingURL(psc, popts), &lsclient.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToMappingVolumeObject(), nil
}
