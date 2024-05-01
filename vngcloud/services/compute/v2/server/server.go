package server

import (
	"fmt"
	"github.com/vngcloud/vngcloud-go-sdk/client"
	lsdkError "github.com/vngcloud/vngcloud-go-sdk/error"
	lserr "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
	"strings"
)

func Get(sc *client.ServiceClient, opts IGetOptsBuilder) (*objects.Server, *lsdkError.SdkError) {
	response := NewGetResponse()
	errResp := lsdkError.NewErrorResponse()
	_, err := sc.Get(getServerURL(sc, opts), &client.RequestOpts{
		JSONResponse: response,
		JSONError:    errResp,
		OkCodes:      []int{200},
	})

	if err != nil {
		if strings.Contains(errResp.Message, patternErrNotFound) {
			return nil, &lsdkError.SdkError{
				Code:    ErrNotFound,
				Message: errResp.Message,
				Error:   err,
			}
		} else {
			return nil, &lsdkError.SdkError{
				Code:    ErrUnknown,
				Message: errResp.Message,
				Error:   err,
			}
		}
	}

	return response.ToServerObject(), nil
}

func Delete(sc *client.ServiceClient, opts IDeleteOptsBuilder) *lsdkError.SdkError {
	errResp := lsdkError.NewErrorResponse()
	_, err := sc.Delete(deleteServerURL(sc, opts), &client.RequestOpts{
		JSONBody:  opts.ToRequestBody(),
		JSONError: errResp,
		OkCodes:   []int{202},
	})

	if err != nil {
		if strings.Contains(errResp.Message, patternErrNotFound) {
			return &lsdkError.SdkError{
				Code:    ErrNotFound,
				Message: errResp.Message,
				Error:   err,
			}
		} else if strings.Contains(errResp.Message, patternDeleteCreatingServer) {
			return &lsdkError.SdkError{
				Code:    ErrDeleteCreatingServer,
				Message: errResp.Message,
				Error:   err,
			}
		} else {
			return &lsdkError.SdkError{
				Code:    ErrUnknown,
				Message: errResp.Message,
				Error:   err,
			}
		}
	}

	return nil
}

func Create(sc *client.ServiceClient, opts ICreateOptsBuilder) (*objects.Server, *lsdkError.SdkError) {
	response := NewCreateResponse()
	errResp := lsdkError.NewErrorResponse()
	_, err := sc.Post(createServerURL(sc, opts), &client.RequestOpts{
		JSONBody:     opts.ToRequestBody(),
		JSONResponse: response,
		JSONError:    errResp,
		OkCodes:      []int{202},
	})

	if err != nil {
		sdkErr := lserr.ErrorHandler(err,
			lserr.WithErrorOutOfPoc(errResp, err))

		return nil, sdkErr
	}

	return response.ToServerObject(), nil
}

func List(sc *client.ServiceClient, opts IListOptsBuilder) ([]*objects.Server, error) {
	url, err := opts.ToListQuery()

	if err != nil {
		return nil, err
	}

	resp := NewListResponse()
	url = fmt.Sprintf("%s%s", listURL(sc, opts), url)
	fmt.Println(url)
	_, err = sc.Get(url, &client.RequestOpts{
		JSONResponse: resp,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}
	return resp.ToServerList()
}

func UpdateSecGroups(sc *client.ServiceClient, opts IUpdateSecGroupsOptsBuilder) (*objects.Server, error) {
	response := NewUpdateSecGroupsResponse()
	_, err := sc.Put(updateSecGroupsServerURL(sc, opts), &client.RequestOpts{
		JSONBody:     opts.ToRequestBody(),
		JSONResponse: response,
		OkCodes:      []int{202},
	})

	if err != nil {
		return nil, err
	}

	return response.ToServerObject(), nil
}
