package listener

func NewCreateOpts(pProjectID, pLbID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLbID
	return pOpts
}

func NewGetBasedLoadBalancerOpts(pProjectID, pLbID string) IGetBasedLoadBalancerOptsBuilder {
	opts := new(GetBasedLoadBalancerOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLbID
	return opts
}

func NewDeleteOpts(pProjectID, pLbID, pListenerID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLbID
	opts.ListenerID = pListenerID
	return opts
}
