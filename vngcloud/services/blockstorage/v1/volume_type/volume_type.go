package volume_type

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Get(pSc *client.ServiceClient, pOpts IGetVolumeTypeOptsBuilder) (*objects.VolumeType, error) {
	response := NewGetVolumeTypeResponses()
	_, err := pSc.Get(
		getVolumeTypeURL(pSc, pOpts),
		&client.RequestOpts{
			JSONResponse: response,
			OkCodes:      []int{200}})

	if err != nil {
		return nil, err
	}

	return response.ToVolumeTypeObject()
}
