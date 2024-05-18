package v2

func NewGetNetworkByIdRequest(pnetworkId string) IGetNetworkByIdRequest {
	opt := new(GetNetworkByIdRequest)
	opt.NetworkId = pnetworkId
	return opt
}

type GetNetworkByIdRequest struct {
	NetworkCommon
}
