package snapshot

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/client"
)

// createURL generates the URL for creating a Snapshot based on the given options.
func createURL(psc *lsclient.ServiceClient, popts ICreateOptsBuilder) string {
	return psc.ServiceURL(
		popts.GetProjectID(),
		"volumes",
		popts.GetVolumeID(),
		"snapshots")
}

func deleteURL(psc *lsclient.ServiceClient, popts IDeleteOptsBuilder) string {
	return psc.ServiceURL(
		popts.GetProjectID(),
		"volumes",
		popts.GetVolumeID(),
		"snapshots",
		popts.GetSnapshotID(),
	)
}

func listVolumeSnapshotURL(psc *lsclient.ServiceClient, popts IListVolumeOptsBuilder) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		popts.GetProjectID(),
		"volumes",
		popts.GetVolumeID(),
		"snapshots",
	) + query
}
