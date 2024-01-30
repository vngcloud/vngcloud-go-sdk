package volume_actions

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Resize(pSc *client.ServiceClient, pOpts IResizeOptsBuilder) (*objects.ResizeVolume, error) {
	response := NewResizeResponse()
	_, err := pSc.Put(resizeURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		JSONBody:     pOpts.ToRequestBody(),
		OkCodes:      []int{202},
	})

	if err != nil {
		return nil, err
	}

	return response.ToResizeVolumeObject(), nil
}

func GetMappingVolume(pSc *client.ServiceClient, pOpts IMappingOptsBuilder) (*objects.MappingVolume, error) {
	response := NewBackendVolumeResponse()
	_, err := pSc.Get(mappingURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToMappingVolumeObject(), nil
}
