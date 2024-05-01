package pool

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
	projectID = "pro-462803f3-6858-466f-bf05-df2b33faa360"
	lbID      = "lb-3dc42c25-2027-4786-a5e9-74072b84c208"
	poolID    = "pool-91075353-680c-4fd2-9e8d-e8e1fc3de1ef"
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vLbURL       = "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway/v2"
		clientID     = "cc136360-709c-4248-9358-e8e9674480a"
		clientSecret = "113ea24eecbf688e5"
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

func TestCreatePool(t *testing.T) {
	vlb := NewSC()

	opt := &CreateOpts{
		Algorithm:     CreateOptsAlgorithmOptSourceIP, // CreateOptsAlgorithmOptLeastConn, CreateOptsAlgorithmOptRoundRobin
		PoolName:      "annd2",
		PoolProtocol:  CreateOptsProtocolOptHTTP,
		Stickiness:    PointerOf(true),
		TLSEncryption: PointerOf(true),
		HealthMonitor: HealthMonitor{
			HealthyThreshold:    2,
			UnhealthyThreshold:  5,
			Interval:            30,
			Timeout:             15,
			HealthCheckProtocol: CreateOptsHealthCheckProtocolOptHTTP,
			// these param is invalid in some case, need to pass in pointer to ignore in the body
			DomainName:        PointerOf[string](""),
			HealthCheckPath:   PointerOf[string]("/heath"),
			SuccessCode:       PointerOf[string]("200,201,202"),
			HealthCheckMethod: PointerOf[CreateOptsHealthCheckMethodOpt]("GET"),
			HttpVersion:       PointerOf[CreateOptsHealthCheckHttpVersionOpt]("1.1"),
		},
		Members: []*Member{
			{Backup: true, Name: "member_1", IpAddress: "10.0.0.9", Port: 70, MonitorPort: 71, Weight: 1},
			{Backup: true, Name: "member_2", IpAddress: "10.0.0.9", Port: 72, MonitorPort: 71, Weight: 1},
		},
	}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID

	resp, err := Create(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
}

func TestListPoolsBasedLoadBalancer(t *testing.T) {
	vlb := NewSC()

	opt := &ListPoolsBasedLoadBalancerOpts{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID

	resp, err := ListPoolsBasedLoadBalancer(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
	for _, pool := range resp {
		fmt.Printf("pool: %s: %+v\n", pool.Name, pool.Members)
	}
}

func TestDelete(t *testing.T) {
	vlb := NewSC()

	opt := &DeleteOpts{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.PoolID = poolID

	err := Delete(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func TestUpdatePoolMembers(t *testing.T) {
	vlb := NewSC()

	opt := &UpdatePoolMembersOpts{
		Members: []*Member{
			{Backup: true, Name: "member_1", IpAddress: "10.0.0.11", Port: 70, MonitorPort: 71, Weight: 1},
			{Backup: true, Name: "member_2", IpAddress: "10.0.0.12", Port: 72, MonitorPort: 71, Weight: 1},
		},
	}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.PoolID = poolID

	err := UpdatePoolMembers(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func TestGet(t *testing.T) {
	vlb := NewSC()

	opt := &GetOpts{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.PoolID = poolID

	resp, err := Get(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
}

func TestGetMember(t *testing.T) {
	vlb := NewSC()

	opt := &GetMemberOpts{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.PoolID = poolID

	resp, err := GetMember(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
	for _, lb := range resp {
		fmt.Println(lb)
	}
}

func TestUpdate(t *testing.T) {
	vlb := NewSC()

	opt := &UpdateOpts{
		Algorithm:     CreateOptsAlgorithmOptLeastConn,
		Stickiness:    PointerOf(false),
		TLSEncryption: PointerOf(true),
		HealthMonitor: HealthMonitor{
			HealthyThreshold:    3,
			UnhealthyThreshold:  5,
			Interval:            30,
			Timeout:             15,
			HealthCheckProtocol: CreateOptsHealthCheckProtocolOptHTTP,
			// these param is invalid in some case, need to pass in pointer to ignore in the body
			DomainName:        PointerOf[string]("annd2.spacee"),
			HealthCheckPath:   PointerOf[string]("/heathz"),
			SuccessCode:       PointerOf[string]("200,201"),
			HealthCheckMethod: PointerOf[CreateOptsHealthCheckMethodOpt]("POST"),
			HttpVersion:       PointerOf[CreateOptsHealthCheckHttpVersionOpt]("1.1"),
		},
	}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.PoolID = poolID

	err := Update(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func TestGetHealthMonitor(t *testing.T) {
	vlb := NewSC()

	opt := &GetOpts{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.PoolID = poolID

	resp, err := GetHealthMonitor(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
}

func TestGetTotal(t *testing.T) {
	vlb := NewSC()

	opt := &GetOpts{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.PoolID = poolID

	resp, err := GetTotal(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
	fmt.Printf("resp: %+v\n", resp.HealthMonitor)
	fmt.Printf("resp: %+v\n", resp.HealthMonitor.DomainName)
	fmt.Printf("resp: %+v\n", *resp.HealthMonitor.DomainName)
}
