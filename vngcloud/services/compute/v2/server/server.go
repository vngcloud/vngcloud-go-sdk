package server

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Get(sc *client.ServiceClient, opts IGetOptsBuilder) (*objects.Server, error) {
	response := NewGetResponse()
	_, err := sc.Get(getServerURL(sc, opts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToServerObject(), nil
}
