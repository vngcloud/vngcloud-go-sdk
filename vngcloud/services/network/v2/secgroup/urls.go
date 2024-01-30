package secgroup

import "github.com/vngcloud/vngcloud-go-sdk/client"

func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(pOpts.GetProjectID(), "secgroups")
}

func deleteURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(pOpts.GetProjectID(), "secgroups", pOpts.GetSecgroupUUID())
}

func getURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(pOpts.GetProjectID(), "secgroups", pOpts.GetSecgroupUUID())
}

func listURL(pSc *client.ServiceClient, pOpts IListOptsBuilder) string {
	query, _ := pOpts.ToListQuery()
	return pSc.ServiceURL(pOpts.GetProjectID(), "secgroups") + query
}
