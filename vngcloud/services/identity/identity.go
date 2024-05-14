package identity

type identityService struct {
}

func NewIdentityService() IIdentityService {
	return &identityService{}
}
