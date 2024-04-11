package volume_attach

import (
	"encoding/json"
	"fmt"
	"github.com/vngcloud/vngcloud-go-sdk/client"
	lsdkError "github.com/vngcloud/vngcloud-go-sdk/error"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
	"strings"
)

// Create creates a volume attachment.
func Create(sc *client.ServiceClient, opts ICreateOptsBuilder) (*objects.VolumeAttach, error) {
	response := NewCreateResponse()
	reqRes, err := sc.Put(createURL(sc, opts), &client.RequestOpts{
		OkCodes:      []int{202},
		JSONResponse: response,
		JSONBody:     opts.ToRequestBody(),
	})

	if err != nil {
		result := make(map[string]interface{})
		err2 := json.Unmarshal(reqRes.Bytes(), &result)
		if err2 == nil {
			if message, _ := result["message"].(string); strings.TrimSpace(message) == "This volume is available" {
				return nil, NewErrAttachNotFound(fmt.Sprintf("volume %s is available", opts.GetVolumeID()))
			}
		}

		return nil, err
	}

	return response.ToVolumeAttachObject(), nil
}

func Attach(sc *client.ServiceClient, opts ICreateOptsBuilder) (*objects.VolumeAttach, error) {
	response := NewCreateResponse()
	reqRes, err := sc.Put(createURL(sc, opts), &client.RequestOpts{
		OkCodes:      []int{202},
		JSONResponse: response,
		JSONBody:     map[string]interface{}{},
	})

	if err != nil {
		fmt.Println(err)
		result := make(map[string]interface{})
		err2 := json.Unmarshal(reqRes.Bytes(), &result)
		if err2 == nil {
			if message, _ := result["message"].(string); strings.TrimSpace(message) == "This volume is available" {
				return nil, NewErrAttachNotFound(fmt.Sprintf("volume %s is available", opts.GetVolumeID()))
			}
		}

		return nil, err
	}

	return response.ToVolumeAttachObject(), nil
}

// Delete deletes a volume attachment.
func Delete(sc *client.ServiceClient, opts IDeleteOptsBuilder) (*objects.VolumeAttach, error) {
	response := NewDeleteResponse()
	reqRes, err := sc.Put(deleteURL(sc, opts), &client.RequestOpts{
		OkCodes:      []int{202},
		JSONResponse: response,
		JSONBody:     opts.ToRequestBody(),
	})

	if err != nil {
		result := make(map[string]interface{})
		err2 := json.Unmarshal(reqRes.Bytes(), &result)
		if err2 == nil {
			if message, _ := result["message"].(string); strings.TrimSpace(message) == "This volume is available" {
				return nil, NewErrAttachNotFound(fmt.Sprintf("volume %s is available", opts.GetVolumeID()))
			}
		}

		return nil, err
	}

	return response.ToVolumeAttachObject(), nil
}

// Delete deletes a volume attachment.
func Detach(sc *client.ServiceClient, opts IDeleteOptsBuilder) (*objects.VolumeAttach, error) {
	response := NewDeleteResponse()
	errResp := lsdkError.NewErrorResponse()
	_, err := sc.Put(deleteURL(sc, opts), &client.RequestOpts{
		OkCodes:      []int{202},
		JSONResponse: response,
		JSONBody:     map[string]interface{}{},
		JSONError:    errResp,
	})

	if err != nil {
		if strings.Contains(errResp.Message, "This volume is available") {
			return nil, nil
		}

		if strings.Contains(errResp.Message, "is not found") {
			return nil, nil
		}

		return nil, err
	}

	return response.ToVolumeAttachObject(), nil
}
