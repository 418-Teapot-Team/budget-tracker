package client

import "github.com/valyala/fasthttp"

const (
	applicationJson = "application/json"
)

type Client struct {
	client  *fasthttp.Client
	request *fasthttp.Request
}

func NewClient() *Client {
	client := &fasthttp.Client{
		DisableHeaderNamesNormalizing: true,
	}

	request := fasthttp.AcquireRequest()
	request.Header.SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")

	return &Client{client: client, request: request}
}
