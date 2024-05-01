package pool

func NewCreateOpts(pProjectID, pLbID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLbID
	return pOpts
}

func NewListPoolsBasedLoadBalancerOpts(pProjectID, pLbID string) IListPoolsBasedLoadBalancerOptsBuilder {
	opts := new(ListPoolsBasedLoadBalancerOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLbID
	return opts
}

func NewDeleteOpts(pProjectID, pLbID, pPoolID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLbID
	opts.PoolID = pPoolID
	return opts
}

func NewUpdatePoolMembersOpts(pProjectID, pLbID, pPoolID string, pOpts *UpdatePoolMembersOpts) IUpdatePoolMembersOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLbID
	pOpts.PoolID = pPoolID
	return pOpts
}

func NewGetMembersOpts(pprojectId, plbId, ppoolId string) IGetMemberOptsBuilder {
	opts := new(GetMemberOpts)
	opts.ProjectID = pprojectId
	opts.LoadBalancerID = plbId
	opts.PoolID = ppoolId
	return opts
}
