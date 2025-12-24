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

type PolicyCommon struct {
	PolicyId string
}

func (s *PolicyCommon) GetPolicyId() string {
	return s.PolicyId
}

type PoolMemberCommon struct {
	PoolMemberId string
}

func (s *PoolMemberCommon) GetPoolMemberId() string {
	return s.PoolMemberId
}
