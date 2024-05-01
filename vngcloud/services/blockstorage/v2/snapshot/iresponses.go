package snapshot

import (
	lso "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type ICreateResponse interface {
	ToSnapshotObject() *lso.Snapshot
}

type IListVolumeResponse interface {
	ToSnapshotListObject() *lso.SnapshotList
}
