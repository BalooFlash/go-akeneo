package client

import (
	"io"
	"net/http"
)

type HttpClient struct {
}

func NewHttpClient() HttpClient {
	return HttpClient{}
}

func (h HttpClient) SendRequest(method, uri string, headers map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
