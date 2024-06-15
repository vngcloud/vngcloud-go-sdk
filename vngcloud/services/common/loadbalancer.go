package common

type LoadBalancerCommon struct {
	LoadBalancerId string
}

func (s *LoadBalancerCommon) GetLoadBalancerId() string {
	return s.LoadBalancerId
}
