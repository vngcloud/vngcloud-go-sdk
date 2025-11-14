package v1

func NewSystemTagRequest(resourceId string, resourceType ResourceType) ICreateSystemTagRequest {
	opt := new(CreateSystemTagRequest)
	opt.ResourceId = resourceId
	opt.ResourceType = resourceType
	return opt
}
