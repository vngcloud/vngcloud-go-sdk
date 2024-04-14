package volume

import "github.com/vngcloud/vngcloud-go-sdk/client"

func listURL(pSc *client.ServiceClient, popts IListOptsBuilder) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return pSc.ServiceURL(
		popts.GetProjectID(),
		"volumes") + query
}

func listAllURL(pSc *client.ServiceClient, pOpts IListAllOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes")
}

func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes")
}

func deleteURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes",
		pOpts.GetVolumeID())
}

func getURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes",
		pOpts.GetVolumeID())
}
