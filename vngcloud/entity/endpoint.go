package entity

type Endpoint struct {
	IPv4Address string
	EndpointUrl string
	Status      string
	VpcId       string
	Id          string
	Name        string
}

func (s *Endpoint) IsUsable() bool {
	return s.Status == "ACTIVE"
}

func (s *Endpoint) GetId() string {
	return s.Id
}

func (s *Endpoint) GetName() string {
	return s.Name
}

func (s *Endpoint) GetIPv4Address() string {
	return s.IPv4Address
}

func (s *Endpoint) GetEndpointUrl() string {
	return s.EndpointUrl
}

func (s *Endpoint) GetVpcId() string {
	return s.VpcId
}

func (s *Endpoint) GetStatus() string {
	return s.Status
}

func (s *Endpoint) IsError() bool {
	return s.Status == "ERROR"
}

type ListEndpoints struct {
	Items     []*Endpoint
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (s *ListEndpoints) Len() int {
	return len(s.Items)
}

func (s *ListEndpoints) At(pidx int) *Endpoint {
	if pidx < 0 || pidx >= s.Len() {
		return nil
	}

	return s.Items[pidx]
}
