package client

import "net/http"

type Client struct {
	httpClient http.Client
}

func (client *Client) GetData(url string) (*http.Response, error) {
	client.httpClient.Get(url)
	return nil, nil
}
