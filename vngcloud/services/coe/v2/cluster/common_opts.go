package cluster

type ClusterCommonOpts struct {
	ClusterID string
}

func (s *ClusterCommonOpts) GetClusterID() string {
	return s.ClusterID
}
