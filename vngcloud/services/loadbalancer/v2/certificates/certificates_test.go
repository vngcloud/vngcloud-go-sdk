package certificates

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
	caID      = ""
)

func PointerOf[T any](t T) *T {
	return &t
}

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

func TestImport(t *testing.T) {
	vlb := NewSC()

	opt := &ImportOpts{
		Name:             "annd2-testtt",
		Type:             "TLS/SSL",
		Certificate:      "",
		CertificateChain: PointerOf[string](""),
		Passphrase:       PointerOf[string](""),
		PrivateKey:       PointerOf[string](""),
	}
	opt.ProjectID = projectID

	resp, err := Import(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	fmt.Printf("%+v\n", resp)
}

func TestGet(t *testing.T) {
	vlb := NewSC()

	opt := &GetOpts{}
	opt.ProjectID = projectID
	opt.CaID = caID

	resp, err := Get(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	fmt.Printf("%+v\n", resp)
}

func TestList(t *testing.T) {
	vlb := NewSC()

	opt := &ListOpts{}
	opt.ProjectID = projectID

	resp, err := List(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	fmt.Printf("%+v\n", resp)
	for _, v := range resp {
		fmt.Printf("%+v\n", v)
	}
}

func TestDelete(t *testing.T) {
	vlb := NewSC()

	opt := &DeleteOpts{}
	opt.ProjectID = projectID
	opt.CaID = caID

	err := Delete(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
