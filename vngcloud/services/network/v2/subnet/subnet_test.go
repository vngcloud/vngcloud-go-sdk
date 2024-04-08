package subnet

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
	networkID = ""
	subnetID  = ""
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = ""
		vServerURL   = ""
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
		vServerURL,
		provider,
		"vlb")
	return vlb
}

func TestGet(t *testing.T) {
	vlb := NewSC()

	opt := NewGetOpts(projectID, networkID, subnetID)

	resp, err := Get(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	fmt.Printf("%+v\n", resp)
}

func TestListByNetworkID(t *testing.T) {
	vlb := NewSC()

	opt := NewListByNetworkIDOpts(projectID, networkID)

	resp, err := ListByNetworkID(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	fmt.Printf("%+v\n", resp)
	for _, s := range resp {
		fmt.Println(s)
	}
}
