package volume_actions

func NewResizeResponse() IResizeResponse {
	return &ResizeResponse{}
}

func NewBackendVolumeResponse() IMappingResponse {
	return &MappingResponse{}
}
