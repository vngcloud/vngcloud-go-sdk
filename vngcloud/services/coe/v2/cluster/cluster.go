package cluster

import (
	"encoding/json"
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
	"strings"
)

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.Cluster, error) {
	response := NewGetResponse()
	reqRes, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: response,
	})

	if err != nil {
		result := make(map[string]interface{})
		err2 := json.Unmarshal(reqRes.Bytes(), &result)
		if err2 == nil {
			if message, _ := result["message"].(string); strings.Contains(strings.ToLower(strings.TrimSpace(message)), "is not found") {
				return nil, NewClusterNotFound(pOpts.GetClusterID(), "")
			}
		}

		return nil, err
	}

	return response.ToClusterObject(), nil
}

func UpdateSecgroup(pSc *client.ServiceClient, pOpts IUpdateSecgroupOptsBuilder) ([]*objects.ClusterSecgroup, error) {
	response := NewUpdateSecgroupResponse()
	_, err := pSc.Put(updateSecgroupURL(pSc, pOpts), &client.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: response,
		JSONBody:     pOpts.ToRequestBody(),
	})

	if err != nil {
		return nil, err
	}

	return response.ToListClusterSecgroupObjects(), nil
}
