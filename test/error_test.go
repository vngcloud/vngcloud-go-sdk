package test

import (
	"fmt"
	"testing"

	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func TestDeleteListener(t *testing.T) {
	errorMsg := "Listener id lis-760f69b4-0b24-4813-9854-9e2c6a85e214 is not belong to load balancer id lb-03410465-5354-4a45-adb0-32c9a75d2a06"
	perrResp := &lserr.NormalErrorResponse{
		Message: errorMsg,
	}

	serr := lserr.ErrorHandler(fmt.Errorf("haha"), lserr.WithErrorListenerNotFound(perrResp))
	t.Log(serr)
}
