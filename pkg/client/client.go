package client

type config struct {
	verbose bool
}

type client struct {
	config *config
}

func (c *client) Info() {

}

func Factory() *client {
	config := &config{}
	client := &client{
		config: config,
	}
	return client
}
