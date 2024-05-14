package client

type client struct {
}

func NewClient() IClient {
	c := new(client)
	return c
}
