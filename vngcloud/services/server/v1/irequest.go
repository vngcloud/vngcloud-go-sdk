package v1

type ICreateSystemTagRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) ICreateSystemTagRequest
	ToMap() map[string]interface{}
	AddTag(pkey, pvalue string) ICreateSystemTagRequest
	ParseUserAgent() string
	GetResourceId() string
	GetResourceType() ResourceType
}
