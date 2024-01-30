package tokens

type AuthOptions struct {
	IdentityEndpoint string
}

func (s *AuthOptions) ToTokenV2BodyMap(interface{}) (interface{}, error) {
	return nil, nil
}

func (s *AuthOptions) ToTokenV2HeadersMap(map[string]interface{}) (map[string]string, error) {
	return nil, nil
}

func (s *AuthOptions) CanReauth() bool {
	return true
}
