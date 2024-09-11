package v2

func NewListAllServersBySecgroupIdRequest(psecgroupId string) IListAllServersBySecgroupIdRequest {
	opt := new(ListAllServersBySecgroupIdRequest)
	opt.SecgroupId = psecgroupId
	return opt
}

type ListAllServersBySecgroupIdRequest struct {
	SecgroupCommon
}
