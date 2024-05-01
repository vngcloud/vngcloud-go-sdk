package volume_attach

import vconError "github.com/vngcloud/vngcloud-go-sdk/error"

func NewErrAttachNotFound(pInfo string) vconError.IErrorBuilder {
	err := new(ErrAttachNotFound)
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrAttachNotFound(pErr error) bool {
	_, ok := pErr.(*ErrAttachNotFound)
	return ok
}

type ErrAttachNotFound struct {
	vconError.BaseError
}

func (s *ErrAttachNotFound) Error() string {
	s.DefaultError = "request options SHOULD NOT be nil"
	return s.ChoseErrString()
}
