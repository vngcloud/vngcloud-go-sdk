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

func Delete(sc *client.ServiceClient, opts IDeleteOptsBuilder) error {
	_, err := sc.Delete(deleteServerURL(sc, opts), &client.RequestOpts{
		JSONBody: opts.ToRequestBody(),
		OkCodes:  []int{202},
	})

	return err
}

func Create(sc *client.ServiceClient, opts ICreateOptsBuilder) (*objects.Server, error) {
	response := NewCreateResponse()
	_, err := sc.Post(createServerURL(sc, opts), &client.RequestOpts{
		JSONBody:     opts.ToRequestBody(),
		JSONResponse: response,
		OkCodes:      []int{202},
	})

	if err != nil {
		return nil, err
	}

	return response.ToServerObject(), nil
}
