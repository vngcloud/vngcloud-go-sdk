package snapshot

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type ICreateResponse interface {
	ToSnapshotObject() *objects.Snapshot
}
