package inter

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
	WithProjectId(pprojectId string) ICreateLoadBalancerRequest
}
