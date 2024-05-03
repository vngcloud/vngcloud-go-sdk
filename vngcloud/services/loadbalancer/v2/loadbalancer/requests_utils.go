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

func NewListOpts(pprojectID string, popts *ListOpts) IListOptsBuilder {
	popts.ProjectID = pprojectID
	return popts
}

func NewUpdateOpts(pProjectID, pLoadBalancerID string, pOpts *UpdateOpts) IUpdateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLoadBalancerID
	return pOpts
}

func NewCreateTagOpts(pprojectID, plbId string, ptags map[string]string) ICreateTagOptsBuilder {
	opts := new(CreateTagOpts)
	opts.ProjectID = pprojectID
	opts.LoadBalancerID = plbId
	opts.ResourceType = "LOAD-BALANCER"
	opts.ResourceID = plbId

	if ptags != nil {
		opts.TagRequestList = make([]CreateTagOptsTag, 0)
		for key, value := range ptags {
			opts.TagRequestList = append(opts.TagRequestList, CreateTagOptsTag{
				Key:   key,
				Value: value,
			})
		}
	}

	return opts
}
