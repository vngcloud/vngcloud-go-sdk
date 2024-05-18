package test

import (
	ltesting "testing"
)

func TestCreateServerFailed(t *ltesting.T) {
	vngcloud := validSdkConfig()
	vngcloud.VServerGateway().V2().ComputeService()
}
