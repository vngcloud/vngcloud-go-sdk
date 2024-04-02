package secgroup

import (
	"encoding/json"
	"github.com/vngcloud/vngcloud-go-sdk/client"
	lsdkError "github.com/vngcloud/vngcloud-go-sdk/error"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
	"strings"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*objects.Secgroup, error) {
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	reqRes, err := pSc.Post(createURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{201},
	})

	if err != nil {
		result := make(map[string]interface{})
		err2 := json.Unmarshal(reqRes.Bytes(), &result)
		if err2 == nil {
			if message, _ := result["message"].(string); strings.Contains(strings.ToLower(strings.TrimSpace(message)), "name of security group already exist") {
				return nil, NewErrNameDuplicate("", "name is already used")
			}
		}

		return nil, err
	}

	return response.ToSecgroupObject(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) *lsdkError.SdkError {
	errResp := lsdkError.NewErrorResponse()
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes:   []int{204},
		JSONError: errResp,
	})

	if err != nil {
		if strings.Contains(errResp.Message, patternErrNotFound) {
			return &lsdkError.SdkError{
				Code:    ErrSecgroupInUse,
				Message: errResp.Message,
				Error:   err,
			}
		}

		return &lsdkError.SdkError{
			Code:    ErrSecgroupUnknown,
			Message: errResp.Message,
			Error:   err,
		}
	}

	return nil
}

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.Secgroup, *lsdkError.SdkError) {
	response := NewGetResponse()
	errResp := lsdkError.NewErrorResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		JSONError:    errResp,
		OkCodes:      []int{200},
	})

	if err != nil {
		if strings.Contains(errResp.Message, patternErrNotFound) {
			return nil, &lsdkError.SdkError{
				Code:    ErrSecgroupNotFound,
				Message: errResp.Message,
				Error:   err,
			}
		}

		return nil, &lsdkError.SdkError{
			Code:    ErrSecgroupUnknown,
			Message: errResp.Message,
			Error:   err,
		}
	}

	return response.ToSecgroupObject(), nil
}

func List(pSc *client.ServiceClient, pOpts IListOptsBuilder) ([]*objects.Secgroup, error) {
	response := NewListResponse()
	_, err := pSc.Get(listURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListSecgroupObjects(), nil
}
