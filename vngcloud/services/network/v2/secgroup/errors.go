package secgroup

import (
	"fmt"
	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
)

const (
	patternErrNotFound   = "Cannot get security group with id secg-"
	patternSecgroupInUse = "SecurityGroupInUse"
)

var (
	ErrSecgroupNotFound vconError.ErrorCode = "SECURITY_GROUP_NOT_FOUND"
	ErrSecgroupUnknown  vconError.ErrorCode = "SECURITY_GROUP_UNKNOWN_ERROR"
	ErrSecgroupInUse    vconError.ErrorCode = "SECURITY_GROUP_IN_USE"
)

func NewErrNameDuplicate(pName, pInfo string) vconError.IErrorBuilder {
	err := new(ErrNameDuplicate)
	err.Name = pName
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrNameDuplicate(pErr error) bool {
	_, ok := pErr.(*ErrNameDuplicate)
	return ok
}

type ErrNameDuplicate struct {
	Name string
	vconError.BaseError
}

func (s *ErrNameDuplicate) Error() string {
	s.DefaultError = fmt.Sprintf("the secgroup [%s] is duplicated by name", s.Name)
	return s.ChoseErrString()
}
