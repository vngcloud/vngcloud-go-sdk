package volume_attach

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
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vServerURL   = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway/v2"
		clientID     = "c3f65a95fc29"
		clientSecret = "8637fea696bf82"
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

func TestAttach(t *testing.T) {
	vserverClient := NewSC()
	serverID := "ins-869a1-fcdf6b3d4bd4"
	volumeID := "vol-cb69a80701c"

	opt := NewCreateOpts(projectID, serverID, volumeID)

	resp, err := Attach(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("resp: %+v\n", resp)
	}
}

func TestDettach(t *testing.T) {
	vserverClient := NewSC()
	serverID := "ins-869ad034b3d4bd4"
	volumeID := "vol-cb69da5b-acc9"

	opt := NewCreateOpts(projectID, serverID, volumeID)

	resp, err := Detach(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("resp: %+v\n", resp)
	}
}
