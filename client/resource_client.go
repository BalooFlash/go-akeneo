package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-akeneo/go-akeneo/routing"
)

type httpClientInterface interface {
	SendRequest(method, uri string, headers map[string]string, body io.Reader) (*http.Response, error)
}

type ResourceClient struct {
	httpClient   httpClientInterface
	uriGenerator routing.UriGenerator
}

func NewResourceClient(httpClient httpClientInterface, uriGenerator routing.UriGenerator) ResourceClient {
	return ResourceClient{
		httpClient:   httpClient,
		uriGenerator: uriGenerator,
	}
}

func (r ResourceClient) GetResource(uri string, uriParams []string) (*http.Response, error) {
	uri = r.uriGenerator.Generate(uri, uriParams)

	headers := map[string]string{"Accept": "*/*"}

	return r.httpClient.SendRequest(http.MethodGet, uri, headers, nil)
}

func (r ResourceClient) CreateResource(uri string, uriParams []string, body json.RawMessage) (*http.Response, error) {
	uri = r.uriGenerator.Generate(uri, uriParams)

	headers := map[string]string{"Content-Type": "application/json"}

	return r.httpClient.SendRequest(http.MethodPost, uri, headers, bytes.NewReader(body))
}

func (r ResourceClient) DeleteResource(uri string, uriParams []string) error {
	uri = r.uriGenerator.Generate(uri, uriParams)

	_, err := r.httpClient.SendRequest(http.MethodDelete, uri, map[string]string{}, nil)
	if err != nil {
		return err
	}

	return nil
}
