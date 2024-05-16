package v2

type IGetAccessTokenRequest interface {
	WithClientId(pclientId string) IGetAccessTokenRequest
	WithClientSecret(pclientSecret string) IGetAccessTokenRequest
	GetClientId() string
	GetClientSecret() string
	ToRequestBody() interface{}
}
