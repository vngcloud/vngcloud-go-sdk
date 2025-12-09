package v1

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func PointerOf[T any](t T) *T {
	return &t
}

type (
	GlobalPoolAlgorithm              string
	GlobalPoolProtocol               string
	GlobalPoolHealthCheckProtocol    string
	GlobalPoolHealthCheckMethod      string
	GlobalPoolHealthCheckHttpVersion string

	GlobalPoolMemberType string
)

const (
	GlobalPoolAlgorithmRoundRobin GlobalPoolAlgorithm = "ROUND_ROBIN"
	GlobalPoolAlgorithmLeastConn  GlobalPoolAlgorithm = "LEAST_CONNECTIONS"
	GlobalPoolAlgorithmSourceIP   GlobalPoolAlgorithm = "SOURCE_IP"
)

const (
	GlobalPoolProtocolTCP   GlobalPoolProtocol = "TCP"
	GlobalPoolProtocolProxy GlobalPoolProtocol = "PROXY"
)

const (
	GlobalPoolHealthCheckProtocolTCP   GlobalPoolHealthCheckProtocol = "TCP"
	GlobalPoolHealthCheckProtocolHTTP  GlobalPoolHealthCheckProtocol = "HTTP"
	GlobalPoolHealthCheckProtocolHTTPs GlobalPoolHealthCheckProtocol = "HTTPS"
)

const (
	GlobalPoolHealthCheckMethodGET  GlobalPoolHealthCheckMethod = "GET"
	GlobalPoolHealthCheckMethodPUT  GlobalPoolHealthCheckMethod = "PUT"
	GlobalPoolHealthCheckMethodPOST GlobalPoolHealthCheckMethod = "POST"
)

const (
	GlobalPoolHealthCheckHttpVersionHttp1       GlobalPoolHealthCheckHttpVersion = "1.0"
	GlobalPoolHealthCheckHttpVersionHttp1Minor1 GlobalPoolHealthCheckHttpVersion = "1.1"
)

const (
	GlobalPoolMemberTypePublic  GlobalPoolMemberType = "PUBLIC"
	GlobalPoolMemberTypePrivate GlobalPoolMemberType = "PRIVATE"
)

// --------------------------------------------------------------------------

var _ IListGlobalPoolsRequest = &ListGlobalPoolsRequest{}

type ListGlobalPoolsRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

func (s *ListGlobalPoolsRequest) WithLoadBalancerId(plbId string) IListGlobalPoolsRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *ListGlobalPoolsRequest) AddUserAgent(pagent ...string) IListGlobalPoolsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewListGlobalPoolsRequest(plbId string) IListGlobalPoolsRequest {
	opts := &ListGlobalPoolsRequest{}
	opts.LoadBalancerId = plbId
	return opts
}

// --------------------------------------------------------------------------

var _ ICreateGlobalPoolRequest = &CreateGlobalPoolRequest{}

type CreateGlobalPoolRequest struct {
	Algorithm         GlobalPoolAlgorithm              `json:"algorithm"`
	Description       string                           `json:"description,omitempty"`
	Name              string                           `json:"name"`
	Protocol          GlobalPoolProtocol               `json:"protocol"`
	Stickiness        *bool                            `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
	TLSEncryption     *bool                            `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
	HealthMonitor     IGlobalHealthMonitorRequest      `json:"health"`
	GlobalPoolMembers []ICreateGlobalPoolMemberRequest `json:"globalPoolMembers"`

	lscommon.LoadBalancerCommon
	lscommon.UserAgent
}

func (s *CreateGlobalPoolRequest) WithAlgorithm(palgorithm GlobalPoolAlgorithm) ICreateGlobalPoolRequest {
	s.Algorithm = palgorithm
	return s
}

func (s *CreateGlobalPoolRequest) WithDescription(pdescription string) ICreateGlobalPoolRequest {
	s.Description = pdescription
	return s
}

func (s *CreateGlobalPoolRequest) WithName(pname string) ICreateGlobalPoolRequest {
	s.Name = pname
	return s
}

func (s *CreateGlobalPoolRequest) WithProtocol(pprotocol GlobalPoolProtocol) ICreateGlobalPoolRequest {
	s.Protocol = pprotocol
	return s
}

func (s *CreateGlobalPoolRequest) WithHealthMonitor(phealth IGlobalHealthMonitorRequest) ICreateGlobalPoolRequest {
	s.HealthMonitor = phealth
	return s
}

func (s *CreateGlobalPoolRequest) WithMembers(pmembers ...ICreateGlobalPoolMemberRequest) ICreateGlobalPoolRequest {
	s.GlobalPoolMembers = append(s.GlobalPoolMembers, pmembers...)
	return s
}

func (s *CreateGlobalPoolRequest) WithLoadBalancerId(plbId string) ICreateGlobalPoolRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *CreateGlobalPoolRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"algorithm":         s.Algorithm,
		"description":       s.Description,
		"name":              s.Name,
		"protocol":          s.Protocol,
		"stickiness":        s.Stickiness,
		"tlsEncryption":     s.TLSEncryption,
		"health":            s.HealthMonitor.ToMap(),
		"globalPoolMembers": make([]map[string]interface{}, 0),
	}

	for _, member := range s.GlobalPoolMembers {
		err["globalPoolMembers"] = append(err["globalPoolMembers"].([]map[string]interface{}), member.ToMap())
	}

	return err
}

func (s *CreateGlobalPoolRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateGlobalPoolRequest) AddUserAgent(pagent ...string) ICreateGlobalPoolRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewCreateGlobalPoolRequest(pname string, pprotocol GlobalPoolProtocol) ICreateGlobalPoolRequest {
	opts := &CreateGlobalPoolRequest{
		Name:              pname,
		Protocol:          pprotocol,
		Algorithm:         GlobalPoolAlgorithmRoundRobin,
		Description:       "",
		Stickiness:        nil,
		TLSEncryption:     nil,
		HealthMonitor:     nil,
		GlobalPoolMembers: make([]ICreateGlobalPoolMemberRequest, 0),
	}

	return opts
}

// --------------------------------------------------------------------------

var _ IGlobalHealthMonitorRequest = &GlobalHealthMonitorRequest{}

type GlobalHealthMonitorRequest struct {
	HealthCheckProtocol GlobalPoolHealthCheckProtocol     `json:"protocol"`
	HealthyThreshold    int                               `json:"healthyThreshold"`
	UnhealthyThreshold  int                               `json:"unhealthyThreshold"`
	Interval            int                               `json:"interval"`
	Timeout             int                               `json:"timeout"`
	HttpMethod          *GlobalPoolHealthCheckMethod      `json:"httpMethod,omitempty"`
	HttpVersion         *GlobalPoolHealthCheckHttpVersion `json:"httpVersion,omitempty"`
	Path                *string                           `json:"path,omitempty"`
	DomainName          *string                           `json:"domainName,omitempty"`
	SuccessCode         *string                           `json:"successCode,omitempty"`

	lscommon.UserAgent
}

func (s *GlobalHealthMonitorRequest) WithHealthyThreshold(pthreshold int) IGlobalHealthMonitorRequest {
	s.HealthyThreshold = pthreshold
	return s
}

func (s *GlobalHealthMonitorRequest) WithUnhealthyThreshold(pthreshold int) IGlobalHealthMonitorRequest {
	s.UnhealthyThreshold = pthreshold
	return s
}

func (s *GlobalHealthMonitorRequest) WithProtocol(pprotocol GlobalPoolHealthCheckProtocol) IGlobalHealthMonitorRequest {
	s.HealthCheckProtocol = pprotocol
	return s
}

func (s *GlobalHealthMonitorRequest) WithInterval(pinterval int) IGlobalHealthMonitorRequest {
	s.Interval = pinterval
	return s
}

func (s *GlobalHealthMonitorRequest) WithTimeout(ptimeout int) IGlobalHealthMonitorRequest {
	s.Timeout = ptimeout
	return s
}

func (s *GlobalHealthMonitorRequest) WithHealthCheckMethod(pmethod *GlobalPoolHealthCheckMethod) IGlobalHealthMonitorRequest {
	s.HttpMethod = pmethod
	return s
}

func (s *GlobalHealthMonitorRequest) WithHttpVersion(pversion *GlobalPoolHealthCheckHttpVersion) IGlobalHealthMonitorRequest {
	s.HttpVersion = pversion
	return s
}

func (s *GlobalHealthMonitorRequest) WithDomainName(pdomain *string) IGlobalHealthMonitorRequest {
	s.DomainName = pdomain
	return s
}

func (s *GlobalHealthMonitorRequest) WithSuccessCode(pcode *string) IGlobalHealthMonitorRequest {
	s.SuccessCode = pcode
	return s
}

func (s *GlobalHealthMonitorRequest) WithPath(ppath *string) IGlobalHealthMonitorRequest {
	s.Path = ppath
	return s
}

func (s *GlobalHealthMonitorRequest) AddUserAgent(pagent ...string) IGlobalHealthMonitorRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *GlobalHealthMonitorRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"protocol":           s.HealthCheckProtocol,
		"healthyThreshold":   s.HealthyThreshold,
		"unhealthyThreshold": s.UnhealthyThreshold,
		"interval":           s.Interval,
		"timeout":            s.Timeout,
		"httpMethod":         s.HttpMethod,
		"httpVersion":        s.HttpVersion,
		"path":               s.Path,
		"domainName":         s.DomainName,
		"successCode":        s.SuccessCode,
	}

	return err
}

func (s *GlobalHealthMonitorRequest) ToRequestBody() interface{} {
	return s
}

func NewGlobalHealthMonitor(pcheckProtocol GlobalPoolHealthCheckProtocol) IGlobalHealthMonitorRequest {
	opts := &GlobalHealthMonitorRequest{
		HealthCheckProtocol: pcheckProtocol,
		HealthyThreshold:    3,
		UnhealthyThreshold:  3,
		Interval:            30,
		Timeout:             5,
		HttpMethod:          nil,
		HttpVersion:         nil,
		Path:                nil,
		DomainName:          nil,
		SuccessCode:         nil,
	}
	if pcheckProtocol == GlobalPoolHealthCheckProtocolHTTP || pcheckProtocol == GlobalPoolHealthCheckProtocolHTTPs {
		opts.HttpMethod = PointerOf(GlobalPoolHealthCheckMethodGET)
		opts.HttpVersion = PointerOf(GlobalPoolHealthCheckHttpVersionHttp1Minor1)
		opts.Path = PointerOf("/")
		opts.DomainName = PointerOf("")
		opts.SuccessCode = PointerOf("200")
	}
	return opts
}

// --------------------------------------------------------------------------

var _ ICreateGlobalPoolMemberRequest = &GlobalPoolMemberRequest{}

type GlobalPoolMemberRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Region      string                 `json:"region"`
	TrafficDial int                    `json:"trafficDial"`
	VPCID       string                 `json:"vpcId"` // only need for private type
	Type        GlobalPoolMemberType   `json:"type"`
	Members     []IGlobalMemberRequest `json:"members"`

	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.UserAgent
}

func (s *GlobalPoolMemberRequest) WithName(pname string) ICreateGlobalPoolMemberRequest {
	s.Name = pname
	return s
}

func (s *GlobalPoolMemberRequest) WithDescription(pdescription string) ICreateGlobalPoolMemberRequest {
	s.Description = pdescription
	return s
}

func (s *GlobalPoolMemberRequest) WithRegion(pregion string) ICreateGlobalPoolMemberRequest {
	s.Region = pregion
	return s
}

func (s *GlobalPoolMemberRequest) WithTrafficDial(pdial int) ICreateGlobalPoolMemberRequest {
	s.TrafficDial = pdial
	return s
}

func (s *GlobalPoolMemberRequest) WithVPCID(pvpcId string) ICreateGlobalPoolMemberRequest {
	s.VPCID = pvpcId
	return s
}

func (s *GlobalPoolMemberRequest) WithType(ptype GlobalPoolMemberType) ICreateGlobalPoolMemberRequest {
	s.Type = ptype
	return s
}

func (s *GlobalPoolMemberRequest) WithMembers(pmembers ...IGlobalMemberRequest) ICreateGlobalPoolMemberRequest {
	s.Members = append(s.Members, pmembers...)
	return s
}

func (s *GlobalPoolMemberRequest) WithLoadBalancerId(plbId string) ICreateGlobalPoolMemberRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *GlobalPoolMemberRequest) WithPoolId(ppoolId string) ICreateGlobalPoolMemberRequest {
	s.PoolId = ppoolId
	return s
}

func (s *GlobalPoolMemberRequest) AddUserAgent(pagent ...string) ICreateGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *GlobalPoolMemberRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"name":        s.Name,
		"description": s.Description,
		"region":      s.Region,
		"trafficDial": s.TrafficDial,
		"vpcId":       s.VPCID,
		"members":     make([]map[string]interface{}, 0),
	}

	for _, member := range s.Members {
		err["members"] = append(err["members"].([]map[string]interface{}), member.ToMap())
	}

	return err
}

func (s *GlobalPoolMemberRequest) ToRequestBody() interface{} {
	return s
}

func NewGlobalPoolMemberRequest(pname, pregion, pvpcId string, pdial int, pType GlobalPoolMemberType) ICreateGlobalPoolMemberRequest {
	opts := &GlobalPoolMemberRequest{
		Name:        pname,
		Description: "",
		Region:      pregion,
		TrafficDial: pdial,
		Type:        pType,
		VPCID:       pvpcId,
		Members:     make([]IGlobalMemberRequest, 0),
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IGlobalMemberRequest = &GlobalMemberRequest{}

type GlobalMemberRequest struct {
	Address     string `json:"address"`
	BackupRole  bool   `json:"backupRole"`
	Description string `json:"description,omitempty"`
	MonitorPort int    `json:"monitorPort"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
	SubnetID    string `json:"subnetId"`
	Weight      int    `json:"weight"`

	lscommon.UserAgent
}

func (s *GlobalMemberRequest) WithAddress(paddress string) IGlobalMemberRequest {
	s.Address = paddress
	return s
}

func (s *GlobalMemberRequest) WithBackupRole(pbackupRole bool) IGlobalMemberRequest {
	s.BackupRole = pbackupRole
	return s
}

func (s *GlobalMemberRequest) WithDescription(pdescription string) IGlobalMemberRequest {
	s.Description = pdescription
	return s
}

func (s *GlobalMemberRequest) WithMonitorPort(pmonitorPort int) IGlobalMemberRequest {
	s.MonitorPort = pmonitorPort
	return s
}

func (s *GlobalMemberRequest) WithName(pname string) IGlobalMemberRequest {
	s.Name = pname
	return s
}

func (s *GlobalMemberRequest) WithPort(pport int) IGlobalMemberRequest {
	s.Port = pport
	return s
}

func (s *GlobalMemberRequest) WithSubnetID(psubnetId string) IGlobalMemberRequest {
	s.SubnetID = psubnetId
	return s
}

func (s *GlobalMemberRequest) WithWeight(pweight int) IGlobalMemberRequest {
	s.Weight = pweight
	return s
}

func (s *GlobalMemberRequest) AddUserAgent(pagent ...string) IGlobalMemberRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *GlobalMemberRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"address":     s.Address,
		"backupRole":  s.BackupRole,
		"description": s.Description,
		"monitorPort": s.MonitorPort,
		"name":        s.Name,
		"port":        s.Port,
		"subnetId":    s.SubnetID,
		"weight":      s.Weight,
	}
	return err
}

func (s *GlobalMemberRequest) ToRequestBody() interface{} {
	return s
}

func NewGlobalMemberRequest(pname, paddress, psubnetId string, pport, pmonitorPort, pweight int, pbackupRole bool) IGlobalMemberRequest {
	opts := &GlobalMemberRequest{
		Name:        pname,
		Address:     paddress,
		BackupRole:  pbackupRole,
		Description: "",
		MonitorPort: pmonitorPort,
		Port:        pport,
		SubnetID:    psubnetId,
		Weight:      pweight,
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IUpdateGlobalPoolRequest = &UpdateGlobalPoolRequest{}

type UpdateGlobalPoolRequest struct {
	Algorithm     GlobalPoolAlgorithm         `json:"algorithm"`
	HealthMonitor IGlobalHealthMonitorRequest `json:"health"`
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.UserAgent
}

func (s *UpdateGlobalPoolRequest) WithAlgorithm(palgorithm GlobalPoolAlgorithm) IUpdateGlobalPoolRequest {
	s.Algorithm = palgorithm
	return s
}

func (s *UpdateGlobalPoolRequest) WithHealthMonitor(pmonitor IGlobalHealthMonitorRequest) IUpdateGlobalPoolRequest {
	s.HealthMonitor = pmonitor
	return s
}

func (s *UpdateGlobalPoolRequest) WithLoadBalancerId(plbId string) IUpdateGlobalPoolRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *UpdateGlobalPoolRequest) WithPoolId(ppoolId string) IUpdateGlobalPoolRequest {
	s.PoolId = ppoolId
	return s
}

func (s *UpdateGlobalPoolRequest) AddUserAgent(pagent ...string) IUpdateGlobalPoolRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *UpdateGlobalPoolRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"algorithm": s.Algorithm,
		"health":    s.HealthMonitor.ToMap(),
	}
	return err
}

func (s *UpdateGlobalPoolRequest) ToRequestBody() interface{} {
	return s
}

func NewUpdateGlobalPoolRequest(plbId, poolId string) IUpdateGlobalPoolRequest {
	opts := &UpdateGlobalPoolRequest{
		Algorithm:     GlobalPoolAlgorithmRoundRobin,
		HealthMonitor: nil,
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		PoolCommon: lscommon.PoolCommon{
			PoolId: poolId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IDeleteGlobalPoolRequest = &DeleteGlobalPoolRequest{}

type DeleteGlobalPoolRequest struct {
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.UserAgent
}

func (s *DeleteGlobalPoolRequest) WithLoadBalancerId(plbId string) IDeleteGlobalPoolRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *DeleteGlobalPoolRequest) WithPoolId(ppoolId string) IDeleteGlobalPoolRequest {
	s.PoolId = ppoolId
	return s
}

func (s *DeleteGlobalPoolRequest) AddUserAgent(pagent ...string) IDeleteGlobalPoolRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewDeleteGlobalPoolRequest(plbId, poolId string) IDeleteGlobalPoolRequest {
	opts := &DeleteGlobalPoolRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		PoolCommon: lscommon.PoolCommon{
			PoolId: poolId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IListGlobalPoolMembersRequest = &ListGlobalPoolMembersRequest{}

type ListGlobalPoolMembersRequest struct {
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.UserAgent
}

func (s *ListGlobalPoolMembersRequest) WithLoadBalancerId(plbId string) IListGlobalPoolMembersRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *ListGlobalPoolMembersRequest) WithPoolId(ppoolId string) IListGlobalPoolMembersRequest {
	s.PoolId = ppoolId
	return s
}

func (s *ListGlobalPoolMembersRequest) AddUserAgent(pagent ...string) IListGlobalPoolMembersRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewListGlobalPoolMembersRequest(plbId, poolId string) IListGlobalPoolMembersRequest {
	opts := &ListGlobalPoolMembersRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		PoolCommon: lscommon.PoolCommon{
			PoolId: poolId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IPatchGlobalPoolMembersRequest = &PatchGlobalPoolMembersRequest{}

type PatchGlobalPoolMembersRequest struct {
	BulkActions []IBulkActionRequest `json:"bulkActions"`
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.UserAgent
}

func (s *PatchGlobalPoolMembersRequest) WithBulkAction(paction ...IBulkActionRequest) IPatchGlobalPoolMembersRequest {
	s.BulkActions = paction
	return s
}

func (s *PatchGlobalPoolMembersRequest) WithLoadBalancerId(plbId string) IPatchGlobalPoolMembersRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *PatchGlobalPoolMembersRequest) WithPoolId(ppoolId string) IPatchGlobalPoolMembersRequest {
	s.PoolId = ppoolId
	return s
}

func (s *PatchGlobalPoolMembersRequest) AddUserAgent(pagent ...string) IPatchGlobalPoolMembersRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *PatchGlobalPoolMembersRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"bulkActions": make([]map[string]interface{}, 0),
	}

	for _, action := range s.BulkActions {
		err["bulkActions"] = append(err["bulkActions"].([]map[string]interface{}), action.ToMap())
	}

	return err
}

func (s *PatchGlobalPoolMembersRequest) ToRequestBody() interface{} {
	return s
}

func NewPatchGlobalPoolMembersRequest(plbId, poolId string) IPatchGlobalPoolMembersRequest {
	opts := &PatchGlobalPoolMembersRequest{
		BulkActions: make([]IBulkActionRequest, 0),
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		PoolCommon: lscommon.PoolCommon{
			PoolId: poolId,
		},
	}
	return opts
}

var _ IBulkActionRequest = &PatchGlobalPoolCreateBulkActionRequest{}

type PatchGlobalPoolCreateBulkActionRequest struct {
	Action           string                         `json:"action"`
	CreatePoolMember ICreateGlobalPoolMemberRequest `json:"createPoolMember"`
}

func (s *PatchGlobalPoolCreateBulkActionRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"action":           s.Action,
		"createPoolMember": s.CreatePoolMember.ToMap(),
	}
	return err
}

func (s *PatchGlobalPoolCreateBulkActionRequest) ToRequestBody() interface{} {
	return s
}

func NewPatchGlobalPoolCreateBulkActionRequest(pmember ICreateGlobalPoolMemberRequest) IBulkActionRequest {
	opts := &PatchGlobalPoolCreateBulkActionRequest{
		Action:           "create",
		CreatePoolMember: pmember,
	}
	return opts
}

var _ IBulkActionRequest = &PatchGlobalPoolDeleteBulkActionRequest{}

type PatchGlobalPoolDeleteBulkActionRequest struct {
	Action string `json:"action"`
	ID     string `json:"id"`
}

func (s *PatchGlobalPoolDeleteBulkActionRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"action": s.Action,
		"id":     s.ID,
	}
	return err
}

func (s *PatchGlobalPoolDeleteBulkActionRequest) ToRequestBody() interface{} {
	return s
}

func NewPatchGlobalPoolDeleteBulkActionRequest(pid string) IBulkActionRequest {
	opts := &PatchGlobalPoolDeleteBulkActionRequest{
		Action: "delete",
		ID:     pid,
	}
	return opts
}

var _ IBulkActionRequest = &PatchGlobalPoolUpdateBulkActionRequest{}

type PatchGlobalPoolUpdateBulkActionRequest struct {
	Action           string                         `json:"action"`
	ID               string                         `json:"id"`
	UpdatePoolMember IUpdateGlobalPoolMemberRequest `json:"updatePoolMember"`
}

func (s *PatchGlobalPoolUpdateBulkActionRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"action":           s.Action,
		"id":               s.ID,
		"updatePoolMember": s.UpdatePoolMember.ToMap(),
	}
	return err
}

func (s *PatchGlobalPoolUpdateBulkActionRequest) ToRequestBody() interface{} {
	return s
}

func NewPatchGlobalPoolUpdateBulkActionRequest(pid string, pmember IUpdateGlobalPoolMemberRequest) IBulkActionRequest {
	opts := &PatchGlobalPoolUpdateBulkActionRequest{
		Action:           "update",
		ID:               pid,
		UpdatePoolMember: pmember,
	}
	return opts
}

type IUpdateGlobalPoolMemberRequest interface {
	WithLoadBalancerId(plbId string) IUpdateGlobalPoolMemberRequest
	WithPoolId(ppoolId string) IUpdateGlobalPoolMemberRequest
	WithPoolMemberId(ppoolMemberId string) IUpdateGlobalPoolMemberRequest
	WithTrafficDial(pdial int) IUpdateGlobalPoolMemberRequest
	WithMembers(pmembers ...IGlobalMemberRequest) IUpdateGlobalPoolMemberRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string
	GetPoolMemberId() string

	AddUserAgent(pagent ...string) IUpdateGlobalPoolMemberRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

var _ IUpdateGlobalPoolMemberRequest = &UpdateGlobalPoolMemberRequest{}

type UpdateGlobalPoolMemberRequest struct {
	TrafficDial int                    `json:"trafficDial"`
	Members     []IGlobalMemberRequest `json:"members"`

	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.PoolMemberCommon
}

func (s *UpdateGlobalPoolMemberRequest) WithTrafficDial(pdial int) IUpdateGlobalPoolMemberRequest {
	s.TrafficDial = pdial
	return s
}

func (s *UpdateGlobalPoolMemberRequest) WithMembers(pmembers ...IGlobalMemberRequest) IUpdateGlobalPoolMemberRequest {
	s.Members = append(s.Members, pmembers...)
	return s
}

func (s *UpdateGlobalPoolMemberRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"trafficDial": s.TrafficDial,
		"members":     make([]map[string]interface{}, 0),
	}

	for _, member := range s.Members {
		err["members"] = append(err["members"].([]map[string]interface{}), member.ToMap())
	}

	return err
}

func (s *UpdateGlobalPoolMemberRequest) WithLoadBalancerId(plbId string) IUpdateGlobalPoolMemberRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *UpdateGlobalPoolMemberRequest) WithPoolId(ppoolId string) IUpdateGlobalPoolMemberRequest {
	s.PoolId = ppoolId
	return s
}

func (s *UpdateGlobalPoolMemberRequest) WithPoolMemberId(ppoolMemberId string) IUpdateGlobalPoolMemberRequest {
	s.PoolMemberId = ppoolMemberId
	return s
}

func (s *UpdateGlobalPoolMemberRequest) AddUserAgent(pagent ...string) IUpdateGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *UpdateGlobalPoolMemberRequest) ToRequestBody() interface{} {
	return s
}

func NewUpdateGlobalPoolMemberRequest(plbId, poolId, poolMemberId string, pdial int) IUpdateGlobalPoolMemberRequest {
	opts := &UpdateGlobalPoolMemberRequest{
		TrafficDial: pdial,
		Members:     make([]IGlobalMemberRequest, 0),
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		PoolCommon: lscommon.PoolCommon{
			PoolId: poolId,
		},
		PoolMemberCommon: lscommon.PoolMemberCommon{
			PoolMemberId: poolMemberId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IDeleteGlobalPoolMemberRequest = &DeleteGlobalPoolMemberRequest{}

type DeleteGlobalPoolMemberRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.PoolMemberCommon
}

func (s *DeleteGlobalPoolMemberRequest) WithLoadBalancerId(plbId string) IDeleteGlobalPoolMemberRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *DeleteGlobalPoolMemberRequest) WithPoolId(ppoolId string) IDeleteGlobalPoolMemberRequest {
	s.PoolId = ppoolId
	return s
}

func (s *DeleteGlobalPoolMemberRequest) WithPoolMemberId(ppoolMemberId string) IDeleteGlobalPoolMemberRequest {
	s.PoolMemberId = ppoolMemberId
	return s
}

func (s *DeleteGlobalPoolMemberRequest) AddUserAgent(pagent ...string) IDeleteGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewDeleteGlobalPoolMemberRequest(plbId, poolId, poolMemberId string) IDeleteGlobalPoolMemberRequest {
	opts := &DeleteGlobalPoolMemberRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		PoolCommon: lscommon.PoolCommon{
			PoolId: poolId,
		},
		PoolMemberCommon: lscommon.PoolMemberCommon{
			PoolMemberId: poolMemberId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IGetGlobalPoolMemberRequest = &GetGlobalPoolMemberRequest{}

type GetGlobalPoolMemberRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.PoolMemberCommon
}

func (s *GetGlobalPoolMemberRequest) WithLoadBalancerId(plbId string) IGetGlobalPoolMemberRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *GetGlobalPoolMemberRequest) WithPoolId(ppoolId string) IGetGlobalPoolMemberRequest {
	s.PoolId = ppoolId
	return s
}

func (s *GetGlobalPoolMemberRequest) WithPoolMemberId(ppoolMemberId string) IGetGlobalPoolMemberRequest {
	s.PoolMemberId = ppoolMemberId
	return s
}

func (s *GetGlobalPoolMemberRequest) AddUserAgent(pagent ...string) IGetGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewGetGlobalPoolMemberRequest(plbId, poolId, poolMemberId string) IGetGlobalPoolMemberRequest {
	opts := &GetGlobalPoolMemberRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{
			LoadBalancerId: plbId,
		},
		PoolCommon: lscommon.PoolCommon{
			PoolId: poolId,
		},
		PoolMemberCommon: lscommon.PoolMemberCommon{
			PoolMemberId: poolMemberId,
		},
	}
	return opts
}
