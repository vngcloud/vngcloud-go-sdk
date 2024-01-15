package subnet

func NewGetOpts(pProjectID, pNetworkUUID, pSubnetUUID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.CommonNetworkUUID = pNetworkUUID
	opts.CommonSubnetUUID = pSubnetUUID
	return opts
}
