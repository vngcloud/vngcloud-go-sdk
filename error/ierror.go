package error

type IErrorBuilder interface {
	SetInfo(pInfo string) IErrorBuilder
	Error() string
}
