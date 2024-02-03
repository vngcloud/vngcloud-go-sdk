package network

func NewGetOpts(pProjectID, pNetworkUUID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.CommonNetworkUUID = pNetworkUUID
	return opts
}
