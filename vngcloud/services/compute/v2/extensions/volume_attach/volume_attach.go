package volume_attach

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vngcloud/vngcloud-go-sdk/client"
	lsdkError "github.com/vngcloud/vngcloud-go-sdk/error"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
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
func Detach(sc *client.ServiceClient, opts IDeleteOptsBuilder) (*objects.VolumeAttach, *lsdkError.SdkError) {
	response := NewDeleteResponse()
	errResp := lsdkError.NewErrorResponse()
	_, err := sc.Put(deleteURL(sc, opts), &client.RequestOpts{
		OkCodes:      []int{202},
		JSONResponse: response,
		JSONBody:     map[string]interface{}{},
		JSONError:    errResp,
	})

	if err != nil {
		sdkErr := errors.ErrorHandler(err,
			errors.WithErrorVolumeAvailable(errResp, err),
			errors.WithErrorNotFound(errResp, err))

		return nil, sdkErr
	}

	return response.ToVolumeAttachObject(), nil
}
