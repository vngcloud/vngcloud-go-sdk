package compute

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lscomputeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute/v2"
)

type IComputeServiceV2 interface {
	CreateServer(popts lscomputeSvcV2.ICreateServerRequest) (*lsentity.Server, lserr.IError)
	GetServerById(popts lscomputeSvcV2.IGetServerByIdRequest) (*lsentity.Server, lserr.IError)
	DeleteServerById(popts lscomputeSvcV2.IDeleteServerByIdRequest) lserr.IError
	UpdateServerSecgroupsByServerId(popts lscomputeSvcV2.IUpdateServerSecgroupsByServerIdRequest) (*lsentity.Server, lserr.IError)
	AttachBlockVolume(popts lscomputeSvcV2.IAttachBlockVolumeRequest) lserr.IError
	DetachBlockVolume(popts lscomputeSvcV2.IDetachBlockVolumeRequest) lserr.IError
	AttachFloatingIp(popts lscomputeSvcV2.IAttachFloatingIpRequest) lserr.IError
	DetachFloatingIp(popts lscomputeSvcV2.IDetachFloatingIpRequest) lserr.IError
	ListServerGroupPolicies(popts lscomputeSvcV2.IListServerGroupPoliciesRequest) (*lsentity.ListServerGroupPolicies, lserr.IError)
	DeleteServerGroupById(popts lscomputeSvcV2.IDeleteServerGroupByIdRequest) lserr.IError
	ListServerGroups(popts lscomputeSvcV2.IListServerGroupsRequest) (*lsentity.ListServerGroups, lserr.IError)
	CreateServerGroup(popts lscomputeSvcV2.ICreateServerGroupRequest) (*lsentity.ServerGroup, lserr.IError)
}
