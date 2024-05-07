package loadbalancer

import (
	"fmt"
	lsListener "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2/listener"
	lsPool "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2/pool"
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
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vLbURL       = "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway/v2"
		clientID     = "b6f68928e8822bcb"
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

func NewSC2() *client.ServiceClient {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vLbURL       = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway/v2"
		clientID     = "b6f688d8e8822bcb"
		clientSecret = "e4db5f1f56813ca"
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
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestUpdate(t *testing.T) {
	vlb := NewSC()

	opt := &UpdateOpts{
		// PackageID: "lbp-f562b658-0fd4-4fa6-9c57-c1a803ccbf86",
		PackageID: "lbp-c531bc55-27d7-4a3e-be0b-eac265658a50",
	}
	opt.ProjectID = projectID
	opt.LoadBalancerID = loadbalancerID

	_, err := Update(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func TestCreateL4Full(t *testing.T) {
	_packageID := "lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa" // NLB_Small
	_subnetID := "sub-403b36d2-39fc-47c4-b40b-8df0ecb71045"
	_projectID := "pro-462803f3-6858-466f-bf05-df2b33faa360"

	vlb := NewSC()

	opt := &CreateOpts{
		Name:      "cuongdm3-test-lb",
		PackageID: _packageID,
		Scheme:    CreateOptsSchemeOptInternet,
		SubnetID:  _subnetID,
		Type:      CreateOptsTypeOptLayer4,
		Listener: &lsListener.CreateOpts{
			AllowedCidrs:         "0.0.0.0/0",
			ListenerName:         "cuongdm3-test-lb-listener",
			ListenerProtocol:     lsListener.CreateOptsListenerProtocolOptTCP,
			ListenerProtocolPort: 80,
			TimeoutConnection:    50,
			TimeoutClient:        50,
			TimeoutMember:        5,
		},
		Pool: &lsPool.CreateOpts{
			Algorithm:    lsPool.CreateOptsAlgorithmOptRoundRobin,
			PoolName:     "cuongdm3-test-lb-pool",
			PoolProtocol: lsPool.CreateOptsProtocolOptTCP,
			HealthMonitor: lsPool.HealthMonitor{
				HealthCheckProtocol: lsPool.CreateOptsHealthCheckProtocolOptTCP,
				HealthyThreshold:    3,
				UnhealthyThreshold:  3,
				Interval:            30,
				Timeout:             5,
			},
			Members: []*lsPool.Member{
				{
					Backup:      false,
					IpAddress:   "10.21.0.18",
					MonitorPort: 80,
					Name:        "cuongdm3-test-lb-member",
					Port:        80,
					Weight:      1,
				},
			},
		},
	}
	opt.ProjectID = _projectID

	resp, err := Create(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("%+v\n", resp)
}

func TestCreateTag(t *testing.T) {
	vlb := NewSC2()
	projectID = "pro-462803f3-6858-466f-bf05-df2b33faa360"
	loadbalancerID = "lb-acc8eaeb-9098-47fb-98f5-656af6a14d2d"

	opt := NewCreateTagOpts(projectID, loadbalancerID, map[string]string{
		"tag1": "value1",
		"tag2": "value2",
	})

	err := CreateTag(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func TestList(t *testing.T) {
	vlb := NewSC()
	projectID = "pro-462803f3-6858-466f-bf05-df2b33faa360"

	opt := NewListOpts(projectID, &ListOpts{
		Page: 1,
		Size: 20,
		Tags: []ListOptsTag{
			{
				Key:   "tag1",
				Value: "value1",
			},
			{
				Key: "tag2",
			},
		},
	})

	resp, err := List(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for _, lb := range resp {
		fmt.Printf("%+v\n", lb.Name)
	}
}

func TestListTags(t *testing.T) {
	vlb := NewSC2()
	projectID = "pro-462803f3-6858-466f-bf05-df2b33faa360"
	loadbalancerID = "lb-76654bb0-72fe-4694-9f43-cc1be0317a5e"

	opt := NewListTagsOpts(projectID, loadbalancerID)

	resp, err := ListTags(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	for _, tag := range resp {
		fmt.Printf("%+v\n", tag)
	}
}

func TestUpdateTag(t *testing.T) {
	vlb := NewSC2()
	projectID = "pro-462803f3-6858-466f-bf05-df2b33faa360"
	loadbalancerID = "lb-3a886fec-c853-4ef9-974b-308d7587880a"

	opt := NewUpdateTagOpts(projectID, loadbalancerID, map[string]string{
		"vks-owned-cluster": "asdadadadasdas",
	})

	err := UpdateTag(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
