package tokens

type AuthOptionsBuilder interface {
	ToTokenV2BodyMap(interface{}) (interface{}, error)
	ToTokenV2HeadersMap(map[string]interface{}) (map[string]string, error)
	CanReauth() bool
}
