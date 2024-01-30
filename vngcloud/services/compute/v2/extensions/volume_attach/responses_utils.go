package volume_attach

func NewCreateResponse() ICreateResponse {
	return new(CreateResponse)
}

func NewDeleteResponse() IDeleteResponse {
	return new(DeleteResponse)
}
