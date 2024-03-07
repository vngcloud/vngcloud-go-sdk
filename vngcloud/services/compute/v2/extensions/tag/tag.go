package tag

import (
	"encoding/json"
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Create(sc *client.ServiceClient, opts ICreateOptsBuilder) ([]*objects.ServerTag, error) {
	reqRes, err := sc.Put(createURL(sc, opts), &client.RequestOpts{
		OkCodes:  []int{200},
		JSONBody: opts.ToRequestBody(),
	})

	if err != nil {
		return nil, err
	}

	response := NewCreateResponse()
	if reqRes.StatusCode == 200 {
		err = json.Unmarshal(reqRes.Bytes(), &response)
		if err != nil {
			return nil, err
		}
	}

	var serverTags []*objects.ServerTag
	if response == nil || len(response) < 1 {
		return serverTags, nil
	}

	for _, item := range response {
		if item != nil {
			serverTags = append(serverTags, item.ToServerTagObject())
		}
	}

	return serverTags, nil
}

func Get(sc *client.ServiceClient, opts IGetOptsBuilder) ([]*objects.ServerTag, error) {
	reqRes, err := sc.Put(getURL(sc, opts), &client.RequestOpts{
		OkCodes: []int{200},
	})

	if err != nil {
		return nil, err
	}

	response := NewCreateResponse()
	if reqRes.StatusCode == 200 {
		err = json.Unmarshal(reqRes.Bytes(), &response)
		if err != nil {
			return nil, err
		}
	}

	var serverTags []*objects.ServerTag
	if response == nil || len(response) < 1 {
		return serverTags, nil
	}

	for _, item := range response {
		if item != nil {
			serverTags = append(serverTags, item.ToServerTagObject())
		}
	}

	return serverTags, nil
}
