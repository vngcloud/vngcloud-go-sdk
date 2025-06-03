package v1

type IGetPortalInfoRequest interface {
	GetBackEndProjectId() string
}

type IListProjectsRequest interface {
	AddUserAgent(pagent ...string) IListProjectsRequest
	ParseUserAgent() string
}
