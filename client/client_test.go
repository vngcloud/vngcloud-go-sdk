package client

import "testing"

func TestNewClient(t *testing.T) {
	vngcloud := NewClient()
	vngcloud.IamGateway().V2().IdentityService()
}
