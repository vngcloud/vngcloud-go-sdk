package v2

type GetQuotaByNameRequest struct {
	Name QuotaName
}

func NewGetQuotaByNameRequest(pname QuotaName) IGetQuotaByNameRequest {
	return &GetQuotaByNameRequest{
		Name: pname,
	}
}

func (s *GetQuotaByNameRequest) GetName() QuotaName {
	return s.Name
}
