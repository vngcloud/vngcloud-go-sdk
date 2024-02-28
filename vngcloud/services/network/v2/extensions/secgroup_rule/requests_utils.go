package secgroup_rule

func NewCreateOpts(pProjectID string, pSecgroupUUID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.SecgroupUUID = pSecgroupUUID
	return pOpts
}

func NewDeleteOpts(pProjectID string, pSecgroupUUID string, pRuleUUID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.SecgroupUUID = pSecgroupUUID
	opts.SecgroupRuleUUID = pRuleUUID
	return opts
}

func NewListRulesBySecgroupIDOpts(pProjectID, pSecgroupUUID string) IListRulesBySecgroupIDOptsBuilder {
	opts := new(ListRulesBySecgroupIDOpts)
	opts.ProjectID = pProjectID
	opts.SecgroupUUID = pSecgroupUUID
	return opts
}
