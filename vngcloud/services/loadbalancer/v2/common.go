package v2

type LoadBalancerV2Common struct {
	LoadBalancerID string
}

func (s *LoadBalancerV2Common) GetLoadBalancerID() string {
	return s.LoadBalancerID
}

type PoolV2Common struct {
	PoolID string
}

func (s *PoolV2Common) GetPoolID() string {
	return s.PoolID
}

type ListenerV2Common struct {
	ListenerID string
}

func (s *ListenerV2Common) GetListenerID() string {
	return s.ListenerID
}
