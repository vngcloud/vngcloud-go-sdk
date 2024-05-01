package subnet

func NewGetOpts(pProjectID, pNetworkUUID, pSubnetUUID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.CommonNetworkUUID = pNetworkUUID
	opts.CommonSubnetUUID = pSubnetUUID
	return opts
}

func NewListByNetworkIDOpts(pProjectID, pNetworkUUID string) IListByNetworkIDOptsBuilder {
	opts := new(ListByNetworkIDOpts)
	opts.ProjectID = pProjectID
	opts.CommonNetworkUUID = pNetworkUUID
	return opts
}
