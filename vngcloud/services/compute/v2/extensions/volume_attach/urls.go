package volume_attach

import "github.com/vngcloud/vngcloud-go-sdk/client"

func createURL(sc *client.ServiceClient, opts ICreateOptsBuilder) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"volumes",
		opts.GetVolumeID(),
		"servers",
		opts.GetInstanceID(),
		"attach",
	)
}

func deleteURL(sc *client.ServiceClient, opts IDeleteOptsBuilder) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"volumes",
		opts.GetVolumeID(),
		"servers",
		opts.GetInstanceID(),
		"detach",
	)
}
