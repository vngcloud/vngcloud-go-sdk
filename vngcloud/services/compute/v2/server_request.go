package v2

import (
	lfmt "fmt"

	ljparser "github.com/cuongpiger/joat/parser"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

const (
	DataDiskEncryptionAesXts256Type DataDiskEncryptionType = "aes-xts-plain64_256"
	DataDiskEncryptionAesXts128Type DataDiskEncryptionType = "aes-xts-plain64_128"
)

type CreateServerRequest struct {
	AttachFloating         bool                     `json:"attachFloating,omitempty"`
	BackupInstancePointId  string                   `json:"backupInstancePointId,omitempty"`
	DataDiskEncryptionType DataDiskEncryptionType   `json:"dataDiskEncryptionType,omitempty"`
	DataDiskName           string                   `json:"dataDiskName,omitempty"`
	DataDiskSize           int                      `json:"dataDiskSize,omitempty"`
	DataDiskTypeId         string                   `json:"dataDiskTypeId,omitempty"`
	EnableBackup           bool                     `json:"enableBackup,omitempty"`
	EncryptionVolume       bool                     `json:"encryptionVolume"`
	ExpirePassword         bool                     `json:"expirePassword,omitempty"`
	FlavorId               string                   `json:"flavorId"`
	ImageId                string                   `json:"imageId"`
	Name                   string                   `json:"name"`
	NetworkId              string                   `json:"networkId,omitempty"`
	SubnetId               string                   `json:"subnetId,omitempty"`
	OsLicence              bool                     `json:"osLicence,omitempty"`
	RestoreBackup          bool                     `json:"restoreBackup,omitempty"`
	RootDiskEncryptionType DataDiskEncryptionType   `json:"rootDiskEncryptionType,omitempty"`
	RootDiskSize           int                      `json:"rootDiskSize"`
	RootDiskTypeId         string                   `json:"rootDiskTypeId"`
	SecurityGroup          []string                 `json:"securityGroup,omitempty"`
	ServerGroupId          string                   `json:"serverGroupId,omitempty"`
	SshKeyId               string                   `json:"sshKeyId,omitempty"`
	UserData               string                   `json:"userData,omitempty"`
	UserDataBase64Encoded  bool                     `json:"userDataBase64Encoded,omitempty"`
	UserName               string                   `json:"userName,omitempty"`
	UserPassword           string                   `json:"userPassword,omitempty"`
	IsPoc                  bool                     `json:"isPoc,omitempty"`
	Product                string                   `json:"product,omitempty"`
	Type                   string                   `json:"type,omitempty"`
	Tags                   []ServerTag              `json:"tags,omitempty"`
	AutoRenew              bool                     `json:"isEnableAutoRenew,omitempty"`
	Networks               []ServerNetworkInterface `json:"networks,omitempty"`
	Zone                   string                   `json:"zoneId,omitempty"`
	lscommon.UserAgent
}

type ServerNetworkInterface struct {
	ProjectId      string `json:"projectId"`
	NetworkId      string `json:"networkId"`
	SubnetId       string `json:"subnetId"`
	AttachFloating bool   `json:"attachFloating"`
}

type AttachBlockVolumeRequest struct {
	lscommon.BlockVolumeCommon
	lscommon.ServerCommon
}

type DetachBlockVolumeRequest struct {
	lscommon.BlockVolumeCommon
	lscommon.ServerCommon
}

type DataDiskEncryptionType string

type ServerTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *CreateServerRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateServerRequest) WithZone(zoneId string) ICreateServerRequest {
	s.Zone = zoneId
	return s
}

func (s *CreateServerRequest) WithServerNetworkInterface(pprojectId, pnetworkId, psubnetId string, pattachFloating bool) ICreateServerRequest {
	s.Networks = append(s.Networks, ServerNetworkInterface{
		ProjectId:      pprojectId,
		NetworkId:      pnetworkId,
		SubnetId:       psubnetId,
		AttachFloating: pattachFloating,
	})

	return s.WithNetwork(s.Networks[0].NetworkId, s.Networks[0].SubnetId)
}

func (s *CreateServerRequest) WithRootDiskEncryptionType(pdataDisk DataDiskEncryptionType) ICreateServerRequest {
	s.EncryptionVolume = true
	s.RootDiskEncryptionType = pdataDisk
	return s
}

func (s *CreateServerRequest) WithEncryptionVolume(penabled bool) ICreateServerRequest {
	s.EncryptionVolume = penabled
	return s
}

func (s *CreateServerRequest) WithUserData(puserData string, pbase64Encode bool) ICreateServerRequest {
	s.UserData = puserData
	s.UserDataBase64Encoded = pbase64Encode
	return s
}

func (s *CreateServerRequest) WithAutoRenew(pval bool) ICreateServerRequest {
	s.AutoRenew = pval
	return s
}

func (s *CreateServerRequest) WithTags(ptags ...string) ICreateServerRequest {
	if s.Tags == nil {
		s.Tags = make([]ServerTag, 0)
	}

	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.Tags = append(s.Tags, ServerTag{Key: ptags[i], Value: ptags[i+1]})
	}
	return s
}

func (s *CreateServerRequest) WithAttachFloating(pattachFloating bool) ICreateServerRequest {
	s.AttachFloating = pattachFloating
	return s
}

func (s *CreateServerRequest) WithSecgroups(psecgroups ...string) ICreateServerRequest {
	s.SecurityGroup = append(s.SecurityGroup, psecgroups...)
	return s
}

func (s *CreateServerRequest) WithServerGroupId(pserverGroupId string) ICreateServerRequest {
	s.ServerGroupId = pserverGroupId
	return s
}

func (s *CreateServerRequest) WithPoc(pisPoc bool) ICreateServerRequest {
	s.IsPoc = pisPoc
	return s
}

func (s *CreateServerRequest) WithType(ptype string) ICreateServerRequest {
	s.Type = ptype
	return s
}

func (s *CreateServerRequest) WithProduct(pproduct string) ICreateServerRequest {
	s.Product = pproduct
	return s
}

func (s *CreateServerRequest) WithNetwork(pnetworkId, psubnetId string) ICreateServerRequest {
	s.NetworkId = pnetworkId
	s.SubnetId = psubnetId

	return s
}

func (s *CreateServerRequest) AddUserAgent(pagent ...string) ICreateServerRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreateServerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"attachFloating":         s.AttachFloating,
		"backupInstancePointId":  s.BackupInstancePointId,
		"dataDiskEncryptionType": s.DataDiskEncryptionType,
		"dataDiskName":           s.DataDiskName,
		"dataDiskSize":           s.DataDiskSize,
		"dataDiskTypeId":         s.DataDiskTypeId,
		"enableBackup":           s.EnableBackup,
		"encryptionVolume":       s.EncryptionVolume,
		"expirePassword":         s.ExpirePassword,
		"flavorId":               s.FlavorId,
		"imageId":                s.ImageId,
		"name":                   s.Name,
		"networkId":              s.NetworkId,
		"subnetId":               s.SubnetId,
		"osLicence":              s.OsLicence,
		"restoreBackup":          s.RestoreBackup,
		"rootDiskEncryptionType": s.RootDiskEncryptionType,
		"rootDiskSize":           s.RootDiskSize,
		"rootDiskTypeId":         s.RootDiskTypeId,
		"securityGroup":          s.SecurityGroup,
		"serverGroupId":          s.ServerGroupId,
		"sshKeyId":               s.SshKeyId,
		"userName":               s.UserName,
		"isPoc":                  s.IsPoc,
		"product":                s.Product,
		"type":                   s.Type,
		"tags":                   s.Tags,
		"autoRenew":              s.AutoRenew,
		"networks":               s.Networks,
	}
}

type GetServerByIdRequest struct {
	lscommon.ServerCommon
	lscommon.UserAgent
}

func (s *GetServerByIdRequest) AddUserAgent(pagent ...string) IGetServerByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *GetServerByIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverId": s.ServerId,
	}
}

type DeleteServerByIdRequest struct {
	DeleteAllVolume bool `json:"deleteAllVolume"`
	lscommon.ServerCommon
	lscommon.UserAgent
}

func (s *DeleteServerByIdRequest) WithDeleteAllVolume(pok bool) IDeleteServerByIdRequest {
	s.DeleteAllVolume = pok
	return s
}

func (s *DeleteServerByIdRequest) AddUserAgent(pagent ...string) IDeleteServerByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *DeleteServerByIdRequest) ToRequestBody() interface{} {
	return s
}

type UpdateServerSecgroupsByServerIdRequest struct {
	Secgroups []string `json:"securityGroup"`

	lscommon.ServerCommon
	lscommon.UserAgent
}

func (s *UpdateServerSecgroupsByServerIdRequest) AddUserAgent(pagent ...string) IUpdateServerSecgroupsByServerIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *UpdateServerSecgroupsByServerIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateServerSecgroupsByServerIdRequest) GetListSecgroupsIds() []string {
	return s.Secgroups
}

type AttachFloatingIpRequest struct {
	NetworkInterfaceId string `json:"networkInterfaceId"`

	lscommon.InternalNetworkInterfaceCommon
	lscommon.ServerCommon
	lscommon.UserAgent
}

func (s *AttachFloatingIpRequest) ToRequestBody() interface{} {
	return s
}

func (s *AttachFloatingIpRequest) AddUserAgent(pagent ...string) IAttachFloatingIpRequest {
	s.Agent = append(s.Agent, pagent...)
	return s
}

func (s *AttachFloatingIpRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverId":                   s.ServerId,
		"internalNetworkInterfaceId": s.InternalNetworkInterfaceId,
		"networkId":                  s.NetworkInterfaceId,
	}
}

type DetachFloatingIpRequest struct {
	NetworkInterfaceId string `json:"networkInterfaceId"`

	lscommon.ServerCommon
	lscommon.WanCommon
	lscommon.InternalNetworkInterfaceCommon
	lscommon.UserAgent
}

func (s *DetachFloatingIpRequest) ToRequestBody() interface{} {
	return s
}

func (s *DetachFloatingIpRequest) AddUserAgent(pagent ...string) IDetachFloatingIpRequest {
	s.Agent = append(s.Agent, pagent...)
	return s
}

func (s *DetachFloatingIpRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverId":                   s.ServerId,
		"internalNetworkInterfaceId": s.InternalNetworkInterfaceId,
		"networkId":                  s.NetworkInterfaceId,
		"wanId":                      s.WanId,
	}
}

type ListServerGroupPoliciesRequest struct {
	lscommon.UserAgent
}

func (s *ListServerGroupPoliciesRequest) AddUserAgent(pagent ...string) IListServerGroupPoliciesRequest {
	s.Agent = append(s.Agent, pagent...)
	return s
}

type DeleteServerGroupByIdRequest struct {
	lscommon.ServerGroupCommon
	lscommon.UserAgent
}

func (s *DeleteServerGroupByIdRequest) AddUserAgent(pagent ...string) IDeleteServerGroupByIdRequest {
	s.Agent = append(s.Agent, pagent...)
	return s
}

func (s *DeleteServerGroupByIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverGroupId": s.ServerGroupId,
	}
}

type ListServerGroupsRequest struct {
	Name string `q:"name,beempty"`
	Page int    `q:"page"`
	Size int    `q:"size"`

	lscommon.UserAgent
}

func (s *ListServerGroupsRequest) WithName(pname string) IListServerGroupsRequest {
	s.Name = pname
	return s
}

func (s *ListServerGroupsRequest) ToListQuery() (string, error) {
	parser, _ := ljparser.GetParser()
	url, err := parser.UrlMe(s)
	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (s *ListServerGroupsRequest) GetDefaultQuery() string {
	return lfmt.Sprintf("offset=%d&limit=%d&name=", defaultOffsetListServerGroups, defaultLimitListServerGroups)
}

func (s *ListServerGroupsRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"page": s.Page,
		"size": s.Size,
		"name": s.Name,
	}
}

func (s *ListServerGroupsRequest) AddUserAgent(pagent ...string) IListServerGroupsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type CreateServerGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	PolicyId    string `json:"policyId,omitempty"`

	lscommon.UserAgent
}

func (s *CreateServerGroupRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateServerGroupRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":        s.Name,
		"description": s.Description,
		"policyId":    s.PolicyId,
	}
}

func (s *CreateServerGroupRequest) AddUserAgent(pagent ...string) ICreateServerGroupRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
