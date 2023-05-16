package client

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

const (
	goverlaUrl = "https://api.goverla.ua/graphql"
)

var goverlaBody = []byte(`{
    "operationName": "Point",
    "variables": {
        "alias": "goverla-ua"
    },
    "query": "query Point($alias: Alias!) {\n  point(alias: $alias) {\n    id\n    rates {\n      id\n      currency {\n        alias\n        name\n        exponent\n        codeAlpha\n        codeNumeric\n        __typename\n      }\n         ask {\n        absolute\n        relative\n        updatedAt\n        __typename\n      }\n      __typename\n    }\n    updatedAt\n    __typename\n  }\n}\n"
}`)

func (app *Client) GetGoverlaRate() (GoverlaOutput, error) {

	var output GoverlaOutput

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	app.request.CopyTo(req)

	req.SetRequestURI(goverlaUrl)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType(applicationJson)
	req.SetBody(goverlaBody)

	err := app.client.Do(req, res)
	defer fasthttp.ReleaseResponse(res)

	if err != nil {
		return output, err
	}

	err = json.Unmarshal(res.Body(), &output)
	if err != nil {
		return output, err
	}
	return output, err
}
