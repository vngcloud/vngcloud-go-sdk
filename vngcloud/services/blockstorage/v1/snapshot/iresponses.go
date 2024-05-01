package snapshot

import (
	lso "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

// ******************************************* Response of List Snapshot API *******************************************

type IListVolumeResponse interface {
	ToSnapshotListObject() *lso.SnapshotList
}
