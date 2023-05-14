package client

import "github.com/valyala/fasthttp"

const (
	applicationJson = "application/json"
)

type Client struct {
	client *fasthttp.Client
}

func NewClient() *Client {
	client := &fasthttp.Client{
		DisableHeaderNamesNormalizing: true,
	}

	return &Client{client: client}
}
