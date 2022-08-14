package client

import (
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

	return r.httpClient.SendRequest("GET", uri, headers, nil)
}
