package pagination

import "github.com/vngcloud/vngcloud-go-sdk/client"

func NewPager(pServiceClient *client.ServiceClient, pInitialURL string, pPageOpts IPageOpts,
	pCreateResponse func() interface{}, pCreatePage func(interface{}) IPage) *Pager {

	return &Pager{
		client:         pServiceClient,
		initialURL:     pInitialURL,
		createPage:     pCreatePage,
		createResponse: pCreateResponse,
		PageOpts:       pPageOpts,
	}
}
