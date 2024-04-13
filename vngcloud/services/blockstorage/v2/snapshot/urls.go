package snapshot

import "github.com/vngcloud/vngcloud-go-sdk/client"

// createURL generates the URL for creating a Snapshot based on the given options.
func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes",
		pOpts.GetVolumeID(),
		"snapshots")
}

func deleteURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes",
		pOpts.GetVolumeID(),
		"snapshots",
		pOpts.GetSnapshotID(),
	)
}

func listVolumeSnapshotURL(psc *client.ServiceClient, popts IListVolumeOptsBuilder) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		popts.GetProjectID(),
		"snapshot-volumes",
	) + query
}
