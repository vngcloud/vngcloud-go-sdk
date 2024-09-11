package v2

func NewListAllServersBySubnetIdRequest(psecgroupId string) IListAllServersBySecgroupIdRequest {
	opt := new(ListAllServersBySecgroupIdRequest)
	opt.SecgroupId = psecgroupId
	return opt
}

type ListAllServersBySecgroupIdRequest struct {
	SecgroupCommon
}
