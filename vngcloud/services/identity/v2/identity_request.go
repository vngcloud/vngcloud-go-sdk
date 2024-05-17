package v2

type GetAccessTokenRequest struct {
	ClientId     string
	ClientSecret string

	GrantType string `json:"grant_type"`
}

func NewGetAccessTokenRequest(pclientId, pclientSecret string) IGetAccessTokenRequest {
	return &GetAccessTokenRequest{
		ClientId:     pclientId,
		ClientSecret: pclientSecret,
		GrantType:    "client_credentials",
	}
}

func (s *GetAccessTokenRequest) WithClientId(pclientId string) IGetAccessTokenRequest {
	s.ClientId = pclientId
	return s
}

func (s *GetAccessTokenRequest) WithClientSecret(pclientSecret string) IGetAccessTokenRequest {
	s.ClientSecret = pclientSecret
	return s
}

func (s *GetAccessTokenRequest) GetClientId() string {
	return s.ClientId
}

func (s *GetAccessTokenRequest) GetClientSecret() string {
	return s.ClientSecret
}

func (s *GetAccessTokenRequest) ToRequestBody() interface{} {
	return s
}
