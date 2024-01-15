package secgroup_rule

import "github.com/vngcloud/vngcloud-go-sdk/client"

func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"secgroups",
		pOpts.GetSecgroupUUID(),
		"secgroupRules")
}
