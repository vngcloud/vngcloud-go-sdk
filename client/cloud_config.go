package client

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//                                                      STRUCTS                                                       //
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Clouds struct {
	Clouds map[string]Cloud `yaml:"clouds" json:"clouds"`
}

type Cloud struct {
	VServerGatewayURL       string `yaml:"vserver_gateway_url,omitempty" json:"vserver_gateway_url,omitempty"`
	VLoadBalancerGatewayURL string `yaml:"vloadbalancer_gateway_url,omitempty" json:"vloadbalancer_gateway_url,omitempty"`
	IamGatewayURL           string `yaml:"iam_gateway_url,omitempty" json:"iam_gateway_url,omitempty"`
	ClienID                 string `yaml:"client_id,omitempty" json:"client_id,omitempty"`
	ClientSecret            string `yaml:"client_secret,omitempty" json:"client_secret,omitempty"`
}
