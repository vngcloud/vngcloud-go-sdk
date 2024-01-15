package secgroup

import (
	"fmt"
	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
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
