package secgroup_rule

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewListRulesBySecgroupIDResponse() IListRulesBySecgroupIDResponse {
	return &ListRulesBySecgroupIDResponse{}
}
