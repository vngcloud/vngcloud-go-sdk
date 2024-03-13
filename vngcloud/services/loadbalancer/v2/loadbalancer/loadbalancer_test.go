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
	projectID      = ""
	subnetID       = ""
	packageId      = ""
	loadbalancerID = ""
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

	opt := new(ListBySubnetIDOpts)
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

func TestCreate(t *testing.T) {
	vlb := NewSC()

	opt := &CreateOpts{
		Name:      "annd2_test",
		PackageID: packageId,
		Scheme:    CreateOptsSchemeOptInternet,
		SubnetID:  subnetID,
		Type:      CreateOptsTypeOptLayer7,
	}
	opt.ProjectID = projectID

	resp, err := Create(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("%+v\n", resp)
}

func TestGet(t *testing.T) {
	vlb := NewSC()

	opt := &GetOpts{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = loadbalancerID

	resp, err := Get(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("%+v\n", resp)
}

func TestResizeListener(t *testing.T) {
	vlb := NewSC()

	opt := &ResizeOpts{
		PackageId: "lbp-f562b658-0fd4-4fa6-9c57-c1a803ccbf86",
		// PackageId: "lbp-c531bc55-27d7-4a3e-be0b-eac265658a50",
	}
	opt.ProjectID = projectID
	opt.LoadBalancerID = loadbalancerID

	err := Resize(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
