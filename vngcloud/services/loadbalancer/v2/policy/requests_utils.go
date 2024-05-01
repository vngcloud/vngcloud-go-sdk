package policy

func NewListOpts(pProjectID string) IListOptsBuilder {
	opts := new(ListOptsBuilder)
	opts.ProjectID = pProjectID
	return opts
}

func NewCreateOpts(pProjectID, pLbID, pLisID string, pOpts *CreateOptsBuilder) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLbID
	pOpts.ListenerID = pLisID
	return pOpts
}

func NewDeleteOpts(pProjectID, pLbID, pListenerID, pPolicyID string) IDeleteOptsBuilder {
	opts := new(DeleteOptsBuilder)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLbID
	opts.ListenerID = pListenerID
	opts.PolicyID = pPolicyID
	return opts
}
