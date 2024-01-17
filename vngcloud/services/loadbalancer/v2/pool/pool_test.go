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
	projectID = ""
	lbID      = ""
	poolID    = ""
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
func TestCreatePool(t *testing.T) {
	vlb := NewSC()

	opt := &CreateOpts{
		Algorithm:     CreateOptsAlgorithmOptSourceIP, // CreateOptsAlgorithmOptLeastConn, CreateOptsAlgorithmOptRoundRobin
		PoolName:      "annd2",
		PoolProtocol:  CreateOptsProtocolOptHTTP,
		Stickiness:    true,
		TLSEncryption: true,
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
		Members: []Member{
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
	for _, lb := range resp {
		fmt.Println(lb)
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
		Members: []Member{
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