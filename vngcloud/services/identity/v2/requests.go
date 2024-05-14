package v2

type IGetAccessTokenRequest interface {
	WithClientId(pclientId string) IGetAccessTokenRequest
	WithClientSecret(pclientSecret string) IGetAccessTokenRequest
}

type getAccessTokenRequest struct {
	ClientId     string
	ClientSecret string
}

func NewGetAccessTokenRequest(pclientId, pclientSecret string) IGetAccessTokenRequest {
	return &getAccessTokenRequest{
		ClientId:     pclientId,
		ClientSecret: pclientSecret,
	}
}

func (s *getAccessTokenRequest) WithClientId(pclientId string) IGetAccessTokenRequest {
	s.ClientId = pclientId
	return s
}

func (s *getAccessTokenRequest) WithClientSecret(pclientSecret string) IGetAccessTokenRequest {
	s.ClientSecret = pclientSecret
	return s
}
