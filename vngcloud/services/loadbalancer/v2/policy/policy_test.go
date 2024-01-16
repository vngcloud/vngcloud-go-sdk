package policy

import (
	"fmt"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
	"testing"
)

func TestListPolicyOfListener(t *testing.T) {
	var (
		clientID     = ""
		clientSecret = ""
		projectID    = ""
		lbID         = ""
		listenerID   = ""
		identityURL  = ""
		vLbURL       = ""
	)

	provider, _ := vcontainer.NewClient(identityURL)
	vcontainer.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: identityURL,
		},
	})

	vlb, _ := vcontainer.NewServiceClient(
		vLbURL,
		provider,
		"vlb")

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
	var (
		clientID     = ""
		clientSecret = ""
		projectID    = ""
		lbID         = ""
		listenerID   = ""
		identityURL  = ""
		vLbURL       = ""
	)

	provider, _ := vcontainer.NewClient(identityURL)
	vcontainer.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: identityURL,
		},
	})

	vlb, _ := vcontainer.NewServiceClient(
		vLbURL,
		provider,
		"vlb")

	opt := &CreateOptsBuilder{
		Name:             "l7policy",
		Action:           CreateOptsActionOptREJECT,
		RedirectPoolID:   "pool-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		RedirectURL:      "https://www.example.com",
		RedirectHTTPCode: 301,
		KeepQueryString:  true,
		Rules: []Rule{
			{
				CompareType: CreateOptsCompareTypeOptCONTAINS,
				RuleType:    CreateOptsRuleTypeOptHOSTNAME,
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
	var (
		clientID     = ""
		clientSecret = ""
		projectID    = ""
		lbID         = ""
		listenerID   = ""
		policyID     = ""
		identityURL  = ""
		vLbURL       = ""
	)

	provider, _ := vcontainer.NewClient(identityURL)
	vcontainer.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: identityURL,
		},
	})

	vlb, _ := vcontainer.NewServiceClient(
		vLbURL,
		provider,
		"vlb")

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
