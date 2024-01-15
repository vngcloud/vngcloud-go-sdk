package error

// ***************************************************** BaseError *****************************************************

type BaseError struct {
	DefaultError string
	Info         string
}

func (s *BaseError) Error() string {
	s.DefaultError = defaultErrorString
	return s.ChoseErrString()
}

func (s *BaseError) SetInfo(pInfo string) IErrorBuilder {
	s.Info = pInfo
	return s
}

func (s *BaseError) ChoseErrString() string {
	if s.Info != "" {
		return s.Info
	}
	return s.DefaultError
}
