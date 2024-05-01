package tag

import "github.com/vngcloud/vngcloud-go-sdk/client"

func createURL(sc *client.ServiceClient, opts ICreateOptsBuilder) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"tag",
		"resource",
		opts.GetResourceID(),
	)
}

func getURL(sc *client.ServiceClient, opts IGetOptsBuilder) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"tag",
		"resource",
		opts.GetResourceID(),
	)
}

func updateURL(sc *client.ServiceClient, opts IGetOptsBuilder) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"tag",
		"resource",
		opts.GetResourceID(),
	)
}