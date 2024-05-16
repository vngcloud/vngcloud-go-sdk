package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type PortalServiceV1 struct {
	PortalClient lsclient.IServiceClient
}

func (s *PortalServiceV1) GetPortalInfo(popts IGetPortalInfoRequest) (*lsentity.Portal, lsdkErr.ISdkError) {
	//resp := new(GetPortalInfoResponse)
	//errResp := lsdkErr.NewErrorResponse()
	//req := popts.ToRequest().WithJsonResponse(resp).WithJsonError(errResp)
	//sdkErr := s.PortalClient.Get(getPortalInfoUrl(s.PortalClient, popts), req)
	//
	//fmt.Println("resp: ", req)
	//
	//if sdkErr != nil {
	//	return nil, lsdkErr.ErrorHandler(sdkErr.GetError())
	//}
	//
	//return resp.ToEntityPortal(), nil
	return nil, nil
}
