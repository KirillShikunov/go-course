package client

type Client struct {
	name string
}

func (c *Client) GetName() string {
	return c.name
}

func NewClient(name string) Client {
	return Client{name}
}
