package cluster

import (
	"fmt"
	"testing"

	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
)

var (
	projectID = ""
	clusterID = ""
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = ""
		vLbURL       = ""
		clientID     = ""
		clientSecret = ""
	)

	provider, _ := vngcloud.NewClient(identityURL)
	vngcloud.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: identityURL,
		},
	})

	vlb, _ := vngcloud.NewServiceClient(
		vLbURL,
		provider,
		"vserver")
	return vlb
}

func TestGet(t *testing.T) {
	vserver := NewSC()

	opts := new(GetOpts)
	opts.ProjectID = projectID
	opts.ClusterID = clusterID

	resp, err := Get(vserver, opts)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("%v\n", resp)
}

func TestCreate(t *testing.T) {
	vserver := NewSC()

	opts := new(UpdateSecgroupOpts)
	opts.ClusterID = clusterID
	opts.Master = false
	opts.SecGroupIds = []string{"secg-31d4488c-xxxxxxxxx-a0ea-526843fe1c0f", "secg-8e015211-xxxxxxxxx-8023-f452cb471123", "secg-23e3b5f5-xxxxxxxxx-ae1e-eb168dbdd7ec"}
	opts.ProjectID = projectID

	resp, err := UpdateSecgroup(vserver, opts)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("%v\n", resp)
}
