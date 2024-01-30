package cluster

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

// ******************************************** Response of GetCluster API *********************************************

type IGetResponse interface {
	ToClusterObject() *objects.Cluster
}

// *************************************** Response of Update Secgroup rule API ****************************************

type IUpdateSecgroupResponse interface {
	ToListClusterSecgroupRuleObjects() []*objects.ClusterSecgroup
}
