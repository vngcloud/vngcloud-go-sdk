package snapshot

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewListVolumeResponse() IListVolumeResponse {
	return new(ListVolumeSnapshotResponse)
}
