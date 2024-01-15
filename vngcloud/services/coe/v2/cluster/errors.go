package cluster

import (
	"fmt"
	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
)

func NewClusterNotFound(pClusterID, pInfo string) vconError.IErrorBuilder {
	err := new(ErrClusterNotFound)
	err.ClusterID = pClusterID
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrClusterNotFound(pErr error) bool {
	_, ok := pErr.(*ErrClusterNotFound)
	return ok
}

type ErrClusterNotFound struct {
	ClusterID string
	vconError.BaseError
}

func (s *ErrClusterNotFound) Error() string {
	s.DefaultError = fmt.Sprintf("cluster %s not found", s.ClusterID)
	return s.ChoseErrString()
}
