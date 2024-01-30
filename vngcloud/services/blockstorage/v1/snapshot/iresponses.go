package snapshot

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

// ******************************************* Response of List Snapshot API *******************************************

type IListResponse interface {
	ToListSnapshotObjects() []*objects.Snapshot
	ToSnapshotObject(pIdx int) *objects.Snapshot
	NextPage() string
}
