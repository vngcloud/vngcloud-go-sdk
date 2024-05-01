package tag

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
	projectID = ""
	serverID  = ""
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
	vserverClient := NewSC()

	opt := NewGetOpts(projectID, serverID)

	resp, err := Get(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("resp: %+v\n", resp)
	for _, tag := range resp {
		fmt.Printf("Tag: %+v\n", tag)
	}
}

func TestUpdate(t *testing.T) {
	vserverClient := NewSC()

	opt := NewUpdateOpts(projectID, serverID, &UpdateOpts{
		ResourceType: "LOAD-BALANCER",
		ResourceID:   serverID,
		TagRequestList: []TagRequestItem{
			{IsEdited: false, Key: "key1", Value: "value1"},
			{IsEdited: false, Key: "key2", Value: "value1"},
			{IsEdited: false, Key: "key3", Value: "value1"},
			{IsEdited: false, Key: "key4", Value: "value1"},
			{IsEdited: false, Key: "key5", Value: "value1"},
		},
	})

	resp, err := Update(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("resp: %+v\n", resp)
	for _, tag := range resp {
		fmt.Printf("Tag: %+v\n", tag)
	}
}
