package flexmem

import (
	"fmt"
	"net/rpc"
)

type Client struct {
	host string

	rc *rpc.Client
}

func NewClient(host string) (*Client, error) {
	c := new(Client)

	rc, err := rpc.Dial("tcp", host)
	if err != nil {
		return nil, fmt.Errorf("unable to connect. %s", err.Error())
	}
	c.rc = rc
	return c, nil
}

func (c *Client) Call(name string, parms ...interface{}) *Response {
	res := new(Response)
	if c.rc == nil {
		res.err = fmt.Errorf("rpc client is not properly initialized")
		return res
	}

	req := new(Request)
	req.Name = name
	req.Parm = parms

	if err := c.rc.Call("RpcProxy.Call", req, res); err != nil {
		res.err = fmt.Errorf("%s error: %s", name, err.Error())
		return res
	}

	return res
}

func (c *Client) Close() {
	if c.rc != nil {
		c.rc.Close()
	}
}
