package secgroup

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
)

func UpdateSecgroups(sc *client.ServiceClient, opts *UpdateSecgroupOpts) error {
	_, err := sc.Put(updateSecgroupURL(sc, opts), &client.RequestOpts{
		OkCodes:  []int{202},
		JSONBody: opts,
	})

	if err != nil {
		return err
	}

	return nil
}
