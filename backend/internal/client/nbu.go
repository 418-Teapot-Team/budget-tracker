package client

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

const (
	nbuUrl = "https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json"
)

func (app *Client) GetOfficialRate() ([]NbuOfficial, error) {
	var output []NbuOfficial

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.SetRequestURI(nbuUrl)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.SetContentType(applicationJson)

	err := app.client.Do(req, res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res.Body(), &output)
	if err != nil {
		return nil, err
	}

	return output, err
}
