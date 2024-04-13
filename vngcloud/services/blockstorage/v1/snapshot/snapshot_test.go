package snapshot

import (
	"fmt"
	"testing"

	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
)

func PointerOf[T any](t T) *T {
	return &t
}

var (
	projectID = "pro-462803f3-6858-466f-bf05-df2b33faa360"
	volID     = "vol-a31e5c00-470a-4165-bf25-d3e469a2ce86"
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vServerURL   = "https://hcm-3.console.vngcloud.vn/vserver/vbackup-gateway/v1"
		clientID     = "cc1364480a"
		clientSecret = "f59560bd6fabe28"
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
		vServerURL,
		provider,
		"vlb")
	return vlb
}

func TestListVolumeSnapshot(t *testing.T) {
	vserverClient := NewSC()

	opt := NewListVolumeOpts(projectID, "", 1, 10)

	resp, err := ListVolumeSnapshot(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("resp: %+v\n", resp)
}
