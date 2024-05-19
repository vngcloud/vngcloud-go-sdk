package v2

func NewCreateServerRequest(pname, pimageId, pflavorId, pnetworkId, psubnetId, prootDiskType string, prootDiskSize int) ICreateServerRequest {
	opt := new(CreateServerRequest)
	opt.Name = pname
	opt.ImageId = pimageId
	opt.FlavorId = pflavorId
	opt.NetworkId = pnetworkId
	opt.SubnetId = psubnetId
	opt.RootDiskTypeId = prootDiskType
	opt.RootDiskSize = prootDiskSize
	return opt
}

func NewGetServerByIdRequest(pserverId string) IGetServerByIdRequest {
	opt := new(GetServerByIdRequest)
	opt.ServerId = pserverId
	return opt
}

func NewDeleteServerByIdRequest(pserverId string) IDeleteServerByIdRequest {
	opt := new(DeleteServerByIdRequest)
	opt.ServerId = pserverId
	opt.DeleteAllVolume = false
	return opt
}

func NewUpdateServerSecgroupsRequest(pserverId string, psecgroups ...string) IUpdateServerSecgroupsByServerIdRequest {
	opt := new(UpdateServerSecgroupsByServerIdRequest)
	opt.ServerId = pserverId
	opt.Secgroups = psecgroups
	return opt
}

const (
	DataDiskEncryptionAesXtsType DataDiskEncryptionType = "aes-xts-plain64_256"
)

type CreateServerRequest struct {
	AttachFloating         bool                   `json:"attachFloating,omitempty"`
	BackupInstancePointId  string                 `json:"backupInstancePointId,omitempty"`
	DataDiskEncryptionType DataDiskEncryptionType `json:"dataDiskEncryptionType,omitempty"`
	DataDiskName           string                 `json:"dataDiskName,omitempty"`
	DataDiskSize           int                    `json:"dataDiskSize,omitempty"`
	DataDiskTypeId         string                 `json:"dataDiskTypeId,omitempty"`
	EnableBackup           bool                   `json:"enableBackup,omitempty"`
	EncryptionVolume       bool                   `json:"encryptionVolume"`
	ExpirePassword         bool                   `json:"expirePassword,omitempty"`
	FlavorId               string                 `json:"flavorId"`
	ImageId                string                 `json:"imageId"`
	Name                   string                 `json:"name"`
	NetworkId              string                 `json:"networkId"`
	OsLicence              bool                   `json:"osLicence,omitempty"`
	RestoreBackup          bool                   `json:"restoreBackup,omitempty"`
	RootDiskEncryptionType DataDiskEncryptionType `json:"rootDiskEncryptionType,omitempty"`
	RootDiskSize           int                    `json:"rootDiskSize"`
	RootDiskTypeId         string                 `json:"rootDiskTypeId"`
	SecurityGroup          []string               `json:"securityGroup,omitempty"`
	ServerGroupId          string                 `json:"serverGroupId,omitempty"`
	SshKeyId               string                 `json:"sshKeyId,omitempty"`
	SubnetId               string                 `json:"subnetId"`
	UserData               string                 `json:"userData,omitempty"`
	UserDataBase64Encoded  bool                   `json:"userDataBase64Encoded,omitempty"`
	UserName               string                 `json:"userName,omitempty"`
	UserPassword           string                 `json:"userPassword,omitempty"`
	IsPoc                  bool                   `json:"isPoc,omitempty"`
	Product                string                 `json:"product,omitempty"`
	Type                   string                 `json:"type,omitempty"`
	Tags                   []ServerTag            `json:"tags,omitempty"`
}

type DataDiskEncryptionType string

type ServerTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *CreateServerRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateServerRequest) WithEncryptionVolume(pencryptionVolume bool) ICreateServerRequest {
	s.EncryptionVolume = pencryptionVolume
	return s
}

func (s *CreateServerRequest) WithUserData(puserData string, pbase64Encode bool) ICreateServerRequest {
	s.UserData = puserData
	s.UserDataBase64Encoded = pbase64Encode
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

func (s *CreateServerRequest) WithPoct(pisPoc bool) ICreateServerRequest {
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

type GetServerByIdRequest struct {
	ServerCommon
}

type DeleteServerByIdRequest struct {
	DeleteAllVolume bool `json:"deleteAllVolume"`
	ServerCommon
}

func (s *DeleteServerByIdRequest) WithDeleteAllVolume(pok bool) IDeleteServerByIdRequest {
	s.DeleteAllVolume = pok
	return s
}

func (s *DeleteServerByIdRequest) ToRequestBody() interface{} {
	return s
}

type UpdateServerSecgroupsByServerIdRequest struct {
	Secgroups []string `json:"securityGroup"`

	ServerCommon
}

func (s *UpdateServerSecgroupsByServerIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateServerSecgroupsByServerIdRequest) GetListSecgroupsIds() []string {
	return s.Secgroups
}
