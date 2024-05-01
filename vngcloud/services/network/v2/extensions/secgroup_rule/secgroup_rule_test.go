package secgroup_rule

import (
	"fmt"
	"testing"

	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
)

var (
	projectID  = ""
	secgroupID = ""
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vLbURL       = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway/v2"
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

func TestListBySecgroupID(t *testing.T) {
	vlb := NewSC()

	opt := new(ListRulesBySecgroupIDOpts)
	opt.ProjectID = projectID
	opt.SecgroupUUID = secgroupID

	resp, err := ListRulesBySecgroupID(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("%+v\n", resp)
}

func TestCreateRule(t *testing.T) {
	vlb := NewSC()

	opt := NewCreateOpts(projectID, secgroupID, &CreateOpts{
		Description:     "test",
		Direction:       "ingress",
		EtherType:       "IPv4",
		PortRangeMax:    8080,
		PortRangeMin:    8080,
		Protocol:        "tcp",
		RemoteIPPrefix:  "0.0.0.0/0",
		SecurityGroupID: secgroupID,
	})

	resp, err := Create(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("%+v\n", resp)
}
