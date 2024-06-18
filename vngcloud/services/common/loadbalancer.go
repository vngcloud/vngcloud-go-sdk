package common

type LoadBalancerCommon struct {
	LoadBalancerId string
}

func (s *LoadBalancerCommon) GetLoadBalancerId() string {
	return s.LoadBalancerId
}

type ListenerCommon struct {
	ListenerId string
}

func (s *ListenerCommon) GetListenerId() string {
	return s.ListenerId
}

type PoolCommon struct {
	PoolId string
}

func (s *PoolCommon) GetPoolId() string {
	return s.PoolId
}
