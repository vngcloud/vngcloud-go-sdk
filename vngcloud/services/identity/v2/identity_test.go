package v2

import (
	"testing"

	"github.com/vngcloud/vngcloud-go-sdk/vngcloud"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
)

func NewSC() {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		clientID     = "1e882593-d431-4765-a429-8f0c17a88a7c"
		clientSecret = "7fbafe5f-a20f-4584-82cc-c2f9279360a9"
	)

	provider, _ := vngcloud.NewClient(identityURL)
	vngcloud.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: identityURL,
		},
	})
}

func TestGetToken(t *testing.T) {
	NewSC()
}
