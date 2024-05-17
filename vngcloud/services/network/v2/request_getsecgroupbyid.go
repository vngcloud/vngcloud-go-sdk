package v2

type GetSecgroupByIdRequest struct {
	SecgroupCommon
}

func NewGetSecgroupByIdRequest(psecgroupId string) IGetSecgroupByIdRequest {
	opt := new(GetSecgroupByIdRequest)
	opt.SecgroupId = psecgroupId
	return opt
}
