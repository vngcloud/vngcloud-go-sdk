package loadbalancer

import (
	"fmt"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
	"testing"
)

func TestListBySubnetID(t *testing.T) {
	// Load file env
	var (
		clientID     string
		clientSecret string
		projectID    string
		subnetID     string
	)

	clientID, clientSecret, projectID, subnetID = "", "", "", ""
	// Override clientID, clientSecret, projectID, subnetID here

	identityURL := "https://iamapis.vngcloud.vn/accounts-api/v2"
	vLbURL := "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway/v2"

	provider, _ := vcontainer.NewClient(identityURL)
	vcontainer.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: identityURL,
		},
	})

	vlb, _ := vcontainer.NewServiceClient(
		vLbURL,
		provider,
		"vlb")

	opt := new(ListBySubnetIDOptsBuilder)
	opt.ProjectID = projectID
	opt.SubnetID = subnetID

	resp, err := ListBySubnetID(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("LB: %s\n", resp[0].Name)
}
