package flavor

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	lsdkError "github.com/vngcloud/vngcloud-go-sdk/error"
	lserr "github.com/vngcloud/vngcloud-go-sdk/vngcloud/errors"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Get(sc *client.ServiceClient, opts IGetOptsBuilder) (*objects.Flavor, *lsdkError.SdkError) {
	response := NewGetResponse()
	errResp := lsdkError.NewErrorResponse()

	_, err := sc.Get(getURL(sc, opts), &client.RequestOpts{
		JSONResponse: response,
		JSONError:    errResp,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, lserr.ErrorHandler(err)
	}

	return response.ToFlavorObject(), nil
}
