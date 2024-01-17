package loadbalancer

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
	subnetID  = ""
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
		"vlb")
	return vlb
}

func TestListBySubnetID(t *testing.T) {
	vlb := NewSC()

	opt := new(ListBySubnetIDOptsBuilder)
	opt.ProjectID = projectID
	opt.SubnetID = subnetID

	resp, err := ListBySubnetID(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	for _, lb := range resp {
		fmt.Printf("%+v\n", lb.Name)
	}
}
