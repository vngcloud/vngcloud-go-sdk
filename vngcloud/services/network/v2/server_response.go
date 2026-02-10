package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type ListAllServersBySecgroupIdResponse struct {
	Data []struct {
		Name   string `json:"name"`
		UUID   string `json:"uuid"`
		Status string `json:"status"`
	} `json:"data"`
}

func (s *ListAllServersBySecgroupIdResponse) ToEntityListServers() *lsentity.ListServers {
	servers := make([]*lsentity.Server, 0, len(s.Data))
	for _, server := range s.Data {
		servers = append(servers, &lsentity.Server{
			Name:   server.Name,
			Uuid:   server.UUID,
			Status: server.Status,
		})
	}
	return &lsentity.ListServers{
		Items: servers,
	}
}
