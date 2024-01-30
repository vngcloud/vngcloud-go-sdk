package snapshot

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/pagination"
)

func List(pSc *client.ServiceClient, pOpts IListOptsBuilder) *pagination.Pager {
	url := listURL(pSc, pOpts)
	return pagination.NewPager(pSc, url, pOpts,
		func() interface{} {
			return NewListResponse()
		},
		func(r interface{}) pagination.IPage {
			resp := r.(*ListResponse)

			if pOpts.GetVolumeID()+pOpts.GetName() != "" {
				result := new(ListResponse)
				result.Items = make([]Snapshot, 0)
				for _, snapshot := range resp.Items {
					if snapshot.VolumeId == pOpts.GetVolumeID() || snapshot.Name == pOpts.GetName() {
						if (pOpts.GetStatus() != "" && pOpts.GetStatus() == snapshot.Status) ||
							pOpts.GetStatus() == "" {
							result.Items = append(result.Items, snapshot)
						}
					}
				}

				result.Page = resp.Page
				result.PageSize = resp.PageSize
				result.TotalPage = resp.TotalPage
				result.TotalItem = resp.TotalItem

				return result
			}

			return resp
		},
	)
}
