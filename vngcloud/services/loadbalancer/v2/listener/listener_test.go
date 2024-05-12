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
	listenerID    = ""
	defaultPoolID = ""
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vLbURL       = "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway/v2"
		clientID     = "b6f68922bcb"
		clientSecret = "e4db50f56813ca"
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
		DefaultPoolId:        nil,
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

func TestCreateTLSListener(t *testing.T) {
	vlb := NewSC()

	opt := &CreateOpts{
		AllowedCidrs:                "0.0.0.0/0",
		DefaultPoolId:               nil,
		ListenerName:                "annd2_test_listener",
		ListenerProtocol:            CreateOptsListenerProtocolOptHTTPS,
		ListenerProtocolPort:        445,
		TimeoutClient:               100,
		TimeoutConnection:           100,
		TimeoutMember:               100,
		CertificateAuthorities:      &[]string{"secret-xxxxxxxxxxxxx", "secret-yyyyyyyyyyyyyyy"},
		ClientCertificate:           PointerOf(""),
		DefaultCertificateAuthority: PointerOf("secret-xxxxxxxxxxxxx"),
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
	for _, item := range resp {
		fmt.Printf("item: %+v\n", item)
	}
}

func TestUpdateListener(t *testing.T) {
	vlb := NewSC()

	opt := &UpdateOpts{
		AllowedCidrs:      "0.0.0.0/0",
		DefaultPoolId:     "",
		TimeoutClient:     100,
		TimeoutConnection: 100,
		TimeoutMember:     100,
		Headers:           []string{"X-Forwarded-For", "X-Forwarded-Proto", "X-Forwarded-Port"},
		// ClientCertificate:           nil,
		// DefaultCertificateAuthority: nil,
		ClientCertificate:           PointerOf[string]("secret-xxxxxxxxxx"),
		DefaultCertificateAuthority: PointerOf[string]("secret-xxxxxxxxxx"),
	}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.ListenerID = listenerID

	err := Update(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func TestDeleteListener(t *testing.T) {
	vlb := NewSC()

	opt := &DeleteOpts{}
	opt.ProjectID = "pro-462803f3-6858-466f-bf05-df2b33faa360"
	opt.LoadBalancerID = "lb-af85f586-5d71-435c-xxxx-ea0c0dc26655"
	opt.ListenerID = "lb-af85f586-5d71-435c-8fe4-ea0c0dc26655"

	err := Delete(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error)
	}
}
