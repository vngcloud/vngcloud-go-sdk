package cluster

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type GetResponse struct {
	Data struct {
		UUID                        string   `json:"uuid"`
		ProjectID                   string   `json:"projectID"`
		ClusterID                   int      `json:"clusterID"`
		Name                        string   `json:"name"`
		Description                 string   `json:"description"`
		K8sVersionID                string   `json:"k8sVersionID"`
		Status                      string   `json:"status"`
		MasterCount                 int      `json:"masterCount"`
		NodeCount                   int      `json:"nodeCount"`
		EnabledLB                   bool     `json:"enabledLB"`
		EtcdVolumeSize              int      `json:"etcdVolumeSize"`
		EtcdVolumeTypeID            string   `json:"etcdVolumeTypeID"`
		MasterInstanceTypeID        string   `json:"masterInstanceTypeID"`
		MasterFlavorID              string   `json:"masterFlavorID"`
		NodeInstanceTypeID          string   `json:"nodeInstanceTypeID"`
		NodeFlavorID                string   `json:"nodeFlavorID"`
		BootVolumeSize              int      `json:"bootVolumeSize"`
		BootVolumeTypeID            string   `json:"bootVolumeTypeID"`
		DockerVolumeSize            int      `json:"dockerVolumeSize"`
		DockerVolumeTypeID          string   `json:"dockerVolumeTypeID"`
		K8sNetworkTypeID            string   `json:"k8sNetworkTypeID"`
		CalicoCidr                  string   `json:"calicoCidr"`
		NetworkID                   string   `json:"networkID"`
		SubnetID                    string   `json:"subnetID"`
		SSHKeyID                    string   `json:"sshKeyID"`
		CreatedAt                   string   `json:"createdAt"`
		ACLID                       string   `json:"aclID"`
		Endpoint                    string   `json:"endpoint"`
		AutoScalingEnabled          bool     `json:"autoScalingEnabled"`
		AutoHealingEnabled          bool     `json:"autoHealingEnabled"`
		MinNodeCount                int      `json:"minNodeCount"`
		MaxNodeCount                int      `json:"maxNodeCount"`
		IngressControllerEnabled    bool     `json:"ingressControllerEnabled"`
		AutoMonitoringEnabled       bool     `json:"autoMonitoringEnabled"`
		MasterFlavorName            string   `json:"masterFlavorName"`
		NodeFlavorName              string   `json:"nodeFlavorName"`
		K8sVersion                  string   `json:"k8sVersion"`
		K8sNetworkType              string   `json:"k8sNetworkType"`
		SSHKeyName                  string   `json:"sshKeyName"`
		NodegroupDefaultID          string   `json:"nodegroupDefaultID"`
		MasterClusterSecGroupIDList []string `json:"masterClusterSecGroupIDList"`
		MinionClusterSecGroupIDList []string `json:"minionClusterSecGroupIDList"`
	} `json:"data"`
}

func (s *GetResponse) ToClusterObject() *objects.Cluster {
	if s == nil {
		return nil
	}

	return &objects.Cluster{
		ID:                          s.Data.UUID,
		VpcID:                       s.Data.NetworkID,
		SubnetID:                    s.Data.SubnetID,
		MasterClusterSecGroupIDList: s.Data.MasterClusterSecGroupIDList,
		MinionClusterSecGroupIDList: s.Data.MinionClusterSecGroupIDList,
	}
}

// ****************************************** Update secgroup for the cluster ******************************************

type UpdateSecgroupResponse struct {
	Data []struct {
		UUID         string `json:"uuid"`
		ClusterID    string `json:"clusterId"`
		SecGroupID   string `json:"secGroupId"`
		ProjectID    string `json:"projectId"`
		Master       bool   `json:"master"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt,omitempty"`
		SecgroupName string `json:"secGroupName,omitempty"`
	} `json:"data"`
}

func (s *UpdateSecgroupResponse) ToListClusterSecgroupObjects() []*objects.ClusterSecgroup {
	if s == nil || s.Data == nil || len(s.Data) < 1 {
		return nil
	}

	var result []*objects.ClusterSecgroup
	for _, secgroup := range s.Data {
		result = append(result, &objects.ClusterSecgroup{
			UUID:         secgroup.UUID,
			ClusterID:    secgroup.ClusterID,
			SecGroupID:   secgroup.SecGroupID,
			Master:       secgroup.Master,
			SecgroupName: secgroup.SecgroupName,
		})
	}

	return result
}
