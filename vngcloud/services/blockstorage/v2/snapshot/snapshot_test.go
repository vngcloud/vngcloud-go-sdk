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
	projectID  = "pro-462803f3-6858-466f-bf05-df2b33faa360"
	volID      = "vol-a31e5c00-470a-4165-bf25-d3e469a2ce86"
	snapshotID = "snap-vol-pt-278a66d3-6a3e-44f0-9df9-06cf714abaa7"
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vServerURL   = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway/v2"
		clientID     = "cc136c74480a"
		clientSecret = "f595635bd6fabe28"
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

	opt := NewListVolumeOpts(projectID, volID, 1, 10)

	resp, err := ListVolumeSnapshot(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("resp: %+v\n", resp)
}

func TestDeleteSnapshot(t *testing.T) {
	vserverClient := NewSC()

	opt := NewDeleteOpts(projectID, volID, snapshotID)

	err := Delete(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("Delete snapshot successfully\n")
}
