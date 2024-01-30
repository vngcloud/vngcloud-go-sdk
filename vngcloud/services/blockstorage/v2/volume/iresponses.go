package volume

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

// ******************************************** Response of CreateVolume API *******************************************

type ICreateResponse interface {
	ToVolumeObject() *objects.Volume
}

// ******************************************** Response of GetVolume API **********************************************

type IGetResponse interface {
	ToVolumeObject() *objects.Volume
}

// ******************************************** Response of ListVolume API *********************************************

type IListResponse interface {
	ToListVolumeObjects() []*objects.Volume
	ToVolumeObject(pIdx int) *objects.Volume
	NextPage() string
}

// ******************************************* Response of ListAllVolume API *******************************************

type IListAllResponse interface {
	ToListVolumeObjects() []*objects.Volume
}
