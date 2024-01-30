package loadbalancer

func NewCreateOpts(pProjectID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	return pOpts
}

func NewGetOpts(pProjectID, pLoadBalancerID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLoadBalancerID
	return opts
}

func NewDeleteOpts(pProjectID, pLoadBalancerID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLoadBalancerID
	return opts
}

func NewListBySubnetIDOpts(pProjectID, pSubnetID string) IListBySubnetIDOptsBuilder {
	opts := new(ListBySubnetIDOpts)
	opts.ProjectID = pProjectID
	opts.SubnetID = pSubnetID
	return opts
}

func NewListOpts(pProjectID string) IListOptsBuilder {
	opts := new(ListOpts)
	opts.ProjectID = pProjectID
	return opts
}

func NewUpdateOpts(pProjectID, pLoadBalancerID string, pOpts *UpdateOpts) IUpdateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLoadBalancerID
	return pOpts
}
