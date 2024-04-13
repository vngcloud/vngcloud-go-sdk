package snapshot

import (
	lsc "github.com/vngcloud/vngcloud-go-sdk/client"
)

// listURL generates the URL to get all the Snapshot.
func listVolumeSnapshotURL(psc *lsc.ServiceClient, popts IListVolumeOptsBuilder) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL("snapshot-volumes") + query
}
