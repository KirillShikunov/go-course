package client

type Client struct {
	name string
}

func (c *Client) Name() string {
	return c.name
}

func NewClient(name string) Client {
	return Client{name}
}
