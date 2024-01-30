package policy

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
	lbID       = ""
	listenerID = ""
	poolID     = ""
	policyID   = ""
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
func TestListPolicyOfListener(t *testing.T) {
	vlb := NewSC()

	opt := &ListOptsBuilder{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.ListenerID = listenerID

	resp, err := List(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
	for _, lb := range resp {
		fmt.Println(lb)
	}
}

func TestCreateL7PolicyOfListener(t *testing.T) {
	vlb := NewSC()

	opt := &CreateOptsBuilder{
		Name:             "l7policy",
		Action:           PolicyOptsActionOptREJECT,
		RedirectPoolID:   poolID,
		RedirectURL:      "https://www.example.com",
		RedirectHTTPCode: 301,
		KeepQueryString:  true,
		Rules: []Rule{
			{
				CompareType: PolicyOptsCompareTypeOptCONTAINS,
				RuleType:    PolicyOptsRuleTypeOptHOSTNAME,
				RuleValue:   "www.example.com",
			},
		},
	}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.ListenerID = listenerID

	resp, err := Create(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
}

func TestDeleteL7PolicyOfListener(t *testing.T) {
	vlb := NewSC()

	opt := &DeleteOptsBuilder{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.ListenerID = listenerID
	opt.PolicyID = policyID

	err := Delete(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func TestUpdateL7PolicyOfListener(t *testing.T) {
	vlb := NewSC()

	opt := &UpdateOptsBuilder{
		Action:         PolicyOptsActionOptREDIRECTTOPOOL,
		RedirectPoolID: poolID,
		Rules: []Rule{
			{
				CompareType: PolicyOptsCompareTypeOptCONTAINS,
				RuleType:    PolicyOptsRuleTypeOptHOSTNAME,
				RuleValue:   "www.example.com",
			},
			{
				CompareType: PolicyOptsCompareTypeOptCONTAINS,
				RuleType:    PolicyOptsRuleTypeOptHOSTNAME,
				RuleValue:   "www.example7.com",
			},
		},
	}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.ListenerID = listenerID
	opt.PolicyID = policyID

	err := Update(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func TestGet(t *testing.T) {
	vlb := NewSC()

	opt := &GetOptsBuilder{}
	opt.ProjectID = projectID
	opt.LoadBalancerID = lbID
	opt.ListenerID = listenerID
	opt.PolicyID = policyID

	resp, err := Get(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
	for _, lb := range resp.L7Rules {
		fmt.Println(lb)
	}
}
