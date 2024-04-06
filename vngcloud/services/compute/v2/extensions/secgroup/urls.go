package secgroup

import "github.com/vngcloud/vngcloud-go-sdk/client"

func updateSecgroupURL(sc *client.ServiceClient, opts IUpdateSecgroupOptsBuilder) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"servers",
		opts.GetServerID(),
		"update-sec-group",
	)
}
