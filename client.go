package flexmem

import (
	"fmt"

	"github.com/ariefdarmawan/rpchub"

	"github.com/ariefdarmawan/rpchub/hubclient"
)

type Client struct {
	host string

	rc *hubclient.Client
}

func NewClient(host string) (*Client, error) {
	c := new(Client)
	rc, err := hubclient.NewClient(host)
	if err != nil {
		return nil, fmt.Errorf("unable to connect. %s", err.Error())
	}
	c.rc = rc
	return c, nil
}

func (c *Client) Call(name string, parms ...interface{}) *rpchub.Response {
	if c == nil || c.rc == nil {
		return rpchub.NewResponseWithErr("rpc client is not properly initialized")
	}

	return c.rc.Call(name, parms...)
}

func (c *Client) Close() {
	if c.rc != nil {
		c.rc.Close()
	}
}
