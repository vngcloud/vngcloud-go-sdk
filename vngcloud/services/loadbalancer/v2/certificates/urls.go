package certificates

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
)

func importURL(pSc *client.ServiceClient, pOpts IImportOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"cas")
}

func getURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"cas",
		pOpts.GetCertificateID())
}

func deleteURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"cas",
		pOpts.GetCertificateID())
}

func listURL(pSc *client.ServiceClient, pOpts IListOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"cas")
}
