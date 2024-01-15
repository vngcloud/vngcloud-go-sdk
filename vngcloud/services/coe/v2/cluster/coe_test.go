package cluster

import (
	"fmt"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
	"testing"
)

func TestCreate(t *testing.T) {
	clientID := ""
	clientSecret := ""
	projectID := ""
	clusterID := ""

	identityURL := "https://iamapis.vngcloud.vn/accounts-api/v2"
	provider, _ := vcontainer.NewClient(identityURL)
	vcontainer.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: "https://iamapis.vngcloud.vn/accounts-api/v2",
		},
	})

	vserver, _ := vcontainer.NewServiceClient(
		"https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway/v2",
		provider,
		"vserver")

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
