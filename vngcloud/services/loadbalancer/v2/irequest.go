package v2

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
}
