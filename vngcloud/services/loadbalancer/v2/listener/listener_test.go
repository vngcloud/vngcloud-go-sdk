package listener

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
	projectID     = ""
	lbID          = ""
	defaultPoolID = ""
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
func TestCreateListener(t *testing.T) {
	vlb := NewSC()

	opt := &CreateOpts{
		AllowedCidrs:         "0.0.0.0/0",
		DefaultPoolId:        defaultPoolID,
		ListenerName:         "annd2_test_listener",
		ListenerProtocol:     CreateOptsListenerProtocolOptHTTP,
		ListenerProtocolPort: 90,
		TimeoutClient:        100,
		TimeoutConnection:    100,
		TimeoutMember:        100,
	}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID

	resp, err := Create(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
}

func TestGetBasedLoadBalancer(t *testing.T) {
	vlb := NewSC()

	opt := &GetBasedLoadBalancerOpts{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID

	resp, err := GetBasedLoadBalancer(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
}
